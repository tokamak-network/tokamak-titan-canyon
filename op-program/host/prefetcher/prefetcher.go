package prefetcher

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	preimage "github.com/tokamak-network/tokamak-thanos/op-preimage"
	"github.com/tokamak-network/tokamak-thanos/op-program/client/l1"
	"github.com/tokamak-network/tokamak-thanos/op-program/client/l2"
	"github.com/tokamak-network/tokamak-thanos/op-program/client/mpt"
	"github.com/tokamak-network/tokamak-thanos/op-program/host/kvstore"
	"github.com/tokamak-network/tokamak-thanos/op-service/eth"
)

type L1Source interface {
	InfoByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, error)
	InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error)
	FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, error)
}

type L2Source interface {
	InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error)
	NodeByHash(ctx context.Context, hash common.Hash) ([]byte, error)
	CodeByHash(ctx context.Context, hash common.Hash) ([]byte, error)
	OutputByRoot(ctx context.Context, root common.Hash) (eth.Output, error)
}

type Prefetcher struct {
	logger    log.Logger
	l1Fetcher L1Source
	l2Fetcher L2Source
	lastHint  string
	kvStore   kvstore.KV
}

func NewPrefetcher(logger log.Logger, l1Fetcher L1Source, l2Fetcher L2Source, kvStore kvstore.KV) *Prefetcher {
	return &Prefetcher{
		logger:    logger,
		l1Fetcher: NewRetryingL1Source(logger, l1Fetcher),
		l2Fetcher: NewRetryingL2Source(logger, l2Fetcher),
		kvStore:   kvStore,
	}
}

func (p *Prefetcher) Hint(hint string) error {
	p.logger.Trace("Received hint", "hint", hint)
	p.lastHint = hint
	return nil
}

func (p *Prefetcher) GetPreimage(ctx context.Context, key common.Hash) ([]byte, error) {
	p.logger.Trace("Pre-image requested", "key", key)
	pre, err := p.kvStore.Get(key)
	// Use a loop to keep retrying the prefetch as long as the key is not found
	// This handles the case where the prefetch downloads a preimage, but it is then deleted unexpectedly
	// before we get to read it.
	for errors.Is(err, kvstore.ErrNotFound) && p.lastHint != "" {
		hint := p.lastHint
		if err := p.prefetch(ctx, hint); err != nil {
			return nil, fmt.Errorf("prefetch failed: %w", err)
		}
		pre, err = p.kvStore.Get(key)
		if err != nil {
			p.logger.Error("Fetched pre-images for last hint but did not find required key", "hint", hint, "key", key)
		}
	}
	return pre, err
}

func (p *Prefetcher) prefetch(ctx context.Context, hint string) error {
	hintType, hash, err := parseHint(hint)
	if err != nil {
		return err
	}
	p.logger.Debug("Prefetching", "type", hintType, "hash", hash)
	switch hintType {
	case l1.HintL1BlockHeader:
		header, err := p.l1Fetcher.InfoByHash(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L1 block %s header: %w", hash, err)
		}
		data, err := header.HeaderRLP()
		if err != nil {
			return fmt.Errorf("marshall header: %w", err)
		}
		return p.kvStore.Put(preimage.Keccak256Key(hash).PreimageKey(), data)
	case l1.HintL1Transactions:
		_, txs, err := p.l1Fetcher.InfoAndTxsByHash(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L1 block %s txs: %w", hash, err)
		}
		return p.storeTransactions(txs)
	case l1.HintL1Receipts:
		_, receipts, err := p.l1Fetcher.FetchReceipts(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L1 block %s receipts: %w", hash, err)
		}
		return p.storeReceipts(receipts)
	case l2.HintL2BlockHeader, l2.HintL2Transactions:
		header, txs, err := p.l2Fetcher.InfoAndTxsByHash(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L2 block %s: %w", hash, err)
		}
		data, err := header.HeaderRLP()
		if err != nil {
			return fmt.Errorf("failed to encode header to RLP: %w", err)
		}
		err = p.kvStore.Put(preimage.Keccak256Key(hash).PreimageKey(), data)
		if err != nil {
			return err
		}
		return p.storeTransactions(txs)
	case l2.HintL2StateNode:
		node, err := p.l2Fetcher.NodeByHash(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L2 state node %s: %w", hash, err)
		}
		return p.kvStore.Put(preimage.Keccak256Key(hash).PreimageKey(), node)
	case l2.HintL2Code:
		code, err := p.l2Fetcher.CodeByHash(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L2 contract code %s: %w", hash, err)
		}
		return p.kvStore.Put(preimage.Keccak256Key(hash).PreimageKey(), code)
	case l2.HintL2Output:
		output, err := p.l2Fetcher.OutputByRoot(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to fetch L2 output root %s: %w", hash, err)
		}
		return p.kvStore.Put(preimage.Keccak256Key(hash).PreimageKey(), output.Marshal())
	}
	return fmt.Errorf("unknown hint type: %v", hintType)
}

func (p *Prefetcher) storeReceipts(receipts types.Receipts) error {
	opaqueReceipts, err := eth.EncodeReceipts(receipts)
	if err != nil {
		return err
	}
	return p.storeTrieNodes(opaqueReceipts)
}

func (p *Prefetcher) storeTransactions(txs types.Transactions) error {
	opaqueTxs, err := eth.EncodeTransactions(txs)
	if err != nil {
		return err
	}
	return p.storeTrieNodes(opaqueTxs)
}

func (p *Prefetcher) storeTrieNodes(values []hexutil.Bytes) error {
	_, nodes := mpt.WriteTrie(values)
	for _, node := range nodes {
		key := preimage.Keccak256Key(crypto.Keccak256Hash(node)).PreimageKey()
		if err := p.kvStore.Put(key, node); err != nil {
			return fmt.Errorf("failed to store node: %w", err)
		}
	}
	return nil
}

// parseHint parses a hint string in wire protocol. Returns the hint type, requested hash and error (if any).
func parseHint(hint string) (string, common.Hash, error) {
	hintType, hashStr, found := strings.Cut(hint, " ")
	if !found {
		return "", common.Hash{}, fmt.Errorf("unsupported hint: %s", hint)
	}
	hash := common.HexToHash(hashStr)
	if hash == (common.Hash{}) {
		return "", common.Hash{}, fmt.Errorf("invalid hash: %s", hashStr)
	}
	return hintType, hash, nil
}

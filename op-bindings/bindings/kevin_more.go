// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const KevinStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2-tokamak/Kevin.sol:Kevin\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/L2-tokamak/Kevin.sol:Kevin\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"src/L2-tokamak/Kevin.sol:Kevin\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L2-tokamak/Kevin.sol:Kevin\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"src/L2-tokamak/Kevin.sol:Kevin\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":1005,\"contract\":\"src/L2-tokamak/Kevin.sol:Kevin\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var KevinStorageLayout = new(solc.StorageLayout)

var KevinDeployedBin = "0x608060405234801561001057600080fd5b506004361061011b5760003560e01c806370a08231116100b257806395d89b4111610081578063a9059cbb11610066578063a9059cbb14610264578063dd62ed3e14610277578063f2fde38b146102bd57600080fd5b806395d89b4114610249578063a457c2d71461025157600080fd5b806370a08231146101d0578063715018a61461020657806379cc67901461020e5780638da5cb5b1461022157600080fd5b8063313ce567116100ee578063313ce56714610186578063395093511461019557806340c10f19146101a857806342966c68146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806318160ddd1461016157806323b872dd14610173575b600080fd5b6101286102d0565b6040516101359190610f18565b60405180910390f35b61015161014c366004610fb4565b610362565b6040519015158152602001610135565b6002545b604051908152602001610135565b610151610181366004610fde565b61037a565b60405160128152602001610135565b6101516101a3366004610fb4565b61039e565b6101bb6101b6366004610fb4565b6103ea565b005b6101bb6101cb36600461101a565b610400565b6101656101de366004611033565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101bb61040d565b6101bb61021c366004610fb4565b610421565b60055460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610128610436565b61015161025f366004610fb4565b610445565b610151610272366004610fb4565b61051b565b610165610285366004611055565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101bb6102cb366004611033565b610529565b6060600380546102df90611088565b80601f016020809104026020016040519081016040528092919081815260200182805461030b90611088565b80156103585780601f1061032d57610100808354040283529160200191610358565b820191906000526020600020905b81548152906001019060200180831161033b57829003601f168201915b5050505050905090565b6000336103708185856105dd565b5060019392505050565b600033610388858285610791565b610393858585610868565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061037090829086906103e590879061110a565b6105dd565b6103f2610b1b565b6103fc8282610b9c565b5050565b61040a3382610cbc565b50565b610415610b1b565b61041f6000610ea1565b565b61042c823383610791565b6103fc8282610cbc565b6060600480546102df90611088565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561050e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61039382868684036105dd565b600033610370818585610868565b610531610b1b565b73ffffffffffffffffffffffffffffffffffffffff81166105d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610505565b61040a81610ea1565b73ffffffffffffffffffffffffffffffffffffffff831661067f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff8216610722576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146108625781811015610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610505565b61086284848484036105dd565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661090b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff82166109ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610a64576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610aa890849061110a565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b0e91815260200190565b60405180910390a3610862565b60055473ffffffffffffffffffffffffffffffffffffffff16331461041f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610505565b73ffffffffffffffffffffffffffffffffffffffff8216610c19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610505565b8060026000828254610c2b919061110a565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610c6590849061110a565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8216610d5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015610e15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610505565b73ffffffffffffffffffffffffffffffffffffffff83166000908152602081905260408120838303905560028054849290610e51908490611122565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610784565b6005805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600060208083528351808285015260005b81811015610f4557858101830151858201604001528201610f29565b81811115610f57576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610faf57600080fd5b919050565b60008060408385031215610fc757600080fd5b610fd083610f8b565b946020939093013593505050565b600080600060608486031215610ff357600080fd5b610ffc84610f8b565b925061100a60208501610f8b565b9150604084013590509250925092565b60006020828403121561102c57600080fd5b5035919050565b60006020828403121561104557600080fd5b61104e82610f8b565b9392505050565b6000806040838503121561106857600080fd5b61107183610f8b565b915061107f60208401610f8b565b90509250929050565b600181811c9082168061109c57607f821691505b6020821081036110d5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561111d5761111d6110db565b500190565b600082821015611134576111346110db565b50039056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(KevinStorageLayoutJSON), KevinStorageLayout); err != nil {
		panic(err)
	}

	layouts["Kevin"] = KevinStorageLayout
	deployedBytecodes["Kevin"] = KevinDeployedBin
}

package routes

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/tokamak-network/tokamak-thanos/indexer/database"
)

// Routes ... Route handler struct
type Routes struct {
	logger log.Logger
	view   database.BridgeTransfersView
	router *chi.Mux
	v      *Validator
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(logger log.Logger, bv database.BridgeTransfersView, r *chi.Mux) Routes {
	return Routes{
		logger: logger,
		view:   bv,
		router: r,
	}
}

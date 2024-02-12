//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package restful

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// Dialer is an interface for the client to dial the server.
type Dialer interface {
	GetStockBySymbol(ctx contextx.Contextx, symbol string) (agg.Stock, error)
}

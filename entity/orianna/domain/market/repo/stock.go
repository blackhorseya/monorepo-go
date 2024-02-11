//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IStockRepo is the interface for stock repository.
type IStockRepo interface {
	List(ctx contextx.Contextx) ([]agg.Stock, error)
	Add(ctx contextx.Contextx, stock agg.Stock) error
	BulkUpsertInfo(ctx contextx.Contextx, stocks []agg.Stock) error
}

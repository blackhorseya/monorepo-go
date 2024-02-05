//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// Storager is the interface that wraps the basic storage methods.
type Storager interface {
	GetBySymbol(ctx contextx.Contextx, symbol string) (info *model.StockInfo, err error)
}
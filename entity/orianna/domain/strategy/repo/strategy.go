//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/strategy/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IStrategyRepo is the interface for strategy repository.
type IStrategyRepo interface {
	Get(ctx contextx.Contextx, id string) (agg.Strategy, error)
	Create(ctx contextx.Contextx, data agg.Strategy) (string, error)
}

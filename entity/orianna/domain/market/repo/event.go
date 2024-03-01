//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListOptions is the options for list.
type ListOptions struct {
	Limit  int
	Offset int
}

// IEventRepo is a interface for event repository.
type IEventRepo interface {
	GetByID(ctx contextx.Contextx, id string) (item agg.Event, err error)
	List(ctx contextx.Contextx, opts ListOptions) (items []agg.Event, total int, err error)
	Create(ctx contextx.Contextx, item agg.Event) (err error)
}

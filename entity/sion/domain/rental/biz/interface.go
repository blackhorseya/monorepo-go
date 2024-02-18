//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/agg"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListByLocationOptions is the options for ListByLocation.
type ListByLocationOptions struct {
	Page int
	Size int
}

// IRentalBiz is an interface for rental biz.
type IRentalBiz interface {
	ListByLocation(
		ctx contextx.Contextx,
		location *model.Location,
		opts ListByLocationOptions,
	) (rentals []*agg.Asset, total int, err error)
}

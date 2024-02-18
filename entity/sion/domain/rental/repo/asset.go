//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListOptions is the options for list.
type ListOptions struct {
	Offset int
	Limit  int
}

// IAssetRepo is an interface for asset repository.
type IAssetRepo interface {
	List(ctx contextx.Contextx, opts ListOptions) ([]*agg.Asset, error)
}

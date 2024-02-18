//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IAssetRepo is an interface for asset repository.
type IAssetRepo interface {
	FetchAvailableCars(ctx contextx.Contextx) ([]*agg.Asset, error)
}

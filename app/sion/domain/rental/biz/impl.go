package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/agg"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/biz"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/model"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	assets repo.IAssetRepo
}

// NewRentalBiz will create a new rental biz.
func NewRentalBiz(assets repo.IAssetRepo) (biz.IRentalBiz, error) {
	return &impl{assets: assets}, nil
}

func (i *impl) ListByLocation(
	ctx contextx.Contextx,
	location *model.Location,
	opts biz.ListByLocationOptions,
) (rentals []*agg.Asset, total int, err error) {
	// TODO implement me
	panic("implement me")
}

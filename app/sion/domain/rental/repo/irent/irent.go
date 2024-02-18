package irent

import (
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/agg"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/repo"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	endpoint string
	version  string
}

// NewAssetRepo is a function to create a new asset repository.
func NewAssetRepo() (repo.IAssetRepo, error) {
	return &impl{
		endpoint: configx.C.IRent.HTTP.URL,
		version:  configx.C.IRent.Version,
	}, nil
}

func (i *impl) FetchAvailableCars(ctx contextx.Contextx) ([]*agg.Asset, error) {
	// TODO implement me
	panic("implement me")
}

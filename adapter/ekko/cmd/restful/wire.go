//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/ekko/domain/workflow/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		mongodbx.NewClient,
		linebot.NewClient,
		httpx.NewServer,
		biz.ProviderSet,
		newService,
	))
}

func NewRestful() (adapterx.Restful, error) {
	panic(wire.Build(
		mongodbx.NewClient,
		linebot.NewClient,
		httpx.NewServer,
		biz.ProviderSet,
		newRestful,
	))
}

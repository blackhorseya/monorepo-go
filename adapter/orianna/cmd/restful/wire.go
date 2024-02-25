//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/orianna/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		linebot.NewClient,
		mongodb.NewClient,
		biz.ProviderSet,
		newService,
	))
}

func NewRestful() (adapterx.Restful, error) {
	panic(wire.Build(
		linebot.NewClient,
		mongodb.NewClient,
		biz.ProviderSet,
		newRestful,
	))
}

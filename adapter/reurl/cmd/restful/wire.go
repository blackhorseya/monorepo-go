//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	"github.com/blackhorseya/monorepo-go/pkg/storage/redis"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		linebot.NewClient,
		redis.NewClient,
		biz.ShortenRedis,
		newService,
	))
}

func NewRestful(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		linebot.NewClient,
		redis.NewClient,
		biz.ShortenRedis,
		newRestful,
	))
}

//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/storage/influxdb"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		finmindx.NewClient,
		mongodb.NewClient,
		influxdb.NewClient,
		biz.ProviderSet,
		newRestful,
	))
}

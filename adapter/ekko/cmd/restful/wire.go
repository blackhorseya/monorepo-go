//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/ekko/domain/workflow/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		mongodbx.NewClient,
		linebot.NewClient,
		biz.ProviderSet,
		newRestful,
	))
}

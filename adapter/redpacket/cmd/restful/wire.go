//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/internal/pkg/logx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	configx.NewWithViper,
	logx.NewWithConfig,

	newImpl,
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

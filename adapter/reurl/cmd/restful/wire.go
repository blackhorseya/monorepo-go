//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		newService,
	))
}

func NewRestful(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		newRestful,
	))
}

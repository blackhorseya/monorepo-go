//go:build wireinject

//go:generate wire

package cronjob

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	newImpl,
)

// New will create a new restful adapter instance
func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

var testProviderSet = wire.NewSet(
	newImpl,
)

// NewExternal will create a new restful adapter instance for external test.
func NewExternal(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(testProviderSet))
}

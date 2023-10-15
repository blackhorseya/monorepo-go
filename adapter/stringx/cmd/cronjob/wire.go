//go:build wireinject

//go:generate wire

package cronjob

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/internal/pkg/logx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	configx.NewExample,
	logx.NewExample,
	newImpl,
)

// New will create a new restful adapter instance
func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

var testProviderSet = wire.NewSet(
	configx.NewExample,
	logx.NewExample,
	newImpl,
)

// NewExternal will create a new restful adapter instance for external test.
func NewExternal(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(testProviderSet))
}

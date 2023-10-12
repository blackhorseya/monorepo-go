//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var providerSet = wire.NewSet(newImpl)

// New will create a new restful adapter instance
func New(viper *viper.Viper, logger *zap.Logger) (adapterx.Restful, error) {
	panic(wire.Build(providerSet))
}

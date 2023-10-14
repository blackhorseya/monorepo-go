//go:build wireinject

//go:generate wire

package grpc

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var providerSet = wire.NewSet(newImpl)

// New will create a new restful adapter instance
func New(v *viper.Viper, logger *zap.Logger) adapterx.Servicer {
	panic(wire.Build(providerSet))
}

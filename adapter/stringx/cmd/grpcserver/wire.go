//go:build wireinject

//go:generate wire

package grpcserver

import (
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var providerSet = wire.NewSet(
	configx.NewExample,
	biz.New,
	newImpl,
)

// New will create a new restful adapter instance
func New(v *viper.Viper, logger *zap.Logger) adapterx.Servicer {
	panic(wire.Build(providerSet))
}

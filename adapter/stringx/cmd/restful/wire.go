//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var providerSet = wire.NewSet(
	configx.NewWithViper,
	newRouter,
	biz.New,
	newImpl,
)

// New will create a new restful adapter instance
func New(v *viper.Viper, logger *zap.Logger) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

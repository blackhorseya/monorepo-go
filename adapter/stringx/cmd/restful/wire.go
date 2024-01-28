//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	biz.New,
	newImpl,
)

// New will create a new restful adapter instance
func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

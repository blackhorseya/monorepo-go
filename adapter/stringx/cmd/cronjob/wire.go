//go:build wireinject

//go:generate wire

package cronjob

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/logx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(newImpl, logx.NewExample)

// New will create a new restful adapter instance
func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

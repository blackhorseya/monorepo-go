//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	newService,
)

func NewService() (adapterx.Servicer, error) {
	panic(wire.Build(providerSet))
}

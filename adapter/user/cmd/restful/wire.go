//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(newImpl)

// New will create a new restful adapter instance
func New() (adapterx.Restful, error) {
	panic(wire.Build(providerSet))
}

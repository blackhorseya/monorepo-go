//go:build wireinject

//go:generate wire

package cronjob

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(newImpl)

// New will create a new grpc adapter instance
func New() (adapterx.Cronjob, error) {
	panic(wire.Build(providerSet))
}

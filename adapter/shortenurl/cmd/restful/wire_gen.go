// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewService(v *viper.Viper) (adapterx.Servicer, error) {
	servicer, err := newService()
	if err != nil {
		return nil, err
	}
	return servicer, nil
}

// wire.go:

var providerSet = wire.NewSet(
	newService,
)

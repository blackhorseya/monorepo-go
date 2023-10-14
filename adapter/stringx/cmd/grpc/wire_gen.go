// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// New will create a new restful adapter instance
func New(v *viper.Viper, logger *zap.Logger) adapterx.Servicer {
	servicer := newImpl(v, logger)
	return servicer
}

// wire.go:

var providerSet = wire.NewSet(newImpl)

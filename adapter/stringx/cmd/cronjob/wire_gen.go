// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cronjob

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/internal/pkg/logx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

// New will create a new restful adapter instance
func New(v *viper.Viper) (adapterx.Servicer, error) {
	config := configx.NewExample()
	logger := logx.NewExample()
	servicer := newImpl(v, config, logger)
	return servicer, nil
}

// NewExternal will create a new restful adapter instance for external test.
func NewExternal(v *viper.Viper) (adapterx.Servicer, error) {
	config := configx.NewExample()
	logger := logx.NewExample()
	servicer := newImpl(v, config, logger)
	return servicer, nil
}

// wire.go:

var providerSet = wire.NewSet(configx.NewExample, logx.NewExample, newImpl)

var testProviderSet = wire.NewSet(configx.NewExample, logx.NewExample, newImpl)

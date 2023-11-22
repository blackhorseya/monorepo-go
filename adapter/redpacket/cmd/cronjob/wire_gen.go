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

func New(v *viper.Viper) (adapterx.Servicer, error) {
	config, err := configx.NewWithViper(v)
	if err != nil {
		return nil, err
	}
	logger, err := logx.NewWithConfig(config)
	if err != nil {
		return nil, err
	}
	servicer, err := newCronjob(v, config, logger)
	if err != nil {
		return nil, err
	}
	return servicer, nil
}

// wire.go:

var providerSet = wire.NewSet(configx.NewWithViper, logx.NewWithConfig, newCronjob)

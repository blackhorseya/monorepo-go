// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/monorepo-go/adapter/orianna/api/docs"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Servicer, error) {
	client, err := linebot.NewClient()
	if err != nil {
		return nil, err
	}
	servicer, err := newRestful(client)
	if err != nil {
		return nil, err
	}
	return servicer, nil
}

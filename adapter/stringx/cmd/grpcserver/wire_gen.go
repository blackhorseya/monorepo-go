// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpcserver

import (
	"github.com/blackhorseya/monorepo-go/app/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

// New will create a new restful adapter instance
func New(v *viper.Viper) (adapterx.Servicer, error) {
	iStringBiz := biz.New()
	servicer := newImpl(v, iStringBiz)
	return servicer, nil
}

// NewExternal will create a new restful adapter instance for external test.
func NewExternal(v *viper.Viper) (adapterx.Servicer, error) {
	iStringBiz := biz.New()
	servicer := newImpl(v, iStringBiz)
	return servicer, nil
}

// wire.go:

var providerSet = wire.NewSet(biz.New, newImpl)

var testProviderSet = wire.NewSet(
	newImpl, biz.New,
)
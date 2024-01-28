// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo/memory"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/monorepo-go/adapter/shortenurl/api/docs"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Servicer, error) {
	application := initAPP()
	storager := memory.NewStorager()
	iShorteningBiz := biz.NewShortening(storager)
	servicer, err := newRestful(v, application, iShorteningBiz)
	if err != nil {
		return nil, err
	}
	return servicer, nil
}

// wire.go:

var providerSet = wire.NewSet(biz.ProviderSet, initAPP,
	newRestful,
)

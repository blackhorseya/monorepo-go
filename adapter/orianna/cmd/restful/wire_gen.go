// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/monorepo-go/app/orianna/domain/market/biz"
	mongodb3 "github.com/blackhorseya/monorepo-go/app/orianna/domain/market/repo/event/mongodb"
	mongodb2 "github.com/blackhorseya/monorepo-go/app/orianna/domain/market/repo/stock/mongodb"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/monorepo-go/adapter/orianna/api/docs"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Servicer, error) {
	server, err := httpx.NewServer()
	if err != nil {
		return nil, err
	}
	client, err := mongodb.NewClient()
	if err != nil {
		return nil, err
	}
	iStockRepo, err := mongodb2.NewStockRepo(client)
	if err != nil {
		return nil, err
	}
	iEventRepo, err := mongodb3.NewEventRepo(client)
	if err != nil {
		return nil, err
	}
	iMarketBiz, err := biz.NewMarketBiz(iStockRepo, iEventRepo)
	if err != nil {
		return nil, err
	}
	linebotClient, err := linebot.NewClient()
	if err != nil {
		return nil, err
	}
	servicer := newService(server, iMarketBiz, linebotClient)
	return servicer, nil
}

func NewRestful() (adapterx.Restful, error) {
	server, err := httpx.NewServer()
	if err != nil {
		return nil, err
	}
	client, err := mongodb.NewClient()
	if err != nil {
		return nil, err
	}
	iStockRepo, err := mongodb2.NewStockRepo(client)
	if err != nil {
		return nil, err
	}
	iEventRepo, err := mongodb3.NewEventRepo(client)
	if err != nil {
		return nil, err
	}
	iMarketBiz, err := biz.NewMarketBiz(iStockRepo, iEventRepo)
	if err != nil {
		return nil, err
	}
	linebotClient, err := linebot.NewClient()
	if err != nil {
		return nil, err
	}
	restful := newRestful(server, iMarketBiz, linebotClient)
	return restful, nil
}

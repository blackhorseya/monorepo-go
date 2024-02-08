// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bot

import (
	"github.com/blackhorseya/monorepo-go/app/domain/market/biz"
	influxdb2 "github.com/blackhorseya/monorepo-go/app/domain/market/repo/influxdb"
	mongodb2 "github.com/blackhorseya/monorepo-go/app/domain/market/repo/mongodb"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/linebot"
	"github.com/blackhorseya/monorepo-go/pkg/storage/influxdb"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
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
	dialer, err := finmindx.NewClient()
	if err != nil {
		return nil, err
	}
	mongoClient, err := mongodb.NewClient()
	if err != nil {
		return nil, err
	}
	storager, err := mongodb2.NewStorager(mongoClient)
	if err != nil {
		return nil, err
	}
	influxdb3Client, err := influxdb.NewClient()
	if err != nil {
		return nil, err
	}
	iQuoteRepo, err := influxdb2.NewQuoteRepo(influxdb3Client)
	if err != nil {
		return nil, err
	}
	iMarketBiz, err := biz.NewMarketBiz(dialer, storager, iQuoteRepo)
	if err != nil {
		return nil, err
	}
	servicer, err := newRestful(client, iMarketBiz)
	if err != nil {
		return nil, err
	}
	return servicer, nil
}

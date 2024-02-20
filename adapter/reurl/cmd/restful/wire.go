//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/biz"
	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	linebotx "github.com/blackhorseya/monorepo-go/pkg/linebot"
	"github.com/blackhorseya/monorepo-go/pkg/storage/redis"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// Injector is the injector for the restful service.
type Injector struct {
	server *httpx.Server
	bot    *linebot.Client
	svc    shortB.IShorteningBiz
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		httpx.NewServer,
		redis.NewClient,
		linebotx.NewClient,
		biz.ShortenRedis,
	))
}

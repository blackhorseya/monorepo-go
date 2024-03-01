//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/monorepo-go/app/orianna/domain/market/repo/stock/mongodb"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/notify"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injector is the injector for main.
type Injector struct {
	finmind  finmindx.Dialer
	rw       *mongo.Client
	notifier notify.Notifier
	repo     repo.IStockRepo
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		finmindx.NewClient,
		mongodbx.NewClient,
		notify.NewLineNotifier,
		mongodb.NewStockRepo,
	))
}

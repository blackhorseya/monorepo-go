//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/monorepo-go/app/orianna/domain/market/repo/mongodb"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/pkg/notify"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
)

// Injector is the injector for main.
type Injector struct {
	notifier notify.Notifier
	repo     repo.IStockRepo
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		notify.NewLineNotifier,
		mongodbx.NewClient,
		mongodb.NewStockRepo,
	))
}
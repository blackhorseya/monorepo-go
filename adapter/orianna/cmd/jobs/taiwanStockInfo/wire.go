//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/notify"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injector is the injector for main.
type Injector struct {
	finmind  finmindx.Dialer
	rw       *mongo.Client
	notifier notify.Notifier
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		finmindx.NewClient,
		mongodb.NewClient,
		notify.NewLineNotifier,
	))
}

//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/notify"
	"github.com/google/wire"
)

// Injector is the injector for main.
type Injector struct {
	notifier notify.Notifier
	finmind  finmindx.Dialer
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		finmindx.NewClient,
		notify.NewLineNotifier,
	))
}

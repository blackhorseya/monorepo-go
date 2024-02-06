//go:build wireinject

//go:generate wire

package main

import (
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/monorepo-go/pkg/notify"
	"github.com/blackhorseya/monorepo-go/pkg/storage/influxdb"
	"github.com/google/wire"
)

// Injector is the injector for main.
type Injector struct {
	notifier notify.Notifier
	client   *influxdb3.Client
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		notify.NewLineNotifier,
		influxdb.NewClient,
	))
}

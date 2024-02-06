// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/notify"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, error) {
	notifier, err := notify.NewLineNotifier()
	if err != nil {
		return nil, err
	}
	dialer, err := finmindx.NewClient()
	if err != nil {
		return nil, err
	}
	mainInjector := &Injector{
		notifier: notifier,
		finmind:  dialer,
	}
	return mainInjector, nil
}

// wire.go:

// Injector is the injector for main.
type Injector struct {
	notifier notify.Notifier
	finmind  finmindx.Dialer
}

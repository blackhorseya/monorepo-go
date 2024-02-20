//go:build wireinject

//go:generate wire

package main

import (
	"github.com/google/wire"
)

// Injector is the wire injector for restful.
type Injector struct {
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
	))
}

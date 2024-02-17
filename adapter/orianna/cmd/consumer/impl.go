package consumer

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
)

type impl struct {
}

func newConsumer() (adapterx.Servicer, error) {
	return &impl{}, nil
}

func (i *impl) Start() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// TODO implement me
	panic("implement me")
}

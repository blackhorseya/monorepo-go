package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
)

type impl struct {
}

func (i *impl) Start() error {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

func newService() (adapterx.Servicer, error) {
	return &impl{}, nil
}

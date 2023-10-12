package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
)

type impl struct {
}

func createRestful() adapterx.Restful {
	return &impl{}
}

func (i *impl) Start() error {
	// todo: 2023/10/12|sean|impl me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// todo: 2023/10/12|sean|impl me
	panic("implement me")
}

func (i *impl) InitRouting() error {
	// todo: 2023/10/12|sean|impl me
	panic("implement me")
}

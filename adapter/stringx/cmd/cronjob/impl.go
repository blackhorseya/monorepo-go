package cronjob

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
)

type impl struct {
}

func newImpl() adapterx.Servicer {
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

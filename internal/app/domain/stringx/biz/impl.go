package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
}

// New create a new string biz.
func New() biz.IStringBiz {
	return &impl{}
}

func (i *impl) Uppercase(ctx contextx.Contextx, value string) (upper string, err error) {
	// todo: 2023/10/13|sean|impl me
	panic("implement me")
}

func (i *impl) Count(ctx contextx.Contextx, value string) int {
	// todo: 2023/10/13|sean|impl me
	panic("implement me")
}

package biz

import (
	"errors"
	"strings"

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
	if value == "" {
		return "", errors.New("value is empty")
	}

	return strings.ToUpper(value), nil
}

func (i *impl) Count(ctx contextx.Contextx, value string) int {
	return len(value)
}

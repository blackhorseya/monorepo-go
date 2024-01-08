package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
}

// New is used to create a new shortening biz instance.
func New() biz.IShorteningBiz {
	return &impl{}
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

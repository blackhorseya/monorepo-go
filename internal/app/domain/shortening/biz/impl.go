package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
}

// NewShortening is used to create a new shortening biz instance.
func NewShortening() biz.IShorteningBiz {
	return &impl{}
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

func (i *impl) CreateShortenedURL(ctx contextx.Contextx, originalURL string) (record *model.ShortenedUrl, err error) {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

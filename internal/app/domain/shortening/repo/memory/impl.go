package memory

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	mapper map[string]*model.ShortenedUrl
}

// NewStorager is used to create a new shortening storage instance.
func NewStorager() repo.Storager {
	return &impl{
		mapper: make(map[string]*model.ShortenedUrl),
	}
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

func (i *impl) CreateURLRecord(ctx contextx.Contextx, record *model.ShortenedUrl) error {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

package redis

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/storage/redis"
)

type impl struct {
	client *redis.Client
}

// NewStorager is used to create a new redis storager instance.
func NewStorager(client *redis.Client) (repo.Storager, error) {
	return &impl{
		client: client,
	}, nil
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	// todo: 2024/1/28|sean|implement me
	panic("implement me")
}

func (i *impl) CreateURLRecord(ctx contextx.Contextx, record *model.ShortenedUrl) error {
	// todo: 2024/1/28|sean|implement me
	panic("implement me")
}

package redis

import (
	"encoding/json"

	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/storage/redis"
)

type impl struct {
	rw *redis.Client
}

// NewStorager is used to create a new redis storager instance.
func NewStorager(rw *redis.Client) (repo.Storager, error) {
	return &impl{
		rw: rw,
	}, nil
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	result, err := i.rw.Get(ctx, shortURL).Result()
	if err != nil {
		return nil, err
	}

	var ret *model.ShortenedUrl
	err = json.Unmarshal([]byte(result), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) CreateURLRecord(ctx contextx.Contextx, record *model.ShortenedUrl) error {
	marshal, err := json.Marshal(record)
	if err != nil {
		return err
	}

	err = i.rw.Set(ctx, record.ShortUrl, string(marshal), 0).Err()
	if err != nil {
		return err
	}

	return nil
}

package memory

import (
	"errors"
	"sync"

	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	mu     *sync.RWMutex
	mapper map[string]*model.ShortenedUrl
}

// NewStorager is used to create a new shortening storage instance.
func NewStorager() repo.Storager {
	return &impl{
		mu:     new(sync.RWMutex),
		mapper: make(map[string]*model.ShortenedUrl),
	}
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	record, exists := i.mapper[shortURL]
	if !exists {
		return nil, errors.New("short url not exists")
	}

	return record, nil
}

func (i *impl) CreateURLRecord(ctx contextx.Contextx, record *model.ShortenedUrl) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, exists := i.mapper[record.ShortUrl]
	if exists {
		return errors.New("short url already exists")
	}

	i.mapper[record.ShortUrl] = record

	return nil
}

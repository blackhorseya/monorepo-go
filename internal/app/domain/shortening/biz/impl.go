package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type impl struct {
	storage repo.Storager
}

// NewShortening is used to create a new shortening biz instance.
func NewShortening(storage repo.Storager) biz.IShorteningBiz {
	return &impl{
		storage: storage,
	}
}

func (i *impl) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error) {
	// todo: 2024/1/9|sean|implement me
	panic("implement me")
}

func (i *impl) CreateShortenedURL(ctx contextx.Contextx, originalURL string) (record *model.ShortenedUrl, err error) {
	now := timestamppb.Now()
	ret := &model.ShortenedUrl{
		Id:          0,
		OriginalUrl: originalURL,
		ShortUrl:    originalURL,
		CreatedAt:   now,
		ExpiredAt:   nil,
	}
	err = i.storage.CreateURLRecord(ctx, ret)
	if err != nil {
		ctx.Error("create url record error", zap.Error(err), zap.Any("record", &ret))
		return nil, err
	}

	return ret, nil
}

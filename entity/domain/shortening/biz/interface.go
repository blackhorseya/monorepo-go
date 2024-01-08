//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IShorteningBiz is a shortening biz interface
type IShorteningBiz interface {
	// GetUrlRecordByShortURL is used to get url record by short url
	GetUrlRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error)
}

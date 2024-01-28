//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// Storager is a shortening storage interface.
type Storager interface {
	// GetURLRecordByShortURL is used to get url record by short url
	GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (record *model.ShortenedUrl, err error)

	// CreateURLRecord is used to create url record
	CreateURLRecord(ctx contextx.Contextx, record *model.ShortenedUrl) error
}

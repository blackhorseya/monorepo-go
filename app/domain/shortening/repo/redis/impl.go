package redis

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo"
)

type impl struct {
}

// NewStorager is used to create a new redis storager instance.
func NewStorager() (repo.Storager, error) {
	panic("implement me")
}

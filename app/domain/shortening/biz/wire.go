package biz

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo/memory"
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo/redis"
	"github.com/google/wire"
)

// ProviderSet is used to provide biz set.
var ProviderSet = wire.NewSet(
	NewShortening,
	memory.NewStorager,
)

// ShortenRedis is used to provide biz set with redis storager.
var ShortenRedis = wire.NewSet(
	NewShortening,
	redis.NewStorager,
)

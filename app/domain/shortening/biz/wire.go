package biz

import (
	"github.com/blackhorseya/monorepo-go/app/domain/shortening/repo/memory"
	"github.com/google/wire"
)

// ProviderSet is used to provide biz set.
var ProviderSet = wire.NewSet(
	NewShortening,
	memory.NewStorager,
)

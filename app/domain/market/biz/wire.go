package biz

import (
	"github.com/blackhorseya/monorepo-go/app/domain/market/repo/mongodb"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewMarketBiz,
	mongodb.NewStorager,
)

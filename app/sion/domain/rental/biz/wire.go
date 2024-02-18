package biz

import (
	"github.com/blackhorseya/monorepo-go/app/sion/domain/rental/repo/irent"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewRentalBiz,
	irent.NewAssetRepo,
)

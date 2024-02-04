package biz

import (
	"github.com/google/wire"
)

// ProvideDefaultSet is used to provide a default set for issue biz.
var ProvideDefaultSet = wire.NewSet(
	NewIssueBiz,
)

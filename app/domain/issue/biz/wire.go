package biz

import (
	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo/mongodb"
	"github.com/google/wire"
)

// ProvideDefaultSet is used to provide a default set for issue biz.
var ProvideDefaultSet = wire.NewSet(
	NewIssueBiz,
	mongodb.NewStorager,
)

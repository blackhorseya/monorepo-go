package biz

import (
	"github.com/blackhorseya/monorepo-go/app/ekko/domain/workflow/repo/mongodb"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewWorkflowBiz,
	mongodb.NewIssueRepoWithMongoDB,
)

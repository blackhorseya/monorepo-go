//go:build wireinject

//go:generate wire

package biz

import (
	"github.com/blackhorseya/monorepo-go/app/ekko/domain/workflow/repo/issue/mongodb"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/biz"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewWorkflowBiz,
	mongodb.NewIssueRepo,
)

func NewForExternal() (biz.IWorkflowBiz, error) {
	panic(wire.Build(
		mongodbx.NewClient,
		mongodb.NewIssueRepo,
		NewWorkflowBiz,
	))
}

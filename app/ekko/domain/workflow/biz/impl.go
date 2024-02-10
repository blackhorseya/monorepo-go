package biz

import (
	idM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/identity/model"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/biz"
	wfM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	issueR "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	issue issueR.IIssueRepo
}

// NewWorkflowBiz creates a new WorkflowBiz.
func NewWorkflowBiz(issue issueR.IIssueRepo) (biz.IWorkflowBiz, error) {
	return &impl{issue: issue}, nil
}

func (i *impl) CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (todo *wfM.Ticket, err error) {
	// todo: 2024/2/11|sean|implement me
	panic("implement me")
}

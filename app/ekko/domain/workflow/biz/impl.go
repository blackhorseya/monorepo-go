package biz

import (
	idM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/identity/model"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/biz"
	wfM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	issueR "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	issues issueR.IIssueRepo
}

// NewWorkflowBiz creates a new WorkflowBiz.
func NewWorkflowBiz(issues issueR.IIssueRepo) (biz.IWorkflowBiz, error) {
	return &impl{issues: issues}, nil
}

func (i *impl) CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (todo *wfM.Ticket, err error) {
	issue, err := agg.NewIssue(who.ID, title)
	if err != nil {
		return nil, err
	}

	id, err := i.issues.Create(ctx, issue)
	if err != nil {
		return nil, err
	}

	return &wfM.Ticket{
		ID:        id,
		Title:     issue.GetTitle(),
		Completed: issue.GetCompleted(),
	}, nil
}

func (i *impl) ListTodos(ctx contextx.Contextx, opts biz.ListTodosOptions) (todos []*wfM.Ticket, total int, err error) {
	got, err := i.issues.List(ctx)
	if err != nil {
		return nil, 0, err
	}

	var ret []*wfM.Ticket
	for _, v := range got {
		ret = append(ret, &wfM.Ticket{
			ID:        v.GetID(),
			Title:     v.GetTitle(),
			Completed: v.GetCompleted(),
		})
	}

	return ret, len(ret), nil
}

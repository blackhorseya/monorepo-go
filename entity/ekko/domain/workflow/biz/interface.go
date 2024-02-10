//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	idM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/identity/model"
	wfM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type ListTodosOptions struct {
	Page int
	Size int
}

// IWorkflowBiz is the interface for workflow biz.
type IWorkflowBiz interface {
	CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (todo *wfM.Ticket, err error)

	ListTodos(ctx contextx.Contextx, options ListTodosOptions) (todos []*wfM.Ticket, total int, err error)
}

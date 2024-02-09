//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	idM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/identity/model"
	wfM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IWorkflowBiz is the interface for workflow biz.
type IWorkflowBiz interface {
	CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (todo *wfM.Ticket, err error)
}

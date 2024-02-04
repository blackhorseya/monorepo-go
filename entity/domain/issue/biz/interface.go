//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IIssueBiz is the interface for issue biz
type IIssueBiz interface {
	// CreateTodo is used to create a new todo.
	CreateTodo(ctx contextx.Contextx, title string) (todo *model.Ticket, err error)

	// ListTodos is used to list all todos.
	ListTodos(ctx contextx.Contextx) (todos []*model.Ticket, total int, err error)

	// CompletedTodoByID is used to complete a todo by id.
	CompletedTodoByID(ctx contextx.Contextx, id uint64) error
}

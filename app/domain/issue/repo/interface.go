//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListOptions is used to define the list options.
type ListOptions struct {
	Limit  int
	Offset int
}

// Storager is used to define the storage interface.
type Storager interface {
	// List is used to list the todos.
	List(ctx contextx.Contextx, opts ListOptions) (todos []*model.Ticket, total int, err error)

	// Create is used to create a todo.
	Create(ctx contextx.Contextx, title string) (todo *model.Ticket, err error)

	// CompleteByID is used to complete a todo by id.
	CompleteByID(ctx contextx.Contextx, id uint64) error
}

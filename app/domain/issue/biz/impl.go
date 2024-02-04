package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
}

// NewIssueBiz is used to create a new issue biz instance.
func NewIssueBiz() (biz.IIssueBiz, error) {
	return &impl{}, nil
}

func (i *impl) CreateTodo(ctx contextx.Contextx, title string) (todo *model.Ticket, err error) {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

func (i *impl) ListTodos(ctx contextx.Contextx) (todos []*model.Ticket, total int, err error) {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

func (i *impl) CompletedTodoByID(ctx contextx.Contextx, id uint64) error {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

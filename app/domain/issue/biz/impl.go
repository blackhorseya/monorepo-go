package biz

import (
	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	storager repo.Storager
}

// NewIssueBiz is used to create a new issue biz instance.
func NewIssueBiz(storager repo.Storager) (biz.IIssueBiz, error) {
	return &impl{
		storager: storager,
	}, nil
}

func (i *impl) CreateTodo(ctx contextx.Contextx, title string) (todo *model.Ticket, err error) {
	ret, err := i.storager.Create(ctx, title)
	if err != nil {
		ctx.Error("create a todo into storage failed", zap.Error(err))
		return nil, err
	}

	return ret, nil
}

func (i *impl) ListTodos(ctx contextx.Contextx) (todos []*model.Ticket, total int, err error) {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

func (i *impl) CompletedTodoByID(ctx contextx.Contextx, id uint64) error {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

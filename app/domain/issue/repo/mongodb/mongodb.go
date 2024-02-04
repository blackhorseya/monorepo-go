package mongodb

import (
	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	rw *mongo.Client
}

// NewStorager is used to create a new issue storage instance.
func NewStorager(rw *mongo.Client) (repo.Storager, error) {
	return &impl{rw: rw}, nil
}

func (i *impl) List(ctx contextx.Contextx, opts repo.ListOptions) (todos []*model.Ticket, total int, err error) {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, title string) (todo *model.Ticket, err error) {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

func (i *impl) CompleteByID(ctx contextx.Contextx, id uint64) error {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	timeoutDuration = 5 * time.Second

	dbName = "ekko"

	collName = "todos"
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
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	now := timestamppb.Now()
	newTask := &model.Ticket{
		Id:        uuid.New().String(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	coll := i.rw.Database(dbName).Collection(collName)
	_, err = coll.InsertOne(timeout, newTask)
	if err != nil {
		ctx.Error("insert todo failed", zap.Error(err), zap.Any("new_task", &newTask))
		return nil, err
	}

	return newTask, nil
}

func (i *impl) CompleteByID(ctx contextx.Contextx, id uint64) error {
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

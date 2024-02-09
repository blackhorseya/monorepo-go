package mongodb

import (
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	rw *mongo.Client
}

// NewIssueRepoWithMongoDB is the constructor of IssueRepo with MongoDB.
func NewIssueRepoWithMongoDB(rw *mongo.Client) (repo.IIssueRepo, error) {
	return &impl{rw: rw}, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (issue agg.Issue, err error) {
	// todo: 2024/2/10|sean|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, item agg.Issue) error {
	// todo: 2024/2/10|sean|implement me
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, item agg.Issue) error {
	// todo: 2024/2/10|sean|implement me
	panic("implement me")
}

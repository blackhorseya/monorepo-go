package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timeoutDuration = 5 * time.Second

	dbName   = "ekko"
	collName = "issues"
)

type impl struct {
	rw *mongo.Client
}

// NewIssueRepoWithMongoDB is the constructor of IssueRepo with MongoDB.
func NewIssueRepoWithMongoDB(rw *mongo.Client) (repo.IIssueRepo, error) {
	return &impl{rw: rw}, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (item agg.Issue, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return agg.Issue{}, err
	}

	coll := i.rw.Database(dbName).Collection(collName)
	filter := bson.M{"_id": hex}
	result := coll.FindOne(timeout, filter)

	var got issue
	err = result.Decode(&got)
	if err != nil {
		return agg.Issue{}, err
	}

	ret, err := got.ToAggregate()
	if err != nil {
		return agg.Issue{}, err
	}

	return ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, item agg.Issue) error {
	// todo: 2024/2/10|sean|implement me
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, item agg.Issue) error {
	// todo: 2024/2/10|sean|implement me
	panic("implement me")
}

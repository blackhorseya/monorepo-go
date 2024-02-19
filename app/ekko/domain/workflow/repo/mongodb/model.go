package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type issue struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	OwnerID   string             `bson:"owner_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func newFromIssue(item agg.Issue) issue {
	now := time.Now()

	return issue{
		ID:        primitive.NewObjectID(),
		Title:     item.GetTitle(),
		Completed: item.GetCompleted(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ToAggregate converts the issue to aggregate.
func (x *issue) ToAggregate() (agg.Issue, error) {
	ret, err := agg.NewIssue(x.OwnerID, x.Title)
	if err != nil {
		return agg.Issue{}, err
	}

	ret.SetID(x.ID.Hex())
	ret.SetCompleted(x.Completed)

	return ret, nil
}

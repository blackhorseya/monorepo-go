package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeoutDuration = 5 * time.Second
	dbName          = "orianna"
	collName        = "stocks"
)

type impl struct {
	client *mongo.Client
}

// NewStockRepo is the constructor of stock repository.
func NewStockRepo(client *mongo.Client) (repo.IStockRepo, error) {
	return &impl{client: client}, nil
}

func (i *impl) List(ctx contextx.Contextx) ([]agg.Stock, error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	coll := i.client.Database(dbName).Collection(collName)
	cursor, err := coll.Find(timeout, bson.M{})
	if err != nil {
		return nil, err
	}

	var ret []agg.Stock
	for cursor.Next(timeout) {
		var got stock
		if err = cursor.Decode(&got); err != nil {
			return nil, err
		}

		ret = append(ret, got.ToAggregate())
	}

	return ret, nil
}

func (i *impl) BulkUpsertInfo(ctx contextx.Contextx, stocks []agg.Stock) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	now := time.Now()
	var models []mongo.WriteModel
	for _, v := range stocks {
		filter := bson.M{"_id": v.GetSymbol()}
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "_id", Value: v.GetSymbol()},
					{Key: "name", Value: v.GetName()},
					{Key: "industry_category", Value: v.GetIndustryCategory()},
					{Key: "exchange_name", Value: v.GetExchangeName()},
					{Key: "updated_at", Value: now},
				}}}).
			SetUpsert(true)
		models = append(models, model)
	}
	opts := options.BulkWrite().SetOrdered(false)

	coll := i.client.Database(dbName).Collection(collName)
	_, err := coll.BulkWrite(timeout, models, opts)
	if err != nil {
		return err
	}

	return nil
}

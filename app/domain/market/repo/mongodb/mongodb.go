package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	timeoutDuration = 5 * time.Second

	dbName   = "orianna"
	collName = "stocks"
)

type impl struct {
	rw *mongo.Client
}

// NewStorager is the factory method to create a Storager instance.
func NewStorager(rw *mongo.Client) (repo.Storager, error) {
	return &impl{rw: rw}, nil
}

func (i *impl) GetBySymbol(ctx contextx.Contextx, symbol string) (info *model.StockInfo, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	coll := i.rw.Database(dbName).Collection(collName)
	filter := bson.M{"_id": symbol}
	var ret *model.StockInfo
	err = coll.FindOne(timeout, filter).Decode(&ret)
	if err != nil {
		ctx.Error("find one error", zap.Error(err))
		return nil, err
	}

	return ret, nil
}

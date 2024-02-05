package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	keyFinmindAPI   = "FINMIND_API"
	keyFinmindToken = "FINMIND_TOKEN"
	keyDatabaseURL  = "DATABASE_URL"

	dbName   = "orianna"
	collName = "stocks"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(c context.Context) (Response, error) {
	uri, err := url.ParseRequestURI(os.Getenv(keyFinmindAPI))
	if err != nil {
		return Response{}, err
	}
	configx.C.Finmind.HTTP.URL = uri.String()

	token := os.Getenv(keyFinmindToken)
	if len(token) == 0 {
		return Response{}, errors.New("finmind token is not set")
	}
	configx.C.Finmind.Token = token

	client, err := finmindx.NewClient()
	if err != nil {
		return Response{}, err
	}

	databaseURL := os.Getenv(keyDatabaseURL)
	if len(databaseURL) == 0 {
		return Response{}, errors.New("database url is not set")
	}
	configx.A.Storage.Mongodb.DSN = databaseURL

	rw, err := mongodb.NewClientWithDSN(databaseURL)
	if err != nil {
		return Response{}, err
	}

	ctx := contextx.Background()
	got, err := client.TaiwanStockInfo(ctx)
	if err != nil {
		return Response{}, err
	}

	var models []mongo.WriteModel
	for _, v := range got {
		var date *timestamppb.Timestamp
		if !v.Date.IsZero() {
			date = timestamppb.New(v.Date)
		}

		filter := bson.M{"_id": v.StockID}
		doc := &model.StockInfo{
			Symbol:           v.StockID,
			Name:             v.StockName,
			IndustryCategory: v.IndustryCategory,
			Type:             v.Type,
			Date:             date,
		}
		models = append(models, mongo.NewReplaceOneModel().
			SetFilter(filter).
			SetReplacement(doc).
			SetUpsert(true),
		)
	}
	opts := options.BulkWrite().SetOrdered(false)

	result, err := rw.Database(dbName).Collection(collName).BulkWrite(ctx, models, opts)
	if err != nil {
		return Response{}, err
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(marshal),
	}, nil
}

func main() {
	zap.ReplaceGlobals(zap.NewExample())
	lambda.Start(Handler)
}

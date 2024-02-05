package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	keyFinmindAPI   = "FINMIND_API"
	keyFinmindToken = "FINMIND_TOKEN"
	keyDatabaseURL  = "DATABASE_URL"
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

	databaseURL := os.Getenv(keyDatabaseURL)
	if len(databaseURL) == 0 {
		return Response{}, errors.New("database url is not set")
	}
	configx.A.Storage.Mongodb.DSN = databaseURL

	client, err := finmindx.NewClient()
	if err != nil {
		return Response{}, err
	}

	ctx := contextx.Background()
	res, err := client.TaiwanStockInfo(ctx)
	if err != nil {
		return Response{}, err
	}

	var ret []*model.StockInfo
	for _, v := range res {
		var date *timestamppb.Timestamp
		if !v.Date.IsZero() {
			date = timestamppb.New(v.Date)
		}

		ret = append(ret, &model.StockInfo{
			Symbol:           v.StockID,
			Name:             v.StockName,
			IndustryCategory: v.IndustryCategory,
			Type:             v.Type,
			Date:             date,
		})
	}

	return Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       fmt.Sprintf("got %d stock info", len(ret)),
	}, nil
}

func main() {
	zap.ReplaceGlobals(zap.NewExample())
	lambda.Start(Handler)
}

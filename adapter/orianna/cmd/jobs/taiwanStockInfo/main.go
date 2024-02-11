package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"go.uber.org/zap"
)

const (
	dbName   = "orianna"
	collName = "stocks"
)

var (
	injector *Injector
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(c context.Context) (Response, error) {
	ctx := contextx.Background()

	got, err := injector.finmind.TaiwanStockInfo(ctx)
	if err != nil {
		return handleError(err)
	}
	ctx.Info("successfully fetch [TaiwanStockInfo] dataset", zap.Int("count", len(got)))

	var stocks []agg.Stock
	for _, v := range got {
		stocks = append(stocks, agg.NewStock(&model.Stock{
			Symbol:           v.StockID,
			Name:             v.StockName,
			IndustryCategory: v.IndustryCategory,
			ExchangeName:     v.Type,
		}))
	}

	err = injector.repo.BulkUpsertInfo(ctx, stocks)
	if err != nil {
		return handleError(err)
	}

	err = injector.notifier.SendText(ctx, "[TaiwanStockInfo] dataset has been updated")
	if err != nil {
		ctx.Warn("failed to send notification", zap.Error(err))
	}

	return Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       "ok",
	}, nil
}

func main() {
	err := configx.Load("", "sean")
	if err != nil {
		log.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		log.Fatal(err)
	}

	injector, err = BuildInjector()
	if err != nil {
		log.Fatal(err)
	}

	lambda.Start(Handler)
}

func handleError(err error) (Response, error) {
	ctx := contextx.Background()

	if injector.notifier != nil {
		_ = injector.notifier.SendText(ctx, fmt.Sprintf("[TaiwanStockInfo] failed to execute the job: %v", err))
	}

	ctx.Error("failed to execute the job", zap.Error(err))

	return Response{}, err
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
	endpoint = "https://openapi.twse.com.tw/v1/exchangeReport/STOCK_DAY_ALL"

	dbName = "stock_quotes"
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return handleError(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return handleError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return handleError(fmt.Errorf("response status code is not 200 OK, got: %d", resp.StatusCode))
	}

	var got []*StockDayResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return handleError(err)
	}

	now := time.Now()
	var stocks []agg.Stock
	for _, v := range got {
		open, _ := strconv.ParseFloat(v.OpeningPrice, 64)
		high, _ := strconv.ParseFloat(v.HighestPrice, 64)
		closePrice, _ := strconv.ParseFloat(v.ClosingPrice, 64)
		low, _ := strconv.ParseFloat(v.LowestPrice, 64)
		volume, _ := strconv.Atoi(v.TradeVolume)

		stocks = append(stocks, agg.NewStockWithQuota(
			&model.Stock{Symbol: v.Code},
			model.NewStockQuota(open, high, closePrice, low, volume, now),
		))
	}

	err = injector.repo.BulkUpdateQuota(ctx, stocks)
	if err != nil {
		return handleError(err)
	}

	err = injector.notifier.SendText(ctx, "[TaiwanDailyQuota] dataset has been updated")
	if err != nil {
		ctx.Warn("failed to send notification", zap.Error(err))
	}

	return Response{StatusCode: http.StatusOK}, nil
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
		_ = injector.notifier.SendText(ctx, fmt.Sprintf("[TaiwanDailyQuote] failed to execute the job: %v", err))
	}

	ctx.Error("failed to execute the job", zap.Error(err))

	return Response{}, err
}

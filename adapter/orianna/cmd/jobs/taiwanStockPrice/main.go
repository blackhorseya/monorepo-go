package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
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
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	ctx := contextx.Background()

	symbol, ok := request.QueryStringParameters["symbol"]
	if !ok {
		return handleError(fmt.Errorf("missing symbol"))
	}

	now := time.Now()
	var (
		err   error
		start = now
		end   = now
	)
	startStr, ok := request.QueryStringParameters["start"]
	if !ok {
		return handleError(fmt.Errorf("missing start"))
	}
	if startStr != "" {
		start, err = time.Parse(time.RFC3339, startStr)
		if err != nil {
			return handleError(fmt.Errorf("invalid start"))
		}
	}

	endStr, ok := request.QueryStringParameters["end"]
	if !ok {
		return handleError(fmt.Errorf("missing end"))
	}
	if endStr != "" {
		end, err = time.Parse(time.RFC3339, endStr)
		if err != nil {
			return handleError(fmt.Errorf("invalid end"))
		}
	}

	got, err := injector.finmind.TaiwanStockPriceV2(ctx, symbol, start, end)
	if err != nil {
		return handleError(err)
	}

	var stocks []*model.Stock
	for _, v := range got {
		var date *timestamppb.Timestamp
		if !v.Date.IsZero() {
			date = timestamppb.New(v.Date)
		}

		stocks = append(stocks, &model.Stock{
			Symbol: v.StockID,
			Name:   "",
			Price:  0,
			DailyQuote: &model.DailyQuote{
				Open:        v.Open,
				High:        v.Max,
				Low:         v.Min,
				Close:       v.Close,
				Volume:      v.TradingVolume,
				Value:       v.TradingMoney,
				Change:      v.Spread,
				Transaction: int64(v.TradingTurnover),
				Date:        date,
			},
		})
	}

	ctx.Info("successfully executed the job", zap.Any("stocks", stocks))

	// todo: 2024/2/7|sean|implement the logic

	return Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       symbol,
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
		_ = injector.notifier.SendText(ctx, fmt.Sprintf("[TaiwanStockPrice] failed to execute the job: %v", err))
	}

	ctx.Error("failed to execute the job", zap.Error(err))

	return Response{StatusCode: http.StatusInternalServerError}, err
}

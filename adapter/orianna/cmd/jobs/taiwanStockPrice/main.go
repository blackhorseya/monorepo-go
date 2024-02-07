package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"go.uber.org/zap"
)

const (
	dbName = "stock_quotes"
)

var (
	injector *Injector
	loc, _   = time.LoadLocation("Asia/Taipei")
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

	var points []*influxdb3.Point
	for _, v := range got {
		timestamp := time.Date(v.Date.Year(), v.Date.Month(), v.Date.Day(), 13, 30, 0, 0, loc)
		point := influxdb3.NewPointWithMeasurement("quotes").
			SetTag("symbol", symbol).
			SetDoubleField("open", v.Open).
			SetDoubleField("high", v.Max).
			SetDoubleField("low", v.Min).
			SetDoubleField("close", v.Close).
			SetDoubleField("change", v.Spread).
			SetIntegerField("volume", v.TradingVolume).
			SetIntegerField("value", v.TradingMoney).
			SetDoubleField("transaction", v.TradingTurnover).
			SetTimestamp(timestamp)

		points = append(points, point)
	}

	opts := &influxdb3.WriteOptions{Database: dbName}
	err = injector.rw.WritePointsWithOptions(ctx, opts, points...)
	if err != nil {
		return handleError(err)
	}

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

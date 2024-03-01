package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/blackhorseya/monorepo-go/pkg/timex"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

const (
	baseURL = "https://mops.twse.com.tw/mops/web/ajax_t100sb02_1"
)

var (
	loc, _   = time.LoadLocation("Asia/Taipei")
	injector *Injector
)

func Handler() (events.APIGatewayProxyResponse, error) {
	ctx := contextx.Background()

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		ctx.Info("visiting", zap.String("url", r.URL.String()))
	})

	var eventList []*agg.Event
	c.OnHTML("#myTable > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			split := strings.Split(tr.ChildText("td:nth-child(3)"), "至")
			if len(split) == 0 {
				ctx.Error(
					"split date string error",
					zap.String("text", tr.ChildText("td:nth-child(3)")),
				)
				return
			}
			atDate, err := time.ParseInLocation(
				"2006/01/02",
				timex.FormatTaiwanDate(strings.TrimSpace(split[len(split)-1])),
				loc,
			)
			if err != nil {
				ctx.Error(
					"parse date error",
					zap.Error(err),
					zap.String("text", tr.ChildText("td:nth-child(3)")),
				)
				return
			}

			timeStr := strings.ReplaceAll(tr.ChildText("td:nth-child(4)"), " ", "")
			atTime, err := time.ParseInLocation("15:04", timeStr, loc)
			if err != nil {
				ctx.Error("parse time error", zap.Error(err))
				return
			}

			at := time.Date(atDate.Year(), atDate.Month(), atDate.Day(), atTime.Hour(), atTime.Minute(), 0, 0, loc)

			symbol := strings.TrimSpace(tr.ChildText("td:nth-child(1)"))
			event := model.NewEvent(symbol, model.EventTypeEarningsCall, at)
			eventList = append(eventList, &agg.Event{Event: event})
		})
	})

	err := c.Post(baseURL, map[string]string{
		"encodeURIComponent": "1",
		"step":               "1",
		"firstin":            "1",
		"off":                "1",
		"TYPEK":              "sii",
		"year":               "113",
		"month":              "02",
	})
	if err != nil {
		return handleError(err)
	}

	err = injector.repo.BulkUpsert(ctx, eventList)
	if err != nil {
		return handleError(err)
	}

	err = injector.notifier.SendText(ctx, "[FetchEarningsCall] dataset has been updated")
	if err != nil {
		ctx.Warn("failed to send notification", zap.Error(err))
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func handleError(err error) (events.APIGatewayProxyResponse, error) {
	ctx := contextx.Background()

	if injector.notifier != nil {
		_ = injector.notifier.SendText(ctx, fmt.Sprintf("[FetchEarningsCall] failed to execute the job: %v", err))
	}

	ctx.Error("failed to execute the job", zap.Error(err))

	return events.APIGatewayProxyResponse{}, err
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

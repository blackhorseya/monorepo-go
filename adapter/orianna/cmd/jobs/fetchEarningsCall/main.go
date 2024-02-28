package main

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/blackhorseya/monorepo-go/pkg/timex"
	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	baseURL = "https://mops.twse.com.tw/mops/web/ajax_t100sb02_1"
)

var (
	loc, _ = time.LoadLocation("Asia/Taipei")
)

func Handler() (events.APIGatewayProxyResponse, error) {
	ctx := contextx.Background()

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		ctx.Info("visiting", zap.String("url", r.URL.String()))
	})

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

			call := model.EarningsCall{
				ID:         uuid.New(),
				Symbol:     tr.ChildText("td:nth-child(1)"),
				Host:       "",
				OccurredAt: at,
			}
			ctx.Info("fetching", zap.Any("call", call))
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
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
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

	lambda.Start(Handler)
}

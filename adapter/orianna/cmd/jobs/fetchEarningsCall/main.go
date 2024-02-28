package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

const (
	baseURL = "https://mops.twse.com.tw/mops/web/ajax_t100sb02_1"
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
			ctx.Info("fetching", zap.String("symbol", tr.ChildText("td:nth-child(1)")))
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

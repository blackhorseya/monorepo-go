package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"go.uber.org/zap"
)

const (
	baseURL = "https://mops.twse.com.tw/mops/web/ajax_t100sb02_1"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	ctx := contextx.Background()

	payload := url.Values{}
	payload.Set("encodeURIComponent", "1")
	payload.Set("step", "1")
	payload.Set("firstin", "1")
	payload.Set("off", "1")
	payload.Set("TYPEK", "sii")
	payload.Set("year", "113")
	payload.Set("month", "02")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return events.APIGatewayProxyResponse{}, errors.New("status code is not 200, got " + resp.Status)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	ctx.Info("got", zap.String("body", string(got)))

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

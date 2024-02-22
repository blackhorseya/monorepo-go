package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/monorepo-go/adapter/reurl/cmd/restful"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/spf13/viper"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	err := configx.Load("", "sean")
	if err != nil {
		log.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.ReURL)

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		log.Fatal(err)
	}

	app, err := restful.NewRestful(viper.GetViper())
	if err != nil {
		log.Fatal(err)
	}

	err = app.InitRouting()
	if err != nil {
		log.Fatal(err)
	}

	ginLambda = ginadapter.New(app.GetRouter())

	lambda.Start(Handler)
}

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/gin-gonic/gin"
)

var (
	injector  *Injector
	ginLambda *ginadapter.GinLambda
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
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

	injector, err = BuildInjector()
	if err != nil {
		log.Fatal(err)
	}
	injector.registerRoutes()

	ginLambda = ginadapter.New(injector.server.Router)

	lambda.Start(Handler)
}

func (i *Injector) registerRoutes() {
	i.server.Router.POST("/callback", func(c *gin.Context) {
		// todo: 2024/2/20|sean|implement callback
		c.JSON(http.StatusOK, response.OK)
	})
	i.server.Router.GET("/:code", func(c *gin.Context) {
		// todo: 2024/2/20|sean|implement redirect
		c.JSON(http.StatusOK, response.OK.WithData(c.Param("code")))
	})
}

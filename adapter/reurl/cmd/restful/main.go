package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
)

var (
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

	server, err := httpx.NewServer(contextx.Background())
	if err != nil {
		log.Fatal(err)
	}

	// register routes
	server.Router.POST("/callback", func(c *gin.Context) {
		// todo: 2024/2/20|sean|implement callback
		c.JSON(http.StatusOK, response.OK)
	})
	server.Router.GET("/:code", func(c *gin.Context) {
		// todo: 2024/2/20|sean|implement redirect
		c.JSON(http.StatusOK, response.OK.WithData(c.Param("code")))
	})

	ginLambda = ginadapter.New(server.Router)

	lambda.Start(Handler)
}

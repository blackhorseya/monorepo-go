package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"go.uber.org/zap"
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
	i.server.Router.POST("/callback", i.callback)
	i.server.Router.GET("/:code", func(c *gin.Context) {
		// todo: 2024/2/20|sean|implement redirect
		c.JSON(http.StatusOK, response.OK.WithData(c.Param("code")))
	})
}
func (i *Injector) callback(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	eventList, err := i.bot.ParseRequest(c.Request)
	if err != nil {
		if errors.Is(err, linebot.ErrInvalidSignature) {
			ctx.Error("invalid line bot signature", zap.Error(err))
			_ = c.Error(err)
		} else {
			ctx.Error("parse line bot request error", zap.Error(err))
			_ = c.Error(err)
		}

		return
	}

	var messages []linebot.SendingMessage
	for _, event := range eventList {
		if event.Type == linebot.EventTypeMessage {
			message, ok := event.Message.(*linebot.TextMessage)
			if !ok {
				continue
			}

			messages, err = i.handleMessage(ctx, event, message)
			if err != nil {
				ctx.Warn("handle message error", zap.Error(err))
				continue
			}

			_, err = i.bot.ReplyMessage(event.ReplyToken, messages...).Do()
			if err != nil {
				_ = c.Error(err)
				return
			}
		}
	}

	c.Status(http.StatusOK)
}

func (i *Injector) handleMessage(
	_ contextx.Contextx,
	event *linebot.Event,
	message *linebot.TextMessage,
) ([]linebot.SendingMessage, error) {
	text := message.Text
	if text == "ping" {
		return []linebot.SendingMessage{
			linebot.NewTextMessage("pong"),
		}, nil
	}

	if text == "whoami" {
		return []linebot.SendingMessage{
			linebot.NewTextMessage(event.Source.UserID),
		}, nil
	}

	return nil, errors.New("unknown message")
}

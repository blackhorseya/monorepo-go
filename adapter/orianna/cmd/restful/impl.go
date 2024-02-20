package restful

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/blackhorseya/monorepo-go/adapter/orianna/api/docs" // swagger docs
	v1 "github.com/blackhorseya/monorepo-go/adapter/orianna/cmd/restful/v1"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	server *httpx.Server
	svc    biz.IMarketBiz
	bot    *linebot.Client
}

func newRestful(svc biz.IMarketBiz, bot *linebot.Client) (adapterx.Servicer, error) {
	ctx := contextx.Background()

	server, err := httpx.NewServerWithContextx(ctx)
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		svc:    svc,
		bot:    bot,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	i.server.Router.POST("/callback", i.callback)

	// register router
	api := i.server.Router.Group("/api")
	{
		api.GET("/healthz", i.healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1.Handle(api.Group("/v1"), i.svc)
	}

	err := i.server.Start(ctx)
	if err != nil {
		return err
	}

	ctx.Info(
		"swagger docs",
		zap.String("url", fmt.Sprintf(
			"http://%s/api/docs/index.html",
			strings.ReplaceAll(configx.A.HTTP.GetAddr(), "0.0.0.0", "localhost"),
		)),
	)

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		err := i.server.Stop(ctx)
		if err != nil {
			ctx.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}

// healthz is used to check the health of the service.
// @Summary healthz
// @Description Check the health of the service.
// @Tags healthz
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /healthz [get]
func (i *impl) healthz(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}

// callback is used to handle the callback from the third-party service.
// @Summary callback
// @Description Handle the callback from the third-party service.
// @Tags callback
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /callback [post]
func (i *impl) callback(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	events, err := i.bot.ParseRequest(c.Request)
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
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			message, ok := event.Message.(*linebot.TextMessage)
			if !ok {
				continue
			}

			messages, err = i.handleMessage(ctx, message)
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

func (i *impl) handleMessage(ctx contextx.Contextx, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	text := message.Text
	if text == "ping" {
		return []linebot.SendingMessage{
			linebot.NewTextMessage("pong"),
		}, nil
	}

	split := strings.Split(message.Text, ".")
	if len(split) == 2 && split[0] == "q" {
		symbol := split[1]

		stock, err := i.svc.GetStockBySymbol(ctx, symbol)
		if err != nil {
			return nil, err
		}

		return []linebot.SendingMessage{
			stock.FlexMessage(),
		}, nil
	}

	return []linebot.SendingMessage{
		linebot.NewTextMessage("I don't understand what you said."),
	}, nil
}

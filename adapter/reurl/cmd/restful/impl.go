package restful

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/blackhorseya/monorepo-go/adapter/reurl/api/docs" // swagger docs
	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
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
	bot    *linebot.Client
	svc    shortB.IShorteningBiz
}

func newService(bot *linebot.Client, svc shortB.IShorteningBiz) (adapterx.Servicer, error) {
	server, err := httpx.NewServer()
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		bot:    bot,
		svc:    svc,
	}, nil
}

func newRestful(bot *linebot.Client, svc shortB.IShorteningBiz) (adapterx.Restful, error) {
	server, err := httpx.NewServer()
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		bot:    bot,
		svc:    svc,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	err := i.InitRouting()
	if err != nil {
		return err
	}

	err = i.server.Start(ctx)
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

func (i *impl) InitRouting() error {
	api := i.server.Router.Group("/api")
	{
		api.GET("/healthz", i.healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		api.POST("/callback", i.callback)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
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

			messages, err = i.handleMessage(ctx, event, message)
			if err != nil {
				ctx.Warn("handle message error", zap.Error(err), zap.String("text", message.Text))
				continue
			}

			_, err = i.bot.ReplyMessage(event.ReplyToken, messages...).Do()
			if err != nil {
				_ = c.Error(err)
				return
			}
		}
	}

	c.JSON(http.StatusOK, response.OK)
}

func (i *impl) handleMessage(
	ctx contextx.Contextx,
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

	uri, err := url.ParseRequestURI(text)
	if err != nil {
		return nil, err
	}

	record, err := i.svc.CreateShortenedURL(ctx, uri.String())
	if err != nil {
		return nil, err
	}

	return []linebot.SendingMessage{
		linebot.NewTextMessage(record.ShortUrl),
	}, nil
}

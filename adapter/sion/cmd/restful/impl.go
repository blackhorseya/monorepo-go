package restful

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/blackhorseya/monorepo-go/adapter/sion/api/docs" // swagger docs
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
}

func newRestful(bot *linebot.Client) (adapterx.Servicer, error) {
	server, err := httpx.NewServer(contextx.Background())
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		bot:    bot,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	// register router
	api := i.server.Router.Group("/api")
	{
		api.GET("/healthz", i.healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		api.POST("/callback", i.callback)
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

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.LocationMessage:
				if _, err = i.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Address)).Do(); err != nil {
					ctx.Error("reply location message error", zap.Error(err))
					_ = c.Error(err)
				}
			default:
				ctx.Warn("unsupported message type", zap.String("type", string(event.Type)))
			}
		}
	}

	c.JSON(http.StatusOK, response.OK)
}

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
	"github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
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
	biz    biz.IMarketBiz
}

func newRestful(bot *linebot.Client, biz biz.IMarketBiz) (adapterx.Servicer, error) {
	ctx := contextx.Background()

	server, err := httpx.NewServer(ctx)
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		bot:    bot,
		biz:    biz,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	// register router
	api := i.server.Router.Group("/api")
	{
		api.GET("/healthz", i.healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	i.server.Router.POST("/callback", i.callback)

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

	err = i.handleEvents(ctx, events)
	if err != nil {
		ctx.Error("handle line bot events error", zap.Error(err))
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

//nolint:gocognit // it's ok
func (i *impl) handleEvents(ctx contextx.Contextx, events []*linebot.Event) error {
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage { //nolint:nestif // for better readability
			message, ok := event.Message.(*linebot.TextMessage)
			if !ok {
				return errors.New("not text message")
			}

			if message.Text == "ping" {
				_, err := i.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("pong")).Do()
				if err != nil {
					return err
				}
			}

			split := strings.Split(message.Text, ".")
			if len(split) == 2 && split[0] == "q" {
				// query stock by symbol
				symbol := split[1]
				stock, err := i.biz.GetStockBySymbol(ctx, symbol)
				if err != nil {
					return err
				}

				_, err = i.bot.ReplyMessage(event.ReplyToken, stock.FlexMessage()).Do()
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

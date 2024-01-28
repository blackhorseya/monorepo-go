package restful

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/blackhorseya/monorepo-go/adapter/shortenurl/api/docs" // swagger docs
	v1 "github.com/blackhorseya/monorepo-go/adapter/shortenurl/cmd/restful/v1"
	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	configx2 "github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func initAPP() *configx2.Application {
	return &configx2.C.ShortenURL
}

type impl struct {
	app    *configx2.Application
	server *httpx.Server
	svc    shortB.IShorteningBiz
}

// New will create a restful service.
func newRestful(viper *viper.Viper, app *configx2.Application, svc shortB.IShorteningBiz) (adapterx.Servicer, error) {
	ctx := contextx.Background()

	server, err := httpx.NewServerWithAPP(ctx, app)
	if err != nil {
		return nil, err
	}

	return &impl{
		app:    app,
		server: server,
		svc:    svc,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

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
			strings.ReplaceAll(i.app.HTTP.GetAddr(), "0.0.0.0", "localhost"),
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

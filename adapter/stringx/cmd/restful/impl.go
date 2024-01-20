package restful

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/blackhorseya/monorepo-go/adapter/stringx/api/docs" // swagger docs
	v1 "github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/restful/v1"
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/internal/pkg/transports/httpx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	server *httpx.Server
	svc    biz.IStringBiz
}

func newRouter() *gin.Engine {
	return gin.New()
}

func newImpl(
	svc biz.IStringBiz,
) (adapterx.Servicer, error) {
	ctx := contextx.Background()

	server, err := httpx.NewServer(ctx)
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		svc:    svc,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	api := i.server.Router.Group("/api")
	{
		api.GET("/healthz", healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1.Handle(api.Group("/v1"), i.svc)
	}

	err := i.server.Start(ctx)
	if err != nil {
		return err
	}

	ctx.Info(
		"swagger docs",
		zap.String("url", fmt.Sprintf("http://localhost:%d/api/docs/index.html", configx.C.HTTP.Port)),
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
// @Router /healthz [get]
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}

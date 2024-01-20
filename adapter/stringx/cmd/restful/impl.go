package restful

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/blackhorseya/monorepo-go/adapter/stringx/api/docs" // swagger docs
	v1 "github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/restful/v1"
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	config *configx.Config

	router *gin.Engine
	server *http.Server
	svc    biz.IStringBiz
}

func newRouter() *gin.Engine {
	return gin.New()
}

func newImpl(
	viper *viper.Viper,
	config *configx.Config,
	svc biz.IStringBiz,
	router *gin.Engine,
) adapterx.Servicer {
	return &impl{
		viper:  viper,
		config: config,
		router: router,
		server: nil,
		svc:    svc,
	}
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	i.router.Use(ginzap.GinzapWithConfig(ctx, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz'"},
		Context:    nil,
	}))
	i.router.Use(ginzap.CustomRecoveryWithZap(ctx, true, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal server error",
		})
	}))

	i.router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := i.router.Group("/api")
	{
		v1.Handle(api.Group("/v1"), i.svc)
	}

	addr := fmt.Sprintf("%s:%d", i.config.HTTP.Host, i.config.HTTP.Port)
	i.server = &http.Server{
		Addr:                         addr,
		Handler:                      i.router,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            3 * time.Second,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	go func() {
		ctx.Info("start restful service", zap.String("addr", i.server.Addr))

		err := i.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Fatal("restful service error", zap.Error(err))
		}
	}()

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		err := i.server.Shutdown(ctx)
		if err != nil {
			ctx.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}

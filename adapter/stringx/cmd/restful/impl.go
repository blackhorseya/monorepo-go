package restful

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/transport"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	logger *zap.Logger
	router *gin.Engine

	server *http.Server
	svc    biz.IStringBiz
}

func newRouter() *gin.Engine {
	return gin.New()
}

func newImpl(viper *viper.Viper, logger *zap.Logger, svc biz.IStringBiz, router *gin.Engine) adapterx.Servicer {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "restful")),
		router: router,
		server: nil,
		svc:    svc,
	}
}

func (i *impl) Start() error {
	i.router.Use(ginzap.GinzapWithConfig(i.logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz'"},
		Context:    nil,
	}))
	i.router.Use(ginzap.CustomRecoveryWithZap(i.logger, true, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal server error",
		})
	}))

	uppercaseHandler := transport.MakeUppercaseHandler(contextx.Background(), endpoints.MakeUppercaseEndpoint(i.svc))
	countHandler := transport.MakeCountHandler(contextx.Background(), endpoints.MakeCountEndpoint(i.svc))

	i.router.POST("/uppercase", gin.WrapH(uppercaseHandler))
	i.router.POST("/count", gin.WrapH(countHandler))

	i.server = &http.Server{
		Addr:                         "0.0.0.0:8080",
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
		i.logger.Info("start restful service", zap.String("addr", i.server.Addr))

		err := i.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			i.logger.Fatal("restful service error", zap.Error(err))
		}
	}()

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		i.logger.Info("receive signal", zap.String("signal", sig.String()))

		timeout, cancelFunc := contextx.WithTimeout(contextx.Background(), 5*time.Second)
		defer cancelFunc()

		err := i.server.Shutdown(timeout)
		if err != nil {
			i.logger.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}

package restful

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	config *configx.Config
	logger *zap.Logger

	router *gin.Engine
	server *http.Server
}

func (i *impl) Start() error {
	addr := fmt.Sprintf("%s:%d", i.config.HTTP.Host, i.config.HTTP.Port)
	i.server = &http.Server{
		Addr:              addr,
		Handler:           i.router,
		ReadHeaderTimeout: 5 * time.Second,
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
			i.logger.Warn("shutdown restful server error", zap.Error(err))
		}
	}

	return nil
}

func newRestful(viper *viper.Viper) (adapterx.Servicer, error) {
	return &impl{
		viper:  viper,
		config: configx.NewExample(),
		logger: zap.NewExample(),
		router: gin.New(),
		server: nil,
	}, nil
}

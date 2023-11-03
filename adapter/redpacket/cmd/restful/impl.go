package restful

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	config *configx.Config
	logger *zap.Logger

	router *gin.Engine
	server *http.Server
}

func newImpl(config *configx.Config, logger *zap.Logger) adapterx.Servicer {
	return &impl{
		config: config,
		logger: logger.With(zap.String("type", "restful")),
		router: gin.New(),
		server: nil,
	}
}

func (i *impl) Start() error {
	// todo: 2023/11/3|sean|implement me
	panic("implement me")
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

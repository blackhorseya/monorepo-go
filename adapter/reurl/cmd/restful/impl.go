package restful

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	server *httpx.Server
}

func newService() (adapterx.Servicer, error) {
	server, err := httpx.NewServer()
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
	}, nil
}

func newRestful() (adapterx.Restful, error) {
	server, err := httpx.NewServer()
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
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
	// todo: 2024/2/23|sean|implement the routing

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

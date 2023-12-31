package grpcserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/grpcserver/s2s"
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	grpcserver "google.golang.org/grpc"
)

type impl struct {
	viper  *viper.Viper
	config *configx.Config
	logger *zap.Logger

	server *grpcserver.Server
	svc    biz.IStringBiz
}

func newImpl(viper *viper.Viper, config *configx.Config, logger *zap.Logger, svc biz.IStringBiz) adapterx.Servicer {
	return &impl{
		viper:  viper,
		config: config,
		logger: logger.With(zap.String("type", "grpc")),
		server: nil,
		svc:    svc,
	}
}

func (i *impl) Start() error {
	i.server = grpcserver.NewServer()

	model.RegisterStringxServiceServer(i.server, s2s.NewServer(
		endpoints.MakeUppercaseEndpoint(i.svc),
		endpoints.MakeCountEndpoint(i.svc),
	))

	addr := fmt.Sprintf("%s:%d", i.config.GRPC.Host, i.config.GRPC.Port)

	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			i.logger.Fatal("listen error", zap.Error(err))
		}

		i.logger.Info("start grpc server", zap.String("address", listen.Addr().String()))
		err = i.server.Serve(listen)
		if err != nil {
			i.logger.Fatal("serve error", zap.Error(err))
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

		i.server.GracefulStop()
	}

	return nil
}

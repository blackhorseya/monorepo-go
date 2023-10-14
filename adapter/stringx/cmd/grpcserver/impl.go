package grpcserver

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/transport/grpc"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	grpcserver "google.golang.org/grpc"
)

type impl struct {
	viper  *viper.Viper
	logger *zap.Logger

	server *grpcserver.Server
	svc    biz.IStringBiz
}

func newImpl(viper *viper.Viper, logger *zap.Logger, svc biz.IStringBiz) adapterx.Servicer {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "grpc")),
		server: nil,
		svc:    svc,
	}
}

func (i *impl) Start() error {
	i.server = grpcserver.NewServer()

	model.RegisterStringxServiceServer(i.server, grpc.New(
		endpoints.MakeUppercaseEndpoint(i.svc),
		endpoints.MakeCountEndpoint(i.svc),
	))

	go func() {
		listen, err := net.Listen("tcp", ":1234") //nolint:gosec // todo: 2023/10/12|sean|impl me
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

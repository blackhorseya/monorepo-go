package grpcserver

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	logger *zap.Logger
}

func newImpl(viper *viper.Viper, logger *zap.Logger) adapterx.Servicer {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "grpc")),
	}
}

func (i *impl) Start() error {
	// todo: 2023/10/12|sean|impl me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// todo: 2023/10/12|sean|impl me
	panic("implement me")
}

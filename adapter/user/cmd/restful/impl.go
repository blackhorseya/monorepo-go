package restful

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	logger *zap.Logger
}

func newImpl(viper *viper.Viper, logger *zap.Logger) adapterx.Restful {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "restful")),
	}
}

func (i *impl) Start() error {
	i.logger.Info("start restful service")

	// todo: 2023/10/12|sean|impl me

	return nil
}

func (i *impl) AwaitSignal() error {
	i.logger.Info("await restful service signal")

	// todo: 2023/10/12|sean|impl me

	return nil
}

func (i *impl) InitRouting() error {
	// todo: 2023/10/12|sean|impl me
	panic("implement me")
}

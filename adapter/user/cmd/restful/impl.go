package restful

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	logger *zap.Logger

	svc biz.IStringBiz
}

func newImpl(viper *viper.Viper, logger *zap.Logger, svc biz.IStringBiz) adapterx.Servicer {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "restful")),
		svc:    svc,
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

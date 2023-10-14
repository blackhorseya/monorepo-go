package cronjob

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	config *configx.Config
	logger *zap.Logger
}

func newImpl(viper *viper.Viper, config *configx.Config, logger *zap.Logger) adapterx.Servicer {
	return &impl{
		viper:  viper,
		config: config,
		logger: logger.With(zap.String("type", "cronjob")),
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

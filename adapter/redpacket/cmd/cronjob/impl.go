package cronjob

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	config *configx.Config
	logger *zap.Logger
}

func newCronjob(v *viper.Viper, config *configx.Config, logger *zap.Logger) (adapterx.Servicer, error) {
	return &impl{
		config: config,
		logger: logger,
	}, nil
}

func (i *impl) Start() error {
	// todo: 2023/11/22|sean|impl me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// todo: 2023/11/22|sean|impl me
	panic("implement me")
}

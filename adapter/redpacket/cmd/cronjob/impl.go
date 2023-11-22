package cronjob

import (
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	config *configx.Config
	logger *zap.Logger

	interval time.Duration
}

func newCronjob(v *viper.Viper, config *configx.Config, logger *zap.Logger) (adapterx.Servicer, error) {
	return &impl{
		config:   config,
		logger:   logger,
		interval: v.GetDuration("interval"),
	}, nil
}

func (i *impl) Start() error {
	i.logger.Info("cronjob service start", zap.Duration("interval", i.interval))

	// todo: 2023/11/22|sean|impl me
	return nil
}

func (i *impl) AwaitSignal() error {
	// todo: 2023/11/22|sean|impl me
	return nil
}

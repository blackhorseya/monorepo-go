package cronjob

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	config *configx.Config
	logger *zap.Logger

	taskC chan time.Time
	done  chan struct{}
}

func newImpl(viper *viper.Viper, config *configx.Config, logger *zap.Logger) adapterx.Servicer {
	return &impl{
		viper:  viper,
		config: config,
		logger: logger.With(zap.String("type", "cronjob")),
		taskC:  make(chan time.Time, 1),
		done:   make(chan struct{}, 0),
	}
}

func (i *impl) Start() error {
	i.logger.Info("start cronjob")

	go func() {
		// todo: 2023/10/15|sean|pass config to cronjob
		ticker := time.NewTicker(1 * time.Second)

		for {
			select {
			case now := <-ticker.C:
				select {
				case i.taskC <- now:
				case <-time.After(50 * time.Millisecond):
					return
				}
			case <-i.done:
				i.logger.Info("stop cronjob")
				ticker.Stop()
				return
			case t := <-i.taskC:
				i.logger.Debug("do cronjob", zap.Time("trigger_at", t))
			}
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

		i.done <- struct{}{}
	}

	return nil
}

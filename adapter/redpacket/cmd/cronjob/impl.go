package cronjob

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	config *configx.Config
	logger *zap.Logger

	interval time.Duration
	taskC    chan contextx.Contextx
	done     chan struct{}
}

func newCronjob(v *viper.Viper, config *configx.Config, logger *zap.Logger) (adapterx.Servicer, error) {
	return &impl{
		config:   config,
		logger:   logger,
		interval: v.GetDuration("interval"),
		taskC:    make(chan contextx.Contextx, 1),
		done:     make(chan struct{}),
	}, nil
}

func (i *impl) Start() error {
	i.logger.Info("cronjob service start", zap.Duration("interval", i.interval))

	ticker := time.NewTicker(i.interval)

	// produce task
	go func() {
		for {
			select {
			case <-i.done:
				break
			case <-ticker.C:
				ctx := contextx.WithLogger(i.logger)

				select {
				case i.taskC <- ctx:
				case <-time.After(50 * time.Millisecond):
					ctx.Warn("task channel is full then drop task")
				}
			}
		}
	}()

	// consume task
	go func() {
		for {
			select {
			case <-i.done:
				break
			case ctx := <-i.taskC:
				ctx.Info("receive task")
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

		// todo: 2023/11/22|sean|impl me
	}

	return nil
}

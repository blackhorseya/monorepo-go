package cronjob

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/google/uuid"
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
				id := uuid.New().String()
				ctx := contextx.WithValue(contextx.WithLogger(i.logger), "id", id)

				ctx.Debug("produce task", zap.String("id", id))

				select {
				case i.taskC <- ctx:
				case <-time.After(50 * time.Millisecond):
					ctx.Warn("task channel is full then drop task", zap.String("id", id))
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
				id := ctx.Value("id").(string)
				delay := time.Duration(1+rand.Intn(10)) * time.Second

				ctx.Info("receive task", zap.String("id", id), zap.Duration("delay", delay))
				time.Sleep(delay)
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

		os.Exit(0)
	}

	return nil
}

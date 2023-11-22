package cronjob

import (
	"crypto/rand"
	"math/big"
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
	ticker   *time.Ticker
	taskC    chan contextx.Contextx
	done     chan struct{}
}

func newCronjob(v *viper.Viper, config *configx.Config, logger *zap.Logger) (adapterx.Servicer, error) {
	return &impl{
		config:   config,
		logger:   logger,
		interval: v.GetDuration("interval"),
		ticker:   time.NewTicker(v.GetDuration("interval")),
		taskC:    make(chan contextx.Contextx, 1),
		done:     make(chan struct{}),
	}, nil
}

func (i *impl) Start() error {
	i.logger.Info("cronjob service start", zap.Duration("interval", i.interval))

	go i.produce()
	go i.consume()

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

func (i *impl) produce() {
	for {
		select {
		case <-i.done:
			break
		case <-i.ticker.C:
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
}

func (i *impl) consume() {
	for {
		select {
		case <-i.done:
			break
		case ctx := <-i.taskC:
			id, ok := ctx.Value("id").(string)
			if !ok {
				ctx.Error("get task id failed")
				return
			}

			n, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				ctx.Error("get random number failed", zap.Error(err))
				return
			}
			delay := time.Duration(1+n.Int64()) * time.Second

			ctx.Info("receive task", zap.String("id", id), zap.Duration("delay", delay))
			time.Sleep(delay)
		}
	}
}

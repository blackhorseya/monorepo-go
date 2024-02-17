package consumer

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	done chan struct{}
}

func newConsumer() (adapterx.Servicer, error) {
	return &impl{
		done: make(chan struct{}),
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	go func() {
		ctx.Info("start consumer...")

		for {
			select {
			case <-i.done:
				ctx.Info("consumer is stopping")
				return
			default:
				ctx.Info("consumer is running")
				time.Sleep(5 * time.Second)
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
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		i.done <- struct{}{}

		ctx.Info("consumer is stopped")
	}

	return nil
}

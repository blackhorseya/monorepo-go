package cronjob

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper *viper.Viper

	taskC chan time.Time
	done  chan struct{}
}

func newImpl(viper *viper.Viper) adapterx.Servicer {
	return &impl{
		viper: viper,
		taskC: make(chan time.Time, 1),
		done:  make(chan struct{}),
	}
}

func (i *impl) Start() error {
	ctx := contextx.Background()
	ctx.Info("start cronjob")

	go func() {
		ticker := time.NewTicker(time.Duration(configx.C.Cronjob.Interval) * time.Second)

		i.taskC <- time.Now()

		for {
			select {
			case now := <-ticker.C:
				select {
				case i.taskC <- now:
				case <-time.After(50 * time.Millisecond):
					return
				}
			case <-i.done:
				ctx.Info("stop cronjob")
				ticker.Stop()
				return
			case t := <-i.taskC:
				ctx.Debug("do cronjob", zap.Time("trigger_at", t))
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
	}

	return nil
}

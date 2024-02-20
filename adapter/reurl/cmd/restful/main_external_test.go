//go:build external

package main

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"go.uber.org/zap"
)

func TestStart(t *testing.T) {
	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.ReURL)

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		t.Fatal(err)
	}

	injector, err = BuildInjector()
	if err != nil {
		t.Fatal(err)
	}

	injector.registerRoutes()

	ctx := contextx.Background()
	err = injector.server.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		err = injector.server.Stop(ctx)
		if err != nil {
			ctx.Error("shutdown restful server error", zap.Error(err))
		}
	}
}

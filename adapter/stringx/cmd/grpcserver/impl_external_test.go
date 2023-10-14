//go:build external

package grpcserver_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/grpcserver"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func TestImpl_Start(t *testing.T) {
	v := viper.GetViper()
	logger := zap.NewExample()

	service := grpcserver.New(v, logger)
	err := service.Start()
	if err != nil {
		t.Errorf("Start() error = %v", err)
	}

	err = service.AwaitSignal()
	if err != nil {
		t.Errorf("AwaitSignal() error = %v", err)
	}

	t.Skip("skip external test")
}

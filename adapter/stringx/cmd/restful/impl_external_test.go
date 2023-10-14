//go:build external

package restful_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/restful"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func TestImpl_Start(t *testing.T) {
	v := viper.GetViper()
	logger := zap.NewExample()

	service, err := restful.New(v, logger)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	err = service.Start()
	if err != nil {
		t.Errorf("Start() error = %v", err)
		return
	}

	err = service.AwaitSignal()
	if err != nil {
		t.Errorf("AwaitSignal() error = %v", err)
		return
	}

	t.Skip("skip external test")
}

//go:build external

package restful_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/adapter/shortenurl/cmd/restful"
	"github.com/spf13/viper"
)

func TestImpl_Start(t *testing.T) {
	v := viper.GetViper()

	service, err := restful.NewService(v)
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

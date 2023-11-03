//go:build external

package restful_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd/restful"
	"github.com/spf13/viper"
)

func TestImpl_Start(t *testing.T) {
	v := viper.GetViper()

	servicer, err := restful.New(v)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	err = servicer.Start()
	if err != nil {
		t.Errorf("Start() error = %v", err)
		return
	}

	err = servicer.AwaitSignal()
	if err != nil {
		t.Errorf("AwaitSignal() error = %v", err)
		return
	}

	t.Skip("skip external test")
}

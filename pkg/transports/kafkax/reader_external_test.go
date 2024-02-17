//go:build external

package kafkax

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.uber.org/zap"
)

func TestNewReader(t *testing.T) {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	reader, err := NewReader()
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	stats := reader.Stats()
	contextx.Background().Debug("stats", zap.Any("stats", &stats))
}

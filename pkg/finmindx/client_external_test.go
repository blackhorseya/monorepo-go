//go:build external

package finmindx_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"go.uber.org/zap"
)

func TestImpl_Do(t *testing.T) {
	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		t.Fatal(err)
	}

	client, err := finmindx.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	var got *finmindx.TaiwanStockPriceResponse
	err = client.Do(ctx, "TaiwanStockPrice", map[string]string{
		"data_id":    "2330",
		"start_date": "2024-02-01",
		"end_date":   "2024-02-01",
	}, &got)
	if err != nil {
		t.Fatal(err)
	}
	ctx.Debug("got", zap.Any("got", &got))

	t.Skip("skipping test in external package")
}

func TestImpl_TaiwanStockInfo(t *testing.T) {
	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		t.Fatal(err)
	}

	client, err := finmindx.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	res, err := client.TaiwanStockInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	ctx.Debug("res", zap.Any("res", res))

	t.Skip("skipping test in external package")
}

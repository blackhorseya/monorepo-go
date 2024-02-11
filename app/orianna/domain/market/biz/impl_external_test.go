//go:build external

package biz

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.uber.org/zap"
)

func TestImpl_ListStocks(t *testing.T) {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	biz, err := newForTest()
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	stocks, err := biz.ListStocks(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ctx.Debug("stocks", zap.Any("stocks", stocks))
}

func TestImpl_GetStockBySymbol(t *testing.T) {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	biz, err := newForTest()
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	stock, err := biz.GetStockBySymbol(ctx, "2330")
	if err != nil {
		t.Fatal(err)
	}

	ctx.Debug("stock", zap.Any("stock", stock))
}

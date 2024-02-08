//go:build external

package influxdb_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/market/repo/influxdb"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	influxdbx "github.com/blackhorseya/monorepo-go/pkg/storage/influxdb"
	"go.uber.org/zap"
)

func TestImpl_GetLatestBySymbol(t *testing.T) {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	client, err := influxdbx.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	quote, err := influxdb.NewQuoteRepo(client)
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	candlestick, err := quote.GetLatestBySymbol(ctx, "2330")
	if err != nil {
		t.Fatal(err)
	}

	ctx.Info("candlestick", zap.Any("candlestick", &candlestick))
}

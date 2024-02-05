//go:build external

package mongodb_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/market/repo/mongodb"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"go.uber.org/zap"
)

func TestImpl_GetBySymbol(t *testing.T) {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	rw, err := mongodbx.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	storager, err := mongodb.NewStorager(rw)
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	stock, err := storager.GetBySymbol(ctx, "2330")
	if err != nil {
		t.Fatal(err)
	}
	ctx.Debug("stock", zap.Any("stock", &stock))

	t.Skip("skip test for now")
}

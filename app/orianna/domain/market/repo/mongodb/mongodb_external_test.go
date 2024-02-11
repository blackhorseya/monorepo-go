//go:build external

package mongodb

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.uber.org/zap"
)

func TestImpl_BulkUpsertInfo(t *testing.T) {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	if err != nil {
		t.Fatal(err)
	}

	configx.ReplaceApplication(configx.C.Orianna)

	repo, err := newStockRepoForTest()
	if err != nil {
		t.Fatal(err)
	}

	ctx := contextx.Background()
	err = repo.BulkUpsertInfo(ctx, []agg.Stock{
		agg.NewStock(&model.Stock{Symbol: "1234"}),
		agg.NewStock(&model.Stock{Symbol: "5678"}),
	})
	if err != nil {
		t.Fatal(err)
	}
}

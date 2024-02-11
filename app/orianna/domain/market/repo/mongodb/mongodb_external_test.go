//go:build external

package mongodb

import (
	"errors"
	"testing"
	"time"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	marketR "github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	repo marketR.IStockRepo
}

func (s *suiteExternal) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	s.Require().NoError(err)

	configx.ReplaceApplication(configx.C.Orianna)

	repo, err := newForTest()
	s.Require().NoError(err)
	s.repo = repo
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) TestImpl_BulkUpsertInfo() {
	ctx := contextx.Background()
	err := s.repo.BulkUpsertInfo(ctx, []agg.Stock{
		agg.NewStock(&model.Stock{Symbol: "1234", Name: "黑松", IndustryCategory: "食品工業", ExchangeName: "twse"}),
	})
	s.Require().NoError(err)
}

func (s *suiteExternal) TestImpl_List() {
	ctx := contextx.Background()

	stocks, err := s.repo.List(ctx)
	s.Require().NoError(err)

	if len(stocks) == 0 {
		s.Error(errors.New("empty stocks"))
	}
}

func (s *suiteExternal) TestImpl_BulkUpdateQuota() {
	ctx := contextx.Background()
	err := s.repo.BulkUpdateQuota(ctx, []agg.Stock{
		agg.NewStockWithQuota(&model.Stock{Symbol: "1234"}, model.NewStockQuota(0, 0, 0, 0, 0, time.Now())),
		agg.NewStockWithQuota(&model.Stock{Symbol: "5678"}, model.NewStockQuota(0, 0, 0, 0, 0, time.Now())),
	})
	s.Require().NoError(err)
}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	marketB "github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	stocks repo.IStockRepo
}

// NewMarketBiz will create a new StockBiz.
func NewMarketBiz(stocks repo.IStockRepo) (marketB.IMarketBiz, error) {
	return &impl{
		stocks: stocks,
	}, nil
}

func (i *impl) ListStocks(ctx contextx.Contextx) ([]agg.Stock, error) {
	ret, err := i.stocks.List(ctx)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) GetStockBySymbol(ctx contextx.Contextx, symbol string) (agg.Stock, error) {
	// TODO implement me
	panic("implement me")
}

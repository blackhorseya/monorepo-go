package biz

import (
	"time"

	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
)

type impl struct {
	finmind finmindx.Dialer
	storage repo.Storager
}

// NewMarketBiz is used to create a new market biz instance.
func NewMarketBiz(finmind finmindx.Dialer, storage repo.Storager) (biz.IMarketBiz, error) {
	return &impl{
		finmind: finmind,
		storage: storage,
	}, nil
}

func (i *impl) GetMarketInfoByType(
	ctx contextx.Contextx,
	typeStr string,
	t *time.Time,
) (info *model.MarketInfo, err error) {
	// todo: 2024/2/8|sean|implement me
	panic("implement me")
}

func (i *impl) GetStockBySymbol(ctx contextx.Contextx, symbol string) (stock *model.Stock, err error) {
	info, err := i.storage.GetBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	ret := &model.Stock{
		Symbol: symbol,
		Name:   info.Name,
		Price:  0,
	}

	return ret, nil
}

func (i *impl) ListStocks(
	ctx contextx.Contextx,
	options biz.ListStocksOptions,
) (stocks []*model.StockInfo, total int, err error) {
	// todo: 2024/2/7|sean|implement me
	panic("implement me")
}

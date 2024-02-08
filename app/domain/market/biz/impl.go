package biz

import (
	"time"

	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type impl struct {
	finmind finmindx.Dialer
	storage repo.Storager
	quote   repo.IQuoteRepo
}

// NewMarketBiz is used to create a new market biz instance.
func NewMarketBiz(finmind finmindx.Dialer, storage repo.Storager, quote repo.IQuoteRepo) (biz.IMarketBiz, error) {
	return &impl{
		finmind: finmind,
		storage: storage,
		quote:   quote,
	}, nil
}

func (i *impl) GetMarketInfoByType(
	ctx contextx.Contextx,
	typeStr string,
	t time.Time,
) (info *model.MarketInfo, err error) {
	ret := &model.MarketInfo{
		Type:       typeStr,
		Name:       "",
		QueriedAt:  timestamppb.New(t),
		IsTradeDay: true,
		IsOpening:  false,
	}

	weekday := t.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		ret.IsTradeDay = false
	}

	return ret, nil
}

func (i *impl) GetStockBySymbol(ctx contextx.Contextx, symbol string) (stock *model.Stock, err error) {
	info, err := i.storage.GetBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	ret := &model.Stock{
		Symbol: symbol,
		Name:   info.Name,
	}

	quote, err := i.quote.GetLatestBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	ret.Price = quote.Close

	return ret, nil
}

func (i *impl) ListStocks(
	ctx contextx.Contextx,
	options biz.ListStocksOptions,
) (stocks []*model.StockInfo, total int, err error) {
	// todo: 2024/2/7|sean|implement me
	panic("implement me")
}

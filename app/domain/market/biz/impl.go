package biz

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
)

type impl struct {
	finmind finmindx.Dialer
}

// NewMarketBiz is used to create a new market biz instance.
func NewMarketBiz(finmind finmindx.Dialer) (biz.IMarketBiz, error) {
	return &impl{
		finmind: finmind,
	}, nil
}

func (i *impl) GetStockBySymbol(ctx contextx.Contextx, symbol string) (stock *model.Stock, err error) {
	ret := &model.Stock{
		Symbol: symbol,
		Name:   "",
		Price:  0,
	}

	var got *finmindx.TaiwanStockPriceResponse
	var retryCount int
	now := time.Now()
	for ret.Price == 0 && retryCount < 5 {
		got, err = i.finmind.TaiwanStockPrice(ctx, symbol, now, now)
		if err != nil {
			return nil, err
		}

		if len(got.Data) > 0 {
			ret.Price = got.Data[0].Close
		}

		now = now.Add(-24 * time.Hour)
		retryCount++
	}

	return ret, nil
}

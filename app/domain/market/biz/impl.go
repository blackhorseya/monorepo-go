package biz

import (
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
	// todo: 2024/2/4|sean|implement me
	panic("implement me")
}

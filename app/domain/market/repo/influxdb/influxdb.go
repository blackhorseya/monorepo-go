package influxdb

import (
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	client *influxdb3.Client
}

// NewQuoteRepo is the factory function to create a new quote repository.
func NewQuoteRepo(client *influxdb3.Client) (repo.IQuoteRepo, error) {
	return &impl{client: client}, nil
}

func (i *impl) GetLatestBySymbol(ctx contextx.Contextx, symbol string) (quote *model.Candlestick, err error) {
	// todo: 2024/2/8|sean|implement me
	panic("implement me")
}

package influxdb

import (
	"fmt"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	timeoutDuration = 5 * time.Second
	dbName          = "stock_quotes"
)

type impl struct {
	client *influxdb3.Client
}

// NewQuoteRepo is the factory function to create a new quote repository.
func NewQuoteRepo(client *influxdb3.Client) (repo.IQuoteRepo, error) {
	return &impl{client: client}, nil
}

func (i *impl) GetLatestBySymbol(ctx contextx.Contextx, symbol string) (quote *model.Candlestick, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	stmt := fmt.Sprintf(`
SELECT *
FROM "quotes"
WHERE
"symbol" = '%s'
ORDER BY time DESC
LIMIT 1
`, symbol)
	opts := &influxdb3.QueryOptions{Database: dbName}
	iterator, err := i.client.QueryWithOptions(timeout, opts, stmt)
	if err != nil {
		return nil, err
	}

	ret := &model.Candlestick{
		Symbol: symbol,
	}
	for iterator.Next() {
		row := iterator.AsPoints()
		ret = &model.Candlestick{
			Symbol:      symbol,
			Open:        *row.GetDoubleField("open"),
			High:        *row.GetDoubleField("high"),
			Low:         *row.GetDoubleField("low"),
			Close:       *row.GetDoubleField("close"),
			Volume:      *row.GetIntegerField("volume"),
			Value:       *row.GetIntegerField("value"),
			Change:      *row.GetDoubleField("change"),
			Transaction: *row.GetDoubleField("transaction"),
			StartAt:     nil,
			EndAt:       timestamppb.New(row.Timestamp),
		}
	}

	return ret, nil
}

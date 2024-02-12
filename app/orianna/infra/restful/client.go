package restful

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	endpoint string
}

// NewClient is to create a new orianna client.
func NewClient() (Dialer, error) {
	return &impl{
		endpoint: configx.C.Orianna.HTTP.URL,
	}, nil
}

func (i *impl) GetStockBySymbol(ctx contextx.Contextx, symbol string) (agg.Stock, error) {
	uri, err := url.ParseRequestURI(i.endpoint + "/api/v1/stocks/" + symbol)
	if err != nil {
		return agg.Stock{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return agg.Stock{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return agg.Stock{}, err
	}
	defer resp.Body.Close()

	var got getStockBySymbolResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return agg.Stock{}, err
	}

	if got.Code != http.StatusOK {
		return agg.Stock{}, errors.New(got.Message)
	}

	return *got.Data, nil
}

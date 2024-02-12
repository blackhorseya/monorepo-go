package restful

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/response"
)

type getStockBySymbolResponse struct {
	*response.Response `json:",inline"`
	Data               *agg.Stock `json:"data,omitempty"`
}

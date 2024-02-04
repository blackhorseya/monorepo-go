package finmindx

// Response is used to represent the response.
type Response struct {
	Message string      `json:"msg"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// TaiwanStockPrice is used to represent the Taiwan stock price.
type TaiwanStockPrice struct {
	Date            string  `json:"date"`
	StockID         string  `json:"stock_id"`
	TradingVolume   int64   `json:"Trading_Volume"`
	TradingMoney    int64   `json:"Trading_money"`
	Open            float64 `json:"open"`
	Max             float64 `json:"max"`
	Min             float64 `json:"min"`
	Close           float64 `json:"close"`
	Spread          float64 `json:"spread"`
	TradingTurnover float64 `json:"Trading_turnover"`
}

// TaiwanStockPriceResponse is used to represent the Taiwan stock price response.
type TaiwanStockPriceResponse struct {
	*Response `json:",inline"`
	Data      []*TaiwanStockPrice `json:"data,omitempty"`
}

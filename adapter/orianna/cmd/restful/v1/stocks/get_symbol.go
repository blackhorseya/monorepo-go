package stocks

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetStockBySymbol is used to get stock by symbol.
// @Summary Get stock by symbol
// @Description get stock by symbol
// @Tags stocks
// @Accept json
// @Produce json
// @Param symbol path string true "Stock symbol"
// @Success 200 {object} response.Response
// @Router /v1/stocks/{symbol} [get]
func (i *impl) GetStockBySymbol(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	symbol := c.Param("symbol")

	ret, err := i.svc.GetStockBySymbol(ctx, symbol)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(&ret))
}

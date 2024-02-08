package stocks

import (
	"net/http"

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
	// todo: 2024/2/8|sean|implement GetStockBySymbol
	c.JSON(http.StatusOK, response.OK.WithMessage("not implemented"))
}

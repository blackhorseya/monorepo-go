package stocks

import (
	"github.com/gin-gonic/gin"
)

// GetList is the handler for the RESTful API.
// @Summary List stocks
// @Description Get a list of stocks
// @Tags stocks
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]model.StockInfo}
// @Failure 500 {object} response.Response
// @Router /v1/stocks/ [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/2/5|sean|implement me
	panic("implement me")
}

package stats

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type impl struct {
	svc biz.IMarketBiz
}

// Handle is the handler for the RESTful API.
func Handle(g *gin.RouterGroup, svc biz.IMarketBiz) {
	instance := &impl{
		svc: svc,
	}

	g.GET("/info", instance.GetInfo)
}

// GetInfo is the handler for the RESTful API.
// @Summary Get Stats
// @Description Get Stats
// @Tags stats
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/stats/info [get]
func (i *impl) GetInfo(c *gin.Context) {
	// todo: 2024/2/8|sean|implement the logic
	c.JSON(http.StatusOK, response.Err.WithMessage("Not implemented"))
}

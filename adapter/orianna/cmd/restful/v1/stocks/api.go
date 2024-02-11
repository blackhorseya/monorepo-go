package stocks

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/biz"
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

	g.GET("/:symbol", instance.GetStockBySymbol)
	g.GET("", instance.GetList)
}

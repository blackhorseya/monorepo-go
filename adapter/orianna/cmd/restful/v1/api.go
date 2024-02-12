package v1

import (
	"github.com/blackhorseya/monorepo-go/adapter/orianna/cmd/restful/v1/stocks"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/biz"
	"github.com/gin-gonic/gin"
)

// Handle is the handler for the RESTful API.
func Handle(g *gin.RouterGroup, svc biz.IMarketBiz) {
	stocks.Handle(g.Group("/stocks"), svc)
}

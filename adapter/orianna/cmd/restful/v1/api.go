package v1

import (
	"github.com/blackhorseya/monorepo-go/adapter/orianna/cmd/restful/v1/stocks"
	"github.com/gin-gonic/gin"
)

// Handle is the handler for the RESTful API.
func Handle(g *gin.RouterGroup) {
	stocks.Handle(g.Group("/stocks"))
}

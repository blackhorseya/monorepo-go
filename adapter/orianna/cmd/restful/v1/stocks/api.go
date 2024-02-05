package stocks

import (
	"github.com/gin-gonic/gin"
)

type impl struct {
}

// Handle is the handler for the RESTful API.
func Handle(g *gin.RouterGroup) {
	instance := &impl{}

	g.GET("/", instance.GetList)
}

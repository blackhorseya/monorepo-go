package v1

import (
	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd/restful/v1/redpackets"
	"github.com/gin-gonic/gin"
)

// Handle register redpacket api.
func Handle(g *gin.RouterGroup) {
	redpackets.Handle(g.Group("/redpackets"))
}

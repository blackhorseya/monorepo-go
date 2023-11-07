package v1

import (
	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd/restful/v1/redpackets"
	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	"github.com/gin-gonic/gin"
)

// Handle register redpacket api.
func Handle(g *gin.RouterGroup, svc eventB.IEventBiz) {
	redpackets.Handle(g.Group("/redpackets"), svc)
}

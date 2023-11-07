package redpackets

import (
	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/gin-gonic/gin"
)

// Handle register redpacket api.
func Handle(g *gin.RouterGroup, svc eventB.IEventBiz) {
	g.GET("/", gin.WrapH(MakeListRedPacketHandler(contextx.Background(), endpoints.MakeListRedPacketEndpoint(svc))))
	g.POST("/", gin.WrapH(MakeCreateRedPacketHandler(contextx.Background(), endpoints.MakeCreateRedPacketEndpoint())))
}

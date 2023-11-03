package redpackets

import (
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/gin-gonic/gin"
)

// Handle register redpacket api.
func Handle(g *gin.RouterGroup) {
	g.GET("/", gin.WrapH(MakeListRedPacketHandler(contextx.Background(), endpoints.MakeListRedPacketEndpoint())))
	g.POST("/", gin.WrapH(MakeCreateRedPacketHandler(contextx.Background(), endpoints.MakeCreateRedPacketEndpoint())))
}

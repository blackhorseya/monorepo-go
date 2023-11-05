package stringx

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, svc biz.IStringBiz) {
	// register POST /uppercase
	g.POST("/uppercase", gin.WrapH(MakeUppercaseHandler(contextx.Background(), endpoints.MakeUppercaseEndpoint(svc))))

	// register POST /count
	g.POST("/count", gin.WrapH(MakeCountHandler(contextx.Background(), endpoints.MakeCountEndpoint(svc))))
}

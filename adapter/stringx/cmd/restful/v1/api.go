package v1

import (
	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/restful/v1/stringx"
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/gin-gonic/gin"
)

// Handle will handle the stringx api.
func Handle(g *gin.RouterGroup, svc biz.IStringBiz) {
	stringx.Handle(g.Group("/stringx"), svc)
}

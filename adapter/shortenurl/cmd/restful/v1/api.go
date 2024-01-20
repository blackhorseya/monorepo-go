package v1

import (
	"github.com/blackhorseya/monorepo-go/adapter/shortenurl/cmd/restful/v1/shorten"
	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/gin-gonic/gin"
)

// Handle will handle the shortenurl api.
func Handle(g *gin.RouterGroup, svc shortB.IShorteningBiz) {
	shorten.Handle(g.Group("/shorten"), svc)
}

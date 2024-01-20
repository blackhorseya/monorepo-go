package shorten

import (
	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	svc shortB.IShorteningBiz
}

// Handle will handle the shortenurl api.
func Handle(g *gin.RouterGroup, svc shortB.IShorteningBiz) {
	instance := &impl{
		svc: svc,
	}

	g.POST("", instance.PostURL)
}

// PostURL will handle the post url request.
func (i *impl) PostURL(c *gin.Context) {
	// todo: 2024/1/20|sean|implement me
	panic("implement me")
}

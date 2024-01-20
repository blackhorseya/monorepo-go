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
// @Summary Shorten a URL
// @Description shorten a url
// @Tags shorten
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/shorten [post]
func (i *impl) PostURL(c *gin.Context) {
	// todo: 2024/1/20|sean|implement me
	panic("implement me")
}

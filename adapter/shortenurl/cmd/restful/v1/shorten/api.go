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

	g.GET("/:short_url", instance.GetShortenURL)
	g.POST("", instance.PostShortenURL)
}

// GetShortenURL will get the shorten url.
// @Summary Get a URL
// @Description get a url
// @Tags shorten
// @Accept json
// @Produce json
// @Param short_url path string true "short url"
// @Success 200 {object} response.Response
// @Router /v1/shorten/{short_url} [get]
func (i *impl) GetShortenURL(c *gin.Context) {
	// todo: 2024/1/28|sean|implement me
	panic("implement me")
}

// PostShortenURL will shorten the url.
// @Summary Shorten a URL
// @Description shorten a url
// @Tags shorten
// @Accept json
// @Produce json
// @Param request body endpoints.CreateShortURLRequest true "shorten url request"
// @Success 200 {object} response.Response
// @Router /v1/shorten [post]
func (i *impl) PostShortenURL(c *gin.Context) {
	// todo: 2024/1/28|sean|implement me
	panic("implement me")
}

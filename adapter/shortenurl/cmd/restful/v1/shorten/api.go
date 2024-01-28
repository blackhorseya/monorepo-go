package shorten

import (
	"errors"
	"net/http"

	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
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
	ctx, ok := c.MustGet(contextx.KeyCtx).(contextx.Contextx)
	if !ok {
		_ = c.Error(errors.New("invalid contextx"))
		return
	}

	record, err := i.svc.GetURLRecordByShortURL(ctx, c.Param("short_url"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(record))
}

// PostShortenURLPayload is the request for post shorten url.
type PostShortenURLPayload struct {
	URL string `json:"url"`
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
	ctx, ok := c.MustGet(contextx.KeyCtx).(contextx.Contextx)
	if !ok {
		_ = c.Error(errors.New("invalid contextx"))
		return
	}

	var payload PostShortenURLPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	record, err := i.svc.CreateShortenedURL(ctx, payload.URL)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(record))
}

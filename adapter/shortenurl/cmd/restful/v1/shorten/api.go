package shorten

import (
	"context"
	"encoding/json"
	"net/http"

	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/endpoints"
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

func decodePostURLRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateShortURLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// PostURL will handle the post url request.
// @Summary Shorten a URL
// @Description shorten a url
// @Tags shorten
// @Accept json
// @Produce json
// @Param request body endpoints.CreateShortURLRequest true "shorten url request"
// @Success 200 {object} response.Response
// @Router /v1/shorten [post]
func (i *impl) PostURL(c *gin.Context) {
	// todo: 2024/1/20|sean|implement me
	panic("implement me")
}

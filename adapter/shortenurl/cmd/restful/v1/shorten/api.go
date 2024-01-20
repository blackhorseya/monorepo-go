package shorten

import (
	"context"
	"encoding/json"
	"net/http"

	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

type impl struct {
	svc shortB.IShorteningBiz
}

// Handle will handle the shortenurl api.
func Handle(g *gin.RouterGroup, svc shortB.IShorteningBiz) {
	g.POST("", gin.WrapH(MakePostURLHandler(svc)))
}

func decodePostURLRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateShortURLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// MakePostURLHandler will handle the post url request.
// @Summary Shorten a URL
// @Description shorten a url
// @Tags shorten
// @Accept json
// @Produce json
// @Param request body endpoints.CreateShortURLRequest true "shorten url request"
// @Success 200 {object} response.Response
// @Router /v1/shorten [post]
func MakePostURLHandler(svc shortB.IShorteningBiz) http.Handler {
	return httptransport.NewServer(
		endpoints.MakeCreateShortURLEndpoint(svc),
		decodePostURLRequest,
		response.EncodeJSON,
	)
}

package restful

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeCountHandler count handler.
// @Summary count
// @Description count
// @Tags String
// @Accept json
// @Produce json
// @Param request body endpoints.CountRequest true "count request"
// @Success 200 {object} endpoints.CountResponse
// @Failure 500 {object} endpoints.CountResponse
// @Router /v1/string/count [post]
func MakeCountHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeCountRequest, encodeResponse)
}

package restful

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeUppercaseHandler uppercase handler.
// @Summary uppercase
// @Description uppercase
// @Tags String
// @Accept json
// @Produce json
// @Param request body endpoints.UppercaseRequest true "uppercase request"
// @Success 200 {object} endpoints.UppercaseResponse
// @Failure 400 {object} endpoints.UppercaseResponse
// @Failure 500 {object} endpoints.UppercaseResponse
// @Router /v1/string/uppercase [post]
func MakeUppercaseHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeUppercaseRequest, encodeResponse)
}

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

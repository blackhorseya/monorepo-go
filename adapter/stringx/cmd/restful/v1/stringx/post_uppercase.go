package stringx

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blackhorseya/monorepo-go/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeUppercaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UppercaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

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
	return httptransport.NewServer(endpoint, decodeUppercaseRequest, response.EncodeJSON)
}

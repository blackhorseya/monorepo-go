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

func decodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
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
	return httptransport.NewServer(endpoint, decodeCountRequest, response.EncodeJSON)
}

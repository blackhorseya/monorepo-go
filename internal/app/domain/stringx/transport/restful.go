package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
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

func decodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// MakeUppercaseHandler uppercase handler.
func MakeUppercaseHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeUppercaseRequest, encodeResponse)
}

// MakeCountHandler count handler.
func MakeCountHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeCountRequest, encodeResponse)
}

package restful

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
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

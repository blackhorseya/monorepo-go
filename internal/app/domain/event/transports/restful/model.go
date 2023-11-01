package restful

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/endpoints"
)

func decodeCreateRedPacketRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateRedPacketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeListRedPacketRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ListRedPacketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

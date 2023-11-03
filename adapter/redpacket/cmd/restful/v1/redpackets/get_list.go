package redpackets

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeListRedPacketRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ListRedPacketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// MakeListRedPacketHandler list red packet handler.
func MakeListRedPacketHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeListRedPacketRequest, response.EncodeJSON)
}

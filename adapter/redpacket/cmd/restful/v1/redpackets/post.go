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

func decodeCreateRedPacketRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateRedPacketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// MakeCreateRedPacketHandler create red packet handler.
// @Summary create red packet
// @Description create red packet
// @Tags RedPacket
// @Accept json
// @Produce json
// @Param payload body endpoints.CreateRedPacketRequest true "create red packet request"
// @Success 200 {object} endpoints.CreateRedPacketResponse
// @Failure 400 {object} endpoints.CreateRedPacketResponse
// @Failure 500 {object} endpoints.CreateRedPacketResponse
// @Router /api/v1/redpackets [post]
func MakeCreateRedPacketHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeCreateRedPacketRequest, response.EncodeJSON)
}

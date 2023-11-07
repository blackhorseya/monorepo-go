package redpackets

import (
	"context"
	"net/http"

	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/endpoints"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeListRedPacketRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ListRedPacketRequest

	return req, nil
}

// MakeListRedPacketHandler list red packet handler.
// @Summary list red packet
// @Description list red packet
// @Tags RedPacket
// @Accept json
// @Produce json
// @Success 200 {object} endpoints.ListRedPacketResponse
// @Failure 400 {object} endpoints.ListRedPacketResponse
// @Failure 500 {object} endpoints.ListRedPacketResponse
// @Router /v1/redpackets [get]
func MakeListRedPacketHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeListRedPacketRequest, response.EncodeJSON)
}

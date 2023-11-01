package restful

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeCreateRedPacketHandler create red packet handler.
func MakeCreateRedPacketHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeCreateRedPacketRequest, encodeResponse)
}

// MakeListRedPacketHandler list red packet handler.
func MakeListRedPacketHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeListRedPacketRequest, encodeResponse)
}

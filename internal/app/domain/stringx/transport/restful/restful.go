package restful

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeUppercaseHandler uppercase handler.
func MakeUppercaseHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeUppercaseRequest, encodeResponse)
}

// MakeCountHandler count handler.
func MakeCountHandler(ctx contextx.Contextx, endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(endpoint, decodeCountRequest, encodeResponse)
}

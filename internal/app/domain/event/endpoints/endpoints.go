package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// ListRedPacketRequest list red packet request struct.
type ListRedPacketRequest struct {
}

// ListRedPacketResponse list red packet response struct.
type ListRedPacketResponse struct {
}

// MakeListRedPacketEndpoint make list red packet endpoint.
func MakeListRedPacketEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		// todo: 2023/10/30|sean|implement list red packet endpoint
		return ListRedPacketResponse{}, errors.New("implement me")
	}
}

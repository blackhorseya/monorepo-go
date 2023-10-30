package endpoints

import (
	"context"
	"errors"

	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/go-kit/kit/endpoint"
)

// ListRedPacketRequest list red packet request struct.
type ListRedPacketRequest struct {
	Page uint32 `json:"page,omitempty"`
	Size uint32 `json:"size,omitempty"`
}

// ListRedPacketResponse list red packet response struct.
type ListRedPacketResponse struct {
	response.Response `json:",inline"`
}

// MakeListRedPacketEndpoint make list red packet endpoint.
func MakeListRedPacketEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		// todo: 2023/10/30|sean|implement list red packet endpoint
		return ListRedPacketResponse{}, errors.New("implement me")
	}
}

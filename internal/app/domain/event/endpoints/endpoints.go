package endpoints

import (
	"context"
	"errors"

	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
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

// CreateRedPacketRequest create red packet request struct.
type CreateRedPacketRequest struct {
	Who    *userM.UserAccount `json:"who,omitempty"`
	Amount uint64             `json:"amount,omitempty"`
	Count  uint32             `json:"count,omitempty"`
}

// CreateRedPacketResponse create red packet response struct.
type CreateRedPacketResponse struct {
	response.Response `json:",inline"`
}

// MakeCreateRedPacketEndpoint make create red packet endpoint.
func MakeCreateRedPacketEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		// todo: 2023/11/1|sean|implement create red packet endpoint
		return CreateRedPacketResponse{}, errors.New("implement me")
	}
}

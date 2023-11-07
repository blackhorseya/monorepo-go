package endpoints

import (
	"context"
	"errors"

	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
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
	*response.Response `json:",inline"`
}

// MakeListRedPacketEndpoint make list red packet endpoint.
func MakeListRedPacketEndpoint(svc eventB.IEventBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (resp interface{}, err error) {
		req, _ := request.(ListRedPacketRequest)
		ctx := contextx.Background()

		ret, err := svc.ListRedPacket(ctx, eventB.ListRedPacketCondition{
			Page:    req.Page,
			PerPage: req.Size,
		})
		if err != nil {
			return ListRedPacketResponse{
				Response: &response.Response{Message: err.Error()},
			}, err
		}

		return ListRedPacketResponse{
			Response: response.OK.WithData(ret),
		}, nil
	}
}

// CreateRedPacketRequest create red packet request struct.
type CreateRedPacketRequest struct {
	Who    *model.UserAccount `json:"who,omitempty"`
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

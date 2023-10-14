package endpoints

import (
	"context"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/go-kit/kit/endpoint"
)

// UppercaseRequest uppercase request struct.
type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func MakeUppercaseEndpoint(svc biz.IStringBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(UppercaseRequest)
		ctx := contextx.Background()

		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}

		return UppercaseResponse{v, ""}, nil
	}
}

// CountRequest count request struct.
type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

func MakeCountEndpoint(svc biz.IStringBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(CountRequest)
		ctx := contextx.Background()

		v := svc.Count(ctx, req.S)

		return CountResponse{v}, nil
	}
}

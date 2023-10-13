package biz

import (
	"context"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/go-kit/kit/endpoint"
)

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func makeUppercaseEndpoint(svc biz.IStringBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(uppercaseRequest)
		ctx := contextx.Background()

		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}

		return uppercaseResponse{v, ""}, nil
	}
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeCountEndpoint(svc biz.IStringBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(countRequest)
		ctx := contextx.Background()

		v := svc.Count(ctx, req.S)

		return countResponse{v}, nil
	}
}

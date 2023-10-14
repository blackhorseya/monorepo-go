package grpc

import (
	"context"
	"errors"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
)

func decodeToUpperRequest(c context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*model.ToUpperRequest)
	if !ok {
		return nil, errors.New("grpc server decode to upper request error")
	}

	return &endpoints.UppercaseRequest{
		S: req.Value,
	}, nil
}

func encodeToUpperResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(*endpoints.UppercaseResponse)
	if !ok {
		return nil, errors.New("grpc server encode to upper response error")
	}

	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}

	return &model.ToUpperResponse{
		Value: resp.V,
	}, nil
}

func decodeCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*model.CountRequest)
	if !ok {
		return nil, errors.New("grpc server decode count request error")
	}

	return &endpoints.CountRequest{
		S: req.Value,
	}, nil
}

func encodeCountResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(*endpoints.CountResponse)
	if !ok {
		return nil, errors.New("grpc server encode count response error")
	}

	return &model.CountResponse{
		Value: int32(resp.V),
	}, nil
}

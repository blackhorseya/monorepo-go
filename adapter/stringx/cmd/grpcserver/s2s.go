package grpcserver

import (
	"context"
	"errors"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/grpc/_grpc_test/pb"
)

type s2s struct {
	pb.UnimplementedTestServer

	toUpper grpctransport.Handler
	count   grpctransport.Handler
}

// NewServer returns a new stringx service server.
func NewServer(uppercase, count endpoint.Endpoint) model.StringxServiceServer {
	return &s2s{
		toUpper: grpctransport.NewServer(uppercase, decodeToUpperRequest, encodeToUpperResponse),
		count:   grpctransport.NewServer(count, decodeCountRequest, encodeCountResponse),
	}
}

func (s *s2s) ToUpper(ctx context.Context, request *model.ToUpperRequest) (*model.ToUpperResponse, error) {
	_, resp, err := s.toUpper.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ToUpperResponse), nil
}

func (s *s2s) Count(ctx context.Context, request *model.CountRequest) (*model.CountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*model.CountResponse), nil
}

func decodeToUpperRequest(c context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*model.ToUpperRequest)
	if !ok {
		return nil, errors.New("grpc server decode to upper request error")
	}

	return endpoints.UppercaseRequest{
		S: req.Value,
	}, nil
}

func encodeToUpperResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoints.UppercaseResponse)
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

	return endpoints.CountRequest{
		S: req.Value,
	}, nil
}

func encodeCountResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoints.CountResponse)
	if !ok {
		return nil, errors.New("grpc server encode count response error")
	}

	return &model.CountResponse{
		Value: int32(resp.V),
	}, nil
}

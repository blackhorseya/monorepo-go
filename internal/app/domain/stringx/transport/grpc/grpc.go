package grpc

import (
	"context"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/model"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/grpc/_grpc_test/pb"
)

type grpc struct {
	pb.UnimplementedTestServer

	toUpper grpctransport.Handler
	count   grpctransport.Handler
}

// New grpc transport.
func New(toUpper, count endpoint.Endpoint) model.StringxServiceServer {
	return &grpc{
		toUpper: grpctransport.NewServer(toUpper, decodeToUpperRequest, encodeToUpperResponse),
		count:   grpctransport.NewServer(count, decodeCountRequest, encodeCountResponse),
	}
}

func (g *grpc) ToUpper(ctx context.Context, request *model.ToUpperRequest) (*model.ToUpperResponse, error) {
	_, resp, err := g.toUpper.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ToUpperResponse), nil
}

func (g *grpc) Count(ctx context.Context, request *model.CountRequest) (*model.CountResponse, error) {
	_, resp, err := g.count.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*model.CountResponse), nil
}

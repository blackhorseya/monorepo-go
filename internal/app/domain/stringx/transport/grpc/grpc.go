package grpc

import (
	"context"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/model"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/grpc/_grpc_test/pb"
)

type grpc struct {
	pb.UnimplementedTestServer

	toUpper grpctransport.Handler
	count   grpctransport.Handler
}

// New grpc transport.
func New() model.StringxServiceServer {
	return &grpc{}
}

func (g *grpc) ToUpper(ctx context.Context, request *model.ToUpperRequest) (*model.ToUpperResponse, error) {
	// todo: 2023/10/14|sean|impl me
	panic("implement me")
}

func (g *grpc) Count(ctx context.Context, request *model.CountRequest) (*model.CountResponse, error) {
	// todo: 2023/10/14|sean|impl me
	panic("implement me")
}

package endpoints

import (
	"context"
	"errors"

	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	responsex "github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/go-kit/kit/endpoint"
)

// CreateShortURLRequest create short url request struct.
type CreateShortURLRequest struct {
	URL string `json:"url"`
}

// MakeCreateShortURLEndpoint returns an endpoint via the passed service.
func MakeCreateShortURLEndpoint(svc biz.IShorteningBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		ctx := contextx.Background()

		req, ok := request.(CreateShortURLRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		record, err := svc.CreateShortenedURL(ctx, req.URL)
		if err != nil {
			return responsex.Err.WrapError(err), nil
		}

		return responsex.OK.WithData(record), nil
	}
}

// GetShortURLRequest get short url request struct.
type GetShortURLRequest struct {
	URL string `json:"url"`
}

// MakeGetShortURLEndpoint returns an endpoint via the passed service.
func MakeGetShortURLEndpoint(svc biz.IShorteningBiz) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		ctx := contextx.Background()

		req, ok := request.(GetShortURLRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		record, err := svc.GetURLRecordByShortURL(ctx, req.URL)
		if err != nil {
			return responsex.Err.WrapError(err), nil
		}

		return responsex.OK.WithData(record), nil
	}
}

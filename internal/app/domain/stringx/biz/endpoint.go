package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
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
	// todo: 2023/10/13|sean|impl me
	panic("implement me")
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeCountEndpoint(svc biz.IStringBiz) endpoint.Endpoint {
	// todo: 2023/10/13|sean|impl me
	panic("implement me")
}

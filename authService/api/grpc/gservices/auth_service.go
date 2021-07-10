package gservices

import (
	"auth-service/api/grpc/pbs"
	"context"
	"google.golang.org/grpc"
)

type AuthService struct {
	pbs.UnimplementedAuthServiceServer
}

func (a AuthService) CheckToken(ctx context.Context, in *pbs.CheckTokenParams, opts ...grpc.CallOption) (*pbs.CheckTokenResp, error) {
	if in.GetToken() == "" {
		return &pbs.CheckTokenResp{Ok: true}, nil
	}
	return &pbs.CheckTokenResp{Ok: false}, nil
}

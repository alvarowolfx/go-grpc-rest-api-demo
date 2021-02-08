package endpoints

import (
	"context"

	"com.aviebrantz.tvtime/internal/application/auth"
	tvtime "com.aviebrantz.tvtime/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthEndpoint interface {
	tvtime.AuthServiceServer
}

type AuthEndpointDeps struct {
	AuthService auth.AuthService
}

type authEndpoint struct {
	AuthEndpointDeps
}

func NewAuthEndpoint(deps AuthEndpointDeps) AuthEndpoint {
	return &authEndpoint{deps}
}

func (e *authEndpoint) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// No Auth
	return ctx, nil
}

func (e *authEndpoint) Login(ctx context.Context, req *tvtime.LoginRequest) (*tvtime.LoginResponse, error) {
	token, err := e.AuthService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "invalid user credential: %v", err)
	}
	return &tvtime.LoginResponse{
		Token: token,
	}, nil
}

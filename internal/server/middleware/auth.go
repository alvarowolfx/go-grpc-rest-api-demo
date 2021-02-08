package middleware

import (
	"context"
	"strings"

	"com.aviebrantz.tvtime/internal/application/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type authContextKey string

const (
	userIDContextKey = authContextKey("userID")
)

func AuthUnaryInterceptor(as auth.AuthService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := authFunc(ctx, as, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

func authFunc(ctx context.Context, as auth.AuthService, method string) (context.Context, error) {
	if strings.HasSuffix(method, "AuthService/Login") ||
		strings.HasSuffix(method, "UserService/RegisterUser") ||
		strings.HasSuffix(method, "CatalogService/Get") ||
		strings.HasSuffix(method, "CatalogService/List") {
		// No Auth
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	err = as.Verify(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	userID, err := as.GetUserID(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	if userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "failed to identify user: %v", err)
	}

	ctx = context.WithValue(ctx, userIDContextKey, userID)

	return ctx, nil
}

func GetUserIDFromContext(ctx context.Context) string {
	userID, ok := ctx.Value(userIDContextKey).(string)
	if !ok {
		return ""
	}
	return userID
}

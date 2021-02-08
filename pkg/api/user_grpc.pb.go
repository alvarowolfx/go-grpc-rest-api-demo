// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tvtime

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	ListFavorites(ctx context.Context, in *ListFavoritesRequest, opts ...grpc.CallOption) (*ListFavoritesResponse, error)
	AddMovieToFavorites(ctx context.Context, in *AddRemoveMovieFavoriteRequest, opts ...grpc.CallOption) (*AddRemoveMovieFavoriteResponse, error)
	RemoveMovieFromFavorites(ctx context.Context, in *AddRemoveMovieFavoriteRequest, opts ...grpc.CallOption) (*AddRemoveMovieFavoriteResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/com.aviebrantz.tvtime.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListFavorites(ctx context.Context, in *ListFavoritesRequest, opts ...grpc.CallOption) (*ListFavoritesResponse, error) {
	out := new(ListFavoritesResponse)
	err := c.cc.Invoke(ctx, "/com.aviebrantz.tvtime.UserService/ListFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddMovieToFavorites(ctx context.Context, in *AddRemoveMovieFavoriteRequest, opts ...grpc.CallOption) (*AddRemoveMovieFavoriteResponse, error) {
	out := new(AddRemoveMovieFavoriteResponse)
	err := c.cc.Invoke(ctx, "/com.aviebrantz.tvtime.UserService/AddMovieToFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveMovieFromFavorites(ctx context.Context, in *AddRemoveMovieFavoriteRequest, opts ...grpc.CallOption) (*AddRemoveMovieFavoriteResponse, error) {
	out := new(AddRemoveMovieFavoriteResponse)
	err := c.cc.Invoke(ctx, "/com.aviebrantz.tvtime.UserService/RemoveMovieFromFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	ListFavorites(context.Context, *ListFavoritesRequest) (*ListFavoritesResponse, error)
	AddMovieToFavorites(context.Context, *AddRemoveMovieFavoriteRequest) (*AddRemoveMovieFavoriteResponse, error)
	RemoveMovieFromFavorites(context.Context, *AddRemoveMovieFavoriteRequest) (*AddRemoveMovieFavoriteResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) ListFavorites(context.Context, *ListFavoritesRequest) (*ListFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFavorites not implemented")
}
func (UnimplementedUserServiceServer) AddMovieToFavorites(context.Context, *AddRemoveMovieFavoriteRequest) (*AddRemoveMovieFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovieToFavorites not implemented")
}
func (UnimplementedUserServiceServer) RemoveMovieFromFavorites(context.Context, *AddRemoveMovieFavoriteRequest) (*AddRemoveMovieFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMovieFromFavorites not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.aviebrantz.tvtime.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.aviebrantz.tvtime.UserService/ListFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListFavorites(ctx, req.(*ListFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddMovieToFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRemoveMovieFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddMovieToFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.aviebrantz.tvtime.UserService/AddMovieToFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddMovieToFavorites(ctx, req.(*AddRemoveMovieFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveMovieFromFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRemoveMovieFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveMovieFromFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.aviebrantz.tvtime.UserService/RemoveMovieFromFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveMovieFromFavorites(ctx, req.(*AddRemoveMovieFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.aviebrantz.tvtime.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "ListFavorites",
			Handler:    _UserService_ListFavorites_Handler,
		},
		{
			MethodName: "AddMovieToFavorites",
			Handler:    _UserService_AddMovieToFavorites_Handler,
		},
		{
			MethodName: "RemoveMovieFromFavorites",
			Handler:    _UserService_RemoveMovieFromFavorites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

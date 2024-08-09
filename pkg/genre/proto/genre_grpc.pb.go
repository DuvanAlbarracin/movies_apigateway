// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pkg/genre/proto/genre.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GenreServiceClient is the client API for GenreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenreServiceClient interface {
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	AddToMovie(ctx context.Context, in *AddToMovieRequest, opts ...grpc.CallOption) (*AddToMovieResponse, error)
	RemoveFromMovie(ctx context.Context, in *RemoveFromMovieRequest, opts ...grpc.CallOption) (*RemoveFromMovieResponse, error)
}

type genreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenreServiceClient(cc grpc.ClientConnInterface) GenreServiceClient {
	return &genreServiceClient{cc}
}

func (c *genreServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/genre.GenreService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/genre.GenreService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) AddToMovie(ctx context.Context, in *AddToMovieRequest, opts ...grpc.CallOption) (*AddToMovieResponse, error) {
	out := new(AddToMovieResponse)
	err := c.cc.Invoke(ctx, "/genre.GenreService/AddToMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) RemoveFromMovie(ctx context.Context, in *RemoveFromMovieRequest, opts ...grpc.CallOption) (*RemoveFromMovieResponse, error) {
	out := new(RemoveFromMovieResponse)
	err := c.cc.Invoke(ctx, "/genre.GenreService/RemoveFromMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenreServiceServer is the server API for GenreService service.
// All implementations must embed UnimplementedGenreServiceServer
// for forward compatibility
type GenreServiceServer interface {
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	AddToMovie(context.Context, *AddToMovieRequest) (*AddToMovieResponse, error)
	RemoveFromMovie(context.Context, *RemoveFromMovieRequest) (*RemoveFromMovieResponse, error)
	mustEmbedUnimplementedGenreServiceServer()
}

// UnimplementedGenreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGenreServiceServer struct {
}

func (UnimplementedGenreServiceServer) GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedGenreServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGenreServiceServer) AddToMovie(context.Context, *AddToMovieRequest) (*AddToMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToMovie not implemented")
}
func (UnimplementedGenreServiceServer) RemoveFromMovie(context.Context, *RemoveFromMovieRequest) (*RemoveFromMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromMovie not implemented")
}
func (UnimplementedGenreServiceServer) mustEmbedUnimplementedGenreServiceServer() {}

// UnsafeGenreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenreServiceServer will
// result in compilation errors.
type UnsafeGenreServiceServer interface {
	mustEmbedUnimplementedGenreServiceServer()
}

func RegisterGenreServiceServer(s grpc.ServiceRegistrar, srv GenreServiceServer) {
	s.RegisterService(&GenreService_ServiceDesc, srv)
}

func _GenreService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genre.GenreService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genre.GenreService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_AddToMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).AddToMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genre.GenreService/AddToMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).AddToMovie(ctx, req.(*AddToMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_RemoveFromMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).RemoveFromMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genre.GenreService/RemoveFromMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).RemoveFromMovie(ctx, req.(*RemoveFromMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenreService_ServiceDesc is the grpc.ServiceDesc for GenreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genre.GenreService",
	HandlerType: (*GenreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _GenreService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _GenreService_GetAll_Handler,
		},
		{
			MethodName: "AddToMovie",
			Handler:    _GenreService_AddToMovie_Handler,
		},
		{
			MethodName: "RemoveFromMovie",
			Handler:    _GenreService_RemoveFromMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/genre/proto/genre.proto",
}

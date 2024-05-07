// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: podcasts.proto

package content

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

// PodcastServiceClient is the client API for PodcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodcastServiceClient interface {
	CreatePodcast(ctx context.Context, in *Podcast, opts ...grpc.CallOption) (*ResOK, error)
	UpdatePodcast(ctx context.Context, in *Podcast, opts ...grpc.CallOption) (*ResOK, error)
	DeletePodcast(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*ResOK, error)
	GetPodcast(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*Podcast, error)
}

type podcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPodcastServiceClient(cc grpc.ClientConnInterface) PodcastServiceClient {
	return &podcastServiceClient{cc}
}

func (c *podcastServiceClient) CreatePodcast(ctx context.Context, in *Podcast, opts ...grpc.CallOption) (*ResOK, error) {
	out := new(ResOK)
	err := c.cc.Invoke(ctx, "/PodcastService/CreatePodcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podcastServiceClient) UpdatePodcast(ctx context.Context, in *Podcast, opts ...grpc.CallOption) (*ResOK, error) {
	out := new(ResOK)
	err := c.cc.Invoke(ctx, "/PodcastService/UpdatePodcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podcastServiceClient) DeletePodcast(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*ResOK, error) {
	out := new(ResOK)
	err := c.cc.Invoke(ctx, "/PodcastService/DeletePodcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podcastServiceClient) GetPodcast(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*Podcast, error) {
	out := new(Podcast)
	err := c.cc.Invoke(ctx, "/PodcastService/GetPodcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodcastServiceServer is the server API for PodcastService service.
// All implementations must embed UnimplementedPodcastServiceServer
// for forward compatibility
type PodcastServiceServer interface {
	CreatePodcast(context.Context, *Podcast) (*ResOK, error)
	UpdatePodcast(context.Context, *Podcast) (*ResOK, error)
	DeletePodcast(context.Context, *ReqID) (*ResOK, error)
	GetPodcast(context.Context, *ReqID) (*Podcast, error)
	mustEmbedUnimplementedPodcastServiceServer()
}

// UnimplementedPodcastServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPodcastServiceServer struct {
}

func (UnimplementedPodcastServiceServer) CreatePodcast(context.Context, *Podcast) (*ResOK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePodcast not implemented")
}
func (UnimplementedPodcastServiceServer) UpdatePodcast(context.Context, *Podcast) (*ResOK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePodcast not implemented")
}
func (UnimplementedPodcastServiceServer) DeletePodcast(context.Context, *ReqID) (*ResOK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePodcast not implemented")
}
func (UnimplementedPodcastServiceServer) GetPodcast(context.Context, *ReqID) (*Podcast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPodcast not implemented")
}
func (UnimplementedPodcastServiceServer) mustEmbedUnimplementedPodcastServiceServer() {}

// UnsafePodcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PodcastServiceServer will
// result in compilation errors.
type UnsafePodcastServiceServer interface {
	mustEmbedUnimplementedPodcastServiceServer()
}

func RegisterPodcastServiceServer(s grpc.ServiceRegistrar, srv PodcastServiceServer) {
	s.RegisterService(&PodcastService_ServiceDesc, srv)
}

func _PodcastService_CreatePodcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Podcast)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodcastServiceServer).CreatePodcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PodcastService/CreatePodcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodcastServiceServer).CreatePodcast(ctx, req.(*Podcast))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodcastService_UpdatePodcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Podcast)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodcastServiceServer).UpdatePodcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PodcastService/UpdatePodcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodcastServiceServer).UpdatePodcast(ctx, req.(*Podcast))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodcastService_DeletePodcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodcastServiceServer).DeletePodcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PodcastService/DeletePodcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodcastServiceServer).DeletePodcast(ctx, req.(*ReqID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodcastService_GetPodcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodcastServiceServer).GetPodcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PodcastService/GetPodcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodcastServiceServer).GetPodcast(ctx, req.(*ReqID))
	}
	return interceptor(ctx, in, info, handler)
}

// PodcastService_ServiceDesc is the grpc.ServiceDesc for PodcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PodcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PodcastService",
	HandlerType: (*PodcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePodcast",
			Handler:    _PodcastService_CreatePodcast_Handler,
		},
		{
			MethodName: "UpdatePodcast",
			Handler:    _PodcastService_UpdatePodcast_Handler,
		},
		{
			MethodName: "DeletePodcast",
			Handler:    _PodcastService_DeletePodcast_Handler,
		},
		{
			MethodName: "GetPodcast",
			Handler:    _PodcastService_GetPodcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "podcasts.proto",
}
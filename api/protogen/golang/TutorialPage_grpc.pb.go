// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: TutorialPage.proto

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

// TutorialPageServiceClient is the client API for TutorialPageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TutorialPageServiceClient interface {
	CreateTutorialPage(ctx context.Context, in *TutorialPage, opts ...grpc.CallOption) (*ReqID, error)
	UpdateTutorialPage(ctx context.Context, in *TutorialPage, opts ...grpc.CallOption) (*ReqID, error)
	DeleteTutorialPage(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*ReqID, error)
	GetTutorialPage(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*TutorialPage, error)
}

type tutorialPageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTutorialPageServiceClient(cc grpc.ClientConnInterface) TutorialPageServiceClient {
	return &tutorialPageServiceClient{cc}
}

func (c *tutorialPageServiceClient) CreateTutorialPage(ctx context.Context, in *TutorialPage, opts ...grpc.CallOption) (*ReqID, error) {
	out := new(ReqID)
	err := c.cc.Invoke(ctx, "/TutorialPageService/CreateTutorialPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialPageServiceClient) UpdateTutorialPage(ctx context.Context, in *TutorialPage, opts ...grpc.CallOption) (*ReqID, error) {
	out := new(ReqID)
	err := c.cc.Invoke(ctx, "/TutorialPageService/UpdateTutorialPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialPageServiceClient) DeleteTutorialPage(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*ReqID, error) {
	out := new(ReqID)
	err := c.cc.Invoke(ctx, "/TutorialPageService/DeleteTutorialPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialPageServiceClient) GetTutorialPage(ctx context.Context, in *ReqID, opts ...grpc.CallOption) (*TutorialPage, error) {
	out := new(TutorialPage)
	err := c.cc.Invoke(ctx, "/TutorialPageService/GetTutorialPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TutorialPageServiceServer is the server API for TutorialPageService service.
// All implementations must embed UnimplementedTutorialPageServiceServer
// for forward compatibility
type TutorialPageServiceServer interface {
	CreateTutorialPage(context.Context, *TutorialPage) (*ReqID, error)
	UpdateTutorialPage(context.Context, *TutorialPage) (*ReqID, error)
	DeleteTutorialPage(context.Context, *ReqID) (*ReqID, error)
	GetTutorialPage(context.Context, *ReqID) (*TutorialPage, error)
	mustEmbedUnimplementedTutorialPageServiceServer()
}

// UnimplementedTutorialPageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTutorialPageServiceServer struct {
}

func (UnimplementedTutorialPageServiceServer) CreateTutorialPage(context.Context, *TutorialPage) (*ReqID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTutorialPage not implemented")
}
func (UnimplementedTutorialPageServiceServer) UpdateTutorialPage(context.Context, *TutorialPage) (*ReqID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTutorialPage not implemented")
}
func (UnimplementedTutorialPageServiceServer) DeleteTutorialPage(context.Context, *ReqID) (*ReqID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTutorialPage not implemented")
}
func (UnimplementedTutorialPageServiceServer) GetTutorialPage(context.Context, *ReqID) (*TutorialPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTutorialPage not implemented")
}
func (UnimplementedTutorialPageServiceServer) mustEmbedUnimplementedTutorialPageServiceServer() {}

// UnsafeTutorialPageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TutorialPageServiceServer will
// result in compilation errors.
type UnsafeTutorialPageServiceServer interface {
	mustEmbedUnimplementedTutorialPageServiceServer()
}

func RegisterTutorialPageServiceServer(s grpc.ServiceRegistrar, srv TutorialPageServiceServer) {
	s.RegisterService(&TutorialPageService_ServiceDesc, srv)
}

func _TutorialPageService_CreateTutorialPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TutorialPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialPageServiceServer).CreateTutorialPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TutorialPageService/CreateTutorialPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialPageServiceServer).CreateTutorialPage(ctx, req.(*TutorialPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TutorialPageService_UpdateTutorialPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TutorialPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialPageServiceServer).UpdateTutorialPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TutorialPageService/UpdateTutorialPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialPageServiceServer).UpdateTutorialPage(ctx, req.(*TutorialPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TutorialPageService_DeleteTutorialPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialPageServiceServer).DeleteTutorialPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TutorialPageService/DeleteTutorialPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialPageServiceServer).DeleteTutorialPage(ctx, req.(*ReqID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TutorialPageService_GetTutorialPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialPageServiceServer).GetTutorialPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TutorialPageService/GetTutorialPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialPageServiceServer).GetTutorialPage(ctx, req.(*ReqID))
	}
	return interceptor(ctx, in, info, handler)
}

// TutorialPageService_ServiceDesc is the grpc.ServiceDesc for TutorialPageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TutorialPageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TutorialPageService",
	HandlerType: (*TutorialPageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTutorialPage",
			Handler:    _TutorialPageService_CreateTutorialPage_Handler,
		},
		{
			MethodName: "UpdateTutorialPage",
			Handler:    _TutorialPageService_UpdateTutorialPage_Handler,
		},
		{
			MethodName: "DeleteTutorialPage",
			Handler:    _TutorialPageService_DeleteTutorialPage_Handler,
		},
		{
			MethodName: "GetTutorialPage",
			Handler:    _TutorialPageService_GetTutorialPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TutorialPage.proto",
}

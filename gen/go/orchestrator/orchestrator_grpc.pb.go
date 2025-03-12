// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: orchestrator.proto

package orchestrator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExpressionService_Create_FullMethodName      = "/orchestrator.ExpressionService/Create"
	ExpressionService_Get_FullMethodName         = "/orchestrator.ExpressionService/Get"
	ExpressionService_GetByUserId_FullMethodName = "/orchestrator.ExpressionService/GetByUserId"
	ExpressionService_Update_FullMethodName      = "/orchestrator.ExpressionService/Update"
	ExpressionService_Delete_FullMethodName      = "/orchestrator.ExpressionService/Delete"
)

// ExpressionServiceClient is the client API for ExpressionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpressionServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ExpressionResponse, error)
	Get(ctx context.Context, in *ExpressionRequest, opts ...grpc.CallOption) (*ExpressionResponse, error)
	GetByUserId(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetByUserIdResponse, error)
	Update(ctx context.Context, in *ExpressionRequest, opts ...grpc.CallOption) (*ExpressionResponse, error)
	Delete(ctx context.Context, in *ExpressionRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type expressionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpressionServiceClient(cc grpc.ClientConnInterface) ExpressionServiceClient {
	return &expressionServiceClient{cc}
}

func (c *expressionServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ExpressionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpressionResponse)
	err := c.cc.Invoke(ctx, ExpressionService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) Get(ctx context.Context, in *ExpressionRequest, opts ...grpc.CallOption) (*ExpressionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpressionResponse)
	err := c.cc.Invoke(ctx, ExpressionService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) GetByUserId(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetByUserIdResponse)
	err := c.cc.Invoke(ctx, ExpressionService_GetByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) Update(ctx context.Context, in *ExpressionRequest, opts ...grpc.CallOption) (*ExpressionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpressionResponse)
	err := c.cc.Invoke(ctx, ExpressionService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) Delete(ctx context.Context, in *ExpressionRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ExpressionService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressionServiceServer is the server API for ExpressionService service.
// All implementations must embed UnimplementedExpressionServiceServer
// for forward compatibility.
type ExpressionServiceServer interface {
	Create(context.Context, *CreateRequest) (*ExpressionResponse, error)
	Get(context.Context, *ExpressionRequest) (*ExpressionResponse, error)
	GetByUserId(context.Context, *GetByUserIdRequest) (*GetByUserIdResponse, error)
	Update(context.Context, *ExpressionRequest) (*ExpressionResponse, error)
	Delete(context.Context, *ExpressionRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedExpressionServiceServer()
}

// UnimplementedExpressionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExpressionServiceServer struct{}

func (UnimplementedExpressionServiceServer) Create(context.Context, *CreateRequest) (*ExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedExpressionServiceServer) Get(context.Context, *ExpressionRequest) (*ExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedExpressionServiceServer) GetByUserId(context.Context, *GetByUserIdRequest) (*GetByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserId not implemented")
}
func (UnimplementedExpressionServiceServer) Update(context.Context, *ExpressionRequest) (*ExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedExpressionServiceServer) Delete(context.Context, *ExpressionRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedExpressionServiceServer) mustEmbedUnimplementedExpressionServiceServer() {}
func (UnimplementedExpressionServiceServer) testEmbeddedByValue()                           {}

// UnsafeExpressionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpressionServiceServer will
// result in compilation errors.
type UnsafeExpressionServiceServer interface {
	mustEmbedUnimplementedExpressionServiceServer()
}

func RegisterExpressionServiceServer(s grpc.ServiceRegistrar, srv ExpressionServiceServer) {
	// If the following call pancis, it indicates UnimplementedExpressionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExpressionService_ServiceDesc, srv)
}

func _ExpressionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).Get(ctx, req.(*ExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressionService_GetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).GetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_GetByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).GetByUserId(ctx, req.(*GetByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).Update(ctx, req.(*ExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpressionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionServiceServer).Delete(ctx, req.(*ExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpressionService_ServiceDesc is the grpc.ServiceDesc for ExpressionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpressionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orchestrator.ExpressionService",
	HandlerType: (*ExpressionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ExpressionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ExpressionService_Get_Handler,
		},
		{
			MethodName: "GetByUserId",
			Handler:    _ExpressionService_GetByUserId_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ExpressionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ExpressionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: idl/pbecs/ecs.proto

package pbecs

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

// EcsServiceClient is the client API for EcsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcsServiceClient interface {
	// 创建多家云ECS
	CreateMultipleEcs(ctx context.Context, in *CreateEcsMultipleReq, opts ...grpc.CallOption) (*CreateEcsMultipleResp, error)
	// 创建ECS
	CreateEcs(ctx context.Context, in *CreateEcsReq, opts ...grpc.CallOption) (*CreateEcsResp, error)
	// 删除ECS
	DeleteEcs(ctx context.Context, in *DeleteEcsReq, opts ...grpc.CallOption) (*DeleteEcsResp, error)
	// 修改ECS
	UpdateEcs(ctx context.Context, in *UpdateEcsReq, opts ...grpc.CallOption) (*UpdateEcsResp, error)
	// 查询ECS明细 - 支持云类型、区域、账户、分页等过滤条件
	ListEcsDetail(ctx context.Context, in *ListDetailReq, opts ...grpc.CallOption) (*ListDetailResp, error)
	// 查询ECS全量 - 根据云类型
	ListEcs(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
	// 查询所有云的ECS
	ListEcsAll(ctx context.Context, in *ListAllReq, opts ...grpc.CallOption) (*ListResp, error)
	//操作ecs(start-stop-restart)
	ActionEcs(ctx context.Context, in *ActionReq, opts ...grpc.CallOption) (*ActionResp, error)
}

type ecsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEcsServiceClient(cc grpc.ClientConnInterface) EcsServiceClient {
	return &ecsServiceClient{cc}
}

func (c *ecsServiceClient) CreateMultipleEcs(ctx context.Context, in *CreateEcsMultipleReq, opts ...grpc.CallOption) (*CreateEcsMultipleResp, error) {
	out := new(CreateEcsMultipleResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/CreateMultipleEcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) CreateEcs(ctx context.Context, in *CreateEcsReq, opts ...grpc.CallOption) (*CreateEcsResp, error) {
	out := new(CreateEcsResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/CreateEcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) DeleteEcs(ctx context.Context, in *DeleteEcsReq, opts ...grpc.CallOption) (*DeleteEcsResp, error) {
	out := new(DeleteEcsResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/DeleteEcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) UpdateEcs(ctx context.Context, in *UpdateEcsReq, opts ...grpc.CallOption) (*UpdateEcsResp, error) {
	out := new(UpdateEcsResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/UpdateEcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) ListEcsDetail(ctx context.Context, in *ListDetailReq, opts ...grpc.CallOption) (*ListDetailResp, error) {
	out := new(ListDetailResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/ListEcsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) ListEcs(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/ListEcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) ListEcsAll(ctx context.Context, in *ListAllReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/ListEcsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecsServiceClient) ActionEcs(ctx context.Context, in *ActionReq, opts ...grpc.CallOption) (*ActionResp, error) {
	out := new(ActionResp)
	err := c.cc.Invoke(ctx, "/pbecs.EcsService/ActionEcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcsServiceServer is the server API for EcsService service.
// All implementations must embed UnimplementedEcsServiceServer
// for forward compatibility
type EcsServiceServer interface {
	// 创建多家云ECS
	CreateMultipleEcs(context.Context, *CreateEcsMultipleReq) (*CreateEcsMultipleResp, error)
	// 创建ECS
	CreateEcs(context.Context, *CreateEcsReq) (*CreateEcsResp, error)
	// 删除ECS
	DeleteEcs(context.Context, *DeleteEcsReq) (*DeleteEcsResp, error)
	// 修改ECS
	UpdateEcs(context.Context, *UpdateEcsReq) (*UpdateEcsResp, error)
	// 查询ECS明细 - 支持云类型、区域、账户、分页等过滤条件
	ListEcsDetail(context.Context, *ListDetailReq) (*ListDetailResp, error)
	// 查询ECS全量 - 根据云类型
	ListEcs(context.Context, *ListReq) (*ListResp, error)
	// 查询所有云的ECS
	ListEcsAll(context.Context, *ListAllReq) (*ListResp, error)
	//操作ecs(start-stop-restart)
	ActionEcs(context.Context, *ActionReq) (*ActionResp, error)
	mustEmbedUnimplementedEcsServiceServer()
}

// UnimplementedEcsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEcsServiceServer struct {
}

func (UnimplementedEcsServiceServer) CreateMultipleEcs(context.Context, *CreateEcsMultipleReq) (*CreateEcsMultipleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMultipleEcs not implemented")
}
func (UnimplementedEcsServiceServer) CreateEcs(context.Context, *CreateEcsReq) (*CreateEcsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEcs not implemented")
}
func (UnimplementedEcsServiceServer) DeleteEcs(context.Context, *DeleteEcsReq) (*DeleteEcsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEcs not implemented")
}
func (UnimplementedEcsServiceServer) UpdateEcs(context.Context, *UpdateEcsReq) (*UpdateEcsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEcs not implemented")
}
func (UnimplementedEcsServiceServer) ListEcsDetail(context.Context, *ListDetailReq) (*ListDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEcsDetail not implemented")
}
func (UnimplementedEcsServiceServer) ListEcs(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEcs not implemented")
}
func (UnimplementedEcsServiceServer) ListEcsAll(context.Context, *ListAllReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEcsAll not implemented")
}
func (UnimplementedEcsServiceServer) ActionEcs(context.Context, *ActionReq) (*ActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionEcs not implemented")
}
func (UnimplementedEcsServiceServer) mustEmbedUnimplementedEcsServiceServer() {}

// UnsafeEcsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcsServiceServer will
// result in compilation errors.
type UnsafeEcsServiceServer interface {
	mustEmbedUnimplementedEcsServiceServer()
}

func RegisterEcsServiceServer(s grpc.ServiceRegistrar, srv EcsServiceServer) {
	s.RegisterService(&EcsService_ServiceDesc, srv)
}

func _EcsService_CreateMultipleEcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEcsMultipleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).CreateMultipleEcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/CreateMultipleEcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).CreateMultipleEcs(ctx, req.(*CreateEcsMultipleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_CreateEcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEcsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).CreateEcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/CreateEcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).CreateEcs(ctx, req.(*CreateEcsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_DeleteEcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEcsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).DeleteEcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/DeleteEcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).DeleteEcs(ctx, req.(*DeleteEcsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_UpdateEcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEcsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).UpdateEcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/UpdateEcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).UpdateEcs(ctx, req.(*UpdateEcsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_ListEcsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).ListEcsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/ListEcsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).ListEcsDetail(ctx, req.(*ListDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_ListEcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).ListEcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/ListEcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).ListEcs(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_ListEcsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).ListEcsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/ListEcsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).ListEcsAll(ctx, req.(*ListAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcsService_ActionEcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcsServiceServer).ActionEcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.EcsService/ActionEcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcsServiceServer).ActionEcs(ctx, req.(*ActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EcsService_ServiceDesc is the grpc.ServiceDesc for EcsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EcsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbecs.EcsService",
	HandlerType: (*EcsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMultipleEcs",
			Handler:    _EcsService_CreateMultipleEcs_Handler,
		},
		{
			MethodName: "CreateEcs",
			Handler:    _EcsService_CreateEcs_Handler,
		},
		{
			MethodName: "DeleteEcs",
			Handler:    _EcsService_DeleteEcs_Handler,
		},
		{
			MethodName: "UpdateEcs",
			Handler:    _EcsService_UpdateEcs_Handler,
		},
		{
			MethodName: "ListEcsDetail",
			Handler:    _EcsService_ListEcsDetail_Handler,
		},
		{
			MethodName: "ListEcs",
			Handler:    _EcsService_ListEcs_Handler,
		},
		{
			MethodName: "ListEcsAll",
			Handler:    _EcsService_ListEcsAll_Handler,
		},
		{
			MethodName: "ActionEcs",
			Handler:    _EcsService_ActionEcs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/pbecs/ecs.proto",
}

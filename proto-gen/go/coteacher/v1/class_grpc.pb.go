// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: coteacher/v1/class.proto

package coteacherv1

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

const (
	ClassService_CreateClass_FullMethodName              = "/coteacher.v1.ClassService/CreateClass"
	ClassService_GetClassByID_FullMethodName             = "/coteacher.v1.ClassService/GetClassByID"
	ClassService_GetClassListByTeacherID_FullMethodName  = "/coteacher.v1.ClassService/GetClassListByTeacherID"
	ClassService_UpdateClass_FullMethodName              = "/coteacher.v1.ClassService/UpdateClass"
	ClassService_DeleteClass_FullMethodName              = "/coteacher.v1.ClassService/DeleteClass"
	ClassService_CheckClassEditPermission_FullMethodName = "/coteacher.v1.ClassService/CheckClassEditPermission"
	ClassService_CheckClassViewPermission_FullMethodName = "/coteacher.v1.ClassService/CheckClassViewPermission"
)

// ClassServiceClient is the client API for ClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassServiceClient interface {
	CreateClass(ctx context.Context, in *CreateClassRequest, opts ...grpc.CallOption) (*CreateClassResponse, error)
	GetClassByID(ctx context.Context, in *GetClassByIDRequest, opts ...grpc.CallOption) (*GetClassByIDResponse, error)
	GetClassListByTeacherID(ctx context.Context, in *GetClassListByTeacherIDRequest, opts ...grpc.CallOption) (*GetClassListByTeacherIDResponse, error)
	UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...grpc.CallOption) (*UpdateClassResponse, error)
	DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...grpc.CallOption) (*DeleteClassResponse, error)
	CheckClassEditPermission(ctx context.Context, in *CheckClassEditPermissionRequest, opts ...grpc.CallOption) (*CheckClassEditPermissionResponse, error)
	CheckClassViewPermission(ctx context.Context, in *CheckClassViewPermissionRequest, opts ...grpc.CallOption) (*CheckClassViewPermissionResponse, error)
}

type classServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassServiceClient(cc grpc.ClientConnInterface) ClassServiceClient {
	return &classServiceClient{cc}
}

func (c *classServiceClient) CreateClass(ctx context.Context, in *CreateClassRequest, opts ...grpc.CallOption) (*CreateClassResponse, error) {
	out := new(CreateClassResponse)
	err := c.cc.Invoke(ctx, ClassService_CreateClass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) GetClassByID(ctx context.Context, in *GetClassByIDRequest, opts ...grpc.CallOption) (*GetClassByIDResponse, error) {
	out := new(GetClassByIDResponse)
	err := c.cc.Invoke(ctx, ClassService_GetClassByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) GetClassListByTeacherID(ctx context.Context, in *GetClassListByTeacherIDRequest, opts ...grpc.CallOption) (*GetClassListByTeacherIDResponse, error) {
	out := new(GetClassListByTeacherIDResponse)
	err := c.cc.Invoke(ctx, ClassService_GetClassListByTeacherID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...grpc.CallOption) (*UpdateClassResponse, error) {
	out := new(UpdateClassResponse)
	err := c.cc.Invoke(ctx, ClassService_UpdateClass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...grpc.CallOption) (*DeleteClassResponse, error) {
	out := new(DeleteClassResponse)
	err := c.cc.Invoke(ctx, ClassService_DeleteClass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) CheckClassEditPermission(ctx context.Context, in *CheckClassEditPermissionRequest, opts ...grpc.CallOption) (*CheckClassEditPermissionResponse, error) {
	out := new(CheckClassEditPermissionResponse)
	err := c.cc.Invoke(ctx, ClassService_CheckClassEditPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) CheckClassViewPermission(ctx context.Context, in *CheckClassViewPermissionRequest, opts ...grpc.CallOption) (*CheckClassViewPermissionResponse, error) {
	out := new(CheckClassViewPermissionResponse)
	err := c.cc.Invoke(ctx, ClassService_CheckClassViewPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServiceServer is the server API for ClassService service.
// All implementations should embed UnimplementedClassServiceServer
// for forward compatibility
type ClassServiceServer interface {
	CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error)
	GetClassByID(context.Context, *GetClassByIDRequest) (*GetClassByIDResponse, error)
	GetClassListByTeacherID(context.Context, *GetClassListByTeacherIDRequest) (*GetClassListByTeacherIDResponse, error)
	UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error)
	DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error)
	CheckClassEditPermission(context.Context, *CheckClassEditPermissionRequest) (*CheckClassEditPermissionResponse, error)
	CheckClassViewPermission(context.Context, *CheckClassViewPermissionRequest) (*CheckClassViewPermissionResponse, error)
}

// UnimplementedClassServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClassServiceServer struct {
}

func (UnimplementedClassServiceServer) CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClass not implemented")
}
func (UnimplementedClassServiceServer) GetClassByID(context.Context, *GetClassByIDRequest) (*GetClassByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassByID not implemented")
}
func (UnimplementedClassServiceServer) GetClassListByTeacherID(context.Context, *GetClassListByTeacherIDRequest) (*GetClassListByTeacherIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassListByTeacherID not implemented")
}
func (UnimplementedClassServiceServer) UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClass not implemented")
}
func (UnimplementedClassServiceServer) DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClass not implemented")
}
func (UnimplementedClassServiceServer) CheckClassEditPermission(context.Context, *CheckClassEditPermissionRequest) (*CheckClassEditPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckClassEditPermission not implemented")
}
func (UnimplementedClassServiceServer) CheckClassViewPermission(context.Context, *CheckClassViewPermissionRequest) (*CheckClassViewPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckClassViewPermission not implemented")
}

// UnsafeClassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassServiceServer will
// result in compilation errors.
type UnsafeClassServiceServer interface {
	mustEmbedUnimplementedClassServiceServer()
}

func RegisterClassServiceServer(s grpc.ServiceRegistrar, srv ClassServiceServer) {
	s.RegisterService(&ClassService_ServiceDesc, srv)
}

func _ClassService_CreateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).CreateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_CreateClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).CreateClass(ctx, req.(*CreateClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_GetClassByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).GetClassByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_GetClassByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).GetClassByID(ctx, req.(*GetClassByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_GetClassListByTeacherID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassListByTeacherIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).GetClassListByTeacherID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_GetClassListByTeacherID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).GetClassListByTeacherID(ctx, req.(*GetClassListByTeacherIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_UpdateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).UpdateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_UpdateClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).UpdateClass(ctx, req.(*UpdateClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_DeleteClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).DeleteClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_DeleteClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).DeleteClass(ctx, req.(*DeleteClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_CheckClassEditPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckClassEditPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).CheckClassEditPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_CheckClassEditPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).CheckClassEditPermission(ctx, req.(*CheckClassEditPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_CheckClassViewPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckClassViewPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).CheckClassViewPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_CheckClassViewPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).CheckClassViewPermission(ctx, req.(*CheckClassViewPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassService_ServiceDesc is the grpc.ServiceDesc for ClassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coteacher.v1.ClassService",
	HandlerType: (*ClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClass",
			Handler:    _ClassService_CreateClass_Handler,
		},
		{
			MethodName: "GetClassByID",
			Handler:    _ClassService_GetClassByID_Handler,
		},
		{
			MethodName: "GetClassListByTeacherID",
			Handler:    _ClassService_GetClassListByTeacherID_Handler,
		},
		{
			MethodName: "UpdateClass",
			Handler:    _ClassService_UpdateClass_Handler,
		},
		{
			MethodName: "DeleteClass",
			Handler:    _ClassService_DeleteClass_Handler,
		},
		{
			MethodName: "CheckClassEditPermission",
			Handler:    _ClassService_CheckClassEditPermission_Handler,
		},
		{
			MethodName: "CheckClassViewPermission",
			Handler:    _ClassService_CheckClassViewPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coteacher/v1/class.proto",
}

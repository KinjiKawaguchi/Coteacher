// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: coteacher/v1/student_class.proto

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
	StudentClassService_CreateStudentClass_FullMethodName      = "/coteacher.v1.StudentClassService/CreateStudentClass"
	StudentClassService_GetStudentListByClassID_FullMethodName = "/coteacher.v1.StudentClassService/GetStudentListByClassID"
	StudentClassService_GetClassListByStudentID_FullMethodName = "/coteacher.v1.StudentClassService/GetClassListByStudentID"
	StudentClassService_DeleteStudentClass_FullMethodName      = "/coteacher.v1.StudentClassService/DeleteStudentClass"
)

// StudentClassServiceClient is the client API for StudentClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClassServiceClient interface {
	CreateStudentClass(ctx context.Context, in *CreateStudentClassRequest, opts ...grpc.CallOption) (*CreateStudentClassResponse, error)
	GetStudentListByClassID(ctx context.Context, in *GetStudentListByClassIDRequest, opts ...grpc.CallOption) (*GetStudentListByClassIDResponse, error)
	GetClassListByStudentID(ctx context.Context, in *GetClassListByStudentIDRequest, opts ...grpc.CallOption) (*GetClassListByStudentIDResponse, error)
	DeleteStudentClass(ctx context.Context, in *DeleteStudentClassRequest, opts ...grpc.CallOption) (*DeleteStudentClassResponse, error)
}

type studentClassServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClassServiceClient(cc grpc.ClientConnInterface) StudentClassServiceClient {
	return &studentClassServiceClient{cc}
}

func (c *studentClassServiceClient) CreateStudentClass(ctx context.Context, in *CreateStudentClassRequest, opts ...grpc.CallOption) (*CreateStudentClassResponse, error) {
	out := new(CreateStudentClassResponse)
	err := c.cc.Invoke(ctx, StudentClassService_CreateStudentClass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassServiceClient) GetStudentListByClassID(ctx context.Context, in *GetStudentListByClassIDRequest, opts ...grpc.CallOption) (*GetStudentListByClassIDResponse, error) {
	out := new(GetStudentListByClassIDResponse)
	err := c.cc.Invoke(ctx, StudentClassService_GetStudentListByClassID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassServiceClient) GetClassListByStudentID(ctx context.Context, in *GetClassListByStudentIDRequest, opts ...grpc.CallOption) (*GetClassListByStudentIDResponse, error) {
	out := new(GetClassListByStudentIDResponse)
	err := c.cc.Invoke(ctx, StudentClassService_GetClassListByStudentID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassServiceClient) DeleteStudentClass(ctx context.Context, in *DeleteStudentClassRequest, opts ...grpc.CallOption) (*DeleteStudentClassResponse, error) {
	out := new(DeleteStudentClassResponse)
	err := c.cc.Invoke(ctx, StudentClassService_DeleteStudentClass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentClassServiceServer is the server API for StudentClassService service.
// All implementations should embed UnimplementedStudentClassServiceServer
// for forward compatibility
type StudentClassServiceServer interface {
	CreateStudentClass(context.Context, *CreateStudentClassRequest) (*CreateStudentClassResponse, error)
	GetStudentListByClassID(context.Context, *GetStudentListByClassIDRequest) (*GetStudentListByClassIDResponse, error)
	GetClassListByStudentID(context.Context, *GetClassListByStudentIDRequest) (*GetClassListByStudentIDResponse, error)
	DeleteStudentClass(context.Context, *DeleteStudentClassRequest) (*DeleteStudentClassResponse, error)
}

// UnimplementedStudentClassServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStudentClassServiceServer struct {
}

func (UnimplementedStudentClassServiceServer) CreateStudentClass(context.Context, *CreateStudentClassRequest) (*CreateStudentClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudentClass not implemented")
}
func (UnimplementedStudentClassServiceServer) GetStudentListByClassID(context.Context, *GetStudentListByClassIDRequest) (*GetStudentListByClassIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentListByClassID not implemented")
}
func (UnimplementedStudentClassServiceServer) GetClassListByStudentID(context.Context, *GetClassListByStudentIDRequest) (*GetClassListByStudentIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassListByStudentID not implemented")
}
func (UnimplementedStudentClassServiceServer) DeleteStudentClass(context.Context, *DeleteStudentClassRequest) (*DeleteStudentClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudentClass not implemented")
}

// UnsafeStudentClassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentClassServiceServer will
// result in compilation errors.
type UnsafeStudentClassServiceServer interface {
	mustEmbedUnimplementedStudentClassServiceServer()
}

func RegisterStudentClassServiceServer(s grpc.ServiceRegistrar, srv StudentClassServiceServer) {
	s.RegisterService(&StudentClassService_ServiceDesc, srv)
}

func _StudentClassService_CreateStudentClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassServiceServer).CreateStudentClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentClassService_CreateStudentClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassServiceServer).CreateStudentClass(ctx, req.(*CreateStudentClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassService_GetStudentListByClassID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentListByClassIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassServiceServer).GetStudentListByClassID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentClassService_GetStudentListByClassID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassServiceServer).GetStudentListByClassID(ctx, req.(*GetStudentListByClassIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassService_GetClassListByStudentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassListByStudentIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassServiceServer).GetClassListByStudentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentClassService_GetClassListByStudentID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassServiceServer).GetClassListByStudentID(ctx, req.(*GetClassListByStudentIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassService_DeleteStudentClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassServiceServer).DeleteStudentClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentClassService_DeleteStudentClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassServiceServer).DeleteStudentClass(ctx, req.(*DeleteStudentClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentClassService_ServiceDesc is the grpc.ServiceDesc for StudentClassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentClassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coteacher.v1.StudentClassService",
	HandlerType: (*StudentClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudentClass",
			Handler:    _StudentClassService_CreateStudentClass_Handler,
		},
		{
			MethodName: "GetStudentListByClassID",
			Handler:    _StudentClassService_GetStudentListByClassID_Handler,
		},
		{
			MethodName: "GetClassListByStudentID",
			Handler:    _StudentClassService_GetClassListByStudentID_Handler,
		},
		{
			MethodName: "DeleteStudentClass",
			Handler:    _StudentClassService_DeleteStudentClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coteacher/v1/student_class.proto",
}

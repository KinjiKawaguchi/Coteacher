// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: coteacher/v1/class_invitation_code.proto

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
	ClassInvitationCodeService_CreateClassInvitationCode_FullMethodName           = "/coteacher.v1.ClassInvitationCodeService/CreateClassInvitationCode"
	ClassInvitationCodeService_GetClassInvitationCodeByID_FullMethodName          = "/coteacher.v1.ClassInvitationCodeService/GetClassInvitationCodeByID"
	ClassInvitationCodeService_GetClassInvitationCodeListByClassID_FullMethodName = "/coteacher.v1.ClassInvitationCodeService/GetClassInvitationCodeListByClassID"
	ClassInvitationCodeService_UpdateClassInvitationCode_FullMethodName           = "/coteacher.v1.ClassInvitationCodeService/UpdateClassInvitationCode"
)

// ClassInvitationCodeServiceClient is the client API for ClassInvitationCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassInvitationCodeServiceClient interface {
	CreateClassInvitationCode(ctx context.Context, in *CreateClassInvitationCodeRequest, opts ...grpc.CallOption) (*CreateClassInvitationCodeResponse, error)
	GetClassInvitationCodeByID(ctx context.Context, in *GetClassInvitationCodeByIDRequest, opts ...grpc.CallOption) (*GetClassInvitationCodeByIDResponse, error)
	GetClassInvitationCodeListByClassID(ctx context.Context, in *GetClassInvitationCodeListByClassIDRequest, opts ...grpc.CallOption) (*GetClassInvitationCodeListByClassIDResponse, error)
	UpdateClassInvitationCode(ctx context.Context, in *UpdateClassInvitationCodeRequest, opts ...grpc.CallOption) (*UpdateClassInvitationCodeResponse, error)
}

type classInvitationCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassInvitationCodeServiceClient(cc grpc.ClientConnInterface) ClassInvitationCodeServiceClient {
	return &classInvitationCodeServiceClient{cc}
}

func (c *classInvitationCodeServiceClient) CreateClassInvitationCode(ctx context.Context, in *CreateClassInvitationCodeRequest, opts ...grpc.CallOption) (*CreateClassInvitationCodeResponse, error) {
	out := new(CreateClassInvitationCodeResponse)
	err := c.cc.Invoke(ctx, ClassInvitationCodeService_CreateClassInvitationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classInvitationCodeServiceClient) GetClassInvitationCodeByID(ctx context.Context, in *GetClassInvitationCodeByIDRequest, opts ...grpc.CallOption) (*GetClassInvitationCodeByIDResponse, error) {
	out := new(GetClassInvitationCodeByIDResponse)
	err := c.cc.Invoke(ctx, ClassInvitationCodeService_GetClassInvitationCodeByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classInvitationCodeServiceClient) GetClassInvitationCodeListByClassID(ctx context.Context, in *GetClassInvitationCodeListByClassIDRequest, opts ...grpc.CallOption) (*GetClassInvitationCodeListByClassIDResponse, error) {
	out := new(GetClassInvitationCodeListByClassIDResponse)
	err := c.cc.Invoke(ctx, ClassInvitationCodeService_GetClassInvitationCodeListByClassID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classInvitationCodeServiceClient) UpdateClassInvitationCode(ctx context.Context, in *UpdateClassInvitationCodeRequest, opts ...grpc.CallOption) (*UpdateClassInvitationCodeResponse, error) {
	out := new(UpdateClassInvitationCodeResponse)
	err := c.cc.Invoke(ctx, ClassInvitationCodeService_UpdateClassInvitationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassInvitationCodeServiceServer is the server API for ClassInvitationCodeService service.
// All implementations should embed UnimplementedClassInvitationCodeServiceServer
// for forward compatibility
type ClassInvitationCodeServiceServer interface {
	CreateClassInvitationCode(context.Context, *CreateClassInvitationCodeRequest) (*CreateClassInvitationCodeResponse, error)
	GetClassInvitationCodeByID(context.Context, *GetClassInvitationCodeByIDRequest) (*GetClassInvitationCodeByIDResponse, error)
	GetClassInvitationCodeListByClassID(context.Context, *GetClassInvitationCodeListByClassIDRequest) (*GetClassInvitationCodeListByClassIDResponse, error)
	UpdateClassInvitationCode(context.Context, *UpdateClassInvitationCodeRequest) (*UpdateClassInvitationCodeResponse, error)
}

// UnimplementedClassInvitationCodeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClassInvitationCodeServiceServer struct {
}

func (UnimplementedClassInvitationCodeServiceServer) CreateClassInvitationCode(context.Context, *CreateClassInvitationCodeRequest) (*CreateClassInvitationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassInvitationCode not implemented")
}
func (UnimplementedClassInvitationCodeServiceServer) GetClassInvitationCodeByID(context.Context, *GetClassInvitationCodeByIDRequest) (*GetClassInvitationCodeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassInvitationCodeByID not implemented")
}
func (UnimplementedClassInvitationCodeServiceServer) GetClassInvitationCodeListByClassID(context.Context, *GetClassInvitationCodeListByClassIDRequest) (*GetClassInvitationCodeListByClassIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassInvitationCodeListByClassID not implemented")
}
func (UnimplementedClassInvitationCodeServiceServer) UpdateClassInvitationCode(context.Context, *UpdateClassInvitationCodeRequest) (*UpdateClassInvitationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClassInvitationCode not implemented")
}

// UnsafeClassInvitationCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassInvitationCodeServiceServer will
// result in compilation errors.
type UnsafeClassInvitationCodeServiceServer interface {
	mustEmbedUnimplementedClassInvitationCodeServiceServer()
}

func RegisterClassInvitationCodeServiceServer(s grpc.ServiceRegistrar, srv ClassInvitationCodeServiceServer) {
	s.RegisterService(&ClassInvitationCodeService_ServiceDesc, srv)
}

func _ClassInvitationCodeService_CreateClassInvitationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassInvitationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassInvitationCodeServiceServer).CreateClassInvitationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassInvitationCodeService_CreateClassInvitationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassInvitationCodeServiceServer).CreateClassInvitationCode(ctx, req.(*CreateClassInvitationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassInvitationCodeService_GetClassInvitationCodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassInvitationCodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassInvitationCodeServiceServer).GetClassInvitationCodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassInvitationCodeService_GetClassInvitationCodeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassInvitationCodeServiceServer).GetClassInvitationCodeByID(ctx, req.(*GetClassInvitationCodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassInvitationCodeService_GetClassInvitationCodeListByClassID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassInvitationCodeListByClassIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassInvitationCodeServiceServer).GetClassInvitationCodeListByClassID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassInvitationCodeService_GetClassInvitationCodeListByClassID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassInvitationCodeServiceServer).GetClassInvitationCodeListByClassID(ctx, req.(*GetClassInvitationCodeListByClassIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassInvitationCodeService_UpdateClassInvitationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassInvitationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassInvitationCodeServiceServer).UpdateClassInvitationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassInvitationCodeService_UpdateClassInvitationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassInvitationCodeServiceServer).UpdateClassInvitationCode(ctx, req.(*UpdateClassInvitationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassInvitationCodeService_ServiceDesc is the grpc.ServiceDesc for ClassInvitationCodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassInvitationCodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coteacher.v1.ClassInvitationCodeService",
	HandlerType: (*ClassInvitationCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClassInvitationCode",
			Handler:    _ClassInvitationCodeService_CreateClassInvitationCode_Handler,
		},
		{
			MethodName: "GetClassInvitationCodeByID",
			Handler:    _ClassInvitationCodeService_GetClassInvitationCodeByID_Handler,
		},
		{
			MethodName: "GetClassInvitationCodeListByClassID",
			Handler:    _ClassInvitationCodeService_GetClassInvitationCodeListByClassID_Handler,
		},
		{
			MethodName: "UpdateClassInvitationCode",
			Handler:    _ClassInvitationCodeService_UpdateClassInvitationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coteacher/v1/class_invitation_code.proto",
}

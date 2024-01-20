// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: coteacher/v1/question.proto

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
	QuestionService_GetQuestionListByFormId_FullMethodName = "/coteacher.v1.QuestionService/GetQuestionListByFormId"
	QuestionService_SaveQuestionList_FullMethodName        = "/coteacher.v1.QuestionService/SaveQuestionList"
)

// QuestionServiceClient is the client API for QuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionServiceClient interface {
	GetQuestionListByFormId(ctx context.Context, in *GetQuestionListByFormIdRequest, opts ...grpc.CallOption) (*GetQuestionListByFormIdResponse, error)
	SaveQuestionList(ctx context.Context, in *SaveQuestionListRequest, opts ...grpc.CallOption) (*SaveQuestionListResponse, error)
}

type questionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionServiceClient(cc grpc.ClientConnInterface) QuestionServiceClient {
	return &questionServiceClient{cc}
}

func (c *questionServiceClient) GetQuestionListByFormId(ctx context.Context, in *GetQuestionListByFormIdRequest, opts ...grpc.CallOption) (*GetQuestionListByFormIdResponse, error) {
	out := new(GetQuestionListByFormIdResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionListByFormId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) SaveQuestionList(ctx context.Context, in *SaveQuestionListRequest, opts ...grpc.CallOption) (*SaveQuestionListResponse, error) {
	out := new(SaveQuestionListResponse)
	err := c.cc.Invoke(ctx, QuestionService_SaveQuestionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiceServer is the server API for QuestionService service.
// All implementations should embed UnimplementedQuestionServiceServer
// for forward compatibility
type QuestionServiceServer interface {
	GetQuestionListByFormId(context.Context, *GetQuestionListByFormIdRequest) (*GetQuestionListByFormIdResponse, error)
	SaveQuestionList(context.Context, *SaveQuestionListRequest) (*SaveQuestionListResponse, error)
}

// UnimplementedQuestionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQuestionServiceServer struct {
}

func (UnimplementedQuestionServiceServer) GetQuestionListByFormId(context.Context, *GetQuestionListByFormIdRequest) (*GetQuestionListByFormIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionListByFormId not implemented")
}
func (UnimplementedQuestionServiceServer) SaveQuestionList(context.Context, *SaveQuestionListRequest) (*SaveQuestionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQuestionList not implemented")
}

// UnsafeQuestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServiceServer will
// result in compilation errors.
type UnsafeQuestionServiceServer interface {
	mustEmbedUnimplementedQuestionServiceServer()
}

func RegisterQuestionServiceServer(s grpc.ServiceRegistrar, srv QuestionServiceServer) {
	s.RegisterService(&QuestionService_ServiceDesc, srv)
}

func _QuestionService_GetQuestionListByFormId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionListByFormIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionListByFormId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionListByFormId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionListByFormId(ctx, req.(*GetQuestionListByFormIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_SaveQuestionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQuestionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).SaveQuestionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_SaveQuestionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).SaveQuestionList(ctx, req.(*SaveQuestionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionService_ServiceDesc is the grpc.ServiceDesc for QuestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coteacher.v1.QuestionService",
	HandlerType: (*QuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestionListByFormId",
			Handler:    _QuestionService_GetQuestionListByFormId_Handler,
		},
		{
			MethodName: "SaveQuestionList",
			Handler:    _QuestionService_SaveQuestionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coteacher/v1/question.proto",
}
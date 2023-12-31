// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: coteacher/v1/class_invitation_code.proto

package coteacherv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ClassInvitationCodeServiceName is the fully-qualified name of the ClassInvitationCodeService
	// service.
	ClassInvitationCodeServiceName = "coteacher.v1.ClassInvitationCodeService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClassInvitationCodeServiceCreateClassInvitationCodeProcedure is the fully-qualified name of the
	// ClassInvitationCodeService's CreateClassInvitationCode RPC.
	ClassInvitationCodeServiceCreateClassInvitationCodeProcedure = "/coteacher.v1.ClassInvitationCodeService/CreateClassInvitationCode"
	// ClassInvitationCodeServiceGetClassInvitationCodeByIDProcedure is the fully-qualified name of the
	// ClassInvitationCodeService's GetClassInvitationCodeByID RPC.
	ClassInvitationCodeServiceGetClassInvitationCodeByIDProcedure = "/coteacher.v1.ClassInvitationCodeService/GetClassInvitationCodeByID"
	// ClassInvitationCodeServiceGetClassInvitationCodeListByClassIDProcedure is the fully-qualified
	// name of the ClassInvitationCodeService's GetClassInvitationCodeListByClassID RPC.
	ClassInvitationCodeServiceGetClassInvitationCodeListByClassIDProcedure = "/coteacher.v1.ClassInvitationCodeService/GetClassInvitationCodeListByClassID"
	// ClassInvitationCodeServiceUpdateClassInvitationCodeProcedure is the fully-qualified name of the
	// ClassInvitationCodeService's UpdateClassInvitationCode RPC.
	ClassInvitationCodeServiceUpdateClassInvitationCodeProcedure = "/coteacher.v1.ClassInvitationCodeService/UpdateClassInvitationCode"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	classInvitationCodeServiceServiceDescriptor                                   = v1.File_coteacher_v1_class_invitation_code_proto.Services().ByName("ClassInvitationCodeService")
	classInvitationCodeServiceCreateClassInvitationCodeMethodDescriptor           = classInvitationCodeServiceServiceDescriptor.Methods().ByName("CreateClassInvitationCode")
	classInvitationCodeServiceGetClassInvitationCodeByIDMethodDescriptor          = classInvitationCodeServiceServiceDescriptor.Methods().ByName("GetClassInvitationCodeByID")
	classInvitationCodeServiceGetClassInvitationCodeListByClassIDMethodDescriptor = classInvitationCodeServiceServiceDescriptor.Methods().ByName("GetClassInvitationCodeListByClassID")
	classInvitationCodeServiceUpdateClassInvitationCodeMethodDescriptor           = classInvitationCodeServiceServiceDescriptor.Methods().ByName("UpdateClassInvitationCode")
)

// ClassInvitationCodeServiceClient is a client for the coteacher.v1.ClassInvitationCodeService
// service.
type ClassInvitationCodeServiceClient interface {
	CreateClassInvitationCode(context.Context, *connect.Request[v1.CreateClassInvitationCodeRequest]) (*connect.Response[v1.CreateClassInvitationCodeResponse], error)
	GetClassInvitationCodeByID(context.Context, *connect.Request[v1.GetClassInvitationCodeByIDRequest]) (*connect.Response[v1.GetClassInvitationCodeByIDResponse], error)
	GetClassInvitationCodeListByClassID(context.Context, *connect.Request[v1.GetClassInvitationCodeListByClassIDRequest]) (*connect.Response[v1.GetClassInvitationCodeListByClassIDResponse], error)
	UpdateClassInvitationCode(context.Context, *connect.Request[v1.UpdateClassInvitationCodeRequest]) (*connect.Response[v1.UpdateClassInvitationCodeResponse], error)
}

// NewClassInvitationCodeServiceClient constructs a client for the
// coteacher.v1.ClassInvitationCodeService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClassInvitationCodeServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ClassInvitationCodeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &classInvitationCodeServiceClient{
		createClassInvitationCode: connect.NewClient[v1.CreateClassInvitationCodeRequest, v1.CreateClassInvitationCodeResponse](
			httpClient,
			baseURL+ClassInvitationCodeServiceCreateClassInvitationCodeProcedure,
			connect.WithSchema(classInvitationCodeServiceCreateClassInvitationCodeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getClassInvitationCodeByID: connect.NewClient[v1.GetClassInvitationCodeByIDRequest, v1.GetClassInvitationCodeByIDResponse](
			httpClient,
			baseURL+ClassInvitationCodeServiceGetClassInvitationCodeByIDProcedure,
			connect.WithSchema(classInvitationCodeServiceGetClassInvitationCodeByIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getClassInvitationCodeListByClassID: connect.NewClient[v1.GetClassInvitationCodeListByClassIDRequest, v1.GetClassInvitationCodeListByClassIDResponse](
			httpClient,
			baseURL+ClassInvitationCodeServiceGetClassInvitationCodeListByClassIDProcedure,
			connect.WithSchema(classInvitationCodeServiceGetClassInvitationCodeListByClassIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateClassInvitationCode: connect.NewClient[v1.UpdateClassInvitationCodeRequest, v1.UpdateClassInvitationCodeResponse](
			httpClient,
			baseURL+ClassInvitationCodeServiceUpdateClassInvitationCodeProcedure,
			connect.WithSchema(classInvitationCodeServiceUpdateClassInvitationCodeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// classInvitationCodeServiceClient implements ClassInvitationCodeServiceClient.
type classInvitationCodeServiceClient struct {
	createClassInvitationCode           *connect.Client[v1.CreateClassInvitationCodeRequest, v1.CreateClassInvitationCodeResponse]
	getClassInvitationCodeByID          *connect.Client[v1.GetClassInvitationCodeByIDRequest, v1.GetClassInvitationCodeByIDResponse]
	getClassInvitationCodeListByClassID *connect.Client[v1.GetClassInvitationCodeListByClassIDRequest, v1.GetClassInvitationCodeListByClassIDResponse]
	updateClassInvitationCode           *connect.Client[v1.UpdateClassInvitationCodeRequest, v1.UpdateClassInvitationCodeResponse]
}

// CreateClassInvitationCode calls
// coteacher.v1.ClassInvitationCodeService.CreateClassInvitationCode.
func (c *classInvitationCodeServiceClient) CreateClassInvitationCode(ctx context.Context, req *connect.Request[v1.CreateClassInvitationCodeRequest]) (*connect.Response[v1.CreateClassInvitationCodeResponse], error) {
	return c.createClassInvitationCode.CallUnary(ctx, req)
}

// GetClassInvitationCodeByID calls
// coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeByID.
func (c *classInvitationCodeServiceClient) GetClassInvitationCodeByID(ctx context.Context, req *connect.Request[v1.GetClassInvitationCodeByIDRequest]) (*connect.Response[v1.GetClassInvitationCodeByIDResponse], error) {
	return c.getClassInvitationCodeByID.CallUnary(ctx, req)
}

// GetClassInvitationCodeListByClassID calls
// coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeListByClassID.
func (c *classInvitationCodeServiceClient) GetClassInvitationCodeListByClassID(ctx context.Context, req *connect.Request[v1.GetClassInvitationCodeListByClassIDRequest]) (*connect.Response[v1.GetClassInvitationCodeListByClassIDResponse], error) {
	return c.getClassInvitationCodeListByClassID.CallUnary(ctx, req)
}

// UpdateClassInvitationCode calls
// coteacher.v1.ClassInvitationCodeService.UpdateClassInvitationCode.
func (c *classInvitationCodeServiceClient) UpdateClassInvitationCode(ctx context.Context, req *connect.Request[v1.UpdateClassInvitationCodeRequest]) (*connect.Response[v1.UpdateClassInvitationCodeResponse], error) {
	return c.updateClassInvitationCode.CallUnary(ctx, req)
}

// ClassInvitationCodeServiceHandler is an implementation of the
// coteacher.v1.ClassInvitationCodeService service.
type ClassInvitationCodeServiceHandler interface {
	CreateClassInvitationCode(context.Context, *connect.Request[v1.CreateClassInvitationCodeRequest]) (*connect.Response[v1.CreateClassInvitationCodeResponse], error)
	GetClassInvitationCodeByID(context.Context, *connect.Request[v1.GetClassInvitationCodeByIDRequest]) (*connect.Response[v1.GetClassInvitationCodeByIDResponse], error)
	GetClassInvitationCodeListByClassID(context.Context, *connect.Request[v1.GetClassInvitationCodeListByClassIDRequest]) (*connect.Response[v1.GetClassInvitationCodeListByClassIDResponse], error)
	UpdateClassInvitationCode(context.Context, *connect.Request[v1.UpdateClassInvitationCodeRequest]) (*connect.Response[v1.UpdateClassInvitationCodeResponse], error)
}

// NewClassInvitationCodeServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClassInvitationCodeServiceHandler(svc ClassInvitationCodeServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	classInvitationCodeServiceCreateClassInvitationCodeHandler := connect.NewUnaryHandler(
		ClassInvitationCodeServiceCreateClassInvitationCodeProcedure,
		svc.CreateClassInvitationCode,
		connect.WithSchema(classInvitationCodeServiceCreateClassInvitationCodeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	classInvitationCodeServiceGetClassInvitationCodeByIDHandler := connect.NewUnaryHandler(
		ClassInvitationCodeServiceGetClassInvitationCodeByIDProcedure,
		svc.GetClassInvitationCodeByID,
		connect.WithSchema(classInvitationCodeServiceGetClassInvitationCodeByIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	classInvitationCodeServiceGetClassInvitationCodeListByClassIDHandler := connect.NewUnaryHandler(
		ClassInvitationCodeServiceGetClassInvitationCodeListByClassIDProcedure,
		svc.GetClassInvitationCodeListByClassID,
		connect.WithSchema(classInvitationCodeServiceGetClassInvitationCodeListByClassIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	classInvitationCodeServiceUpdateClassInvitationCodeHandler := connect.NewUnaryHandler(
		ClassInvitationCodeServiceUpdateClassInvitationCodeProcedure,
		svc.UpdateClassInvitationCode,
		connect.WithSchema(classInvitationCodeServiceUpdateClassInvitationCodeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/coteacher.v1.ClassInvitationCodeService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ClassInvitationCodeServiceCreateClassInvitationCodeProcedure:
			classInvitationCodeServiceCreateClassInvitationCodeHandler.ServeHTTP(w, r)
		case ClassInvitationCodeServiceGetClassInvitationCodeByIDProcedure:
			classInvitationCodeServiceGetClassInvitationCodeByIDHandler.ServeHTTP(w, r)
		case ClassInvitationCodeServiceGetClassInvitationCodeListByClassIDProcedure:
			classInvitationCodeServiceGetClassInvitationCodeListByClassIDHandler.ServeHTTP(w, r)
		case ClassInvitationCodeServiceUpdateClassInvitationCodeProcedure:
			classInvitationCodeServiceUpdateClassInvitationCodeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedClassInvitationCodeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClassInvitationCodeServiceHandler struct{}

func (UnimplementedClassInvitationCodeServiceHandler) CreateClassInvitationCode(context.Context, *connect.Request[v1.CreateClassInvitationCodeRequest]) (*connect.Response[v1.CreateClassInvitationCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.ClassInvitationCodeService.CreateClassInvitationCode is not implemented"))
}

func (UnimplementedClassInvitationCodeServiceHandler) GetClassInvitationCodeByID(context.Context, *connect.Request[v1.GetClassInvitationCodeByIDRequest]) (*connect.Response[v1.GetClassInvitationCodeByIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeByID is not implemented"))
}

func (UnimplementedClassInvitationCodeServiceHandler) GetClassInvitationCodeListByClassID(context.Context, *connect.Request[v1.GetClassInvitationCodeListByClassIDRequest]) (*connect.Response[v1.GetClassInvitationCodeListByClassIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeListByClassID is not implemented"))
}

func (UnimplementedClassInvitationCodeServiceHandler) UpdateClassInvitationCode(context.Context, *connect.Request[v1.UpdateClassInvitationCodeRequest]) (*connect.Response[v1.UpdateClassInvitationCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.ClassInvitationCodeService.UpdateClassInvitationCode is not implemented"))
}

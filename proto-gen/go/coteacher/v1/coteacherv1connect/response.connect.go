// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: coteacher/v1/response.proto

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
	// ResponseServiceName is the fully-qualified name of the ResponseService service.
	ResponseServiceName = "coteacher.v1.ResponseService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ResponseServiceGetNumberOfResponsesByStudentIDProcedure is the fully-qualified name of the
	// ResponseService's GetNumberOfResponsesByStudentID RPC.
	ResponseServiceGetNumberOfResponsesByStudentIDProcedure = "/coteacher.v1.ResponseService/GetNumberOfResponsesByStudentID"
	// ResponseServiceGetNumberOfResponsesByFormIDProcedure is the fully-qualified name of the
	// ResponseService's GetNumberOfResponsesByFormID RPC.
	ResponseServiceGetNumberOfResponsesByFormIDProcedure = "/coteacher.v1.ResponseService/GetNumberOfResponsesByFormID"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	responseServiceServiceDescriptor                               = v1.File_coteacher_v1_response_proto.Services().ByName("ResponseService")
	responseServiceGetNumberOfResponsesByStudentIDMethodDescriptor = responseServiceServiceDescriptor.Methods().ByName("GetNumberOfResponsesByStudentID")
	responseServiceGetNumberOfResponsesByFormIDMethodDescriptor    = responseServiceServiceDescriptor.Methods().ByName("GetNumberOfResponsesByFormID")
)

// ResponseServiceClient is a client for the coteacher.v1.ResponseService service.
type ResponseServiceClient interface {
	GetNumberOfResponsesByStudentID(context.Context, *connect.Request[v1.GetNumberOfResponsesByStudentIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByStudentIDResponse], error)
	GetNumberOfResponsesByFormID(context.Context, *connect.Request[v1.GetNumberOfResponsesByFormIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByFormIDResponse], error)
}

// NewResponseServiceClient constructs a client for the coteacher.v1.ResponseService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewResponseServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ResponseServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &responseServiceClient{
		getNumberOfResponsesByStudentID: connect.NewClient[v1.GetNumberOfResponsesByStudentIDRequest, v1.GetNumberOfResponsesByStudentIDResponse](
			httpClient,
			baseURL+ResponseServiceGetNumberOfResponsesByStudentIDProcedure,
			connect.WithSchema(responseServiceGetNumberOfResponsesByStudentIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getNumberOfResponsesByFormID: connect.NewClient[v1.GetNumberOfResponsesByFormIDRequest, v1.GetNumberOfResponsesByFormIDResponse](
			httpClient,
			baseURL+ResponseServiceGetNumberOfResponsesByFormIDProcedure,
			connect.WithSchema(responseServiceGetNumberOfResponsesByFormIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// responseServiceClient implements ResponseServiceClient.
type responseServiceClient struct {
	getNumberOfResponsesByStudentID *connect.Client[v1.GetNumberOfResponsesByStudentIDRequest, v1.GetNumberOfResponsesByStudentIDResponse]
	getNumberOfResponsesByFormID    *connect.Client[v1.GetNumberOfResponsesByFormIDRequest, v1.GetNumberOfResponsesByFormIDResponse]
}

// GetNumberOfResponsesByStudentID calls
// coteacher.v1.ResponseService.GetNumberOfResponsesByStudentID.
func (c *responseServiceClient) GetNumberOfResponsesByStudentID(ctx context.Context, req *connect.Request[v1.GetNumberOfResponsesByStudentIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByStudentIDResponse], error) {
	return c.getNumberOfResponsesByStudentID.CallUnary(ctx, req)
}

// GetNumberOfResponsesByFormID calls coteacher.v1.ResponseService.GetNumberOfResponsesByFormID.
func (c *responseServiceClient) GetNumberOfResponsesByFormID(ctx context.Context, req *connect.Request[v1.GetNumberOfResponsesByFormIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByFormIDResponse], error) {
	return c.getNumberOfResponsesByFormID.CallUnary(ctx, req)
}

// ResponseServiceHandler is an implementation of the coteacher.v1.ResponseService service.
type ResponseServiceHandler interface {
	GetNumberOfResponsesByStudentID(context.Context, *connect.Request[v1.GetNumberOfResponsesByStudentIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByStudentIDResponse], error)
	GetNumberOfResponsesByFormID(context.Context, *connect.Request[v1.GetNumberOfResponsesByFormIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByFormIDResponse], error)
}

// NewResponseServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewResponseServiceHandler(svc ResponseServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	responseServiceGetNumberOfResponsesByStudentIDHandler := connect.NewUnaryHandler(
		ResponseServiceGetNumberOfResponsesByStudentIDProcedure,
		svc.GetNumberOfResponsesByStudentID,
		connect.WithSchema(responseServiceGetNumberOfResponsesByStudentIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	responseServiceGetNumberOfResponsesByFormIDHandler := connect.NewUnaryHandler(
		ResponseServiceGetNumberOfResponsesByFormIDProcedure,
		svc.GetNumberOfResponsesByFormID,
		connect.WithSchema(responseServiceGetNumberOfResponsesByFormIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/coteacher.v1.ResponseService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ResponseServiceGetNumberOfResponsesByStudentIDProcedure:
			responseServiceGetNumberOfResponsesByStudentIDHandler.ServeHTTP(w, r)
		case ResponseServiceGetNumberOfResponsesByFormIDProcedure:
			responseServiceGetNumberOfResponsesByFormIDHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedResponseServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedResponseServiceHandler struct{}

func (UnimplementedResponseServiceHandler) GetNumberOfResponsesByStudentID(context.Context, *connect.Request[v1.GetNumberOfResponsesByStudentIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByStudentIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.ResponseService.GetNumberOfResponsesByStudentID is not implemented"))
}

func (UnimplementedResponseServiceHandler) GetNumberOfResponsesByFormID(context.Context, *connect.Request[v1.GetNumberOfResponsesByFormIDRequest]) (*connect.Response[v1.GetNumberOfResponsesByFormIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.ResponseService.GetNumberOfResponsesByFormID is not implemented"))
}
// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: coteacher/v1/form.proto

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
	// FormServiceName is the fully-qualified name of the FormService service.
	FormServiceName = "coteacher.v1.FormService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FormServiceCreateFormProcedure is the fully-qualified name of the FormService's CreateForm RPC.
	FormServiceCreateFormProcedure = "/coteacher.v1.FormService/CreateForm"
	// FormServiceGetFormByIDProcedure is the fully-qualified name of the FormService's GetFormByID RPC.
	FormServiceGetFormByIDProcedure = "/coteacher.v1.FormService/GetFormByID"
	// FormServiceGetFormListByClassIDProcedure is the fully-qualified name of the FormService's
	// GetFormListByClassID RPC.
	FormServiceGetFormListByClassIDProcedure = "/coteacher.v1.FormService/GetFormListByClassID"
	// FormServiceUpdateFormProcedure is the fully-qualified name of the FormService's UpdateForm RPC.
	FormServiceUpdateFormProcedure = "/coteacher.v1.FormService/UpdateForm"
	// FormServiceDeleteFormProcedure is the fully-qualified name of the FormService's DeleteForm RPC.
	FormServiceDeleteFormProcedure = "/coteacher.v1.FormService/DeleteForm"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	formServiceServiceDescriptor                    = v1.File_coteacher_v1_form_proto.Services().ByName("FormService")
	formServiceCreateFormMethodDescriptor           = formServiceServiceDescriptor.Methods().ByName("CreateForm")
	formServiceGetFormByIDMethodDescriptor          = formServiceServiceDescriptor.Methods().ByName("GetFormByID")
	formServiceGetFormListByClassIDMethodDescriptor = formServiceServiceDescriptor.Methods().ByName("GetFormListByClassID")
	formServiceUpdateFormMethodDescriptor           = formServiceServiceDescriptor.Methods().ByName("UpdateForm")
	formServiceDeleteFormMethodDescriptor           = formServiceServiceDescriptor.Methods().ByName("DeleteForm")
)

// FormServiceClient is a client for the coteacher.v1.FormService service.
type FormServiceClient interface {
	CreateForm(context.Context, *connect.Request[v1.CreateFormRequest]) (*connect.Response[v1.CreateFormResponse], error)
	GetFormByID(context.Context, *connect.Request[v1.GetFormByIDRequest]) (*connect.Response[v1.GetFormByIDResponse], error)
	GetFormListByClassID(context.Context, *connect.Request[v1.GetFormListByClassIDRequest]) (*connect.Response[v1.GetFormListByClassIDResponse], error)
	UpdateForm(context.Context, *connect.Request[v1.UpdateFormRequest]) (*connect.Response[v1.UpdateFormResponse], error)
	DeleteForm(context.Context, *connect.Request[v1.DeleteFormRequest]) (*connect.Response[v1.DeleteFormResponse], error)
}

// NewFormServiceClient constructs a client for the coteacher.v1.FormService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFormServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FormServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &formServiceClient{
		createForm: connect.NewClient[v1.CreateFormRequest, v1.CreateFormResponse](
			httpClient,
			baseURL+FormServiceCreateFormProcedure,
			connect.WithSchema(formServiceCreateFormMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getFormByID: connect.NewClient[v1.GetFormByIDRequest, v1.GetFormByIDResponse](
			httpClient,
			baseURL+FormServiceGetFormByIDProcedure,
			connect.WithSchema(formServiceGetFormByIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getFormListByClassID: connect.NewClient[v1.GetFormListByClassIDRequest, v1.GetFormListByClassIDResponse](
			httpClient,
			baseURL+FormServiceGetFormListByClassIDProcedure,
			connect.WithSchema(formServiceGetFormListByClassIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateForm: connect.NewClient[v1.UpdateFormRequest, v1.UpdateFormResponse](
			httpClient,
			baseURL+FormServiceUpdateFormProcedure,
			connect.WithSchema(formServiceUpdateFormMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteForm: connect.NewClient[v1.DeleteFormRequest, v1.DeleteFormResponse](
			httpClient,
			baseURL+FormServiceDeleteFormProcedure,
			connect.WithSchema(formServiceDeleteFormMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// formServiceClient implements FormServiceClient.
type formServiceClient struct {
	createForm           *connect.Client[v1.CreateFormRequest, v1.CreateFormResponse]
	getFormByID          *connect.Client[v1.GetFormByIDRequest, v1.GetFormByIDResponse]
	getFormListByClassID *connect.Client[v1.GetFormListByClassIDRequest, v1.GetFormListByClassIDResponse]
	updateForm           *connect.Client[v1.UpdateFormRequest, v1.UpdateFormResponse]
	deleteForm           *connect.Client[v1.DeleteFormRequest, v1.DeleteFormResponse]
}

// CreateForm calls coteacher.v1.FormService.CreateForm.
func (c *formServiceClient) CreateForm(ctx context.Context, req *connect.Request[v1.CreateFormRequest]) (*connect.Response[v1.CreateFormResponse], error) {
	return c.createForm.CallUnary(ctx, req)
}

// GetFormByID calls coteacher.v1.FormService.GetFormByID.
func (c *formServiceClient) GetFormByID(ctx context.Context, req *connect.Request[v1.GetFormByIDRequest]) (*connect.Response[v1.GetFormByIDResponse], error) {
	return c.getFormByID.CallUnary(ctx, req)
}

// GetFormListByClassID calls coteacher.v1.FormService.GetFormListByClassID.
func (c *formServiceClient) GetFormListByClassID(ctx context.Context, req *connect.Request[v1.GetFormListByClassIDRequest]) (*connect.Response[v1.GetFormListByClassIDResponse], error) {
	return c.getFormListByClassID.CallUnary(ctx, req)
}

// UpdateForm calls coteacher.v1.FormService.UpdateForm.
func (c *formServiceClient) UpdateForm(ctx context.Context, req *connect.Request[v1.UpdateFormRequest]) (*connect.Response[v1.UpdateFormResponse], error) {
	return c.updateForm.CallUnary(ctx, req)
}

// DeleteForm calls coteacher.v1.FormService.DeleteForm.
func (c *formServiceClient) DeleteForm(ctx context.Context, req *connect.Request[v1.DeleteFormRequest]) (*connect.Response[v1.DeleteFormResponse], error) {
	return c.deleteForm.CallUnary(ctx, req)
}

// FormServiceHandler is an implementation of the coteacher.v1.FormService service.
type FormServiceHandler interface {
	CreateForm(context.Context, *connect.Request[v1.CreateFormRequest]) (*connect.Response[v1.CreateFormResponse], error)
	GetFormByID(context.Context, *connect.Request[v1.GetFormByIDRequest]) (*connect.Response[v1.GetFormByIDResponse], error)
	GetFormListByClassID(context.Context, *connect.Request[v1.GetFormListByClassIDRequest]) (*connect.Response[v1.GetFormListByClassIDResponse], error)
	UpdateForm(context.Context, *connect.Request[v1.UpdateFormRequest]) (*connect.Response[v1.UpdateFormResponse], error)
	DeleteForm(context.Context, *connect.Request[v1.DeleteFormRequest]) (*connect.Response[v1.DeleteFormResponse], error)
}

// NewFormServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFormServiceHandler(svc FormServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	formServiceCreateFormHandler := connect.NewUnaryHandler(
		FormServiceCreateFormProcedure,
		svc.CreateForm,
		connect.WithSchema(formServiceCreateFormMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formServiceGetFormByIDHandler := connect.NewUnaryHandler(
		FormServiceGetFormByIDProcedure,
		svc.GetFormByID,
		connect.WithSchema(formServiceGetFormByIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formServiceGetFormListByClassIDHandler := connect.NewUnaryHandler(
		FormServiceGetFormListByClassIDProcedure,
		svc.GetFormListByClassID,
		connect.WithSchema(formServiceGetFormListByClassIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formServiceUpdateFormHandler := connect.NewUnaryHandler(
		FormServiceUpdateFormProcedure,
		svc.UpdateForm,
		connect.WithSchema(formServiceUpdateFormMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formServiceDeleteFormHandler := connect.NewUnaryHandler(
		FormServiceDeleteFormProcedure,
		svc.DeleteForm,
		connect.WithSchema(formServiceDeleteFormMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/coteacher.v1.FormService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FormServiceCreateFormProcedure:
			formServiceCreateFormHandler.ServeHTTP(w, r)
		case FormServiceGetFormByIDProcedure:
			formServiceGetFormByIDHandler.ServeHTTP(w, r)
		case FormServiceGetFormListByClassIDProcedure:
			formServiceGetFormListByClassIDHandler.ServeHTTP(w, r)
		case FormServiceUpdateFormProcedure:
			formServiceUpdateFormHandler.ServeHTTP(w, r)
		case FormServiceDeleteFormProcedure:
			formServiceDeleteFormHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFormServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFormServiceHandler struct{}

func (UnimplementedFormServiceHandler) CreateForm(context.Context, *connect.Request[v1.CreateFormRequest]) (*connect.Response[v1.CreateFormResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.FormService.CreateForm is not implemented"))
}

func (UnimplementedFormServiceHandler) GetFormByID(context.Context, *connect.Request[v1.GetFormByIDRequest]) (*connect.Response[v1.GetFormByIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.FormService.GetFormByID is not implemented"))
}

func (UnimplementedFormServiceHandler) GetFormListByClassID(context.Context, *connect.Request[v1.GetFormListByClassIDRequest]) (*connect.Response[v1.GetFormListByClassIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.FormService.GetFormListByClassID is not implemented"))
}

func (UnimplementedFormServiceHandler) UpdateForm(context.Context, *connect.Request[v1.UpdateFormRequest]) (*connect.Response[v1.UpdateFormResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.FormService.UpdateForm is not implemented"))
}

func (UnimplementedFormServiceHandler) DeleteForm(context.Context, *connect.Request[v1.DeleteFormRequest]) (*connect.Response[v1.DeleteFormResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("coteacher.v1.FormService.DeleteForm is not implemented"))
}
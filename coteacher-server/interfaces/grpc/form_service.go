package grpc

import (
	"context"

	"connectrpc.com/connect"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/form"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type formServiceServer struct {
	formInteractor *form.Interactor
}

func NewFormServiceServer(interactor *form.Interactor) coteacherv1connect.FormServiceHandler {
	return &formServiceServer{interactor}
}

// CreateForm implements coteacherv1connect.FormServiceHandler.
func (s *formServiceServer) CreateForm(ctx context.Context, req *connect.Request[coteacherv1.CreateFormRequest]) (*connect.Response[coteacherv1.CreateFormResponse], error) {
	resp, err := s.formInteractor.CreateForm(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// DeleteForm implements coteacherv1connect.FormServiceHandler.
func (*formServiceServer) DeleteForm(context.Context, *connect.Request[coteacherv1.DeleteFormRequest]) (*connect.Response[coteacherv1.DeleteFormResponse], error) {
	panic("unimplemented")
}

// GetFormByID implements coteacherv1connect.FormServiceHandler.
func (s *formServiceServer) GetFormByID(ctx context.Context, req *connect.Request[coteacherv1.GetFormByIDRequest]) (*connect.Response[coteacherv1.GetFormByIDResponse], error) {
	resp, err := s.formInteractor.GetFormByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// GetFormListByClassID implements coteacherv1connect.FormServiceHandler.
func (s *formServiceServer) GetFormListByClassID(ctx context.Context, req *connect.Request[coteacherv1.GetFormListByClassIDRequest]) (*connect.Response[coteacherv1.GetFormListByClassIDResponse], error) {
	resp, err := s.formInteractor.GetFormListByClassID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// UpdateForm implements coteacherv1connect.FormServiceHandler.
func (*formServiceServer) UpdateForm(context.Context, *connect.Request[coteacherv1.UpdateFormRequest]) (*connect.Response[coteacherv1.UpdateFormResponse], error) {
	panic("unimplemented")
}

// CheckFormEditPermission implements coteacherv1connect.FormServiceHandler.
func (s *formServiceServer) CheckFormEditPermission(ctx context.Context, req *connect.Request[coteacherv1.CheckFormEditPermissionRequest]) (*connect.Response[coteacherv1.CheckFormEditPermissionResponse], error) {
	resp, err := s.formInteractor.CheckFormEditPermission(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// CheckFormViewPermission implements coteacherv1connect.FormServiceHandler.
func (s *formServiceServer) CheckFormViewPermission(ctx context.Context, req *connect.Request[coteacherv1.CheckFormViewPermissionRequest]) (*connect.Response[coteacherv1.CheckFormViewPermissionResponse], error) {
	resp, err := s.formInteractor.CheckFormViewPermission(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

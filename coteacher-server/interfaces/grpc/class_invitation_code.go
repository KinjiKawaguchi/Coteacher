package grpc

import (
	"context"

	"connectrpc.com/connect"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class_invitation_code"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type classInvitationCodeServiceServer struct {
	interactor *class_invitation_code.Interactor
}

func NewClassInvitationCodeServiceServer(interactor *class_invitation_code.Interactor) coteacherv1connect.ClassInvitationCodeServiceHandler {
	return &classInvitationCodeServiceServer{interactor}
}

func (s *classInvitationCodeServiceServer) CreateClassInvitationCode(ctx context.Context, req *connect.Request[coteacherv1.CreateClassInvitationCodeRequest]) (*connect.Response[coteacherv1.CreateClassInvitationCodeResponse], error) {
	resp, err := s.interactor.CreateClassInvitationCode(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classInvitationCodeServiceServer) GetClassInvitationCodeByID(ctx context.Context, req *connect.Request[coteacherv1.GetClassInvitationCodeByIDRequest]) (*connect.Response[coteacherv1.GetClassInvitationCodeByIDResponse], error) {
	resp, err := s.interactor.GetClassInvitationCodeByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classInvitationCodeServiceServer) GetClassInvitationCodeListByClassID(ctx context.Context, req *connect.Request[coteacherv1.GetClassInvitationCodeListByClassIDRequest]) (*connect.Response[coteacherv1.GetClassInvitationCodeListByClassIDResponse], error) {
	resp, err := s.interactor.GetClassInvitationCodeListByClassID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classInvitationCodeServiceServer) UpdateClassInvitationCode(ctx context.Context, req *connect.Request[coteacherv1.UpdateClassInvitationCodeRequest]) (*connect.Response[coteacherv1.UpdateClassInvitationCodeResponse], error) {
	resp, err := s.interactor.UpdateClassInvitationCode(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

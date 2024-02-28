package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/student"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type studentServiceServer struct {
	interactor *student.Interactor
}

func NewStudentServiceServer(interactor *student.Interactor) coteacherv1connect.StudentServiceHandler {
	return &studentServiceServer{interactor}
}

func (s *studentServiceServer) CheckStudentExistsByID(ctx context.Context, req *connect.Request[coteacherv1.CheckStudentExistsByIDRequest]) (*connect.Response[coteacherv1.CheckStudentExistsByIDResponse], error) {
	resp, err := s.interactor.CheckStudentExistsByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *studentServiceServer) CheckStudentExistsByEmail(ctx context.Context, req *connect.Request[coteacherv1.CheckStudentExistsByEmailRequest]) (*connect.Response[coteacherv1.CheckStudentExistsByEmailResponse], error) {
	resp, err := s.interactor.CheckStudentExistsByEmail(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *studentServiceServer) ParticipateClass(ctx context.Context, req *connect.Request[coteacherv1.ParticipateClassRequest]) (*connect.Response[coteacherv1.ParticipateClassResponse], error) {
	resp, err := s.interactor.ParticipateClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *studentServiceServer) QuitClass(ctx context.Context, req *connect.Request[coteacherv1.QuitClassRequest]) (*connect.Response[coteacherv1.QuitClassResponse], error) {
	resp, err := s.interactor.QuitClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/student_class"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type studentClassServiceServer struct {
	interactor *student_class.Interactor
}

func NewStudentClassServiceServer(interactor *student_class.Interactor) coteacherv1connect.StudentClassServiceHandler {
	return &studentClassServiceServer{interactor}
}

func (s *studentClassServiceServer) CreateStudentClass(ctx context.Context, req *connect.Request[coteacherv1.CreateStudentClassRequest]) (*connect.Response[coteacherv1.CreateStudentClassResponse], error) {
	resp, err := s.interactor.CreateStudentClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *studentClassServiceServer) GetStudentListByClassID(ctx context.Context, req *connect.Request[coteacherv1.GetStudentListByClassIDRequest]) (*connect.Response[coteacherv1.GetStudentListByClassIDResponse], error) {
	resp, err := s.interactor.GetStudentListByClassID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *studentClassServiceServer) GetClassListByStudentID(ctx context.Context, req *connect.Request[coteacherv1.GetClassListByStudentIDRequest]) (*connect.Response[coteacherv1.GetClassListByStudentIDResponse], error) {
	resp, err := s.interactor.GetClassListByStudentID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *studentClassServiceServer) DeleteStudentClass(ctx context.Context, req *connect.Request[coteacherv1.DeleteStudentClassRequest]) (*connect.Response[coteacherv1.DeleteStudentClassResponse], error) {
	resp, err := s.interactor.DeleteStudentClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

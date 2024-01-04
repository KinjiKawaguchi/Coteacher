package grpc

import (
	"context"

	"connectrpc.com/connect"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type classServiceServer struct {
	classInteractor *class.Interactor
}

func NewClassServiceServer(interactor *class.Interactor) coteacherv1connect.ClassServiceHandler {
	return &classServiceServer{interactor}
}

func (s *classServiceServer) CreateClass(ctx context.Context, req *connect.Request[coteacherv1.CreateClassRequest]) (*connect.Response[coteacherv1.CreateClassResponse], error) {
	resp, err := s.classInteractor.CreateClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classServiceServer) GetClassByID(ctx context.Context, req *connect.Request[coteacherv1.GetClassByIDRequest]) (*connect.Response[coteacherv1.GetClassByIDResponse], error) {
	resp, err := s.classInteractor.GetClassByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classServiceServer) GetClassListByTeacherID(ctx context.Context, req *connect.Request[coteacherv1.GetClassListByTeacherIDRequest]) (*connect.Response[coteacherv1.GetClassListByTeacherIDResponse], error) {
	resp, err := s.classInteractor.GetClassListByTeacherID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classServiceServer) UpdateClass(ctx context.Context, req *connect.Request[coteacherv1.UpdateClassRequest]) (*connect.Response[coteacherv1.UpdateClassResponse], error) {
	resp, err := s.classInteractor.UpdateClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *classServiceServer) DeleteClass(ctx context.Context, req *connect.Request[coteacherv1.DeleteClassRequest]) (*connect.Response[coteacherv1.DeleteClassResponse], error) {
	resp, err := s.classInteractor.DeleteClass(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

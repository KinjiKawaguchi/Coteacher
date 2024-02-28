package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/teacher"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type teacherServiceServer struct {
	interactor *teacher.Interactor
}

func NewTeacherServiceServer(interactor *teacher.Interactor) coteacherv1connect.TeacherServiceHandler {
	return &teacherServiceServer{interactor}
}

// CheckTeacherExistsByID implements coteacherv1connect.TeacherServiceHandler.
func (s *teacherServiceServer) CheckTeacherExistsByID(ctx context.Context, req *connect.Request[coteacherv1.CheckTeacherExistsByIDRequest]) (*connect.Response[coteacherv1.CheckTeacherExistsByIDResponse], error) {
	resp, err := s.interactor.CheckTeacherExistsByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *teacherServiceServer) CheckTeacherExistsByEmail(ctx context.Context, req *connect.Request[coteacherv1.CheckTeacherExistsByEmailRequest]) (*connect.Response[coteacherv1.CheckTeacherExistsByEmailResponse], error) {
	resp, err := s.interactor.CheckTeacherExistsByEmail(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

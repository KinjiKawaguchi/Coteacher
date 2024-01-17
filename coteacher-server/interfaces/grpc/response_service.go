package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/response"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type responseServiceServer struct {
	interactor *response.Interactor
}
func NewResponseServiceServer(interactor *response.Interactor) coteacherv1connect.ResponseServiceHandler {
	return &responseServiceServer{interactor}
}

// GetNumberOfResponsesByFormID implements coteacherv1connect.ResponseServiceHandler.
func (s *responseServiceServer) GetNumberOfResponsesByFormID(ctx context.Context, req *connect.Request[coteacherv1.GetNumberOfResponsesByFormIDRequest]) (*connect.Response[coteacherv1.GetNumberOfResponsesByFormIDResponse], error) {
	resp, err := s.interactor.GetNumberOfResponsesByFormID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// GetNumberOfResponsesByStudentID implements coteacherv1connect.ResponseServiceHandler.
func (s *responseServiceServer) GetNumberOfResponsesByStudentID(ctx context.Context, req *connect.Request[coteacherv1.GetNumberOfResponsesByStudentIDRequest]) (*connect.Response[coteacherv1.GetNumberOfResponsesByStudentIDResponse], error) {
	resp, err := s.interactor.GetNumberOfResponsesByStudentID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}


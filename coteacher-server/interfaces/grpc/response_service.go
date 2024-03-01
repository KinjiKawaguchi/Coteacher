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

// GetResponseListByFormID implements coteacherv1connect.ResponseServiceHandler.
func (s *responseServiceServer) GetResponseListByFormID(ctx context.Context, req *connect.Request[coteacherv1.GetResponseListByFormIDRequest]) (*connect.Response[coteacherv1.GetResponseListByFormIDResponse], error) {
	resp, err := s.interactor.GetResponseListByFormID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// SubmitResponse implements coteacherv1connect.ResponseServiceHandler.
func (s *responseServiceServer) SubmitResponse(ctx context.Context, req *connect.Request[coteacherv1.SubmitResponseRequest]) (*connect.Response[coteacherv1.SubmitResponseResponse], error) {
	resp, err := s.interactor.SubmitResponse(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// SubmitAIResponse implements coteacherv1connect.ResponseServiceHandler.
func (s *responseServiceServer) SubmitAIResponse(ctx context.Context, req *connect.Request[coteacherv1.SubmitAIResponseRequest]) (*connect.Response[coteacherv1.SubmitAIResponseResponse], error) {
	resp, err := s.interactor.SubmitAIResponse(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// CreateDataset implements coteacherv1connect.ResponseServiceHandler.
func (s *responseServiceServer) CreateDataset(ctx context.Context, req *connect.Request[coteacherv1.CreateDatasetRequest]) (*connect.Response[coteacherv1.CreateDatasetResponse], error) {
	resp, err := s.interactor.CreateDataset(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

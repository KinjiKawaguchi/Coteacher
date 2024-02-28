package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/question"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type questionServiceServer struct {
	interactor *question.Interactor
}

func NewQuestionServiceServer(interactor *question.Interactor) coteacherv1connect.QuestionServiceHandler {
	return &questionServiceServer{interactor}
}

// SaveQuestionList implements coteacherv1connect.QuestionServiceHandler.
func (s *questionServiceServer) SaveQuestionList(ctx context.Context, req *connect.Request[coteacherv1.SaveQuestionListRequest]) (*connect.Response[coteacherv1.SaveQuestionListResponse], error) {
	resp, err := s.interactor.SaveQuestionList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// GetQuestionListByFormId implements coteacherv1connect.QuestionServiceHandler.
func (s *questionServiceServer) GetQuestionListByFormId(ctx context.Context, req *connect.Request[coteacherv1.GetQuestionListByFormIdRequest]) (*connect.Response[coteacherv1.GetQuestionListByFormIdResponse], error) {
	resp, err := s.interactor.GetQuestionListByFormID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

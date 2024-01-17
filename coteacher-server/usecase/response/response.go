package response

import (
	"context"

	"connectrpc.com/connect"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) GetNumberOfResponsesByStudentID(ctx context.Context, req *pb.GetNumberOfResponsesByStudentIDRequest) (*pb.GetNumberOfResponsesByStudentIDResponse, error) {
	// Get all forms by class id.
	studentID, err := uuid.Parse(req.StudentId)
	responses, err := i.entClient.Response.Query().Where(response.StudentID(studentID)).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 個数を返す
	return &pb.GetNumberOfResponsesByStudentIDResponse{
		NumberOfResponses: int32(len(responses)),
	}, nil
}

func (i *Interactor) GetNumberOfResponsesByFormID(ctx context.Context, req *pb.GetNumberOfResponsesByFormIDRequest) (*pb.GetNumberOfResponsesByFormIDResponse, error) {
	// Get all forms by class id.
	formID, err := uuid.Parse(req.FormId)
	responses, err := i.entClient.Response.Query().Where(response.FormID(formID)).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 個数を返す
	return &pb.GetNumberOfResponsesByFormIDResponse{
		NumberOfResponses: int32(len(responses)),
	}, nil
}

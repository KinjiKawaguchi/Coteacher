package teacher

import (
	"connectrpc.com/connect"
	"golang.org/x/exp/slog"

	"context"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	entteacher "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/teacher"
	entuser "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/user"

	"github.com/google/uuid"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CheckTeacherExistsByID(ctx context.Context, req *pb.CheckTeacherExistsByIDRequest) (*pb.CheckTeacherExistsByIDResponse, error) {
	teacherID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	exists, err := i.entClient.Teacher.Query().
		Where(entteacher.ID(teacherID)).
		Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CheckTeacherExistsByIDResponse{
		Exists: exists,
	}, nil
}

func (i *Interactor) CheckTeacherExistsByEmail(ctx context.Context, req *pb.CheckTeacherExistsByEmailRequest) (*pb.CheckTeacherExistsByEmailResponse, error) {
	//EmailがUserTableでマッチするUserIDががTeacherIDに存在するかを返す

	id := i.entClient.User.Query().
		Where(entuser.Email(req.Email)).
		Select(entuser.FieldID).
		FirstIDX(ctx)

	exists, err := i.entClient.Teacher.Query().
		Where(entteacher.ID(id)).
		Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CheckTeacherExistsByEmailResponse{
		Exists: exists,
	}, nil
}

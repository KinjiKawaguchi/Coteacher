package student

import (
	"time"

	"connectrpc.com/connect"
	"golang.org/x/exp/slog"

	"context"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	entclass "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	entcic "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/classinvitationcode"
	entstudent "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/student"
	entstudentclass "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/studentclass"
	entuser "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/user"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/utils"

	"github.com/google/uuid"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CheckStudentExistsByID(ctx context.Context, req *pb.CheckStudentExistsByIDRequest) (*pb.CheckStudentExistsByIDResponse, error) {
	studentID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	exists, err := i.entClient.Student.Query().
		Where(entstudent.ID(studentID)).
		Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CheckStudentExistsByIDResponse{
		Exists: exists,
	}, nil
}

func (i *Interactor) CheckStudentExistsByEmail(ctx context.Context, req *pb.CheckStudentExistsByEmailRequest) (*pb.CheckStudentExistsByEmailResponse, error) {
	//EmailがUserTableでマッチするUserIDががstudentIDに存在するかを返す

	id := i.entClient.User.Query().
		Where(entuser.Email(req.Email)).
		Select(entuser.FieldID).
		FirstIDX(ctx)

	exists, err := i.entClient.Student.Query().
		Where(entstudent.ID(id)).
		Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CheckStudentExistsByEmailResponse{
		Exists: exists,
	}, nil
}

func (i *Interactor) ParticipateClass(ctx context.Context, req *pb.ParticipateClassRequest) (*pb.ParticipateClassResponse, error) {
	studentID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	classInvitationCode, err := i.entClient.ClassInvitationCode.Query().
		Where(entcic.InvitationCodeEQ(req.InvitaionCode)).
		Where(entcic.ExpirationDateGT(time.Now())).
		Where(entcic.IsActiveEQ(true)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			// 有効な招待コードが見つからない
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 既に参加している場合はエラーを返す
	exists, err := i.entClient.StudentClass.
		Query().
		Where(entstudentclass.StudentID(studentID)).
		Where(entstudentclass.ClassID(classInvitationCode.ClassID)).
		Exist(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if exists {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	// 参加する
	_, err = i.entClient.StudentClass.Create().
		SetStudentID(studentID).
		SetClassID(classInvitationCode.ClassID).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	class, err := i.entClient.Class.Query().
		Where(entclass.ID(classInvitationCode.ClassID)).
		Only(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.ParticipateClassResponse{
		Class: utils.ToPbClass(class),
	}, nil
}

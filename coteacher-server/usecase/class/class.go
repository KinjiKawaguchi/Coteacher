package class

import (
	utils "coteacher/usecase/utils"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	"context"
	"coteacher/domain/repository/ent"
	entclass "coteacher/domain/repository/ent/class"
	pb "coteacher/proto-gen/go/coteacher/v1"

	"golang.org/x/exp/slog"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateClass(ctx context.Context, req *pb.CreateClassRequest) (*pb.CreateClassResponse, error) {

	now := time.Now()
	teacherID, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid teacher ID"))
	}

	q := i.entClient.Class.Create().
		SetName(req.Name).
		SetTeacherID(teacherID).
		SetCreatedAt(now).
		SetUpdatedAt(now)

	class, err := q.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the task was not found"))
		}
		i.logger.Error("failed to get the task", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CreateClassResponse{
		Class: utils.ToPbClass(class),
	}, nil
}

func (i *Interactor) GetClassByID(ctx context.Context, req *pb.GetClassByIDRequest) (*pb.GetClassByIDResponse, error) {
	q := i.entClient.Class.Query()
	class, err := q.Where(entclass.ID(uuid.MustParse(req.Id))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
	}
	return &pb.GetClassByIDResponse{
		Class: utils.ToPbClass(class),
	}, nil
}

func (i *Interactor) GetClassListByTeacherID(ctx context.Context, req *pb.GetClassListByTeacherIDRequest) (*pb.GetClassListByTeacherIDResponse, error) {
	teacherID, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid teacher ID"))
	}

	classes, err := i.entClient.Class.
		Query().
		Where(entclass.TeacherID(teacherID)).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("classes not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	pbClasses := make([]*pb.Class, len(classes))
	for i, class := range classes {
		pbClasses[i] = utils.ToPbClass(class)
	}

	return &pb.GetClassListByTeacherIDResponse{
		Classes: pbClasses,
	}, nil
}

func (i *Interactor) UpdateClass(ctx context.Context, req *pb.UpdateClassRequest) (*pb.UpdateClassResponse, error) {
	classID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid class ID"))
	}

	teacherID, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid teacher ID"))
	}

	class, err := i.entClient.Class.UpdateOneID(classID).
		SetName(req.Name).
		SetTeacherID(teacherID).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
	}
	return &pb.UpdateClassResponse{
		Class: utils.ToPbClass(class),
	}, nil
}

func (i *Interactor) DeleteClass(ctx context.Context, req *pb.DeleteClassRequest) (*pb.DeleteClassResponse, error) {
	classID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid class ID"))
	}

	err = i.entClient.Class.DeleteOneID(classID).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
	}
	return &pb.DeleteClassResponse{}, nil
}

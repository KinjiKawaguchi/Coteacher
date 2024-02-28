package class

import (
	"errors"
	"time"

	utils "github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/utils"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	"context"

	pb "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	entclass "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	entclassinvitationcode "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/classinvitationcode"
	entstudentclass "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/studentclass"

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
		return nil, connect.NewError(connect.CodeInternal, err)
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
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
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

	class, err := i.entClient.Class.Query().Where(entclass.ID(classID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 関係のあるclassinvitationcodeを削除
	_, err = i.entClient.ClassInvitationCode.Delete().Where(entclassinvitationcode.ClassID(classID)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// studentclassとの関係を削除
	_, err = i.entClient.StudentClass.Delete().Where(entstudentclass.ClassID(classID)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	err = i.entClient.Class.DeleteOneID(classID).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("class not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &pb.DeleteClassResponse{
		Class: utils.ToPbClass(class),
	}, nil
}

func (i *Interactor) CheckClassEditPermission(ctx context.Context, req *pb.CheckClassEditPermissionRequest) (*pb.CheckClassEditPermissionResponse, error) {
	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid class ID"))
	}

	teacherID, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid teacher ID"))
	}

	class, err := i.entClient.Class.Query().Where(entclass.ID(classID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &pb.CheckClassEditPermissionResponse{
				HasPermission: false,
			}, nil
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if class.TeacherID != teacherID {
		return &pb.CheckClassEditPermissionResponse{
			HasPermission: false,
		}, nil
	}

	return &pb.CheckClassEditPermissionResponse{
		HasPermission: true,
	}, nil
}

func (i *Interactor) CheckClassViewPermission(ctx context.Context, req *pb.CheckClassViewPermissionRequest) (*pb.CheckClassViewPermissionResponse, error) {
	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid class ID"))
	}

	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid student ID"))
	}

	studentClass, err := i.entClient.StudentClass.Query().Where(entstudentclass.ClassID(classID), entstudentclass.StudentID(studentID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &pb.CheckClassViewPermissionResponse{
				HasPermission: false,
			}, nil
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if studentClass == nil {
		return &pb.CheckClassViewPermissionResponse{
			HasPermission: false,
		}, nil
	}

	return &pb.CheckClassViewPermissionResponse{
		HasPermission: true,
	}, nil
}

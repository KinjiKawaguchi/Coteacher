package form

import (
	"context"
	"time"

	"connectrpc.com/connect"
	utils "github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/utils"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	form "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/form"
	studentClass "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/studentclass"

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

func (i *Interactor) GetFormListByClassID(ctx context.Context, req *pb.GetFormListByClassIDRequest) (*pb.GetFormListByClassIDResponse, error) {
	// Get all forms by class id.
	classID, err := uuid.Parse(req.ClassId)
	forms, err := i.entClient.Form.Query().Where(form.ClassID(classID)).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Convert to pb.
	pbForms := make([]*pb.Form, len(forms))
	for i, form := range forms {
		pbForms[i] = utils.ToPbForm(form)
	}

	return &pb.GetFormListByClassIDResponse{
		Forms: pbForms,
	}, nil
}

func (i *Interactor) CreateForm(ctx context.Context, req *pb.CreateFormRequest) (*pb.CreateFormResponse, error) {
	now := time.Now()
	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	form, err := i.entClient.Form.Create().
		SetClassID(classID).
		SetName(req.Name).
		SetDescription(req.Description).
		SetUsageLimit(int(req.UsageLimit)).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CreateFormResponse{
		Form: utils.ToPbForm(form),
	}, nil
}

func (i *Interactor) CheckFormEditPermission(ctx context.Context, req *pb.CheckFormEditPermissionRequest) (*pb.CheckFormEditPermissionResponse, error) {
	// Get form by id.
	formID, err := uuid.Parse(req.FormId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	teacherID, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	form, err := i.entClient.Form.Query().Where(form.ID(formID)).Only(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	class, err := i.entClient.Class.Query().Where(class.ID(form.ClassID)).Only(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if class.TeacherID != teacherID {
		return &pb.CheckFormEditPermissionResponse{
			HasPermission: false,
		}, nil
	}

	return &pb.CheckFormEditPermissionResponse{
		HasPermission: true,
	}, nil
}

func (i *Interactor) CheckFormViewPermission(ctx context.Context, req *pb.CheckFormViewPermissionRequest) (*pb.CheckFormViewPermissionResponse, error) {
	// Get form by id.
	formID, err := uuid.Parse(req.FormId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	form, err := i.entClient.Form.Query().Where(form.ID(formID)).Only(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	_, err = i.entClient.StudentClass.Query().
		Where(studentClass.
			ClassID(form.ClassID), studentClass.StudentID(studentID)).
		Exist(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return &pb.CheckFormViewPermissionResponse{
				HasPermission: false,
			}, nil
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.CheckFormViewPermissionResponse{
		HasPermission: true,
	}, nil
}

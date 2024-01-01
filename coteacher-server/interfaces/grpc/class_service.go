package grpc

import (
	"context"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"golang.org/x/exp/slog"
)

type classServiceServer struct {
	classInteractor *class.Interactor
}

func NewClassServiceServer(entClient *ent.Client, logger *slog.Logger) coteacherv1.ClassServiceServer {
	return &classServiceServer{
		classInteractor: class.NewInteractor(entClient, logger),
	}
}

// CreateClass implements coteacherv1.ClassServiceServer.
func (i *classServiceServer) CreateClass(ctx context.Context, req *coteacherv1.CreateClassRequest) (*coteacherv1.CreateClassResponse, error) {
	return i.classInteractor.CreateClass(ctx, req)
}

// GetClassByID implements coteacherv1.ClassServiceServer.
func (i *classServiceServer) GetClassByID(ctx context.Context, req *coteacherv1.GetClassByIDRequest) (*coteacherv1.GetClassByIDResponse, error) {
	return i.classInteractor.GetClassByID(ctx, req)
}

// GetClassListByTeacherID implements coteacherv1.ClassServiceServer.
func (i *classServiceServer) GetClassListByTeacherID(ctx context.Context, req *coteacherv1.GetClassListByTeacherIDRequest) (*coteacherv1.GetClassListByTeacherIDResponse, error) {
	return i.classInteractor.GetClassListByTeacherID(ctx, req)
}

// UpdateClass implements coteacherv1.ClassServiceServer.
func (i *classServiceServer) UpdateClass(ctx context.Context, req *coteacherv1.UpdateClassRequest) (*coteacherv1.UpdateClassResponse, error) {
	return i.classInteractor.UpdateClass(ctx, req)
}

// DeleteClass implements coteacherv1.ClassServiceServer.
func (i *classServiceServer) DeleteClass(ctx context.Context, req *coteacherv1.DeleteClassRequest) (*coteacherv1.DeleteClassResponse, error) {
	return i.classInteractor.DeleteClass(ctx, req)
}

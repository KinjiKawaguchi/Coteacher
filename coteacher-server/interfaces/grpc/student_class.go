package grpc

import (
	"context"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/student_class"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"golang.org/x/exp/slog"
)

type studentClassServiceServer struct {
	studentClassInteractor *student_class.Interactor
}

func NewStudentClassServiceServer(entClient *ent.Client, logger *slog.Logger) coteacherv1.StudentClassServiceServer {
	return &studentClassServiceServer{
		studentClassInteractor: student_class.NewInteractor(entClient, logger),
	}
}

// CreateStudentClass implements coteacherv1.StudentClassServiceServer.
func (i *studentClassServiceServer) CreateStudentClass(ctx context.Context, req *coteacherv1.CreateStudentClassRequest) (*coteacherv1.CreateStudentClassResponse, error) {
	return i.studentClassInteractor.CreateStudentClass(ctx, req)
}

// GetStudentListByCLassID implements coteacherv1.StudentClassServiceServer.
func (i *studentClassServiceServer) GetStudentListByClassID(ctx context.Context, req *coteacherv1.GetStudentListByClassIDRequest) (*coteacherv1.GetStudentListByClassIDResponse, error) {
	return i.studentClassInteractor.GetStudentListByClassID(ctx, req)
}

// GetClassListByStudentID implements coteacherv1.StudentClassServiceServer.
func (i *studentClassServiceServer) GetClassListByStudentID(ctx context.Context, req *coteacherv1.GetClassListByStudentIDRequest) (*coteacherv1.GetClassListByStudentIDResponse, error) {
	return i.studentClassInteractor.GetClassListByStudentID(ctx, req)
}

// DeleteStudentClass implements coteacherv1.StudentClassServiceServer.
func (i *studentClassServiceServer) DeleteStudentClass(ctx context.Context, req *coteacherv1.DeleteStudentClassRequest) (*coteacherv1.DeleteStudentClassResponse, error) {
	return i.studentClassInteractor.DeleteStudentClass(ctx, req)
}

package grpc

import (
	"context"

	"coteacher/domain/repository/ent"
	coteacherv1 "coteacher/proto-gen/go/coteacher/v1"
	"coteacher/usecase/student_class"

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

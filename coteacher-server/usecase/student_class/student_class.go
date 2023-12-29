package student_class

import (
	"time"

	"golang.org/x/exp/slog"

	utils "coteacher/usecase/utils"

	"context"
	"coteacher/domain/repository/ent"
	entclass "coteacher/domain/repository/ent/class"
	entstudentclass "coteacher/domain/repository/ent/studentclass"
	entuser "coteacher/domain/repository/ent/user"
	pb "coteacher/proto-gen/go/coteacher/v1"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateStudentClass(ctx context.Context, req *pb.CreateStudentClassRequest) (*pb.CreateStudentClassResponse, error) {
	now := time.Now()
	q := i.entClient.StudentClass.Create().
		SetStudentID(req.StudentId).
		SetClassID(req.ClassId).
		SetCreatedAt(now).
		SetUpdatedAt(now)

	student_class, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateStudentClassResponse{
		StudentClass: utils.ToPbStudentClass(student_class),
	}, nil
}

func (i *Interactor) GetStudentListByClassID(ctx context.Context, req *pb.GetStudentListByClassIDRequest) (*pb.GetStudentListByClassIDResponse, error) {
	studentclasses, err := i.entClient.StudentClass.
		Query().
		Where(entstudentclass.ClassID(req.ClassId)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	studentIDs := make([]string, len(studentclasses))
	for i, studentclass := range studentclasses {
		studentIDs[i] = studentclass.StudentID
	}

	students, err := i.entClient.User.
		Query().
		Where(entuser.IDIn(studentIDs...)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	pbStudents := make([]*pb.User, len(students))
	for i, student := range students {
		pbStudents[i] = utils.ToPbUser(student)
	}

	return &pb.GetStudentListByClassIDResponse{
		Students: pbStudents,
	}, nil
}

func (i *Interactor) GetClassListByStudentID(ctx context.Context, req *pb.GetClassListByStudentIDRequest) (*pb.GetClassListByStudentIDResponse, error) {
	studentclasses, err := i.entClient.StudentClass.
		Query().
		Where(entstudentclass.StudentID(req.StudentId)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	classIds := make([]string, len(studentclasses))
	for i, studentclass := range studentclasses {
		classIds[i] = studentclass.ClassID
	}

	classes, err := i.entClient.Class.
		Query().
		Where(entclass.IDIn(classIds...)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	pbClasses := make([]*pb.Class, len(classes))
	for i, class := range classes {
		pbClasses[i] = utils.ToPbClass(class)
	}

	return &pb.GetClassListByStudentIDResponse{
		Classes: pbClasses,
	}, nil
}

func (i *Interactor) DeleteStudentClass(ctx context.Context, req *pb.DeleteStudentClassRequest) (*pb.DeleteStudentClassResponse, error) {
	err := i.entClient.StudentClass.DeleteOne(&ent.StudentClass{
		StudentID: req.StudentId,
		ClassID:   req.ClassId,
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentClassResponse{}, nil
}

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

	"github.com/google/uuid"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateStudentClass(ctx context.Context, req *pb.CreateStudentClassRequest) (*pb.CreateStudentClassResponse, error) {

	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, err
	}

	q := i.entClient.StudentClass.Create().
		SetStudentID(studentID).
		SetClassID(classID).
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
	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, err
	}

	studentclasses, err := i.entClient.StudentClass.
		Query().
		Where(entstudentclass.ClassID(classID)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	studentIDs := make([]uuid.UUID, len(studentclasses))
	for i, studentclass := range studentclasses {
		studentID, err := uuid.Parse(studentclass.StudentID.String())
		if err != nil {
			return nil, err
		}
		studentIDs[i] = studentID
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
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, err
	}

	studentclasses, err := i.entClient.StudentClass.
		Query().
		Where(entstudentclass.StudentID(studentID)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	classIDs := make([]uuid.UUID, len(studentclasses))
	for i, studentclass := range studentclasses {
		classIDs[i], err = uuid.Parse(studentclass.ClassID.String())
		if err != nil {
			return nil, err
		}
	}

	classes, err := i.entClient.Class.
		Query().
		Where(entclass.IDIn(classIDs...)).
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
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, err
	}

	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, err
	}

	err = i.entClient.StudentClass.DeleteOne(&ent.StudentClass{
		StudentID: studentID,
		ClassID:   classID,
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentClassResponse{}, nil
}

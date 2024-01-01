package user

import (
	"errors"
	"time"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class"

	utils "github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/utils"

	"connectrpc.com/connect"

	"context"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	entstudent "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/student"
	entteacher "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/teacher"
	entuser "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/user"

	"golang.org/x/exp/slog"

	"github.com/google/uuid"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	now := time.Now()
	// Start a transaction because you may need to create a user and a teacher/student.
	tx, err := i.entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Create a base user.
	user, err := tx.User.
		Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	switch {
	case req.UserType == pb.UserType_USER_TYPE_TEACHER:
		_, err = tx.Teacher.
			Create().
			SetID(user.ID).
			SetUser(user). // Set the user edge here
			Save(ctx)
	case req.UserType == pb.UserType_USER_TYPE_STUDENT:
		_, err = tx.Student.
			Create().
			SetID(user.ID).
			SetUser(user). // Set the user edge here
			Save(ctx)
	default:
		tx.Rollback()
		return nil, errors.New("user role is not specified")
	}
	if err != nil {
		return nil, err
	}

	// Commit the transaction if everything is successful.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Return the created user.
	return &pb.CreateUserResponse{
		User:     utils.ToPbUser(user),
		UserType: getUserType(i, ctx, user)}, nil
}

func (i *Interactor) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	q := i.entClient.User.Query()
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	user, err := q.Where(entuser.ID(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
	}

	return &pb.GetUserByIDResponse{
		User:     utils.ToPbUser(user),
		UserType: getUserType(i, ctx, user),
	}, nil
}

func (i *Interactor) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	q := i.entClient.User.Query()
	user, err := q.Where(entuser.Email(req.Email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
	}
	return &pb.GetUserByEmailResponse{
		User:     utils.ToPbUser(user),
		UserType: getUserType(i, ctx, user),
	}, nil
}

func (i *Interactor) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	now := time.Now()
	tx, err := i.entClient.Tx(ctx)
	id, err := uuid.Parse(req.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user, err := i.entClient.User.UpdateOneID(id).
		SetName(req.Name).
		SetEmail(req.Email).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := i.updateUserRole(ctx, user, req); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		User:     utils.ToPbUser(user),
		UserType: getUserType(i, ctx, user),
	}, nil
}

func (i *Interactor) updateUserRole(ctx context.Context, user *ent.User, req *pb.UpdateUserRequest) error {
	switch {
	case req.UserType == pb.UserType_USER_TYPE_TEACHER:
		return i.updateTeacherRole(ctx, user, req)
	case req.UserType == pb.UserType_USER_TYPE_STUDENT:
		return i.updateStudentRole(ctx, user, req)
	}
	return nil
}

func (i *Interactor) updateTeacherRole(ctx context.Context, user *ent.User, req *pb.UpdateUserRequest) error {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return err
	}
	exists := i.entClient.Teacher.Query().Where(entteacher.ID(id)).ExistX(ctx)
	if !exists {
		// 生徒データが存在する場合は削除する
		if i.entClient.Student.Query().Where(entstudent.ID(id)).ExistX(ctx) {
			err := i.entClient.Student.DeleteOneID(id).Exec(ctx)
			if err != nil {
				return err
			}
		}

		// 教師として新しくデータを作成する
		_, err := i.entClient.Teacher.
			Create().
			SetID(user.ID).
			SetUser(user). // Set the user edge here
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interactor) updateStudentRole(ctx context.Context, user *ent.User, req *pb.UpdateUserRequest) error {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return err
	}
	exists := i.entClient.Student.Query().Where(entstudent.ID(id)).ExistX(ctx)
	if !exists {
		// 教師データが存在する場合は削除する
		if i.entClient.Teacher.Query().Where(entteacher.ID(id)).ExistX(ctx) {
			err := i.entClient.Teacher.DeleteOneID(id).Exec(ctx)
			if err != nil {
				return err
			}
		}

		// 生徒として新しくデータを作成する
		_, err := i.entClient.Student.
			Create().
			SetID(user.ID).
			SetUser(user). // Set the user edge here
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Interactor) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id, err := uuid.Parse(req.Id)

	if err != nil {
		return nil, err
	}

	user, err := i.entClient.User.Query().Where(entuser.ID(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
	}
	userType := getUserType(i, ctx, user)

	switch {
	case userType == pb.UserType_USER_TYPE_TEACHER:
		err = i.entClient.Teacher.DeleteOneID(id).Exec(ctx)
		class.NewInteractor(i.entClient, i.logger).DeleteClass(ctx, &pb.DeleteClassRequest{
			Id: req.Id,
		})
	case userType == pb.UserType_USER_TYPE_STUDENT:
		err = i.entClient.Student.DeleteOneID(id).Exec(ctx)
	default:
		return nil, errors.New("user role is not specified")
	}

	if err != nil {
		return nil, err
	}

	err = i.entClient.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		User:     utils.ToPbUser(user),
		UserType: userType,
	}, nil
}

func getUserType(i *Interactor, ctx context.Context, user *ent.User) pb.UserType {
	switch {
	case i.entClient.Teacher.Query().Where(entteacher.ID(user.ID)).ExistX(ctx):
		return pb.UserType_USER_TYPE_TEACHER
	case i.entClient.Student.Query().Where(entstudent.ID(user.ID)).ExistX(ctx):
		return pb.UserType_USER_TYPE_STUDENT
	default:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	}
}

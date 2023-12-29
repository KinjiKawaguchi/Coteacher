package user

import (
	"errors"
	"time"

	utils "coteacher/usecase/utils"

	"connectrpc.com/connect"

	"context"
	"coteacher/domain/repository/ent"
	entstudent "coteacher/domain/repository/ent/student"
	entteacher "coteacher/domain/repository/ent/teacher"
	entuser "coteacher/domain/repository/ent/user"
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

	// Check what type of user is being created and set the appropriate fields.
	switch {
	case req.GetTeacher() != nil:
		_, err = tx.Teacher.
			Create().
			SetID(user.ID).
			// Set additional fields specific to the Teacher here.
			Save(ctx)
	case req.GetStudent() != nil:
		_, err = tx.Student.
			Create().
			SetID(user.ID).
			// Set additional fields specific to the Student here.
			Save(ctx)
	default:
		// Rollback the transaction if no role is provided.
		tx.Rollback()
		return nil, errors.New("user role is not specified")
	}

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction if everything is successful.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Return the created user.
	return &pb.CreateUserResponse{User: utils.ToPbUser(user)}, nil
}

func (i *Interactor) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	q := i.entClient.User.Query()
	user, err := q.Where(entuser.ID(req.Id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
	}
	return &pb.GetUserByIDResponse{
		User: utils.ToPbUser(user),
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
		User: utils.ToPbUser(user),
	}, nil
}

func (i *Interactor) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	now := time.Now()
	tx, err := i.entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	user, err := i.entClient.User.UpdateOneID(req.Id).
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
		User: utils.ToPbUser(user),
	}, nil
}

func (i *Interactor) updateUserRole(ctx context.Context, user *ent.User, req *pb.UpdateUserRequest) error {
	switch {
	case req.GetTeacher() != nil:
		return i.updateTeacherRole(ctx, user, req)
	case req.GetStudent() != nil:
		return i.updateStudentRole(ctx, user, req)
	}
	return nil
}

func (i *Interactor) updateTeacherRole(ctx context.Context, user *ent.User, req *pb.UpdateUserRequest) error {
	exists := i.entClient.Teacher.Query().Where(entteacher.ID(req.Id)).ExistX(ctx)
	if exists {
		// 既存の教師データを更新するロジック（必要に応じて）
	} else {
		// 教師として新しくデータを作成する
		_, err := i.entClient.Teacher.
			Create().
			SetID(user.ID).
			// ここで教師固有のフィールドを設定
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interactor) updateStudentRole(ctx context.Context, user *ent.User, req *pb.UpdateUserRequest) error {
	exists := i.entClient.Student.Query().Where(entstudent.ID(req.Id)).ExistX(ctx)
	if exists {
		// 既存の生徒データを更新するロジック（必要に応じて）
	} else {
		// 生徒として新しくデータを作成する
		_, err := i.entClient.Student.
			Create().
			SetID(user.ID).
			// ここで生徒固有のフィールドを設定
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interactor) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	tx, err := i.entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	err = i.entClient.User.DeleteOneID(req.Id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
		tx.Rollback()
	}
	if i.entClient.Student.Query().Where(entstudent.ID(req.Id)).ExistX(ctx) {
		err := i.entClient.Student.DeleteOneID(req.Id).Exec(ctx)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}
	if i.entClient.Teacher.Query().Where(entteacher.ID(req.Id)).ExistX(ctx) {
		err := i.entClient.Teacher.DeleteOneID(req.Id).Exec(ctx)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{}, nil
}

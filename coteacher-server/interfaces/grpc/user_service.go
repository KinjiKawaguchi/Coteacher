package grpc

import (
	"context"

	"coteacher/domain/repository/ent"
	coteacherv1 "coteacher/proto-gen/go/coteacher/v1"
	"coteacher/usecase/user"

	"golang.org/x/exp/slog"
)

type userServiceServer struct {
	userInteractor *user.Interactor
}

func NewUserServiceServer(entClient *ent.Client, logger *slog.Logger) coteacherv1.UserServiceServer {
	return &userServiceServer{
		userInteractor: user.NewInteractor(entClient, logger),
	}
}

// CreateUser implements coteacherv1.UserServiceServer.
func (i *userServiceServer) CreateUser(ctx context.Context, req *coteacherv1.CreateUserRequest) (*coteacherv1.CreateUserResponse, error) {
	return i.userInteractor.CreateUser(ctx, req)
}

// DeleteUser implements coteacherv1.UserServiceServer.
func (i *userServiceServer) DeleteUser(ctx context.Context, req *coteacherv1.DeleteUserRequest) (*coteacherv1.DeleteUserResponse, error) {
	return i.userInteractor.DeleteUser(ctx, req)
}

// GetUserByEmail implements coteacherv1.UserServiceServer.
func (i *userServiceServer) GetUserByEmail(ctx context.Context, req *coteacherv1.GetUserByEmailRequest) (*coteacherv1.GetUserByEmailResponse, error) {
	return i.userInteractor.GetUserByEmail(ctx, req)
}

// GetUserByID implements coteacherv1.UserServiceServer.
func (i *userServiceServer) GetUserByID(ctx context.Context, req *coteacherv1.GetUserByIDRequest) (*coteacherv1.GetUserByIDResponse, error) {
	return i.userInteractor.GetUserByID(ctx, req)
}

// UpdateUser implements coteacherv1.UserServiceServer.
func (i *userServiceServer) UpdateUser(ctx context.Context, req *coteacherv1.UpdateUserRequest) (*coteacherv1.UpdateUserResponse, error) {
	return i.userInteractor.UpdateUser(ctx, req)
}

package grpc

import (
	"context"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/user"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

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

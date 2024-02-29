package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/user"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type userServiceServer struct {
	interactor *user.Interactor
}

func NewUserServiceServer(interactor *user.Interactor) coteacherv1connect.UserServiceHandler {
	return &userServiceServer{interactor}
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *connect.Request[coteacherv1.CreateUserRequest]) (*connect.Response[coteacherv1.CreateUserResponse], error) {
	resp, err := s.interactor.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) GetUserByID(ctx context.Context, req *connect.Request[coteacherv1.GetUserByIDRequest]) (*connect.Response[coteacherv1.GetUserByIDResponse], error) {
	resp, err := s.interactor.GetUserByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) GetUserByEmail(ctx context.Context, req *connect.Request[coteacherv1.GetUserByEmailRequest]) (*connect.Response[coteacherv1.GetUserByEmailResponse], error) {
	resp, err := s.interactor.GetUserByEmail(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) UpdateUser(ctx context.Context, req *connect.Request[coteacherv1.UpdateUserRequest]) (*connect.Response[coteacherv1.UpdateUserResponse], error) {
	resp, err := s.interactor.UpdateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) DeleteUser(ctx context.Context, req *connect.Request[coteacherv1.DeleteUserRequest]) (*connect.Response[coteacherv1.DeleteUserResponse], error) {
	resp, err := s.interactor.DeleteUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) CheckUserExistsByEmail(ctx context.Context, req *connect.Request[coteacherv1.CheckUserExistsByEmailRequest]) (*connect.Response[coteacherv1.CheckUserExistsByEmailResponse], error) {
	resp, err := s.interactor.CheckUserExistsByEmail(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

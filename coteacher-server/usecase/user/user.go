package usecase

import (
	"log/slog"

	"coteacher/domain/repository/ent"
	entuser "coteacher/domain/repository/ent/user"
	pb "coteacher/proto-gen/go/coteacher/v1"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var entUserType string
	switch req.UserType {
	case pb.UserType_USER_TYPE_TEACHER:
		entUserType = "teacher"
	case pb.UserType_USER_TYPE_STUDENT:
		entUserType = "student"
	default:
		entUserType = "" // Handle default case or return an error
	}

	user, err := i.entClient.User.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetUserType(entUserType).
		Save(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Map the Ent user type to the protobuf UserType
	var pbUserType pb.UserType
	switch user.UserType {
	case "teacher":
		pbUserType = pb.UserType_USER_TYPE_TEACHER
	case "student":
		pbUserType = pb.UserType_USER_TYPE_STUDENT
	default:
		pbUserType = pb.UserType_USER_TYPE_UNSPECIFIED
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			UserType: pbUserType,
		},
	}
}

func (i *Interactor) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	// Query the user from the database
	user, err := i.entClient.User.Query().Where(entuser.ID(req.Id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the user is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Map the Ent user type to the protobuf UserType
	var userType pb.UserType
	switch user.UserType {
	case "teacher":
		userType = pb.UserType_USER_TYPE_TEACHER
	case "student":
		userType = pb.UserType_USER_TYPE_STUDENT
	default:
		userType = pb.UserType_USER_TYPE_UNSPECIFIED
	}

	// Create the response
	return &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			UserType: userType, // Add the user type to the response
		},
	}, nil
}

func (i *Interactor) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	user, err := i.entClient.User.Query().Where(entuser.Email(req.Email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the user is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Map the Ent user type to the protobuf UserType
	var userType pb.UserType
	switch user.UserType {
	case "teacher":
		userType = pb.UserType_USER_TYPE_TEACHER
	case "student":
		userType = pb.UserType_USER_TYPE_STUDENT
	default:
		userType = pb.UserType_USER_TYPE_UNSPECIFIED
	}

	// Create the response
	return &pb.GetUserByEmailResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			UserType: userType, // Add the user type to the response
		},
	}, nil
}

func (i *Interactor) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// Query the user from the database
	user, err := i.entClient.User.Query().Where(entuser.ID(req.Id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the user is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Update the user
	user, err = user.Update().
		SetName(req.Name).
		SetEmail(req.Email).
		SetUserType(req.UserType). // Set the user type
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Map the Ent user type to the protobuf UserType
	var userType pb.UserType
	switch user.UserType {
	case "teacher":
		userType = pb.UserType_USER_TYPE_TEACHER
	case "student":
		userType = pb.UserType_USER_TYPE_STUDENT
	default:
		userType = pb.UserType_USER_TYPE_UNSPECIFIED
	}

	// Create the response
	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			UserType: userType, // Add the user type to the response
		},
	}, nil
}

func (i *Interactor) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Delete the user from the database
	err := i.entClient.User.DeleteOneID(req.Id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the user is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Create the response
	return &pb.DeleteUserResponse{}, nil
}

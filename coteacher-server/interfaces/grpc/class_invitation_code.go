package grpc

import (
	"context"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class_invitation_code"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"golang.org/x/exp/slog"
)

type classInvitationCodeServiceServer struct {
	classInvitationCodeInteractor *class_invitation_code.Interactor
}

func NewClassInvitationCodeServiceServer(entClient *ent.Client, logger *slog.Logger) coteacherv1.ClassInvitationCodeServiceServer {
	return &classInvitationCodeServiceServer{
		classInvitationCodeInteractor: class_invitation_code.NewInteractor(entClient, logger),
	}
}

// CreateClassInvitationCode implements coteacherv1.ClassInvitationCodeServiceServer.
func (i *classInvitationCodeServiceServer) CreateClassInvitationCode(ctx context.Context, req *coteacherv1.CreateClassInvitationCodeRequest) (*coteacherv1.CreateClassInvitationCodeResponse, error) {
	return i.classInvitationCodeInteractor.CreateClassInvitationCode(ctx, req)
}

// GetClassInvitationCodeByID implements coteacherv1.ClassInvitationCodeServiceServer.
func (i *classInvitationCodeServiceServer) GetClassInvitationCodeByID(ctx context.Context, req *coteacherv1.GetClassInvitationCodeByIDRequest) (*coteacherv1.GetClassInvitationCodeByIDResponse, error) {
	return i.classInvitationCodeInteractor.GetClassInvitationCodeByID(ctx, req)
}

// GetClassInvitationCodeListByClassID implements coteacherv1.ClassInvitationCodeServiceServer.
func (i *classInvitationCodeServiceServer) GetClassInvitationCodeListByClassID(ctx context.Context, req *coteacherv1.GetClassInvitationCodeListByClassIDRequest) (*coteacherv1.GetClassInvitationCodeListByClassIDResponse, error) {
	return i.classInvitationCodeInteractor.GetClassInvitationCodeListByClassID(ctx, req)
}

// UpdateClassInvitationCode implements coteacherv1.ClassInvitationCodeServiceServer.
func (i *classInvitationCodeServiceServer) UpdateClassInvitationCode(ctx context.Context, req *coteacherv1.UpdateClassInvitationCodeRequest) (*coteacherv1.UpdateClassInvitationCodeResponse, error) {
	return i.classInvitationCodeInteractor.UpdateClassInvitationCode(ctx, req)
}

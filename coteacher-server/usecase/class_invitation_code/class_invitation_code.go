package class_invitation_code

import (
	"time"

	"golang.org/x/exp/slog"

	utils "coteacher/usecase/utils"

	"context"
	"coteacher/domain/repository/ent"
	entcic "coteacher/domain/repository/ent/classinvitationcode"
	pb "coteacher/proto-gen/go/coteacher/v1"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateClassInvitationCode(ctx context.Context, req *pb.CreateClassInvitationCodeRequest) (*pb.CreateClassInvitationCodeResponse, error) {
	now := time.Now()
	cic, err := i.entClient.ClassInvitationCode.Create().
		SetClassID(req.ClassId).
		SetExpirationDate(req.ExpirationDate.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateClassInvitationCodeResponse{
		ClassInvitationCode: utils.ToPbClassInvitationCode(cic),
	}, nil
}

func (i *Interactor) GetClassInvitationCodeByID(ctx context.Context, req *pb.GetClassInvitationCodeByIDRequest) (*pb.GetClassInvitationCodeByIDResponse, error) {
	cic, err := i.entClient.ClassInvitationCode.Query().Where(entcic.ID(req.Id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetClassInvitationCodeByIDResponse{
		ClassInvitationCode: utils.ToPbClassInvitationCode(cic),
	}, nil
}

func (i *Interactor) GetClassInvitationCodeListByClassID(ctx context.Context, req *pb.GetClassInvitationCodeListByClassIDRequest) (*pb.GetClassInvitationCodeListByClassIDResponse, error) {
	cicList, err := i.entClient.ClassInvitationCode.Query().Where(entcic.ClassID(req.ClassId)).All(ctx)
	if err != nil {
		return nil, err
	}

	pbCicList := make([]*pb.ClassInvitationCode, len(cicList))
	for i, cic := range cicList {
		pbCicList[i] = utils.ToPbClassInvitationCode(cic)
	}

	return &pb.GetClassInvitationCodeListByClassIDResponse{
		ClassInvitationCodeList: pbCicList,
	}, nil
}

func (i *Interactor) UpdateClassInvitationCode(ctx context.Context, req *pb.UpdateClassInvitationCodeRequest) (*pb.UpdateClassInvitationCodeResponse, error) {
	cic, err := i.entClient.ClassInvitationCode.UpdateOneID((req.Id)).
		SetIsActive(req.IsActive).
		SetExpirationDate(req.ExpirationDate.AsTime()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateClassInvitationCodeResponse{
		ClassInvitationCode: utils.ToPbClassInvitationCode(cic),
	}, nil
}

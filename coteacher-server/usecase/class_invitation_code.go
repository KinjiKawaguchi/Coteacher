package usecase

import (
	"coteacher/domain/repository/ent"
	classinvitationcode "coteacher/domain/repository/ent/classinvitationcode"
	pb "coteacher/proto-gen/go/coteacher/v1"
	"log/slog"

	"golang.org/x/net/context"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) CreateClassInvitationCodeRequest(ctx context.Context, req *pb.CreateClassInvitationCodeRequest) *pb.CreateClassInvitationCodeRes {
	var entIsActive bool
	switch req.IsActive {
	case pb.IsActive_IS_ACTIVE_TRUE:
		entIsActive = true
	case pb.IsActive_IS_ACTIVE_FALSE:
		entIsActive = false
	default:
		entIsActive = false
	}

	classInvitationCode, err := i.entClient.ClassInvitationCode.Create().
		SetInvitationCode(req.InvitationCode).
		SetExpirationDate(req.ExpirationDate.AsTime()).
		SetIsActive(entIsActive).
		Save(ctx)

	if err != nil {
		return &pb.CreateClassInvitationCodeRes{
			Success: false,
			Message: err.Error(),
		}
	}

	return &pb.CreateClassInvitationCodeRes{
		Success: true,
		Message: "success",
		ClassInvitationCode: &pb.ClassInvitationCode{
			Id:             classInvitationCode.ID,
			InvitationCode: classInvitationCode.InvitationCode,
			ExpirationDate: &pb.Timestamp{
				Seconds: classInvitationCode.ExpirationDate.Unix(),
			},
			IsActive: classInvitationCode.IsActive,
		},
	}
}

func (i *Interactor) GetClassInvitationCodeByIDRequest(ctx context.Context, req *pb.GetClassInvitationCodeByIDRequest) *pb.GetClassInvitationCodeByIDResponse {
	classInvitationCode, err := i.entClient.ClassInvitationCode.Query().
		Where(classinvitationcode.ID(req.Id)).
		Only(ctx)

	if err != nil {
		return &pb.GetClassInvitationCodeByIDResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &pb.GetClassInvitationCodeByIDResponse{
		Success: true,
		Message: "success",
		ClassInvitationCode: &pb.ClassInvitationCode{
			Id:             classInvitationCode.ID,
			InvitationCode: classInvitationCode.InvitationCode,
			ExpirationDate: &pb.Timestamp{
				Seconds: classInvitationCode.ExpirationDate.Unix(),
			},
			IsActive: classInvitationCode.IsActive,
		},
	}
}

func (i *Interactor) GetClassInvitationCodeListByClassIDRequest(ctx context.COntext, req *pb.GetClassInvitationCodeListByClassIDRequest) *pb.GetClassInvitationCodeListByClassIDResponse {
	classInvitationCodeList, err := i.entClient.ClassInvitationCode.Query().
		Where(classinvitationcode.ClassID(req.ClassId)).
		All(ctx)

	if err != nil {
		return &pb.GetClassInvitationCodeListByClassIDResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	var pbClassInvitationCodeList []*pb.ClassInvitationCode
	for _, classInvitationCode := range classInvitationCodeList {
		pbClassInvitationCodeList = append(pbClassInvitationCodeList, &pb.ClassInvitationCode{
			Id:             classInvitationCode.ID,
			InvitationCode: classInvitationCode.InvitationCode,
			ExpirationDate: &pb.Timestamp{
				Seconds: classInvitationCode.ExpirationDate.Unix(),
			},
			IsActive: classInvitationCode.IsActive,
		})
	}

	return &pb.GetClassInvitationCodeListByClassIDResponse{
		Success:                 true,
		Message:                 "success",
		ClassInvitationCodeList: pbClassInvitationCodeList,
	}
}

func (i *Interactor) UpdateClassInvitationCodeRequest(ctx context.Context, req *pb.UpdateClassInvitationCodeRequest) *pb.UpdateClassInvitationCodeResponse {
	var entIsActive bool
	switch req.IsActive {
	case pb.IsActive_IS_ACTIVE_TRUE:
		entIsActive = true
	case pb.IsActive_IS_ACTIVE_FALSE:
		entIsActive = false
	default:
		entIsActive = false
	}

	classInvitationCode, err := i.entClient.ClassInvitationCode.UpdateOneID(req.Id).
		SetInvitationCode(req.InvitationCode).
		SetExpirationDate(req.ExpirationDate.AsTime()).
		SetIsActive(entIsActive).
		Save(ctx)

	if err != nil {
		return &pb.UpdateClassInvitationCodeResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &pb.UpdateClassInvitationCodeResponse{
		Success: true,
		Message: "success",
		ClassInvitationCode: &pb.ClassInvitationCode{
			Id:             classInvitationCode.ID,
			InvitationCode: classInvitationCode.InvitationCode,
			ExpirationDate: &pb.Timestamp{
				Seconds: classInvitationCode.ExpirationDate.Unix(),
			},
			IsActive: classInvitationCode.IsActive,
		},
	}
}

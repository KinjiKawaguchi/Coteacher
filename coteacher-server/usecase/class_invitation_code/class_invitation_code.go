package class_invitation_code

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
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

	var invitationCode string
	for {
		// Generate a random 8-character string
		invitationCode = RandomString(8)

		// Check if the code is unique
		exists, err := i.entClient.ClassInvitationCode.Query().
			Where(entcic.InvitationCode(invitationCode)).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if !exists {
			break // Unique code found, break the loop
		}
	}

	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, err
	}

	cic, err := i.entClient.ClassInvitationCode.Create().
		SetClassID(classID).
		SetInvitationCode(invitationCode).
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
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	cic, err := i.entClient.ClassInvitationCode.Query().Where(entcic.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetClassInvitationCodeByIDResponse{
		ClassInvitationCode: utils.ToPbClassInvitationCode(cic),
	}, nil
}

func (i *Interactor) GetClassInvitationCodeListByClassID(ctx context.Context, req *pb.GetClassInvitationCodeListByClassIDRequest) (*pb.GetClassInvitationCodeListByClassIDResponse, error) {
	classID, err := uuid.Parse(req.ClassId)
	if err != nil {
		return nil, err
	}

	cicList, err := i.entClient.ClassInvitationCode.Query().Where(entcic.ClassID(classID)).All(ctx)
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
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	cic, err := i.entClient.ClassInvitationCode.UpdateOneID(id).
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandomString generates a random string of n length.
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	sb := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(sb)
}

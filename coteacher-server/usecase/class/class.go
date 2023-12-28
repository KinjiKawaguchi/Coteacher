package usecase

import (
	"coteacher/domain/repository/ent"
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

func (i *Interactor) CreateClass(ctx context.Context, req *pb.CreateClassRequest) (*pb.CreateClassResponse, error) {
	// クラスの作成と保存
	class, err := i.entClient.Class.Create().
		SetName(req.Name).
		AddUserIDs(req.TeacherId).
		Save(ctx)

	// エラーが発生した場合、エラーを返す
	if err != nil {
		i.logger.Error("failed to create class", slog.Any("error", err))
		return nil, err // エラーを返す際には、通常はエラー情報をラップする
	}

	// 成功した場合、作成されたクラスの情報を含むレスポンスを返す
	return &pb.CreateClassResponse{
		Class: &pb.Class{
			Id:        class.ID,
			Name:      class.Name,
			TeacherId: class.Edges.Users[0].ID,
		},
	}, nil
}

package utils

import (
	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbStudentClass(t *ent.StudentClass) *pb.StudentClass {
	return &pb.StudentClass{
		StudentId: t.StudentID.String(),
		ClassId:   t.ClassID.String(),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func ToPbUser(t *ent.User) *pb.User {
	return &pb.User{
		Id:        t.ID.String(),
		Name:      t.Name,
		Email:     t.Email,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func ToPbClassInvitationCode(t *ent.ClassInvitationCode) *pb.ClassInvitationCode {
	return &pb.ClassInvitationCode{
		Id:             t.ID.String(),
		ClassId:        t.ClassID.String(),
		InvitationCode: t.InvitationCode,
		ExpirationDate: timestamppb.New(t.ExpirationDate),
		IsActive:       t.IsActive,
		CreatedAt:      timestamppb.New(t.CreatedAt),
		UpdatedAt:      timestamppb.New(t.UpdatedAt),
	}
}

func ToPbClass(t *ent.Class) *pb.Class {
	return &pb.Class{
		Id:        t.ID.String(),
		Name:      t.Name,
		TeacherId: t.TeacherID.String(),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

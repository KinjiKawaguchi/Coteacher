package utils

import (
	"coteacher/domain/repository/ent"
	pb "coteacher/proto-gen/go/coteacher/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbStudentClass(t *ent.StudentClass) *pb.StudentClass {
	return &pb.StudentClass{
		StudentId: t.StudentID,
		ClassId:   t.ClassID,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func ToPbUser(t *ent.User) *pb.User {
	return &pb.User{
		Id:        t.ID,
		Name:      t.Name,
		Email:     t.Email,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func ToPbClassInvitationCode(t *ent.ClassInvitationCode) *pb.ClassInvitationCode {
	return &pb.ClassInvitationCode{
		Id:             t.ID,
		ClassId:        t.ClassID,
		InvitationCode: t.InvitationCode,
		ExpirationDate: timestamppb.New(t.ExpirationDate),
		IsActive:       t.IsActive,
		CreatedAt:      timestamppb.New(t.CreatedAt),
		UpdatedAt:      timestamppb.New(t.UpdatedAt),
	}
}

func ToPbClass(t *ent.Class) *pb.Class {
	return &pb.Class{
		Id:        t.ID,
		Name:      t.Name,
		TeacherId: t.TeacherID,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

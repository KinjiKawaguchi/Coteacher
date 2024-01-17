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

func ToPbForm(t *ent.Form) *pb.Form {
	return &pb.Form{
		Id:          t.ID.String(),
		ClassId:     t.ClassID.String(),
		Name:        t.Name,
		Description: t.Description,
		UsageLimit:  int32(t.UsageLimit),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
	}
}

func ToPbQuestion(t *ent.Question) *pb.Question {
	// Convert QuestionType to its corresponding enum in pb.QuestionType
	questionType := pb.Question_QuestionType(pb.Question_QuestionType_value[t.QuestionType.String()])

	// Create a slice of *pb.Question_QuestionOption for options
	var pbOptions []*pb.QuestionOption
	for _, option := range t.Edges.QuestionOption {
		pbOptions = append(pbOptions, &pb.QuestionOption{
			Id:         option.ID.String(),
			QuestionId: t.ID.String(),
			OptionText: option.OptionText,
			Order:      int32(option.Order),
		})
	}

	// Create the TextQuestion if applicable
	var textQuestion *pb.Question_TextQuestion
	if len(t.Edges.TextQuestion) > 0 {
		firstTextQuestion := t.Edges.TextQuestion[0]
		textQuestion = &pb.Question_TextQuestion{
			Id:         firstTextQuestion.ID.String(),
			QuestionId: t.ID.String(),
			MaxLength:  int32(firstTextQuestion.MaxLength),
		}
	}

	// Construct and return the pb.Question
	return &pb.Question{
		Id:              t.ID.String(),
		FormId:          t.FormID.String(),
		QuestionType:    questionType,
		QuestionText:    t.QuestionText,
		IsRequired:      t.IsRequired,
		ForAiProcessing: t.ForAiProcessing,
		Order:           int32(t.Order),
		CreatedAt:       timestamppb.New(t.CreatedAt),
		UpdatedAt:       timestamppb.New(t.UpdatedAt),
		Options:         pbOptions,
		TextQuestion:    textQuestion,
	}
}

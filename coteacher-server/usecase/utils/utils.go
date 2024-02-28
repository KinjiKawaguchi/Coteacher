package utils

import (
	pb "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"

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
		Id:           t.ID.String(),
		ClassId:      t.ClassID.String(),
		Name:         t.Name,
		Description:  t.Description,
		UsageLimit:   int32(t.UsageLimit),
		SystemPrompt: t.SystemPrompt,
		CreatedAt:    timestamppb.New(t.CreatedAt),
		UpdatedAt:    timestamppb.New(t.UpdatedAt),
	}
}

func ToPbQuestion(t *ent.Question) *pb.Question {
	// Create a slice of *pb.QuestionOption for options
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

	// Convert QuestionType to its corresponding enum in pb.QuestionType
	questionType := convertQuestionTypeToPbQuestionType(t.QuestionType)

	// Construct and return the pb.Question
	return &pb.Question{
		Id:              t.ID.String(),
		FormId:          t.FormID.String(),
		QuestionType:    questionType,
		QuestionText:    t.QuestionText,
		TextQuestion:    textQuestion,
		Options:         pbOptions,
		IsRequired:      t.IsRequired,
		ForAiProcessing: t.ForAiProcessing,
		Order:           int32(t.Order),
		CreatedAt:       timestamppb.New(t.CreatedAt),
		UpdatedAt:       timestamppb.New(t.UpdatedAt),
	}
}

func ToPbResponse(t *ent.Response) *pb.Response {
	return &pb.Response{
		Id:         t.ID.String(),
		FormId:     t.FormID.String(),
		StudentId:  t.StudentID.String(),
		Answers:    ToPbAnswers(t.Edges.Answer),
		AiResponse: t.AiResponse,
		CreatedAt:  timestamppb.New(t.CreatedAt),
		UpdatedAt:  timestamppb.New(t.UpdatedAt),
	}
}

func ToPbAnswers(t []*ent.Answer) []*pb.Response_Answer {
	var pbAnswers []*pb.Response_Answer
	for _, answer := range t {
		pbAnswers = append(pbAnswers, &pb.Response_Answer{
			Id:         answer.ID.String(),
			ResponseId: answer.ResponseID.String(),
			QuestionId: answer.QuestionID.String(),
			Order:      int32(answer.Order),
			AnswerText: answer.AnswerText,
		})
	}
	return pbAnswers
}

func convertQuestionTypeToPbQuestionType(questionType question.QuestionType) pb.Question_QuestionType {
	switch questionType {
	case question.QuestionTypeText:
		return pb.Question_QUESTION_TYPE_TEXT
	case question.QuestionTypeParagraphText:
		return pb.Question_QUESTION_TYPE_PARAGRAPH_TEXT
	case question.QuestionTypeCheckbox:
		return pb.Question_QUESTION_TYPE_CHECKBOX
	case question.QuestionTypeRadio:
		return pb.Question_QUESTION_TYPE_RADIO
	case question.QuestionTypeList:
		return pb.Question_QUESTION_TYPE_LIST
	case question.QuestionTypeMultipleChoice:
		return pb.Question_QUESTION_TYPE_MULTIPLE_CHOICE
	default:
		return pb.Question_QUESTION_TYPE_UNSPECIFIED
	}
}

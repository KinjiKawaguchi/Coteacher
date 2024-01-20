package form

import (
	"context"
	"log"

	"connectrpc.com/connect"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	ent_question "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	question "github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/questionoption"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/textquestion"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/utils"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) GetQuestionListByFormID(ctx context.Context, req *pb.GetQuestionListByFormIdRequest) (*pb.GetQuestionListByFormIdResponse, error) {
	// Parse form ID.
	formID, err := uuid.Parse(req.FormId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// Get all questions by form id.
	questions, err := i.entClient.Question.Query().
		Where(question.FormID(formID)).
		WithTextQuestion().
		WithQuestionOption().
		All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Process questions.
	var pbQuestions []*pb.Question
	for _, q := range questions {
		pbQuestions = append(pbQuestions, utils.ToPbQuestion(q))
	}

	return &pb.GetQuestionListByFormIdResponse{
		Questions: pbQuestions,
	}, nil
}

func (i *Interactor) SaveQuestionList(ctx context.Context, req *pb.SaveQuestionListRequest) (*pb.SaveQuestionListResponse, error) {
	log.Println("SaveQuestionList started") // 関数開始時のログ

	var pbQuestions []*pb.Question

	for idx, question := range req.Questions {
		log.Printf("Processing question %d: %+v\n", idx, question) // 各質問の処理開始ログ

		if len(req.IsDeleted) > idx && req.IsDeleted[idx] {
			log.Printf("Deleting question %d: %+v\n", idx, question)

			if question.Id != "" {
				questionID, err := uuid.Parse(question.Id)
				if err != nil {
					return nil, connect.NewError(connect.CodeInvalidArgument, err)
				}

				// Delete the question's related records
				switch question.QuestionType {
				case pb.Question_QUESTION_TYPE_TEXT, pb.Question_QUESTION_TYPE_PARAGRAPH_TEXT:
					_, err = i.entClient.TextQuestion.Delete().
						Where(textquestion.QuestionID(questionID)).Exec(ctx)
					if err != nil {
						return nil, connect.NewError(connect.CodeInternal, err)
					}

				case pb.Question_QUESTION_TYPE_CHECKBOX, pb.Question_QUESTION_TYPE_RADIO, pb.Question_QUESTION_TYPE_LIST, pb.Question_QUESTION_TYPE_MULTIPLE_CHOICE:
					_, err = i.entClient.QuestionOption.Delete().
						Where(questionoption.QuestionID(questionID)).Exec(ctx)
					if err != nil {
						return nil, connect.NewError(connect.CodeInternal, err)
					}
				}

				// Delete the question
				err = i.entClient.Question.DeleteOneID(questionID).Exec(ctx)
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, err)
				}
			}
			continue // Skip further processing for this question
		}

		var dbQuestion *ent.Question
		var err error

		// Determine if we're updating an existing question or creating a new one
		if question.Id != "" {
			log.Printf("Updating question %d: %+v\n", idx, question)
			questionID, err := uuid.Parse(question.Id)
			if err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, err)
			}

			FormID, err := uuid.Parse(question.FormId)
			if err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, err)
			}

			// Update existing question
			updateQuery := i.entClient.Question.UpdateOneID(questionID).
				SetQuestionType(convertPbQuestionTypeToEntQuestionType(question.QuestionType)).
				SetQuestionText(question.QuestionText).
				SetIsRequired(question.IsRequired).
				SetForAiProcessing(question.ForAiProcessing).
				SetOrder(int(question.Order)).
				SetFormID(FormID)

			// Additional fields updates can go here

			dbQuestion, err = updateQuery.Save(ctx)
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, err)
			}
		} else {
			log.Printf("Creating question %d: %+v\n", idx, question)
			// Create new question
			formID, err := uuid.Parse(question.FormId)
			createQuery := i.entClient.Question.
				Create().
				SetQuestionType(convertPbQuestionTypeToEntQuestionType(question.QuestionType)).
				SetQuestionText(question.QuestionText).
				SetIsRequired(question.IsRequired).
				SetForAiProcessing(question.ForAiProcessing).
				SetOrder(int(question.Order)).
				SetFormID(formID)

			dbQuestion, err = createQuery.Save(ctx)
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, err)
			}

			log.Printf("Created question %d: %+v\n", idx, question)
		}

		// Handle additional details based on QuestionType
		switch question.QuestionType {
		case pb.Question_QUESTION_TYPE_TEXT, pb.Question_QUESTION_TYPE_PARAGRAPH_TEXT:
			if question.TextQuestion != nil {
				if question.TextQuestion.Id != "" {
					textQuestionID, err := uuid.Parse(question.TextQuestion.Id)
					if err != nil {
						return nil, connect.NewError(connect.CodeInvalidArgument, err)
					}

					// Update existing TextQuestion
					_, err = i.entClient.TextQuestion.UpdateOneID(textQuestionID).
						SetMaxLength(int(question.TextQuestion.MaxLength)).
						Save(ctx)
					if err != nil {
						return nil, connect.NewError(connect.CodeInternal, err)
					}
				} else {
					log.Printf("Creating TextQuestion for question %d: %+v\n", idx, question)
					// Create new TextQuestion
					_, err = i.entClient.TextQuestion.
						Create().
						SetQuestionID(dbQuestion.ID).
						SetMaxLength(int(question.TextQuestion.MaxLength)).
						Save(ctx)
					if err != nil {
						return nil, connect.NewError(connect.CodeInternal, err)
					}
					log.Printf("Created TextQuestion for question %d: %+v\n", idx, question)
				}
			}

		case pb.Question_QUESTION_TYPE_CHECKBOX, pb.Question_QUESTION_TYPE_RADIO, pb.Question_QUESTION_TYPE_LIST, pb.Question_QUESTION_TYPE_MULTIPLE_CHOICE:
			// Assuming question.Options is a slice of QuestionOption
			for _, option := range question.Options {
				if option.Id != "" {
					optionID, err := uuid.Parse(option.Id)
					if err != nil {
						return nil, connect.NewError(connect.CodeInvalidArgument, err)
					}

					// Update existing QuestionOption
					_, err = i.entClient.QuestionOption.UpdateOneID(optionID).
						SetOptionText(option.OptionText).
						SetOrder(int(option.Order)).
						Save(ctx)
					if err != nil {
						return nil, connect.NewError(connect.CodeInternal, err)
					}
				} else {
					// Create new QuestionOption
					_, err = i.entClient.QuestionOption.
						Create().
						SetQuestionID(dbQuestion.ID).
						SetOptionText(option.OptionText).
						SetOrder(int(option.Order)).
						Save(ctx)
					if err != nil {
						return nil, connect.NewError(connect.CodeInternal, err)
					}
				}
			}
		}

		log.Printf("Retrieving question %d: %+v\n", idx, question)
		dbQuestion, err = i.entClient.Question.Query().
			Where(ent_question.ID(dbQuestion.ID)).
			WithTextQuestion().
			WithQuestionOption().
			Only(ctx)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}

		// Convert the updated dbQuestion to protobuf Question
		pbQuestion := utils.ToPbQuestion(dbQuestion)
		pbQuestions = append(pbQuestions, pbQuestion)

		log.Println("Question processed successfully") // 質問の処理が成功したことを示すログ
	}

	return &pb.SaveQuestionListResponse{
		Questions: pbQuestions,
	}, nil
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

func convertPbQuestionTypeToEntQuestionType(pbType pb.Question_QuestionType) question.QuestionType {
	switch pbType {
	case pb.Question_QUESTION_TYPE_TEXT:
		return question.QuestionTypeText
	case pb.Question_QUESTION_TYPE_PARAGRAPH_TEXT:
		return question.QuestionTypeParagraphText
	case pb.Question_QUESTION_TYPE_CHECKBOX:
		return question.QuestionTypeCheckbox
	case pb.Question_QUESTION_TYPE_RADIO:
		return question.QuestionTypeRadio
	case pb.Question_QUESTION_TYPE_LIST:
		return question.QuestionTypeList
	case pb.Question_QUESTION_TYPE_MULTIPLE_CHOICE:
		return question.QuestionTypeMultipleChoice
	default:
		return question.QuestionTypeUnspecified
	}
}

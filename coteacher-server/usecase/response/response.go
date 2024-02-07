package response

import (
	"context"

	"connectrpc.com/connect"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
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

func (i *Interactor) GetNumberOfResponsesByStudentID(ctx context.Context, req *pb.GetNumberOfResponsesByStudentIDRequest) (*pb.GetNumberOfResponsesByStudentIDResponse, error) {
	// Get all forms by class id.
	studentID, err := uuid.Parse(req.StudentId)
	responses, err := i.entClient.Response.Query().Where(response.StudentID(studentID)).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 個数を返す
	return &pb.GetNumberOfResponsesByStudentIDResponse{
		NumberOfResponses: int32(len(responses)),
	}, nil
}

func (i *Interactor) GetNumberOfResponsesByFormID(ctx context.Context, req *pb.GetNumberOfResponsesByFormIDRequest) (*pb.GetNumberOfResponsesByFormIDResponse, error) {
	// Get all forms by class id.
	formID, err := uuid.Parse(req.FormId)
	responses, err := i.entClient.Response.Query().Where(response.FormID(formID)).All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 個数を返す
	return &pb.GetNumberOfResponsesByFormIDResponse{
		NumberOfResponses: int32(len(responses)),
	}, nil
}

func (i *Interactor) SubmitResponse(ctx context.Context, req *pb.SubmitResponseRequest) (*pb.SubmitResponseResponse, error) {
	// Get all forms by class id.
	formID, err := uuid.Parse(req.FormId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := i.entClient.Response.Create().
		SetFormID(formID).
		SetStudentID(studentID).
		SetAiResponse(req.AiResponse).
		Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	for _, answer := range req.Answers {
		questionID, err := uuid.Parse(answer.QuestionId)
		_, err = i.entClient.Answer.Create().
			SetResponseID(response.ID).
			SetAnswerText(answer.AnswerText).
			SetQuestionID(questionID).
			SetOrder(int(answer.Order)).
			Save(ctx)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}
	return &pb.SubmitResponseResponse{
		ResponseId: response.ID.String(),
		Success:    true,
	}, nil
}

func (i *Interactor) GetResponseListByFormID(ctx context.Context, req *pb.GetResponseListByFormIDRequest) (*pb.GetResponseListByFormIDResponse, error) {
	// Parse form ID from request
	formID, err := uuid.Parse(req.FormId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Query responses with associated entities
	responses, err := i.entClient.Response.
		Query().
		Where(response.FormID(formID)).
		WithStudent(func(q *ent.StudentQuery) {
			q.WithUser()
		}).
		WithAnswer(func(q *ent.AnswerQuery) {
			q.WithQuestion()
		}).
		All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	studentMap := make(map[uuid.UUID]*pb.User)
	questionMap := make(map[uuid.UUID]*pb.Question)
	pbResponses := make([]*pb.Response, 0, len(responses))

	// Process each response and populate maps and pbResponses slice
	for _, resp := range responses {
		// Convert and add the response to the pbResponses slice
		pbResponses = append(pbResponses, utils.ToPbResponse(resp))

		// Process the student and add to studentMap if not already present
		student := resp.Edges.Student.Edges.User
		if _, exists := studentMap[student.ID]; !exists {
			studentMap[student.ID] = utils.ToPbUser(student)
		}

		// Process each answer and add its question to questionMap if not already present
		for _, answer := range resp.Edges.Answer {
			question := answer.Edges.Question
			if _, exists := questionMap[question.ID]; !exists {
				questionMap[question.ID] = utils.ToPbQuestion(question)
			}
		}
	}

	// Convert maps to slices for the response
	pbStudents := make([]*pb.User, 0, len(studentMap))
	for _, student := range studentMap {
		pbStudents = append(pbStudents, student)
	}

	pbQuestions := make([]*pb.Question, 0, len(questionMap))
	for _, question := range questionMap {
		pbQuestions = append(pbQuestions, question)
	}

	// Return the assembled response
	return &pb.GetResponseListByFormIDResponse{
		Responses: pbResponses,
		Questions: pbQuestions,
		Students:  pbStudents,
	}, nil
}

func (i *Interactor) SubmitAIResponse(ctx context.Context, req *pb.SubmitAIResponseRequest) (*pb.SubmitAIResponseResponse, error) {
	// Get all forms by class id.

	responseID, err := uuid.Parse(req.ResponseId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	_, err = i.entClient.Response.UpdateOneID(responseID).
		SetAiResponse(req.AiResponse).
		Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &pb.SubmitAIResponseResponse{
		Success: true,
	}, nil
}

package response

import (
	"context"
	"sort"

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

func (i *Interactor) CreateDataset(ctx context.Context, req *pb.CreateDatasetRequest) (*pb.CreateDatasetResponse, error) {
	system_prompt := "あなたは、多くの生徒を持つ先生が拾いきれない疑問、質問、意見などに対して代わりに生徒と対話を行います。\n" +
		"先生が用意したAIに提供する情報の生徒からの入力は以下のフォーマットによって与えられます。\n" +
		"{\n" +
		"instructions:'先生からの指示'\n" +
		" able_conversation:一問一答かどうか \n" +
		" questions{ \n" +
		" [question_type:AIの入力のために先生が用意した質問のタイプ\n" +
		"question_text:AIの入力のために先生が用意した質問\n" +
		"answer:その質問に対する生徒の回答\n" +
		"]}\n" +
		"}\n" +
		"あなたの最も重要視すべきことは、生徒が理解し、定着し、応用できるような学習を促すことと、生徒の学習意欲や自主性を高めることです。生徒の理解度や意欲を常に把握し、それに応じた対応をしましょう。生徒が自分で考え、答えを導き出すことができるよう、ヒントやサポートを与えましょう。生徒が学ぶ楽しさや喜びを感じられるよう、コミュニケーションを大切にしましょう。とにかく元気にポジティブな返答をしましょう。\n"
	responses, err := i.entClient.Response.Query().
		WithAnswer(
			func(q *ent.AnswerQuery) {
				q.WithQuestion()
			},
		).
		WithForm().
		WithStudent(
			func(q *ent.StudentQuery) {
				q.WithUser()
			},
		).
		All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var reepatedMessages []*pb.Messages
	for _, response := range responses {
		user_prompt := "{\n" +
			"instructions:" + "'" + response.Edges.Form.SystemPrompt + "',\n" +
			"able_conversation: false,\n" +
			"questions: [\n"
		answers := response.Edges.Answer
		// answerのorderでソートして繰り返す
		sort.Slice(answers, func(i, j int) bool {
			return answers[i].Order < answers[j].Order
		})
		for _, answer := range answers {
			question := answer.Edges.Question
			user_prompt += "{\n" +
				"question_type:" + "'5',\n" +
				"question_text:" + "'" + question.QuestionText + "',\n" +
				"answer:" + "'" + answer.AnswerText + "'\n" +
				"},\n"
		}
		user_prompt += "]\n" +
			"}"
		// answer_prompt := response.Edges.Student.Edges.User.Name + "さん、先生の補助をする元気でポジティブなGPT先生だよ!よろしく!" +
		// 	"ただし解決のための直接的な回答は出来ないから、これから教えるヒントを活かしてがんばって!"

		var messages *pb.Messages
		system_message := &pb.Message{
			Role:    "system",
			Content: system_prompt,
		}

		user_message := &pb.Message{
			Role:    "user",
			Content: user_prompt,
		}

		assistant_message := &pb.Message{
			Role:    "assistant",
			Content: response.AiResponse,
		}

		messages = &pb.Messages{
			Messages: []*pb.Message{
				system_message,
				user_message,
				assistant_message,
			},
		}

		reepatedMessages = append(reepatedMessages, messages)
	}
	return &pb.CreateDatasetResponse{
		RepeatedMessages: reepatedMessages,
	}, nil
}

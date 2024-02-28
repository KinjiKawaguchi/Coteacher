import { Response_Answer } from '@/gen/proto/coteacher/v1/resources_pb';
import {
  Answer as INTERFACE_Answer,
  Response as INTERFACE_Response,
  Question,
  User,
} from '@/interfaces';
import { UserType } from '@/gen/proto/coteacher/v1/user_pb';
import { GetResponseListByFormIDResponse } from '@/gen/proto/coteacher/v1/response_pb';

export function transformAnswerToResponseAnswer(
  a: INTERFACE_Answer | undefined
): Response_Answer {
  const answer = new Response_Answer();
  answer.answerText = a?.answerText || '';
  if (a?.question?.id) answer.questionId = a.question.id;
  if (a?.order) answer.order = a.order;
  if (a?.selectedOptionList) {
    answer.answerOptionIds = a.selectedOptionList
      .filter(o => o !== undefined)
      .map(o => o.id || '');
  }
  return answer;
}

export function transformGrpcResponseToResponse(
  res: GetResponseListByFormIDResponse
): INTERFACE_Response[] {
  return res.responses.flatMap(r => {
    const matchedStudent = res.students.find(s => s.id === r.studentId);
    if (!matchedStudent) return [];

    const student: User = {
      id: matchedStudent.id,
      name: matchedStudent.name,
      email: matchedStudent.email,
      user_type: UserType.STUDENT,
    };

    return [
      {
        id: r.id,
        formId: r.formId,
        student: student,
        aiResponse: r.aiResponse,
        answerList: r.answers.flatMap(a => {
          const matchedQuestion = res.questions?.find(
            q => q.id === a.questionId
          );
          if (!matchedQuestion) return [];

          const question: Question = {
            id: matchedQuestion.id,
            formId: matchedQuestion.formId,
            questionType: matchedQuestion.questionType,
            questionText: matchedQuestion.questionText,
            isRequired: matchedQuestion.isRequired,
            forAiProcessing: matchedQuestion.forAiProcessing,
            order: matchedQuestion.order,
            options: matchedQuestion.options,
          };

          return [
            {
              id: a.id,
              responseId: a.responseId,
              question: question,
              order: a.order,
              answerText: a.answerText,
              selectedOptionList: a.answerOptionIds.flatMap(optionId => {
                const matchedOption = question.options?.find(
                  option => option.id === optionId
                );
                if (!matchedOption) return [];
                return [
                  {
                    id: matchedOption.id,
                    questionId: matchedOption.questionId,
                    optionText: matchedOption.optionText,
                    order: matchedOption.order,
                  },
                ];
              }),
            },
          ];
        }),
      },
    ];
  });
}

import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import { ResponseService } from '@/gen/proto/coteacher/v1/response_connect';
import {
  GetNumberOfResponsesByFormIDRequest,
  GetNumberOfResponsesByStudentIDRequest,
  GetResponseListByFormIDRequest,
  // GetResponseListByFormIDRequest,
  SubmitResponseRequest,
} from '@/gen/proto/coteacher/v1/response_pb';
import {
  Answer as INTERFACE_Answer,
  Response as INTERFACE_Response,
  Question,
  User,
} from '@/interfaces';
import { Response_Answer } from '@/gen/proto/coteacher/v1/resources_pb';
import { UserType } from '@/gen/proto/coteacher/v1/user_pb';

class ResponseRepository {
  readonly cli: PromiseClient<typeof ResponseService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(ResponseService, t);
  }

  async getNumberOfResponseByFormId(formId: string): Promise<number> {
    const req = new GetNumberOfResponsesByFormIDRequest();
    req.formId = formId;
    const res = await this.cli.getNumberOfResponsesByFormID(req);
    return res.numberOfResponses;
  }

  async getNumberOfResponseByStudentId(
    formId: string,
    studentId?: string
  ): Promise<number> {
    const req = new GetNumberOfResponsesByStudentIDRequest();
    req.formId = formId;
    if (studentId) {
      req.studentId = studentId;
    } else {
      req.studentId = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.getNumberOfResponsesByFormID(req);
    return res.numberOfResponses;
  }

  async submitResponse(
    formId: string,
    answerList: INTERFACE_Answer[],
    aiResponse: string,
    studentId?: string
  ): Promise<boolean> {
    const req = new SubmitResponseRequest();
    req.formId = formId;
    if (studentId) {
      req.studentId = studentId;
    } else {
      req.studentId = localStorage.getItem('UserID') || '';
    }
    req.answers = answerList.map(a => {
      const answer = new Response_Answer();
      answer.answerText = a.answerText ?? '';
      if (a.question.id) answer.questionId = a.question.id;
      if (a.selectedOptionList) {
        answer.answerOptionIds = a.selectedOptionList
          .filter(o => o !== undefined)
          .map(o => o.id || '');
      }
      return answer;
    });
    req.aiResponse = aiResponse;
    const res = await this.cli.submitResponse(req);
    return res.success;
  }

  async getResponseListByFormId(formId: string): Promise<INTERFACE_Response[]> {
    const req = new GetResponseListByFormIDRequest();
    req.formId = formId;
    const res = await this.cli.getResponseListByFormID(req);

    const responseList = res.responses.flatMap(r => {
      const matchedStudent = res.students.find(s => s.id === r.studentId);
      if (!matchedStudent) return [];

      const student: User = {
        id: matchedStudent.id,
        name: matchedStudent.name,
        email: matchedStudent.email,
        user_type: UserType.STUDENT,
      };

      const response: INTERFACE_Response = {
        id: r.id,
        formId: r.formId,
        student: student,
        aiResponse: r.aiResponse,
        answerList: r.answers.flatMap(a => {
          if (!res.questions) return [];
          const matchedQuestion = res.questions.find(
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

          const answer: INTERFACE_Answer = {
            id: a.id,
            responseId: a.responseId,
            question: question,
            answerText: a.answerText,
            selectedOptionList: a.answerOptionIds.flatMap(optionId => {
              if (!question.options) return [];
              const matchedOption = question.options.find(
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
          };
          return [answer];
        }),
      };
      return [response];
    });

    return responseList;
  }

  // async getResponseListByFormId(
  //   formId: string,
  // ): Promise<INTERFACE_Response[]> {
  //   const req = new GetResponseListByFormIDRequest();
  //   req.formId = formId;
  //   const response = await this.cli.getResponseListByFormID(req);
  //   const responseList = res.responses.map(r => {
  //     const response: Response = {
  //       id: r.id,
  //       formId: r.formId,
  //       studentId: r.studentId,
  //       answers: r.answers.map(a => {
  //         const answer: INTERFACE_Answer = {
  //           answerText: a.answerText,

  //           selectedOptionList: a.answerOptionIdsList.map(o => {
  //             const option: INTERFACE_Option = {
  //               id: o
  //               tionText: '',
  //             };
  //             return option;
  //           }),
  //         };
  //         return answer;
  //       }),
  //       aiResponse: r.aiResponse,
  //     };
  //     return response;
  //   }
  // }
}

export const responseRepo = new ResponseRepository(backendGrpcTransport);

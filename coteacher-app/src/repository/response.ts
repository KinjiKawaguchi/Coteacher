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
  // GetResponseListByFormIDRequest,
  SubmitResponseRequest,
} from '@/gen/proto/coteacher/v1/response_pb';
import {
  Answer as INTERFACE_Answer,
  // Response as INTERFACE_Response,
  // Option as INTERFACE_Option,
} from '@/interfaces';
import { Response_Answer } from '@/gen/proto/coteacher/v1/resources_pb';

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
      answer.answerText = a.answerText;
      if (a.question.id) answer.questionId = a.question.id;
      if (a.selectedOptionList) {
        answer.answerOptionIds = a.selectedOptionList
          // まず、未定義の値をフィルタリングして、定義済みのオブジェクトのみを処理します。
          .filter(o => o !== undefined)
          // 次に、フィルタリングされた配列に対してマップを実行します。各要素が定義済みであることが確実なので、安全に 'id' にアクセスできます。
          // 'id' が未定義の場合は、デフォルト値（空の文字列など）を使用します。
          .map(o => o.id || '');
      }
      return answer;
    });
    req.aiResponse = aiResponse;
    const res = await this.cli.submitResponse(req);
    return res.success;
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

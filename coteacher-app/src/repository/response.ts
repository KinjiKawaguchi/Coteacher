import { createPromiseClient, Transport } from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import { ResponseService } from '@/gen/proto/coteacher/v1/response_connect';
import {
  GetNumberOfResponsesByFormIDRequest,
  GetNumberOfResponsesByStudentIDRequest,
  GetResponseListByFormIDRequest,
  SubmitAIResponseRequest,
  SubmitResponseRequest,
  SubmitResponseResponse,
} from '@/gen/proto/coteacher/v1/response_pb';
import {
  Response as INTERFACE_Response,
  Answer as INTERFACE_Answer,
} from '@/interfaces';
import {
  transformAnswerToResponseAnswer,
  transformGrpcResponseToResponse,
} from './utils';

class ResponseRepository {
  readonly cli;

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
    req.studentId = studentId || localStorage.getItem('UserID') || '';
    const res = await this.cli.getNumberOfResponsesByFormID(req);
    return res.numberOfResponses;
  }

  async submitResponse(
    formId: string,
    answerList: INTERFACE_Answer[],
    studentId?: string
  ): Promise<SubmitResponseResponse> {
    const req = new SubmitResponseRequest();
    req.formId = formId;
    req.studentId = studentId || localStorage.getItem('UserID') || '';
    req.answers = answerList.map(transformAnswerToResponseAnswer);
    const res = await this.cli.submitResponse(req);
    return res;
  }

  async submitAiResponse(
    responseId: string,
    aiResponse: string
  ): Promise<boolean> {
    const req = new SubmitAIResponseRequest();
    req.responseId = responseId;
    req.aiResponse = aiResponse;
    const res = await this.cli.submitAIResponse(req);
    return res.success;
  }

  async getResponseListByFormId(formId: string): Promise<INTERFACE_Response[]> {
    const req = new GetResponseListByFormIDRequest();
    req.formId = formId;
    const res = await this.cli.getResponseListByFormID(req);
    return transformGrpcResponseToResponse(res);
  }
}

export const responseRepo = new ResponseRepository(backendGrpcTransport);

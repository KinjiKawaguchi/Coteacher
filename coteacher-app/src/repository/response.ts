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
} from '@/gen/proto/coteacher/v1/response_pb';

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
    studentId?: string,
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
}

export const responseRepo = new ResponseRepository(backendGrpcTransport);

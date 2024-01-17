import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import { FormService } from '@/gen/proto/coteacher/v1/form_connect';
import { Form } from '@/interfaces';
import { GetFormListByClassIDRequest } from '@/gen/proto/coteacher/v1/form_pb';

class FormRepository {
  readonly cli: PromiseClient<typeof FormService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(FormService, t);
  }

  async getFormListByClassId(classId: string): Promise<Form[]> {
    const req = new GetFormListByClassIDRequest();
    req.classId = classId;
    const res = await this.cli.getFormListByClassID(req);
    const forms: Form[] = res.forms.map(f => {
      return {
        id: f.id,
        classId: f.classId,
        name: f.name,
        description: f.description,
        usageLimit: f.usageLimit,
      };
    });
    return forms;
  }
}

export const formRepo = new FormRepository(backendGrpcTransport);

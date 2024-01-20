import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import { FormService } from '@/gen/proto/coteacher/v1/form_connect';
import { Form } from '@/interfaces';
import {
  CheckFormEditPermissionRequest,
  CheckFormViewPermissionRequest,
  CreateFormRequest,
  GetFormListByClassIDRequest,
} from '@/gen/proto/coteacher/v1/form_pb';

export interface CreateFormInput {
  classId: string;
  name: string;
  description: string;
  usageLimit: number;
}
class FormRepository {
  readonly cli: PromiseClient<typeof FormService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(FormService, t);
  }

  async createForm(createFormInput: CreateFormInput) {
    const req = new CreateFormRequest();
    req.classId = createFormInput.classId;
    req.name = createFormInput.name;
    req.description = createFormInput.description;
    req.usageLimit = createFormInput.usageLimit;
    const res = await this.cli.createForm(req);
    if (!res.form) {
      throw new Error('form is null');
    }
    const form: Form = {
      id: res.form.id,
      classId: res.form.classId,
      name: res.form.name,
      description: res.form.description,
      usageLimit: res.form.usageLimit,
    };
    return form;
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

  async checkFormEditPermission(formId: string, teacherId: string | null) {
    const req = new CheckFormEditPermissionRequest();
    req.formId = formId;
    if (teacherId) {
      req.teacherId = teacherId;
    } else {
      req.teacherId = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.checkFormEditPermission(req);
    return res.hasPermission;
  }

  async checkFormViewPermission(formId: string, studentId: string | null) {
    const req = new CheckFormViewPermissionRequest();
    req.formId = formId;
    if (studentId) {
      req.studentId = studentId;
    } else {
      req.studentId = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.checkFormViewPermission(req);
    return res.hasPermission;
  }
}

export const formRepo = new FormRepository(backendGrpcTransport);

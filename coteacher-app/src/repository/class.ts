import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import { ClassService } from '@/gen/proto/coteacher/v1/class_connect';
import {
  CreateClassRequest,
  GetClassListByTeacherIDRequest,
  DeleteClassRequest,
} from '@/gen/proto/coteacher/v1/class_pb';
import { Class } from '@/interfaces';

class ClassRepository {
  readonly cli: PromiseClient<typeof ClassService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(ClassService, t);
  }

  async createClass(className: string): Promise<Class> {
    const req = new CreateClassRequest();
    req.name = className;
    req.teacherId = localStorage.getItem('UserID') || '';
    const res = await this.cli.createClass(req);
    if (!res.class) {
      throw new Error('Class not created');
    }
    const c: Class = {
      id: res.class.id,
      name: res.class.name,
      teacherId: res.class.teacherId,
    };
    return c;
  }

  async deleteClass(classId: string): Promise<void> {
    const req = new DeleteClassRequest();
    req.id = classId;
    await this.cli.deleteClass(req);
  }

  async getClassListByTeacherId(teacherId: string | null): Promise<Class[]> {
    const req = new GetClassListByTeacherIDRequest();
    if (teacherId) {
      req.teacherId = teacherId;
    } else {
      req.teacherId = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.getClassListByTeacherID(req);
    const classes: Class[] = res.classes.map(c => {
      return {
        id: c.id,
        name: c.name,
        teacherId: c.teacherId,
      };
    });
    return classes;
  }
}

export const classRepo = new ClassRepository(backendGrpcTransport);

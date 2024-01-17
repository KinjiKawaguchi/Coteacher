import { backendGrpcTransport } from '@/config/connectRpc';
import { TeacherService } from '@/gen/proto/coteacher/v1/teacher_connect';
import {
  CheckTeacherExistsByEmailRequest,
  CheckTeacherExistsByIDRequest,
} from '@/gen/proto/coteacher/v1/teacher_pb';
import {
  PromiseClient,
  Transport,
  createPromiseClient,
} from '@bufbuild/connect';

class TeacherRepository {
  readonly cli: PromiseClient<typeof TeacherService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(TeacherService, t);
  }

  async checkTeacherExistsByID(teacherId: string | null): Promise<boolean> {
    const req = new CheckTeacherExistsByIDRequest();
    if (teacherId) {
      req.id = teacherId;
    } else {
      req.id = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.checkTeacherExistsByID(req);
    return res.exists;
  }

  async checkTeacherExistsByEmail(email: string | null): Promise<boolean> {
    const req = new CheckTeacherExistsByEmailRequest();
    if (email) {
      req.email = email;
    } else {
      req.email = localStorage.getItem('UserEmail') || '';
    }
    const res = await this.cli.checkTeacherExistsByEmail(req);
    return res.exists;
  }
}

export const teacherRepo = new TeacherRepository(backendGrpcTransport);

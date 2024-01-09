import { backendGrpcTransport } from '@/config/connectRpc';
import { StudentService } from '@/gen/proto/coteacher/v1/student_connect';
import {
  CheckStudentExistsByEmailRequest,
  CheckStudentExistsByIDRequest,
  ParticipateClassRequest,
  QuitClassRequest,
} from '@/gen/proto/coteacher/v1/student_pb';
import { Class } from '@/interfaces';
import {
  PromiseClient,
  Transport,
  createPromiseClient,
} from '@bufbuild/connect';

class StudentRepository {
  readonly cli: PromiseClient<typeof StudentService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(StudentService, t);
  }

  async checkStudentExistsByID(studentId: string | null): Promise<boolean> {
    const req = new CheckStudentExistsByIDRequest();
    if (studentId) {
      req.id = studentId;
    } else {
      req.id = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.checkStudentExistsByID(req);
    return res.exists;
  }

  async checkStudentExistsByEmail(email: string | null): Promise<boolean> {
    const req = new CheckStudentExistsByEmailRequest();
    if (email) {
      req.email = email;
    } else {
      req.email = localStorage.getItem('UserEmail') || '';
    }
    const res = await this.cli.checkStudentExistsByEmail(req);
    return res.exists;
  }

  async participateClass(classInvitationCode: string): Promise<Class> {
    const req = new ParticipateClassRequest();
    req.invitaionCode = classInvitationCode;
    req.userId = localStorage.getItem('UserID') || '';
    const res = await this.cli.participateClass(req);
    if (!res.class) {
      throw new Error('Class not found');
    }
    const c: Class = {
      id: res.class.id,
      name: res.class.name,
      teacherId: res.class.teacherId,
    };
    return c;
  }

  async quitClass(classId: string): Promise<void> {
    const req = new QuitClassRequest();
    req.classId = classId;
    req.userId = localStorage.getItem('UserID') || '';
    await this.cli.quitClass(req);

    return;
  }
}

export const studentRepo = new StudentRepository(backendGrpcTransport);

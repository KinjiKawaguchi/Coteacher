import { backendGrpcTransport } from '@/config/connectRpc';
import { Class } from '@/interfaces';
import { StudentClassService } from '@/gen/proto/coteacher/v1/student_class_connect';
import { GetClassListByStudentIDRequest } from '@/gen/proto/coteacher/v1/student_class_pb';
import {
  PromiseClient,
  Transport,
  createPromiseClient,
} from '@bufbuild/connect';

class StudentClassRepository {
  readonly cli: PromiseClient<typeof StudentClassService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(StudentClassService, t);
  }

  async getClassListByStudentId(studentId: string | null): Promise<Class[]> {
    const req = new GetClassListByStudentIDRequest();
    if (studentId) {
      req.studentId = studentId;
    } else {
      req.studentId = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.getClassListByStudentID(req);
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

export const studentClassRepo = new StudentClassRepository(
  backendGrpcTransport
);

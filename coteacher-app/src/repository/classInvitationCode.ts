import { backendGrpcTransport } from '@/config/connectRpc';
import { ClassInvitationCodeService } from '@/gen/proto/coteacher/v1/class_invitation_code_connect';
import { CreateClassInvitationCodeRequest } from '@/gen/proto/coteacher/v1/class_invitation_code_pb';
import { ClassInvitationCode } from '@/interfaces';
import {
  PromiseClient,
  Transport,
  createPromiseClient,
} from '@bufbuild/connect';
import { Timestamp } from '@bufbuild/protobuf';

class ClassInvitationCodeRepository {
  readonly cli: PromiseClient<typeof ClassInvitationCodeService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(ClassInvitationCodeService, t);
  }

  async issueClassInvitationCode(
    classId: string
  ): Promise<ClassInvitationCode> {
    const req = new CreateClassInvitationCodeRequest();
    req.classId = classId;
    // 一年先にexpirationDateを設定
    req.expirationDate = Timestamp.fromDate(
      new Date(new Date().setFullYear(new Date().getFullYear() + 1))
    );
    const res = await this.cli.createClassInvitationCode(req);
    if (!res.classInvitationCode) {
      throw new Error('classInvitationCode is null');
    }
    const cic: ClassInvitationCode = {
      id: res.classInvitationCode.id,
      classId: res.classInvitationCode.classId,
      invitationCode: res.classInvitationCode.invitationCode,
      expirationDate:
        res.classInvitationCode.expirationDate?.toDate() || new Date(), // TODO: 修正
      isActive: res.classInvitationCode.isActive,
    };
    return cic;
  }
}

export const classInvitationCodeRepo = new ClassInvitationCodeRepository(
  backendGrpcTransport
);

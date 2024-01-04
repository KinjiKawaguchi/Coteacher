import {
  CheckUserExistsByEmailRequest,
  CheckUserExistsByEmailResponse,
  GetUserByEmailRequest,
  GetUserByEmailResponse,
} from '@/gen/proto/coteacher/v1/user_pb';
import { UserService } from '@/gen/proto/coteacher/v1/user_connect';
import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';

class UserRepository {
  readonly cli: PromiseClient<typeof UserService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(UserService, t);
  }

  async getUserByEmail(
    req: GetUserByEmailRequest
  ): Promise<GetUserByEmailResponse> {
    const res = await this.cli.getUserByEmail(req);
    return res;
  }

  async checkUserExistsByEmail(
    req: CheckUserExistsByEmailRequest
  ): Promise<CheckUserExistsByEmailResponse> {
    const res = await this.cli.checkUserExistsByEmail(req);
    console.log(res);
    return res;
  }
}

export const userRepo = new UserRepository(backendGrpcTransport);

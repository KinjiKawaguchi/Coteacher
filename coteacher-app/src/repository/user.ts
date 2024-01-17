import {
  CheckUserExistsByEmailRequest,
  CreateUserRequest,
  GetUserByEmailRequest,
  GetUserByIDRequest,
} from '@/gen/proto/coteacher/v1/user_pb';
import { UserService } from '@/gen/proto/coteacher/v1/user_connect';
import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import { User } from '@/interfaces';

class UserRepository {
  readonly cli: PromiseClient<typeof UserService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(UserService, t);
  }

  async createUser(name: string, userType: number): Promise<User> {
    const req = new CreateUserRequest();
    req.name = name;
    req.email = localStorage.getItem('email') || '';
    req.userType = userType;
    const res = await this.cli.createUser(req);
    if (!res.user) {
      throw new Error('User not created');
    }

    localStorage.setItem('UserID', res.user.id);
    localStorage.setItem('UserEmail', res.user.email);
    localStorage.setItem('UserName', res.user.name);
    localStorage.setItem('UserType', res.userType.toString());
    const user: User = {
      id: res.user.id,
      name: res.user.name,
      email: res.user.email,
      user_type: res.userType, //TODO: 戻値の確認
    };
    return user;
  }

  async getUserById(id: string | null): Promise<User> {
    const req = new GetUserByIDRequest();
    if (id) {
      req.id = id;
    } else {
      req.id = localStorage.getItem('UserID') || '';
    }
    const res = await this.cli.getUserByID(req);
    if (!res.user) {
      throw new Error('User not found');
    }
    const user: User = {
      id: res.user.id,
      name: res.user.name,
      email: res.user.email,
      user_type: res.userType, //TODO: 戻値の確認
    };
    return user;
  }

  async getUserByEmail(email: string | null): Promise<User> {
    const req = new GetUserByEmailRequest();
    if (email) {
      req.email = email;
    } else {
      req.email = localStorage.getItem('UserEmail') || '';
    }
    const res = await this.cli.getUserByEmail(req);
    if (!res.user) {
      throw new Error('User not found');
    }
    // TODO: ここはまじで良くないので治す。
    localStorage.setItem('UserID', res.user.id);
    localStorage.setItem('UserEmail', res.user.email);
    localStorage.setItem('UserName', res.user.name);
    localStorage.setItem('UserType', res.userType.toString());
    const user: User = {
      id: res.user.id,
      name: res.user.name,
      email: res.user.email,
      user_type: res.userType, //TODO: 戻値の確認
    };
    return user;
  }

  async checkUserExistsByEmail(email: string | null): Promise<boolean> {
    const req = new CheckUserExistsByEmailRequest();
    if (email) {
      req.email = email;
    } else {
      req.email = localStorage.getItem('UserEmail') || '';
    }
    const res = await this.cli.checkUserExistsByEmail(req);
    return res.exists;
  }
}

export const userRepo = new UserRepository(backendGrpcTransport);

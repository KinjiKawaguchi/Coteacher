// GENERATED CODE -- DO NOT EDIT!

// package: coteacher.v1
// file: coteacher/v1/user.proto

import * as coteacher_v1_user_pb from "../../coteacher/v1/user_pb";
import * as grpc from "@grpc/grpc-js";

interface IUserServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createUser: grpc.MethodDefinition<coteacher_v1_user_pb.CreateUserRequest, coteacher_v1_user_pb.CreateUserResponse>;
  getUserByID: grpc.MethodDefinition<coteacher_v1_user_pb.GetUserByIDRequest, coteacher_v1_user_pb.GetUserByIDResponse>;
  getUserByEmail: grpc.MethodDefinition<coteacher_v1_user_pb.GetUserByEmailRequest, coteacher_v1_user_pb.GetUserByEmailResponse>;
  updateUser: grpc.MethodDefinition<coteacher_v1_user_pb.UpdateUserRequest, coteacher_v1_user_pb.UpdateUserResponse>;
  deleteUser: grpc.MethodDefinition<coteacher_v1_user_pb.DeleteUserRequest, coteacher_v1_user_pb.DeleteUserResponse>;
}

export const UserServiceService: IUserServiceService;

export interface IUserServiceServer extends grpc.UntypedServiceImplementation {
  createUser: grpc.handleUnaryCall<coteacher_v1_user_pb.CreateUserRequest, coteacher_v1_user_pb.CreateUserResponse>;
  getUserByID: grpc.handleUnaryCall<coteacher_v1_user_pb.GetUserByIDRequest, coteacher_v1_user_pb.GetUserByIDResponse>;
  getUserByEmail: grpc.handleUnaryCall<coteacher_v1_user_pb.GetUserByEmailRequest, coteacher_v1_user_pb.GetUserByEmailResponse>;
  updateUser: grpc.handleUnaryCall<coteacher_v1_user_pb.UpdateUserRequest, coteacher_v1_user_pb.UpdateUserResponse>;
  deleteUser: grpc.handleUnaryCall<coteacher_v1_user_pb.DeleteUserRequest, coteacher_v1_user_pb.DeleteUserResponse>;
}

export class UserServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createUser(argument: coteacher_v1_user_pb.CreateUserRequest, callback: grpc.requestCallback<coteacher_v1_user_pb.CreateUserResponse>): grpc.ClientUnaryCall;
  createUser(argument: coteacher_v1_user_pb.CreateUserRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.CreateUserResponse>): grpc.ClientUnaryCall;
  createUser(argument: coteacher_v1_user_pb.CreateUserRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.CreateUserResponse>): grpc.ClientUnaryCall;
  getUserByID(argument: coteacher_v1_user_pb.GetUserByIDRequest, callback: grpc.requestCallback<coteacher_v1_user_pb.GetUserByIDResponse>): grpc.ClientUnaryCall;
  getUserByID(argument: coteacher_v1_user_pb.GetUserByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.GetUserByIDResponse>): grpc.ClientUnaryCall;
  getUserByID(argument: coteacher_v1_user_pb.GetUserByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.GetUserByIDResponse>): grpc.ClientUnaryCall;
  getUserByEmail(argument: coteacher_v1_user_pb.GetUserByEmailRequest, callback: grpc.requestCallback<coteacher_v1_user_pb.GetUserByEmailResponse>): grpc.ClientUnaryCall;
  getUserByEmail(argument: coteacher_v1_user_pb.GetUserByEmailRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.GetUserByEmailResponse>): grpc.ClientUnaryCall;
  getUserByEmail(argument: coteacher_v1_user_pb.GetUserByEmailRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.GetUserByEmailResponse>): grpc.ClientUnaryCall;
  updateUser(argument: coteacher_v1_user_pb.UpdateUserRequest, callback: grpc.requestCallback<coteacher_v1_user_pb.UpdateUserResponse>): grpc.ClientUnaryCall;
  updateUser(argument: coteacher_v1_user_pb.UpdateUserRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.UpdateUserResponse>): grpc.ClientUnaryCall;
  updateUser(argument: coteacher_v1_user_pb.UpdateUserRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.UpdateUserResponse>): grpc.ClientUnaryCall;
  deleteUser(argument: coteacher_v1_user_pb.DeleteUserRequest, callback: grpc.requestCallback<coteacher_v1_user_pb.DeleteUserResponse>): grpc.ClientUnaryCall;
  deleteUser(argument: coteacher_v1_user_pb.DeleteUserRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.DeleteUserResponse>): grpc.ClientUnaryCall;
  deleteUser(argument: coteacher_v1_user_pb.DeleteUserRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_user_pb.DeleteUserResponse>): grpc.ClientUnaryCall;
}

// package: coteacher.v1
// file: coteacher/v1/user.proto

import * as jspb from "google-protobuf";
import * as coteacher_v1_resources_pb from "../../coteacher/v1/resources_pb";

export class CreateUserRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    name: string,
    email: string,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export class CreateUserResponse extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): coteacher_v1_resources_pb.User | undefined;
  setUser(value?: coteacher_v1_resources_pb.User): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
  static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
  export type AsObject = {
    user?: coteacher_v1_resources_pb.User.AsObject,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export class GetUserByIDRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByIDRequest): GetUserByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByIDRequest;
  static deserializeBinaryFromReader(message: GetUserByIDRequest, reader: jspb.BinaryReader): GetUserByIDRequest;
}

export namespace GetUserByIDRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetUserByIDResponse extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): coteacher_v1_resources_pb.User | undefined;
  setUser(value?: coteacher_v1_resources_pb.User): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByIDResponse): GetUserByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByIDResponse;
  static deserializeBinaryFromReader(message: GetUserByIDResponse, reader: jspb.BinaryReader): GetUserByIDResponse;
}

export namespace GetUserByIDResponse {
  export type AsObject = {
    user?: coteacher_v1_resources_pb.User.AsObject,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export class GetUserByEmailRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByEmailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByEmailRequest): GetUserByEmailRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserByEmailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByEmailRequest;
  static deserializeBinaryFromReader(message: GetUserByEmailRequest, reader: jspb.BinaryReader): GetUserByEmailRequest;
}

export namespace GetUserByEmailRequest {
  export type AsObject = {
    email: string,
  }
}

export class GetUserByEmailResponse extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): coteacher_v1_resources_pb.User | undefined;
  setUser(value?: coteacher_v1_resources_pb.User): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByEmailResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByEmailResponse): GetUserByEmailResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserByEmailResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByEmailResponse;
  static deserializeBinaryFromReader(message: GetUserByEmailResponse, reader: jspb.BinaryReader): GetUserByEmailResponse;
}

export namespace GetUserByEmailResponse {
  export type AsObject = {
    user?: coteacher_v1_resources_pb.User.AsObject,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export class UpdateUserRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
  static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    id: string,
    name: string,
    email: string,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export class UpdateUserResponse extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): coteacher_v1_resources_pb.User | undefined;
  setUser(value?: coteacher_v1_resources_pb.User): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserResponse): UpdateUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserResponse;
  static deserializeBinaryFromReader(message: UpdateUserResponse, reader: jspb.BinaryReader): UpdateUserResponse;
}

export namespace UpdateUserResponse {
  export type AsObject = {
    user?: coteacher_v1_resources_pb.User.AsObject,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export class DeleteUserRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
  static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteUserResponse extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): coteacher_v1_resources_pb.User | undefined;
  setUser(value?: coteacher_v1_resources_pb.User): void;

  getUserType(): UserTypeMap[keyof UserTypeMap];
  setUserType(value: UserTypeMap[keyof UserTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserResponse): DeleteUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserResponse;
  static deserializeBinaryFromReader(message: DeleteUserResponse, reader: jspb.BinaryReader): DeleteUserResponse;
}

export namespace DeleteUserResponse {
  export type AsObject = {
    user?: coteacher_v1_resources_pb.User.AsObject,
    userType: UserTypeMap[keyof UserTypeMap],
  }
}

export interface UserTypeMap {
  USER_TYPE_UNSPECIFIED: 0;
  USER_TYPE_STUDENT: 1;
  USER_TYPE_TEACHER: 2;
}

export const UserType: UserTypeMap;


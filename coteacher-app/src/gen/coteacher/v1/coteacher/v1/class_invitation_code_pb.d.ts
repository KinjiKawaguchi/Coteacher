// package: coteacher.v1
// file: coteacher/v1/class_invitation_code.proto

import * as jspb from "google-protobuf";
import * as coteacher_v1_resources_pb from "../../coteacher/v1/resources_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class CreateClassInvitationCodeRequest extends jspb.Message {
  getClassId(): string;
  setClassId(value: string): void;

  hasExpirationDate(): boolean;
  clearExpirationDate(): void;
  getExpirationDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExpirationDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateClassInvitationCodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateClassInvitationCodeRequest): CreateClassInvitationCodeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateClassInvitationCodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateClassInvitationCodeRequest;
  static deserializeBinaryFromReader(message: CreateClassInvitationCodeRequest, reader: jspb.BinaryReader): CreateClassInvitationCodeRequest;
}

export namespace CreateClassInvitationCodeRequest {
  export type AsObject = {
    classId: string,
    expirationDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateClassInvitationCodeResponse extends jspb.Message {
  hasClassInvitationCode(): boolean;
  clearClassInvitationCode(): void;
  getClassInvitationCode(): coteacher_v1_resources_pb.ClassInvitationCode | undefined;
  setClassInvitationCode(value?: coteacher_v1_resources_pb.ClassInvitationCode): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateClassInvitationCodeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateClassInvitationCodeResponse): CreateClassInvitationCodeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateClassInvitationCodeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateClassInvitationCodeResponse;
  static deserializeBinaryFromReader(message: CreateClassInvitationCodeResponse, reader: jspb.BinaryReader): CreateClassInvitationCodeResponse;
}

export namespace CreateClassInvitationCodeResponse {
  export type AsObject = {
    classInvitationCode?: coteacher_v1_resources_pb.ClassInvitationCode.AsObject,
  }
}

export class GetClassInvitationCodeByIDRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassInvitationCodeByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassInvitationCodeByIDRequest): GetClassInvitationCodeByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassInvitationCodeByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassInvitationCodeByIDRequest;
  static deserializeBinaryFromReader(message: GetClassInvitationCodeByIDRequest, reader: jspb.BinaryReader): GetClassInvitationCodeByIDRequest;
}

export namespace GetClassInvitationCodeByIDRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetClassInvitationCodeByIDResponse extends jspb.Message {
  hasClassInvitationCode(): boolean;
  clearClassInvitationCode(): void;
  getClassInvitationCode(): coteacher_v1_resources_pb.ClassInvitationCode | undefined;
  setClassInvitationCode(value?: coteacher_v1_resources_pb.ClassInvitationCode): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassInvitationCodeByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassInvitationCodeByIDResponse): GetClassInvitationCodeByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassInvitationCodeByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassInvitationCodeByIDResponse;
  static deserializeBinaryFromReader(message: GetClassInvitationCodeByIDResponse, reader: jspb.BinaryReader): GetClassInvitationCodeByIDResponse;
}

export namespace GetClassInvitationCodeByIDResponse {
  export type AsObject = {
    classInvitationCode?: coteacher_v1_resources_pb.ClassInvitationCode.AsObject,
  }
}

export class GetClassInvitationCodeListByClassIDRequest extends jspb.Message {
  getClassId(): string;
  setClassId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassInvitationCodeListByClassIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassInvitationCodeListByClassIDRequest): GetClassInvitationCodeListByClassIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassInvitationCodeListByClassIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassInvitationCodeListByClassIDRequest;
  static deserializeBinaryFromReader(message: GetClassInvitationCodeListByClassIDRequest, reader: jspb.BinaryReader): GetClassInvitationCodeListByClassIDRequest;
}

export namespace GetClassInvitationCodeListByClassIDRequest {
  export type AsObject = {
    classId: string,
  }
}

export class GetClassInvitationCodeListByClassIDResponse extends jspb.Message {
  clearClassInvitationCodeListList(): void;
  getClassInvitationCodeListList(): Array<coteacher_v1_resources_pb.ClassInvitationCode>;
  setClassInvitationCodeListList(value: Array<coteacher_v1_resources_pb.ClassInvitationCode>): void;
  addClassInvitationCodeList(value?: coteacher_v1_resources_pb.ClassInvitationCode, index?: number): coteacher_v1_resources_pb.ClassInvitationCode;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassInvitationCodeListByClassIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassInvitationCodeListByClassIDResponse): GetClassInvitationCodeListByClassIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassInvitationCodeListByClassIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassInvitationCodeListByClassIDResponse;
  static deserializeBinaryFromReader(message: GetClassInvitationCodeListByClassIDResponse, reader: jspb.BinaryReader): GetClassInvitationCodeListByClassIDResponse;
}

export namespace GetClassInvitationCodeListByClassIDResponse {
  export type AsObject = {
    classInvitationCodeListList: Array<coteacher_v1_resources_pb.ClassInvitationCode.AsObject>,
  }
}

export class UpdateClassInvitationCodeRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasExpirationDate(): boolean;
  clearExpirationDate(): void;
  getExpirationDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExpirationDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getIsActive(): boolean;
  setIsActive(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateClassInvitationCodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateClassInvitationCodeRequest): UpdateClassInvitationCodeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateClassInvitationCodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateClassInvitationCodeRequest;
  static deserializeBinaryFromReader(message: UpdateClassInvitationCodeRequest, reader: jspb.BinaryReader): UpdateClassInvitationCodeRequest;
}

export namespace UpdateClassInvitationCodeRequest {
  export type AsObject = {
    id: string,
    expirationDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    isActive: boolean,
  }
}

export class UpdateClassInvitationCodeResponse extends jspb.Message {
  hasClassInvitationCode(): boolean;
  clearClassInvitationCode(): void;
  getClassInvitationCode(): coteacher_v1_resources_pb.ClassInvitationCode | undefined;
  setClassInvitationCode(value?: coteacher_v1_resources_pb.ClassInvitationCode): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateClassInvitationCodeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateClassInvitationCodeResponse): UpdateClassInvitationCodeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateClassInvitationCodeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateClassInvitationCodeResponse;
  static deserializeBinaryFromReader(message: UpdateClassInvitationCodeResponse, reader: jspb.BinaryReader): UpdateClassInvitationCodeResponse;
}

export namespace UpdateClassInvitationCodeResponse {
  export type AsObject = {
    classInvitationCode?: coteacher_v1_resources_pb.ClassInvitationCode.AsObject,
  }
}


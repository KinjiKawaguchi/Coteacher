// package: coteacher.v1
// file: coteacher/v1/student_class.proto

import * as jspb from "google-protobuf";
import * as coteacher_v1_resources_pb from "../../coteacher/v1/resources_pb";

export class CreateStudentClassRequest extends jspb.Message {
  getStudentId(): string;
  setStudentId(value: string): void;

  getClassId(): string;
  setClassId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateStudentClassRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateStudentClassRequest): CreateStudentClassRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateStudentClassRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateStudentClassRequest;
  static deserializeBinaryFromReader(message: CreateStudentClassRequest, reader: jspb.BinaryReader): CreateStudentClassRequest;
}

export namespace CreateStudentClassRequest {
  export type AsObject = {
    studentId: string,
    classId: string,
  }
}

export class CreateStudentClassResponse extends jspb.Message {
  hasStudentClass(): boolean;
  clearStudentClass(): void;
  getStudentClass(): coteacher_v1_resources_pb.StudentClass | undefined;
  setStudentClass(value?: coteacher_v1_resources_pb.StudentClass): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateStudentClassResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateStudentClassResponse): CreateStudentClassResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateStudentClassResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateStudentClassResponse;
  static deserializeBinaryFromReader(message: CreateStudentClassResponse, reader: jspb.BinaryReader): CreateStudentClassResponse;
}

export namespace CreateStudentClassResponse {
  export type AsObject = {
    studentClass?: coteacher_v1_resources_pb.StudentClass.AsObject,
  }
}

export class GetStudentListByClassIDRequest extends jspb.Message {
  getClassId(): string;
  setClassId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStudentListByClassIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetStudentListByClassIDRequest): GetStudentListByClassIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetStudentListByClassIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStudentListByClassIDRequest;
  static deserializeBinaryFromReader(message: GetStudentListByClassIDRequest, reader: jspb.BinaryReader): GetStudentListByClassIDRequest;
}

export namespace GetStudentListByClassIDRequest {
  export type AsObject = {
    classId: string,
  }
}

export class GetStudentListByClassIDResponse extends jspb.Message {
  clearStudentsList(): void;
  getStudentsList(): Array<coteacher_v1_resources_pb.User>;
  setStudentsList(value: Array<coteacher_v1_resources_pb.User>): void;
  addStudents(value?: coteacher_v1_resources_pb.User, index?: number): coteacher_v1_resources_pb.User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStudentListByClassIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetStudentListByClassIDResponse): GetStudentListByClassIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetStudentListByClassIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStudentListByClassIDResponse;
  static deserializeBinaryFromReader(message: GetStudentListByClassIDResponse, reader: jspb.BinaryReader): GetStudentListByClassIDResponse;
}

export namespace GetStudentListByClassIDResponse {
  export type AsObject = {
    studentsList: Array<coteacher_v1_resources_pb.User.AsObject>,
  }
}

export class GetClassListByStudentIDRequest extends jspb.Message {
  getStudentId(): string;
  setStudentId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassListByStudentIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassListByStudentIDRequest): GetClassListByStudentIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassListByStudentIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassListByStudentIDRequest;
  static deserializeBinaryFromReader(message: GetClassListByStudentIDRequest, reader: jspb.BinaryReader): GetClassListByStudentIDRequest;
}

export namespace GetClassListByStudentIDRequest {
  export type AsObject = {
    studentId: string,
  }
}

export class GetClassListByStudentIDResponse extends jspb.Message {
  clearClassesList(): void;
  getClassesList(): Array<coteacher_v1_resources_pb.Class>;
  setClassesList(value: Array<coteacher_v1_resources_pb.Class>): void;
  addClasses(value?: coteacher_v1_resources_pb.Class, index?: number): coteacher_v1_resources_pb.Class;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassListByStudentIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassListByStudentIDResponse): GetClassListByStudentIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassListByStudentIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassListByStudentIDResponse;
  static deserializeBinaryFromReader(message: GetClassListByStudentIDResponse, reader: jspb.BinaryReader): GetClassListByStudentIDResponse;
}

export namespace GetClassListByStudentIDResponse {
  export type AsObject = {
    classesList: Array<coteacher_v1_resources_pb.Class.AsObject>,
  }
}

export class DeleteStudentClassRequest extends jspb.Message {
  getStudentId(): string;
  setStudentId(value: string): void;

  getClassId(): string;
  setClassId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteStudentClassRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteStudentClassRequest): DeleteStudentClassRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteStudentClassRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteStudentClassRequest;
  static deserializeBinaryFromReader(message: DeleteStudentClassRequest, reader: jspb.BinaryReader): DeleteStudentClassRequest;
}

export namespace DeleteStudentClassRequest {
  export type AsObject = {
    studentId: string,
    classId: string,
  }
}

export class DeleteStudentClassResponse extends jspb.Message {
  hasStudentClass(): boolean;
  clearStudentClass(): void;
  getStudentClass(): coteacher_v1_resources_pb.StudentClass | undefined;
  setStudentClass(value?: coteacher_v1_resources_pb.StudentClass): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteStudentClassResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteStudentClassResponse): DeleteStudentClassResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteStudentClassResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteStudentClassResponse;
  static deserializeBinaryFromReader(message: DeleteStudentClassResponse, reader: jspb.BinaryReader): DeleteStudentClassResponse;
}

export namespace DeleteStudentClassResponse {
  export type AsObject = {
    studentClass?: coteacher_v1_resources_pb.StudentClass.AsObject,
  }
}


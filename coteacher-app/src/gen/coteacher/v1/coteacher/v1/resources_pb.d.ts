// package: coteacher.v1
// file: coteacher/v1/resources.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  hasTeacher(): boolean;
  clearTeacher(): void;
  getTeacher(): Teacher | undefined;
  setTeacher(value?: Teacher): void;

  hasStudent(): boolean;
  clearStudent(): void;
  getStudent(): Student | undefined;
  setStudent(value?: Student): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getRoleCase(): User.RoleCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    name: string,
    email: string,
    teacher?: Teacher.AsObject,
    student?: Student.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export enum RoleCase {
    ROLE_NOT_SET = 0,
    TEACHER = 4,
    STUDENT = 5,
  }
}

export class Teacher extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Teacher.AsObject;
  static toObject(includeInstance: boolean, msg: Teacher): Teacher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Teacher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Teacher;
  static deserializeBinaryFromReader(message: Teacher, reader: jspb.BinaryReader): Teacher;
}

export namespace Teacher {
  export type AsObject = {
    id: string,
  }
}

export class Student extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Student.AsObject;
  static toObject(includeInstance: boolean, msg: Student): Student.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Student, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Student;
  static deserializeBinaryFromReader(message: Student, reader: jspb.BinaryReader): Student;
}

export namespace Student {
  export type AsObject = {
    id: string,
  }
}

export class ClassInvitationCode extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getClassId(): string;
  setClassId(value: string): void;

  getInvitationCode(): string;
  setInvitationCode(value: string): void;

  hasExpirationDate(): boolean;
  clearExpirationDate(): void;
  getExpirationDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExpirationDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getIsActive(): boolean;
  setIsActive(value: boolean): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClassInvitationCode.AsObject;
  static toObject(includeInstance: boolean, msg: ClassInvitationCode): ClassInvitationCode.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClassInvitationCode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClassInvitationCode;
  static deserializeBinaryFromReader(message: ClassInvitationCode, reader: jspb.BinaryReader): ClassInvitationCode;
}

export namespace ClassInvitationCode {
  export type AsObject = {
    id: string,
    classId: string,
    invitationCode: string,
    expirationDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    isActive: boolean,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Class extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getTeacherId(): string;
  setTeacherId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Class.AsObject;
  static toObject(includeInstance: boolean, msg: Class): Class.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Class, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Class;
  static deserializeBinaryFromReader(message: Class, reader: jspb.BinaryReader): Class;
}

export namespace Class {
  export type AsObject = {
    id: string,
    name: string,
    teacherId: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class StudentClass extends jspb.Message {
  getStudentId(): string;
  setStudentId(value: string): void;

  getClassId(): string;
  setClassId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StudentClass.AsObject;
  static toObject(includeInstance: boolean, msg: StudentClass): StudentClass.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StudentClass, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StudentClass;
  static deserializeBinaryFromReader(message: StudentClass, reader: jspb.BinaryReader): StudentClass;
}

export namespace StudentClass {
  export type AsObject = {
    studentId: string,
    classId: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}


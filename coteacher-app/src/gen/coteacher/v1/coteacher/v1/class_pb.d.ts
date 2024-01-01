// package: coteacher.v1
// file: coteacher/v1/class.proto

import * as jspb from "google-protobuf";
import * as coteacher_v1_resources_pb from "../../coteacher/v1/resources_pb";

export class CreateClassRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getTeacherId(): string;
  setTeacherId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateClassRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateClassRequest): CreateClassRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateClassRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateClassRequest;
  static deserializeBinaryFromReader(message: CreateClassRequest, reader: jspb.BinaryReader): CreateClassRequest;
}

export namespace CreateClassRequest {
  export type AsObject = {
    name: string,
    teacherId: string,
  }
}

export class CreateClassResponse extends jspb.Message {
  hasClass(): boolean;
  clearClass(): void;
  getClass(): coteacher_v1_resources_pb.Class | undefined;
  setClass(value?: coteacher_v1_resources_pb.Class): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateClassResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateClassResponse): CreateClassResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateClassResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateClassResponse;
  static deserializeBinaryFromReader(message: CreateClassResponse, reader: jspb.BinaryReader): CreateClassResponse;
}

export namespace CreateClassResponse {
  export type AsObject = {
    pb_class?: coteacher_v1_resources_pb.Class.AsObject,
  }
}

export class GetClassByIDRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassByIDRequest): GetClassByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassByIDRequest;
  static deserializeBinaryFromReader(message: GetClassByIDRequest, reader: jspb.BinaryReader): GetClassByIDRequest;
}

export namespace GetClassByIDRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetClassByIDResponse extends jspb.Message {
  hasClass(): boolean;
  clearClass(): void;
  getClass(): coteacher_v1_resources_pb.Class | undefined;
  setClass(value?: coteacher_v1_resources_pb.Class): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassByIDResponse): GetClassByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassByIDResponse;
  static deserializeBinaryFromReader(message: GetClassByIDResponse, reader: jspb.BinaryReader): GetClassByIDResponse;
}

export namespace GetClassByIDResponse {
  export type AsObject = {
    pb_class?: coteacher_v1_resources_pb.Class.AsObject,
  }
}

export class GetClassListByTeacherIDRequest extends jspb.Message {
  getTeacherId(): string;
  setTeacherId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassListByTeacherIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassListByTeacherIDRequest): GetClassListByTeacherIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassListByTeacherIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassListByTeacherIDRequest;
  static deserializeBinaryFromReader(message: GetClassListByTeacherIDRequest, reader: jspb.BinaryReader): GetClassListByTeacherIDRequest;
}

export namespace GetClassListByTeacherIDRequest {
  export type AsObject = {
    teacherId: string,
  }
}

export class GetClassListByTeacherIDResponse extends jspb.Message {
  clearClassesList(): void;
  getClassesList(): Array<coteacher_v1_resources_pb.Class>;
  setClassesList(value: Array<coteacher_v1_resources_pb.Class>): void;
  addClasses(value?: coteacher_v1_resources_pb.Class, index?: number): coteacher_v1_resources_pb.Class;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClassListByTeacherIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetClassListByTeacherIDResponse): GetClassListByTeacherIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetClassListByTeacherIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClassListByTeacherIDResponse;
  static deserializeBinaryFromReader(message: GetClassListByTeacherIDResponse, reader: jspb.BinaryReader): GetClassListByTeacherIDResponse;
}

export namespace GetClassListByTeacherIDResponse {
  export type AsObject = {
    classesList: Array<coteacher_v1_resources_pb.Class.AsObject>,
  }
}

export class UpdateClassRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getTeacherId(): string;
  setTeacherId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateClassRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateClassRequest): UpdateClassRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateClassRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateClassRequest;
  static deserializeBinaryFromReader(message: UpdateClassRequest, reader: jspb.BinaryReader): UpdateClassRequest;
}

export namespace UpdateClassRequest {
  export type AsObject = {
    id: string,
    name: string,
    teacherId: string,
  }
}

export class UpdateClassResponse extends jspb.Message {
  hasClass(): boolean;
  clearClass(): void;
  getClass(): coteacher_v1_resources_pb.Class | undefined;
  setClass(value?: coteacher_v1_resources_pb.Class): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateClassResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateClassResponse): UpdateClassResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateClassResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateClassResponse;
  static deserializeBinaryFromReader(message: UpdateClassResponse, reader: jspb.BinaryReader): UpdateClassResponse;
}

export namespace UpdateClassResponse {
  export type AsObject = {
    pb_class?: coteacher_v1_resources_pb.Class.AsObject,
  }
}

export class DeleteClassRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteClassRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteClassRequest): DeleteClassRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteClassRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteClassRequest;
  static deserializeBinaryFromReader(message: DeleteClassRequest, reader: jspb.BinaryReader): DeleteClassRequest;
}

export namespace DeleteClassRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteClassResponse extends jspb.Message {
  hasClass(): boolean;
  clearClass(): void;
  getClass(): coteacher_v1_resources_pb.Class | undefined;
  setClass(value?: coteacher_v1_resources_pb.Class): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteClassResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteClassResponse): DeleteClassResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteClassResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteClassResponse;
  static deserializeBinaryFromReader(message: DeleteClassResponse, reader: jspb.BinaryReader): DeleteClassResponse;
}

export namespace DeleteClassResponse {
  export type AsObject = {
    pb_class?: coteacher_v1_resources_pb.Class.AsObject,
  }
}


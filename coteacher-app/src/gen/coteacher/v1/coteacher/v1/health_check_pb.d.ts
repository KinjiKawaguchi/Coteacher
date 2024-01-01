// package: coteacher.v1
// file: coteacher/v1/health_check.proto

import * as jspb from "google-protobuf";

export class UnaryRequest extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryRequest): UnaryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryRequest;
  static deserializeBinaryFromReader(message: UnaryRequest, reader: jspb.BinaryReader): UnaryRequest;
}

export namespace UnaryRequest {
  export type AsObject = {
    msg: string,
  }
}

export class UnaryResponse extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryResponse): UnaryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryResponse;
  static deserializeBinaryFromReader(message: UnaryResponse, reader: jspb.BinaryReader): UnaryResponse;
}

export namespace UnaryResponse {
  export type AsObject = {
    msg: string,
  }
}

export class ServerStreamingRequest extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerStreamingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ServerStreamingRequest): ServerStreamingRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServerStreamingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerStreamingRequest;
  static deserializeBinaryFromReader(message: ServerStreamingRequest, reader: jspb.BinaryReader): ServerStreamingRequest;
}

export namespace ServerStreamingRequest {
  export type AsObject = {
    msg: string,
  }
}

export class ServerStreamingResponse extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerStreamingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ServerStreamingResponse): ServerStreamingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServerStreamingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerStreamingResponse;
  static deserializeBinaryFromReader(message: ServerStreamingResponse, reader: jspb.BinaryReader): ServerStreamingResponse;
}

export namespace ServerStreamingResponse {
  export type AsObject = {
    msg: string,
  }
}

export class ClientStreamingRequest extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientStreamingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ClientStreamingRequest): ClientStreamingRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientStreamingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientStreamingRequest;
  static deserializeBinaryFromReader(message: ClientStreamingRequest, reader: jspb.BinaryReader): ClientStreamingRequest;
}

export namespace ClientStreamingRequest {
  export type AsObject = {
    msg: string,
  }
}

export class ClientStreamingResponse extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientStreamingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ClientStreamingResponse): ClientStreamingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientStreamingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientStreamingResponse;
  static deserializeBinaryFromReader(message: ClientStreamingResponse, reader: jspb.BinaryReader): ClientStreamingResponse;
}

export namespace ClientStreamingResponse {
  export type AsObject = {
    msg: string,
  }
}

export class BidirectionalStreamingRequest extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BidirectionalStreamingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BidirectionalStreamingRequest): BidirectionalStreamingRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BidirectionalStreamingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BidirectionalStreamingRequest;
  static deserializeBinaryFromReader(message: BidirectionalStreamingRequest, reader: jspb.BinaryReader): BidirectionalStreamingRequest;
}

export namespace BidirectionalStreamingRequest {
  export type AsObject = {
    msg: string,
  }
}

export class BidirectionalStreamingResponse extends jspb.Message {
  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BidirectionalStreamingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BidirectionalStreamingResponse): BidirectionalStreamingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BidirectionalStreamingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BidirectionalStreamingResponse;
  static deserializeBinaryFromReader(message: BidirectionalStreamingResponse, reader: jspb.BinaryReader): BidirectionalStreamingResponse;
}

export namespace BidirectionalStreamingResponse {
  export type AsObject = {
    msg: string,
  }
}


// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file coteacher/v1/response.proto (package coteacher.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Question, Response, Response_Answer, User } from "./resources_pb";

/**
 * @generated from message coteacher.v1.GetNumberOfResponsesByStudentIDRequest
 */
export class GetNumberOfResponsesByStudentIDRequest extends Message<GetNumberOfResponsesByStudentIDRequest> {
  /**
   * @generated from field: string student_id = 1;
   */
  studentId = "";

  /**
   * @generated from field: string form_id = 2;
   */
  formId = "";

  constructor(data?: PartialMessage<GetNumberOfResponsesByStudentIDRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.GetNumberOfResponsesByStudentIDRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "student_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "form_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNumberOfResponsesByStudentIDRequest {
    return new GetNumberOfResponsesByStudentIDRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByStudentIDRequest {
    return new GetNumberOfResponsesByStudentIDRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByStudentIDRequest {
    return new GetNumberOfResponsesByStudentIDRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetNumberOfResponsesByStudentIDRequest | PlainMessage<GetNumberOfResponsesByStudentIDRequest> | undefined, b: GetNumberOfResponsesByStudentIDRequest | PlainMessage<GetNumberOfResponsesByStudentIDRequest> | undefined): boolean {
    return proto3.util.equals(GetNumberOfResponsesByStudentIDRequest, a, b);
  }
}

/**
 * @generated from message coteacher.v1.GetNumberOfResponsesByStudentIDResponse
 */
export class GetNumberOfResponsesByStudentIDResponse extends Message<GetNumberOfResponsesByStudentIDResponse> {
  /**
   * @generated from field: int32 number_of_responses = 1;
   */
  numberOfResponses = 0;

  constructor(data?: PartialMessage<GetNumberOfResponsesByStudentIDResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.GetNumberOfResponsesByStudentIDResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "number_of_responses", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNumberOfResponsesByStudentIDResponse {
    return new GetNumberOfResponsesByStudentIDResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByStudentIDResponse {
    return new GetNumberOfResponsesByStudentIDResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByStudentIDResponse {
    return new GetNumberOfResponsesByStudentIDResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetNumberOfResponsesByStudentIDResponse | PlainMessage<GetNumberOfResponsesByStudentIDResponse> | undefined, b: GetNumberOfResponsesByStudentIDResponse | PlainMessage<GetNumberOfResponsesByStudentIDResponse> | undefined): boolean {
    return proto3.util.equals(GetNumberOfResponsesByStudentIDResponse, a, b);
  }
}

/**
 * @generated from message coteacher.v1.GetNumberOfResponsesByFormIDRequest
 */
export class GetNumberOfResponsesByFormIDRequest extends Message<GetNumberOfResponsesByFormIDRequest> {
  /**
   * @generated from field: string form_id = 1;
   */
  formId = "";

  constructor(data?: PartialMessage<GetNumberOfResponsesByFormIDRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.GetNumberOfResponsesByFormIDRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "form_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNumberOfResponsesByFormIDRequest {
    return new GetNumberOfResponsesByFormIDRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByFormIDRequest {
    return new GetNumberOfResponsesByFormIDRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByFormIDRequest {
    return new GetNumberOfResponsesByFormIDRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetNumberOfResponsesByFormIDRequest | PlainMessage<GetNumberOfResponsesByFormIDRequest> | undefined, b: GetNumberOfResponsesByFormIDRequest | PlainMessage<GetNumberOfResponsesByFormIDRequest> | undefined): boolean {
    return proto3.util.equals(GetNumberOfResponsesByFormIDRequest, a, b);
  }
}

/**
 * @generated from message coteacher.v1.GetNumberOfResponsesByFormIDResponse
 */
export class GetNumberOfResponsesByFormIDResponse extends Message<GetNumberOfResponsesByFormIDResponse> {
  /**
   * @generated from field: int32 number_of_responses = 1;
   */
  numberOfResponses = 0;

  constructor(data?: PartialMessage<GetNumberOfResponsesByFormIDResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.GetNumberOfResponsesByFormIDResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "number_of_responses", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNumberOfResponsesByFormIDResponse {
    return new GetNumberOfResponsesByFormIDResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByFormIDResponse {
    return new GetNumberOfResponsesByFormIDResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNumberOfResponsesByFormIDResponse {
    return new GetNumberOfResponsesByFormIDResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetNumberOfResponsesByFormIDResponse | PlainMessage<GetNumberOfResponsesByFormIDResponse> | undefined, b: GetNumberOfResponsesByFormIDResponse | PlainMessage<GetNumberOfResponsesByFormIDResponse> | undefined): boolean {
    return proto3.util.equals(GetNumberOfResponsesByFormIDResponse, a, b);
  }
}

/**
 * @generated from message coteacher.v1.GetResponseListByFormIDRequest
 */
export class GetResponseListByFormIDRequest extends Message<GetResponseListByFormIDRequest> {
  /**
   * @generated from field: string form_id = 1;
   */
  formId = "";

  constructor(data?: PartialMessage<GetResponseListByFormIDRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.GetResponseListByFormIDRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "form_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetResponseListByFormIDRequest {
    return new GetResponseListByFormIDRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetResponseListByFormIDRequest {
    return new GetResponseListByFormIDRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetResponseListByFormIDRequest {
    return new GetResponseListByFormIDRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetResponseListByFormIDRequest | PlainMessage<GetResponseListByFormIDRequest> | undefined, b: GetResponseListByFormIDRequest | PlainMessage<GetResponseListByFormIDRequest> | undefined): boolean {
    return proto3.util.equals(GetResponseListByFormIDRequest, a, b);
  }
}

/**
 * @generated from message coteacher.v1.GetResponseListByFormIDResponse
 */
export class GetResponseListByFormIDResponse extends Message<GetResponseListByFormIDResponse> {
  /**
   * @generated from field: repeated coteacher.v1.Response responses = 1;
   */
  responses: Response[] = [];

  /**
   * @generated from field: repeated coteacher.v1.Question questions = 2;
   */
  questions: Question[] = [];

  /**
   * @generated from field: repeated coteacher.v1.User students = 3;
   */
  students: User[] = [];

  constructor(data?: PartialMessage<GetResponseListByFormIDResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.GetResponseListByFormIDResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "responses", kind: "message", T: Response, repeated: true },
    { no: 2, name: "questions", kind: "message", T: Question, repeated: true },
    { no: 3, name: "students", kind: "message", T: User, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetResponseListByFormIDResponse {
    return new GetResponseListByFormIDResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetResponseListByFormIDResponse {
    return new GetResponseListByFormIDResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetResponseListByFormIDResponse {
    return new GetResponseListByFormIDResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetResponseListByFormIDResponse | PlainMessage<GetResponseListByFormIDResponse> | undefined, b: GetResponseListByFormIDResponse | PlainMessage<GetResponseListByFormIDResponse> | undefined): boolean {
    return proto3.util.equals(GetResponseListByFormIDResponse, a, b);
  }
}

/**
 * @generated from message coteacher.v1.SubmitResponseRequest
 */
export class SubmitResponseRequest extends Message<SubmitResponseRequest> {
  /**
   * @generated from field: string form_id = 1;
   */
  formId = "";

  /**
   * @generated from field: string student_id = 2;
   */
  studentId = "";

  /**
   * @generated from field: repeated coteacher.v1.Response.Answer answers = 3;
   */
  answers: Response_Answer[] = [];

  /**
   * @generated from field: string ai_response = 4;
   */
  aiResponse = "";

  constructor(data?: PartialMessage<SubmitResponseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.SubmitResponseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "form_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "student_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "answers", kind: "message", T: Response_Answer, repeated: true },
    { no: 4, name: "ai_response", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmitResponseRequest {
    return new SubmitResponseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmitResponseRequest {
    return new SubmitResponseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmitResponseRequest {
    return new SubmitResponseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SubmitResponseRequest | PlainMessage<SubmitResponseRequest> | undefined, b: SubmitResponseRequest | PlainMessage<SubmitResponseRequest> | undefined): boolean {
    return proto3.util.equals(SubmitResponseRequest, a, b);
  }
}

/**
 * @generated from message coteacher.v1.SubmitResponseResponse
 */
export class SubmitResponseResponse extends Message<SubmitResponseResponse> {
  /**
   * @generated from field: bool success = 1;
   */
  success = false;

  /**
   * @generated from field: string response_id = 2;
   */
  responseId = "";

  constructor(data?: PartialMessage<SubmitResponseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.SubmitResponseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "success", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "response_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmitResponseResponse {
    return new SubmitResponseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmitResponseResponse {
    return new SubmitResponseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmitResponseResponse {
    return new SubmitResponseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SubmitResponseResponse | PlainMessage<SubmitResponseResponse> | undefined, b: SubmitResponseResponse | PlainMessage<SubmitResponseResponse> | undefined): boolean {
    return proto3.util.equals(SubmitResponseResponse, a, b);
  }
}

/**
 * @generated from message coteacher.v1.SubmitAIResponseRequest
 */
export class SubmitAIResponseRequest extends Message<SubmitAIResponseRequest> {
  /**
   * @generated from field: string response_id = 1;
   */
  responseId = "";

  /**
   * @generated from field: string ai_response = 2;
   */
  aiResponse = "";

  constructor(data?: PartialMessage<SubmitAIResponseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.SubmitAIResponseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "response_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ai_response", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmitAIResponseRequest {
    return new SubmitAIResponseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmitAIResponseRequest {
    return new SubmitAIResponseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmitAIResponseRequest {
    return new SubmitAIResponseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SubmitAIResponseRequest | PlainMessage<SubmitAIResponseRequest> | undefined, b: SubmitAIResponseRequest | PlainMessage<SubmitAIResponseRequest> | undefined): boolean {
    return proto3.util.equals(SubmitAIResponseRequest, a, b);
  }
}

/**
 * @generated from message coteacher.v1.SubmitAIResponseResponse
 */
export class SubmitAIResponseResponse extends Message<SubmitAIResponseResponse> {
  /**
   * @generated from field: bool success = 1;
   */
  success = false;

  constructor(data?: PartialMessage<SubmitAIResponseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.SubmitAIResponseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "success", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmitAIResponseResponse {
    return new SubmitAIResponseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmitAIResponseResponse {
    return new SubmitAIResponseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmitAIResponseResponse {
    return new SubmitAIResponseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SubmitAIResponseResponse | PlainMessage<SubmitAIResponseResponse> | undefined, b: SubmitAIResponseResponse | PlainMessage<SubmitAIResponseResponse> | undefined): boolean {
    return proto3.util.equals(SubmitAIResponseResponse, a, b);
  }
}

/**
 * @generated from message coteacher.v1.CreateDatasetRequest
 */
export class CreateDatasetRequest extends Message<CreateDatasetRequest> {
  constructor(data?: PartialMessage<CreateDatasetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.CreateDatasetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDatasetRequest {
    return new CreateDatasetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDatasetRequest {
    return new CreateDatasetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDatasetRequest {
    return new CreateDatasetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDatasetRequest | PlainMessage<CreateDatasetRequest> | undefined, b: CreateDatasetRequest | PlainMessage<CreateDatasetRequest> | undefined): boolean {
    return proto3.util.equals(CreateDatasetRequest, a, b);
  }
}

/**
 * @generated from message coteacher.v1.CreateDatasetResponse
 */
export class CreateDatasetResponse extends Message<CreateDatasetResponse> {
  /**
   * @generated from field: repeated string ai_input = 1;
   */
  aiInput: string[] = [];

  /**
   * @generated from field: repeated string ai_output = 2;
   */
  aiOutput: string[] = [];

  constructor(data?: PartialMessage<CreateDatasetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.CreateDatasetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ai_input", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "ai_output", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDatasetResponse {
    return new CreateDatasetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDatasetResponse {
    return new CreateDatasetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDatasetResponse {
    return new CreateDatasetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDatasetResponse | PlainMessage<CreateDatasetResponse> | undefined, b: CreateDatasetResponse | PlainMessage<CreateDatasetResponse> | undefined): boolean {
    return proto3.util.equals(CreateDatasetResponse, a, b);
  }
}


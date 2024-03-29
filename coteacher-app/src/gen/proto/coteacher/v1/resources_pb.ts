// @generated by protoc-gen-es v1.3.1 with parameter "target=ts,import_extension=none"
// @generated from file coteacher/v1/resources.proto (package coteacher.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * Base user message
 *
 * @generated from message coteacher.v1.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string email = 3;
   */
  email = "";

  /**
   * @generated from oneof coteacher.v1.User.role
   */
  role: {
    /**
     * @generated from field: coteacher.v1.Teacher teacher = 4;
     */
    value: Teacher;
    case: "teacher";
  } | {
    /**
     * @generated from field: coteacher.v1.Student student = 5;
     */
    value: Student;
    case: "student";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 6;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<User>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.User";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "teacher", kind: "message", T: Teacher, oneof: "role" },
    { no: 5, name: "student", kind: "message", T: Student, oneof: "role" },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User {
    return new User().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User {
    return new User().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User {
    return new User().fromJsonString(jsonString, options);
  }

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean {
    return proto3.util.equals(User, a, b);
  }
}

/**
 * Teacher message with a reference to the User ID
 *
 * @generated from message coteacher.v1.Teacher
 */
export class Teacher extends Message<Teacher> {
  /**
   * This should match the User ID if it's a 1-to-1 relationship
   *
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<Teacher>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Teacher";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Teacher {
    return new Teacher().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Teacher {
    return new Teacher().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Teacher {
    return new Teacher().fromJsonString(jsonString, options);
  }

  static equals(a: Teacher | PlainMessage<Teacher> | undefined, b: Teacher | PlainMessage<Teacher> | undefined): boolean {
    return proto3.util.equals(Teacher, a, b);
  }
}

/**
 * Student message with a reference to the User ID
 *
 * @generated from message coteacher.v1.Student
 */
export class Student extends Message<Student> {
  /**
   * This should match the User ID if it's a 1-to-1 relationship
   *
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<Student>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Student";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Student {
    return new Student().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Student {
    return new Student().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Student {
    return new Student().fromJsonString(jsonString, options);
  }

  static equals(a: Student | PlainMessage<Student> | undefined, b: Student | PlainMessage<Student> | undefined): boolean {
    return proto3.util.equals(Student, a, b);
  }
}

/**
 * @generated from message coteacher.v1.ClassInvitationCode
 */
export class ClassInvitationCode extends Message<ClassInvitationCode> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string class_id = 2;
   */
  classId = "";

  /**
   * @generated from field: string invitation_code = 3;
   */
  invitationCode = "";

  /**
   * @generated from field: google.protobuf.Timestamp expiration_date = 4;
   */
  expirationDate?: Timestamp;

  /**
   * @generated from field: bool is_active = 5;
   */
  isActive = false;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 6;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<ClassInvitationCode>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.ClassInvitationCode";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "class_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "invitation_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "expiration_date", kind: "message", T: Timestamp },
    { no: 5, name: "is_active", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClassInvitationCode {
    return new ClassInvitationCode().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClassInvitationCode {
    return new ClassInvitationCode().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClassInvitationCode {
    return new ClassInvitationCode().fromJsonString(jsonString, options);
  }

  static equals(a: ClassInvitationCode | PlainMessage<ClassInvitationCode> | undefined, b: ClassInvitationCode | PlainMessage<ClassInvitationCode> | undefined): boolean {
    return proto3.util.equals(ClassInvitationCode, a, b);
  }
}

/**
 * @generated from message coteacher.v1.Class
 */
export class Class extends Message<Class> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string teacher_id = 2;
   */
  teacherId = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 4;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 5;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Class>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Class";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "teacher_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "created_at", kind: "message", T: Timestamp },
    { no: 5, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Class {
    return new Class().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Class {
    return new Class().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Class {
    return new Class().fromJsonString(jsonString, options);
  }

  static equals(a: Class | PlainMessage<Class> | undefined, b: Class | PlainMessage<Class> | undefined): boolean {
    return proto3.util.equals(Class, a, b);
  }
}

/**
 * @generated from message coteacher.v1.StudentClass
 */
export class StudentClass extends Message<StudentClass> {
  /**
   * @generated from field: string student_id = 1;
   */
  studentId = "";

  /**
   * @generated from field: string class_id = 2;
   */
  classId = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 3;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 4;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<StudentClass>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.StudentClass";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "student_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "class_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "message", T: Timestamp },
    { no: 4, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StudentClass {
    return new StudentClass().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StudentClass {
    return new StudentClass().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StudentClass {
    return new StudentClass().fromJsonString(jsonString, options);
  }

  static equals(a: StudentClass | PlainMessage<StudentClass> | undefined, b: StudentClass | PlainMessage<StudentClass> | undefined): boolean {
    return proto3.util.equals(StudentClass, a, b);
  }
}

/**
 * @generated from message coteacher.v1.Form
 */
export class Form extends Message<Form> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string class_id = 2;
   */
  classId = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * @generated from field: int32 usage_limit = 5;
   */
  usageLimit = 0;

  /**
   * @generated from field: string system_prompt = 6;
   */
  systemPrompt = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 7;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 8;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Form>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Form";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "class_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "usage_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "system_prompt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "created_at", kind: "message", T: Timestamp },
    { no: 8, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Form {
    return new Form().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Form {
    return new Form().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Form {
    return new Form().fromJsonString(jsonString, options);
  }

  static equals(a: Form | PlainMessage<Form> | undefined, b: Form | PlainMessage<Form> | undefined): boolean {
    return proto3.util.equals(Form, a, b);
  }
}

/**
 * @generated from message coteacher.v1.Question
 */
export class Question extends Message<Question> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string form_id = 2;
   */
  formId = "";

  /**
   * @generated from field: coteacher.v1.Question.QuestionType question_type = 3;
   */
  questionType = Question_QuestionType.UNSPECIFIED;

  /**
   * @generated from field: string question_text = 4;
   */
  questionText = "";

  /**
   * @generated from field: coteacher.v1.Question.TextQuestion text_question = 5;
   */
  textQuestion?: Question_TextQuestion;

  /**
   * @generated from field: repeated coteacher.v1.QuestionOption options = 6;
   */
  options: QuestionOption[] = [];

  /**
   * @generated from field: bool is_required = 7;
   */
  isRequired = false;

  /**
   * @generated from field: bool for_ai_processing = 8;
   */
  forAiProcessing = false;

  /**
   * @generated from field: int32 order = 9;
   */
  order = 0;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 10;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 11;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Question>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Question";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "form_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "question_type", kind: "enum", T: proto3.getEnumType(Question_QuestionType) },
    { no: 4, name: "question_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "text_question", kind: "message", T: Question_TextQuestion },
    { no: 6, name: "options", kind: "message", T: QuestionOption, repeated: true },
    { no: 7, name: "is_required", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 8, name: "for_ai_processing", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 9, name: "order", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "created_at", kind: "message", T: Timestamp },
    { no: 11, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Question {
    return new Question().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Question {
    return new Question().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Question {
    return new Question().fromJsonString(jsonString, options);
  }

  static equals(a: Question | PlainMessage<Question> | undefined, b: Question | PlainMessage<Question> | undefined): boolean {
    return proto3.util.equals(Question, a, b);
  }
}

/**
 * @generated from enum coteacher.v1.Question.QuestionType
 */
export enum Question_QuestionType {
  /**
   * @generated from enum value: QUESTION_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: QUESTION_TYPE_CHECKBOX = 1;
   */
  CHECKBOX = 1,

  /**
   * @generated from enum value: QUESTION_TYPE_LIST = 2;
   */
  LIST = 2,

  /**
   * @generated from enum value: QUESTION_TYPE_RADIO = 3;
   */
  RADIO = 3,

  /**
   * @generated from enum value: QUESTION_TYPE_MULTIPLE_CHOICE = 4;
   */
  MULTIPLE_CHOICE = 4,

  /**
   * @generated from enum value: QUESTION_TYPE_PARAGRAPH_TEXT = 5;
   */
  PARAGRAPH_TEXT = 5,

  /**
   * @generated from enum value: QUESTION_TYPE_TEXT = 6;
   */
  TEXT = 6,
}
// Retrieve enum metadata with: proto3.getEnumType(Question_QuestionType)
proto3.util.setEnumType(Question_QuestionType, "coteacher.v1.Question.QuestionType", [
  { no: 0, name: "QUESTION_TYPE_UNSPECIFIED" },
  { no: 1, name: "QUESTION_TYPE_CHECKBOX" },
  { no: 2, name: "QUESTION_TYPE_LIST" },
  { no: 3, name: "QUESTION_TYPE_RADIO" },
  { no: 4, name: "QUESTION_TYPE_MULTIPLE_CHOICE" },
  { no: 5, name: "QUESTION_TYPE_PARAGRAPH_TEXT" },
  { no: 6, name: "QUESTION_TYPE_TEXT" },
]);

/**
 * @generated from message coteacher.v1.Question.TextQuestion
 */
export class Question_TextQuestion extends Message<Question_TextQuestion> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string question_id = 2;
   */
  questionId = "";

  /**
   * @generated from field: int32 max_length = 3;
   */
  maxLength = 0;

  constructor(data?: PartialMessage<Question_TextQuestion>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Question.TextQuestion";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "question_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "max_length", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Question_TextQuestion {
    return new Question_TextQuestion().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Question_TextQuestion {
    return new Question_TextQuestion().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Question_TextQuestion {
    return new Question_TextQuestion().fromJsonString(jsonString, options);
  }

  static equals(a: Question_TextQuestion | PlainMessage<Question_TextQuestion> | undefined, b: Question_TextQuestion | PlainMessage<Question_TextQuestion> | undefined): boolean {
    return proto3.util.equals(Question_TextQuestion, a, b);
  }
}

/**
 * @generated from message coteacher.v1.QuestionOption
 */
export class QuestionOption extends Message<QuestionOption> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string question_id = 2;
   */
  questionId = "";

  /**
   * @generated from field: string option_text = 3;
   */
  optionText = "";

  /**
   * @generated from field: int32 order = 4;
   */
  order = 0;

  constructor(data?: PartialMessage<QuestionOption>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.QuestionOption";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "question_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "option_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "order", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QuestionOption {
    return new QuestionOption().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QuestionOption {
    return new QuestionOption().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QuestionOption {
    return new QuestionOption().fromJsonString(jsonString, options);
  }

  static equals(a: QuestionOption | PlainMessage<QuestionOption> | undefined, b: QuestionOption | PlainMessage<QuestionOption> | undefined): boolean {
    return proto3.util.equals(QuestionOption, a, b);
  }
}

/**
 * @generated from message coteacher.v1.Response
 */
export class Response extends Message<Response> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string form_id = 2;
   */
  formId = "";

  /**
   * @generated from field: string student_id = 3;
   */
  studentId = "";

  /**
   * @generated from field: repeated coteacher.v1.Response.Answer answers = 4;
   */
  answers: Response_Answer[] = [];

  /**
   * @generated from field: string ai_response = 5;
   */
  aiResponse = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 6;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "form_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "student_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "answers", kind: "message", T: Response_Answer, repeated: true },
    { no: 5, name: "ai_response", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Response {
    return new Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Response {
    return new Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Response {
    return new Response().fromJsonString(jsonString, options);
  }

  static equals(a: Response | PlainMessage<Response> | undefined, b: Response | PlainMessage<Response> | undefined): boolean {
    return proto3.util.equals(Response, a, b);
  }
}

/**
 * @generated from message coteacher.v1.Response.Answer
 */
export class Response_Answer extends Message<Response_Answer> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string response_id = 2;
   */
  responseId = "";

  /**
   * @generated from field: string question_id = 3;
   */
  questionId = "";

  /**
   * @generated from field: int32 order = 4;
   */
  order = 0;

  /**
   * @generated from field: string answer_text = 5;
   */
  answerText = "";

  /**
   * @generated from field: repeated string answer_option_ids = 6;
   */
  answerOptionIds: string[] = [];

  constructor(data?: PartialMessage<Response_Answer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "coteacher.v1.Response.Answer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "response_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "question_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "order", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "answer_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "answer_option_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Response_Answer {
    return new Response_Answer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Response_Answer {
    return new Response_Answer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Response_Answer {
    return new Response_Answer().fromJsonString(jsonString, options);
  }

  static equals(a: Response_Answer | PlainMessage<Response_Answer> | undefined, b: Response_Answer | PlainMessage<Response_Answer> | undefined): boolean {
    return proto3.util.equals(Response_Answer, a, b);
  }
}


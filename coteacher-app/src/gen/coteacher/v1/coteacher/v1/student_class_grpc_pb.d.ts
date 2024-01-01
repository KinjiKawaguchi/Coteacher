// GENERATED CODE -- DO NOT EDIT!

// package: coteacher.v1
// file: coteacher/v1/student_class.proto

import * as coteacher_v1_student_class_pb from "../../coteacher/v1/student_class_pb";
import * as grpc from "@grpc/grpc-js";

interface IStudentClassServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createStudentClass: grpc.MethodDefinition<coteacher_v1_student_class_pb.CreateStudentClassRequest, coteacher_v1_student_class_pb.CreateStudentClassResponse>;
  getStudentListByClassID: grpc.MethodDefinition<coteacher_v1_student_class_pb.GetStudentListByClassIDRequest, coteacher_v1_student_class_pb.GetStudentListByClassIDResponse>;
  getClassListByStudentID: grpc.MethodDefinition<coteacher_v1_student_class_pb.GetClassListByStudentIDRequest, coteacher_v1_student_class_pb.GetClassListByStudentIDResponse>;
  deleteStudentClass: grpc.MethodDefinition<coteacher_v1_student_class_pb.DeleteStudentClassRequest, coteacher_v1_student_class_pb.DeleteStudentClassResponse>;
}

export const StudentClassServiceService: IStudentClassServiceService;

export interface IStudentClassServiceServer extends grpc.UntypedServiceImplementation {
  createStudentClass: grpc.handleUnaryCall<coteacher_v1_student_class_pb.CreateStudentClassRequest, coteacher_v1_student_class_pb.CreateStudentClassResponse>;
  getStudentListByClassID: grpc.handleUnaryCall<coteacher_v1_student_class_pb.GetStudentListByClassIDRequest, coteacher_v1_student_class_pb.GetStudentListByClassIDResponse>;
  getClassListByStudentID: grpc.handleUnaryCall<coteacher_v1_student_class_pb.GetClassListByStudentIDRequest, coteacher_v1_student_class_pb.GetClassListByStudentIDResponse>;
  deleteStudentClass: grpc.handleUnaryCall<coteacher_v1_student_class_pb.DeleteStudentClassRequest, coteacher_v1_student_class_pb.DeleteStudentClassResponse>;
}

export class StudentClassServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createStudentClass(argument: coteacher_v1_student_class_pb.CreateStudentClassRequest, callback: grpc.requestCallback<coteacher_v1_student_class_pb.CreateStudentClassResponse>): grpc.ClientUnaryCall;
  createStudentClass(argument: coteacher_v1_student_class_pb.CreateStudentClassRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.CreateStudentClassResponse>): grpc.ClientUnaryCall;
  createStudentClass(argument: coteacher_v1_student_class_pb.CreateStudentClassRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.CreateStudentClassResponse>): grpc.ClientUnaryCall;
  getStudentListByClassID(argument: coteacher_v1_student_class_pb.GetStudentListByClassIDRequest, callback: grpc.requestCallback<coteacher_v1_student_class_pb.GetStudentListByClassIDResponse>): grpc.ClientUnaryCall;
  getStudentListByClassID(argument: coteacher_v1_student_class_pb.GetStudentListByClassIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.GetStudentListByClassIDResponse>): grpc.ClientUnaryCall;
  getStudentListByClassID(argument: coteacher_v1_student_class_pb.GetStudentListByClassIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.GetStudentListByClassIDResponse>): grpc.ClientUnaryCall;
  getClassListByStudentID(argument: coteacher_v1_student_class_pb.GetClassListByStudentIDRequest, callback: grpc.requestCallback<coteacher_v1_student_class_pb.GetClassListByStudentIDResponse>): grpc.ClientUnaryCall;
  getClassListByStudentID(argument: coteacher_v1_student_class_pb.GetClassListByStudentIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.GetClassListByStudentIDResponse>): grpc.ClientUnaryCall;
  getClassListByStudentID(argument: coteacher_v1_student_class_pb.GetClassListByStudentIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.GetClassListByStudentIDResponse>): grpc.ClientUnaryCall;
  deleteStudentClass(argument: coteacher_v1_student_class_pb.DeleteStudentClassRequest, callback: grpc.requestCallback<coteacher_v1_student_class_pb.DeleteStudentClassResponse>): grpc.ClientUnaryCall;
  deleteStudentClass(argument: coteacher_v1_student_class_pb.DeleteStudentClassRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.DeleteStudentClassResponse>): grpc.ClientUnaryCall;
  deleteStudentClass(argument: coteacher_v1_student_class_pb.DeleteStudentClassRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_student_class_pb.DeleteStudentClassResponse>): grpc.ClientUnaryCall;
}

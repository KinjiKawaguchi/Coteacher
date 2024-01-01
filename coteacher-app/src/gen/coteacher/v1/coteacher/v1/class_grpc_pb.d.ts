// GENERATED CODE -- DO NOT EDIT!

// package: coteacher.v1
// file: coteacher/v1/class.proto

import * as coteacher_v1_class_pb from "../../coteacher/v1/class_pb";
import * as grpc from "@grpc/grpc-js";

interface IClassServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createClass: grpc.MethodDefinition<coteacher_v1_class_pb.CreateClassRequest, coteacher_v1_class_pb.CreateClassResponse>;
  getClassByID: grpc.MethodDefinition<coteacher_v1_class_pb.GetClassByIDRequest, coteacher_v1_class_pb.GetClassByIDResponse>;
  getClassListByTeacherID: grpc.MethodDefinition<coteacher_v1_class_pb.GetClassListByTeacherIDRequest, coteacher_v1_class_pb.GetClassListByTeacherIDResponse>;
  updateClass: grpc.MethodDefinition<coteacher_v1_class_pb.UpdateClassRequest, coteacher_v1_class_pb.UpdateClassResponse>;
  deleteClass: grpc.MethodDefinition<coteacher_v1_class_pb.DeleteClassRequest, coteacher_v1_class_pb.DeleteClassResponse>;
}

export const ClassServiceService: IClassServiceService;

export interface IClassServiceServer extends grpc.UntypedServiceImplementation {
  createClass: grpc.handleUnaryCall<coteacher_v1_class_pb.CreateClassRequest, coteacher_v1_class_pb.CreateClassResponse>;
  getClassByID: grpc.handleUnaryCall<coteacher_v1_class_pb.GetClassByIDRequest, coteacher_v1_class_pb.GetClassByIDResponse>;
  getClassListByTeacherID: grpc.handleUnaryCall<coteacher_v1_class_pb.GetClassListByTeacherIDRequest, coteacher_v1_class_pb.GetClassListByTeacherIDResponse>;
  updateClass: grpc.handleUnaryCall<coteacher_v1_class_pb.UpdateClassRequest, coteacher_v1_class_pb.UpdateClassResponse>;
  deleteClass: grpc.handleUnaryCall<coteacher_v1_class_pb.DeleteClassRequest, coteacher_v1_class_pb.DeleteClassResponse>;
}

export class ClassServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createClass(argument: coteacher_v1_class_pb.CreateClassRequest, callback: grpc.requestCallback<coteacher_v1_class_pb.CreateClassResponse>): grpc.ClientUnaryCall;
  createClass(argument: coteacher_v1_class_pb.CreateClassRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.CreateClassResponse>): grpc.ClientUnaryCall;
  createClass(argument: coteacher_v1_class_pb.CreateClassRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.CreateClassResponse>): grpc.ClientUnaryCall;
  getClassByID(argument: coteacher_v1_class_pb.GetClassByIDRequest, callback: grpc.requestCallback<coteacher_v1_class_pb.GetClassByIDResponse>): grpc.ClientUnaryCall;
  getClassByID(argument: coteacher_v1_class_pb.GetClassByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.GetClassByIDResponse>): grpc.ClientUnaryCall;
  getClassByID(argument: coteacher_v1_class_pb.GetClassByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.GetClassByIDResponse>): grpc.ClientUnaryCall;
  getClassListByTeacherID(argument: coteacher_v1_class_pb.GetClassListByTeacherIDRequest, callback: grpc.requestCallback<coteacher_v1_class_pb.GetClassListByTeacherIDResponse>): grpc.ClientUnaryCall;
  getClassListByTeacherID(argument: coteacher_v1_class_pb.GetClassListByTeacherIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.GetClassListByTeacherIDResponse>): grpc.ClientUnaryCall;
  getClassListByTeacherID(argument: coteacher_v1_class_pb.GetClassListByTeacherIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.GetClassListByTeacherIDResponse>): grpc.ClientUnaryCall;
  updateClass(argument: coteacher_v1_class_pb.UpdateClassRequest, callback: grpc.requestCallback<coteacher_v1_class_pb.UpdateClassResponse>): grpc.ClientUnaryCall;
  updateClass(argument: coteacher_v1_class_pb.UpdateClassRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.UpdateClassResponse>): grpc.ClientUnaryCall;
  updateClass(argument: coteacher_v1_class_pb.UpdateClassRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.UpdateClassResponse>): grpc.ClientUnaryCall;
  deleteClass(argument: coteacher_v1_class_pb.DeleteClassRequest, callback: grpc.requestCallback<coteacher_v1_class_pb.DeleteClassResponse>): grpc.ClientUnaryCall;
  deleteClass(argument: coteacher_v1_class_pb.DeleteClassRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.DeleteClassResponse>): grpc.ClientUnaryCall;
  deleteClass(argument: coteacher_v1_class_pb.DeleteClassRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_pb.DeleteClassResponse>): grpc.ClientUnaryCall;
}

// GENERATED CODE -- DO NOT EDIT!

// package: coteacher.v1
// file: coteacher/v1/class_invitation_code.proto

import * as coteacher_v1_class_invitation_code_pb from "../../coteacher/v1/class_invitation_code_pb";
import * as grpc from "@grpc/grpc-js";

interface IClassInvitationCodeServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createClassInvitationCode: grpc.MethodDefinition<coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeRequest, coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeResponse>;
  getClassInvitationCodeByID: grpc.MethodDefinition<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDRequest, coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDResponse>;
  getClassInvitationCodeListByClassID: grpc.MethodDefinition<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDRequest, coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDResponse>;
  updateClassInvitationCode: grpc.MethodDefinition<coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeRequest, coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeResponse>;
}

export const ClassInvitationCodeServiceService: IClassInvitationCodeServiceService;

export interface IClassInvitationCodeServiceServer extends grpc.UntypedServiceImplementation {
  createClassInvitationCode: grpc.handleUnaryCall<coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeRequest, coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeResponse>;
  getClassInvitationCodeByID: grpc.handleUnaryCall<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDRequest, coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDResponse>;
  getClassInvitationCodeListByClassID: grpc.handleUnaryCall<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDRequest, coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDResponse>;
  updateClassInvitationCode: grpc.handleUnaryCall<coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeRequest, coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeResponse>;
}

export class ClassInvitationCodeServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createClassInvitationCode(argument: coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeRequest, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeResponse>): grpc.ClientUnaryCall;
  createClassInvitationCode(argument: coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeResponse>): grpc.ClientUnaryCall;
  createClassInvitationCode(argument: coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.CreateClassInvitationCodeResponse>): grpc.ClientUnaryCall;
  getClassInvitationCodeByID(argument: coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDRequest, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDResponse>): grpc.ClientUnaryCall;
  getClassInvitationCodeByID(argument: coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDResponse>): grpc.ClientUnaryCall;
  getClassInvitationCodeByID(argument: coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeByIDResponse>): grpc.ClientUnaryCall;
  getClassInvitationCodeListByClassID(argument: coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDRequest, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDResponse>): grpc.ClientUnaryCall;
  getClassInvitationCodeListByClassID(argument: coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDResponse>): grpc.ClientUnaryCall;
  getClassInvitationCodeListByClassID(argument: coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.GetClassInvitationCodeListByClassIDResponse>): grpc.ClientUnaryCall;
  updateClassInvitationCode(argument: coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeRequest, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeResponse>): grpc.ClientUnaryCall;
  updateClassInvitationCode(argument: coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeResponse>): grpc.ClientUnaryCall;
  updateClassInvitationCode(argument: coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_class_invitation_code_pb.UpdateClassInvitationCodeResponse>): grpc.ClientUnaryCall;
}

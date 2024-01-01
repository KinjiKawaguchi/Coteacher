// GENERATED CODE -- DO NOT EDIT!

// package: coteacher.v1
// file: coteacher/v1/health_check.proto

import * as coteacher_v1_health_check_pb from "../../coteacher/v1/health_check_pb";
import * as grpc from "@grpc/grpc-js";

interface IHealthcheckServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  unary: grpc.MethodDefinition<coteacher_v1_health_check_pb.UnaryRequest, coteacher_v1_health_check_pb.UnaryResponse>;
  serverStreaming: grpc.MethodDefinition<coteacher_v1_health_check_pb.ServerStreamingRequest, coteacher_v1_health_check_pb.ServerStreamingResponse>;
  clientStreaming: grpc.MethodDefinition<coteacher_v1_health_check_pb.ClientStreamingRequest, coteacher_v1_health_check_pb.ClientStreamingResponse>;
  bidirectionalStreaming: grpc.MethodDefinition<coteacher_v1_health_check_pb.BidirectionalStreamingRequest, coteacher_v1_health_check_pb.BidirectionalStreamingResponse>;
}

export const HealthcheckServiceService: IHealthcheckServiceService;

export interface IHealthcheckServiceServer extends grpc.UntypedServiceImplementation {
  unary: grpc.handleUnaryCall<coteacher_v1_health_check_pb.UnaryRequest, coteacher_v1_health_check_pb.UnaryResponse>;
  serverStreaming: grpc.handleServerStreamingCall<coteacher_v1_health_check_pb.ServerStreamingRequest, coteacher_v1_health_check_pb.ServerStreamingResponse>;
  clientStreaming: grpc.handleClientStreamingCall<coteacher_v1_health_check_pb.ClientStreamingRequest, coteacher_v1_health_check_pb.ClientStreamingResponse>;
  bidirectionalStreaming: grpc.handleBidiStreamingCall<coteacher_v1_health_check_pb.BidirectionalStreamingRequest, coteacher_v1_health_check_pb.BidirectionalStreamingResponse>;
}

export class HealthcheckServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  unary(argument: coteacher_v1_health_check_pb.UnaryRequest, callback: grpc.requestCallback<coteacher_v1_health_check_pb.UnaryResponse>): grpc.ClientUnaryCall;
  unary(argument: coteacher_v1_health_check_pb.UnaryRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_health_check_pb.UnaryResponse>): grpc.ClientUnaryCall;
  unary(argument: coteacher_v1_health_check_pb.UnaryRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_health_check_pb.UnaryResponse>): grpc.ClientUnaryCall;
  serverStreaming(argument: coteacher_v1_health_check_pb.ServerStreamingRequest, metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientReadableStream<coteacher_v1_health_check_pb.ServerStreamingResponse>;
  serverStreaming(argument: coteacher_v1_health_check_pb.ServerStreamingRequest, metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientReadableStream<coteacher_v1_health_check_pb.ServerStreamingResponse>;
  clientStreaming(callback: grpc.requestCallback<coteacher_v1_health_check_pb.ClientStreamingResponse>): grpc.ClientWritableStream<coteacher_v1_health_check_pb.ClientStreamingRequest>;
  clientStreaming(metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_health_check_pb.ClientStreamingResponse>): grpc.ClientWritableStream<coteacher_v1_health_check_pb.ClientStreamingRequest>;
  clientStreaming(metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<coteacher_v1_health_check_pb.ClientStreamingResponse>): grpc.ClientWritableStream<coteacher_v1_health_check_pb.ClientStreamingRequest>;
  bidirectionalStreaming(metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientDuplexStream<coteacher_v1_health_check_pb.BidirectionalStreamingRequest, coteacher_v1_health_check_pb.BidirectionalStreamingResponse>;
  bidirectionalStreaming(metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientDuplexStream<coteacher_v1_health_check_pb.BidirectionalStreamingRequest, coteacher_v1_health_check_pb.BidirectionalStreamingResponse>;
}

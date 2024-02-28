# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from coteacher.v1 import question_pb2 as coteacher_dot_v1_dot_question__pb2


class QuestionServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetQuestionListByFormId = channel.unary_unary(
                '/coteacher.v1.QuestionService/GetQuestionListByFormId',
                request_serializer=coteacher_dot_v1_dot_question__pb2.GetQuestionListByFormIdRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_question__pb2.GetQuestionListByFormIdResponse.FromString,
                )
        self.SaveQuestionList = channel.unary_unary(
                '/coteacher.v1.QuestionService/SaveQuestionList',
                request_serializer=coteacher_dot_v1_dot_question__pb2.SaveQuestionListRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_question__pb2.SaveQuestionListResponse.FromString,
                )


class QuestionServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetQuestionListByFormId(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SaveQuestionList(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QuestionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetQuestionListByFormId': grpc.unary_unary_rpc_method_handler(
                    servicer.GetQuestionListByFormId,
                    request_deserializer=coteacher_dot_v1_dot_question__pb2.GetQuestionListByFormIdRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_question__pb2.GetQuestionListByFormIdResponse.SerializeToString,
            ),
            'SaveQuestionList': grpc.unary_unary_rpc_method_handler(
                    servicer.SaveQuestionList,
                    request_deserializer=coteacher_dot_v1_dot_question__pb2.SaveQuestionListRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_question__pb2.SaveQuestionListResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'coteacher.v1.QuestionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class QuestionService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetQuestionListByFormId(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.QuestionService/GetQuestionListByFormId',
            coteacher_dot_v1_dot_question__pb2.GetQuestionListByFormIdRequest.SerializeToString,
            coteacher_dot_v1_dot_question__pb2.GetQuestionListByFormIdResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SaveQuestionList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.QuestionService/SaveQuestionList',
            coteacher_dot_v1_dot_question__pb2.SaveQuestionListRequest.SerializeToString,
            coteacher_dot_v1_dot_question__pb2.SaveQuestionListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

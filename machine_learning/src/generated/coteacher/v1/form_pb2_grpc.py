# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from coteacher.v1 import form_pb2 as coteacher_dot_v1_dot_form__pb2


class FormServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateForm = channel.unary_unary(
                '/coteacher.v1.FormService/CreateForm',
                request_serializer=coteacher_dot_v1_dot_form__pb2.CreateFormRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.CreateFormResponse.FromString,
                )
        self.GetFormByID = channel.unary_unary(
                '/coteacher.v1.FormService/GetFormByID',
                request_serializer=coteacher_dot_v1_dot_form__pb2.GetFormByIDRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.GetFormByIDResponse.FromString,
                )
        self.GetFormListByClassID = channel.unary_unary(
                '/coteacher.v1.FormService/GetFormListByClassID',
                request_serializer=coteacher_dot_v1_dot_form__pb2.GetFormListByClassIDRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.GetFormListByClassIDResponse.FromString,
                )
        self.UpdateForm = channel.unary_unary(
                '/coteacher.v1.FormService/UpdateForm',
                request_serializer=coteacher_dot_v1_dot_form__pb2.UpdateFormRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.UpdateFormResponse.FromString,
                )
        self.DeleteForm = channel.unary_unary(
                '/coteacher.v1.FormService/DeleteForm',
                request_serializer=coteacher_dot_v1_dot_form__pb2.DeleteFormRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.DeleteFormResponse.FromString,
                )
        self.CheckFormEditPermission = channel.unary_unary(
                '/coteacher.v1.FormService/CheckFormEditPermission',
                request_serializer=coteacher_dot_v1_dot_form__pb2.CheckFormEditPermissionRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.CheckFormEditPermissionResponse.FromString,
                )
        self.CheckFormViewPermission = channel.unary_unary(
                '/coteacher.v1.FormService/CheckFormViewPermission',
                request_serializer=coteacher_dot_v1_dot_form__pb2.CheckFormViewPermissionRequest.SerializeToString,
                response_deserializer=coteacher_dot_v1_dot_form__pb2.CheckFormViewPermissionResponse.FromString,
                )


class FormServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateForm(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFormByID(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFormListByClassID(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateForm(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteForm(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckFormEditPermission(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckFormViewPermission(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_FormServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateForm': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateForm,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.CreateFormRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.CreateFormResponse.SerializeToString,
            ),
            'GetFormByID': grpc.unary_unary_rpc_method_handler(
                    servicer.GetFormByID,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.GetFormByIDRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.GetFormByIDResponse.SerializeToString,
            ),
            'GetFormListByClassID': grpc.unary_unary_rpc_method_handler(
                    servicer.GetFormListByClassID,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.GetFormListByClassIDRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.GetFormListByClassIDResponse.SerializeToString,
            ),
            'UpdateForm': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateForm,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.UpdateFormRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.UpdateFormResponse.SerializeToString,
            ),
            'DeleteForm': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteForm,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.DeleteFormRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.DeleteFormResponse.SerializeToString,
            ),
            'CheckFormEditPermission': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckFormEditPermission,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.CheckFormEditPermissionRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.CheckFormEditPermissionResponse.SerializeToString,
            ),
            'CheckFormViewPermission': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckFormViewPermission,
                    request_deserializer=coteacher_dot_v1_dot_form__pb2.CheckFormViewPermissionRequest.FromString,
                    response_serializer=coteacher_dot_v1_dot_form__pb2.CheckFormViewPermissionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'coteacher.v1.FormService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class FormService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateForm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/CreateForm',
            coteacher_dot_v1_dot_form__pb2.CreateFormRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.CreateFormResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFormByID(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/GetFormByID',
            coteacher_dot_v1_dot_form__pb2.GetFormByIDRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.GetFormByIDResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFormListByClassID(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/GetFormListByClassID',
            coteacher_dot_v1_dot_form__pb2.GetFormListByClassIDRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.GetFormListByClassIDResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateForm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/UpdateForm',
            coteacher_dot_v1_dot_form__pb2.UpdateFormRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.UpdateFormResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteForm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/DeleteForm',
            coteacher_dot_v1_dot_form__pb2.DeleteFormRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.DeleteFormResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CheckFormEditPermission(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/CheckFormEditPermission',
            coteacher_dot_v1_dot_form__pb2.CheckFormEditPermissionRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.CheckFormEditPermissionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CheckFormViewPermission(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/coteacher.v1.FormService/CheckFormViewPermission',
            coteacher_dot_v1_dot_form__pb2.CheckFormViewPermissionRequest.SerializeToString,
            coteacher_dot_v1_dot_form__pb2.CheckFormViewPermissionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

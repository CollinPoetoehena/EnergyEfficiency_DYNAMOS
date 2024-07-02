# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import microserviceCommunication_pb2 as microserviceCommunication__pb2
import rabbitMQ_pb2 as rabbitMQ__pb2

GRPC_GENERATED_VERSION = '1.64.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in rabbitMQ_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class SideCarStub(object):
    """The sidecar definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.InitRabbitMq = channel.unary_unary(
                '/dynamos.SideCar/InitRabbitMq',
                request_serializer=rabbitMQ__pb2.InitRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.Consume = channel.unary_stream(
                '/dynamos.SideCar/Consume',
                request_serializer=rabbitMQ__pb2.ConsumeRequest.SerializeToString,
                response_deserializer=rabbitMQ__pb2.SideCarMessage.FromString,
                _registered_method=True)
        self.ChainConsume = channel.unary_stream(
                '/dynamos.SideCar/ChainConsume',
                request_serializer=rabbitMQ__pb2.ConsumeRequest.SerializeToString,
                response_deserializer=rabbitMQ__pb2.SideCarMessage.FromString,
                _registered_method=True)
        self.SendRequestApproval = channel.unary_unary(
                '/dynamos.SideCar/SendRequestApproval',
                request_serializer=rabbitMQ__pb2.RequestApproval.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendValidationResponse = channel.unary_unary(
                '/dynamos.SideCar/SendValidationResponse',
                request_serializer=rabbitMQ__pb2.ValidationResponse.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendCompositionRequest = channel.unary_unary(
                '/dynamos.SideCar/SendCompositionRequest',
                request_serializer=rabbitMQ__pb2.CompositionRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendSqlDataRequest = channel.unary_unary(
                '/dynamos.SideCar/SendSqlDataRequest',
                request_serializer=rabbitMQ__pb2.SqlDataRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendPolicyUpdate = channel.unary_unary(
                '/dynamos.SideCar/SendPolicyUpdate',
                request_serializer=rabbitMQ__pb2.PolicyUpdate.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendTest = channel.unary_unary(
                '/dynamos.SideCar/SendTest',
                request_serializer=rabbitMQ__pb2.SqlDataRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendMicroserviceComm = channel.unary_unary(
                '/dynamos.SideCar/SendMicroserviceComm',
                request_serializer=microserviceCommunication__pb2.MicroserviceCommunication.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.CreateQueue = channel.unary_unary(
                '/dynamos.SideCar/CreateQueue',
                request_serializer=rabbitMQ__pb2.QueueInfo.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.DeleteQueue = channel.unary_unary(
                '/dynamos.SideCar/DeleteQueue',
                request_serializer=rabbitMQ__pb2.QueueInfo.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendRequestApprovalResponse = channel.unary_unary(
                '/dynamos.SideCar/SendRequestApprovalResponse',
                request_serializer=rabbitMQ__pb2.RequestApprovalResponse.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.SendRequestApprovalRequest = channel.unary_unary(
                '/dynamos.SideCar/SendRequestApprovalRequest',
                request_serializer=rabbitMQ__pb2.RequestApproval.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)


class SideCarServicer(object):
    """The sidecar definition.
    """

    def InitRabbitMq(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Consume(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChainConsume(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendRequestApproval(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendValidationResponse(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendCompositionRequest(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendSqlDataRequest(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendPolicyUpdate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendTest(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendMicroserviceComm(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateQueue(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteQueue(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendRequestApprovalResponse(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendRequestApprovalRequest(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SideCarServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'InitRabbitMq': grpc.unary_unary_rpc_method_handler(
                    servicer.InitRabbitMq,
                    request_deserializer=rabbitMQ__pb2.InitRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'Consume': grpc.unary_stream_rpc_method_handler(
                    servicer.Consume,
                    request_deserializer=rabbitMQ__pb2.ConsumeRequest.FromString,
                    response_serializer=rabbitMQ__pb2.SideCarMessage.SerializeToString,
            ),
            'ChainConsume': grpc.unary_stream_rpc_method_handler(
                    servicer.ChainConsume,
                    request_deserializer=rabbitMQ__pb2.ConsumeRequest.FromString,
                    response_serializer=rabbitMQ__pb2.SideCarMessage.SerializeToString,
            ),
            'SendRequestApproval': grpc.unary_unary_rpc_method_handler(
                    servicer.SendRequestApproval,
                    request_deserializer=rabbitMQ__pb2.RequestApproval.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendValidationResponse': grpc.unary_unary_rpc_method_handler(
                    servicer.SendValidationResponse,
                    request_deserializer=rabbitMQ__pb2.ValidationResponse.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendCompositionRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.SendCompositionRequest,
                    request_deserializer=rabbitMQ__pb2.CompositionRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendSqlDataRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.SendSqlDataRequest,
                    request_deserializer=rabbitMQ__pb2.SqlDataRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendPolicyUpdate': grpc.unary_unary_rpc_method_handler(
                    servicer.SendPolicyUpdate,
                    request_deserializer=rabbitMQ__pb2.PolicyUpdate.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendTest': grpc.unary_unary_rpc_method_handler(
                    servicer.SendTest,
                    request_deserializer=rabbitMQ__pb2.SqlDataRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendMicroserviceComm': grpc.unary_unary_rpc_method_handler(
                    servicer.SendMicroserviceComm,
                    request_deserializer=microserviceCommunication__pb2.MicroserviceCommunication.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'CreateQueue': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateQueue,
                    request_deserializer=rabbitMQ__pb2.QueueInfo.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'DeleteQueue': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteQueue,
                    request_deserializer=rabbitMQ__pb2.QueueInfo.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendRequestApprovalResponse': grpc.unary_unary_rpc_method_handler(
                    servicer.SendRequestApprovalResponse,
                    request_deserializer=rabbitMQ__pb2.RequestApprovalResponse.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendRequestApprovalRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.SendRequestApprovalRequest,
                    request_deserializer=rabbitMQ__pb2.RequestApproval.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'dynamos.SideCar', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('dynamos.SideCar', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SideCar(object):
    """The sidecar definition.
    """

    @staticmethod
    def InitRabbitMq(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/InitRabbitMq',
            rabbitMQ__pb2.InitRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Consume(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/dynamos.SideCar/Consume',
            rabbitMQ__pb2.ConsumeRequest.SerializeToString,
            rabbitMQ__pb2.SideCarMessage.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ChainConsume(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/dynamos.SideCar/ChainConsume',
            rabbitMQ__pb2.ConsumeRequest.SerializeToString,
            rabbitMQ__pb2.SideCarMessage.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendRequestApproval(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendRequestApproval',
            rabbitMQ__pb2.RequestApproval.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendValidationResponse(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendValidationResponse',
            rabbitMQ__pb2.ValidationResponse.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendCompositionRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendCompositionRequest',
            rabbitMQ__pb2.CompositionRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendSqlDataRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendSqlDataRequest',
            rabbitMQ__pb2.SqlDataRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendPolicyUpdate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendPolicyUpdate',
            rabbitMQ__pb2.PolicyUpdate.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendTest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendTest',
            rabbitMQ__pb2.SqlDataRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendMicroserviceComm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendMicroserviceComm',
            microserviceCommunication__pb2.MicroserviceCommunication.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CreateQueue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/CreateQueue',
            rabbitMQ__pb2.QueueInfo.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def DeleteQueue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/DeleteQueue',
            rabbitMQ__pb2.QueueInfo.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendRequestApprovalResponse(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendRequestApprovalResponse',
            rabbitMQ__pb2.RequestApprovalResponse.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendRequestApprovalRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/dynamos.SideCar/SendRequestApprovalRequest',
            rabbitMQ__pb2.RequestApproval.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

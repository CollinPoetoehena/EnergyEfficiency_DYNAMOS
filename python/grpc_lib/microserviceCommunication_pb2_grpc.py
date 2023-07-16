# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import microserviceCommunication_pb2 as microserviceCommunication__pb2


class MicroserviceStub(object):
    """The sidecar definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendData = channel.unary_unary(
                '/proto.Microservice/SendData',
                request_serializer=microserviceCommunication__pb2.MicroserviceCommunication.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.SendShutdownSignal = channel.unary_unary(
                '/proto.Microservice/SendShutdownSignal',
                request_serializer=microserviceCommunication__pb2.ShutDown.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class MicroserviceServicer(object):
    """The sidecar definition.
    """

    def SendData(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendShutdownSignal(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MicroserviceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendData': grpc.unary_unary_rpc_method_handler(
                    servicer.SendData,
                    request_deserializer=microserviceCommunication__pb2.MicroserviceCommunication.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'SendShutdownSignal': grpc.unary_unary_rpc_method_handler(
                    servicer.SendShutdownSignal,
                    request_deserializer=microserviceCommunication__pb2.ShutDown.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.Microservice', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Microservice(object):
    """The sidecar definition.
    """

    @staticmethod
    def SendData(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Microservice/SendData',
            microserviceCommunication__pb2.MicroserviceCommunication.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SendShutdownSignal(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Microservice/SendShutdownSignal',
            microserviceCommunication__pb2.ShutDown.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

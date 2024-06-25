# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import etcd_pb2 as etcd__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class EtcdStub(object):
    """The sidecar definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.InitEtcd = channel.unary_unary(
                '/proto.Etcd/InitEtcd',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.GetDatasetMetadata = channel.unary_unary(
                '/proto.Etcd/GetDatasetMetadata',
                request_serializer=etcd__pb2.EtcdKey.SerializeToString,
                response_deserializer=etcd__pb2.Dataset.FromString,
                )


class EtcdServicer(object):
    """The sidecar definition.
    """

    def InitEtcd(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDatasetMetadata(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EtcdServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'InitEtcd': grpc.unary_unary_rpc_method_handler(
                    servicer.InitEtcd,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetDatasetMetadata': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDatasetMetadata,
                    request_deserializer=etcd__pb2.EtcdKey.FromString,
                    response_serializer=etcd__pb2.Dataset.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.Etcd', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Etcd(object):
    """The sidecar definition.
    """

    @staticmethod
    def InitEtcd(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Etcd/InitEtcd',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetDatasetMetadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Etcd/GetDatasetMetadata',
            etcd__pb2.EtcdKey.SerializeToString,
            etcd__pb2.Dataset.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

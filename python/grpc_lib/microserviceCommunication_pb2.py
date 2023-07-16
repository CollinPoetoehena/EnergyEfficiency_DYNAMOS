# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: microserviceCommunication.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
import generic_pb2 as generic__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fmicroserviceCommunication.proto\x12\x05proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x19google/protobuf/any.proto\x1a\rgeneric.proto\"\xb7\x02\n\x19MicroserviceCommunication\x12\x0c\n\x04type\x18\x01 \x01(\t\x12\x14\n\x0crequest_type\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12@\n\x08metadata\x18\x04 \x03(\x0b\x32..proto.MicroserviceCommunication.MetadataEntry\x12.\n\x10original_request\x18\x05 \x01(\x0b\x32\x14.google.protobuf.Any\x12,\n\x0erequest_metada\x18\x06 \x01(\x0b\x32\x14.proto.RequestMetada\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x18\n\x08ShutDown\x12\x0c\n\x04name\x18\x01 \x01(\t2\x97\x01\n\x0cMicroservice\x12\x46\n\x08SendData\x12 .proto.MicroserviceCommunication\x1a\x16.google.protobuf.Empty\"\x00\x12?\n\x12SendShutdownSignal\x12\x0f.proto.ShutDown\x1a\x16.google.protobuf.Empty\"\x00\x42\'Z%github.com/Jorrit05/DYNAMOS/pkg/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'microserviceCommunication_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z%github.com/Jorrit05/DYNAMOS/pkg/proto'
  _MICROSERVICECOMMUNICATION_METADATAENTRY._options = None
  _MICROSERVICECOMMUNICATION_METADATAENTRY._serialized_options = b'8\001'
  _globals['_MICROSERVICECOMMUNICATION']._serialized_start=144
  _globals['_MICROSERVICECOMMUNICATION']._serialized_end=455
  _globals['_MICROSERVICECOMMUNICATION_METADATAENTRY']._serialized_start=408
  _globals['_MICROSERVICECOMMUNICATION_METADATAENTRY']._serialized_end=455
  _globals['_SHUTDOWN']._serialized_start=457
  _globals['_SHUTDOWN']._serialized_end=481
  _globals['_MICROSERVICE']._serialized_start=484
  _globals['_MICROSERVICE']._serialized_end=635
# @@protoc_insertion_point(module_scope)

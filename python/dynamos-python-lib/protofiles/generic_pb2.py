# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: generic.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'generic.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rgeneric.proto\x12\x07\x64ynamos\x1a\x1bgoogle/protobuf/empty.proto\"~\n\x0fRequestMetadata\x12\x16\n\x0e\x63orrelation_id\x18\x01 \x01(\t\x12\x19\n\x11\x64\x65stination_queue\x18\x02 \x01(\t\x12\x10\n\x08job_name\x18\x03 \x01(\t\x12\x16\n\x0ereturn_address\x18\x04 \x01(\t\x12\x0e\n\x06job_id\x18\x05 \x01(\t\"#\n\x0bServiceName\x12\x14\n\x0cservice_name\x18\x01 \x01(\t2G\n\x07Generic\x12<\n\nInitTracer\x12\x14.dynamos.ServiceName\x1a\x16.google.protobuf.Empty\"\x00\x42\'Z%github.com/Jorrit05/DYNAMOS/pkg/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'generic_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z%github.com/Jorrit05/DYNAMOS/pkg/proto'
  _globals['_REQUESTMETADATA']._serialized_start=55
  _globals['_REQUESTMETADATA']._serialized_end=181
  _globals['_SERVICENAME']._serialized_start=183
  _globals['_SERVICENAME']._serialized_end=218
  _globals['_GENERIC']._serialized_start=220
  _globals['_GENERIC']._serialized_end=291
# @@protoc_insertion_point(module_scope)

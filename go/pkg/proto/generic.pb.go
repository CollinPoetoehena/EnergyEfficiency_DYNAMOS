// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.12.4
// source: generic.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestMetadata struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	CorrelationId    string                 `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	DestinationQueue string                 `protobuf:"bytes,2,opt,name=destination_queue,json=destinationQueue,proto3" json:"destination_queue,omitempty"`
	JobName          string                 `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	ReturnAddress    string                 `protobuf:"bytes,4,opt,name=return_address,json=returnAddress,proto3" json:"return_address,omitempty"`
	JobId            string                 `protobuf:"bytes,5,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	mi := &file_generic_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMetadata) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RequestMetadata) GetDestinationQueue() string {
	if x != nil {
		return x.DestinationQueue
	}
	return ""
}

func (x *RequestMetadata) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *RequestMetadata) GetReturnAddress() string {
	if x != nil {
		return x.ReturnAddress
	}
	return ""
}

func (x *RequestMetadata) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type ServiceName struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceName   string                 `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceName) Reset() {
	*x = ServiceName{}
	mi := &file_generic_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceName) ProtoMessage() {}

func (x *ServiceName) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceName.ProtoReflect.Descriptor instead.
func (*ServiceName) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceName) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

var File_generic_proto protoreflect.FileDescriptor

var file_generic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x47, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x12, 0x3c, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4a, 0x6f, 0x72, 0x72, 0x69, 0x74, 0x30, 0x35, 0x2f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x4f, 0x53,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_generic_proto_rawDescOnce sync.Once
	file_generic_proto_rawDescData = file_generic_proto_rawDesc
)

func file_generic_proto_rawDescGZIP() []byte {
	file_generic_proto_rawDescOnce.Do(func() {
		file_generic_proto_rawDescData = protoimpl.X.CompressGZIP(file_generic_proto_rawDescData)
	})
	return file_generic_proto_rawDescData
}

var file_generic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_generic_proto_goTypes = []any{
	(*RequestMetadata)(nil), // 0: dynamos.RequestMetadata
	(*ServiceName)(nil),     // 1: dynamos.ServiceName
	(*empty.Empty)(nil),     // 2: google.protobuf.Empty
}
var file_generic_proto_depIdxs = []int32{
	1, // 0: dynamos.Generic.InitTracer:input_type -> dynamos.ServiceName
	2, // 1: dynamos.Generic.InitTracer:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_generic_proto_init() }
func file_generic_proto_init() {
	if File_generic_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_generic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generic_proto_goTypes,
		DependencyIndexes: file_generic_proto_depIdxs,
		MessageInfos:      file_generic_proto_msgTypes,
	}.Build()
	File_generic_proto = out.File
	file_generic_proto_rawDesc = nil
	file_generic_proto_goTypes = nil
	file_generic_proto_depIdxs = nil
}

// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        (unknown)
// source: configuration.proto

package hlmirbft

import (
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

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "mirbft".
type ConfigMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consenters []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options    *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ConfigMetadata) Reset() {
	*x = ConfigMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMetadata) ProtoMessage() {}

func (x *ConfigMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMetadata.ProtoReflect.Descriptor instead.
func (*ConfigMetadata) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigMetadata) GetConsenters() []*Consenter {
	if x != nil {
		return x.Consenters
	}
	return nil
}

func (x *ConfigMetadata) GetOptions() *Options {
	if x != nil {
		return x.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host          string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port          uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ClientTlsCert []byte `protobuf:"bytes,3,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert []byte `protobuf:"bytes,4,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
}

func (x *Consenter) Reset() {
	*x = Consenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consenter) ProtoMessage() {}

func (x *Consenter) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consenter.ProtoReflect.Descriptor instead.
func (*Consenter) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *Consenter) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Consenter) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Consenter) GetClientTlsCert() []byte {
	if x != nil {
		return x.ClientTlsCert
	}
	return nil
}

func (x *Consenter) GetServerTlsCert() []byte {
	if x != nil {
		return x.ServerTlsCert
	}
	return nil
}

// Options to be specified for all the hyperledger-labs/mirbft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HeartbeatTicks is the number of ticks before a heartbeat is emitted
	// by a leader.
	HeartbeatTicks uint32 `protobuf:"varint,1,opt,name=HeartbeatTicks,json=heartbeatTicks,proto3" json:"HeartbeatTicks,omitempty"`
	// SuspectTicks is the number of ticks a bucket may not progress before
	// the node suspects the epoch has gone bad.
	SuspectTicks uint32 `protobuf:"varint,2,opt,name=SuspectTicks,json=suspectTicks,proto3" json:"SuspectTicks,omitempty"`
	// NewEpochTimeoutTicks is the number of ticks a replica will wait until
	// it suspects the epoch leader has failed.  This value must be greater
	// than 1, as rebroadcast ticks are computed as half this value.
	NewEpochTimeoutTicks uint32 `protobuf:"varint,3,opt,name=NewEpochTimeoutTicks,json=newEpochTimeoutTicks,proto3" json:"NewEpochTimeoutTicks,omitempty"`
	// BufferSize is the total size of messages which can be held by the state
	// machine, pending application, for each node.  This is necessary because
	// there may be dependencies between messages (for instance, until a checkpoint
	// result is computed, watermarks cannot advance).  This should be set
	// to a minimum of a few MB.
	BufferSize uint32 `protobuf:"varint,4,opt,name=BufferSize,json=bufferSize,proto3" json:"BufferSize,omitempty"`
	//JIRA FLY2-103 Adding Protocol parameters to config metadata
	CheckpointInterval int32  `protobuf:"varint,5,opt,name=CheckpointInterval,json=checkpointInterval,proto3" json:"CheckpointInterval,omitempty"`
	MaxEpochLength     uint64 `protobuf:"varint,6,opt,name=MaxEpochLength,json=maxEpochLength,proto3" json:"MaxEpochLength,omitempty"`
	NumberOfBuckets    int32  `protobuf:"varint,7,opt,name=NumberOfBuckets,json=numberOfBuckets,proto3" json:"NumberOfBuckets,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *Options) GetHeartbeatTicks() uint32 {
	if x != nil {
		return x.HeartbeatTicks
	}
	return 0
}

func (x *Options) GetSuspectTicks() uint32 {
	if x != nil {
		return x.SuspectTicks
	}
	return 0
}

func (x *Options) GetNewEpochTimeoutTicks() uint32 {
	if x != nil {
		return x.NewEpochTimeoutTicks
	}
	return 0
}

func (x *Options) GetBufferSize() uint32 {
	if x != nil {
		return x.BufferSize
	}
	return 0
}

func (x *Options) GetCheckpointInterval() int32 {
	if x != nil {
		return x.CheckpointInterval
	}
	return 0
}

func (x *Options) GetMaxEpochLength() uint64 {
	if x != nil {
		return x.MaxEpochLength
	}
	return 0
}

func (x *Options) GetNumberOfBuckets() int32 {
	if x != nil {
		return x.NumberOfBuckets
	}
	return 0
}

var File_configuration_proto protoreflect.FileDescriptor

var file_configuration_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x6c, 0x6d, 0x69, 0x72, 0x62, 0x66, 0x74, 0x22,
	0x72, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x6c, 0x6d, 0x69, 0x72, 0x62, 0x66, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x6c, 0x6d, 0x69, 0x72, 0x62,
	0x66, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6c, 0x73, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x53, 0x75, 0x73, 0x70, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x73, 0x70, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x73, 0x12, 0x32, 0x0a, 0x14, 0x4e, 0x65, 0x77, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x14, 0x6e, 0x65, 0x77, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6d,
	0x61, 0x78, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x28, 0x0a,
	0x0f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x64, 0x0a, 0x2b, 0x6f, 0x72, 0x67, 0x2e, 0x66,
	0x6c, 0x79, 0x32, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x68, 0x6c,
	0x6d, 0x69, 0x72, 0x62, 0x66, 0x74, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x32, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x2f, 0x68, 0x6c, 0x6d, 0x69, 0x72, 0x62, 0x66, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configuration_proto_rawDescOnce sync.Once
	file_configuration_proto_rawDescData = file_configuration_proto_rawDesc
)

func file_configuration_proto_rawDescGZIP() []byte {
	file_configuration_proto_rawDescOnce.Do(func() {
		file_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_configuration_proto_rawDescData)
	})
	return file_configuration_proto_rawDescData
}

var file_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_configuration_proto_goTypes = []interface{}{
	(*ConfigMetadata)(nil), // 0: hlmirbft.ConfigMetadata
	(*Consenter)(nil),      // 1: hlmirbft.Consenter
	(*Options)(nil),        // 2: hlmirbft.Options
}
var file_configuration_proto_depIdxs = []int32{
	1, // 0: hlmirbft.ConfigMetadata.consenters:type_name -> hlmirbft.Consenter
	2, // 1: hlmirbft.ConfigMetadata.options:type_name -> hlmirbft.Options
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_configuration_proto_init() }
func file_configuration_proto_init() {
	if File_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consenter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_configuration_proto_goTypes,
		DependencyIndexes: file_configuration_proto_depIdxs,
		MessageInfos:      file_configuration_proto_msgTypes,
	}.Build()
	File_configuration_proto = out.File
	file_configuration_proto_rawDesc = nil
	file_configuration_proto_goTypes = nil
	file_configuration_proto_depIdxs = nil
}

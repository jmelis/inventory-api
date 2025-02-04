// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/reporter_data.proto

package v1beta1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReporterData_ReporterType int32

const (
	ReporterData_REPORTER_TYPE_UNSPECIFIED ReporterData_ReporterType = 0
	ReporterData_REPORTER_TYPE_OTHER       ReporterData_ReporterType = 1
	ReporterData_ACM                       ReporterData_ReporterType = 2
	ReporterData_HBI                       ReporterData_ReporterType = 3
	ReporterData_OCM                       ReporterData_ReporterType = 4
	ReporterData_NOTIFICATIONS             ReporterData_ReporterType = 5
)

// Enum value maps for ReporterData_ReporterType.
var (
	ReporterData_ReporterType_name = map[int32]string{
		0: "REPORTER_TYPE_UNSPECIFIED",
		1: "REPORTER_TYPE_OTHER",
		2: "ACM",
		3: "HBI",
		4: "OCM",
		5: "NOTIFICATIONS",
	}
	ReporterData_ReporterType_value = map[string]int32{
		"REPORTER_TYPE_UNSPECIFIED": 0,
		"REPORTER_TYPE_OTHER":       1,
		"ACM":                       2,
		"HBI":                       3,
		"OCM":                       4,
		"NOTIFICATIONS":             5,
	}
)

func (x ReporterData_ReporterType) Enum() *ReporterData_ReporterType {
	p := new(ReporterData_ReporterType)
	*p = x
	return p
}

func (x ReporterData_ReporterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReporterData_ReporterType) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta1_reporter_data_proto_enumTypes[0].Descriptor()
}

func (ReporterData_ReporterType) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta1_reporter_data_proto_enumTypes[0]
}

func (x ReporterData_ReporterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReporterData_ReporterType.Descriptor instead.
func (ReporterData_ReporterType) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_reporter_data_proto_rawDescGZIP(), []int{0, 0}
}

type ReporterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReporterType ReporterData_ReporterType `protobuf:"varint,245778392,opt,name=reporter_type,json=reporterType,proto3,enum=kessel.inventory.v1beta1.ReporterData_ReporterType" json:"reporter_type,omitempty"`
	// The ID of the instance of the reporter. This is derived from the
	// authentication mechanism, i.e. authentication token.
	ReporterInstanceId string `protobuf:"bytes,241085112,opt,name=reporter_instance_id,json=reporterInstanceId,proto3" json:"reporter_instance_id,omitempty"`
	// Date and time when the inventory item was first reported by this reporter
	FirstReported *timestamppb.Timestamp `protobuf:"bytes,13874816,opt,name=first_reported,json=firstReported,proto3" json:"first_reported,omitempty"`
	// Date and time when the inventory item was last updated by this reporter
	LastReported *timestamppb.Timestamp `protobuf:"bytes,436473483,opt,name=last_reported,json=lastReported,proto3" json:"last_reported,omitempty"`
	// The URL for this resource in the reporter's management UI console. For
	// example this would be the cluster URL in the HCC Console for an OCM
	// reported cluster.
	ConsoleHref string `protobuf:"bytes,145854740,opt,name=console_href,json=consoleHref,proto3" json:"console_href,omitempty"`
	// Reporter specific API link to the resource.
	ApiHref string `protobuf:"bytes,430210609,opt,name=api_href,json=apiHref,proto3" json:"api_href,omitempty"`
	// The ID assigned to this resource by the reporter, for example OCM cluster
	// ID, HBI's host id, or ACM managed cluster name etc.
	LocalResourceId string `protobuf:"bytes,508976189,opt,name=local_resource_id,json=localResourceId,proto3" json:"local_resource_id,omitempty"`
	// version of the reporter
	ReporterVersion string `protobuf:"bytes,269629306,opt,name=reporter_version,json=reporterVersion,proto3" json:"reporter_version,omitempty"`
}

func (x *ReporterData) Reset() {
	*x = ReporterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_reporter_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReporterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReporterData) ProtoMessage() {}

func (x *ReporterData) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_reporter_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReporterData.ProtoReflect.Descriptor instead.
func (*ReporterData) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_reporter_data_proto_rawDescGZIP(), []int{0}
}

func (x *ReporterData) GetReporterType() ReporterData_ReporterType {
	if x != nil {
		return x.ReporterType
	}
	return ReporterData_REPORTER_TYPE_UNSPECIFIED
}

func (x *ReporterData) GetReporterInstanceId() string {
	if x != nil {
		return x.ReporterInstanceId
	}
	return ""
}

func (x *ReporterData) GetFirstReported() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstReported
	}
	return nil
}

func (x *ReporterData) GetLastReported() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReported
	}
	return nil
}

func (x *ReporterData) GetConsoleHref() string {
	if x != nil {
		return x.ConsoleHref
	}
	return ""
}

func (x *ReporterData) GetApiHref() string {
	if x != nil {
		return x.ApiHref
	}
	return ""
}

func (x *ReporterData) GetLocalResourceId() string {
	if x != nil {
		return x.LocalResourceId
	}
	return ""
}

func (x *ReporterData) GetReporterVersion() string {
	if x != nil {
		return x.ReporterVersion
	}
	return ""
}

var File_kessel_inventory_v1beta1_reporter_data_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_reporter_data_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x04, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x67, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xd8, 0x8f, 0x99, 0x75, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52,
	0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a,
	0x14, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xb8, 0xd5, 0xfa, 0x72, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x80, 0xed, 0xce, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x12, 0x48, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x8b, 0x9d, 0x90, 0xd0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5f, 0x68, 0x72, 0x65, 0x66, 0x18, 0x94, 0xa2, 0xc6,
	0x45, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x48, 0x72,
	0x65, 0x66, 0x12, 0x1d, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x72, 0x65, 0x66, 0x18, 0xb1,
	0xfc, 0x91, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x48, 0x72, 0x65,
	0x66, 0x12, 0x37, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xbd, 0xb8, 0xd9, 0xf2, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xfa,
	0xee, 0xc8, 0x80, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x0c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x50, 0x4f,
	0x52, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4d, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x42,
	0x49, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x43, 0x4d, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x05, 0x42,
	0x72, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_reporter_data_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_reporter_data_proto_rawDescData = file_kessel_inventory_v1beta1_reporter_data_proto_rawDesc
)

func file_kessel_inventory_v1beta1_reporter_data_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_reporter_data_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_reporter_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_reporter_data_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_reporter_data_proto_rawDescData
}

var file_kessel_inventory_v1beta1_reporter_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kessel_inventory_v1beta1_reporter_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_reporter_data_proto_goTypes = []any{
	(ReporterData_ReporterType)(0), // 0: kessel.inventory.v1beta1.ReporterData.ReporterType
	(*ReporterData)(nil),           // 1: kessel.inventory.v1beta1.ReporterData
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_kessel_inventory_v1beta1_reporter_data_proto_depIdxs = []int32{
	0, // 0: kessel.inventory.v1beta1.ReporterData.reporter_type:type_name -> kessel.inventory.v1beta1.ReporterData.ReporterType
	2, // 1: kessel.inventory.v1beta1.ReporterData.first_reported:type_name -> google.protobuf.Timestamp
	2, // 2: kessel.inventory.v1beta1.ReporterData.last_reported:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_reporter_data_proto_init() }
func file_kessel_inventory_v1beta1_reporter_data_proto_init() {
	if File_kessel_inventory_v1beta1_reporter_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_reporter_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReporterData); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_reporter_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_reporter_data_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_reporter_data_proto_depIdxs,
		EnumInfos:         file_kessel_inventory_v1beta1_reporter_data_proto_enumTypes,
		MessageInfos:      file_kessel_inventory_v1beta1_reporter_data_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_reporter_data_proto = out.File
	file_kessel_inventory_v1beta1_reporter_data_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_reporter_data_proto_goTypes = nil
	file_kessel_inventory_v1beta1_reporter_data_proto_depIdxs = nil
}

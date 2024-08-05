// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/resource_tags.proto

package v1beta1

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

type ResourceTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*ResourceTag `protobuf:"bytes,3552281,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ResourceTags) Reset() {
	*x = ResourceTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resource_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceTags) ProtoMessage() {}

func (x *ResourceTags) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resource_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceTags.ProtoReflect.Descriptor instead.
func (*ResourceTags) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resource_tags_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceTags) GetTags() []*ResourceTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_kessel_inventory_v1beta1_resource_tags_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_resource_tags_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2b, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x99, 0xe8,
	0xd8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x42, 0x72, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_resource_tags_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_resource_tags_proto_rawDescData = file_kessel_inventory_v1beta1_resource_tags_proto_rawDesc
)

func file_kessel_inventory_v1beta1_resource_tags_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_resource_tags_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_resource_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_resource_tags_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_resource_tags_proto_rawDescData
}

var file_kessel_inventory_v1beta1_resource_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_resource_tags_proto_goTypes = []any{
	(*ResourceTags)(nil), // 0: kessel.inventory.v1beta1.ResourceTags
	(*ResourceTag)(nil),  // 1: kessel.inventory.v1beta1.ResourceTag
}
var file_kessel_inventory_v1beta1_resource_tags_proto_depIdxs = []int32{
	1, // 0: kessel.inventory.v1beta1.ResourceTags.tags:type_name -> kessel.inventory.v1beta1.ResourceTag
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_resource_tags_proto_init() }
func file_kessel_inventory_v1beta1_resource_tags_proto_init() {
	if File_kessel_inventory_v1beta1_resource_tags_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_resource_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_resource_tags_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceTags); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_resource_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_resource_tags_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_resource_tags_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_resource_tags_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_resource_tags_proto = out.File
	file_kessel_inventory_v1beta1_resource_tags_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_resource_tags_proto_goTypes = nil
	file_kessel_inventory_v1beta1_resource_tags_proto_depIdxs = nil
}

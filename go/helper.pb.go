// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: helper.proto

package __

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

type HelperVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelperVar) Reset() {
	*x = HelperVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelperVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelperVar) ProtoMessage() {}

func (x *HelperVar) ProtoReflect() protoreflect.Message {
	mi := &file_helper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelperVar.ProtoReflect.Descriptor instead.
func (*HelperVar) Descriptor() ([]byte, []int) {
	return file_helper_proto_rawDescGZIP(), []int{0}
}

func (x *HelperVar) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HelperVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Accounting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateBy int64  `protobuf:"varint,1,opt,name=createBy,proto3" json:"createBy,omitempty"`
	CreateAt string `protobuf:"bytes,2,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateBy int64  `protobuf:"varint,3,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	UpdateAt string `protobuf:"bytes,4,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	DeleteBy int64  `protobuf:"varint,5,opt,name=deleteBy,proto3" json:"deleteBy,omitempty"`
	DeleteAt string `protobuf:"bytes,6,opt,name=deleteAt,proto3" json:"deleteAt,omitempty"`
}

func (x *Accounting) Reset() {
	*x = Accounting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accounting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accounting) ProtoMessage() {}

func (x *Accounting) ProtoReflect() protoreflect.Message {
	mi := &file_helper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accounting.ProtoReflect.Descriptor instead.
func (*Accounting) Descriptor() ([]byte, []int) {
	return file_helper_proto_rawDescGZIP(), []int{1}
}

func (x *Accounting) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *Accounting) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *Accounting) GetUpdateBy() int64 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *Accounting) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *Accounting) GetDeleteBy() int64 {
	if x != nil {
		return x.DeleteBy
	}
	return 0
}

func (x *Accounting) GetDeleteAt() string {
	if x != nil {
		return x.DeleteAt
	}
	return ""
}

var File_helper_proto protoreflect.FileDescriptor

var file_helper_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x73, 0x69, 0x6d, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56,
	0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helper_proto_rawDescOnce sync.Once
	file_helper_proto_rawDescData = file_helper_proto_rawDesc
)

func file_helper_proto_rawDescGZIP() []byte {
	file_helper_proto_rawDescOnce.Do(func() {
		file_helper_proto_rawDescData = protoimpl.X.CompressGZIP(file_helper_proto_rawDescData)
	})
	return file_helper_proto_rawDescData
}

var file_helper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_helper_proto_goTypes = []interface{}{
	(*HelperVar)(nil),  // 0: simrs.HelperVar
	(*Accounting)(nil), // 1: simrs.Accounting
}
var file_helper_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helper_proto_init() }
func file_helper_proto_init() {
	if File_helper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelperVar); i {
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
		file_helper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accounting); i {
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
			RawDescriptor: file_helper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helper_proto_goTypes,
		DependencyIndexes: file_helper_proto_depIdxs,
		MessageInfos:      file_helper_proto_msgTypes,
	}.Build()
	File_helper_proto = out.File
	file_helper_proto_rawDesc = nil
	file_helper_proto_goTypes = nil
	file_helper_proto_depIdxs = nil
}

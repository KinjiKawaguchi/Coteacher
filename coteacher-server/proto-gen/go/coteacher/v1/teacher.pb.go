// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: coteacher/v1/teacher.proto

package coteacherv1

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

type CheckTeacherExistsByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckTeacherExistsByIDRequest) Reset() {
	*x = CheckTeacherExistsByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTeacherExistsByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTeacherExistsByIDRequest) ProtoMessage() {}

func (x *CheckTeacherExistsByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTeacherExistsByIDRequest.ProtoReflect.Descriptor instead.
func (*CheckTeacherExistsByIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *CheckTeacherExistsByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckTeacherExistsByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *CheckTeacherExistsByIDResponse) Reset() {
	*x = CheckTeacherExistsByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTeacherExistsByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTeacherExistsByIDResponse) ProtoMessage() {}

func (x *CheckTeacherExistsByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTeacherExistsByIDResponse.ProtoReflect.Descriptor instead.
func (*CheckTeacherExistsByIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *CheckTeacherExistsByIDResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type CheckTeacherExistsByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CheckTeacherExistsByEmailRequest) Reset() {
	*x = CheckTeacherExistsByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTeacherExistsByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTeacherExistsByEmailRequest) ProtoMessage() {}

func (x *CheckTeacherExistsByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTeacherExistsByEmailRequest.ProtoReflect.Descriptor instead.
func (*CheckTeacherExistsByEmailRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *CheckTeacherExistsByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CheckTeacherExistsByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *CheckTeacherExistsByEmailResponse) Reset() {
	*x = CheckTeacherExistsByEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTeacherExistsByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTeacherExistsByEmailResponse) ProtoMessage() {}

func (x *CheckTeacherExistsByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTeacherExistsByEmailResponse.ProtoReflect.Descriptor instead.
func (*CheckTeacherExistsByEmailResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *CheckTeacherExistsByEmailResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_coteacher_v1_teacher_proto protoreflect.FileDescriptor

var file_coteacher_v1_teacher_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x2f, 0x0a, 0x1d, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x1e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x3b, 0x0a, 0x21, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0x83, 0x02, 0x0a,
	0x0e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x73, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xbc, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x6e, 0x6a, 0x69, 0x4b, 0x61, 0x77, 0x61, 0x67, 0x75, 0x63,
	0x68, 0x69, 0x2f, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0d, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coteacher_v1_teacher_proto_rawDescOnce sync.Once
	file_coteacher_v1_teacher_proto_rawDescData = file_coteacher_v1_teacher_proto_rawDesc
)

func file_coteacher_v1_teacher_proto_rawDescGZIP() []byte {
	file_coteacher_v1_teacher_proto_rawDescOnce.Do(func() {
		file_coteacher_v1_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_coteacher_v1_teacher_proto_rawDescData)
	})
	return file_coteacher_v1_teacher_proto_rawDescData
}

var file_coteacher_v1_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coteacher_v1_teacher_proto_goTypes = []interface{}{
	(*CheckTeacherExistsByIDRequest)(nil),     // 0: coteacher.v1.CheckTeacherExistsByIDRequest
	(*CheckTeacherExistsByIDResponse)(nil),    // 1: coteacher.v1.CheckTeacherExistsByIDResponse
	(*CheckTeacherExistsByEmailRequest)(nil),  // 2: coteacher.v1.CheckTeacherExistsByEmailRequest
	(*CheckTeacherExistsByEmailResponse)(nil), // 3: coteacher.v1.CheckTeacherExistsByEmailResponse
}
var file_coteacher_v1_teacher_proto_depIdxs = []int32{
	0, // 0: coteacher.v1.TeacherService.CheckTeacherExistsByID:input_type -> coteacher.v1.CheckTeacherExistsByIDRequest
	2, // 1: coteacher.v1.TeacherService.CheckTeacherExistsByEmail:input_type -> coteacher.v1.CheckTeacherExistsByEmailRequest
	1, // 2: coteacher.v1.TeacherService.CheckTeacherExistsByID:output_type -> coteacher.v1.CheckTeacherExistsByIDResponse
	3, // 3: coteacher.v1.TeacherService.CheckTeacherExistsByEmail:output_type -> coteacher.v1.CheckTeacherExistsByEmailResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coteacher_v1_teacher_proto_init() }
func file_coteacher_v1_teacher_proto_init() {
	if File_coteacher_v1_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coteacher_v1_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTeacherExistsByIDRequest); i {
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
		file_coteacher_v1_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTeacherExistsByIDResponse); i {
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
		file_coteacher_v1_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTeacherExistsByEmailRequest); i {
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
		file_coteacher_v1_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTeacherExistsByEmailResponse); i {
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
			RawDescriptor: file_coteacher_v1_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coteacher_v1_teacher_proto_goTypes,
		DependencyIndexes: file_coteacher_v1_teacher_proto_depIdxs,
		MessageInfos:      file_coteacher_v1_teacher_proto_msgTypes,
	}.Build()
	File_coteacher_v1_teacher_proto = out.File
	file_coteacher_v1_teacher_proto_rawDesc = nil
	file_coteacher_v1_teacher_proto_goTypes = nil
	file_coteacher_v1_teacher_proto_depIdxs = nil
}

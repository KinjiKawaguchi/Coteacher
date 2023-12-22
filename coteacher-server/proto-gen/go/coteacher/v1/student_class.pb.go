// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: coteacher/v1/student_class.proto

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

type CreateStudentClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	ClassId   string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *CreateStudentClassRequest) Reset() {
	*x = CreateStudentClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentClassRequest) ProtoMessage() {}

func (x *CreateStudentClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentClassRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentClassRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStudentClassRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateStudentClassRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type CreateStudentClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	ClassId   string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *CreateStudentClassResponse) Reset() {
	*x = CreateStudentClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentClassResponse) ProtoMessage() {}

func (x *CreateStudentClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentClassResponse.ProtoReflect.Descriptor instead.
func (*CreateStudentClassResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStudentClassResponse) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateStudentClassResponse) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type GetStudentListByClassIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *GetStudentListByClassIDRequest) Reset() {
	*x = GetStudentListByClassIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentListByClassIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentListByClassIDRequest) ProtoMessage() {}

func (x *GetStudentListByClassIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentListByClassIDRequest.ProtoReflect.Descriptor instead.
func (*GetStudentListByClassIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{2}
}

func (x *GetStudentListByClassIDRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type GetStudentListByClassIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*User `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetStudentListByClassIDResponse) Reset() {
	*x = GetStudentListByClassIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentListByClassIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentListByClassIDResponse) ProtoMessage() {}

func (x *GetStudentListByClassIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentListByClassIDResponse.ProtoReflect.Descriptor instead.
func (*GetStudentListByClassIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{3}
}

func (x *GetStudentListByClassIDResponse) GetStudents() []*User {
	if x != nil {
		return x.Students
	}
	return nil
}

type GetClassListByStudentIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *GetClassListByStudentIDRequest) Reset() {
	*x = GetClassListByStudentIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassListByStudentIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassListByStudentIDRequest) ProtoMessage() {}

func (x *GetClassListByStudentIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassListByStudentIDRequest.ProtoReflect.Descriptor instead.
func (*GetClassListByStudentIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{4}
}

func (x *GetClassListByStudentIDRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type GetClassListByStudentIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes []*Class `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
}

func (x *GetClassListByStudentIDResponse) Reset() {
	*x = GetClassListByStudentIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassListByStudentIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassListByStudentIDResponse) ProtoMessage() {}

func (x *GetClassListByStudentIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassListByStudentIDResponse.ProtoReflect.Descriptor instead.
func (*GetClassListByStudentIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{5}
}

func (x *GetClassListByStudentIDResponse) GetClasses() []*Class {
	if x != nil {
		return x.Classes
	}
	return nil
}

type DeleteStudentClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	ClassId   string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *DeleteStudentClassRequest) Reset() {
	*x = DeleteStudentClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentClassRequest) ProtoMessage() {}

func (x *DeleteStudentClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentClassRequest.ProtoReflect.Descriptor instead.
func (*DeleteStudentClassRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteStudentClassRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *DeleteStudentClassRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type DeleteStudentClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	ClassId   string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *DeleteStudentClassResponse) Reset() {
	*x = DeleteStudentClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_student_class_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentClassResponse) ProtoMessage() {}

func (x *DeleteStudentClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_student_class_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentClassResponse.ProtoReflect.Descriptor instead.
func (*DeleteStudentClassResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_student_class_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteStudentClassResponse) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *DeleteStudentClassResponse) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

var File_coteacher_v1_student_class_proto protoreflect.FileDescriptor

var file_coteacher_v1_student_class_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x1f, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3f, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x50,
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x22, 0x55, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x32,
	0xd4, 0x03, 0x0a, 0x10, 0x43, 0x6f, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12,
	0x2c, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63,
	0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd2, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6e,
	0x6a, 0x69, 0x6b, 0x61, 0x77, 0x61, 0x67, 0x75, 0x63, 0x68, 0x69, 0x2f, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x58, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x18, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x6f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_coteacher_v1_student_class_proto_rawDescOnce sync.Once
	file_coteacher_v1_student_class_proto_rawDescData = file_coteacher_v1_student_class_proto_rawDesc
)

func file_coteacher_v1_student_class_proto_rawDescGZIP() []byte {
	file_coteacher_v1_student_class_proto_rawDescOnce.Do(func() {
		file_coteacher_v1_student_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_coteacher_v1_student_class_proto_rawDescData)
	})
	return file_coteacher_v1_student_class_proto_rawDescData
}

var file_coteacher_v1_student_class_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_coteacher_v1_student_class_proto_goTypes = []interface{}{
	(*CreateStudentClassRequest)(nil),       // 0: coteacher.v1.CreateStudentClassRequest
	(*CreateStudentClassResponse)(nil),      // 1: coteacher.v1.CreateStudentClassResponse
	(*GetStudentListByClassIDRequest)(nil),  // 2: coteacher.v1.GetStudentListByClassIDRequest
	(*GetStudentListByClassIDResponse)(nil), // 3: coteacher.v1.GetStudentListByClassIDResponse
	(*GetClassListByStudentIDRequest)(nil),  // 4: coteacher.v1.GetClassListByStudentIDRequest
	(*GetClassListByStudentIDResponse)(nil), // 5: coteacher.v1.GetClassListByStudentIDResponse
	(*DeleteStudentClassRequest)(nil),       // 6: coteacher.v1.DeleteStudentClassRequest
	(*DeleteStudentClassResponse)(nil),      // 7: coteacher.v1.DeleteStudentClassResponse
	(*User)(nil),                            // 8: coteacher.v1.User
	(*Class)(nil),                           // 9: coteacher.v1.Class
}
var file_coteacher_v1_student_class_proto_depIdxs = []int32{
	8, // 0: coteacher.v1.GetStudentListByClassIDResponse.students:type_name -> coteacher.v1.User
	9, // 1: coteacher.v1.GetClassListByStudentIDResponse.classes:type_name -> coteacher.v1.Class
	2, // 2: coteacher.v1.CoTeacherService.GetStudentListByClassID:input_type -> coteacher.v1.GetStudentListByClassIDRequest
	4, // 3: coteacher.v1.CoTeacherService.GetClassListByStudentID:input_type -> coteacher.v1.GetClassListByStudentIDRequest
	0, // 4: coteacher.v1.CoTeacherService.CreateStudentClass:input_type -> coteacher.v1.CreateStudentClassRequest
	6, // 5: coteacher.v1.CoTeacherService.DeleteStudentClass:input_type -> coteacher.v1.DeleteStudentClassRequest
	3, // 6: coteacher.v1.CoTeacherService.GetStudentListByClassID:output_type -> coteacher.v1.GetStudentListByClassIDResponse
	5, // 7: coteacher.v1.CoTeacherService.GetClassListByStudentID:output_type -> coteacher.v1.GetClassListByStudentIDResponse
	1, // 8: coteacher.v1.CoTeacherService.CreateStudentClass:output_type -> coteacher.v1.CreateStudentClassResponse
	7, // 9: coteacher.v1.CoTeacherService.DeleteStudentClass:output_type -> coteacher.v1.DeleteStudentClassResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coteacher_v1_student_class_proto_init() }
func file_coteacher_v1_student_class_proto_init() {
	if File_coteacher_v1_student_class_proto != nil {
		return
	}
	file_coteacher_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coteacher_v1_student_class_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentClassRequest); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentClassResponse); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentListByClassIDRequest); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentListByClassIDResponse); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassListByStudentIDRequest); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassListByStudentIDResponse); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentClassRequest); i {
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
		file_coteacher_v1_student_class_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentClassResponse); i {
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
			RawDescriptor: file_coteacher_v1_student_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coteacher_v1_student_class_proto_goTypes,
		DependencyIndexes: file_coteacher_v1_student_class_proto_depIdxs,
		MessageInfos:      file_coteacher_v1_student_class_proto_msgTypes,
	}.Build()
	File_coteacher_v1_student_class_proto = out.File
	file_coteacher_v1_student_class_proto_rawDesc = nil
	file_coteacher_v1_student_class_proto_goTypes = nil
	file_coteacher_v1_student_class_proto_depIdxs = nil
}

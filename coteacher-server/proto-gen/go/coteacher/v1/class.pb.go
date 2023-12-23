// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: coteacher/v1/class.proto

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

type CreateClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TeacherId string `protobuf:"bytes,2,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *CreateClassRequest) Reset() {
	*x = CreateClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassRequest) ProtoMessage() {}

func (x *CreateClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassRequest.ProtoReflect.Descriptor instead.
func (*CreateClassRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClassRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type CreateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class *Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *CreateClassResponse) Reset() {
	*x = CreateClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassResponse) ProtoMessage() {}

func (x *CreateClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassResponse.ProtoReflect.Descriptor instead.
func (*CreateClassResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClassResponse) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type GetClassByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClassByIDRequest) Reset() {
	*x = GetClassByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByIDRequest) ProtoMessage() {}

func (x *GetClassByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByIDRequest.ProtoReflect.Descriptor instead.
func (*GetClassByIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{2}
}

func (x *GetClassByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetClassByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class *Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *GetClassByIDResponse) Reset() {
	*x = GetClassByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByIDResponse) ProtoMessage() {}

func (x *GetClassByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByIDResponse.ProtoReflect.Descriptor instead.
func (*GetClassByIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{3}
}

func (x *GetClassByIDResponse) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type GetClassByTeacherIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *GetClassByTeacherIDRequest) Reset() {
	*x = GetClassByTeacherIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassByTeacherIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByTeacherIDRequest) ProtoMessage() {}

func (x *GetClassByTeacherIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByTeacherIDRequest.ProtoReflect.Descriptor instead.
func (*GetClassByTeacherIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{4}
}

func (x *GetClassByTeacherIDRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type GetClassByTeacherIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes []*Class `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
}

func (x *GetClassByTeacherIDResponse) Reset() {
	*x = GetClassByTeacherIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassByTeacherIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByTeacherIDResponse) ProtoMessage() {}

func (x *GetClassByTeacherIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByTeacherIDResponse.ProtoReflect.Descriptor instead.
func (*GetClassByTeacherIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{5}
}

func (x *GetClassByTeacherIDResponse) GetClasses() []*Class {
	if x != nil {
		return x.Classes
	}
	return nil
}

type UpdateClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeacherId string `protobuf:"bytes,3,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *UpdateClassRequest) Reset() {
	*x = UpdateClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClassRequest) ProtoMessage() {}

func (x *UpdateClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClassRequest.ProtoReflect.Descriptor instead.
func (*UpdateClassRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateClassRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateClassRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type UpdateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class *Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *UpdateClassResponse) Reset() {
	*x = UpdateClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClassResponse) ProtoMessage() {}

func (x *UpdateClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClassResponse.ProtoReflect.Descriptor instead.
func (*UpdateClassResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateClassResponse) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type DeleteClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteClassRequest) Reset() {
	*x = DeleteClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClassRequest) ProtoMessage() {}

func (x *DeleteClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClassRequest.ProtoReflect.Descriptor instead.
func (*DeleteClassRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteClassRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class *Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *DeleteClassResponse) Reset() {
	*x = DeleteClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClassResponse) ProtoMessage() {}

func (x *DeleteClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClassResponse.ProtoReflect.Descriptor instead.
func (*DeleteClassResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteClassResponse) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

var File_coteacher_v1_class_proto protoreflect.FileDescriptor

var file_coteacher_v1_class_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x40, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x07, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x40, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x32, 0xa5, 0x02, 0x0a, 0x0c, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x42, 0x79, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x63,
	0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xcb, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x69, 0x6e, 0x6a, 0x69, 0x6b, 0x61, 0x77, 0x61, 0x67, 0x75, 0x63, 0x68, 0x69, 0x2f,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0d, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coteacher_v1_class_proto_rawDescOnce sync.Once
	file_coteacher_v1_class_proto_rawDescData = file_coteacher_v1_class_proto_rawDesc
)

func file_coteacher_v1_class_proto_rawDescGZIP() []byte {
	file_coteacher_v1_class_proto_rawDescOnce.Do(func() {
		file_coteacher_v1_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_coteacher_v1_class_proto_rawDescData)
	})
	return file_coteacher_v1_class_proto_rawDescData
}

var file_coteacher_v1_class_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_coteacher_v1_class_proto_goTypes = []interface{}{
	(*CreateClassRequest)(nil),          // 0: coteacher.v1.CreateClassRequest
	(*CreateClassResponse)(nil),         // 1: coteacher.v1.CreateClassResponse
	(*GetClassByIDRequest)(nil),         // 2: coteacher.v1.GetClassByIDRequest
	(*GetClassByIDResponse)(nil),        // 3: coteacher.v1.GetClassByIDResponse
	(*GetClassByTeacherIDRequest)(nil),  // 4: coteacher.v1.GetClassByTeacherIDRequest
	(*GetClassByTeacherIDResponse)(nil), // 5: coteacher.v1.GetClassByTeacherIDResponse
	(*UpdateClassRequest)(nil),          // 6: coteacher.v1.UpdateClassRequest
	(*UpdateClassResponse)(nil),         // 7: coteacher.v1.UpdateClassResponse
	(*DeleteClassRequest)(nil),          // 8: coteacher.v1.DeleteClassRequest
	(*DeleteClassResponse)(nil),         // 9: coteacher.v1.DeleteClassResponse
	(*Class)(nil),                       // 10: coteacher.v1.Class
}
var file_coteacher_v1_class_proto_depIdxs = []int32{
	10, // 0: coteacher.v1.CreateClassResponse.class:type_name -> coteacher.v1.Class
	10, // 1: coteacher.v1.GetClassByIDResponse.class:type_name -> coteacher.v1.Class
	10, // 2: coteacher.v1.GetClassByTeacherIDResponse.classes:type_name -> coteacher.v1.Class
	10, // 3: coteacher.v1.UpdateClassResponse.class:type_name -> coteacher.v1.Class
	10, // 4: coteacher.v1.DeleteClassResponse.class:type_name -> coteacher.v1.Class
	2,  // 5: coteacher.v1.ClassService.GetClassByID:input_type -> coteacher.v1.GetClassByIDRequest
	4,  // 6: coteacher.v1.ClassService.GetClassByTeacherID:input_type -> coteacher.v1.GetClassByTeacherIDRequest
	0,  // 7: coteacher.v1.ClassService.CreateClass:input_type -> coteacher.v1.CreateClassRequest
	3,  // 8: coteacher.v1.ClassService.GetClassByID:output_type -> coteacher.v1.GetClassByIDResponse
	5,  // 9: coteacher.v1.ClassService.GetClassByTeacherID:output_type -> coteacher.v1.GetClassByTeacherIDResponse
	1,  // 10: coteacher.v1.ClassService.CreateClass:output_type -> coteacher.v1.CreateClassResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_coteacher_v1_class_proto_init() }
func file_coteacher_v1_class_proto_init() {
	if File_coteacher_v1_class_proto != nil {
		return
	}
	file_coteacher_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coteacher_v1_class_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassRequest); i {
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
		file_coteacher_v1_class_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassResponse); i {
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
		file_coteacher_v1_class_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassByIDRequest); i {
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
		file_coteacher_v1_class_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassByIDResponse); i {
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
		file_coteacher_v1_class_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassByTeacherIDRequest); i {
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
		file_coteacher_v1_class_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassByTeacherIDResponse); i {
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
		file_coteacher_v1_class_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClassRequest); i {
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
		file_coteacher_v1_class_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClassResponse); i {
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
		file_coteacher_v1_class_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClassRequest); i {
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
		file_coteacher_v1_class_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClassResponse); i {
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
			RawDescriptor: file_coteacher_v1_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coteacher_v1_class_proto_goTypes,
		DependencyIndexes: file_coteacher_v1_class_proto_depIdxs,
		MessageInfos:      file_coteacher_v1_class_proto_msgTypes,
	}.Build()
	File_coteacher_v1_class_proto = out.File
	file_coteacher_v1_class_proto_rawDesc = nil
	file_coteacher_v1_class_proto_goTypes = nil
	file_coteacher_v1_class_proto_depIdxs = nil
}

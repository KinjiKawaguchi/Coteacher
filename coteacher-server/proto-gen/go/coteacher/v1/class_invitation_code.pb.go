// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: coteacher/v1/class_invitation_code.proto

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

type CreateClassInvitationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId        string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	ExpirationDate string `protobuf:"bytes,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
}

func (x *CreateClassInvitationCodeRequest) Reset() {
	*x = CreateClassInvitationCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassInvitationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassInvitationCodeRequest) ProtoMessage() {}

func (x *CreateClassInvitationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassInvitationCodeRequest.ProtoReflect.Descriptor instead.
func (*CreateClassInvitationCodeRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClassInvitationCodeRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *CreateClassInvitationCodeRequest) GetExpirationDate() string {
	if x != nil {
		return x.ExpirationDate
	}
	return ""
}

type CreateClassInvitationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassInvitationCode *ClassInvitationCode `protobuf:"bytes,1,opt,name=class_invitation_code,json=classInvitationCode,proto3" json:"class_invitation_code,omitempty"`
}

func (x *CreateClassInvitationCodeResponse) Reset() {
	*x = CreateClassInvitationCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassInvitationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassInvitationCodeResponse) ProtoMessage() {}

func (x *CreateClassInvitationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassInvitationCodeResponse.ProtoReflect.Descriptor instead.
func (*CreateClassInvitationCodeResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClassInvitationCodeResponse) GetClassInvitationCode() *ClassInvitationCode {
	if x != nil {
		return x.ClassInvitationCode
	}
	return nil
}

type GetClassInvitationCodeByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClassInvitationCodeByIDRequest) Reset() {
	*x = GetClassInvitationCodeByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassInvitationCodeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassInvitationCodeByIDRequest) ProtoMessage() {}

func (x *GetClassInvitationCodeByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassInvitationCodeByIDRequest.ProtoReflect.Descriptor instead.
func (*GetClassInvitationCodeByIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{2}
}

func (x *GetClassInvitationCodeByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetClassInvitationCodeByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassInvitationCode *ClassInvitationCode `protobuf:"bytes,1,opt,name=class_invitation_code,json=classInvitationCode,proto3" json:"class_invitation_code,omitempty"`
}

func (x *GetClassInvitationCodeByIDResponse) Reset() {
	*x = GetClassInvitationCodeByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassInvitationCodeByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassInvitationCodeByIDResponse) ProtoMessage() {}

func (x *GetClassInvitationCodeByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassInvitationCodeByIDResponse.ProtoReflect.Descriptor instead.
func (*GetClassInvitationCodeByIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{3}
}

func (x *GetClassInvitationCodeByIDResponse) GetClassInvitationCode() *ClassInvitationCode {
	if x != nil {
		return x.ClassInvitationCode
	}
	return nil
}

type GetClassInvitationCodeListByClassIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *GetClassInvitationCodeListByClassIDRequest) Reset() {
	*x = GetClassInvitationCodeListByClassIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassInvitationCodeListByClassIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassInvitationCodeListByClassIDRequest) ProtoMessage() {}

func (x *GetClassInvitationCodeListByClassIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassInvitationCodeListByClassIDRequest.ProtoReflect.Descriptor instead.
func (*GetClassInvitationCodeListByClassIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{4}
}

func (x *GetClassInvitationCodeListByClassIDRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type GetClassInvitationCodeListByClassIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassInvitationCodeList []*ClassInvitationCode `protobuf:"bytes,1,rep,name=class_invitation_code_list,json=classInvitationCodeList,proto3" json:"class_invitation_code_list,omitempty"`
}

func (x *GetClassInvitationCodeListByClassIDResponse) Reset() {
	*x = GetClassInvitationCodeListByClassIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassInvitationCodeListByClassIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassInvitationCodeListByClassIDResponse) ProtoMessage() {}

func (x *GetClassInvitationCodeListByClassIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassInvitationCodeListByClassIDResponse.ProtoReflect.Descriptor instead.
func (*GetClassInvitationCodeListByClassIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{5}
}

func (x *GetClassInvitationCodeListByClassIDResponse) GetClassInvitationCodeList() []*ClassInvitationCode {
	if x != nil {
		return x.ClassInvitationCodeList
	}
	return nil
}

type UpdateClassInvitationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InvitationCode string `protobuf:"bytes,3,opt,name=invitation_code,json=invitationCode,proto3" json:"invitation_code,omitempty"`
	ExpirationDate string `protobuf:"bytes,4,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	IsActive       bool   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *UpdateClassInvitationCodeRequest) Reset() {
	*x = UpdateClassInvitationCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClassInvitationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClassInvitationCodeRequest) ProtoMessage() {}

func (x *UpdateClassInvitationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClassInvitationCodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateClassInvitationCodeRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateClassInvitationCodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateClassInvitationCodeRequest) GetInvitationCode() string {
	if x != nil {
		return x.InvitationCode
	}
	return ""
}

func (x *UpdateClassInvitationCodeRequest) GetExpirationDate() string {
	if x != nil {
		return x.ExpirationDate
	}
	return ""
}

func (x *UpdateClassInvitationCodeRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type UpdateClassInvitationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassInvitationCode *ClassInvitationCode `protobuf:"bytes,1,opt,name=class_invitation_code,json=classInvitationCode,proto3" json:"class_invitation_code,omitempty"`
}

func (x *UpdateClassInvitationCodeResponse) Reset() {
	*x = UpdateClassInvitationCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClassInvitationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClassInvitationCodeResponse) ProtoMessage() {}

func (x *UpdateClassInvitationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_class_invitation_code_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClassInvitationCodeResponse.ProtoReflect.Descriptor instead.
func (*UpdateClassInvitationCodeResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_class_invitation_code_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateClassInvitationCodeResponse) GetClassInvitationCode() *ClassInvitationCode {
	if x != nil {
		return x.ClassInvitationCode
	}
	return nil
}

var File_coteacher_v1_class_invitation_code_proto protoreflect.FileDescriptor

var file_coteacher_v1_class_invitation_code_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x7a,
	0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x33, 0x0a, 0x21, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x7b, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x47, 0x0a, 0x2a,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x2b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x1a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x17, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x7a, 0x0a, 0x21, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xb6, 0x04, 0x0a, 0x1a, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9a, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x38, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x7c, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2e, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7c, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd9,
	0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x18, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6e, 0x6a,
	0x69, 0x6b, 0x61, 0x77, 0x61, 0x67, 0x75, 0x63, 0x68, 0x69, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58,
	0x58, 0xaa, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x18, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x6f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_coteacher_v1_class_invitation_code_proto_rawDescOnce sync.Once
	file_coteacher_v1_class_invitation_code_proto_rawDescData = file_coteacher_v1_class_invitation_code_proto_rawDesc
)

func file_coteacher_v1_class_invitation_code_proto_rawDescGZIP() []byte {
	file_coteacher_v1_class_invitation_code_proto_rawDescOnce.Do(func() {
		file_coteacher_v1_class_invitation_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_coteacher_v1_class_invitation_code_proto_rawDescData)
	})
	return file_coteacher_v1_class_invitation_code_proto_rawDescData
}

var file_coteacher_v1_class_invitation_code_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_coteacher_v1_class_invitation_code_proto_goTypes = []interface{}{
	(*CreateClassInvitationCodeRequest)(nil),            // 0: coteacher.v1.CreateClassInvitationCodeRequest
	(*CreateClassInvitationCodeResponse)(nil),           // 1: coteacher.v1.CreateClassInvitationCodeResponse
	(*GetClassInvitationCodeByIDRequest)(nil),           // 2: coteacher.v1.GetClassInvitationCodeByIDRequest
	(*GetClassInvitationCodeByIDResponse)(nil),          // 3: coteacher.v1.GetClassInvitationCodeByIDResponse
	(*GetClassInvitationCodeListByClassIDRequest)(nil),  // 4: coteacher.v1.GetClassInvitationCodeListByClassIDRequest
	(*GetClassInvitationCodeListByClassIDResponse)(nil), // 5: coteacher.v1.GetClassInvitationCodeListByClassIDResponse
	(*UpdateClassInvitationCodeRequest)(nil),            // 6: coteacher.v1.UpdateClassInvitationCodeRequest
	(*UpdateClassInvitationCodeResponse)(nil),           // 7: coteacher.v1.UpdateClassInvitationCodeResponse
	(*ClassInvitationCode)(nil),                         // 8: coteacher.v1.ClassInvitationCode
}
var file_coteacher_v1_class_invitation_code_proto_depIdxs = []int32{
	8, // 0: coteacher.v1.CreateClassInvitationCodeResponse.class_invitation_code:type_name -> coteacher.v1.ClassInvitationCode
	8, // 1: coteacher.v1.GetClassInvitationCodeByIDResponse.class_invitation_code:type_name -> coteacher.v1.ClassInvitationCode
	8, // 2: coteacher.v1.GetClassInvitationCodeListByClassIDResponse.class_invitation_code_list:type_name -> coteacher.v1.ClassInvitationCode
	8, // 3: coteacher.v1.UpdateClassInvitationCodeResponse.class_invitation_code:type_name -> coteacher.v1.ClassInvitationCode
	2, // 4: coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeByID:input_type -> coteacher.v1.GetClassInvitationCodeByIDRequest
	4, // 5: coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeListByClassID:input_type -> coteacher.v1.GetClassInvitationCodeListByClassIDRequest
	0, // 6: coteacher.v1.ClassInvitationCodeService.CreateClassInvitationCode:input_type -> coteacher.v1.CreateClassInvitationCodeRequest
	6, // 7: coteacher.v1.ClassInvitationCodeService.UpdateClassInvitationCode:input_type -> coteacher.v1.UpdateClassInvitationCodeRequest
	3, // 8: coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeByID:output_type -> coteacher.v1.GetClassInvitationCodeByIDResponse
	5, // 9: coteacher.v1.ClassInvitationCodeService.GetClassInvitationCodeListByClassID:output_type -> coteacher.v1.GetClassInvitationCodeListByClassIDResponse
	1, // 10: coteacher.v1.ClassInvitationCodeService.CreateClassInvitationCode:output_type -> coteacher.v1.CreateClassInvitationCodeResponse
	7, // 11: coteacher.v1.ClassInvitationCodeService.UpdateClassInvitationCode:output_type -> coteacher.v1.UpdateClassInvitationCodeResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_coteacher_v1_class_invitation_code_proto_init() }
func file_coteacher_v1_class_invitation_code_proto_init() {
	if File_coteacher_v1_class_invitation_code_proto != nil {
		return
	}
	file_coteacher_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coteacher_v1_class_invitation_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassInvitationCodeRequest); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassInvitationCodeResponse); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassInvitationCodeByIDRequest); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassInvitationCodeByIDResponse); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassInvitationCodeListByClassIDRequest); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassInvitationCodeListByClassIDResponse); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClassInvitationCodeRequest); i {
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
		file_coteacher_v1_class_invitation_code_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClassInvitationCodeResponse); i {
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
			RawDescriptor: file_coteacher_v1_class_invitation_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coteacher_v1_class_invitation_code_proto_goTypes,
		DependencyIndexes: file_coteacher_v1_class_invitation_code_proto_depIdxs,
		MessageInfos:      file_coteacher_v1_class_invitation_code_proto_msgTypes,
	}.Build()
	File_coteacher_v1_class_invitation_code_proto = out.File
	file_coteacher_v1_class_invitation_code_proto_rawDesc = nil
	file_coteacher_v1_class_invitation_code_proto_goTypes = nil
	file_coteacher_v1_class_invitation_code_proto_depIdxs = nil
}

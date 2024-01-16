// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: coteacher/v1/form.proto

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

type CreateFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Form    *Form  `protobuf:"bytes,2,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *CreateFormRequest) Reset() {
	*x = CreateFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFormRequest) ProtoMessage() {}

func (x *CreateFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFormRequest.ProtoReflect.Descriptor instead.
func (*CreateFormRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFormRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *CreateFormRequest) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type CreateFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *CreateFormResponse) Reset() {
	*x = CreateFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFormResponse) ProtoMessage() {}

func (x *CreateFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFormResponse.ProtoReflect.Descriptor instead.
func (*CreateFormResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFormResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type GetFormByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormId string `protobuf:"bytes,1,opt,name=form_id,json=formId,proto3" json:"form_id,omitempty"`
}

func (x *GetFormByIDRequest) Reset() {
	*x = GetFormByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormByIDRequest) ProtoMessage() {}

func (x *GetFormByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormByIDRequest.ProtoReflect.Descriptor instead.
func (*GetFormByIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{2}
}

func (x *GetFormByIDRequest) GetFormId() string {
	if x != nil {
		return x.FormId
	}
	return ""
}

type GetFormByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *GetFormByIDResponse) Reset() {
	*x = GetFormByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormByIDResponse) ProtoMessage() {}

func (x *GetFormByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormByIDResponse.ProtoReflect.Descriptor instead.
func (*GetFormByIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{3}
}

func (x *GetFormByIDResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type GetFormListByClassIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *GetFormListByClassIDRequest) Reset() {
	*x = GetFormListByClassIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormListByClassIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormListByClassIDRequest) ProtoMessage() {}

func (x *GetFormListByClassIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormListByClassIDRequest.ProtoReflect.Descriptor instead.
func (*GetFormListByClassIDRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{4}
}

func (x *GetFormListByClassIDRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

type GetFormListByClassIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forms []*Form `protobuf:"bytes,1,rep,name=forms,proto3" json:"forms,omitempty"`
}

func (x *GetFormListByClassIDResponse) Reset() {
	*x = GetFormListByClassIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormListByClassIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormListByClassIDResponse) ProtoMessage() {}

func (x *GetFormListByClassIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormListByClassIDResponse.ProtoReflect.Descriptor instead.
func (*GetFormListByClassIDResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{5}
}

func (x *GetFormListByClassIDResponse) GetForms() []*Form {
	if x != nil {
		return x.Forms
	}
	return nil
}

type UpdateFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Form    *Form  `protobuf:"bytes,2,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *UpdateFormRequest) Reset() {
	*x = UpdateFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFormRequest) ProtoMessage() {}

func (x *UpdateFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFormRequest.ProtoReflect.Descriptor instead.
func (*UpdateFormRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFormRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *UpdateFormRequest) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type UpdateFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *UpdateFormResponse) Reset() {
	*x = UpdateFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFormResponse) ProtoMessage() {}

func (x *UpdateFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFormResponse.ProtoReflect.Descriptor instead.
func (*UpdateFormResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateFormResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type DeleteFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	FormId  string `protobuf:"bytes,2,opt,name=form_id,json=formId,proto3" json:"form_id,omitempty"`
}

func (x *DeleteFormRequest) Reset() {
	*x = DeleteFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFormRequest) ProtoMessage() {}

func (x *DeleteFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFormRequest.ProtoReflect.Descriptor instead.
func (*DeleteFormRequest) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFormRequest) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *DeleteFormRequest) GetFormId() string {
	if x != nil {
		return x.FormId
	}
	return ""
}

type DeleteFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Form `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *DeleteFormResponse) Reset() {
	*x = DeleteFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coteacher_v1_form_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFormResponse) ProtoMessage() {}

func (x *DeleteFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coteacher_v1_form_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFormResponse.ProtoReflect.Descriptor instead.
func (*DeleteFormResponse) Descriptor() ([]byte, []int) {
	return file_coteacher_v1_form_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteFormResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

var File_coteacher_v1_form_proto protoreflect.FileDescriptor

var file_coteacher_v1_form_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x3c, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x2d, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x38, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x22, 0x56, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52,
	0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66,
	0x6f, 0x72, 0x6d, 0x22, 0x47, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x32, 0xc3, 0x03, 0x0a, 0x0b, 0x46,
	0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72,
	0x6d, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x72, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1f, 0x2e, 0x63,
	0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1f, 0x2e,
	0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xb9, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x46, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b,
	0x69, 0x6e, 0x6a, 0x69, 0x4b, 0x61, 0x77, 0x61, 0x67, 0x75, 0x63, 0x68, 0x69, 0x2f, 0x43, 0x6f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0c, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x18, 0x43, 0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43,
	0x6f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coteacher_v1_form_proto_rawDescOnce sync.Once
	file_coteacher_v1_form_proto_rawDescData = file_coteacher_v1_form_proto_rawDesc
)

func file_coteacher_v1_form_proto_rawDescGZIP() []byte {
	file_coteacher_v1_form_proto_rawDescOnce.Do(func() {
		file_coteacher_v1_form_proto_rawDescData = protoimpl.X.CompressGZIP(file_coteacher_v1_form_proto_rawDescData)
	})
	return file_coteacher_v1_form_proto_rawDescData
}

var file_coteacher_v1_form_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_coteacher_v1_form_proto_goTypes = []interface{}{
	(*CreateFormRequest)(nil),            // 0: coteacher.v1.CreateFormRequest
	(*CreateFormResponse)(nil),           // 1: coteacher.v1.CreateFormResponse
	(*GetFormByIDRequest)(nil),           // 2: coteacher.v1.GetFormByIDRequest
	(*GetFormByIDResponse)(nil),          // 3: coteacher.v1.GetFormByIDResponse
	(*GetFormListByClassIDRequest)(nil),  // 4: coteacher.v1.GetFormListByClassIDRequest
	(*GetFormListByClassIDResponse)(nil), // 5: coteacher.v1.GetFormListByClassIDResponse
	(*UpdateFormRequest)(nil),            // 6: coteacher.v1.UpdateFormRequest
	(*UpdateFormResponse)(nil),           // 7: coteacher.v1.UpdateFormResponse
	(*DeleteFormRequest)(nil),            // 8: coteacher.v1.DeleteFormRequest
	(*DeleteFormResponse)(nil),           // 9: coteacher.v1.DeleteFormResponse
	(*Form)(nil),                         // 10: coteacher.v1.Form
}
var file_coteacher_v1_form_proto_depIdxs = []int32{
	10, // 0: coteacher.v1.CreateFormRequest.form:type_name -> coteacher.v1.Form
	10, // 1: coteacher.v1.CreateFormResponse.form:type_name -> coteacher.v1.Form
	10, // 2: coteacher.v1.GetFormByIDResponse.form:type_name -> coteacher.v1.Form
	10, // 3: coteacher.v1.GetFormListByClassIDResponse.forms:type_name -> coteacher.v1.Form
	10, // 4: coteacher.v1.UpdateFormRequest.form:type_name -> coteacher.v1.Form
	10, // 5: coteacher.v1.UpdateFormResponse.form:type_name -> coteacher.v1.Form
	10, // 6: coteacher.v1.DeleteFormResponse.form:type_name -> coteacher.v1.Form
	0,  // 7: coteacher.v1.FormService.CreateForm:input_type -> coteacher.v1.CreateFormRequest
	2,  // 8: coteacher.v1.FormService.GetFormByID:input_type -> coteacher.v1.GetFormByIDRequest
	4,  // 9: coteacher.v1.FormService.GetFormListByClassID:input_type -> coteacher.v1.GetFormListByClassIDRequest
	6,  // 10: coteacher.v1.FormService.UpdateForm:input_type -> coteacher.v1.UpdateFormRequest
	8,  // 11: coteacher.v1.FormService.DeleteForm:input_type -> coteacher.v1.DeleteFormRequest
	1,  // 12: coteacher.v1.FormService.CreateForm:output_type -> coteacher.v1.CreateFormResponse
	3,  // 13: coteacher.v1.FormService.GetFormByID:output_type -> coteacher.v1.GetFormByIDResponse
	5,  // 14: coteacher.v1.FormService.GetFormListByClassID:output_type -> coteacher.v1.GetFormListByClassIDResponse
	7,  // 15: coteacher.v1.FormService.UpdateForm:output_type -> coteacher.v1.UpdateFormResponse
	9,  // 16: coteacher.v1.FormService.DeleteForm:output_type -> coteacher.v1.DeleteFormResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_coteacher_v1_form_proto_init() }
func file_coteacher_v1_form_proto_init() {
	if File_coteacher_v1_form_proto != nil {
		return
	}
	file_coteacher_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coteacher_v1_form_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFormRequest); i {
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
		file_coteacher_v1_form_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFormResponse); i {
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
		file_coteacher_v1_form_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormByIDRequest); i {
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
		file_coteacher_v1_form_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormByIDResponse); i {
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
		file_coteacher_v1_form_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormListByClassIDRequest); i {
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
		file_coteacher_v1_form_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormListByClassIDResponse); i {
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
		file_coteacher_v1_form_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFormRequest); i {
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
		file_coteacher_v1_form_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFormResponse); i {
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
		file_coteacher_v1_form_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFormRequest); i {
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
		file_coteacher_v1_form_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFormResponse); i {
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
			RawDescriptor: file_coteacher_v1_form_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coteacher_v1_form_proto_goTypes,
		DependencyIndexes: file_coteacher_v1_form_proto_depIdxs,
		MessageInfos:      file_coteacher_v1_form_proto_msgTypes,
	}.Build()
	File_coteacher_v1_form_proto = out.File
	file_coteacher_v1_form_proto_rawDesc = nil
	file_coteacher_v1_form_proto_goTypes = nil
	file_coteacher_v1_form_proto_depIdxs = nil
}

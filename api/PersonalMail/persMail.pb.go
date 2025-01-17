// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: persMail.proto

package PersonalMail

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreatePersMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string       `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Requestbody *RequestBody `protobuf:"bytes,2,opt,name=requestbody,proto3" json:"requestbody,omitempty"`
}

func (x *CreatePersMailRequest) Reset() {
	*x = CreatePersMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersMailRequest) ProtoMessage() {}

func (x *CreatePersMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersMailRequest.ProtoReflect.Descriptor instead.
func (*CreatePersMailRequest) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePersMailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreatePersMailRequest) GetRequestbody() *RequestBody {
	if x != nil {
		return x.Requestbody
	}
	return nil
}

type CreatePersMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
}

func (x *CreatePersMailResponse) Reset() {
	*x = CreatePersMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersMailResponse) ProtoMessage() {}

func (x *CreatePersMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersMailResponse.ProtoReflect.Descriptor instead.
func (*CreatePersMailResponse) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersMailResponse) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

type GetPersMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
}

func (x *GetPersMailRequest) Reset() {
	*x = GetPersMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersMailRequest) ProtoMessage() {}

func (x *GetPersMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersMailRequest.ProtoReflect.Descriptor instead.
func (*GetPersMailRequest) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersMailRequest) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

type GetPersMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetPersMailResponse) Reset() {
	*x = GetPersMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersMailResponse) ProtoMessage() {}

func (x *GetPersMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersMailResponse.ProtoReflect.Descriptor instead.
func (*GetPersMailResponse) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersMailResponse) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

func (x *GetPersMailResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdatePersMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId    uint64       `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
	Email       string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Requestbody *RequestBody `protobuf:"bytes,3,opt,name=requestbody,proto3" json:"requestbody,omitempty"`
}

func (x *UpdatePersMailRequest) Reset() {
	*x = UpdatePersMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersMailRequest) ProtoMessage() {}

func (x *UpdatePersMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersMailRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersMailRequest) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePersMailRequest) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

func (x *UpdatePersMailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdatePersMailRequest) GetRequestbody() *RequestBody {
	if x != nil {
		return x.Requestbody
	}
	return nil
}

type UpdatePersMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
}

func (x *UpdatePersMailResponse) Reset() {
	*x = UpdatePersMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersMailResponse) ProtoMessage() {}

func (x *UpdatePersMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersMailResponse.ProtoReflect.Descriptor instead.
func (*UpdatePersMailResponse) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePersMailResponse) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

type DeletePersMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
}

func (x *DeletePersMailRequest) Reset() {
	*x = DeletePersMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersMailRequest) ProtoMessage() {}

func (x *DeletePersMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersMailRequest.ProtoReflect.Descriptor instead.
func (*DeletePersMailRequest) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePersMailRequest) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

type DeletePersMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
}

func (x *DeletePersMailResponse) Reset() {
	*x = DeletePersMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersMailResponse) ProtoMessage() {}

func (x *DeletePersMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersMailResponse.ProtoReflect.Descriptor instead.
func (*DeletePersMailResponse) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePersMailResponse) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

type GetPersMailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPersMailsRequest) Reset() {
	*x = GetPersMailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersMailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersMailsRequest) ProtoMessage() {}

func (x *GetPersMailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersMailsRequest.ProtoReflect.Descriptor instead.
func (*GetPersMailsRequest) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{8}
}

type GetPersMailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persemails []*PersMail `protobuf:"bytes,1,rep,name=persemails,proto3" json:"persemails,omitempty"`
}

func (x *GetPersMailsResponse) Reset() {
	*x = GetPersMailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersMailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersMailsResponse) ProtoMessage() {}

func (x *GetPersMailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersMailsResponse.ProtoReflect.Descriptor instead.
func (*GetPersMailsResponse) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{9}
}

func (x *GetPersMailsResponse) GetPersemails() []*PersMail {
	if x != nil {
		return x.Persemails
	}
	return nil
}

type PersMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersEmId uint64 `protobuf:"varint,1,opt,name=pers_em_id,json=persEmId,proto3" json:"pers_em_id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *PersMail) Reset() {
	*x = PersMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersMail) ProtoMessage() {}

func (x *PersMail) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersMail.ProtoReflect.Descriptor instead.
func (*PersMail) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{10}
}

func (x *PersMail) GetPersEmId() uint64 {
	if x != nil {
		return x.PersEmId
	}
	return 0
}

func (x *PersMail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persMail_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_persMail_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBody.ProtoReflect.Descriptor instead.
func (*RequestBody) Descriptor() ([]byte, []int) {
	return file_persMail_proto_rawDescGZIP(), []int{11}
}

func (x *RequestBody) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_persMail_proto protoreflect.FileDescriptor

var file_persMail_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x36, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x45, 0x6d, 0x49, 0x64,
	0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x45, 0x6d, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x45, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x88, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x73, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x65, 0x72, 0x73, 0x45, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69,
	0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x36, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x45, 0x6d,
	0x49, 0x64, 0x22, 0x35, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x45, 0x6d, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x45, 0x6d, 0x49,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d,
	0x61, 0x69, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x73, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3e, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x73,
	0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x45, 0x6d,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x23, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xfe, 0x04,
	0x0a, 0x0f, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61,
	0x69, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x0b, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x74, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x73, 0x5f,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x18, 0x2f, 0x70, 0x65, 0x72, 0x73,
	0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x21, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x69,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x12, 0x0b, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x12,
	0x5a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_persMail_proto_rawDescOnce sync.Once
	file_persMail_proto_rawDescData = file_persMail_proto_rawDesc
)

func file_persMail_proto_rawDescGZIP() []byte {
	file_persMail_proto_rawDescOnce.Do(func() {
		file_persMail_proto_rawDescData = protoimpl.X.CompressGZIP(file_persMail_proto_rawDescData)
	})
	return file_persMail_proto_rawDescData
}

var file_persMail_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_persMail_proto_goTypes = []interface{}{
	(*CreatePersMailRequest)(nil),  // 0: PersonalMail.CreatePersMailRequest
	(*CreatePersMailResponse)(nil), // 1: PersonalMail.CreatePersMailResponse
	(*GetPersMailRequest)(nil),     // 2: PersonalMail.GetPersMailRequest
	(*GetPersMailResponse)(nil),    // 3: PersonalMail.GetPersMailResponse
	(*UpdatePersMailRequest)(nil),  // 4: PersonalMail.UpdatePersMailRequest
	(*UpdatePersMailResponse)(nil), // 5: PersonalMail.UpdatePersMailResponse
	(*DeletePersMailRequest)(nil),  // 6: PersonalMail.DeletePersMailRequest
	(*DeletePersMailResponse)(nil), // 7: PersonalMail.DeletePersMailResponse
	(*GetPersMailsRequest)(nil),    // 8: PersonalMail.GetPersMailsRequest
	(*GetPersMailsResponse)(nil),   // 9: PersonalMail.GetPersMailsResponse
	(*PersMail)(nil),               // 10: PersonalMail.PersMail
	(*RequestBody)(nil),            // 11: PersonalMail.RequestBody
}
var file_persMail_proto_depIdxs = []int32{
	11, // 0: PersonalMail.CreatePersMailRequest.requestbody:type_name -> PersonalMail.RequestBody
	11, // 1: PersonalMail.UpdatePersMailRequest.requestbody:type_name -> PersonalMail.RequestBody
	10, // 2: PersonalMail.GetPersMailsResponse.persemails:type_name -> PersonalMail.PersMail
	0,  // 3: PersonalMail.PersMailService.CreatePersMail:input_type -> PersonalMail.CreatePersMailRequest
	2,  // 4: PersonalMail.PersMailService.GetPersMail:input_type -> PersonalMail.GetPersMailRequest
	4,  // 5: PersonalMail.PersMailService.UpdatePersMail:input_type -> PersonalMail.UpdatePersMailRequest
	6,  // 6: PersonalMail.PersMailService.DeletePersMail:input_type -> PersonalMail.DeletePersMailRequest
	8,  // 7: PersonalMail.PersMailService.GetPersMails:input_type -> PersonalMail.GetPersMailsRequest
	1,  // 8: PersonalMail.PersMailService.CreatePersMail:output_type -> PersonalMail.CreatePersMailResponse
	3,  // 9: PersonalMail.PersMailService.GetPersMail:output_type -> PersonalMail.GetPersMailResponse
	5,  // 10: PersonalMail.PersMailService.UpdatePersMail:output_type -> PersonalMail.UpdatePersMailResponse
	7,  // 11: PersonalMail.PersMailService.DeletePersMail:output_type -> PersonalMail.DeletePersMailResponse
	9,  // 12: PersonalMail.PersMailService.GetPersMails:output_type -> PersonalMail.GetPersMailsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_persMail_proto_init() }
func file_persMail_proto_init() {
	if File_persMail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_persMail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersMailRequest); i {
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
		file_persMail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersMailResponse); i {
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
		file_persMail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersMailRequest); i {
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
		file_persMail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersMailResponse); i {
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
		file_persMail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersMailRequest); i {
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
		file_persMail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersMailResponse); i {
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
		file_persMail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersMailRequest); i {
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
		file_persMail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersMailResponse); i {
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
		file_persMail_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersMailsRequest); i {
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
		file_persMail_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersMailsResponse); i {
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
		file_persMail_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersMail); i {
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
		file_persMail_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBody); i {
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
			RawDescriptor: file_persMail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_persMail_proto_goTypes,
		DependencyIndexes: file_persMail_proto_depIdxs,
		MessageInfos:      file_persMail_proto_msgTypes,
	}.Build()
	File_persMail_proto = out.File
	file_persMail_proto_rawDesc = nil
	file_persMail_proto_goTypes = nil
	file_persMail_proto_depIdxs = nil
}

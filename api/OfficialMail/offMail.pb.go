// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: offMail.proto

package OfficialMail

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

type CreateOffMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string       `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Requestbody *RequestBody `protobuf:"bytes,2,opt,name=requestbody,proto3" json:"requestbody,omitempty"`
}

func (x *CreateOffMailRequest) Reset() {
	*x = CreateOffMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOffMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOffMailRequest) ProtoMessage() {}

func (x *CreateOffMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOffMailRequest.ProtoReflect.Descriptor instead.
func (*CreateOffMailRequest) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOffMailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateOffMailRequest) GetRequestbody() *RequestBody {
	if x != nil {
		return x.Requestbody
	}
	return nil
}

type CreateOffMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
}

func (x *CreateOffMailResponse) Reset() {
	*x = CreateOffMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOffMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOffMailResponse) ProtoMessage() {}

func (x *CreateOffMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOffMailResponse.ProtoReflect.Descriptor instead.
func (*CreateOffMailResponse) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOffMailResponse) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

type GetOffMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
}

func (x *GetOffMailRequest) Reset() {
	*x = GetOffMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOffMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOffMailRequest) ProtoMessage() {}

func (x *GetOffMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOffMailRequest.ProtoReflect.Descriptor instead.
func (*GetOffMailRequest) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{2}
}

func (x *GetOffMailRequest) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

type GetOffMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetOffMailResponse) Reset() {
	*x = GetOffMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOffMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOffMailResponse) ProtoMessage() {}

func (x *GetOffMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOffMailResponse.ProtoReflect.Descriptor instead.
func (*GetOffMailResponse) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{3}
}

func (x *GetOffMailResponse) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

func (x *GetOffMailResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateOffMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId     uint64       `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
	Email       string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Requestbody *RequestBody `protobuf:"bytes,3,opt,name=requestbody,proto3" json:"requestbody,omitempty"`
}

func (x *UpdateOffMailRequest) Reset() {
	*x = UpdateOffMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOffMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOffMailRequest) ProtoMessage() {}

func (x *UpdateOffMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOffMailRequest.ProtoReflect.Descriptor instead.
func (*UpdateOffMailRequest) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOffMailRequest) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

func (x *UpdateOffMailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateOffMailRequest) GetRequestbody() *RequestBody {
	if x != nil {
		return x.Requestbody
	}
	return nil
}

type UpdateOffMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
}

func (x *UpdateOffMailResponse) Reset() {
	*x = UpdateOffMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOffMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOffMailResponse) ProtoMessage() {}

func (x *UpdateOffMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOffMailResponse.ProtoReflect.Descriptor instead.
func (*UpdateOffMailResponse) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOffMailResponse) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

type DeleteOffMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
}

func (x *DeleteOffMailRequest) Reset() {
	*x = DeleteOffMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOffMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOffMailRequest) ProtoMessage() {}

func (x *DeleteOffMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOffMailRequest.ProtoReflect.Descriptor instead.
func (*DeleteOffMailRequest) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteOffMailRequest) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

type DeleteOffMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
}

func (x *DeleteOffMailResponse) Reset() {
	*x = DeleteOffMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOffMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOffMailResponse) ProtoMessage() {}

func (x *DeleteOffMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOffMailResponse.ProtoReflect.Descriptor instead.
func (*DeleteOffMailResponse) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteOffMailResponse) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

type GetOffMailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOffMailsRequest) Reset() {
	*x = GetOffMailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOffMailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOffMailsRequest) ProtoMessage() {}

func (x *GetOffMailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOffMailsRequest.ProtoReflect.Descriptor instead.
func (*GetOffMailsRequest) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{8}
}

type GetOffMailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offemails []*OffMail `protobuf:"bytes,1,rep,name=offemails,proto3" json:"offemails,omitempty"`
}

func (x *GetOffMailsResponse) Reset() {
	*x = GetOffMailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOffMailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOffMailsResponse) ProtoMessage() {}

func (x *GetOffMailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOffMailsResponse.ProtoReflect.Descriptor instead.
func (*GetOffMailsResponse) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{9}
}

func (x *GetOffMailsResponse) GetOffemails() []*OffMail {
	if x != nil {
		return x.Offemails
	}
	return nil
}

type OffMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffEmId uint64 `protobuf:"varint,1,opt,name=off_em_id,json=offEmId,proto3" json:"off_em_id,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *OffMail) Reset() {
	*x = OffMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffMail) ProtoMessage() {}

func (x *OffMail) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffMail.ProtoReflect.Descriptor instead.
func (*OffMail) Descriptor() ([]byte, []int) {
	return file_offMail_proto_rawDescGZIP(), []int{10}
}

func (x *OffMail) GetOffEmId() uint64 {
	if x != nil {
		return x.OffEmId
	}
	return 0
}

func (x *OffMail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offMail_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_offMail_proto_msgTypes[11]
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
	return file_offMail_proto_rawDescGZIP(), []int{11}
}

func (x *RequestBody) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_offMail_proto protoreflect.FileDescriptor

var file_offMail_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x33, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x45, 0x6d, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x45, 0x6d, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x45, 0x6d, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x6f, 0x66, 0x66, 0x45, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d,
	0x61, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x33, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x45, 0x6d, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x09, 0x6f, 0x66, 0x66,
	0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x66,
	0x66, 0x45, 0x6d, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x45, 0x6d, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x4a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x09, 0x6f, 0x66, 0x66, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3b, 0x0a, 0x07,
	0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x5f, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x45,
	0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x23, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xe6,
	0x04, 0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x79, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x22, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69,
	0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x3a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x0a, 0x2f, 0x6f, 0x66, 0x66, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6f, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66,
	0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x6f, 0x66, 0x66, 0x5f, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x2f, 0x7b, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x85, 0x01,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x22, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61,
	0x69, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x3a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x16, 0x2f,
	0x6f, 0x66, 0x66, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x6f, 0x66, 0x66, 0x5f, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x78, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x6f, 0x66, 0x66, 0x5f, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x6f, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x66, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20,
	0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x6f, 0x66,
	0x66, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_offMail_proto_rawDescOnce sync.Once
	file_offMail_proto_rawDescData = file_offMail_proto_rawDesc
)

func file_offMail_proto_rawDescGZIP() []byte {
	file_offMail_proto_rawDescOnce.Do(func() {
		file_offMail_proto_rawDescData = protoimpl.X.CompressGZIP(file_offMail_proto_rawDescData)
	})
	return file_offMail_proto_rawDescData
}

var file_offMail_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_offMail_proto_goTypes = []interface{}{
	(*CreateOffMailRequest)(nil),  // 0: OfficialMail.CreateOffMailRequest
	(*CreateOffMailResponse)(nil), // 1: OfficialMail.CreateOffMailResponse
	(*GetOffMailRequest)(nil),     // 2: OfficialMail.GetOffMailRequest
	(*GetOffMailResponse)(nil),    // 3: OfficialMail.GetOffMailResponse
	(*UpdateOffMailRequest)(nil),  // 4: OfficialMail.UpdateOffMailRequest
	(*UpdateOffMailResponse)(nil), // 5: OfficialMail.UpdateOffMailResponse
	(*DeleteOffMailRequest)(nil),  // 6: OfficialMail.DeleteOffMailRequest
	(*DeleteOffMailResponse)(nil), // 7: OfficialMail.DeleteOffMailResponse
	(*GetOffMailsRequest)(nil),    // 8: OfficialMail.GetOffMailsRequest
	(*GetOffMailsResponse)(nil),   // 9: OfficialMail.GetOffMailsResponse
	(*OffMail)(nil),               // 10: OfficialMail.OffMail
	(*RequestBody)(nil),           // 11: OfficialMail.RequestBody
}
var file_offMail_proto_depIdxs = []int32{
	11, // 0: OfficialMail.CreateOffMailRequest.requestbody:type_name -> OfficialMail.RequestBody
	11, // 1: OfficialMail.UpdateOffMailRequest.requestbody:type_name -> OfficialMail.RequestBody
	10, // 2: OfficialMail.GetOffMailsResponse.offemails:type_name -> OfficialMail.OffMail
	0,  // 3: OfficialMail.OffMailService.CreateOffMail:input_type -> OfficialMail.CreateOffMailRequest
	2,  // 4: OfficialMail.OffMailService.GetOffMail:input_type -> OfficialMail.GetOffMailRequest
	4,  // 5: OfficialMail.OffMailService.UpdateOffMail:input_type -> OfficialMail.UpdateOffMailRequest
	6,  // 6: OfficialMail.OffMailService.DeleteOffMail:input_type -> OfficialMail.DeleteOffMailRequest
	8,  // 7: OfficialMail.OffMailService.GetOffMails:input_type -> OfficialMail.GetOffMailsRequest
	1,  // 8: OfficialMail.OffMailService.CreateOffMail:output_type -> OfficialMail.CreateOffMailResponse
	3,  // 9: OfficialMail.OffMailService.GetOffMail:output_type -> OfficialMail.GetOffMailResponse
	5,  // 10: OfficialMail.OffMailService.UpdateOffMail:output_type -> OfficialMail.UpdateOffMailResponse
	7,  // 11: OfficialMail.OffMailService.DeleteOffMail:output_type -> OfficialMail.DeleteOffMailResponse
	9,  // 12: OfficialMail.OffMailService.GetOffMails:output_type -> OfficialMail.GetOffMailsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_offMail_proto_init() }
func file_offMail_proto_init() {
	if File_offMail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_offMail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOffMailRequest); i {
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
		file_offMail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOffMailResponse); i {
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
		file_offMail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOffMailRequest); i {
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
		file_offMail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOffMailResponse); i {
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
		file_offMail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOffMailRequest); i {
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
		file_offMail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOffMailResponse); i {
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
		file_offMail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOffMailRequest); i {
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
		file_offMail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOffMailResponse); i {
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
		file_offMail_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOffMailsRequest); i {
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
		file_offMail_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOffMailsResponse); i {
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
		file_offMail_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffMail); i {
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
		file_offMail_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_offMail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_offMail_proto_goTypes,
		DependencyIndexes: file_offMail_proto_depIdxs,
		MessageInfos:      file_offMail_proto_msgTypes,
	}.Build()
	File_offMail_proto = out.File
	file_offMail_proto_rawDesc = nil
	file_offMail_proto_goTypes = nil
	file_offMail_proto_depIdxs = nil
}

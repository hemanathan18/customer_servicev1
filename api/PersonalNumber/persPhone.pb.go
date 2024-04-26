// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: persPhone.proto

package PersonalNumber

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

type CreatePersPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode int32        `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Number      int64        `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Requestbody *RequestBody `protobuf:"bytes,3,opt,name=requestbody,proto3" json:"requestbody,omitempty"`
}

func (x *CreatePersPhoneNumberRequest) Reset() {
	*x = CreatePersPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersPhoneNumberRequest) ProtoMessage() {}

func (x *CreatePersPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*CreatePersPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePersPhoneNumberRequest) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *CreatePersPhoneNumberRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *CreatePersPhoneNumberRequest) GetRequestbody() *RequestBody {
	if x != nil {
		return x.Requestbody
	}
	return nil
}

type CreatePersPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
}

func (x *CreatePersPhoneNumberResponse) Reset() {
	*x = CreatePersPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersPhoneNumberResponse) ProtoMessage() {}

func (x *CreatePersPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*CreatePersPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersPhoneNumberResponse) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

type GetPersPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
}

func (x *GetPersPhoneNumberRequest) Reset() {
	*x = GetPersPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersPhoneNumberRequest) ProtoMessage() {}

func (x *GetPersPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*GetPersPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersPhoneNumberRequest) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

type GetPersPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId    uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
	CountryCode int32  `protobuf:"varint,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Number      int64  `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetPersPhoneNumberResponse) Reset() {
	*x = GetPersPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersPhoneNumberResponse) ProtoMessage() {}

func (x *GetPersPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*GetPersPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersPhoneNumberResponse) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

func (x *GetPersPhoneNumberResponse) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *GetPersPhoneNumberResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type UpdatePersPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId    uint64       `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
	CountryCode int32        `protobuf:"varint,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Number      int64        `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Requestbody *RequestBody `protobuf:"bytes,4,opt,name=requestbody,proto3" json:"requestbody,omitempty"`
}

func (x *UpdatePersPhoneNumberRequest) Reset() {
	*x = UpdatePersPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersPhoneNumberRequest) ProtoMessage() {}

func (x *UpdatePersPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePersPhoneNumberRequest) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

func (x *UpdatePersPhoneNumberRequest) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *UpdatePersPhoneNumberRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *UpdatePersPhoneNumberRequest) GetRequestbody() *RequestBody {
	if x != nil {
		return x.Requestbody
	}
	return nil
}

type UpdatePersPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
}

func (x *UpdatePersPhoneNumberResponse) Reset() {
	*x = UpdatePersPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersPhoneNumberResponse) ProtoMessage() {}

func (x *UpdatePersPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*UpdatePersPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePersPhoneNumberResponse) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

type DeletePersPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
}

func (x *DeletePersPhoneNumberRequest) Reset() {
	*x = DeletePersPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersPhoneNumberRequest) ProtoMessage() {}

func (x *DeletePersPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*DeletePersPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePersPhoneNumberRequest) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

type DeletePersPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
}

func (x *DeletePersPhoneNumberResponse) Reset() {
	*x = DeletePersPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersPhoneNumberResponse) ProtoMessage() {}

func (x *DeletePersPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*DeletePersPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePersPhoneNumberResponse) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

type GetPersPhoneNumbersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPersPhoneNumbersRequest) Reset() {
	*x = GetPersPhoneNumbersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersPhoneNumbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersPhoneNumbersRequest) ProtoMessage() {}

func (x *GetPersPhoneNumbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersPhoneNumbersRequest.ProtoReflect.Descriptor instead.
func (*GetPersPhoneNumbersRequest) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{8}
}

type GetPersPhoneNumbersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persphonenumbers []*PersPhoneNumber `protobuf:"bytes,1,rep,name=persphonenumbers,proto3" json:"persphonenumbers,omitempty"`
}

func (x *GetPersPhoneNumbersResponse) Reset() {
	*x = GetPersPhoneNumbersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersPhoneNumbersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersPhoneNumbersResponse) ProtoMessage() {}

func (x *GetPersPhoneNumbersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersPhoneNumbersResponse.ProtoReflect.Descriptor instead.
func (*GetPersPhoneNumbersResponse) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{9}
}

func (x *GetPersPhoneNumbersResponse) GetPersphonenumbers() []*PersPhoneNumber {
	if x != nil {
		return x.Persphonenumbers
	}
	return nil
}

type PersPhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersPhId    uint64 `protobuf:"varint,1,opt,name=pers_ph_id,json=persPhId,proto3" json:"pers_ph_id,omitempty"`
	CountryCode int32  `protobuf:"varint,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Number      int64  `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PersPhoneNumber) Reset() {
	*x = PersPhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersPhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersPhoneNumber) ProtoMessage() {}

func (x *PersPhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersPhoneNumber.ProtoReflect.Descriptor instead.
func (*PersPhoneNumber) Descriptor() ([]byte, []int) {
	return file_persPhone_proto_rawDescGZIP(), []int{10}
}

func (x *PersPhoneNumber) GetPersPhId() uint64 {
	if x != nil {
		return x.PersPhId
	}
	return 0
}

func (x *PersPhoneNumber) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *PersPhoneNumber) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode int32 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Number      int64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persPhone_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_persPhone_proto_msgTypes[11]
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
	return file_persPhone_proto_rawDescGZIP(), []int{11}
}

func (x *RequestBody) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *RequestBody) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_persPhone_proto protoreflect.FileDescriptor

var file_persPhone_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x98, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3d, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x50, 0x68, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xb6, 0x01, 0x0a, 0x1c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x3d, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x50,
	0x68, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x49,
	0x64, 0x22, 0x3d, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x49, 0x64,
	0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x10, 0x70, 0x65, 0x72, 0x73, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x10, 0x70, 0x65, 0x72, 0x73, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x6a, 0x0a, 0x0f, 0x50, 0x65,
	0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x50, 0x68, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x32, 0xa3, 0x06, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x2c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x3a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x12, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x94, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x70, 0x65,
	0x72, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xaa, 0x01, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x1f, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x65,
	0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9d, 0x01, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x65,
	0x72, 0x73, 0x5f, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x2a, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x12, 0x12, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_persPhone_proto_rawDescOnce sync.Once
	file_persPhone_proto_rawDescData = file_persPhone_proto_rawDesc
)

func file_persPhone_proto_rawDescGZIP() []byte {
	file_persPhone_proto_rawDescOnce.Do(func() {
		file_persPhone_proto_rawDescData = protoimpl.X.CompressGZIP(file_persPhone_proto_rawDescData)
	})
	return file_persPhone_proto_rawDescData
}

var file_persPhone_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_persPhone_proto_goTypes = []interface{}{
	(*CreatePersPhoneNumberRequest)(nil),  // 0: PersonalNumber.CreatePersPhoneNumberRequest
	(*CreatePersPhoneNumberResponse)(nil), // 1: PersonalNumber.CreatePersPhoneNumberResponse
	(*GetPersPhoneNumberRequest)(nil),     // 2: PersonalNumber.GetPersPhoneNumberRequest
	(*GetPersPhoneNumberResponse)(nil),    // 3: PersonalNumber.GetPersPhoneNumberResponse
	(*UpdatePersPhoneNumberRequest)(nil),  // 4: PersonalNumber.UpdatePersPhoneNumberRequest
	(*UpdatePersPhoneNumberResponse)(nil), // 5: PersonalNumber.UpdatePersPhoneNumberResponse
	(*DeletePersPhoneNumberRequest)(nil),  // 6: PersonalNumber.DeletePersPhoneNumberRequest
	(*DeletePersPhoneNumberResponse)(nil), // 7: PersonalNumber.DeletePersPhoneNumberResponse
	(*GetPersPhoneNumbersRequest)(nil),    // 8: PersonalNumber.GetPersPhoneNumbersRequest
	(*GetPersPhoneNumbersResponse)(nil),   // 9: PersonalNumber.GetPersPhoneNumbersResponse
	(*PersPhoneNumber)(nil),               // 10: PersonalNumber.PersPhoneNumber
	(*RequestBody)(nil),                   // 11: PersonalNumber.RequestBody
}
var file_persPhone_proto_depIdxs = []int32{
	11, // 0: PersonalNumber.CreatePersPhoneNumberRequest.requestbody:type_name -> PersonalNumber.RequestBody
	11, // 1: PersonalNumber.UpdatePersPhoneNumberRequest.requestbody:type_name -> PersonalNumber.RequestBody
	10, // 2: PersonalNumber.GetPersPhoneNumbersResponse.persphonenumbers:type_name -> PersonalNumber.PersPhoneNumber
	0,  // 3: PersonalNumber.PersPhoneService.CreatePersPhoneNumber:input_type -> PersonalNumber.CreatePersPhoneNumberRequest
	2,  // 4: PersonalNumber.PersPhoneService.GetPersPhoneNumber:input_type -> PersonalNumber.GetPersPhoneNumberRequest
	4,  // 5: PersonalNumber.PersPhoneService.UpdatePersPhoneNumber:input_type -> PersonalNumber.UpdatePersPhoneNumberRequest
	6,  // 6: PersonalNumber.PersPhoneService.DeletePersPhoneNumber:input_type -> PersonalNumber.DeletePersPhoneNumberRequest
	8,  // 7: PersonalNumber.PersPhoneService.GetPersPhoneNumbers:input_type -> PersonalNumber.GetPersPhoneNumbersRequest
	1,  // 8: PersonalNumber.PersPhoneService.CreatePersPhoneNumber:output_type -> PersonalNumber.CreatePersPhoneNumberResponse
	3,  // 9: PersonalNumber.PersPhoneService.GetPersPhoneNumber:output_type -> PersonalNumber.GetPersPhoneNumberResponse
	5,  // 10: PersonalNumber.PersPhoneService.UpdatePersPhoneNumber:output_type -> PersonalNumber.UpdatePersPhoneNumberResponse
	7,  // 11: PersonalNumber.PersPhoneService.DeletePersPhoneNumber:output_type -> PersonalNumber.DeletePersPhoneNumberResponse
	9,  // 12: PersonalNumber.PersPhoneService.GetPersPhoneNumbers:output_type -> PersonalNumber.GetPersPhoneNumbersResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_persPhone_proto_init() }
func file_persPhone_proto_init() {
	if File_persPhone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_persPhone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersPhoneNumberRequest); i {
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
		file_persPhone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersPhoneNumberResponse); i {
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
		file_persPhone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersPhoneNumberRequest); i {
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
		file_persPhone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersPhoneNumberResponse); i {
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
		file_persPhone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersPhoneNumberRequest); i {
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
		file_persPhone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersPhoneNumberResponse); i {
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
		file_persPhone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersPhoneNumberRequest); i {
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
		file_persPhone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersPhoneNumberResponse); i {
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
		file_persPhone_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersPhoneNumbersRequest); i {
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
		file_persPhone_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersPhoneNumbersResponse); i {
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
		file_persPhone_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersPhoneNumber); i {
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
		file_persPhone_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_persPhone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_persPhone_proto_goTypes,
		DependencyIndexes: file_persPhone_proto_depIdxs,
		MessageInfos:      file_persPhone_proto_msgTypes,
	}.Build()
	File_persPhone_proto = out.File
	file_persPhone_proto_rawDesc = nil
	file_persPhone_proto_goTypes = nil
	file_persPhone_proto_depIdxs = nil
}
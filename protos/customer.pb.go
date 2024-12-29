// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/customer.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nik string `protobuf:"bytes,1,opt,name=nik,proto3" json:"nik,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerRequest) GetNik() string {
	if x != nil {
		return x.Nik
	}
	return ""
}

type GetCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nik          string                 `protobuf:"bytes,2,opt,name=nik,proto3" json:"nik,omitempty"`
	FullName     string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	LegalName    string                 `protobuf:"bytes,4,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	PlaceOfBirth string                 `protobuf:"bytes,5,opt,name=place_of_birth,json=placeOfBirth,proto3" json:"place_of_birth,omitempty"`
	DateOfBirth  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Salary       float64                `protobuf:"fixed64,7,opt,name=salary,proto3" json:"salary,omitempty"`
	PhotoKtp     string                 `protobuf:"bytes,8,opt,name=photo_ktp,json=photoKtp,proto3" json:"photo_ktp,omitempty"`
	PhotoSelfie  string                 `protobuf:"bytes,9,opt,name=photo_selfie,json=photoSelfie,proto3" json:"photo_selfie,omitempty"`
}

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCustomerResponse) GetNik() string {
	if x != nil {
		return x.Nik
	}
	return ""
}

func (x *GetCustomerResponse) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *GetCustomerResponse) GetLegalName() string {
	if x != nil {
		return x.LegalName
	}
	return ""
}

func (x *GetCustomerResponse) GetPlaceOfBirth() string {
	if x != nil {
		return x.PlaceOfBirth
	}
	return ""
}

func (x *GetCustomerResponse) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *GetCustomerResponse) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *GetCustomerResponse) GetPhotoKtp() string {
	if x != nil {
		return x.PhotoKtp
	}
	return ""
}

func (x *GetCustomerResponse) GetPhotoSelfie() string {
	if x != nil {
		return x.PhotoSelfie
	}
	return ""
}

type CreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nik          string                 `protobuf:"bytes,1,opt,name=nik,proto3" json:"nik,omitempty"`
	FullName     string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	LegalName    string                 `protobuf:"bytes,3,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	PlaceOfBirth string                 `protobuf:"bytes,4,opt,name=place_of_birth,json=placeOfBirth,proto3" json:"place_of_birth,omitempty"`
	DateOfBirth  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Salary       float64                `protobuf:"fixed64,6,opt,name=salary,proto3" json:"salary,omitempty"`
	PhotoKtp     string                 `protobuf:"bytes,7,opt,name=photo_ktp,json=photoKtp,proto3" json:"photo_ktp,omitempty"`
	PhotoSelfie  string                 `protobuf:"bytes,8,opt,name=photo_selfie,json=photoSelfie,proto3" json:"photo_selfie,omitempty"`
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomerRequest) GetNik() string {
	if x != nil {
		return x.Nik
	}
	return ""
}

func (x *CreateCustomerRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateCustomerRequest) GetLegalName() string {
	if x != nil {
		return x.LegalName
	}
	return ""
}

func (x *CreateCustomerRequest) GetPlaceOfBirth() string {
	if x != nil {
		return x.PlaceOfBirth
	}
	return ""
}

func (x *CreateCustomerRequest) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *CreateCustomerRequest) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateCustomerRequest) GetPhotoKtp() string {
	if x != nil {
		return x.PhotoKtp
	}
	return ""
}

func (x *CreateCustomerRequest) GetPhotoSelfie() string {
	if x != nil {
		return x.PhotoSelfie
	}
	return ""
}

type UpdateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nik          string                 `protobuf:"bytes,2,opt,name=nik,proto3" json:"nik,omitempty"`
	FullName     string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	LegalName    string                 `protobuf:"bytes,4,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	PlaceOfBirth string                 `protobuf:"bytes,5,opt,name=place_of_birth,json=placeOfBirth,proto3" json:"place_of_birth,omitempty"`
	DateOfBirth  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Salary       float64                `protobuf:"fixed64,7,opt,name=salary,proto3" json:"salary,omitempty"` // 64-bit floating-point number
	PhotoKtp     string                 `protobuf:"bytes,8,opt,name=photo_ktp,json=photoKtp,proto3" json:"photo_ktp,omitempty"`
	PhotoSelfie  string                 `protobuf:"bytes,9,opt,name=photo_selfie,json=photoSelfie,proto3" json:"photo_selfie,omitempty"`
}

func (x *UpdateCustomerRequest) Reset() {
	*x = UpdateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerRequest) ProtoMessage() {}

func (x *UpdateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCustomerRequest) GetNik() string {
	if x != nil {
		return x.Nik
	}
	return ""
}

func (x *UpdateCustomerRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateCustomerRequest) GetLegalName() string {
	if x != nil {
		return x.LegalName
	}
	return ""
}

func (x *UpdateCustomerRequest) GetPlaceOfBirth() string {
	if x != nil {
		return x.PlaceOfBirth
	}
	return ""
}

func (x *UpdateCustomerRequest) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *UpdateCustomerRequest) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateCustomerRequest) GetPhotoKtp() string {
	if x != nil {
		return x.PhotoKtp
	}
	return ""
}

func (x *UpdateCustomerRequest) GetPhotoSelfie() string {
	if x != nil {
		return x.PhotoSelfie
	}
	return ""
}

type DeleteCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nik string `protobuf:"bytes,1,opt,name=nik,proto3" json:"nik,omitempty"`
}

func (x *DeleteCustomerRequest) Reset() {
	*x = DeleteCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerRequest) ProtoMessage() {}

func (x *DeleteCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerRequest.ProtoReflect.Descriptor instead.
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCustomerRequest) GetNik() string {
	if x != nil {
		return x.Nik
	}
	return ""
}

type GetCustomerByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerByIdRequest) Reset() {
	*x = GetCustomerByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdRequest) ProtoMessage() {}

func (x *GetCustomerByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetCustomerByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerLimitRequest) Reset() {
	*x = GetCustomerLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerLimitRequest) ProtoMessage() {}

func (x *GetCustomerLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerLimitRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerLimitRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{6}
}

func (x *GetCustomerLimitRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId    string  `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Tenor_1Month  float64 `protobuf:"fixed64,2,opt,name=tenor_1_month,json=tenor1Month,proto3" json:"tenor_1_month,omitempty"`
	Tenor_2Months float64 `protobuf:"fixed64,3,opt,name=tenor_2_months,json=tenor2Months,proto3" json:"tenor_2_months,omitempty"`
	Tenor_3Months float64 `protobuf:"fixed64,4,opt,name=tenor_3_months,json=tenor3Months,proto3" json:"tenor_3_months,omitempty"`
	Tenor_4Months float64 `protobuf:"fixed64,5,opt,name=tenor_4_months,json=tenor4Months,proto3" json:"tenor_4_months,omitempty"`
}

func (x *GetCustomerLimitResponse) Reset() {
	*x = GetCustomerLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerLimitResponse) ProtoMessage() {}

func (x *GetCustomerLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerLimitResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerLimitResponse) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{7}
}

func (x *GetCustomerLimitResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *GetCustomerLimitResponse) GetTenor_1Month() float64 {
	if x != nil {
		return x.Tenor_1Month
	}
	return 0
}

func (x *GetCustomerLimitResponse) GetTenor_2Months() float64 {
	if x != nil {
		return x.Tenor_2Months
	}
	return 0
}

func (x *GetCustomerLimitResponse) GetTenor_3Months() float64 {
	if x != nil {
		return x.Tenor_3Months
	}
	return 0
}

func (x *GetCustomerLimitResponse) GetTenor_4Months() float64 {
	if x != nil {
		return x.Tenor_4Months
	}
	return 0
}

type UpdateCustomerLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId    string  `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Tenor_1Month  float64 `protobuf:"fixed64,2,opt,name=tenor_1_month,json=tenor1Month,proto3" json:"tenor_1_month,omitempty"`
	Tenor_2Months float64 `protobuf:"fixed64,3,opt,name=tenor_2_months,json=tenor2Months,proto3" json:"tenor_2_months,omitempty"`
	Tenor_3Months float64 `protobuf:"fixed64,4,opt,name=tenor_3_months,json=tenor3Months,proto3" json:"tenor_3_months,omitempty"`
	Tenor_4Months float64 `protobuf:"fixed64,5,opt,name=tenor_4_months,json=tenor4Months,proto3" json:"tenor_4_months,omitempty"`
}

func (x *UpdateCustomerLimitRequest) Reset() {
	*x = UpdateCustomerLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerLimitRequest) ProtoMessage() {}

func (x *UpdateCustomerLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerLimitRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerLimitRequest) Descriptor() ([]byte, []int) {
	return file_protos_customer_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCustomerLimitRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *UpdateCustomerLimitRequest) GetTenor_1Month() float64 {
	if x != nil {
		return x.Tenor_1Month
	}
	return 0
}

func (x *UpdateCustomerLimitRequest) GetTenor_2Months() float64 {
	if x != nil {
		return x.Tenor_2Months
	}
	return 0
}

func (x *UpdateCustomerLimitRequest) GetTenor_3Months() float64 {
	if x != nil {
		return x.Tenor_3Months
	}
	return 0
}

func (x *UpdateCustomerLimitRequest) GetTenor_4Months() float64 {
	if x != nil {
		return x.Tenor_4Months
	}
	return 0
}

var File_protos_customer_proto protoreflect.FileDescriptor

var file_protos_customer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6e, 0x69, 0x6b, 0x22, 0xb1, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x69, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x6b, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x65, 0x67, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x5f, 0x6b, 0x74, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x4b, 0x74, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f,
	0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x22, 0xa3, 0x02, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6e, 0x69, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x6b, 0x74, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4b, 0x74, 0x70, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x22,
	0xb3, 0x02, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x67, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65,
	0x67, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x3e, 0x0a,
	0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x6b,
	0x74, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4b,
	0x74, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6c, 0x66,
	0x69, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x53,
	0x65, 0x6c, 0x66, 0x69, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x69, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x6b,
	0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd1, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x5f, 0x31, 0x5f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6e, 0x6f,
	0x72, 0x31, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x6e, 0x6f, 0x72,
	0x5f, 0x32, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0c, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x32, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x5f, 0x33, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x33, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x5f, 0x34, 0x5f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x65, 0x6e,
	0x6f, 0x72, 0x34, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x65, 0x6e,
	0x6f, 0x72, 0x5f, 0x31, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x31, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x5f, 0x32, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x32, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x5f, 0x33, 0x5f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x65, 0x6e,
	0x6f, 0x72, 0x33, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x6e,
	0x6f, 0x72, 0x5f, 0x34, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x34, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x32,
	0xae, 0x04, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_customer_proto_rawDescOnce sync.Once
	file_protos_customer_proto_rawDescData = file_protos_customer_proto_rawDesc
)

func file_protos_customer_proto_rawDescGZIP() []byte {
	file_protos_customer_proto_rawDescOnce.Do(func() {
		file_protos_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_customer_proto_rawDescData)
	})
	return file_protos_customer_proto_rawDescData
}

var file_protos_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_customer_proto_goTypes = []interface{}{
	(*GetCustomerRequest)(nil),         // 0: protos.GetCustomerRequest
	(*GetCustomerResponse)(nil),        // 1: protos.GetCustomerResponse
	(*CreateCustomerRequest)(nil),      // 2: protos.CreateCustomerRequest
	(*UpdateCustomerRequest)(nil),      // 3: protos.UpdateCustomerRequest
	(*DeleteCustomerRequest)(nil),      // 4: protos.DeleteCustomerRequest
	(*GetCustomerByIdRequest)(nil),     // 5: protos.GetCustomerByIdRequest
	(*GetCustomerLimitRequest)(nil),    // 6: protos.GetCustomerLimitRequest
	(*GetCustomerLimitResponse)(nil),   // 7: protos.GetCustomerLimitResponse
	(*UpdateCustomerLimitRequest)(nil), // 8: protos.UpdateCustomerLimitRequest
	(*timestamppb.Timestamp)(nil),      // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 10: google.protobuf.Empty
}
var file_protos_customer_proto_depIdxs = []int32{
	9,  // 0: protos.GetCustomerResponse.date_of_birth:type_name -> google.protobuf.Timestamp
	9,  // 1: protos.CreateCustomerRequest.date_of_birth:type_name -> google.protobuf.Timestamp
	9,  // 2: protos.UpdateCustomerRequest.date_of_birth:type_name -> google.protobuf.Timestamp
	0,  // 3: protos.customerService.GetCustomer:input_type -> protos.GetCustomerRequest
	5,  // 4: protos.customerService.GetCustomerById:input_type -> protos.GetCustomerByIdRequest
	2,  // 5: protos.customerService.CreateCustomer:input_type -> protos.CreateCustomerRequest
	3,  // 6: protos.customerService.UpdateCustomer:input_type -> protos.UpdateCustomerRequest
	4,  // 7: protos.customerService.DeleteCustomer:input_type -> protos.DeleteCustomerRequest
	6,  // 8: protos.customerService.GetCustomerLimit:input_type -> protos.GetCustomerLimitRequest
	8,  // 9: protos.customerService.UpdateCustomerLimit:input_type -> protos.UpdateCustomerLimitRequest
	1,  // 10: protos.customerService.GetCustomer:output_type -> protos.GetCustomerResponse
	1,  // 11: protos.customerService.GetCustomerById:output_type -> protos.GetCustomerResponse
	10, // 12: protos.customerService.CreateCustomer:output_type -> google.protobuf.Empty
	10, // 13: protos.customerService.UpdateCustomer:output_type -> google.protobuf.Empty
	10, // 14: protos.customerService.DeleteCustomer:output_type -> google.protobuf.Empty
	7,  // 15: protos.customerService.GetCustomerLimit:output_type -> protos.GetCustomerLimitResponse
	10, // 16: protos.customerService.UpdateCustomerLimit:output_type -> google.protobuf.Empty
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protos_customer_proto_init() }
func file_protos_customer_proto_init() {
	if File_protos_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_protos_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerResponse); i {
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
		file_protos_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerRequest); i {
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
		file_protos_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerRequest); i {
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
		file_protos_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCustomerRequest); i {
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
		file_protos_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByIdRequest); i {
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
		file_protos_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerLimitRequest); i {
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
		file_protos_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerLimitResponse); i {
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
		file_protos_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerLimitRequest); i {
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
			RawDescriptor: file_protos_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_customer_proto_goTypes,
		DependencyIndexes: file_protos_customer_proto_depIdxs,
		MessageInfos:      file_protos_customer_proto_msgTypes,
	}.Build()
	File_protos_customer_proto = out.File
	file_protos_customer_proto_rawDesc = nil
	file_protos_customer_proto_goTypes = nil
	file_protos_customer_proto_depIdxs = nil
}

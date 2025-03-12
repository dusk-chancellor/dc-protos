// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: orchestrator.proto

package orchestrator

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Status int32

const (
	Status_PENDING    Status = 0
	Status_CALCULATED Status = 1
	Status_ERROR      Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDING",
		1: "CALCULATED",
		2: "ERROR",
	}
	Status_value = map[string]int32{
		"PENDING":    0,
		"CALCULATED": 1,
		"ERROR":      2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestrator_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_orchestrator_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

type Expression struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                       // uuid
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // uuid
	Expression    string                 `protobuf:"bytes,3,opt,name=expression,proto3" json:"expression,omitempty"`
	Answer        string                 `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
	Status        Status                 `protobuf:"varint,5,opt,name=status,proto3,enum=orchestrator.Status" json:"status,omitempty"`
	CreatedAt     *timestamp.Timestamp   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Expression) Reset() {
	*x = Expression{}
	mi := &file_orchestrator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression) ProtoMessage() {}

func (x *Expression) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression.ProtoReflect.Descriptor instead.
func (*Expression) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *Expression) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Expression) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Expression) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *Expression) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Expression) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

func (x *Expression) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ExpressionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpressionRequest) Reset() {
	*x = ExpressionRequest{}
	mi := &file_orchestrator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpressionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionRequest) ProtoMessage() {}

func (x *ExpressionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionRequest.ProtoReflect.Descriptor instead.
func (*ExpressionRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *ExpressionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExpressionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ExpressionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Expression    *Expression            `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpressionResponse) Reset() {
	*x = ExpressionResponse{}
	mi := &file_orchestrator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpressionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionResponse) ProtoMessage() {}

func (x *ExpressionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionResponse.ProtoReflect.Descriptor instead.
func (*ExpressionResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{2}
}

func (x *ExpressionResponse) GetExpression() *Expression {
	if x != nil {
		return x.Expression
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Expression    string                 `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_orchestrator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRequest) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type GetByUserIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetByUserIdRequest) Reset() {
	*x = GetByUserIdRequest{}
	mi := &file_orchestrator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserIdRequest) ProtoMessage() {}

func (x *GetByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{4}
}

func (x *GetByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetByUserIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Expression    []*Expression          `protobuf:"bytes,1,rep,name=expression,proto3" json:"expression,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetByUserIdResponse) Reset() {
	*x = GetByUserIdResponse{}
	mi := &file_orchestrator_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserIdResponse) ProtoMessage() {}

func (x *GetByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{5}
}

func (x *GetByUserIdResponse) GetExpression() []*Expression {
	if x != nil {
		return x.Expression
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_orchestrator_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_orchestrator_proto protoreflect.FileDescriptor

var file_orchestrator_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3c, 0x0a, 0x11,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x12, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2a, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x4c, 0x43, 0x55,
	0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x02, 0x32, 0x90, 0x03, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orchestrator_proto_rawDescOnce sync.Once
	file_orchestrator_proto_rawDescData = file_orchestrator_proto_rawDesc
)

func file_orchestrator_proto_rawDescGZIP() []byte {
	file_orchestrator_proto_rawDescOnce.Do(func() {
		file_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_proto_rawDescData)
	})
	return file_orchestrator_proto_rawDescData
}

var file_orchestrator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_orchestrator_proto_goTypes = []any{
	(Status)(0),                 // 0: orchestrator.Status
	(*Expression)(nil),          // 1: orchestrator.Expression
	(*ExpressionRequest)(nil),   // 2: orchestrator.ExpressionRequest
	(*ExpressionResponse)(nil),  // 3: orchestrator.ExpressionResponse
	(*CreateRequest)(nil),       // 4: orchestrator.CreateRequest
	(*GetByUserIdRequest)(nil),  // 5: orchestrator.GetByUserIdRequest
	(*GetByUserIdResponse)(nil), // 6: orchestrator.GetByUserIdResponse
	(*DeleteResponse)(nil),      // 7: orchestrator.DeleteResponse
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_orchestrator_proto_depIdxs = []int32{
	0, // 0: orchestrator.Expression.status:type_name -> orchestrator.Status
	8, // 1: orchestrator.Expression.created_at:type_name -> google.protobuf.Timestamp
	1, // 2: orchestrator.ExpressionResponse.expression:type_name -> orchestrator.Expression
	1, // 3: orchestrator.GetByUserIdResponse.expression:type_name -> orchestrator.Expression
	4, // 4: orchestrator.ExpressionService.Create:input_type -> orchestrator.CreateRequest
	2, // 5: orchestrator.ExpressionService.Get:input_type -> orchestrator.ExpressionRequest
	5, // 6: orchestrator.ExpressionService.GetByUserId:input_type -> orchestrator.GetByUserIdRequest
	2, // 7: orchestrator.ExpressionService.Update:input_type -> orchestrator.ExpressionRequest
	2, // 8: orchestrator.ExpressionService.Delete:input_type -> orchestrator.ExpressionRequest
	3, // 9: orchestrator.ExpressionService.Create:output_type -> orchestrator.ExpressionResponse
	3, // 10: orchestrator.ExpressionService.Get:output_type -> orchestrator.ExpressionResponse
	6, // 11: orchestrator.ExpressionService.GetByUserId:output_type -> orchestrator.GetByUserIdResponse
	3, // 12: orchestrator.ExpressionService.Update:output_type -> orchestrator.ExpressionResponse
	7, // 13: orchestrator.ExpressionService.Delete:output_type -> orchestrator.DeleteResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orchestrator_proto_init() }
func file_orchestrator_proto_init() {
	if File_orchestrator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orchestrator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestrator_proto_goTypes,
		DependencyIndexes: file_orchestrator_proto_depIdxs,
		EnumInfos:         file_orchestrator_proto_enumTypes,
		MessageInfos:      file_orchestrator_proto_msgTypes,
	}.Build()
	File_orchestrator_proto = out.File
	file_orchestrator_proto_rawDesc = nil
	file_orchestrator_proto_goTypes = nil
	file_orchestrator_proto_depIdxs = nil
}

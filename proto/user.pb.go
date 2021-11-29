// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/user.proto

package proto

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

type RequestInf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Valor   int32  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *RequestInf) Reset() {
	*x = RequestInf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInf) ProtoMessage() {}

func (x *RequestInf) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInf.ProtoReflect.Descriptor instead.
func (*RequestInf) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *RequestInf) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RequestInf) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *RequestInf) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

type RequestDel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
}

func (x *RequestDel) Reset() {
	*x = RequestDel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDel) ProtoMessage() {}

func (x *RequestDel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDel.ProtoReflect.Descriptor instead.
func (*RequestDel) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *RequestDel) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RequestDel) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

type ResponseBroker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ResponseBroker) Reset() {
	*x = ResponseBroker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBroker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBroker) ProtoMessage() {}

func (x *ResponseBroker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBroker.ProtoReflect.Descriptor instead.
func (*ResponseBroker) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseBroker) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ResponseFulcrum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector string `protobuf:"bytes,1,opt,name=vector,proto3" json:"vector,omitempty"`
}

func (x *ResponseFulcrum) Reset() {
	*x = ResponseFulcrum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFulcrum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFulcrum) ProtoMessage() {}

func (x *ResponseFulcrum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFulcrum.ProtoReflect.Descriptor instead.
func (*ResponseFulcrum) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseFulcrum) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

type RequestLeia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
}

func (x *RequestLeia) Reset() {
	*x = RequestLeia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLeia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLeia) ProtoMessage() {}

func (x *RequestLeia) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLeia.ProtoReflect.Descriptor instead.
func (*RequestLeia) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *RequestLeia) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RequestLeia) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

type ResponseRebelds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor  int32  `protobuf:"varint,1,opt,name=valor,proto3" json:"valor,omitempty"`
	Vector string `protobuf:"bytes,2,opt,name=vector,proto3" json:"vector,omitempty"`
}

func (x *ResponseRebelds) Reset() {
	*x = ResponseRebelds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRebelds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRebelds) ProtoMessage() {}

func (x *ResponseRebelds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRebelds.ProtoReflect.Descriptor instead.
func (*ResponseRebelds) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseRebelds) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *ResponseRebelds) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x54, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x3e,
	0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x22, 0x2a,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x29, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x65, 0x69, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0xaf, 0x02, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x1a,
	0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65,
	0x6c, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x11, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x69, 0x61,
	0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x00, 0x32, 0xf4, 0x01, 0x0a, 0x0f, 0x46, 0x75,
	0x6c, 0x63, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x1a, 0x15,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x75,
	0x6c, 0x63, 0x72, 0x75, 0x6d, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x4c, 0x61, 0x62, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData = file_proto_user_proto_rawDesc
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_proto_rawDescData)
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_user_proto_goTypes = []interface{}{
	(*RequestInf)(nil),      // 0: grpc.RequestInf
	(*RequestDel)(nil),      // 1: grpc.RequestDel
	(*ResponseBroker)(nil),  // 2: grpc.ResponseBroker
	(*ResponseFulcrum)(nil), // 3: grpc.ResponseFulcrum
	(*RequestLeia)(nil),     // 4: grpc.RequestLeia
	(*ResponseRebelds)(nil), // 5: grpc.ResponseRebelds
}
var file_proto_user_proto_depIdxs = []int32{
	0, // 0: grpc.BrokerServices.AddCity:input_type -> grpc.RequestInf
	0, // 1: grpc.BrokerServices.UpdateName:input_type -> grpc.RequestInf
	0, // 2: grpc.BrokerServices.UpdateNumber:input_type -> grpc.RequestInf
	1, // 3: grpc.BrokerServices.DeleteCity:input_type -> grpc.RequestDel
	4, // 4: grpc.BrokerServices.GetNumberRebelds:input_type -> grpc.RequestLeia
	0, // 5: grpc.FulcramServices.AddCity:input_type -> grpc.RequestInf
	0, // 6: grpc.FulcramServices.UpdateName:input_type -> grpc.RequestInf
	0, // 7: grpc.FulcramServices.UpdateNumber:input_type -> grpc.RequestInf
	1, // 8: grpc.FulcramServices.DeleteCity:input_type -> grpc.RequestDel
	2, // 9: grpc.BrokerServices.AddCity:output_type -> grpc.ResponseBroker
	2, // 10: grpc.BrokerServices.UpdateName:output_type -> grpc.ResponseBroker
	2, // 11: grpc.BrokerServices.UpdateNumber:output_type -> grpc.ResponseBroker
	2, // 12: grpc.BrokerServices.DeleteCity:output_type -> grpc.ResponseBroker
	5, // 13: grpc.BrokerServices.GetNumberRebelds:output_type -> grpc.ResponseRebelds
	3, // 14: grpc.FulcramServices.AddCity:output_type -> grpc.ResponseFulcrum
	3, // 15: grpc.FulcramServices.UpdateName:output_type -> grpc.ResponseFulcrum
	3, // 16: grpc.FulcramServices.UpdateNumber:output_type -> grpc.ResponseFulcrum
	3, // 17: grpc.FulcramServices.DeleteCity:output_type -> grpc.ResponseFulcrum
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInf); i {
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
		file_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDel); i {
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
		file_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBroker); i {
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
		file_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFulcrum); i {
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
		file_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLeia); i {
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
		file_proto_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRebelds); i {
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
			RawDescriptor: file_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_rawDesc = nil
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}
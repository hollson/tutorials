// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CallRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74a5129edc93dca, []int{0}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallResponse struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74a5129edc93dca, []int{1}
}

func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (m *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(m, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CallRequest)(nil), "CallRequest")
	proto.RegisterType((*CallResponse)(nil), "CallResponse")
}

func init() { proto.RegisterFile("proto/rpc.proto", fileDescriptor_d74a5129edc93dca) }

var fileDescriptor_d74a5129edc93dca = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0xb3, 0x94, 0x14, 0xb9, 0xb8, 0x9d, 0x13, 0x73, 0x72,
	0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x0d, 0x2e, 0x1e, 0x88, 0x92, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09,
	0x26, 0xb0, 0x32, 0x18, 0xd7, 0x48, 0x97, 0x8b, 0xc5, 0x25, 0x35, 0x37, 0x5f, 0x48, 0x95, 0x8b,
	0x05, 0xa4, 0x43, 0x88, 0x47, 0x0f, 0xc9, 0x6c, 0x29, 0x5e, 0x3d, 0x64, 0x63, 0x94, 0x18, 0x92,
	0xd8, 0xc0, 0x4e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5a, 0x4c, 0x8f, 0x95, 0x00,
	0x00, 0x00,
}

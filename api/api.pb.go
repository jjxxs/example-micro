// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting             string   `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type IncRequest struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncRequest) Reset()         { *m = IncRequest{} }
func (m *IncRequest) String() string { return proto.CompactTextString(m) }
func (*IncRequest) ProtoMessage()    {}
func (*IncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *IncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncRequest.Unmarshal(m, b)
}
func (m *IncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncRequest.Marshal(b, m, deterministic)
}
func (m *IncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncRequest.Merge(m, src)
}
func (m *IncRequest) XXX_Size() int {
	return xxx_messageInfo_IncRequest.Size(m)
}
func (m *IncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncRequest proto.InternalMessageInfo

func (m *IncRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SumResponse struct {
	Counter              int32    `protobuf:"varint,4,opt,name=counter,proto3" json:"counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
	proto.RegisterType((*IncRequest)(nil), "IncRequest")
	proto.RegisterType((*SumResponse)(nil), "SumResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0xce, 0x82, 0x30,
	0x14, 0x85, 0xe1, 0x07, 0x7e, 0xe4, 0x02, 0x0e, 0x9d, 0x1a, 0x26, 0xd2, 0x45, 0x12, 0x93, 0x0e,
	0xf0, 0x08, 0x0e, 0xca, 0x8a, 0x4f, 0x80, 0xe4, 0x86, 0x90, 0x40, 0x8b, 0xa5, 0xbc, 0xbf, 0xb1,
	0x22, 0xa2, 0xdb, 0x3d, 0xc9, 0xc9, 0x97, 0xef, 0x1e, 0x08, 0xea, 0xb1, 0xe3, 0xa3, 0x92, 0x5a,
	0x32, 0x06, 0xd1, 0x05, 0xfb, 0x5e, 0x56, 0x78, 0x9f, 0x71, 0xd2, 0x84, 0x80, 0x2b, 0xea, 0x01,
	0xa9, 0x9d, 0xda, 0x59, 0x50, 0x99, 0x9b, 0x1d, 0x21, 0x5e, 0x3a, 0xd3, 0x28, 0xc5, 0x84, 0x24,
	0x81, 0x5d, 0xab, 0x10, 0x75, 0x27, 0x5a, 0xfa, 0x67, 0x8a, 0x6b, 0x66, 0x29, 0x40, 0x29, 0x9a,
	0x5f, 0x9c, 0xb3, 0xc1, 0x1d, 0x20, 0xbc, 0xce, 0xc3, 0x0a, 0xa3, 0xe0, 0x37, 0x72, 0x16, 0x1a,
	0x15, 0x75, 0x53, 0x3b, 0xf3, 0xaa, 0x77, 0xcc, 0x0b, 0xf0, 0xcf, 0x4f, 0x2c, 0x2a, 0x92, 0x81,
	0x67, 0x14, 0x48, 0xcc, 0xb7, 0xba, 0xc9, 0x9e, 0x7f, 0x99, 0x31, 0x2b, 0xcf, 0x97, 0x87, 0x4e,
	0x2f, 0x08, 0x61, 0xe0, 0x94, 0xa2, 0x21, 0x21, 0xff, 0x58, 0x25, 0x11, 0xdf, 0x08, 0x30, 0xeb,
	0xf6, 0x6f, 0xb6, 0x28, 0x1e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0xd5, 0x38, 0x04, 0x18, 0x01,
	0x00, 0x00,
}
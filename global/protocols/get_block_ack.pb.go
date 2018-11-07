// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_block_ack.proto

package protocols

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetBlockAck struct {
	Block                []byte   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockAck) Reset()         { *m = GetBlockAck{} }
func (m *GetBlockAck) String() string { return proto.CompactTextString(m) }
func (*GetBlockAck) ProtoMessage()    {}
func (*GetBlockAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_block_ack_4ecc37e042fecca0, []int{0}
}
func (m *GetBlockAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockAck.Unmarshal(m, b)
}
func (m *GetBlockAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockAck.Marshal(b, m, deterministic)
}
func (dst *GetBlockAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockAck.Merge(dst, src)
}
func (m *GetBlockAck) XXX_Size() int {
	return xxx_messageInfo_GetBlockAck.Size(m)
}
func (m *GetBlockAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockAck.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockAck proto.InternalMessageInfo

func (m *GetBlockAck) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBlockAck)(nil), "protocols.GetBlockAck")
}

func init() { proto.RegisterFile("get_block_ack.proto", fileDescriptor_get_block_ack_4ecc37e042fecca0) }

var fileDescriptor_get_block_ack_4ecc37e042fecca0 = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4f, 0x2d, 0x89,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x8e, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0x52, 0xe6, 0xe2, 0x76, 0x4f, 0x2d,
	0x71, 0x02, 0xc9, 0x3b, 0x26, 0x67, 0x0b, 0x89, 0x70, 0xb1, 0x82, 0xd5, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0xf0, 0x04, 0x41, 0x38, 0x49, 0x6c, 0x60, 0x5d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x31, 0xab, 0x58, 0x40, 0x4c, 0x00, 0x00, 0x00,
}

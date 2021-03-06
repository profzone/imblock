// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat_ack.proto

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

type HeartbeatAck struct {
	Sequence             int64    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatAck) Reset()         { *m = HeartbeatAck{} }
func (m *HeartbeatAck) String() string { return proto.CompactTextString(m) }
func (*HeartbeatAck) ProtoMessage()    {}
func (*HeartbeatAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_ack_5265d35157de3c26, []int{0}
}
func (m *HeartbeatAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatAck.Unmarshal(m, b)
}
func (m *HeartbeatAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatAck.Marshal(b, m, deterministic)
}
func (dst *HeartbeatAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatAck.Merge(dst, src)
}
func (m *HeartbeatAck) XXX_Size() int {
	return xxx_messageInfo_HeartbeatAck.Size(m)
}
func (m *HeartbeatAck) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatAck.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatAck proto.InternalMessageInfo

func (m *HeartbeatAck) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*HeartbeatAck)(nil), "protocols.HeartbeatAck")
}

func init() { proto.RegisterFile("heartbeat_ack.proto", fileDescriptor_heartbeat_ack_5265d35157de3c26) }

var fileDescriptor_heartbeat_ack_5265d35157de3c26 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0x89, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0xd2, 0xe2, 0xe2, 0xf1, 0x80, 0x29,
	0x70, 0x4c, 0xce, 0x16, 0x92, 0xe2, 0xe2, 0x28, 0x4e, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf3, 0x93, 0xd8, 0xc0, 0x9a, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x12, 0xe9, 0x30, 0x1e, 0x53, 0x00, 0x00, 0x00,
}

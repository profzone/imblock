// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broadcast_new_peer_to_peers.proto

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

type BroadcastNewPeerToPeers struct {
	Ip                   []byte   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Guid                 int64    `protobuf:"varint,3,opt,name=guid,proto3" json:"guid,omitempty"`
	Version              uint32   `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastNewPeerToPeers) Reset()         { *m = BroadcastNewPeerToPeers{} }
func (m *BroadcastNewPeerToPeers) String() string { return proto.CompactTextString(m) }
func (*BroadcastNewPeerToPeers) ProtoMessage()    {}
func (*BroadcastNewPeerToPeers) Descriptor() ([]byte, []int) {
	return fileDescriptor_broadcast_new_peer_to_peers_cab8d60bbe394b90, []int{0}
}
func (m *BroadcastNewPeerToPeers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastNewPeerToPeers.Unmarshal(m, b)
}
func (m *BroadcastNewPeerToPeers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastNewPeerToPeers.Marshal(b, m, deterministic)
}
func (dst *BroadcastNewPeerToPeers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastNewPeerToPeers.Merge(dst, src)
}
func (m *BroadcastNewPeerToPeers) XXX_Size() int {
	return xxx_messageInfo_BroadcastNewPeerToPeers.Size(m)
}
func (m *BroadcastNewPeerToPeers) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastNewPeerToPeers.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastNewPeerToPeers proto.InternalMessageInfo

func (m *BroadcastNewPeerToPeers) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *BroadcastNewPeerToPeers) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *BroadcastNewPeerToPeers) GetGuid() int64 {
	if m != nil {
		return m.Guid
	}
	return 0
}

func (m *BroadcastNewPeerToPeers) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*BroadcastNewPeerToPeers)(nil), "protocols.BroadcastNewPeerToPeers")
}

func init() {
	proto.RegisterFile("broadcast_new_peer_to_peers.proto", fileDescriptor_broadcast_new_peer_to_peers_cab8d60bbe394b90)
}

var fileDescriptor_broadcast_new_peer_to_peers_cab8d60bbe394b90 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2a, 0xca, 0x4f,
	0x4c, 0x49, 0x4e, 0x2c, 0x2e, 0x89, 0xcf, 0x4b, 0x2d, 0x8f, 0x2f, 0x48, 0x4d, 0x2d, 0x8a, 0x2f,
	0xc9, 0x07, 0xd3, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x1c, 0xb9, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0xa9, 0xc5, 0x4a, 0xd9, 0x5c, 0xe2, 0x4e, 0x30, 0xe5, 0x7e, 0xa9, 0xe5, 0x01, 0xa9,
	0xa9, 0x45, 0x21, 0xf9, 0x20, 0xb2, 0x58, 0x88, 0x8f, 0x8b, 0x29, 0xb3, 0x40, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x27, 0x88, 0x29, 0xb3, 0x40, 0x48, 0x88, 0x8b, 0xa5, 0x20, 0xbf, 0xa8, 0x44, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x37, 0x08, 0xcc, 0x06, 0x89, 0xa5, 0x97, 0x66, 0xa6, 0x48, 0x30, 0x2b,
	0x30, 0x6a, 0x30, 0x07, 0x81, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9,
	0x79, 0x12, 0x2c, 0x60, 0xa5, 0x30, 0x6e, 0x12, 0x1b, 0xd8, 0x76, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0x2b, 0x77, 0xb7, 0xa2, 0x00, 0x00, 0x00,
}
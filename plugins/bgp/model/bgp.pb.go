// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bgp.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PeerConf_RemovePrivateAs int32

const (
	PeerConf_NONE    PeerConf_RemovePrivateAs = 0
	PeerConf_ALL     PeerConf_RemovePrivateAs = 1
	PeerConf_REPLACE PeerConf_RemovePrivateAs = 2
)

var PeerConf_RemovePrivateAs_name = map[int32]string{
	0: "NONE",
	1: "ALL",
	2: "REPLACE",
}

var PeerConf_RemovePrivateAs_value = map[string]int32{
	"NONE":    0,
	"ALL":     1,
	"REPLACE": 2,
}

func (x PeerConf_RemovePrivateAs) String() string {
	return proto.EnumName(PeerConf_RemovePrivateAs_name, int32(x))
}

func (PeerConf_RemovePrivateAs) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{2, 0}
}

//BGP configuration
type BgpConf struct {
	Global               *GlobalConf `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Peers                []*PeerConf `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BgpConf) Reset()         { *m = BgpConf{} }
func (m *BgpConf) String() string { return proto.CompactTextString(m) }
func (*BgpConf) ProtoMessage()    {}
func (*BgpConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{0}
}
func (m *BgpConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConf.Unmarshal(m, b)
}
func (m *BgpConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConf.Marshal(b, m, deterministic)
}
func (m *BgpConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConf.Merge(m, src)
}
func (m *BgpConf) XXX_Size() int {
	return xxx_messageInfo_BgpConf.Size(m)
}
func (m *BgpConf) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConf.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConf proto.InternalMessageInfo

func (m *BgpConf) GetGlobal() *GlobalConf {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *BgpConf) GetPeers() []*PeerConf {
	if m != nil {
		return m.Peers
	}
	return nil
}

// global configuration
type GlobalConf struct {
	As                   uint32   `protobuf:"varint,1,opt,name=as,proto3" json:"as,omitempty"`
	RouterId             string   `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	ListenPort           int32    `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	ListenAddresses      []string `protobuf:"bytes,4,rep,name=listen_addresses,json=listenAddresses,proto3" json:"listen_addresses,omitempty"`
	Families             []uint32 `protobuf:"varint,5,rep,packed,name=families,proto3" json:"families,omitempty"`
	UseMultiplePaths     bool     `protobuf:"varint,6,opt,name=use_multiple_paths,json=useMultiplePaths,proto3" json:"use_multiple_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalConf) Reset()         { *m = GlobalConf{} }
func (m *GlobalConf) String() string { return proto.CompactTextString(m) }
func (*GlobalConf) ProtoMessage()    {}
func (*GlobalConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{1}
}
func (m *GlobalConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConf.Unmarshal(m, b)
}
func (m *GlobalConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConf.Marshal(b, m, deterministic)
}
func (m *GlobalConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConf.Merge(m, src)
}
func (m *GlobalConf) XXX_Size() int {
	return xxx_messageInfo_GlobalConf.Size(m)
}
func (m *GlobalConf) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConf.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConf proto.InternalMessageInfo

func (m *GlobalConf) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *GlobalConf) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *GlobalConf) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *GlobalConf) GetListenAddresses() []string {
	if m != nil {
		return m.ListenAddresses
	}
	return nil
}

func (m *GlobalConf) GetFamilies() []uint32 {
	if m != nil {
		return m.Families
	}
	return nil
}

func (m *GlobalConf) GetUseMultiplePaths() bool {
	if m != nil {
		return m.UseMultiplePaths
	}
	return false
}

// neighbor configuration, one struct will be created per peer
type PeerConf struct {
	Name                 string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AuthPassword         string                   `protobuf:"bytes,2,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	Description          string                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	LocalAs              uint32                   `protobuf:"varint,4,opt,name=local_as,json=localAs,proto3" json:"local_as,omitempty"`
	NeighborAddress      string                   `protobuf:"bytes,5,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	PeerAs               uint32                   `protobuf:"varint,6,opt,name=peer_as,json=peerAs,proto3" json:"peer_as,omitempty"`
	PeerGroup            string                   `protobuf:"bytes,7,opt,name=peer_group,json=peerGroup,proto3" json:"peer_group,omitempty"`
	PeerType             uint32                   `protobuf:"varint,8,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	RemovePrivateAs      PeerConf_RemovePrivateAs `protobuf:"varint,9,opt,name=remove_private_as,json=removePrivateAs,proto3,enum=model.PeerConf_RemovePrivateAs" json:"remove_private_as,omitempty"`
	RouteFlapDamping     bool                     `protobuf:"varint,10,opt,name=route_flap_damping,json=routeFlapDamping,proto3" json:"route_flap_damping,omitempty"`
	SendCommunity        uint32                   `protobuf:"varint,11,opt,name=send_community,json=sendCommunity,proto3" json:"send_community,omitempty"`
	NeighborInterface    string                   `protobuf:"bytes,12,opt,name=neighbor_interface,json=neighborInterface,proto3" json:"neighbor_interface,omitempty"`
	Vrf                  string                   `protobuf:"bytes,13,opt,name=vrf,proto3" json:"vrf,omitempty"`
	AllowOwnAs           uint32                   `protobuf:"varint,14,opt,name=allow_own_as,json=allowOwnAs,proto3" json:"allow_own_as,omitempty"`
	ReplacePeerAs        bool                     `protobuf:"varint,15,opt,name=replace_peer_as,json=replacePeerAs,proto3" json:"replace_peer_as,omitempty"`
	AdminDown            bool                     `protobuf:"varint,16,opt,name=admin_down,json=adminDown,proto3" json:"admin_down,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PeerConf) Reset()         { *m = PeerConf{} }
func (m *PeerConf) String() string { return proto.CompactTextString(m) }
func (*PeerConf) ProtoMessage()    {}
func (*PeerConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{2}
}
func (m *PeerConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerConf.Unmarshal(m, b)
}
func (m *PeerConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerConf.Marshal(b, m, deterministic)
}
func (m *PeerConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerConf.Merge(m, src)
}
func (m *PeerConf) XXX_Size() int {
	return xxx_messageInfo_PeerConf.Size(m)
}
func (m *PeerConf) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerConf.DiscardUnknown(m)
}

var xxx_messageInfo_PeerConf proto.InternalMessageInfo

func (m *PeerConf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PeerConf) GetAuthPassword() string {
	if m != nil {
		return m.AuthPassword
	}
	return ""
}

func (m *PeerConf) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PeerConf) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *PeerConf) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *PeerConf) GetPeerAs() uint32 {
	if m != nil {
		return m.PeerAs
	}
	return 0
}

func (m *PeerConf) GetPeerGroup() string {
	if m != nil {
		return m.PeerGroup
	}
	return ""
}

func (m *PeerConf) GetPeerType() uint32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *PeerConf) GetRemovePrivateAs() PeerConf_RemovePrivateAs {
	if m != nil {
		return m.RemovePrivateAs
	}
	return PeerConf_NONE
}

func (m *PeerConf) GetRouteFlapDamping() bool {
	if m != nil {
		return m.RouteFlapDamping
	}
	return false
}

func (m *PeerConf) GetSendCommunity() uint32 {
	if m != nil {
		return m.SendCommunity
	}
	return 0
}

func (m *PeerConf) GetNeighborInterface() string {
	if m != nil {
		return m.NeighborInterface
	}
	return ""
}

func (m *PeerConf) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *PeerConf) GetAllowOwnAs() uint32 {
	if m != nil {
		return m.AllowOwnAs
	}
	return 0
}

func (m *PeerConf) GetReplacePeerAs() bool {
	if m != nil {
		return m.ReplacePeerAs
	}
	return false
}

func (m *PeerConf) GetAdminDown() bool {
	if m != nil {
		return m.AdminDown
	}
	return false
}

func init() {
	proto.RegisterEnum("model.PeerConf_RemovePrivateAs", PeerConf_RemovePrivateAs_name, PeerConf_RemovePrivateAs_value)
	proto.RegisterType((*BgpConf)(nil), "model.BgpConf")
	proto.RegisterType((*GlobalConf)(nil), "model.GlobalConf")
	proto.RegisterType((*PeerConf)(nil), "model.PeerConf")
}

func init() { proto.RegisterFile("bgp.proto", fileDescriptor_2e12fe45eec524a6) }

var fileDescriptor_2e12fe45eec524a6 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x5f, 0x6f, 0xd3, 0x3a,
	0x18, 0xc6, 0x4f, 0xfa, 0x37, 0x79, 0xbb, 0xb6, 0xd9, 0x7b, 0x73, 0x7c, 0x0e, 0x42, 0x8b, 0x8a,
	0x86, 0x32, 0x09, 0x2a, 0x31, 0x3e, 0x41, 0xd8, 0xc6, 0x34, 0x31, 0xb6, 0x28, 0xe2, 0x8e, 0x8b,
	0xc8, 0x6d, 0xdc, 0x2e, 0x92, 0x63, 0x5b, 0x76, 0xb2, 0x6a, 0x9f, 0x92, 0x8f, 0x04, 0xb2, 0x93,
	0x0c, 0xd8, 0x9d, 0xfd, 0x7b, 0x1e, 0xbf, 0xf5, 0xf3, 0xb8, 0x81, 0x60, 0xb3, 0x57, 0x6b, 0xa5,
	0x65, 0x2d, 0x71, 0x5c, 0xc9, 0x82, 0xf1, 0xd5, 0x77, 0x98, 0x7e, 0xda, 0xab, 0x0b, 0x29, 0x76,
	0x78, 0x06, 0x93, 0x3d, 0x97, 0x1b, 0xca, 0x89, 0x17, 0x79, 0xf1, 0xec, 0xfc, 0x78, 0xed, 0x2c,
	0xeb, 0x6b, 0x07, 0xad, 0x25, 0xeb, 0x0c, 0x78, 0x0a, 0x63, 0xc5, 0x98, 0x36, 0x64, 0x10, 0x0d,
	0xe3, 0xd9, 0xf9, 0xb2, 0x73, 0xa6, 0x8c, 0x69, 0xe7, 0x6b, 0xd5, 0xd5, 0x0f, 0x0f, 0xe0, 0xf7,
	0x69, 0x5c, 0xc0, 0x80, 0x1a, 0x37, 0x7c, 0x9e, 0x0d, 0xa8, 0xc1, 0x57, 0x10, 0x68, 0xd9, 0xd4,
	0x4c, 0xe7, 0x65, 0x41, 0x06, 0x91, 0x17, 0x07, 0x99, 0xdf, 0x82, 0x9b, 0x02, 0x4f, 0x60, 0xc6,
	0x4b, 0x53, 0x33, 0x91, 0x2b, 0xa9, 0x6b, 0x32, 0x8c, 0xbc, 0x78, 0x9c, 0x41, 0x8b, 0x52, 0xa9,
	0x6b, 0x3c, 0x83, 0xb0, 0x33, 0xd0, 0xa2, 0xd0, 0xcc, 0x18, 0x66, 0xc8, 0x28, 0x1a, 0xc6, 0x41,
	0xb6, 0x6c, 0x79, 0xd2, 0x63, 0xfc, 0x1f, 0xfc, 0x1d, 0xad, 0x4a, 0x5e, 0x32, 0x43, 0xc6, 0xd1,
	0x30, 0x9e, 0x67, 0xcf, 0x7b, 0x7c, 0x07, 0xd8, 0x18, 0x96, 0x57, 0x0d, 0xaf, 0x4b, 0xc5, 0x59,
	0xae, 0x68, 0xfd, 0x60, 0xc8, 0x24, 0xf2, 0x62, 0x3f, 0x0b, 0x1b, 0xc3, 0xbe, 0x76, 0x42, 0x6a,
	0xf9, 0xea, 0xe7, 0x08, 0xfc, 0x3e, 0x25, 0x22, 0x8c, 0x04, 0xad, 0x98, 0x4b, 0x14, 0x64, 0x6e,
	0x8d, 0x6f, 0x60, 0x4e, 0x9b, 0xfa, 0x21, 0x57, 0xd4, 0x98, 0x83, 0xd4, 0x7d, 0xae, 0x23, 0x0b,
	0xd3, 0x8e, 0x61, 0x04, 0xb3, 0x82, 0x99, 0xad, 0x2e, 0x55, 0x5d, 0x4a, 0xe1, 0xb2, 0x05, 0xd9,
	0x9f, 0x08, 0xff, 0x03, 0x9f, 0xcb, 0x2d, 0xe5, 0x39, 0xb5, 0xa1, 0x6c, 0x61, 0x53, 0xb7, 0x4f,
	0x8c, 0xcd, 0x2d, 0x58, 0xb9, 0x7f, 0xd8, 0x48, 0xdd, 0x27, 0x27, 0x63, 0x37, 0x61, 0xd9, 0xf3,
	0x2e, 0x39, 0xfe, 0x0b, 0x53, 0xfb, 0x10, 0x76, 0xc8, 0xc4, 0x0d, 0x99, 0xd8, 0x6d, 0x62, 0xf0,
	0x35, 0x80, 0x13, 0xf6, 0x5a, 0x36, 0x8a, 0x4c, 0xdd, 0xe9, 0xc0, 0x92, 0x6b, 0x0b, 0xec, 0xc3,
	0x38, 0xb9, 0x7e, 0x52, 0x8c, 0xf8, 0xee, 0xa4, 0x6f, 0xc1, 0xb7, 0x27, 0xc5, 0xf0, 0x0b, 0x1c,
	0x6b, 0x56, 0xc9, 0x47, 0x96, 0x2b, 0x5d, 0x3e, 0xd2, 0x9a, 0xd9, 0xf1, 0x41, 0xe4, 0xc5, 0x8b,
	0xf3, 0x93, 0x17, 0xff, 0x83, 0x75, 0xe6, 0x8c, 0x69, 0xeb, 0x4b, 0x4c, 0xb6, 0xd4, 0x7f, 0x03,
	0xdb, 0xbe, 0x7b, 0xf1, 0x7c, 0xc7, 0xa9, 0xca, 0x0b, 0x5a, 0xa9, 0x52, 0xec, 0x09, 0xb4, 0xed,
	0x3b, 0xe5, 0x33, 0xa7, 0xea, 0xb2, 0xe5, 0x78, 0x0a, 0x0b, 0xc3, 0x44, 0x91, 0x6f, 0x65, 0x55,
	0x35, 0xa2, 0xac, 0x9f, 0xc8, 0xcc, 0x5d, 0x6e, 0x6e, 0xe9, 0x45, 0x0f, 0xf1, 0x3d, 0xe0, 0x73,
	0x43, 0xa5, 0xa8, 0x99, 0xde, 0xd1, 0x2d, 0x23, 0x47, 0x2e, 0xe5, 0x71, 0xaf, 0xdc, 0xf4, 0x02,
	0x86, 0x30, 0x7c, 0xd4, 0x3b, 0x32, 0x77, 0xba, 0x5d, 0x62, 0x04, 0x47, 0x94, 0x73, 0x79, 0xc8,
	0xe5, 0x41, 0xd8, 0x74, 0x0b, 0xf7, 0x2b, 0xe0, 0xd8, 0xfd, 0x41, 0x24, 0x06, 0xdf, 0xc2, 0x52,
	0x33, 0xc5, 0xe9, 0x96, 0xe5, 0x7d, 0xc3, 0x4b, 0x77, 0xe9, 0x79, 0x87, 0xd3, 0xe7, 0xa2, 0x69,
	0x51, 0x95, 0x22, 0x2f, 0xe4, 0x41, 0x90, 0xd0, 0x59, 0x02, 0x47, 0x2e, 0xe5, 0x41, 0xac, 0x3e,
	0xc0, 0xf2, 0x45, 0x45, 0xe8, 0xc3, 0xe8, 0xee, 0xfe, 0xee, 0x2a, 0xfc, 0x07, 0xa7, 0x30, 0x4c,
	0x6e, 0x6f, 0x43, 0x0f, 0x67, 0x30, 0xcd, 0xae, 0xd2, 0xdb, 0xe4, 0xe2, 0x2a, 0x1c, 0x6c, 0x26,
	0xee, 0xf3, 0xfd, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x51, 0x45, 0x32, 0xcb, 0x03, 0x00,
	0x00,
}

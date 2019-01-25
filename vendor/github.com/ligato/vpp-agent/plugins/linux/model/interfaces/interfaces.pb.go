// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interfaces.proto

package interfaces

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LinuxInterfaces_InterfaceType int32

const (
	LinuxInterfaces_VETH     LinuxInterfaces_InterfaceType = 0
	LinuxInterfaces_AUTO_TAP LinuxInterfaces_InterfaceType = 1
)

var LinuxInterfaces_InterfaceType_name = map[int32]string{
	0: "VETH",
	1: "AUTO_TAP",
}
var LinuxInterfaces_InterfaceType_value = map[string]int32{
	"VETH":     0,
	"AUTO_TAP": 1,
}

func (x LinuxInterfaces_InterfaceType) String() string {
	return proto.EnumName(LinuxInterfaces_InterfaceType_name, int32(x))
}
func (LinuxInterfaces_InterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0, 0}
}

type LinuxInterfaces_Interface_Namespace_NamespaceType int32

const (
	LinuxInterfaces_Interface_Namespace_PID_REF_NS          LinuxInterfaces_Interface_Namespace_NamespaceType = 0
	LinuxInterfaces_Interface_Namespace_MICROSERVICE_REF_NS LinuxInterfaces_Interface_Namespace_NamespaceType = 1
	LinuxInterfaces_Interface_Namespace_NAMED_NS            LinuxInterfaces_Interface_Namespace_NamespaceType = 2
	LinuxInterfaces_Interface_Namespace_FILE_REF_NS         LinuxInterfaces_Interface_Namespace_NamespaceType = 3
)

var LinuxInterfaces_Interface_Namespace_NamespaceType_name = map[int32]string{
	0: "PID_REF_NS",
	1: "MICROSERVICE_REF_NS",
	2: "NAMED_NS",
	3: "FILE_REF_NS",
}
var LinuxInterfaces_Interface_Namespace_NamespaceType_value = map[string]int32{
	"PID_REF_NS":          0,
	"MICROSERVICE_REF_NS": 1,
	"NAMED_NS":            2,
	"FILE_REF_NS":         3,
}

func (x LinuxInterfaces_Interface_Namespace_NamespaceType) String() string {
	return proto.EnumName(LinuxInterfaces_Interface_Namespace_NamespaceType_name, int32(x))
}
func (LinuxInterfaces_Interface_Namespace_NamespaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0, 0, 0, 0}
}

type LinuxInterfaces struct {
	Interface            []*LinuxInterfaces_Interface `protobuf:"bytes,1,rep,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *LinuxInterfaces) Reset()         { *m = LinuxInterfaces{} }
func (m *LinuxInterfaces) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces) ProtoMessage()    {}
func (*LinuxInterfaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0}
}
func (m *LinuxInterfaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterfaces.Unmarshal(m, b)
}
func (m *LinuxInterfaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterfaces.Marshal(b, m, deterministic)
}
func (dst *LinuxInterfaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterfaces.Merge(dst, src)
}
func (m *LinuxInterfaces) XXX_Size() int {
	return xxx_messageInfo_LinuxInterfaces.Size(m)
}
func (m *LinuxInterfaces) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterfaces.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterfaces proto.InternalMessageInfo

func (m *LinuxInterfaces) GetInterface() []*LinuxInterfaces_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type LinuxInterfaces_Interface struct {
	Name                 string                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 LinuxInterfaces_InterfaceType        `protobuf:"varint,3,opt,name=type,proto3,enum=interfaces.LinuxInterfaces_InterfaceType" json:"type,omitempty"`
	Enabled              bool                                 `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IpAddresses          []string                             `protobuf:"bytes,5,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	PhysAddress          string                               `protobuf:"bytes,6,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Mtu                  uint32                               `protobuf:"varint,7,opt,name=mtu,proto3" json:"mtu,omitempty"`
	HostIfName           string                               `protobuf:"bytes,8,opt,name=host_if_name,json=hostIfName,proto3" json:"host_if_name,omitempty"`
	Namespace            *LinuxInterfaces_Interface_Namespace `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Veth                 *LinuxInterfaces_Interface_Veth      `protobuf:"bytes,10,opt,name=veth,proto3" json:"veth,omitempty"`
	Tap                  *LinuxInterfaces_Interface_Tap       `protobuf:"bytes,11,opt,name=tap,proto3" json:"tap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *LinuxInterfaces_Interface) Reset()         { *m = LinuxInterfaces_Interface{} }
func (m *LinuxInterfaces_Interface) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface) ProtoMessage()    {}
func (*LinuxInterfaces_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0, 0}
}
func (m *LinuxInterfaces_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterfaces_Interface.Unmarshal(m, b)
}
func (m *LinuxInterfaces_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterfaces_Interface.Marshal(b, m, deterministic)
}
func (dst *LinuxInterfaces_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterfaces_Interface.Merge(dst, src)
}
func (m *LinuxInterfaces_Interface) XXX_Size() int {
	return xxx_messageInfo_LinuxInterfaces_Interface.Size(m)
}
func (m *LinuxInterfaces_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterfaces_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterfaces_Interface proto.InternalMessageInfo

func (m *LinuxInterfaces_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LinuxInterfaces_Interface) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LinuxInterfaces_Interface) GetType() LinuxInterfaces_InterfaceType {
	if m != nil {
		return m.Type
	}
	return LinuxInterfaces_VETH
}

func (m *LinuxInterfaces_Interface) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *LinuxInterfaces_Interface) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

func (m *LinuxInterfaces_Interface) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *LinuxInterfaces_Interface) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *LinuxInterfaces_Interface) GetHostIfName() string {
	if m != nil {
		return m.HostIfName
	}
	return ""
}

func (m *LinuxInterfaces_Interface) GetNamespace() *LinuxInterfaces_Interface_Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *LinuxInterfaces_Interface) GetVeth() *LinuxInterfaces_Interface_Veth {
	if m != nil {
		return m.Veth
	}
	return nil
}

func (m *LinuxInterfaces_Interface) GetTap() *LinuxInterfaces_Interface_Tap {
	if m != nil {
		return m.Tap
	}
	return nil
}

type LinuxInterfaces_Interface_Namespace struct {
	Type                 LinuxInterfaces_Interface_Namespace_NamespaceType `protobuf:"varint,1,opt,name=type,proto3,enum=interfaces.LinuxInterfaces_Interface_Namespace_NamespaceType" json:"type,omitempty"`
	Pid                  uint32                                            `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Microservice         string                                            `protobuf:"bytes,3,opt,name=microservice,proto3" json:"microservice,omitempty"`
	Name                 string                                            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Filepath             string                                            `protobuf:"bytes,5,opt,name=filepath,proto3" json:"filepath,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *LinuxInterfaces_Interface_Namespace) Reset()         { *m = LinuxInterfaces_Interface_Namespace{} }
func (m *LinuxInterfaces_Interface_Namespace) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface_Namespace) ProtoMessage()    {}
func (*LinuxInterfaces_Interface_Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0, 0, 0}
}
func (m *LinuxInterfaces_Interface_Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterfaces_Interface_Namespace.Unmarshal(m, b)
}
func (m *LinuxInterfaces_Interface_Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterfaces_Interface_Namespace.Marshal(b, m, deterministic)
}
func (dst *LinuxInterfaces_Interface_Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterfaces_Interface_Namespace.Merge(dst, src)
}
func (m *LinuxInterfaces_Interface_Namespace) XXX_Size() int {
	return xxx_messageInfo_LinuxInterfaces_Interface_Namespace.Size(m)
}
func (m *LinuxInterfaces_Interface_Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterfaces_Interface_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterfaces_Interface_Namespace proto.InternalMessageInfo

func (m *LinuxInterfaces_Interface_Namespace) GetType() LinuxInterfaces_Interface_Namespace_NamespaceType {
	if m != nil {
		return m.Type
	}
	return LinuxInterfaces_Interface_Namespace_PID_REF_NS
}

func (m *LinuxInterfaces_Interface_Namespace) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *LinuxInterfaces_Interface_Namespace) GetMicroservice() string {
	if m != nil {
		return m.Microservice
	}
	return ""
}

func (m *LinuxInterfaces_Interface_Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LinuxInterfaces_Interface_Namespace) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

type LinuxInterfaces_Interface_Veth struct {
	PeerIfName           string   `protobuf:"bytes,1,opt,name=peer_if_name,json=peerIfName,proto3" json:"peer_if_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinuxInterfaces_Interface_Veth) Reset()         { *m = LinuxInterfaces_Interface_Veth{} }
func (m *LinuxInterfaces_Interface_Veth) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface_Veth) ProtoMessage()    {}
func (*LinuxInterfaces_Interface_Veth) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0, 0, 1}
}
func (m *LinuxInterfaces_Interface_Veth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterfaces_Interface_Veth.Unmarshal(m, b)
}
func (m *LinuxInterfaces_Interface_Veth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterfaces_Interface_Veth.Marshal(b, m, deterministic)
}
func (dst *LinuxInterfaces_Interface_Veth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterfaces_Interface_Veth.Merge(dst, src)
}
func (m *LinuxInterfaces_Interface_Veth) XXX_Size() int {
	return xxx_messageInfo_LinuxInterfaces_Interface_Veth.Size(m)
}
func (m *LinuxInterfaces_Interface_Veth) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterfaces_Interface_Veth.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterfaces_Interface_Veth proto.InternalMessageInfo

func (m *LinuxInterfaces_Interface_Veth) GetPeerIfName() string {
	if m != nil {
		return m.PeerIfName
	}
	return ""
}

type LinuxInterfaces_Interface_Tap struct {
	TempIfName           string   `protobuf:"bytes,1,opt,name=temp_if_name,json=tempIfName,proto3" json:"temp_if_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinuxInterfaces_Interface_Tap) Reset()         { *m = LinuxInterfaces_Interface_Tap{} }
func (m *LinuxInterfaces_Interface_Tap) String() string { return proto.CompactTextString(m) }
func (*LinuxInterfaces_Interface_Tap) ProtoMessage()    {}
func (*LinuxInterfaces_Interface_Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e1580d882dd65c0f, []int{0, 0, 2}
}
func (m *LinuxInterfaces_Interface_Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterfaces_Interface_Tap.Unmarshal(m, b)
}
func (m *LinuxInterfaces_Interface_Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterfaces_Interface_Tap.Marshal(b, m, deterministic)
}
func (dst *LinuxInterfaces_Interface_Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterfaces_Interface_Tap.Merge(dst, src)
}
func (m *LinuxInterfaces_Interface_Tap) XXX_Size() int {
	return xxx_messageInfo_LinuxInterfaces_Interface_Tap.Size(m)
}
func (m *LinuxInterfaces_Interface_Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterfaces_Interface_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterfaces_Interface_Tap proto.InternalMessageInfo

func (m *LinuxInterfaces_Interface_Tap) GetTempIfName() string {
	if m != nil {
		return m.TempIfName
	}
	return ""
}

func init() {
	proto.RegisterType((*LinuxInterfaces)(nil), "interfaces.LinuxInterfaces")
	proto.RegisterType((*LinuxInterfaces_Interface)(nil), "interfaces.LinuxInterfaces.Interface")
	proto.RegisterType((*LinuxInterfaces_Interface_Namespace)(nil), "interfaces.LinuxInterfaces.Interface.Namespace")
	proto.RegisterType((*LinuxInterfaces_Interface_Veth)(nil), "interfaces.LinuxInterfaces.Interface.Veth")
	proto.RegisterType((*LinuxInterfaces_Interface_Tap)(nil), "interfaces.LinuxInterfaces.Interface.Tap")
	proto.RegisterEnum("interfaces.LinuxInterfaces_InterfaceType", LinuxInterfaces_InterfaceType_name, LinuxInterfaces_InterfaceType_value)
	proto.RegisterEnum("interfaces.LinuxInterfaces_Interface_Namespace_NamespaceType", LinuxInterfaces_Interface_Namespace_NamespaceType_name, LinuxInterfaces_Interface_Namespace_NamespaceType_value)
}

func init() { proto.RegisterFile("interfaces.proto", fileDescriptor_interfaces_e1580d882dd65c0f) }

var fileDescriptor_interfaces_e1580d882dd65c0f = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0x9b, 0x40,
	0x10, 0x0d, 0x86, 0xd8, 0x30, 0xd8, 0x09, 0xda, 0x1e, 0xba, 0xe2, 0x44, 0x2d, 0x55, 0xa1, 0x3d,
	0x50, 0xc9, 0x3d, 0x56, 0xa9, 0x64, 0x39, 0x44, 0x45, 0x8a, 0x9d, 0x74, 0x43, 0xdd, 0x23, 0x22,
	0x66, 0x2d, 0x56, 0x8a, 0x61, 0x05, 0x9b, 0xa8, 0xfe, 0x98, 0x7e, 0x55, 0xd5, 0xff, 0xa9, 0x76,
	0x09, 0xd8, 0xce, 0xc9, 0xbd, 0xcd, 0xbe, 0x79, 0x6f, 0x76, 0x66, 0xe7, 0x2d, 0x38, 0xac, 0x10,
	0xb4, 0x5a, 0xa7, 0x2b, 0x5a, 0x07, 0xbc, 0x2a, 0x45, 0x89, 0x60, 0x87, 0x8c, 0xff, 0x0e, 0xe0,
	0xfc, 0x86, 0x15, 0x4f, 0xbf, 0xa2, 0x0e, 0x43, 0x33, 0xb0, 0x3a, 0x06, 0xd6, 0x3c, 0xdd, 0xb7,
	0x27, 0xef, 0x83, 0xbd, 0x2a, 0xaf, 0xf8, 0x41, 0x17, 0x92, 0x9d, 0xce, 0xfd, 0xd3, 0x07, 0xab,
	0x4b, 0x20, 0x04, 0x46, 0x91, 0x6e, 0x64, 0x35, 0xcd, 0xb7, 0x88, 0x8a, 0x91, 0x07, 0x76, 0x46,
	0xeb, 0x55, 0xc5, 0xb8, 0x60, 0x65, 0x81, 0x7b, 0x2a, 0xb5, 0x0f, 0xa1, 0x4b, 0x30, 0xc4, 0x96,
	0x53, 0xac, 0x7b, 0x9a, 0x7f, 0x36, 0xf9, 0x70, 0x54, 0x0f, 0xf1, 0x96, 0x53, 0xa2, 0x64, 0x08,
	0xc3, 0x80, 0x16, 0xe9, 0xc3, 0x23, 0xcd, 0xb0, 0xe1, 0x69, 0xbe, 0x49, 0xda, 0x23, 0x7a, 0x07,
	0x43, 0xc6, 0x93, 0x34, 0xcb, 0x2a, 0x5a, 0xd7, 0xb4, 0xc6, 0xa7, 0x9e, 0x2e, 0xef, 0x66, 0x7c,
	0xda, 0x42, 0x92, 0xc2, 0xf3, 0x6d, 0xdd, 0x92, 0x70, 0xbf, 0x69, 0x4f, 0x62, 0x2f, 0x24, 0xe4,
	0x80, 0xbe, 0x11, 0x4f, 0x78, 0xe0, 0x69, 0xfe, 0x88, 0xc8, 0x10, 0x79, 0x30, 0xcc, 0xcb, 0x5a,
	0x24, 0x6c, 0x9d, 0xa8, 0x71, 0x4d, 0x25, 0x02, 0x89, 0x45, 0xeb, 0x85, 0x1c, 0x7a, 0x0e, 0x96,
	0xcc, 0xd4, 0x5c, 0xbe, 0xad, 0xe5, 0x69, 0xbe, 0x3d, 0xf9, 0x74, 0xd4, 0x5c, 0xc1, 0xa2, 0x95,
	0x91, 0x5d, 0x05, 0xf4, 0x15, 0x8c, 0x67, 0x2a, 0x72, 0x0c, 0xaa, 0xd2, 0xc7, 0xe3, 0x2a, 0x2d,
	0xa9, 0xc8, 0x89, 0xd2, 0xa1, 0x2f, 0xa0, 0x8b, 0x94, 0x63, 0x5b, 0xc9, 0x8f, 0x7b, 0xe0, 0x20,
	0x4e, 0x39, 0x91, 0x2a, 0xf7, 0x77, 0x0f, 0xac, 0xae, 0x2b, 0xf4, 0xfd, 0x65, 0x59, 0x9a, 0x5a,
	0xd6, 0xe5, 0x7f, 0x0e, 0xb5, 0x8b, 0xf6, 0x16, 0xe8, 0x80, 0xce, 0x59, 0xa6, 0x9c, 0x31, 0x22,
	0x32, 0x44, 0x63, 0x18, 0x6e, 0xd8, 0xaa, 0x2a, 0x6b, 0x5a, 0x3d, 0xb3, 0x55, 0xe3, 0x0c, 0x8b,
	0x1c, 0x60, 0x9d, 0xd7, 0x8c, 0x3d, 0xaf, 0xb9, 0x60, 0xae, 0xd9, 0x23, 0xe5, 0xa9, 0xc8, 0xf1,
	0xa9, 0xc2, 0xbb, 0xf3, 0xf8, 0x27, 0x8c, 0x0e, 0x2e, 0x47, 0x67, 0x00, 0x77, 0xd1, 0x55, 0x42,
	0xc2, 0xeb, 0x64, 0x71, 0xef, 0x9c, 0xa0, 0xb7, 0xf0, 0x66, 0x1e, 0xcd, 0xc8, 0xed, 0x7d, 0x48,
	0x96, 0xd1, 0x2c, 0x6c, 0x13, 0x1a, 0x1a, 0x82, 0xb9, 0x98, 0xce, 0xc3, 0x2b, 0x79, 0xea, 0xa1,
	0x73, 0xb0, 0xaf, 0xa3, 0x9b, 0x2e, 0xad, 0xbb, 0x3e, 0x18, 0xf2, 0xa9, 0xa5, 0x2b, 0x38, 0xa5,
	0x55, 0xe7, 0x8a, 0xe6, 0x13, 0x80, 0xc4, 0x1a, 0x57, 0xb8, 0x17, 0xa0, 0xc7, 0x29, 0x97, 0x44,
	0x41, 0x37, 0xfc, 0x35, 0x51, 0x62, 0x0d, 0x71, 0x7c, 0x01, 0xa3, 0x03, 0xa7, 0x23, 0x13, 0x8c,
	0x65, 0x18, 0x7f, 0x73, 0x4e, 0x64, 0x33, 0xd3, 0x1f, 0xf1, 0x6d, 0x12, 0x4f, 0xef, 0x1c, 0xed,
	0xa1, 0xaf, 0xbe, 0xfa, 0xe7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x5b, 0x72, 0x29, 0xfe,
	0x03, 0x00, 0x00,
}

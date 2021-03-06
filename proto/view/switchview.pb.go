// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/proto/view/switchview.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SwitchView struct {
	Value   bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Enabled bool `protobuf:"varint,2,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *SwitchView) Reset()                    { *m = SwitchView{} }
func (m *SwitchView) String() string            { return proto.CompactTextString(m) }
func (*SwitchView) ProtoMessage()               {}
func (*SwitchView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SwitchView) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func (m *SwitchView) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type SwitchEvent struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *SwitchEvent) Reset()                    { *m = SwitchEvent{} }
func (m *SwitchEvent) String() string            { return proto.CompactTextString(m) }
func (*SwitchEvent) ProtoMessage()               {}
func (*SwitchEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SwitchEvent) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*SwitchView)(nil), "matcha.view.SwitchView")
	proto.RegisterType((*SwitchEvent)(nil), "matcha.view.SwitchEvent")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/proto/view/switchview.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2,
	0xf5, 0xcb, 0x32, 0x53, 0xcb, 0xf5, 0x8b, 0xcb, 0x33, 0x4b, 0x92, 0x33, 0x40, 0x4c, 0x3d, 0xb0,
	0xa8, 0x10, 0x37, 0x54, 0x35, 0x48, 0x48, 0xc9, 0x86, 0x8b, 0x2b, 0x18, 0xac, 0x20, 0x2c, 0x33,
	0xb5, 0x5c, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x23, 0x08, 0xc2, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0xcd, 0x4b, 0x4c, 0xca, 0x49, 0x4d, 0x91, 0x60,
	0x02, 0x8b, 0xc3, 0xb8, 0x4a, 0xca, 0x5c, 0xdc, 0x10, 0xdd, 0xae, 0x65, 0xa9, 0x79, 0x25, 0xd8,
	0xb5, 0x3b, 0x39, 0x71, 0xc9, 0x66, 0xe6, 0x23, 0x1c, 0x09, 0xa5, 0xc0, 0x6e, 0x01, 0xbb, 0xc1,
	0x89, 0x27, 0x20, 0x09, 0xe1, 0x86, 0x28, 0x16, 0x90, 0xd8, 0x22, 0x26, 0x1e, 0x5f, 0xb0, 0x3a,
	0x90, 0x50, 0x40, 0x52, 0x12, 0x1b, 0x58, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x54,
	0x17, 0xa7, 0xec, 0x00, 0x00, 0x00,
}

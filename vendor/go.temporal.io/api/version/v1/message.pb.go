// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/version/v1/message.proto

package version

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	v1 "go.temporal.io/api/enums/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ReleaseInfo contains information about specific version of temporal.
type ReleaseInfo struct {
	Version     string     `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ReleaseTime *time.Time `protobuf:"bytes,2,opt,name=release_time,json=releaseTime,proto3,stdtime" json:"release_time,omitempty"`
	Notes       string     `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (m *ReleaseInfo) Reset()      { *m = ReleaseInfo{} }
func (*ReleaseInfo) ProtoMessage() {}
func (*ReleaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_48bb9b63ea24e063, []int{0}
}
func (m *ReleaseInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReleaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReleaseInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReleaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseInfo.Merge(m, src)
}
func (m *ReleaseInfo) XXX_Size() int {
	return m.Size()
}
func (m *ReleaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseInfo proto.InternalMessageInfo

func (m *ReleaseInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ReleaseInfo) GetReleaseTime() *time.Time {
	if m != nil {
		return m.ReleaseTime
	}
	return nil
}

func (m *ReleaseInfo) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

// Alert contains notification and severity.
type Alert struct {
	Message  string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Severity v1.Severity `protobuf:"varint,2,opt,name=severity,proto3,enum=temporal.api.enums.v1.Severity" json:"severity,omitempty"`
}

func (m *Alert) Reset()      { *m = Alert{} }
func (*Alert) ProtoMessage() {}
func (*Alert) Descriptor() ([]byte, []int) {
	return fileDescriptor_48bb9b63ea24e063, []int{1}
}
func (m *Alert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Alert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Alert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Alert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alert.Merge(m, src)
}
func (m *Alert) XXX_Size() int {
	return m.Size()
}
func (m *Alert) XXX_DiscardUnknown() {
	xxx_messageInfo_Alert.DiscardUnknown(m)
}

var xxx_messageInfo_Alert proto.InternalMessageInfo

func (m *Alert) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Alert) GetSeverity() v1.Severity {
	if m != nil {
		return m.Severity
	}
	return v1.SEVERITY_UNSPECIFIED
}

// VersionInfo contains details about current and recommended release versions as well as alerts and upgrade instructions.
type VersionInfo struct {
	Current        *ReleaseInfo `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Recommended    *ReleaseInfo `protobuf:"bytes,2,opt,name=recommended,proto3" json:"recommended,omitempty"`
	Instructions   string       `protobuf:"bytes,3,opt,name=instructions,proto3" json:"instructions,omitempty"`
	Alerts         []*Alert     `protobuf:"bytes,4,rep,name=alerts,proto3" json:"alerts,omitempty"`
	LastUpdateTime *time.Time   `protobuf:"bytes,5,opt,name=last_update_time,json=lastUpdateTime,proto3,stdtime" json:"last_update_time,omitempty"`
}

func (m *VersionInfo) Reset()      { *m = VersionInfo{} }
func (*VersionInfo) ProtoMessage() {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_48bb9b63ea24e063, []int{2}
}
func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(m, src)
}
func (m *VersionInfo) XXX_Size() int {
	return m.Size()
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

func (m *VersionInfo) GetCurrent() *ReleaseInfo {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *VersionInfo) GetRecommended() *ReleaseInfo {
	if m != nil {
		return m.Recommended
	}
	return nil
}

func (m *VersionInfo) GetInstructions() string {
	if m != nil {
		return m.Instructions
	}
	return ""
}

func (m *VersionInfo) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

func (m *VersionInfo) GetLastUpdateTime() *time.Time {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ReleaseInfo)(nil), "temporal.api.version.v1.ReleaseInfo")
	proto.RegisterType((*Alert)(nil), "temporal.api.version.v1.Alert")
	proto.RegisterType((*VersionInfo)(nil), "temporal.api.version.v1.VersionInfo")
}

func init() {
	proto.RegisterFile("temporal/api/version/v1/message.proto", fileDescriptor_48bb9b63ea24e063)
}

var fileDescriptor_48bb9b63ea24e063 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x8b, 0x13, 0x4d,
	0x10, 0x9d, 0xce, 0x6e, 0x76, 0xbf, 0xaf, 0x27, 0x2c, 0x32, 0x08, 0xc6, 0x20, 0x9d, 0x18, 0x5c,
	0xc8, 0xa9, 0x87, 0x44, 0xf0, 0x30, 0x0b, 0x42, 0x22, 0x08, 0x0a, 0xc2, 0x32, 0xae, 0x39, 0x78,
	0x70, 0x99, 0x4d, 0x6a, 0x87, 0x86, 0x4c, 0xf7, 0xd0, 0xdd, 0x09, 0x78, 0x13, 0x7f, 0xc1, 0xfe,
	0x0c, 0xf1, 0x3f, 0x78, 0xf1, 0xe4, 0x31, 0xc7, 0xbd, 0x69, 0x26, 0x17, 0xf1, 0xb4, 0x3f, 0x41,
	0xba, 0xa7, 0x47, 0x13, 0x21, 0xb0, 0xb7, 0xa9, 0xaa, 0xf7, 0xba, 0x5e, 0xbd, 0xc7, 0xe0, 0x63,
	0x0d, 0x59, 0x2e, 0x64, 0x32, 0x0b, 0x93, 0x9c, 0x85, 0x0b, 0x90, 0x8a, 0x09, 0x1e, 0x2e, 0xfa,
	0x61, 0x06, 0x4a, 0x25, 0x29, 0xd0, 0x5c, 0x0a, 0x2d, 0x82, 0x7b, 0x15, 0x8c, 0x26, 0x39, 0xa3,
	0x0e, 0x46, 0x17, 0xfd, 0x56, 0x3b, 0x15, 0x22, 0x9d, 0x41, 0x68, 0x61, 0x17, 0xf3, 0xcb, 0x50,
	0xb3, 0x0c, 0x94, 0x4e, 0xb2, 0xbc, 0x64, 0xb6, 0x1e, 0x4e, 0x21, 0x07, 0x3e, 0x05, 0x3e, 0x61,
	0xa0, 0xc2, 0x54, 0xa4, 0xc2, 0xf6, 0xed, 0x97, 0x83, 0x74, 0xb7, 0x34, 0x00, 0x9f, 0x67, 0xca,
	0x28, 0x98, 0x88, 0x2c, 0x13, 0xbc, 0xc4, 0x74, 0x3f, 0x22, 0xec, 0xc7, 0x30, 0x83, 0x44, 0xc1,
	0x0b, 0x7e, 0x29, 0x82, 0x26, 0x3e, 0x74, 0x2a, 0x9a, 0xa8, 0x83, 0x7a, 0xff, 0xc7, 0x55, 0x19,
	0x3c, 0xc3, 0x0d, 0x59, 0x02, 0xcf, 0x8d, 0x96, 0x66, 0xad, 0x83, 0x7a, 0xfe, 0xa0, 0x45, 0x4b,
	0xa1, 0xb4, 0x12, 0x4a, 0xcf, 0x2a, 0xa1, 0xa3, 0xfd, 0xab, 0xef, 0x6d, 0x14, 0xfb, 0x8e, 0x65,
	0xfa, 0xc1, 0x5d, 0x5c, 0xe7, 0x42, 0x83, 0x6a, 0xee, 0xd9, 0xc7, 0xcb, 0xa2, 0xfb, 0x0e, 0xd7,
	0x87, 0x33, 0x90, 0xda, 0x6c, 0x77, 0xfe, 0x54, 0xdb, 0x5d, 0x19, 0x9c, 0xe0, 0xff, 0x14, 0x2c,
	0x40, 0x32, 0xfd, 0xde, 0x6e, 0x3e, 0x1a, 0xb4, 0xe9, 0x96, 0x77, 0xf6, 0x3c, 0xba, 0xe8, 0xd3,
	0xd7, 0x0e, 0x16, 0xff, 0x21, 0x74, 0xbf, 0xd4, 0xb0, 0x3f, 0x2e, 0xcf, 0xb0, 0x47, 0x3e, 0xc5,
	0x87, 0x93, 0xb9, 0x94, 0xc0, 0xb5, 0x5d, 0xe3, 0x0f, 0x1e, 0xd1, 0x1d, 0x39, 0xd0, 0x0d, 0x6f,
	0xe2, 0x8a, 0x14, 0x3c, 0xc7, 0xbe, 0x04, 0x63, 0xa3, 0x09, 0x60, 0xea, 0x9c, 0xb8, 0xdd, 0x1b,
	0x9b, 0xc4, 0xa0, 0x8b, 0x1b, 0x8c, 0x2b, 0x2d, 0xe7, 0x13, 0xcd, 0x04, 0xaf, 0x4c, 0xd9, 0xea,
	0x05, 0x4f, 0xf0, 0x41, 0x62, 0xbc, 0x51, 0xcd, 0xfd, 0xce, 0x5e, 0xcf, 0x1f, 0x90, 0x9d, 0x6b,
	0xac, 0x85, 0xb1, 0x43, 0x07, 0x2f, 0xf1, 0x9d, 0x59, 0xa2, 0xf4, 0xf9, 0x3c, 0x9f, 0x26, 0xda,
	0x45, 0x56, 0xbf, 0x65, 0x64, 0x47, 0x86, 0xf9, 0xc6, 0x12, 0xcd, 0x68, 0xf4, 0x15, 0x2d, 0x57,
	0xc4, 0xbb, 0x5e, 0x11, 0xef, 0x66, 0x45, 0xd0, 0x87, 0x82, 0xa0, 0x4f, 0x05, 0x41, 0xdf, 0x0a,
	0x82, 0x96, 0x05, 0x41, 0x3f, 0x0a, 0x82, 0x7e, 0x16, 0xc4, 0xbb, 0x29, 0x08, 0xba, 0x5a, 0x13,
	0x6f, 0xb9, 0x26, 0xde, 0xf5, 0x9a, 0x78, 0xb8, 0xc5, 0xc4, 0x2e, 0xb1, 0xa3, 0xc6, 0xab, 0x32,
	0xd8, 0x53, 0xa3, 0xe1, 0x14, 0xbd, 0x3d, 0x4e, 0x37, 0xb0, 0x4c, 0xfc, 0xf3, 0xd7, 0x9c, 0xb8,
	0xcf, 0xcf, 0xb5, 0xfb, 0x67, 0x0e, 0xc4, 0x04, 0x1d, 0xe6, 0x8c, 0xba, 0x58, 0xe9, 0xb8, 0xff,
	0xab, 0xf6, 0xe0, 0xef, 0x2c, 0x8a, 0x86, 0x39, 0x8b, 0x22, 0x37, 0x8d, 0xa2, 0x71, 0xff, 0xe2,
	0xc0, 0x9e, 0xfb, 0xf8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xaf, 0xf5, 0x21, 0x9a, 0x03,
	0x00, 0x00,
}

func (this *ReleaseInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReleaseInfo)
	if !ok {
		that2, ok := that.(ReleaseInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if that1.ReleaseTime == nil {
		if this.ReleaseTime != nil {
			return false
		}
	} else if !this.ReleaseTime.Equal(*that1.ReleaseTime) {
		return false
	}
	if this.Notes != that1.Notes {
		return false
	}
	return true
}
func (this *Alert) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Alert)
	if !ok {
		that2, ok := that.(Alert)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.Severity != that1.Severity {
		return false
	}
	return true
}
func (this *VersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VersionInfo)
	if !ok {
		that2, ok := that.(VersionInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Current.Equal(that1.Current) {
		return false
	}
	if !this.Recommended.Equal(that1.Recommended) {
		return false
	}
	if this.Instructions != that1.Instructions {
		return false
	}
	if len(this.Alerts) != len(that1.Alerts) {
		return false
	}
	for i := range this.Alerts {
		if !this.Alerts[i].Equal(that1.Alerts[i]) {
			return false
		}
	}
	if that1.LastUpdateTime == nil {
		if this.LastUpdateTime != nil {
			return false
		}
	} else if !this.LastUpdateTime.Equal(*that1.LastUpdateTime) {
		return false
	}
	return true
}
func (this *ReleaseInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&version.ReleaseInfo{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "ReleaseTime: "+fmt.Sprintf("%#v", this.ReleaseTime)+",\n")
	s = append(s, "Notes: "+fmt.Sprintf("%#v", this.Notes)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Alert) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&version.Alert{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "Severity: "+fmt.Sprintf("%#v", this.Severity)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VersionInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&version.VersionInfo{")
	if this.Current != nil {
		s = append(s, "Current: "+fmt.Sprintf("%#v", this.Current)+",\n")
	}
	if this.Recommended != nil {
		s = append(s, "Recommended: "+fmt.Sprintf("%#v", this.Recommended)+",\n")
	}
	s = append(s, "Instructions: "+fmt.Sprintf("%#v", this.Instructions)+",\n")
	if this.Alerts != nil {
		s = append(s, "Alerts: "+fmt.Sprintf("%#v", this.Alerts)+",\n")
	}
	s = append(s, "LastUpdateTime: "+fmt.Sprintf("%#v", this.LastUpdateTime)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ReleaseInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReleaseInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReleaseInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Notes) > 0 {
		i -= len(m.Notes)
		copy(dAtA[i:], m.Notes)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Notes)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ReleaseTime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ReleaseTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ReleaseTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintMessage(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Alert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Alert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Alert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Severity != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Severity))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VersionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdateTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.LastUpdateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastUpdateTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintMessage(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Alerts) > 0 {
		for iNdEx := len(m.Alerts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Alerts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Instructions) > 0 {
		i -= len(m.Instructions)
		copy(dAtA[i:], m.Instructions)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Instructions)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Recommended != nil {
		{
			size, err := m.Recommended.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Current != nil {
		{
			size, err := m.Current.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReleaseInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ReleaseTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ReleaseTime)
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Notes)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Alert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Severity != 0 {
		n += 1 + sovMessage(uint64(m.Severity))
	}
	return n
}

func (m *VersionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Current != nil {
		l = m.Current.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Recommended != nil {
		l = m.Recommended.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Instructions)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Alerts) > 0 {
		for _, e := range m.Alerts {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.LastUpdateTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastUpdateTime)
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReleaseInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReleaseInfo{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`ReleaseTime:` + strings.Replace(fmt.Sprintf("%v", this.ReleaseTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`Notes:` + fmt.Sprintf("%v", this.Notes) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Alert) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Alert{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Severity:` + fmt.Sprintf("%v", this.Severity) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VersionInfo) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAlerts := "[]*Alert{"
	for _, f := range this.Alerts {
		repeatedStringForAlerts += strings.Replace(f.String(), "Alert", "Alert", 1) + ","
	}
	repeatedStringForAlerts += "}"
	s := strings.Join([]string{`&VersionInfo{`,
		`Current:` + strings.Replace(this.Current.String(), "ReleaseInfo", "ReleaseInfo", 1) + `,`,
		`Recommended:` + strings.Replace(this.Recommended.String(), "ReleaseInfo", "ReleaseInfo", 1) + `,`,
		`Instructions:` + fmt.Sprintf("%v", this.Instructions) + `,`,
		`Alerts:` + repeatedStringForAlerts + `,`,
		`LastUpdateTime:` + strings.Replace(fmt.Sprintf("%v", this.LastUpdateTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReleaseInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReleaseInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReleaseInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReleaseTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReleaseTime == nil {
				m.ReleaseTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ReleaseTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Alert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Alert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Alert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			m.Severity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Severity |= v1.Severity(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VersionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Current", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Current == nil {
				m.Current = &ReleaseInfo{}
			}
			if err := m.Current.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recommended", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Recommended == nil {
				m.Recommended = &ReleaseInfo{}
			}
			if err := m.Recommended.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instructions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instructions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alerts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alerts = append(m.Alerts, &Alert{})
			if err := m.Alerts[len(m.Alerts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastUpdateTime == nil {
				m.LastUpdateTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.LastUpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)

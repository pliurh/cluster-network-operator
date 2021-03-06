// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/openshift/api/project/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	k8s_io_api_core_v1 "k8s.io/api/core/v1"
	v11 "k8s.io/api/core/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *Project) Reset()      { *m = Project{} }
func (*Project) ProtoMessage() {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf46eaac05029bf, []int{0}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return m.Size()
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *ProjectList) Reset()      { *m = ProjectList{} }
func (*ProjectList) ProtoMessage() {}
func (*ProjectList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf46eaac05029bf, []int{1}
}
func (m *ProjectList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ProjectList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectList.Merge(m, src)
}
func (m *ProjectList) XXX_Size() int {
	return m.Size()
}
func (m *ProjectList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectList.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectList proto.InternalMessageInfo

func (m *ProjectRequest) Reset()      { *m = ProjectRequest{} }
func (*ProjectRequest) ProtoMessage() {}
func (*ProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf46eaac05029bf, []int{2}
}
func (m *ProjectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRequest.Merge(m, src)
}
func (m *ProjectRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRequest proto.InternalMessageInfo

func (m *ProjectSpec) Reset()      { *m = ProjectSpec{} }
func (*ProjectSpec) ProtoMessage() {}
func (*ProjectSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf46eaac05029bf, []int{3}
}
func (m *ProjectSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ProjectSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectSpec.Merge(m, src)
}
func (m *ProjectSpec) XXX_Size() int {
	return m.Size()
}
func (m *ProjectSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectSpec proto.InternalMessageInfo

func (m *ProjectStatus) Reset()      { *m = ProjectStatus{} }
func (*ProjectStatus) ProtoMessage() {}
func (*ProjectStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf46eaac05029bf, []int{4}
}
func (m *ProjectStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ProjectStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectStatus.Merge(m, src)
}
func (m *ProjectStatus) XXX_Size() int {
	return m.Size()
}
func (m *ProjectStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Project)(nil), "github.com.openshift.api.project.v1.Project")
	proto.RegisterType((*ProjectList)(nil), "github.com.openshift.api.project.v1.ProjectList")
	proto.RegisterType((*ProjectRequest)(nil), "github.com.openshift.api.project.v1.ProjectRequest")
	proto.RegisterType((*ProjectSpec)(nil), "github.com.openshift.api.project.v1.ProjectSpec")
	proto.RegisterType((*ProjectStatus)(nil), "github.com.openshift.api.project.v1.ProjectStatus")
}

func init() {
	proto.RegisterFile("github.com/openshift/api/project/v1/generated.proto", fileDescriptor_fbf46eaac05029bf)
}

var fileDescriptor_fbf46eaac05029bf = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0xf6, 0x7a, 0x5c, 0x5d, 0xee, 0x84, 0xc2, 0x52, 0x75, 0x48, 0x4b, 0x90, 0x50,
	0x07, 0x70, 0x68, 0x79, 0x11, 0x73, 0x40, 0x08, 0x24, 0x5e, 0x0e, 0xb3, 0x55, 0x0c, 0xb8, 0xa9,
	0x9b, 0x9a, 0x5e, 0x62, 0x13, 0xbb, 0x95, 0x8e, 0x89, 0x8f, 0xc0, 0xce, 0xe7, 0x60, 0x65, 0xee,
	0x78, 0xe3, 0x4d, 0xd5, 0x35, 0x7c, 0x8b, 0x9b, 0x90, 0x1d, 0x37, 0x09, 0x5c, 0x91, 0xee, 0x16,
	0xb6, 0xfa, 0xc9, 0xff, 0xf7, 0xb3, 0xfd, 0x3c, 0x2e, 0x78, 0x10, 0x52, 0x39, 0x9d, 0x8f, 0x60,
	0xc0, 0x22, 0x8f, 0x71, 0x12, 0x8b, 0x29, 0x9d, 0x48, 0x0f, 0x73, 0xea, 0xf1, 0x84, 0x7d, 0x22,
	0x81, 0xf4, 0x16, 0x7d, 0x2f, 0x24, 0x31, 0x49, 0xb0, 0x24, 0x63, 0xc8, 0x13, 0x26, 0x99, 0x7d,
	0xbb, 0x80, 0x60, 0x0e, 0x41, 0xcc, 0x29, 0x34, 0x10, 0x5c, 0xf4, 0xdb, 0xf7, 0x4a, 0xe6, 0x90,
	0x85, 0xcc, 0xd3, 0xec, 0x68, 0x3e, 0xd1, 0x2b, 0xbd, 0xd0, 0xbf, 0x32, 0x67, 0xdb, 0x9d, 0x3d,
	0x11, 0x90, 0x32, 0xbd, 0x75, 0xc0, 0x12, 0xb2, 0x65, 0xdf, 0xf6, 0xc3, 0x22, 0x13, 0xe1, 0x60,
	0x4a, 0x63, 0x92, 0x1c, 0x7b, 0x7c, 0x16, 0xaa, 0x82, 0xf0, 0x22, 0x22, 0xf1, 0x36, 0xea, 0xf1,
	0xbf, 0xa8, 0x64, 0x1e, 0x4b, 0x1a, 0x11, 0x4f, 0x04, 0x53, 0x12, 0xe1, 0xbf, 0x39, 0xf7, 0x7b,
	0x15, 0x5c, 0x3b, 0xcc, 0xee, 0x63, 0x7f, 0x04, 0x7b, 0x4a, 0x3f, 0xc6, 0x12, 0xb7, 0xac, 0xae,
	0xd5, 0x6b, 0x0e, 0xee, 0xc3, 0x4c, 0x0b, 0xcb, 0x5a, 0xc8, 0x67, 0xa1, 0x2a, 0x08, 0xa8, 0xd2,
	0x70, 0xd1, 0x87, 0x6f, 0x47, 0x8a, 0x7f, 0x4d, 0x24, 0xf6, 0xed, 0xe5, 0xaa, 0x53, 0x49, 0x57,
	0x1d, 0x50, 0xd4, 0x50, 0x6e, 0xb5, 0x11, 0xd8, 0x11, 0x9c, 0x04, 0xad, 0xaa, 0xb1, 0x5f, 0xa2,
	0xc5, 0xd0, 0x9c, 0xee, 0x3d, 0x27, 0x81, 0x7f, 0xdd, 0xd8, 0x77, 0xd4, 0x0a, 0x69, 0x97, 0x3d,
	0x04, 0xbb, 0x42, 0x62, 0x39, 0x17, 0xad, 0x9a, 0xb6, 0x0e, 0xae, 0x64, 0xd5, 0xa4, 0x7f, 0x60,
	0xbc, 0xbb, 0xd9, 0x1a, 0x19, 0xa3, 0xfb, 0xd3, 0x02, 0x4d, 0x93, 0x7c, 0x45, 0x85, 0xb4, 0x3f,
	0x5c, 0xe8, 0x10, 0xbc, 0x5c, 0x87, 0x14, 0xad, 0xfb, 0x73, 0xc3, 0xec, 0xb4, 0xb7, 0xa9, 0x94,
	0xba, 0xf3, 0x0e, 0xd4, 0xa9, 0x24, 0x91, 0x68, 0x55, 0xbb, 0xb5, 0x5e, 0x73, 0x70, 0xf7, 0x2a,
	0x17, 0xf1, 0xf7, 0x8d, 0xb8, 0xfe, 0x52, 0x29, 0x50, 0x66, 0x72, 0xcf, 0x2c, 0x70, 0x60, 0x12,
	0x88, 0x7c, 0x9e, 0x13, 0xf1, 0x3f, 0xa6, 0xfc, 0x08, 0x34, 0xc7, 0x54, 0xf0, 0x23, 0x7c, 0xfc,
	0x06, 0x47, 0x44, 0x0f, 0xbb, 0xe1, 0xdf, 0x34, 0x48, 0xf3, 0x59, 0xf1, 0x09, 0x95, 0x73, 0x1a,
	0x23, 0x22, 0x48, 0x28, 0x97, 0x94, 0xc5, 0x7a, 0x9a, 0x65, 0xac, 0xf8, 0x84, 0xca, 0x39, 0x17,
	0xe7, 0x23, 0x52, 0x8f, 0xc2, 0x46, 0x00, 0x4c, 0x68, 0x8c, 0x8f, 0xe8, 0x17, 0x92, 0x88, 0x96,
	0xd5, 0xad, 0xf5, 0x1a, 0xfe, 0x40, 0x1d, 0xf5, 0x79, 0x5e, 0x3d, 0x5f, 0x75, 0xba, 0x17, 0xff,
	0x88, 0x30, 0x0f, 0xe8, 0xa3, 0x95, 0x2c, 0xee, 0x0f, 0x0b, 0xec, 0xff, 0xf1, 0x60, 0xec, 0x17,
	0xa0, 0xce, 0xa7, 0x58, 0x10, 0xdd, 0xc1, 0x86, 0x3f, 0xd8, 0x34, 0xff, 0x50, 0x15, 0xcf, 0x57,
	0x9d, 0x5b, 0x5b, 0xfc, 0x4a, 0x2b, 0x38, 0x0e, 0x88, 0x0e, 0xa1, 0x4c, 0x60, 0x0f, 0x01, 0x08,
	0x58, 0x3c, 0xa6, 0xea, 0x2e, 0x9b, 0xc9, 0xdf, 0x29, 0x0d, 0x04, 0x2a, 0x1c, 0x96, 0xf1, 0xa7,
	0x9b, 0x78, 0x31, 0x86, 0xbc, 0x24, 0x50, 0xc9, 0xe6, 0xf7, 0x96, 0x6b, 0xa7, 0x72, 0xb2, 0x76,
	0x2a, 0xa7, 0x6b, 0xa7, 0xf2, 0x35, 0x75, 0xac, 0x65, 0xea, 0x58, 0x27, 0xa9, 0x63, 0x9d, 0xa6,
	0x8e, 0x75, 0x96, 0x3a, 0xd6, 0xb7, 0x5f, 0x4e, 0x65, 0x58, 0x5d, 0xf4, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0xd0, 0xf2, 0xe0, 0x22, 0x05, 0x00, 0x00,
}

func (m *Project) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Project) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Project) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ProjectList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ProjectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Description)
	copy(dAtA[i:], m.Description)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Description)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.DisplayName)
	copy(dAtA[i:], m.DisplayName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.DisplayName)))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ProjectSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Finalizers) > 0 {
		for iNdEx := len(m.Finalizers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Finalizers[iNdEx])
			copy(dAtA[i:], m.Finalizers[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Finalizers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProjectStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for iNdEx := len(m.Conditions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conditions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	i -= len(m.Phase)
	copy(dAtA[i:], m.Phase)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Phase)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Project) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ProjectList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ProjectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.DisplayName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ProjectSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Finalizers) > 0 {
		for _, s := range m.Finalizers {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ProjectStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Phase)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Project) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Project{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ProjectSpec", "ProjectSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ProjectStatus", "ProjectStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Project{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Project", "Project", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&ProjectList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProjectRequest{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProjectSpec{`,
		`Finalizers:` + fmt.Sprintf("%v", this.Finalizers) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectStatus) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForConditions := "[]NamespaceCondition{"
	for _, f := range this.Conditions {
		repeatedStringForConditions += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForConditions += "}"
	s := strings.Join([]string{`&ProjectStatus{`,
		`Phase:` + fmt.Sprintf("%v", this.Phase) + `,`,
		`Conditions:` + repeatedStringForConditions + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Project) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ProjectList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ProjectList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Project{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ProjectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ProjectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ProjectSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ProjectSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalizers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Finalizers = append(m.Finalizers, k8s_io_api_core_v1.FinalizerName(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ProjectStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ProjectStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phase = k8s_io_api_core_v1.NamespacePhase(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, v11.NamespaceCondition{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)

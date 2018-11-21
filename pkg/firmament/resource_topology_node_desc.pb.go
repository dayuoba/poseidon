/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_topology_node_desc.proto

package firmament

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

// ResourceTopologyNodeDescriptor describe a node resource topology information.
type ResourceTopologyNodeDescriptor struct {
	ResourceDesc *ResourceDescriptor `protobuf:"bytes,1,opt,name=resource_desc,json=resourceDesc,proto3" json:"resource_desc,omitempty"`
	// children stores the children nodes.
	Children []*ResourceTopologyNodeDescriptor `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
	// parent_id is the id of its parent topology node.
	ParentId             string   `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceTopologyNodeDescriptor) Reset()         { *m = ResourceTopologyNodeDescriptor{} }
func (m *ResourceTopologyNodeDescriptor) String() string { return proto.CompactTextString(m) }
func (*ResourceTopologyNodeDescriptor) ProtoMessage()    {}
func (*ResourceTopologyNodeDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7efed64a16c51499, []int{0}
}

func (m *ResourceTopologyNodeDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceTopologyNodeDescriptor.Unmarshal(m, b)
}
func (m *ResourceTopologyNodeDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceTopologyNodeDescriptor.Marshal(b, m, deterministic)
}
func (m *ResourceTopologyNodeDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceTopologyNodeDescriptor.Merge(m, src)
}
func (m *ResourceTopologyNodeDescriptor) XXX_Size() int {
	return xxx_messageInfo_ResourceTopologyNodeDescriptor.Size(m)
}
func (m *ResourceTopologyNodeDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceTopologyNodeDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceTopologyNodeDescriptor proto.InternalMessageInfo

func (m *ResourceTopologyNodeDescriptor) GetResourceDesc() *ResourceDescriptor {
	if m != nil {
		return m.ResourceDesc
	}
	return nil
}

func (m *ResourceTopologyNodeDescriptor) GetChildren() []*ResourceTopologyNodeDescriptor {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *ResourceTopologyNodeDescriptor) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceTopologyNodeDescriptor)(nil), "firmament.ResourceTopologyNodeDescriptor")
}

func init() { proto.RegisterFile("resource_topology_node_desc.proto", fileDescriptor_7efed64a16c51499) }

var fileDescriptor_7efed64a16c51499 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0x2f, 0xc9, 0x2f, 0xc8, 0xcf, 0xc9, 0x4f, 0xaf, 0x8c, 0xcf, 0xcb,
	0x4f, 0x49, 0x8d, 0x4f, 0x49, 0x2d, 0x4e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c,
	0xcb, 0x2c, 0xca, 0x4d, 0xcc, 0x4d, 0xcd, 0x2b, 0x91, 0x12, 0x86, 0xab, 0x46, 0xc8, 0x2b, 0x9d,
	0x60, 0xe4, 0x92, 0x0b, 0x82, 0x8a, 0x87, 0x40, 0x0d, 0xf1, 0xcb, 0x4f, 0x49, 0x75, 0x49, 0x2d,
	0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x12, 0x72, 0xe2, 0xe2, 0x45, 0xd1, 0x29, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xab, 0x07, 0x37, 0x5a, 0x0f, 0x66, 0x02, 0x42, 0x57, 0x10, 0x4f,
	0x11, 0x92, 0x98, 0x90, 0x2b, 0x17, 0x47, 0x72, 0x46, 0x66, 0x4e, 0x4a, 0x51, 0x6a, 0x9e, 0x04,
	0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x26, 0x16, 0xed, 0xd8, 0x1d, 0x10, 0x04, 0xd7, 0x2a, 0x24,
	0xcd, 0xc5, 0x59, 0x90, 0x58, 0x94, 0x9a, 0x57, 0x12, 0x9f, 0x99, 0x22, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x01, 0x11, 0xf0, 0x4c, 0x49, 0x62, 0x03, 0xfb, 0xc8, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x1f, 0x73, 0x43, 0x16, 0x01, 0x00, 0x00,
}

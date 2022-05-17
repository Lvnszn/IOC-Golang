/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: debug/api/ioc_golang/boot/debug.proto

package boot

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceMetadata []*ServiceMetadata `protobuf:"bytes,1,rep,name=serviceMetadata,proto3" json:"serviceMetadata,omitempty"`
}

func (x *ListServiceResponse) Reset() {
	*x = ListServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResponse) ProtoMessage() {}

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResponse.ProtoReflect.Descriptor instead.
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{0}
}

func (x *ListServiceResponse) GetServiceMetadata() []*ServiceMetadata {
	if x != nil {
		return x.ServiceMetadata
	}
	return nil
}

type ServiceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName      string   `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ImplementationName string   `protobuf:"bytes,2,opt,name=implementationName,proto3" json:"implementationName,omitempty"`
	Methods            []string `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *ServiceMetadata) Reset() {
	*x = ServiceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMetadata) ProtoMessage() {}

func (x *ServiceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMetadata.ProtoReflect.Descriptor instead.
func (*ServiceMetadata) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceMetadata) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *ServiceMetadata) GetImplementationName() string {
	if x != nil {
		return x.ImplementationName
	}
	return ""
}

func (x *ServiceMetadata) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type WatchEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName      string         `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ImplementationName string         `protobuf:"bytes,2,opt,name=implementationName,proto3" json:"implementationName,omitempty"`
	Method             string         `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	IsParam            bool           `protobuf:"varint,4,opt,name=isParam,proto3" json:"isParam,omitempty"`
	IsEdit             bool           `protobuf:"varint,5,opt,name=isEdit,proto3" json:"isEdit,omitempty"`
	Matchers           []*Matcher     `protobuf:"bytes,7,rep,name=matchers,proto3" json:"matchers,omitempty"`
	EditRequests       []*EditRequest `protobuf:"bytes,8,rep,name=editRequests,proto3" json:"editRequests,omitempty"`
}

func (x *WatchEditRequest) Reset() {
	*x = WatchEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchEditRequest) ProtoMessage() {}

func (x *WatchEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchEditRequest.ProtoReflect.Descriptor instead.
func (*WatchEditRequest) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{2}
}

func (x *WatchEditRequest) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *WatchEditRequest) GetImplementationName() string {
	if x != nil {
		return x.ImplementationName
	}
	return ""
}

func (x *WatchEditRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *WatchEditRequest) GetIsParam() bool {
	if x != nil {
		return x.IsParam
	}
	return false
}

func (x *WatchEditRequest) GetIsEdit() bool {
	if x != nil {
		return x.IsEdit
	}
	return false
}

func (x *WatchEditRequest) GetMatchers() []*Matcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

func (x *WatchEditRequest) GetEditRequests() []*EditRequest {
	if x != nil {
		return x.EditRequests
	}
	return nil
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName      string     `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ImplementationName string     `protobuf:"bytes,2,opt,name=implementationName,proto3" json:"implementationName,omitempty"`
	Method             string     `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Input              bool       `protobuf:"varint,4,opt,name=input,proto3" json:"input,omitempty"`
	Output             bool       `protobuf:"varint,5,opt,name=output,proto3" json:"output,omitempty"`
	Matchers           []*Matcher `protobuf:"bytes,6,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{3}
}

func (x *WatchRequest) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *WatchRequest) GetImplementationName() string {
	if x != nil {
		return x.ImplementationName
	}
	return ""
}

func (x *WatchRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *WatchRequest) GetInput() bool {
	if x != nil {
		return x.Input
	}
	return false
}

func (x *WatchRequest) GetOutput() bool {
	if x != nil {
		return x.Output
	}
	return false
}

func (x *WatchRequest) GetMatchers() []*Matcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type EditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EditRequest) Reset() {
	*x = EditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRequest) ProtoMessage() {}

func (x *EditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRequest.ProtoReflect.Descriptor instead.
func (*EditRequest) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{4}
}

func (x *EditRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *EditRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *EditRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	MatchPath  string `protobuf:"bytes,2,opt,name=matchPath,proto3" json:"matchPath,omitempty"`
	MatchValue string `protobuf:"bytes,3,opt,name=matchValue,proto3" json:"matchValue,omitempty"`
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{5}
}

func (x *Matcher) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Matcher) GetMatchPath() string {
	if x != nil {
		return x.MatchPath
	}
	return ""
}

func (x *Matcher) GetMatchValue() string {
	if x != nil {
		return x.MatchValue
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName      string   `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ImplementationName string   `protobuf:"bytes,2,opt,name=implementationName,proto3" json:"implementationName,omitempty"`
	MethodName         string   `protobuf:"bytes,3,opt,name=methodName,proto3" json:"methodName,omitempty"`
	IsParam            bool     `protobuf:"varint,4,opt,name=isParam,proto3" json:"isParam,omitempty"`
	Params             []string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_boot_debug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP(), []int{6}
}

func (x *WatchResponse) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *WatchResponse) GetImplementationName() string {
	if x != nil {
		return x.ImplementationName
	}
	return ""
}

func (x *WatchResponse) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *WatchResponse) GetIsParam() bool {
	if x != nil {
		return x.IsParam
	}
	return false
}

func (x *WatchResponse) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_debug_api_ioc_golang_boot_debug_proto protoreflect.FileDescriptor

var file_debug_api_ioc_golang_boot_debug_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x2f, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0xaa, 0x02, 0x0a,
	0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x45,
	0x64, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x45, 0x64, 0x69,
	0x74, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x65, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x65, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x0b,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5d, 0x0a, 0x07, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x32, 0x80, 0x02, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d,
	0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x4e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x69, 0x6f, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x54, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x45, 0x64, 0x69, 0x74, 0x12, 0x21,
	0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x62,
	0x6f, 0x6f, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x6f, 0x63, 0x5f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_debug_api_ioc_golang_boot_debug_proto_rawDescOnce sync.Once
	file_debug_api_ioc_golang_boot_debug_proto_rawDescData = file_debug_api_ioc_golang_boot_debug_proto_rawDesc
)

func file_debug_api_ioc_golang_boot_debug_proto_rawDescGZIP() []byte {
	file_debug_api_ioc_golang_boot_debug_proto_rawDescOnce.Do(func() {
		file_debug_api_ioc_golang_boot_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_debug_api_ioc_golang_boot_debug_proto_rawDescData)
	})
	return file_debug_api_ioc_golang_boot_debug_proto_rawDescData
}

var file_debug_api_ioc_golang_boot_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_debug_api_ioc_golang_boot_debug_proto_goTypes = []interface{}{
	(*ListServiceResponse)(nil), // 0: ioc_golang.boot.ListServiceResponse
	(*ServiceMetadata)(nil),     // 1: ioc_golang.boot.ServiceMetadata
	(*WatchEditRequest)(nil),    // 2: ioc_golang.boot.WatchEditRequest
	(*WatchRequest)(nil),        // 3: ioc_golang.boot.WatchRequest
	(*EditRequest)(nil),         // 4: ioc_golang.boot.EditRequest
	(*Matcher)(nil),             // 5: ioc_golang.boot.Matcher
	(*WatchResponse)(nil),       // 6: ioc_golang.boot.WatchResponse
	(*emptypb.Empty)(nil),       // 7: google.protobuf.Empty
}
var file_debug_api_ioc_golang_boot_debug_proto_depIdxs = []int32{
	1, // 0: ioc_golang.boot.ListServiceResponse.serviceMetadata:type_name -> ioc_golang.boot.ServiceMetadata
	5, // 1: ioc_golang.boot.WatchEditRequest.matchers:type_name -> ioc_golang.boot.Matcher
	4, // 2: ioc_golang.boot.WatchEditRequest.editRequests:type_name -> ioc_golang.boot.EditRequest
	5, // 3: ioc_golang.boot.WatchRequest.matchers:type_name -> ioc_golang.boot.Matcher
	3, // 4: ioc_golang.boot.DebugService.Watch:input_type -> ioc_golang.boot.WatchRequest
	7, // 5: ioc_golang.boot.DebugService.ListServices:input_type -> google.protobuf.Empty
	2, // 6: ioc_golang.boot.DebugService.WatchEdit:input_type -> ioc_golang.boot.WatchEditRequest
	6, // 7: ioc_golang.boot.DebugService.Watch:output_type -> ioc_golang.boot.WatchResponse
	0, // 8: ioc_golang.boot.DebugService.ListServices:output_type -> ioc_golang.boot.ListServiceResponse
	6, // 9: ioc_golang.boot.DebugService.WatchEdit:output_type -> ioc_golang.boot.WatchResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_debug_api_ioc_golang_boot_debug_proto_init() }
func file_debug_api_ioc_golang_boot_debug_proto_init() {
	if File_debug_api_ioc_golang_boot_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchEditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_debug_api_ioc_golang_boot_debug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_debug_api_ioc_golang_boot_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_debug_api_ioc_golang_boot_debug_proto_goTypes,
		DependencyIndexes: file_debug_api_ioc_golang_boot_debug_proto_depIdxs,
		MessageInfos:      file_debug_api_ioc_golang_boot_debug_proto_msgTypes,
	}.Build()
	File_debug_api_ioc_golang_boot_debug_proto = out.File
	file_debug_api_ioc_golang_boot_debug_proto_rawDesc = nil
	file_debug_api_ioc_golang_boot_debug_proto_goTypes = nil
	file_debug_api_ioc_golang_boot_debug_proto_depIdxs = nil
}

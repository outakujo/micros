// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: v1/goods.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Count int64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreateGoodsRequest) Reset() {
	*x = CreateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsRequest) ProtoMessage() {}

func (x *CreateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGoodsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoodsRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateGoodsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Count int64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UpdateGoodsRequest) Reset() {
	*x = UpdateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsRequest) ProtoMessage() {}

func (x *UpdateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGoodsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGoodsRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateGoodsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGoodsRequest) Reset() {
	*x = DeleteGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsRequest) ProtoMessage() {}

func (x *DeleteGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoodsReply) Reset() {
	*x = DeleteGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsReply) ProtoMessage() {}

func (x *DeleteGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsReply.ProtoReflect.Descriptor instead.
func (*DeleteGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{3}
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GetGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoodsModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Count int64   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GoodsModel) Reset() {
	*x = GoodsModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsModel) ProtoMessage() {}

func (x *GoodsModel) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsModel.ProtoReflect.Descriptor instead.
func (*GoodsModel) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{5}
}

func (x *GoodsModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsModel) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodsModel) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListGoodsRequest) Reset() {
	*x = ListGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsRequest) ProtoMessage() {}

func (x *ListGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsRequest.ProtoReflect.Descriptor instead.
func (*ListGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{6}
}

func (x *ListGoodsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListGoodsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List    []*GoodsModel `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count   int64         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	MaxPage int64         `protobuf:"varint,3,opt,name=maxPage,proto3" json:"maxPage,omitempty"`
}

func (x *ListGoodsReply) Reset() {
	*x = ListGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsReply) ProtoMessage() {}

func (x *ListGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsReply.ProtoReflect.Descriptor instead.
func (*ListGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{7}
}

func (x *ListGoodsReply) GetList() []*GoodsModel {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListGoodsReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListGoodsReply) GetMaxPage() int64 {
	if x != nil {
		return x.MaxPage
	}
	return 0
}

var File_v1_goods_proto protoreflect.FileDescriptor

var file_v1_goods_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x54, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x54,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5c, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x3a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x68, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61,
	0x78, 0x50, 0x61, 0x67, 0x65, 0x32, 0xcc, 0x02, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12,
	0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_goods_proto_rawDescOnce sync.Once
	file_v1_goods_proto_rawDescData = file_v1_goods_proto_rawDesc
)

func file_v1_goods_proto_rawDescGZIP() []byte {
	file_v1_goods_proto_rawDescOnce.Do(func() {
		file_v1_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_goods_proto_rawDescData)
	})
	return file_v1_goods_proto_rawDescData
}

var file_v1_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_goods_proto_goTypes = []interface{}{
	(*CreateGoodsRequest)(nil), // 0: api.v1.CreateGoodsRequest
	(*UpdateGoodsRequest)(nil), // 1: api.v1.UpdateGoodsRequest
	(*DeleteGoodsRequest)(nil), // 2: api.v1.DeleteGoodsRequest
	(*DeleteGoodsReply)(nil),   // 3: api.v1.DeleteGoodsReply
	(*GetGoodsRequest)(nil),    // 4: api.v1.GetGoodsRequest
	(*GoodsModel)(nil),         // 5: api.v1.GoodsModel
	(*ListGoodsRequest)(nil),   // 6: api.v1.ListGoodsRequest
	(*ListGoodsReply)(nil),     // 7: api.v1.ListGoodsReply
}
var file_v1_goods_proto_depIdxs = []int32{
	5, // 0: api.v1.ListGoodsReply.list:type_name -> api.v1.GoodsModel
	0, // 1: api.v1.Goods.CreateGoods:input_type -> api.v1.CreateGoodsRequest
	1, // 2: api.v1.Goods.UpdateGoods:input_type -> api.v1.UpdateGoodsRequest
	2, // 3: api.v1.Goods.DeleteGoods:input_type -> api.v1.DeleteGoodsRequest
	4, // 4: api.v1.Goods.GetGoods:input_type -> api.v1.GetGoodsRequest
	6, // 5: api.v1.Goods.ListGoods:input_type -> api.v1.ListGoodsRequest
	5, // 6: api.v1.Goods.CreateGoods:output_type -> api.v1.GoodsModel
	5, // 7: api.v1.Goods.UpdateGoods:output_type -> api.v1.GoodsModel
	3, // 8: api.v1.Goods.DeleteGoods:output_type -> api.v1.DeleteGoodsReply
	5, // 9: api.v1.Goods.GetGoods:output_type -> api.v1.GoodsModel
	7, // 10: api.v1.Goods.ListGoods:output_type -> api.v1.ListGoodsReply
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_goods_proto_init() }
func file_v1_goods_proto_init() {
	if File_v1_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsReply); i {
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
		file_v1_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsModel); i {
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
		file_v1_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsReply); i {
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
			RawDescriptor: file_v1_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_goods_proto_goTypes,
		DependencyIndexes: file_v1_goods_proto_depIdxs,
		MessageInfos:      file_v1_goods_proto_msgTypes,
	}.Build()
	File_v1_goods_proto = out.File
	file_v1_goods_proto_rawDesc = nil
	file_v1_goods_proto_goTypes = nil
	file_v1_goods_proto_depIdxs = nil
}

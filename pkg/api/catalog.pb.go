// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: catalog.proto

package tvtime

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CatalogType int32

const (
	CatalogType_MOVIE   CatalogType = 0
	CatalogType_TV_SHOW CatalogType = 1
)

// Enum value maps for CatalogType.
var (
	CatalogType_name = map[int32]string{
		0: "MOVIE",
		1: "TV_SHOW",
	}
	CatalogType_value = map[string]int32{
		"MOVIE":   0,
		"TV_SHOW": 1,
	}
)

func (x CatalogType) Enum() *CatalogType {
	p := new(CatalogType)
	*p = x
	return p
}

func (x CatalogType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CatalogType) Descriptor() protoreflect.EnumDescriptor {
	return file_catalog_proto_enumTypes[0].Descriptor()
}

func (CatalogType) Type() protoreflect.EnumType {
	return &file_catalog_proto_enumTypes[0]
}

func (x CatalogType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CatalogType.Descriptor instead.
func (CatalogType) EnumDescriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{0}
}

type CatalogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug     string      `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name     string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string      `protobuf:"bytes,4,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Type     CatalogType `protobuf:"varint,5,opt,name=type,proto3,enum=com.aviebrantz.tvtime.CatalogType" json:"type,omitempty"`
}

func (x *CatalogEntry) Reset() {
	*x = CatalogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogEntry) ProtoMessage() {}

func (x *CatalogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogEntry.ProtoReflect.Descriptor instead.
func (*CatalogEntry) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *CatalogEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CatalogEntry) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CatalogEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CatalogEntry) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CatalogEntry) GetType() CatalogType {
	if x != nil {
		return x.Type
	}
	return CatalogType_MOVIE
}

type CreateEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string      `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Type     CatalogType `protobuf:"varint,3,opt,name=type,proto3,enum=com.aviebrantz.tvtime.CatalogType" json:"type,omitempty"`
}

func (x *CreateEntryRequest) Reset() {
	*x = CreateEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntryRequest) ProtoMessage() {}

func (x *CreateEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntryRequest.ProtoReflect.Descriptor instead.
func (*CreateEntryRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEntryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEntryRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateEntryRequest) GetType() CatalogType {
	if x != nil {
		return x.Type
	}
	return CatalogType_MOVIE
}

type CreateEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *CatalogEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *CreateEntryResponse) Reset() {
	*x = CreateEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntryResponse) ProtoMessage() {}

func (x *CreateEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntryResponse.ProtoReflect.Descriptor instead.
func (*CreateEntryResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEntryResponse) GetEntry() *CatalogEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type SearchCatalogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    *wrappers.StringValue `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	Type    CatalogType           `protobuf:"varint,2,opt,name=type,proto3,enum=com.aviebrantz.tvtime.CatalogType" json:"type,omitempty"`
	Page    int64                 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64                 `protobuf:"varint,4,opt,name=perPage,proto3" json:"perPage,omitempty"`
}

func (x *SearchCatalogFilter) Reset() {
	*x = SearchCatalogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCatalogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCatalogFilter) ProtoMessage() {}

func (x *SearchCatalogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCatalogFilter.ProtoReflect.Descriptor instead.
func (*SearchCatalogFilter) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *SearchCatalogFilter) GetTerm() *wrappers.StringValue {
	if x != nil {
		return x.Term
	}
	return nil
}

func (x *SearchCatalogFilter) GetType() CatalogType {
	if x != nil {
		return x.Type
	}
	return CatalogType_MOVIE
}

func (x *SearchCatalogFilter) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchCatalogFilter) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*CatalogEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	Total   int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page    int32           `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32           `protobuf:"varint,4,opt,name=perPage,proto3" json:"perPage,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetEntries() []*CatalogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListResponse) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type GetCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *GetCatalogRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry     *CatalogEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Favorited int64         `protobuf:"varint,2,opt,name=favorited,proto3" json:"favorited,omitempty"`
}

func (x *GetCatalogResponse) Reset() {
	*x = GetCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogResponse) ProtoMessage() {}

func (x *GetCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogResponse.ProtoReflect.Descriptor instead.
func (*GetCatalogResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *GetCatalogResponse) GetEntry() *CatalogEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *GetCatalogResponse) GetFavorited() int64 {
	if x != nil {
		return x.Favorited
	}
	return 0
}

var File_catalog_proto protoreflect.FileDescriptor

var file_catalog_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e,
	0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76,
	0x69, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x7c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65,
	0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x50, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65,
	0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x36, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x76, 0x69, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72,
	0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x6d,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72,
	0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x2a, 0x25, 0x0a,
	0x0b, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x4d, 0x4f, 0x56, 0x49, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x56, 0x5f, 0x53, 0x48,
	0x4f, 0x57, 0x10, 0x01, 0x32, 0x81, 0x03, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72, 0x61,
	0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74,
	0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x92, 0x41, 0x0e, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a,
	0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69,
	0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x12, 0x76, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76,
	0x69, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x65, 0x62, 0x72, 0x61, 0x6e,
	0x74, 0x7a, 0x2e, 0x74, 0x76, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2f, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x74, 0x76,
	0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_proto_rawDescOnce sync.Once
	file_catalog_proto_rawDescData = file_catalog_proto_rawDesc
)

func file_catalog_proto_rawDescGZIP() []byte {
	file_catalog_proto_rawDescOnce.Do(func() {
		file_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_proto_rawDescData)
	})
	return file_catalog_proto_rawDescData
}

var file_catalog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_catalog_proto_goTypes = []interface{}{
	(CatalogType)(0),             // 0: com.aviebrantz.tvtime.CatalogType
	(*CatalogEntry)(nil),         // 1: com.aviebrantz.tvtime.CatalogEntry
	(*CreateEntryRequest)(nil),   // 2: com.aviebrantz.tvtime.CreateEntryRequest
	(*CreateEntryResponse)(nil),  // 3: com.aviebrantz.tvtime.CreateEntryResponse
	(*SearchCatalogFilter)(nil),  // 4: com.aviebrantz.tvtime.SearchCatalogFilter
	(*ListResponse)(nil),         // 5: com.aviebrantz.tvtime.ListResponse
	(*GetCatalogRequest)(nil),    // 6: com.aviebrantz.tvtime.GetCatalogRequest
	(*GetCatalogResponse)(nil),   // 7: com.aviebrantz.tvtime.GetCatalogResponse
	(*wrappers.StringValue)(nil), // 8: google.protobuf.StringValue
}
var file_catalog_proto_depIdxs = []int32{
	0,  // 0: com.aviebrantz.tvtime.CatalogEntry.type:type_name -> com.aviebrantz.tvtime.CatalogType
	0,  // 1: com.aviebrantz.tvtime.CreateEntryRequest.type:type_name -> com.aviebrantz.tvtime.CatalogType
	1,  // 2: com.aviebrantz.tvtime.CreateEntryResponse.entry:type_name -> com.aviebrantz.tvtime.CatalogEntry
	8,  // 3: com.aviebrantz.tvtime.SearchCatalogFilter.term:type_name -> google.protobuf.StringValue
	0,  // 4: com.aviebrantz.tvtime.SearchCatalogFilter.type:type_name -> com.aviebrantz.tvtime.CatalogType
	1,  // 5: com.aviebrantz.tvtime.ListResponse.entries:type_name -> com.aviebrantz.tvtime.CatalogEntry
	1,  // 6: com.aviebrantz.tvtime.GetCatalogResponse.entry:type_name -> com.aviebrantz.tvtime.CatalogEntry
	2,  // 7: com.aviebrantz.tvtime.CatalogService.Create:input_type -> com.aviebrantz.tvtime.CreateEntryRequest
	4,  // 8: com.aviebrantz.tvtime.CatalogService.List:input_type -> com.aviebrantz.tvtime.SearchCatalogFilter
	6,  // 9: com.aviebrantz.tvtime.CatalogService.Get:input_type -> com.aviebrantz.tvtime.GetCatalogRequest
	3,  // 10: com.aviebrantz.tvtime.CatalogService.Create:output_type -> com.aviebrantz.tvtime.CreateEntryResponse
	5,  // 11: com.aviebrantz.tvtime.CatalogService.List:output_type -> com.aviebrantz.tvtime.ListResponse
	7,  // 12: com.aviebrantz.tvtime.CatalogService.Get:output_type -> com.aviebrantz.tvtime.GetCatalogResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_catalog_proto_init() }
func file_catalog_proto_init() {
	if File_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatalogEntry); i {
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
		file_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntryRequest); i {
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
		file_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntryResponse); i {
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
		file_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCatalogFilter); i {
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
		file_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogRequest); i {
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
		file_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogResponse); i {
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
			RawDescriptor: file_catalog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_proto_goTypes,
		DependencyIndexes: file_catalog_proto_depIdxs,
		EnumInfos:         file_catalog_proto_enumTypes,
		MessageInfos:      file_catalog_proto_msgTypes,
	}.Build()
	File_catalog_proto = out.File
	file_catalog_proto_rawDesc = nil
	file_catalog_proto_goTypes = nil
	file_catalog_proto_depIdxs = nil
}

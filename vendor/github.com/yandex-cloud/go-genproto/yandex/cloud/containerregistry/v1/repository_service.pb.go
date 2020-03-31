// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/repository_service.proto

package containerregistry

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetRepositoryRequest struct {
	// ID of the Repository resource to return.
	//
	// To get the repository ID use a [RepositoryService.List] request.
	RepositoryId         string   `protobuf:"bytes,1,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRepositoryRequest) Reset()         { *m = GetRepositoryRequest{} }
func (m *GetRepositoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetRepositoryRequest) ProtoMessage()    {}
func (*GetRepositoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ad489b83930188, []int{0}
}

func (m *GetRepositoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRepositoryRequest.Unmarshal(m, b)
}
func (m *GetRepositoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRepositoryRequest.Marshal(b, m, deterministic)
}
func (m *GetRepositoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRepositoryRequest.Merge(m, src)
}
func (m *GetRepositoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetRepositoryRequest.Size(m)
}
func (m *GetRepositoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRepositoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRepositoryRequest proto.InternalMessageInfo

func (m *GetRepositoryRequest) GetRepositoryId() string {
	if m != nil {
		return m.RepositoryId
	}
	return ""
}

type GetRepositoryByNameRequest struct {
	// Name of the Repository resource to return.
	//
	// To get the repository name use a [RepositoryService.List] request.
	RepositoryName       string   `protobuf:"bytes,1,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRepositoryByNameRequest) Reset()         { *m = GetRepositoryByNameRequest{} }
func (m *GetRepositoryByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetRepositoryByNameRequest) ProtoMessage()    {}
func (*GetRepositoryByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ad489b83930188, []int{1}
}

func (m *GetRepositoryByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRepositoryByNameRequest.Unmarshal(m, b)
}
func (m *GetRepositoryByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRepositoryByNameRequest.Marshal(b, m, deterministic)
}
func (m *GetRepositoryByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRepositoryByNameRequest.Merge(m, src)
}
func (m *GetRepositoryByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetRepositoryByNameRequest.Size(m)
}
func (m *GetRepositoryByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRepositoryByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRepositoryByNameRequest proto.InternalMessageInfo

func (m *GetRepositoryByNameRequest) GetRepositoryName() string {
	if m != nil {
		return m.RepositoryName
	}
	return ""
}

type ListRepositoriesRequest struct {
	// ID of the registry to list repositories in.
	//
	// To get the registry ID use a [RegistryService.List] request.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// ID of the folder to list registries in.
	//
	// [folder_id] is ignored if a [ListImagesRequest.registry_id] is specified in the request.
	//
	// To get the folder ID use a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,6,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListRepositoriesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListRepositoriesResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Repository.name] field.
	// 2. An operator. Can be either `=` or `!=` for single values, `IN` or `NOT IN` for lists of values.
	// 3. Value or a list of values to compare against the values of the field.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy              string   `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesRequest) Reset()         { *m = ListRepositoriesRequest{} }
func (m *ListRepositoriesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesRequest) ProtoMessage()    {}
func (*ListRepositoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ad489b83930188, []int{2}
}

func (m *ListRepositoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesRequest.Unmarshal(m, b)
}
func (m *ListRepositoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesRequest.Marshal(b, m, deterministic)
}
func (m *ListRepositoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesRequest.Merge(m, src)
}
func (m *ListRepositoriesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesRequest.Size(m)
}
func (m *ListRepositoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesRequest proto.InternalMessageInfo

func (m *ListRepositoriesRequest) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *ListRepositoriesRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *ListRepositoriesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRepositoriesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListRepositoriesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListRepositoriesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type ListRepositoriesResponse struct {
	// List of Repository resources.
	Repositories []*Repository `protobuf:"bytes,1,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListRepositoriesRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListRepositoriesRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesResponse) Reset()         { *m = ListRepositoriesResponse{} }
func (m *ListRepositoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesResponse) ProtoMessage()    {}
func (*ListRepositoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ad489b83930188, []int{3}
}

func (m *ListRepositoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesResponse.Unmarshal(m, b)
}
func (m *ListRepositoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesResponse.Marshal(b, m, deterministic)
}
func (m *ListRepositoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesResponse.Merge(m, src)
}
func (m *ListRepositoriesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesResponse.Size(m)
}
func (m *ListRepositoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesResponse proto.InternalMessageInfo

func (m *ListRepositoriesResponse) GetRepositories() []*Repository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func (m *ListRepositoriesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRepositoryRequest)(nil), "yandex.cloud.containerregistry.v1.GetRepositoryRequest")
	proto.RegisterType((*GetRepositoryByNameRequest)(nil), "yandex.cloud.containerregistry.v1.GetRepositoryByNameRequest")
	proto.RegisterType((*ListRepositoriesRequest)(nil), "yandex.cloud.containerregistry.v1.ListRepositoriesRequest")
	proto.RegisterType((*ListRepositoriesResponse)(nil), "yandex.cloud.containerregistry.v1.ListRepositoriesResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/repository_service.proto", fileDescriptor_b3ad489b83930188)
}

var fileDescriptor_b3ad489b83930188 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x4f, 0x3b, 0x55,
	0x14, 0xcd, 0xa3, 0x50, 0xdb, 0x07, 0x48, 0x78, 0xc1, 0xd8, 0x34, 0x9a, 0x94, 0x41, 0xb0, 0x2d,
	0x99, 0x99, 0x0e, 0x84, 0x18, 0x0a, 0x68, 0xac, 0x0a, 0x36, 0xe2, 0xbf, 0xa2, 0x31, 0x91, 0x90,
	0x66, 0xda, 0xb9, 0x8c, 0x2f, 0xb6, 0xf3, 0xc6, 0x79, 0xaf, 0x0d, 0x83, 0x31, 0x31, 0xae, 0x0c,
	0xc6, 0x85, 0x31, 0x7e, 0x07, 0x63, 0x5c, 0xf9, 0x15, 0x5c, 0xc0, 0x5a, 0xbf, 0x82, 0x0b, 0xd7,
	0xba, 0x73, 0xe3, 0x2f, 0xf3, 0xaf, 0xd3, 0x69, 0x0b, 0xb4, 0xbf, 0x55, 0x9b, 0x77, 0xcf, 0x3d,
	0xf7, 0x9c, 0xfb, 0xde, 0xbd, 0x2d, 0xae, 0xba, 0xba, 0x65, 0xc0, 0x95, 0xda, 0xee, 0xb0, 0x9e,
	0xa1, 0xb6, 0x99, 0x25, 0x74, 0x6a, 0x81, 0xe3, 0x80, 0x49, 0xb9, 0x70, 0x5c, 0xb5, 0xaf, 0xa9,
	0x0e, 0xd8, 0x8c, 0x53, 0xc1, 0x1c, 0xb7, 0xc9, 0xc1, 0xe9, 0xd3, 0x36, 0x28, 0xb6, 0xc3, 0x04,
	0x23, 0xeb, 0x41, 0xae, 0xe2, 0xe7, 0x2a, 0x63, 0xb9, 0x4a, 0x5f, 0xcb, 0x17, 0x12, 0xf4, 0xba,
	0x4d, 0x55, 0x66, 0x83, 0xa3, 0x0b, 0xca, 0xac, 0x80, 0x64, 0x14, 0xd1, 0x6e, 0x03, 0xe7, 0xe1,
	0x47, 0x88, 0xd8, 0x99, 0x45, 0x62, 0x98, 0xb3, 0x95, 0xc8, 0x19, 0xd4, 0x1c, 0xab, 0xfe, 0x62,
	0x02, 0xd7, 0xd7, 0x3b, 0xd4, 0x18, 0x0e, 0xbf, 0x60, 0x32, 0x66, 0x76, 0xc0, 0x17, 0xae, 0x5b,
	0x16, 0x13, 0x7e, 0x30, 0x14, 0x26, 0xd5, 0xf1, 0xda, 0x09, 0x88, 0xc6, 0xa0, 0x76, 0x03, 0xbe,
	0xe8, 0x01, 0x17, 0x44, 0xc3, 0xcb, 0x43, 0x3d, 0xa3, 0x46, 0x0e, 0x15, 0x50, 0x31, 0x5b, 0x5b,
	0xfa, 0xfb, 0x56, 0x43, 0x37, 0x77, 0xda, 0xfc, 0xe1, 0xd1, 0x5e, 0xa5, 0xb1, 0x14, 0x43, 0xea,
	0x86, 0xf4, 0x3d, 0xc2, 0xf9, 0x04, 0x57, 0xcd, 0x7d, 0x4f, 0xef, 0x42, 0xc4, 0xc8, 0xf0, 0xca,
	0x10, 0xa3, 0xa5, 0x77, 0x21, 0xe4, 0x3c, 0xf6, 0x38, 0xff, 0xb9, 0xd5, 0x5e, 0x3d, 0xd7, 0xe5,
	0xeb, 0x8a, 0xbc, 0x7f, 0xb1, 0x5d, 0x7c, 0xad, 0x7a, 0xae, 0x34, 0xe5, 0x8b, 0xc1, 0x41, 0xa9,
	0x5c, 0x54, 0x8b, 0x0f, 0x85, 0x4b, 0xa5, 0x72, 0xe3, 0xd9, 0x98, 0xde, 0xab, 0x2b, 0x7d, 0x3b,
	0x87, 0x9f, 0x3f, 0xa5, 0x3c, 0x16, 0x44, 0x81, 0x47, 0x62, 0x4a, 0x78, 0x31, 0xea, 0x7d, 0x6c,
	0x2e, 0x33, 0x30, 0x86, 0xa3, 0x60, 0xdd, 0x20, 0x9b, 0x38, 0x7b, 0xc9, 0x3a, 0x06, 0x38, 0x1e,
	0x30, 0x3d, 0x02, 0xcc, 0x04, 0xa1, 0xba, 0x41, 0x5e, 0xc6, 0x59, 0x5b, 0x37, 0xa1, 0xc9, 0xe9,
	0x35, 0xe4, 0xe6, 0x0a, 0xa8, 0x98, 0xaa, 0xe1, 0xff, 0x6e, 0xb5, 0xf4, 0xe1, 0x91, 0x56, 0xa9,
	0x54, 0x1a, 0x19, 0x2f, 0x78, 0x46, 0xaf, 0x81, 0x14, 0x31, 0xf6, 0x81, 0x82, 0x7d, 0x0e, 0x56,
	0x2e, 0xe5, 0x13, 0x66, 0x6f, 0xee, 0xb4, 0x05, 0x1f, 0xd9, 0xf0, 0x59, 0x3e, 0xf2, 0x62, 0x44,
	0xc2, 0xe9, 0x4b, 0xda, 0x11, 0xe0, 0xe4, 0xe6, 0x7d, 0x14, 0xbe, 0xb9, 0x1b, 0xf0, 0x85, 0x11,
	0xf2, 0x12, 0xce, 0x30, 0xc7, 0x13, 0xd7, 0x72, 0x73, 0x0b, 0xa3, 0x5c, 0xcf, 0xf8, 0xa1, 0x9a,
	0x2b, 0xfd, 0x84, 0x70, 0x6e, 0xbc, 0x15, 0xdc, 0x66, 0x16, 0x07, 0xf2, 0x21, 0x8e, 0xef, 0x91,
	0x02, 0xcf, 0xa1, 0x42, 0xaa, 0xb8, 0xb8, 0x23, 0x2b, 0x8f, 0x4e, 0x86, 0x32, 0xf4, 0x6c, 0x12,
	0x14, 0x64, 0x0b, 0xaf, 0x58, 0x70, 0x25, 0x9a, 0x43, 0x46, 0xbd, 0x96, 0x64, 0x1b, 0xcb, 0xde,
	0xf1, 0x07, 0x91, 0xc3, 0x9d, 0x9f, 0xb3, 0x78, 0x35, 0x26, 0x39, 0x0b, 0x26, 0x93, 0xfc, 0x82,
	0x70, 0xea, 0x04, 0x04, 0x79, 0x65, 0x0a, 0x09, 0x93, 0x1e, 0x6f, 0x7e, 0x36, 0xed, 0xd2, 0xc1,
	0x37, 0x7f, 0xfe, 0xf5, 0xe3, 0xdc, 0x1e, 0xd9, 0x8d, 0x07, 0x53, 0x9e, 0x38, 0x99, 0x14, 0xb8,
	0xfa, 0x65, 0x62, 0x2c, 0xbe, 0x22, 0x2e, 0xce, 0x9e, 0x80, 0x08, 0x9e, 0x3a, 0x39, 0x9a, 0x55,
	0x71, 0x62, 0x44, 0x66, 0xd4, 0x4d, 0x7e, 0x45, 0x78, 0xde, 0xbb, 0x55, 0x52, 0x9d, 0x22, 0xef,
	0x9e, 0x49, 0xc8, 0x1f, 0x3c, 0x55, 0x6e, 0xf0, 0x74, 0xa4, 0x6d, 0xbf, 0x73, 0x9b, 0x64, 0x63,
	0x8a, 0xce, 0x91, 0xdf, 0x11, 0x26, 0x1e, 0xd3, 0xeb, 0xfe, 0x62, 0xac, 0x51, 0xcb, 0xa0, 0x96,
	0xc9, 0x89, 0x92, 0x14, 0x10, 0xae, 0xcd, 0x71, 0x60, 0x24, 0x58, 0x9d, 0x1a, 0x1f, 0x8a, 0x3c,
	0xf5, 0x45, 0x1e, 0x93, 0x37, 0xa7, 0xbc, 0x5e, 0xce, 0x7a, 0x4e, 0x1b, 0xbc, 0xcb, 0xad, 0x76,
	0xc6, 0xe5, 0xfe, 0x8b, 0xf0, 0xea, 0x19, 0x8c, 0x9e, 0xca, 0x13, 0x45, 0x8d, 0xe1, 0x22, 0x0f,
	0xeb, 0x49, 0x78, 0xbc, 0xd1, 0xdf, 0x8f, 0xbe, 0x49, 0xdf, 0xa1, 0xdf, 0xfe, 0x28, 0xef, 0xe3,
	0xc2, 0x7d, 0x54, 0xef, 0x82, 0xd0, 0x0d, 0x5d, 0xe8, 0xe4, 0xb9, 0x60, 0xbf, 0x07, 0xfb, 0xbc,
	0xd5, 0xbb, 0x54, 0xde, 0xea, 0xda, 0xc2, 0xf5, 0x3d, 0xbf, 0x2d, 0xbd, 0x31, 0xbb, 0x67, 0x3e,
	0x5a, 0xa9, 0x8a, 0xca, 0xe4, 0x7f, 0x84, 0xd7, 0x3e, 0xb6, 0x0d, 0x5d, 0xc0, 0x88, 0xf1, 0xca,
	0x44, 0xe3, 0x93, 0xa0, 0x33, 0x78, 0xff, 0xc1, 0xf3, 0x7e, 0x88, 0x37, 0x1e, 0x60, 0x9b, 0xc6,
	0xfe, 0x3b, 0xd2, 0xf1, 0xec, 0xf6, 0x7b, 0x13, 0x8a, 0x55, 0x51, 0xb9, 0xf6, 0x35, 0xc2, 0x9b,
	0x49, 0xab, 0x36, 0x9d, 0x38, 0x2d, 0x9f, 0x7e, 0x62, 0x52, 0xf1, 0x59, 0xaf, 0xa5, 0xb4, 0x59,
	0x57, 0x0d, 0x32, 0xe4, 0xe0, 0xa7, 0xd9, 0x64, 0xb2, 0x09, 0x96, 0x2f, 0x54, 0x7d, 0xf4, 0xff,
	0xc0, 0xc1, 0xd8, 0x61, 0x2b, 0xed, 0xa7, 0xee, 0x3e, 0x09, 0x00, 0x00, 0xff, 0xff, 0xee, 0xd5,
	0x3a, 0x62, 0xf0, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	// Returns the specified Repository resource.
	//
	// To get the list of available Repository resources, make a [List] request.
	Get(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*Repository, error)
	// Returns the specified Repository resource.
	//
	// To get the list of available Repository resources, make a [List] request.
	GetByName(ctx context.Context, in *GetRepositoryByNameRequest, opts ...grpc.CallOption) (*Repository, error)
	// Retrieves the list of Repository resources in the specified registry.
	List(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error)
	// Lists access bindings for the specified repository.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified repository.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified repository.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type repositoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryServiceClient(cc grpc.ClientConnInterface) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) Get(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) GetByName(ctx context.Context, in *GetRepositoryByNameRequest, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) List(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error) {
	out := new(ListRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/ListAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/SetAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/UpdateAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
type RepositoryServiceServer interface {
	// Returns the specified Repository resource.
	//
	// To get the list of available Repository resources, make a [List] request.
	Get(context.Context, *GetRepositoryRequest) (*Repository, error)
	// Returns the specified Repository resource.
	//
	// To get the list of available Repository resources, make a [List] request.
	GetByName(context.Context, *GetRepositoryByNameRequest) (*Repository, error)
	// Retrieves the list of Repository resources in the specified registry.
	List(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error)
	// Lists access bindings for the specified repository.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified repository.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified repository.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedRepositoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepositoryServiceServer struct {
}

func (*UnimplementedRepositoryServiceServer) Get(ctx context.Context, req *GetRepositoryRequest) (*Repository, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRepositoryServiceServer) GetByName(ctx context.Context, req *GetRepositoryByNameRequest) (*Repository, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (*UnimplementedRepositoryServiceServer) List(ctx context.Context, req *ListRepositoriesRequest) (*ListRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedRepositoryServiceServer) ListAccessBindings(ctx context.Context, req *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (*UnimplementedRepositoryServiceServer) SetAccessBindings(ctx context.Context, req *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (*UnimplementedRepositoryServiceServer) UpdateAccessBindings(ctx context.Context, req *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Get(ctx, req.(*GetRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GetByName(ctx, req.(*GetRepositoryByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).List(ctx, req.(*ListRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/ListAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/SetAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/UpdateAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.containerregistry.v1.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RepositoryService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _RepositoryService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RepositoryService_List_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _RepositoryService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _RepositoryService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _RepositoryService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/containerregistry/v1/repository_service.proto",
}

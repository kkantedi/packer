// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/redis/v1/resource_preset_service.proto

package redis

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
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

type GetResourcePresetRequest struct {
	// ID of the resource preset to return.
	// To get the resource preset ID, use a [ResourcePresetService.List] request.
	ResourcePresetId     string   `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourcePresetRequest) Reset()         { *m = GetResourcePresetRequest{} }
func (m *GetResourcePresetRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourcePresetRequest) ProtoMessage()    {}
func (*GetResourcePresetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fffc675283e9ab1, []int{0}
}

func (m *GetResourcePresetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourcePresetRequest.Unmarshal(m, b)
}
func (m *GetResourcePresetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourcePresetRequest.Marshal(b, m, deterministic)
}
func (m *GetResourcePresetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourcePresetRequest.Merge(m, src)
}
func (m *GetResourcePresetRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourcePresetRequest.Size(m)
}
func (m *GetResourcePresetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourcePresetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourcePresetRequest proto.InternalMessageInfo

func (m *GetResourcePresetRequest) GetResourcePresetId() string {
	if m != nil {
		return m.ResourcePresetId
	}
	return ""
}

type ListResourcePresetsRequest struct {
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListResourcePresetsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the [ListResourcePresetsResponse.next_page_token]
	// returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResourcePresetsRequest) Reset()         { *m = ListResourcePresetsRequest{} }
func (m *ListResourcePresetsRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourcePresetsRequest) ProtoMessage()    {}
func (*ListResourcePresetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fffc675283e9ab1, []int{1}
}

func (m *ListResourcePresetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcePresetsRequest.Unmarshal(m, b)
}
func (m *ListResourcePresetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcePresetsRequest.Marshal(b, m, deterministic)
}
func (m *ListResourcePresetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcePresetsRequest.Merge(m, src)
}
func (m *ListResourcePresetsRequest) XXX_Size() int {
	return xxx_messageInfo_ListResourcePresetsRequest.Size(m)
}
func (m *ListResourcePresetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcePresetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcePresetsRequest proto.InternalMessageInfo

func (m *ListResourcePresetsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListResourcePresetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListResourcePresetsResponse struct {
	// List of resource presets.
	ResourcePresets []*ResourcePreset `protobuf:"bytes,1,rep,name=resource_presets,json=resourcePresets,proto3" json:"resource_presets,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListResourcePresetsRequest.page_size], use the [next_page_token] as the value
	// for the [ListResourcePresetsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResourcePresetsResponse) Reset()         { *m = ListResourcePresetsResponse{} }
func (m *ListResourcePresetsResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourcePresetsResponse) ProtoMessage()    {}
func (*ListResourcePresetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fffc675283e9ab1, []int{2}
}

func (m *ListResourcePresetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcePresetsResponse.Unmarshal(m, b)
}
func (m *ListResourcePresetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcePresetsResponse.Marshal(b, m, deterministic)
}
func (m *ListResourcePresetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcePresetsResponse.Merge(m, src)
}
func (m *ListResourcePresetsResponse) XXX_Size() int {
	return xxx_messageInfo_ListResourcePresetsResponse.Size(m)
}
func (m *ListResourcePresetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcePresetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcePresetsResponse proto.InternalMessageInfo

func (m *ListResourcePresetsResponse) GetResourcePresets() []*ResourcePreset {
	if m != nil {
		return m.ResourcePresets
	}
	return nil
}

func (m *ListResourcePresetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetResourcePresetRequest)(nil), "yandex.cloud.mdb.redis.v1.GetResourcePresetRequest")
	proto.RegisterType((*ListResourcePresetsRequest)(nil), "yandex.cloud.mdb.redis.v1.ListResourcePresetsRequest")
	proto.RegisterType((*ListResourcePresetsResponse)(nil), "yandex.cloud.mdb.redis.v1.ListResourcePresetsResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/redis/v1/resource_preset_service.proto", fileDescriptor_1fffc675283e9ab1)
}

var fileDescriptor_1fffc675283e9ab1 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6a, 0xd5, 0x40,
	0x14, 0x86, 0x99, 0x9b, 0x22, 0x76, 0x44, 0x5a, 0x06, 0x84, 0x98, 0x5a, 0xb8, 0x46, 0x90, 0xdb,
	0xc5, 0x9d, 0xe1, 0xde, 0x62, 0x15, 0x04, 0x17, 0x15, 0x29, 0x82, 0x48, 0x49, 0xbb, 0x10, 0x37,
	0x61, 0x92, 0x39, 0xc4, 0xc1, 0x9b, 0x99, 0x98, 0x99, 0x84, 0x5a, 0x71, 0xe3, 0x2b, 0xb8, 0x74,
	0xe3, 0xc6, 0x67, 0xf1, 0x01, 0x7c, 0x03, 0x71, 0xe1, 0x63, 0x48, 0x26, 0x29, 0x98, 0xd4, 0xd4,
	0x76, 0x17, 0x4e, 0xe6, 0xfb, 0xcf, 0x7f, 0xe6, 0x3f, 0x83, 0x1f, 0xbe, 0xe7, 0x4a, 0xc0, 0x09,
	0x4b, 0x57, 0xba, 0x12, 0x2c, 0x17, 0x09, 0x2b, 0x41, 0x48, 0xc3, 0xea, 0x05, 0x2b, 0xc1, 0xe8,
	0xaa, 0x4c, 0x21, 0x2e, 0x4a, 0x30, 0x60, 0x63, 0x03, 0x65, 0x2d, 0x53, 0xa0, 0x45, 0xa9, 0xad,
	0x26, 0xb7, 0x5b, 0x90, 0x3a, 0x90, 0xe6, 0x22, 0xa1, 0x0e, 0xa4, 0xf5, 0x22, 0xb8, 0x93, 0x69,
	0x9d, 0xad, 0x80, 0xf1, 0x42, 0x32, 0xae, 0x94, 0xb6, 0xdc, 0x4a, 0xad, 0x4c, 0x0b, 0x06, 0xec,
	0xd2, 0x1d, 0x3b, 0x60, 0xbb, 0x07, 0xd4, 0x7c, 0x25, 0x85, 0x13, 0x6c, 0x7f, 0x87, 0x2f, 0xb1,
	0x7f, 0x00, 0x36, 0xea, 0xd0, 0x43, 0x47, 0x46, 0xf0, 0xae, 0x02, 0x63, 0xc9, 0x12, 0x93, 0xe1,
	0x14, 0x52, 0xf8, 0x68, 0x8a, 0x66, 0xeb, 0xfb, 0x6b, 0xbf, 0xbf, 0x2f, 0x50, 0xb4, 0x59, 0xf6,
	0xc0, 0xe7, 0x22, 0x7c, 0x85, 0x83, 0x17, 0xd2, 0x0c, 0x04, 0xcd, 0x99, 0xe2, 0x16, 0x5e, 0x2f,
	0x78, 0x06, 0xb1, 0x91, 0xa7, 0xe0, 0x4f, 0xa6, 0x68, 0xe6, 0x45, 0xd7, 0x9b, 0xc2, 0x91, 0x3c,
	0x05, 0xb2, 0x8d, 0xb1, 0xfb, 0x69, 0xf5, 0x5b, 0x50, 0xbe, 0xd7, 0xb4, 0x89, 0xdc, 0xf1, 0xe3,
	0xa6, 0x10, 0x7e, 0x41, 0x78, 0xeb, 0x9f, 0xd2, 0xa6, 0xd0, 0xca, 0x00, 0x39, 0xc6, 0x9b, 0x03,
	0xb7, 0xc6, 0x47, 0x53, 0x6f, 0x76, 0x63, 0xb9, 0x43, 0x47, 0x6f, 0x9b, 0x0e, 0x26, 0xdf, 0xe8,
	0x0f, 0x64, 0xc8, 0x7d, 0xbc, 0xa1, 0xe0, 0xc4, 0xc6, 0x7f, 0x39, 0x9b, 0x38, 0x67, 0x37, 0x9b,
	0xf2, 0xe1, 0x99, 0xbb, 0xe5, 0xcf, 0x09, 0xbe, 0xd5, 0xd7, 0x3a, 0x6a, 0x03, 0x27, 0xdf, 0x10,
	0xf6, 0x0e, 0xc0, 0x92, 0xdd, 0x0b, 0x5c, 0x8c, 0x45, 0x10, 0x5c, 0xde, 0x7a, 0xf8, 0xe4, 0xd3,
	0x8f, 0x5f, 0x9f, 0x27, 0x8f, 0xc8, 0x1e, 0xcb, 0xb9, 0xe2, 0x19, 0x88, 0xf9, 0xb9, 0xcd, 0xe8,
	0x86, 0x62, 0x1f, 0xce, 0xc7, 0xfa, 0x91, 0x7c, 0x45, 0x78, 0xad, 0xb9, 0x5f, 0xf2, 0xe0, 0x82,
	0x9e, 0xe3, 0xd9, 0x06, 0x7b, 0x57, 0xc5, 0xda, 0xdc, 0xc2, 0x1d, 0xe7, 0xfb, 0x1e, 0xb9, 0xfb,
	0x5f, 0xdf, 0xfb, 0xcf, 0x5e, 0x3f, 0xcd, 0xa4, 0x7d, 0x53, 0x25, 0x34, 0xd5, 0x79, 0xf7, 0x12,
	0xe6, 0xed, 0x62, 0x67, 0x7a, 0x9e, 0x81, 0x72, 0x3b, 0x3d, 0xfe, 0x44, 0x1e, 0xbb, 0x8f, 0xe4,
	0x9a, 0x3b, 0xb6, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x67, 0x6e, 0xb1, 0xfb, 0xbe, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourcePresetServiceClient is the client API for ResourcePresetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourcePresetServiceClient interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error)
}

type resourcePresetServiceClient struct {
	cc *grpc.ClientConn
}

func NewResourcePresetServiceClient(cc *grpc.ClientConn) ResourcePresetServiceClient {
	return &resourcePresetServiceClient{cc}
}

func (c *resourcePresetServiceClient) Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error) {
	out := new(ResourcePreset)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.redis.v1.ResourcePresetService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcePresetServiceClient) List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error) {
	out := new(ListResourcePresetsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.redis.v1.ResourcePresetService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcePresetServiceServer is the server API for ResourcePresetService service.
type ResourcePresetServiceServer interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error)
}

// UnimplementedResourcePresetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResourcePresetServiceServer struct {
}

func (*UnimplementedResourcePresetServiceServer) Get(ctx context.Context, req *GetResourcePresetRequest) (*ResourcePreset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedResourcePresetServiceServer) List(ctx context.Context, req *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterResourcePresetServiceServer(s *grpc.Server, srv ResourcePresetServiceServer) {
	s.RegisterService(&_ResourcePresetService_serviceDesc, srv)
}

func _ResourcePresetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.redis.v1.ResourcePresetService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).Get(ctx, req.(*GetResourcePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourcePresetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcePresetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.redis.v1.ResourcePresetService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).List(ctx, req.(*ListResourcePresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourcePresetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.redis.v1.ResourcePresetService",
	HandlerType: (*ResourcePresetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ResourcePresetService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResourcePresetService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/redis/v1/resource_preset_service.proto",
}

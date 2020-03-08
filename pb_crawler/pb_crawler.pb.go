// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_crawler.proto

package pb_crawler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type StartRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac1a34fb274b16d, []int{0}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac1a34fb274b16d, []int{1}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac1a34fb274b16d, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ControlResponse struct {
	Started              bool     `protobuf:"varint,1,opt,name=started,proto3" json:"started,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlResponse) Reset()         { *m = ControlResponse{} }
func (m *ControlResponse) String() string { return proto.CompactTextString(m) }
func (*ControlResponse) ProtoMessage()    {}
func (*ControlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac1a34fb274b16d, []int{3}
}

func (m *ControlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlResponse.Unmarshal(m, b)
}
func (m *ControlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlResponse.Marshal(b, m, deterministic)
}
func (m *ControlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlResponse.Merge(m, src)
}
func (m *ControlResponse) XXX_Size() int {
	return xxx_messageInfo_ControlResponse.Size(m)
}
func (m *ControlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlResponse proto.InternalMessageInfo

func (m *ControlResponse) GetStarted() bool {
	if m != nil {
		return m.Started
	}
	return false
}

type ListResponse struct {
	Urls                 string   `protobuf:"bytes,1,opt,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac1a34fb274b16d, []int{4}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetUrls() string {
	if m != nil {
		return m.Urls
	}
	return ""
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "pb_crawler.StartRequest")
	proto.RegisterType((*StopRequest)(nil), "pb_crawler.StopRequest")
	proto.RegisterType((*ListRequest)(nil), "pb_crawler.ListRequest")
	proto.RegisterType((*ControlResponse)(nil), "pb_crawler.ControlResponse")
	proto.RegisterType((*ListResponse)(nil), "pb_crawler.ListResponse")
}

func init() {
	proto.RegisterFile("pb_crawler.proto", fileDescriptor_cac1a34fb274b16d)
}

var fileDescriptor_cac1a34fb274b16d = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0x8a, 0x4f,
	0x2e, 0x4a, 0x2c, 0xcf, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x70, 0xf1, 0x04, 0x97, 0x24, 0x16, 0x95, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x70, 0x31, 0x97, 0x16, 0xe5, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98,
	0x4a, 0xbc, 0x5c, 0xdc, 0xc1, 0x25, 0xf9, 0x05, 0x50, 0x05, 0x20, 0xae, 0x4f, 0x66, 0x31, 0x4c,
	0xbd, 0x92, 0x36, 0x17, 0xbf, 0x73, 0x7e, 0x5e, 0x49, 0x51, 0x7e, 0x4e, 0x50, 0x6a, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x31, 0xc8, 0xc8, 0xd4, 0x14, 0xb0, 0x31, 0x1c,
	0x41, 0x30, 0xae, 0x92, 0x12, 0x17, 0x0f, 0x44, 0x2f, 0x54, 0xa5, 0x10, 0x17, 0x4b, 0x69, 0x51,
	0x4e, 0x31, 0xd4, 0x36, 0x30, 0xdb, 0xe8, 0x16, 0x23, 0x17, 0xbb, 0x33, 0xc4, 0x71, 0x42, 0xee,
	0x5c, 0x3c, 0x50, 0x26, 0xd8, 0x8d, 0x42, 0x12, 0x7a, 0x48, 0x7e, 0x41, 0x76, 0xb6, 0x94, 0x34,
	0xb2, 0x0c, 0x9a, 0x83, 0x94, 0x18, 0x84, 0x5c, 0xb9, 0xb8, 0xe1, 0x06, 0xe5, 0x17, 0x08, 0x89,
	0xa3, 0x9a, 0x03, 0xf7, 0x1c, 0x21, 0x63, 0xec, 0xb9, 0x38, 0x40, 0xee, 0x0f, 0x29, 0x4a, 0x4d,
	0x45, 0x35, 0x03, 0x29, 0x44, 0xa4, 0x24, 0x30, 0x25, 0x60, 0x06, 0x24, 0xb1, 0x81, 0x23, 0xc0,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x6f, 0x26, 0xdf, 0x94, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrawlerClient is the client API for Crawler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrawlerClient interface {
	// Start the crawler with url
	CrawlerStart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*ControlResponse, error)
	// Stop the crawler
	CrawlerStop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*ControlResponse, error)
	// List the Urls tree
	ListTree(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type crawlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlerClient(cc grpc.ClientConnInterface) CrawlerClient {
	return &crawlerClient{cc}
}

func (c *crawlerClient) CrawlerStart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/pb_crawler.Crawler/CrawlerStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerClient) CrawlerStop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/pb_crawler.Crawler/CrawlerStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerClient) ListTree(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pb_crawler.Crawler/ListTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrawlerServer is the server API for Crawler service.
type CrawlerServer interface {
	// Start the crawler with url
	CrawlerStart(context.Context, *StartRequest) (*ControlResponse, error)
	// Stop the crawler
	CrawlerStop(context.Context, *StopRequest) (*ControlResponse, error)
	// List the Urls tree
	ListTree(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedCrawlerServer can be embedded to have forward compatible implementations.
type UnimplementedCrawlerServer struct {
}

func (*UnimplementedCrawlerServer) CrawlerStart(ctx context.Context, req *StartRequest) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrawlerStart not implemented")
}
func (*UnimplementedCrawlerServer) CrawlerStop(ctx context.Context, req *StopRequest) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrawlerStop not implemented")
}
func (*UnimplementedCrawlerServer) ListTree(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTree not implemented")
}

func RegisterCrawlerServer(s *grpc.Server, srv CrawlerServer) {
	s.RegisterService(&_Crawler_serviceDesc, srv)
}

func _Crawler_CrawlerStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServer).CrawlerStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_crawler.Crawler/CrawlerStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServer).CrawlerStart(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crawler_CrawlerStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServer).CrawlerStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_crawler.Crawler/CrawlerStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServer).CrawlerStop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crawler_ListTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServer).ListTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_crawler.Crawler/ListTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServer).ListTree(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crawler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_crawler.Crawler",
	HandlerType: (*CrawlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrawlerStart",
			Handler:    _Crawler_CrawlerStart_Handler,
		},
		{
			MethodName: "CrawlerStop",
			Handler:    _Crawler_CrawlerStop_Handler,
		},
		{
			MethodName: "ListTree",
			Handler:    _Crawler_ListTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_crawler.proto",
}

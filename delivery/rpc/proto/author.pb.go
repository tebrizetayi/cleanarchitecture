// Code generated by protoc-gen-go. DO NOT EDIT.
// source: author.proto

package author

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

type AuthorId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorId) Reset()         { *m = AuthorId{} }
func (m *AuthorId) String() string { return proto.CompactTextString(m) }
func (*AuthorId) ProtoMessage()    {}
func (*AuthorId) Descriptor() ([]byte, []int) {
	return fileDescriptor_54cb4b5a5dd8ee3d, []int{0}
}

func (m *AuthorId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorId.Unmarshal(m, b)
}
func (m *AuthorId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorId.Marshal(b, m, deterministic)
}
func (m *AuthorId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorId.Merge(m, src)
}
func (m *AuthorId) XXX_Size() int {
	return xxx_messageInfo_AuthorId.Size(m)
}
func (m *AuthorId) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorId.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorId proto.InternalMessageInfo

func (m *AuthorId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AuthorIds struct {
	AuthorId             []*AuthorId `protobuf:"bytes,1,rep,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AuthorIds) Reset()         { *m = AuthorIds{} }
func (m *AuthorIds) String() string { return proto.CompactTextString(m) }
func (*AuthorIds) ProtoMessage()    {}
func (*AuthorIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_54cb4b5a5dd8ee3d, []int{1}
}

func (m *AuthorIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorIds.Unmarshal(m, b)
}
func (m *AuthorIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorIds.Marshal(b, m, deterministic)
}
func (m *AuthorIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorIds.Merge(m, src)
}
func (m *AuthorIds) XXX_Size() int {
	return xxx_messageInfo_AuthorIds.Size(m)
}
func (m *AuthorIds) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorIds.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorIds proto.InternalMessageInfo

func (m *AuthorIds) GetAuthorId() []*AuthorId {
	if m != nil {
		return m.AuthorId
	}
	return nil
}

type Author struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_54cb4b5a5dd8ee3d, []int{2}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Authors struct {
	Author               []*Author `protobuf:"bytes,1,rep,name=Author,proto3" json:"Author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Authors) Reset()         { *m = Authors{} }
func (m *Authors) String() string { return proto.CompactTextString(m) }
func (*Authors) ProtoMessage()    {}
func (*Authors) Descriptor() ([]byte, []int) {
	return fileDescriptor_54cb4b5a5dd8ee3d, []int{3}
}

func (m *Authors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authors.Unmarshal(m, b)
}
func (m *Authors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authors.Marshal(b, m, deterministic)
}
func (m *Authors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authors.Merge(m, src)
}
func (m *Authors) XXX_Size() int {
	return xxx_messageInfo_Authors.Size(m)
}
func (m *Authors) XXX_DiscardUnknown() {
	xxx_messageInfo_Authors.DiscardUnknown(m)
}

var xxx_messageInfo_Authors proto.InternalMessageInfo

func (m *Authors) GetAuthor() []*Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthorId)(nil), "author.AuthorId")
	proto.RegisterType((*AuthorIds)(nil), "author.AuthorIds")
	proto.RegisterType((*Author)(nil), "author.Author")
	proto.RegisterType((*Authors)(nil), "author.Authors")
}

func init() {
	proto.RegisterFile("author.proto", fileDescriptor_54cb4b5a5dd8ee3d)
}

var fileDescriptor_54cb4b5a5dd8ee3d = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xa4, 0xb8, 0x38,
	0x1c, 0xc1, 0x2c, 0xcf, 0x14, 0x21, 0x3e, 0x2e, 0x26, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x20, 0x4b, 0xc9, 0x92, 0x8b, 0x13, 0x26, 0x57, 0x2c, 0xa4, 0x83, 0x50, 0x08, 0x54,
	0xc2, 0xac, 0xc1, 0x6d, 0x24, 0xa0, 0x07, 0x35, 0x11, 0x26, 0x1e, 0x04, 0x57, 0xa1, 0xa4, 0xc3,
	0xc5, 0x06, 0x61, 0xa3, 0x1b, 0x2a, 0x24, 0xc4, 0xc5, 0xe2, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04,
	0x16, 0x01, 0xb3, 0x95, 0x0c, 0xb9, 0xd8, 0x21, 0xaa, 0x8b, 0x85, 0xd4, 0x60, 0x1a, 0xa1, 0x96,
	0xf0, 0xa1, 0x5a, 0x12, 0x04, 0x95, 0x35, 0xca, 0xe1, 0xe2, 0x82, 0xb0, 0xdc, 0x8b, 0x0a, 0x92,
	0x85, 0x0c, 0xb8, 0x38, 0xdc, 0x53, 0x4b, 0x9c, 0x2a, 0x41, 0x0e, 0x15, 0x44, 0x77, 0x56, 0xb1,
	0x14, 0x3f, 0xaa, 0x50, 0xb1, 0x12, 0x03, 0xd0, 0x3b, 0x6c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9,
	0x42, 0xe8, 0x92, 0x58, 0x54, 0x27, 0xb1, 0x81, 0x03, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0xb6, 0xb4, 0xff, 0x44, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorGrpcClient is the client API for AuthorGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorGrpcClient interface {
	GetByIds(ctx context.Context, in *AuthorIds, opts ...grpc.CallOption) (*Authors, error)
	Create(ctx context.Context, in *Authors, opts ...grpc.CallOption) (*Authors, error)
}

type authorGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorGrpcClient(cc grpc.ClientConnInterface) AuthorGrpcClient {
	return &authorGrpcClient{cc}
}

func (c *authorGrpcClient) GetByIds(ctx context.Context, in *AuthorIds, opts ...grpc.CallOption) (*Authors, error) {
	out := new(Authors)
	err := c.cc.Invoke(ctx, "/author.AuthorGrpc/GetByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorGrpcClient) Create(ctx context.Context, in *Authors, opts ...grpc.CallOption) (*Authors, error) {
	out := new(Authors)
	err := c.cc.Invoke(ctx, "/author.AuthorGrpc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorGrpcServer is the server API for AuthorGrpc service.
type AuthorGrpcServer interface {
	GetByIds(context.Context, *AuthorIds) (*Authors, error)
	Create(context.Context, *Authors) (*Authors, error)
}

// UnimplementedAuthorGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorGrpcServer struct {
}

func (*UnimplementedAuthorGrpcServer) GetByIds(ctx context.Context, req *AuthorIds) (*Authors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIds not implemented")
}
func (*UnimplementedAuthorGrpcServer) Create(ctx context.Context, req *Authors) (*Authors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterAuthorGrpcServer(s *grpc.Server, srv AuthorGrpcServer) {
	s.RegisterService(&_AuthorGrpc_serviceDesc, srv)
}

func _AuthorGrpc_GetByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorGrpcServer).GetByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author.AuthorGrpc/GetByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorGrpcServer).GetByIds(ctx, req.(*AuthorIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorGrpc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorGrpcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author.AuthorGrpc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorGrpcServer).Create(ctx, req.(*Authors))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "author.AuthorGrpc",
	HandlerType: (*AuthorGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByIds",
			Handler:    _AuthorGrpc_GetByIds_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AuthorGrpc_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author.proto",
}

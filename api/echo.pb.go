// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

type Pong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Ping)(nil), "api.Ping")
	proto.RegisterType((*Pong)(nil), "api.Pong")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x62, 0xe3, 0x62, 0x09,
	0xc8, 0xcc, 0x4b, 0x07, 0xd3, 0xf9, 0x79, 0xe9, 0x46, 0x3a, 0x5c, 0x5c, 0xc1, 0xa9, 0xc9, 0xa5,
	0x45, 0xa9, 0x20, 0x51, 0x21, 0x39, 0x2e, 0x16, 0xd7, 0xe4, 0x8c, 0x7c, 0x21, 0x4e, 0xbd, 0xc4,
	0x82, 0x4c, 0x3d, 0x90, 0x90, 0x14, 0x94, 0x99, 0x9f, 0x97, 0xae, 0xc4, 0x90, 0xc4, 0x06, 0x36,
	0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x23, 0x0f, 0x72, 0x55, 0x57, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecurePingClient is the client API for SecurePing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecurePingClient interface {
	Echo(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type securePingClient struct {
	cc *grpc.ClientConn
}

func NewSecurePingClient(cc *grpc.ClientConn) SecurePingClient {
	return &securePingClient{cc}
}

func (c *securePingClient) Echo(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/api.SecurePing/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurePingServer is the server API for SecurePing service.
type SecurePingServer interface {
	Echo(context.Context, *Ping) (*Pong, error)
}

func RegisterSecurePingServer(s *grpc.Server, srv SecurePingServer) {
	s.RegisterService(&_SecurePing_serviceDesc, srv)
}

func _SecurePing_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurePingServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SecurePing/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurePingServer).Echo(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurePing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SecurePing",
	HandlerType: (*SecurePingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _SecurePing_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

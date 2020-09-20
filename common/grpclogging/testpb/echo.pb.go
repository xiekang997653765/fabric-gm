// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package testpb

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

type Message struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Sequence             int32    `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Message) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "testpb.Message")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4b, 0xc5, 0x30,
	0x10, 0xc6, 0x89, 0xe8, 0x7b, 0x7a, 0x0e, 0x42, 0xa6, 0xd2, 0xa9, 0x38, 0x65, 0x4a, 0xa4, 0x0e,
	0xe2, 0x24, 0x08, 0x8e, 0x2e, 0x75, 0x73, 0x6b, 0xce, 0x33, 0x09, 0x36, 0xbd, 0x98, 0xa4, 0x82,
	0xff, 0xbd, 0xd8, 0xaa, 0x83, 0xcb, 0xdb, 0xee, 0x77, 0xdc, 0xfd, 0x3e, 0x3e, 0x00, 0x42, 0xcf,
	0x3a, 0x65, 0xae, 0x2c, 0x77, 0x95, 0x4a, 0x4d, 0xf6, 0xf2, 0x0e, 0xf6, 0x8f, 0x54, 0xca, 0xe8,
	0x48, 0x36, 0xb0, 0x8f, 0xdb, 0xd8, 0x88, 0x4e, 0xa8, 0xb3, 0xe1, 0x17, 0x65, 0x0b, 0xa7, 0x85,
	0xde, 0x17, 0x9a, 0x91, 0x9a, 0xa3, 0x4e, 0xa8, 0x93, 0xe1, 0x8f, 0xfb, 0x37, 0x38, 0x7f, 0x40,
	0xcf, 0x4f, 0x94, 0x3f, 0x02, 0x92, 0x54, 0x70, 0xfc, 0x8d, 0xf2, 0x42, 0x6f, 0x01, 0xfa, 0xc7,
	0xde, 0xfe, 0x5f, 0xc8, 0x1e, 0x60, 0x7d, 0xac, 0x99, 0xc6, 0x78, 0xf8, 0x5e, 0x89, 0x2b, 0x71,
	0x7f, 0xfb, 0x7c, 0xe3, 0x42, 0xf5, 0x8b, 0xd5, 0xc8, 0xd1, 0xf8, 0xcf, 0x44, 0x79, 0xa2, 0x17,
	0x47, 0xd9, 0xbc, 0x8e, 0x36, 0x07, 0x34, 0xc8, 0x31, 0xf2, 0x6c, 0x5c, 0x4e, 0x38, 0xb1, 0x73,
	0x61, 0x76, 0x66, 0xd3, 0xd8, 0xdd, 0xda, 0xfb, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x33,
	0xa2, 0xe6, 0x05, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EchoStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoStreamClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/testpb.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoStream(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[0], "/testpb.EchoService/EchoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoStreamClient{stream}
	return x, nil
}

type EchoService_EchoStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type echoServiceEchoStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	Echo(context.Context, *Message) (*Message, error)
	EchoStream(EchoService_EchoStreamServer) error
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testpb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).EchoStream(&echoServiceEchoStreamServer{stream})
}

type EchoService_EchoStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type echoServiceEchoStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testpb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoStream",
			Handler:       _EchoService_EchoStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: counter.proto

package counter

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

type Command struct {
	CommandName          string   `protobuf:"bytes,1,opt,name=commandName,proto3" json:"commandName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCommandName() string {
	if m != nil {
		return m.CommandName
	}
	return ""
}

type StreamField struct {
	FieldName            string   `protobuf:"bytes,1,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamField) Reset()         { *m = StreamField{} }
func (m *StreamField) String() string { return proto.CompactTextString(m) }
func (*StreamField) ProtoMessage()    {}
func (*StreamField) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{1}
}

func (m *StreamField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamField.Unmarshal(m, b)
}
func (m *StreamField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamField.Marshal(b, m, deterministic)
}
func (m *StreamField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamField.Merge(m, src)
}
func (m *StreamField) XXX_Size() int {
	return xxx_messageInfo_StreamField.Size(m)
}
func (m *StreamField) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamField.DiscardUnknown(m)
}

var xxx_messageInfo_StreamField proto.InternalMessageInfo

func (m *StreamField) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

type Count struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{2}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Count) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type CountNote struct {
	Instance             *Count   `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	CommandName          string   `protobuf:"bytes,2,opt,name=commandName,proto3" json:"commandName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountNote) Reset()         { *m = CountNote{} }
func (m *CountNote) String() string { return proto.CompactTextString(m) }
func (*CountNote) ProtoMessage()    {}
func (*CountNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{3}
}

func (m *CountNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountNote.Unmarshal(m, b)
}
func (m *CountNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountNote.Marshal(b, m, deterministic)
}
func (m *CountNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountNote.Merge(m, src)
}
func (m *CountNote) XXX_Size() int {
	return xxx_messageInfo_CountNote.Size(m)
}
func (m *CountNote) XXX_DiscardUnknown() {
	xxx_messageInfo_CountNote.DiscardUnknown(m)
}

var xxx_messageInfo_CountNote proto.InternalMessageInfo

func (m *CountNote) GetInstance() *Count {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CountNote) GetCommandName() string {
	if m != nil {
		return m.CommandName
	}
	return ""
}

type Statistics struct {
	AmountValue          int32    `protobuf:"varint,1,opt,name=amountValue,proto3" json:"amountValue,omitempty"`
	MinValue             int32    `protobuf:"varint,2,opt,name=minValue,proto3" json:"minValue,omitempty"`
	MaxValue             int32    `protobuf:"varint,3,opt,name=maxValue,proto3" json:"maxValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statistics) Reset()         { *m = Statistics{} }
func (m *Statistics) String() string { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()    {}
func (*Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{4}
}

func (m *Statistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statistics.Unmarshal(m, b)
}
func (m *Statistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statistics.Marshal(b, m, deterministic)
}
func (m *Statistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statistics.Merge(m, src)
}
func (m *Statistics) XXX_Size() int {
	return xxx_messageInfo_Statistics.Size(m)
}
func (m *Statistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Statistics.DiscardUnknown(m)
}

var xxx_messageInfo_Statistics proto.InternalMessageInfo

func (m *Statistics) GetAmountValue() int32 {
	if m != nil {
		return m.AmountValue
	}
	return 0
}

func (m *Statistics) GetMinValue() int32 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *Statistics) GetMaxValue() int32 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func init() {
	proto.RegisterType((*Command)(nil), "counter.Command")
	proto.RegisterType((*StreamField)(nil), "counter.StreamField")
	proto.RegisterType((*Count)(nil), "counter.Count")
	proto.RegisterType((*CountNote)(nil), "counter.CountNote")
	proto.RegisterType((*Statistics)(nil), "counter.Statistics")
}

func init() { proto.RegisterFile("counter.proto", fileDescriptor_75dcd656fce7132f) }

var fileDescriptor_75dcd656fce7132f = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcb, 0x6b, 0x32, 0x31,
	0x10, 0x37, 0x7e, 0xf8, 0xa9, 0x63, 0x5f, 0x4c, 0x3d, 0xc8, 0xd2, 0x83, 0xe4, 0x24, 0x15, 0x44,
	0xec, 0xd1, 0x42, 0x29, 0x42, 0xdb, 0x93, 0x07, 0x85, 0x42, 0x8f, 0x71, 0x37, 0x42, 0x60, 0x93,
	0xc8, 0x6e, 0xb6, 0xf4, 0x1f, 0xef, 0xbd, 0xe4, 0xe1, 0x9a, 0xae, 0xbd, 0x65, 0x7e, 0x8f, 0xc9,
	0xfc, 0x26, 0x81, 0xcb, 0x54, 0x57, 0xca, 0xf0, 0x62, 0x76, 0x28, 0xb4, 0xd1, 0xd8, 0x0d, 0x25,
	0x9d, 0x42, 0x77, 0xa5, 0xa5, 0x64, 0x2a, 0xc3, 0x31, 0x0c, 0x52, 0x7f, 0x5c, 0x33, 0xc9, 0x47,
	0x64, 0x4c, 0x26, 0xfd, 0x4d, 0x0c, 0xd1, 0x29, 0x0c, 0xb6, 0xa6, 0xe0, 0x4c, 0xbe, 0x08, 0x9e,
	0x67, 0x78, 0x07, 0xfd, 0xbd, 0x3d, 0x44, 0xf2, 0x13, 0x40, 0x97, 0xd0, 0x59, 0xd9, 0x4b, 0x70,
	0x08, 0x9d, 0x4f, 0x96, 0x57, 0x5e, 0xd2, 0xd9, 0xf8, 0xc2, 0x9a, 0x8d, 0x90, 0xbc, 0x34, 0x4c,
	0x1e, 0x46, 0x6d, 0x6f, 0xae, 0x01, 0xfa, 0x01, 0x7d, 0x67, 0x5e, 0x6b, 0xc3, 0xf1, 0x1e, 0x7a,
	0x42, 0x95, 0x86, 0xa9, 0xd4, 0xf7, 0x18, 0x2c, 0xae, 0x66, 0xc7, 0x38, 0x4e, 0xb5, 0xa9, 0xf9,
	0x66, 0x88, 0xf6, 0x79, 0x88, 0x3d, 0xc0, 0xd6, 0x30, 0x23, 0x4a, 0x23, 0xd2, 0xd2, 0xea, 0x99,
	0xb4, 0x3d, 0xde, 0xa3, 0x11, 0x63, 0x08, 0x13, 0xe8, 0x49, 0xa1, 0x3c, 0xdd, 0x76, 0x74, 0x5d,
	0x3b, 0x8e, 0x7d, 0x79, 0xee, 0x5f, 0xe0, 0x42, 0xbd, 0xf8, 0x26, 0x76, 0xb5, 0x6e, 0x4a, 0x9c,
	0x41, 0xef, 0x95, 0x1b, 0xbf, 0x8e, 0x9b, 0x68, 0x76, 0x37, 0x55, 0xd2, 0x48, 0x43, 0x5b, 0xb8,
	0x84, 0xeb, 0xb0, 0xe8, 0x42, 0xcb, 0xb0, 0xc5, 0x5a, 0x14, 0x3d, 0xc1, 0xb9, 0x75, 0x4e, 0xf0,
	0x09, 0x86, 0xa7, 0x80, 0xcf, 0x3b, 0x5d, 0x85, 0x8b, 0x1b, 0xda, 0xe4, 0x36, 0xea, 0x78, 0x94,
	0xd3, 0xd6, 0x84, 0xe0, 0x23, 0x5c, 0x38, 0xc5, 0x1b, 0x53, 0x59, 0xce, 0x0b, 0xc4, 0xdf, 0x46,
	0xfb, 0x26, 0xc9, 0x1f, 0x98, 0xf5, 0xce, 0xc9, 0xee, 0xbf, 0xfb, 0x61, 0x0f, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0x6f, 0x64, 0x49, 0x72, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterClient interface {
	GetCount(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Count, error)
	StreamFromCount(ctx context.Context, in *StreamField, opts ...grpc.CallOption) (Counter_StreamFromCountClient, error)
	StatisticsAboutCount(ctx context.Context, opts ...grpc.CallOption) (Counter_StatisticsAboutCountClient, error)
	CountHandler(ctx context.Context, opts ...grpc.CallOption) (Counter_CountHandlerClient, error)
}

type counterClient struct {
	cc *grpc.ClientConn
}

func NewCounterClient(cc *grpc.ClientConn) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) GetCount(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/counter.Counter/GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) StreamFromCount(ctx context.Context, in *StreamField, opts ...grpc.CallOption) (Counter_StreamFromCountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Counter_serviceDesc.Streams[0], "/counter.Counter/StreamFromCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterStreamFromCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Counter_StreamFromCountClient interface {
	Recv() (*Count, error)
	grpc.ClientStream
}

type counterStreamFromCountClient struct {
	grpc.ClientStream
}

func (x *counterStreamFromCountClient) Recv() (*Count, error) {
	m := new(Count)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *counterClient) StatisticsAboutCount(ctx context.Context, opts ...grpc.CallOption) (Counter_StatisticsAboutCountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Counter_serviceDesc.Streams[1], "/counter.Counter/StatisticsAboutCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterStatisticsAboutCountClient{stream}
	return x, nil
}

type Counter_StatisticsAboutCountClient interface {
	Send(*Count) error
	CloseAndRecv() (*Statistics, error)
	grpc.ClientStream
}

type counterStatisticsAboutCountClient struct {
	grpc.ClientStream
}

func (x *counterStatisticsAboutCountClient) Send(m *Count) error {
	return x.ClientStream.SendMsg(m)
}

func (x *counterStatisticsAboutCountClient) CloseAndRecv() (*Statistics, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Statistics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *counterClient) CountHandler(ctx context.Context, opts ...grpc.CallOption) (Counter_CountHandlerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Counter_serviceDesc.Streams[2], "/counter.Counter/CountHandler", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterCountHandlerClient{stream}
	return x, nil
}

type Counter_CountHandlerClient interface {
	Send(*CountNote) error
	Recv() (*CountNote, error)
	grpc.ClientStream
}

type counterCountHandlerClient struct {
	grpc.ClientStream
}

func (x *counterCountHandlerClient) Send(m *CountNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *counterCountHandlerClient) Recv() (*CountNote, error) {
	m := new(CountNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CounterServer is the server API for Counter service.
type CounterServer interface {
	GetCount(context.Context, *Command) (*Count, error)
	StreamFromCount(*StreamField, Counter_StreamFromCountServer) error
	StatisticsAboutCount(Counter_StatisticsAboutCountServer) error
	CountHandler(Counter_CountHandlerServer) error
}

// UnimplementedCounterServer can be embedded to have forward compatible implementations.
type UnimplementedCounterServer struct {
}

func (*UnimplementedCounterServer) GetCount(ctx context.Context, req *Command) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}
func (*UnimplementedCounterServer) StreamFromCount(req *StreamField, srv Counter_StreamFromCountServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamFromCount not implemented")
}
func (*UnimplementedCounterServer) StatisticsAboutCount(srv Counter_StatisticsAboutCountServer) error {
	return status.Errorf(codes.Unimplemented, "method StatisticsAboutCount not implemented")
}
func (*UnimplementedCounterServer) CountHandler(srv Counter_CountHandlerServer) error {
	return status.Errorf(codes.Unimplemented, "method CountHandler not implemented")
}

func RegisterCounterServer(s *grpc.Server, srv CounterServer) {
	s.RegisterService(&_Counter_serviceDesc, srv)
}

func _Counter_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Counter/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).GetCount(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_StreamFromCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamField)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CounterServer).StreamFromCount(m, &counterStreamFromCountServer{stream})
}

type Counter_StreamFromCountServer interface {
	Send(*Count) error
	grpc.ServerStream
}

type counterStreamFromCountServer struct {
	grpc.ServerStream
}

func (x *counterStreamFromCountServer) Send(m *Count) error {
	return x.ServerStream.SendMsg(m)
}

func _Counter_StatisticsAboutCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CounterServer).StatisticsAboutCount(&counterStatisticsAboutCountServer{stream})
}

type Counter_StatisticsAboutCountServer interface {
	SendAndClose(*Statistics) error
	Recv() (*Count, error)
	grpc.ServerStream
}

type counterStatisticsAboutCountServer struct {
	grpc.ServerStream
}

func (x *counterStatisticsAboutCountServer) SendAndClose(m *Statistics) error {
	return x.ServerStream.SendMsg(m)
}

func (x *counterStatisticsAboutCountServer) Recv() (*Count, error) {
	m := new(Count)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Counter_CountHandler_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CounterServer).CountHandler(&counterCountHandlerServer{stream})
}

type Counter_CountHandlerServer interface {
	Send(*CountNote) error
	Recv() (*CountNote, error)
	grpc.ServerStream
}

type counterCountHandlerServer struct {
	grpc.ServerStream
}

func (x *counterCountHandlerServer) Send(m *CountNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *counterCountHandlerServer) Recv() (*CountNote, error) {
	m := new(CountNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Counter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "counter.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _Counter_GetCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamFromCount",
			Handler:       _Counter_StreamFromCount_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StatisticsAboutCount",
			Handler:       _Counter_StatisticsAboutCount_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CountHandler",
			Handler:       _Counter_CountHandler_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "counter.proto",
}

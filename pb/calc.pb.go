// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/calc.proto

/*
Package calcpb is a generated protocol buffer package.

It is generated from these files:
	pb/calc.proto

It has these top-level messages:
	FindMaximumRequest
	FindMaximumResponse
	ComputeAverageRequest
	ComputeAverageResponse
	PrimeNumberDecomposerRequest
	PrimeNumberDecomposerResponse
	CalculateRequest
	CalculateResponse
	CalculateWithDeadlineRequest
	CalculateWithDeadlineResponse
*/
package calcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Operation int32

const (
	Operation_UNSET    Operation = 0
	Operation_SUM      Operation = 1
	Operation_DIFF     Operation = 2
	Operation_DIVIDE   Operation = 3
	Operation_MULTIPLY Operation = 4
)

var Operation_name = map[int32]string{
	0: "UNSET",
	1: "SUM",
	2: "DIFF",
	3: "DIVIDE",
	4: "MULTIPLY",
}
var Operation_value = map[string]int32{
	"UNSET":    0,
	"SUM":      1,
	"DIFF":     2,
	"DIVIDE":   3,
	"MULTIPLY": 4,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FindMaximumRequest struct {
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *FindMaximumRequest) Reset()                    { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string            { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()               {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FindMaximumRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FindMaximumResponse struct {
	Max int32 `protobuf:"varint,1,opt,name=max" json:"max,omitempty"`
}

func (m *FindMaximumResponse) Reset()                    { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string            { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()               {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FindMaximumResponse) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type ComputeAverageRequest struct {
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *ComputeAverageRequest) Reset()                    { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string            { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()               {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ComputeAverageRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ComputeAverageResponse struct {
	Result float32 `protobuf:"fixed32,1,opt,name=result" json:"result,omitempty"`
}

func (m *ComputeAverageResponse) Reset()                    { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string            { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()               {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ComputeAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberDecomposerRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *PrimeNumberDecomposerRequest) Reset()                    { *m = PrimeNumberDecomposerRequest{} }
func (m *PrimeNumberDecomposerRequest) String() string            { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposerRequest) ProtoMessage()               {}
func (*PrimeNumberDecomposerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PrimeNumberDecomposerRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecomposerResponse struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *PrimeNumberDecomposerResponse) Reset()                    { *m = PrimeNumberDecomposerResponse{} }
func (m *PrimeNumberDecomposerResponse) String() string            { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposerResponse) ProtoMessage()               {}
func (*PrimeNumberDecomposerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PrimeNumberDecomposerResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculateRequest struct {
	Num1      float32   `protobuf:"fixed32,1,opt,name=num1" json:"num1,omitempty"`
	Num2      float32   `protobuf:"fixed32,2,opt,name=num2" json:"num2,omitempty"`
	Operation Operation `protobuf:"varint,3,opt,name=operation,enum=calc.Operation" json:"operation,omitempty"`
}

func (m *CalculateRequest) Reset()                    { *m = CalculateRequest{} }
func (m *CalculateRequest) String() string            { return proto.CompactTextString(m) }
func (*CalculateRequest) ProtoMessage()               {}
func (*CalculateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CalculateRequest) GetNum1() float32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *CalculateRequest) GetNum2() float32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

func (m *CalculateRequest) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_UNSET
}

type CalculateResponse struct {
	Result float32 `protobuf:"fixed32,1,opt,name=result" json:"result,omitempty"`
}

func (m *CalculateResponse) Reset()                    { *m = CalculateResponse{} }
func (m *CalculateResponse) String() string            { return proto.CompactTextString(m) }
func (*CalculateResponse) ProtoMessage()               {}
func (*CalculateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CalculateResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculateWithDeadlineRequest struct {
	Num1      float32   `protobuf:"fixed32,1,opt,name=num1" json:"num1,omitempty"`
	Num2      float32   `protobuf:"fixed32,2,opt,name=num2" json:"num2,omitempty"`
	Operation Operation `protobuf:"varint,3,opt,name=operation,enum=calc.Operation" json:"operation,omitempty"`
}

func (m *CalculateWithDeadlineRequest) Reset()                    { *m = CalculateWithDeadlineRequest{} }
func (m *CalculateWithDeadlineRequest) String() string            { return proto.CompactTextString(m) }
func (*CalculateWithDeadlineRequest) ProtoMessage()               {}
func (*CalculateWithDeadlineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CalculateWithDeadlineRequest) GetNum1() float32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *CalculateWithDeadlineRequest) GetNum2() float32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

func (m *CalculateWithDeadlineRequest) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_UNSET
}

type CalculateWithDeadlineResponse struct {
	Result float32 `protobuf:"fixed32,1,opt,name=result" json:"result,omitempty"`
}

func (m *CalculateWithDeadlineResponse) Reset()                    { *m = CalculateWithDeadlineResponse{} }
func (m *CalculateWithDeadlineResponse) String() string            { return proto.CompactTextString(m) }
func (*CalculateWithDeadlineResponse) ProtoMessage()               {}
func (*CalculateWithDeadlineResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CalculateWithDeadlineResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*FindMaximumRequest)(nil), "calc.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calc.FindMaximumResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calc.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calc.ComputeAverageResponse")
	proto.RegisterType((*PrimeNumberDecomposerRequest)(nil), "calc.PrimeNumberDecomposerRequest")
	proto.RegisterType((*PrimeNumberDecomposerResponse)(nil), "calc.PrimeNumberDecomposerResponse")
	proto.RegisterType((*CalculateRequest)(nil), "calc.CalculateRequest")
	proto.RegisterType((*CalculateResponse)(nil), "calc.CalculateResponse")
	proto.RegisterType((*CalculateWithDeadlineRequest)(nil), "calc.CalculateWithDeadlineRequest")
	proto.RegisterType((*CalculateWithDeadlineResponse)(nil), "calc.CalculateWithDeadlineResponse")
	proto.RegisterEnum("calc.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calc service

type CalcClient interface {
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (Calc_FindMaximumClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (Calc_ComputeAverageClient, error)
	PrimeNumberDecompose(ctx context.Context, in *PrimeNumberDecomposerRequest, opts ...grpc.CallOption) (Calc_PrimeNumberDecomposeClient, error)
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	CalculateWithDeadline(ctx context.Context, in *CalculateWithDeadlineRequest, opts ...grpc.CallOption) (*CalculateWithDeadlineResponse, error)
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (Calc_FindMaximumClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Calc_serviceDesc.Streams[0], c.cc, "/calc.Calc/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcFindMaximumClient{stream}
	return x, nil
}

type Calc_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calcFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calcFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (Calc_ComputeAverageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Calc_serviceDesc.Streams[1], c.cc, "/calc.Calc/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcComputeAverageClient{stream}
	return x, nil
}

type Calc_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calcComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calcComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcClient) PrimeNumberDecompose(ctx context.Context, in *PrimeNumberDecomposerRequest, opts ...grpc.CallOption) (Calc_PrimeNumberDecomposeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Calc_serviceDesc.Streams[2], c.cc, "/calc.Calc/PrimeNumberDecompose", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcPrimeNumberDecomposeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calc_PrimeNumberDecomposeClient interface {
	Recv() (*PrimeNumberDecomposerResponse, error)
	grpc.ClientStream
}

type calcPrimeNumberDecomposeClient struct {
	grpc.ClientStream
}

func (x *calcPrimeNumberDecomposeClient) Recv() (*PrimeNumberDecomposerResponse, error) {
	m := new(PrimeNumberDecomposerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := grpc.Invoke(ctx, "/calc.Calc/Calculate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) CalculateWithDeadline(ctx context.Context, in *CalculateWithDeadlineRequest, opts ...grpc.CallOption) (*CalculateWithDeadlineResponse, error) {
	out := new(CalculateWithDeadlineResponse)
	err := grpc.Invoke(ctx, "/calc.Calc/CalculateWithDeadline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calc service

type CalcServer interface {
	FindMaximum(Calc_FindMaximumServer) error
	ComputeAverage(Calc_ComputeAverageServer) error
	PrimeNumberDecompose(*PrimeNumberDecomposerRequest, Calc_PrimeNumberDecomposeServer) error
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	CalculateWithDeadline(context.Context, *CalculateWithDeadlineRequest) (*CalculateWithDeadlineResponse, error)
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServer).FindMaximum(&calcFindMaximumServer{stream})
}

type Calc_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calcFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calcFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calc_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServer).ComputeAverage(&calcComputeAverageServer{stream})
}

type Calc_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calcComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calcComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calc_PrimeNumberDecompose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecomposerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServer).PrimeNumberDecompose(m, &calcPrimeNumberDecomposeServer{stream})
}

type Calc_PrimeNumberDecomposeServer interface {
	Send(*PrimeNumberDecomposerResponse) error
	grpc.ServerStream
}

type calcPrimeNumberDecomposeServer struct {
	grpc.ServerStream
}

func (x *calcPrimeNumberDecomposeServer) Send(m *PrimeNumberDecomposerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calc_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_CalculateWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).CalculateWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/CalculateWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).CalculateWithDeadline(ctx, req.(*CalculateWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Calc_Calculate_Handler,
		},
		{
			MethodName: "CalculateWithDeadline",
			Handler:    _Calc_CalculateWithDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaximum",
			Handler:       _Calc_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _Calc_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PrimeNumberDecompose",
			Handler:       _Calc_PrimeNumberDecompose_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/calc.proto",
}

func init() { proto.RegisterFile("pb/calc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x6f, 0xd3, 0x40,
	0x10, 0x8c, 0x63, 0xc7, 0x24, 0x0b, 0x14, 0xb3, 0xd0, 0x10, 0x42, 0x2a, 0x55, 0x87, 0x04, 0x01,
	0x44, 0x09, 0x46, 0xa2, 0x6f, 0x48, 0x10, 0x37, 0xc2, 0x52, 0xd3, 0x56, 0x6e, 0x03, 0x82, 0xb7,
	0xb3, 0x7b, 0x02, 0x4b, 0x3e, 0x9f, 0xb1, 0x7d, 0x28, 0xbf, 0x89, 0x5f, 0x89, 0xfc, 0xd9, 0x2f,
	0xc7, 0x3c, 0xf1, 0xb6, 0xb7, 0x3b, 0x3b, 0x33, 0xc9, 0x8e, 0x0c, 0x77, 0x23, 0xf7, 0x8d, 0x47,
	0x03, 0x6f, 0x2f, 0x8a, 0x45, 0x2a, 0x50, 0xcb, 0x6a, 0xf2, 0x0c, 0x70, 0xe1, 0x87, 0xe7, 0x4b,
	0xba, 0xf6, 0xb9, 0xe4, 0x0e, 0xfb, 0x25, 0x59, 0x92, 0xa2, 0x01, 0x6a, 0x28, 0xf9, 0x48, 0xd9,
	0x55, 0xa6, 0x3d, 0x27, 0x2b, 0xc9, 0x73, 0x78, 0x70, 0x05, 0x97, 0x44, 0x22, 0x4c, 0x58, 0x06,
	0xe4, 0x74, 0x5d, 0x01, 0x39, 0x5d, 0x93, 0x17, 0xb0, 0x3d, 0x17, 0x3c, 0x92, 0x29, 0xfb, 0xf8,
	0x9b, 0xc5, 0xf4, 0x07, 0xdb, 0xcc, 0x39, 0x83, 0xe1, 0x75, 0x68, 0x49, 0x3b, 0x04, 0x3d, 0x66,
	0x89, 0x0c, 0xd2, 0x1c, 0xde, 0x75, 0xca, 0x17, 0x79, 0x0f, 0x93, 0x93, 0xd8, 0xe7, 0xec, 0x48,
	0x72, 0x97, 0xc5, 0x16, 0xf3, 0x04, 0x8f, 0x44, 0xc2, 0xe2, 0x4a, 0x63, 0x08, 0x7a, 0x98, 0x8f,
	0x4a, 0x99, 0xf2, 0x45, 0xf6, 0x61, 0x67, 0xc3, 0x5e, 0xa3, 0x60, 0xaf, 0x16, 0xf4, 0xc1, 0x98,
	0xd3, 0xc0, 0x93, 0x01, 0x4d, 0xeb, 0x1f, 0x82, 0xa0, 0x85, 0x92, 0xbf, 0x2d, 0xad, 0xe5, 0x75,
	0xd9, 0x33, 0x47, 0xdd, 0xba, 0x67, 0xe2, 0x6b, 0x18, 0x88, 0x88, 0xc5, 0x34, 0xf5, 0x45, 0x38,
	0x52, 0x77, 0x95, 0xe9, 0x96, 0x79, 0x6f, 0x2f, 0x3f, 0xc0, 0x71, 0xd5, 0x76, 0x2e, 0x10, 0xe4,
	0x15, 0xdc, 0xbf, 0x24, 0xf5, 0x8f, 0x3f, 0x42, 0xc2, 0xa4, 0x06, 0x7f, 0xf5, 0xd3, 0x9f, 0x16,
	0xa3, 0xe7, 0x81, 0x1f, 0xfe, 0x6f, 0x8f, 0xfb, 0xb0, 0xb3, 0x41, 0xb6, 0xdd, 0xef, 0xcb, 0x39,
	0x0c, 0x6a, 0x42, 0x1c, 0x40, 0x6f, 0x75, 0x74, 0x7a, 0x70, 0x66, 0x74, 0xf0, 0x16, 0xa8, 0xa7,
	0xab, 0xa5, 0xa1, 0x60, 0x1f, 0x34, 0xcb, 0x5e, 0x2c, 0x8c, 0x2e, 0x02, 0xe8, 0x96, 0xfd, 0xc5,
	0xb6, 0x0e, 0x0c, 0x15, 0xef, 0x40, 0x7f, 0xb9, 0x3a, 0x3c, 0xb3, 0x4f, 0x0e, 0xbf, 0x19, 0x9a,
	0xf9, 0x47, 0x05, 0x2d, 0x93, 0xc7, 0xcf, 0x70, 0xfb, 0x52, 0x18, 0x71, 0x54, 0x38, 0xbe, 0x99,
	0xe3, 0xf1, 0xe3, 0x86, 0x49, 0xe1, 0x94, 0x74, 0xa6, 0xca, 0x4c, 0xc1, 0x63, 0xd8, 0xba, 0x1a,
	0x41, 0x7c, 0x52, 0xac, 0x34, 0x66, 0x78, 0x3c, 0x69, 0x1e, 0x5e, 0x50, 0xa2, 0x07, 0x0f, 0x9b,
	0x92, 0x86, 0xa4, 0xd8, 0x6c, 0x4b, 0xef, 0xf8, 0x69, 0x2b, 0xa6, 0x12, 0x99, 0x29, 0xf8, 0x01,
	0x06, 0xf5, 0x19, 0x70, 0x58, 0x7a, 0xba, 0x16, 0xd3, 0xf1, 0xa3, 0x1b, 0xfd, 0x8a, 0x01, 0x5d,
	0xd8, 0x6e, 0x3c, 0x63, 0xe5, 0xb2, 0x2d, 0x5a, 0x95, 0xcb, 0xd6, 0x1c, 0x90, 0xce, 0xa7, 0xfe,
	0x77, 0x3d, 0xc3, 0x45, 0xae, 0xab, 0xe7, 0xdf, 0x9b, 0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x22, 0xa8, 0x2b, 0x3a, 0x80, 0x04, 0x00, 0x00,
}
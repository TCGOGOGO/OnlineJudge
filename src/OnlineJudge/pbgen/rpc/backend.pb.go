// Code generated by protoc-gen-go.
// source: rpc/backend.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc/backend.proto

It has these top-level messages:
	SubmitCodeRequest
	SubmitCodeResponse
*/
package rpc

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

type SubmitCodeRequest struct {
	RunId int64 `protobuf:"varint,1,opt,name=run_id,json=runId" json:"run_id,omitempty"`
}

func (m *SubmitCodeRequest) Reset()                    { *m = SubmitCodeRequest{} }
func (m *SubmitCodeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitCodeRequest) ProtoMessage()               {}
func (*SubmitCodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubmitCodeRequest) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

type SubmitCodeResponse struct {
	Received bool `protobuf:"varint,1,opt,name=received" json:"received,omitempty"`
}

func (m *SubmitCodeResponse) Reset()                    { *m = SubmitCodeResponse{} }
func (m *SubmitCodeResponse) String() string            { return proto.CompactTextString(m) }
func (*SubmitCodeResponse) ProtoMessage()               {}
func (*SubmitCodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubmitCodeResponse) GetReceived() bool {
	if m != nil {
		return m.Received
	}
	return false
}

func init() {
	proto.RegisterType((*SubmitCodeRequest)(nil), "rpc.SubmitCodeRequest")
	proto.RegisterType((*SubmitCodeResponse)(nil), "rpc.SubmitCodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BackendHelper service

type BackendHelperClient interface {
	Submit(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error)
}

type backendHelperClient struct {
	cc *grpc.ClientConn
}

func NewBackendHelperClient(cc *grpc.ClientConn) BackendHelperClient {
	return &backendHelperClient{cc}
}

func (c *backendHelperClient) Submit(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error) {
	out := new(SubmitCodeResponse)
	err := grpc.Invoke(ctx, "/rpc.BackendHelper/Submit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BackendHelper service

type BackendHelperServer interface {
	Submit(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error)
}

func RegisterBackendHelperServer(s *grpc.Server, srv BackendHelperServer) {
	s.RegisterService(&_BackendHelper_serviceDesc, srv)
}

func _BackendHelper_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendHelperServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.BackendHelper/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendHelperServer).Submit(ctx, req.(*SubmitCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackendHelper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.BackendHelper",
	HandlerType: (*BackendHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _BackendHelper_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/backend.proto",
}

func init() { proto.RegisterFile("rpc/backend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2a, 0x48, 0xd6,
	0x4f, 0x4a, 0x4c, 0xce, 0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e,
	0x2a, 0x48, 0x56, 0xd2, 0xe2, 0x12, 0x0c, 0x2e, 0x4d, 0xca, 0xcd, 0x2c, 0x71, 0xce, 0x4f, 0x49,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe5, 0x62, 0x2b, 0x2a, 0xcd, 0x8b, 0xcf,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0x2d, 0x2a, 0xcd, 0xf3, 0x4c, 0x51, 0x32,
	0xe0, 0x12, 0x42, 0x56, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc5, 0xc5, 0x51, 0x94,
	0x9a, 0x9c, 0x9a, 0x59, 0x96, 0x0a, 0x51, 0xce, 0x11, 0x04, 0xe7, 0x1b, 0xf9, 0x70, 0xf1, 0x3a,
	0x41, 0xec, 0xf4, 0x48, 0xcd, 0x29, 0x48, 0x2d, 0x12, 0xb2, 0xe6, 0x62, 0x83, 0x18, 0x21, 0x24,
	0xa6, 0x57, 0x54, 0x90, 0xac, 0x87, 0x61, 0xb7, 0x94, 0x38, 0x86, 0x38, 0xc4, 0x1e, 0x25, 0x86,
	0x24, 0x36, 0xb0, 0xbb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x15, 0xee, 0xda, 0xcc,
	0x00, 0x00, 0x00,
}
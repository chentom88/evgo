// Code generated by protoc-gen-go.
// source: proto/api.proto
// DO NOT EDIT!

/*
Package api_proto is a generated protocol buffer package.

It is generated from these files:
	proto/api.proto

It has these top-level messages:
	ApiResponse
	Destination
*/
package api_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

//import _ "google/api"

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

type ApiResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ApiResponse) Reset()                    { *m = ApiResponse{} }
func (m *ApiResponse) String() string            { return proto.CompactTextString(m) }
func (*ApiResponse) ProtoMessage()               {}
func (*ApiResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Destination struct {
	Name        string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Lat         float64 `protobuf:"fixed64,3,opt,name=lat" json:"lat,omitempty"`
	Long        float64 `protobuf:"fixed64,4,opt,name=long" json:"long,omitempty"`
	Phone       string  `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Url         string  `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
}

func (m *Destination) Reset()                    { *m = Destination{} }
func (m *Destination) String() string            { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()               {}
func (*Destination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ApiResponse)(nil), "api_proto.ApiResponse")
	proto.RegisterType((*Destination)(nil), "api_proto.Destination")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DestinationAPI service

type DestinationAPIClient interface {
	AddDestination(ctx context.Context, in *Destination, opts ...grpc.CallOption) (*ApiResponse, error)
}

type destinationAPIClient struct {
	cc *grpc.ClientConn
}

func NewDestinationAPIClient(cc *grpc.ClientConn) DestinationAPIClient {
	return &destinationAPIClient{cc}
}

func (c *destinationAPIClient) AddDestination(ctx context.Context, in *Destination, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := grpc.Invoke(ctx, "/api_proto.DestinationAPI/AddDestination", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DestinationAPI service

type DestinationAPIServer interface {
	AddDestination(context.Context, *Destination) (*ApiResponse, error)
}

func RegisterDestinationAPIServer(s *grpc.Server, srv DestinationAPIServer) {
	s.RegisterService(&_DestinationAPI_serviceDesc, srv)
}

func _DestinationAPI_AddDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Destination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationAPIServer).AddDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_proto.DestinationAPI/AddDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationAPIServer).AddDestination(ctx, req.(*Destination))
	}
	return interceptor(ctx, in, info, handler)
}

var _DestinationAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api_proto.DestinationAPI",
	HandlerType: (*DestinationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDestination",
			Handler:    _DestinationAPI_AddDestination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

func init() { proto.RegisterFile("proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xe9, 0xcc, 0x74, 0x3e, 0x7a, 0x0a, 0xf3, 0x0d, 0x41, 0x25, 0x8c, 0x2e, 0x86, 0xae,
	0xc4, 0x45, 0x8b, 0xba, 0xd3, 0x55, 0xc1, 0x8d, 0x3b, 0xc9, 0x5e, 0x24, 0x4e, 0x42, 0x0d, 0x74,
	0x72, 0x42, 0x13, 0xbd, 0x00, 0xef, 0x40, 0xbc, 0x34, 0x6f, 0xc1, 0x0b, 0xb1, 0x39, 0xf5, 0x27,
	0xbb, 0x27, 0x4f, 0xce, 0xfb, 0x72, 0x38, 0xf0, 0xdf, 0x0d, 0x18, 0xb0, 0x91, 0xce, 0xd4, 0x44,
	0xac, 0x18, 0xf1, 0x81, 0x70, 0x73, 0xd2, 0x21, 0x76, 0xbd, 0x8e, 0x9f, 0x8d, 0xb4, 0x16, 0x83,
	0x0c, 0x06, 0xad, 0x9f, 0x06, 0xab, 0x6b, 0x28, 0x5b, 0x67, 0x84, 0xf6, 0x6e, 0x74, 0x9a, 0x31,
	0x58, 0xec, 0x50, 0x69, 0x9e, 0x6d, 0xb3, 0xd3, 0x5c, 0x10, 0x33, 0x0e, 0xff, 0xf6, 0xda, 0x7b,
	0xd9, 0x69, 0x3e, 0x1b, 0x75, 0x21, 0x7e, 0x9e, 0xd5, 0x5b, 0x06, 0xe5, 0x8d, 0xf6, 0xc1, 0x58,
	0xea, 0x8c, 0x69, 0x2b, 0xf7, 0x53, 0xba, 0x10, 0xc4, 0x6c, 0x0b, 0xa5, 0xd2, 0x7e, 0x37, 0x18,
	0x17, 0x47, 0xbe, 0x1b, 0x52, 0xc5, 0xd6, 0x30, 0xef, 0x65, 0xe0, 0xf3, 0xf1, 0x27, 0x13, 0x11,
	0x63, 0x4f, 0x8f, 0xb6, 0xe3, 0x0b, 0x52, 0xc4, 0xec, 0x00, 0x72, 0xf7, 0x84, 0x56, 0xf3, 0x9c,
	0x1a, 0xa6, 0x47, 0xcc, 0x3e, 0x0f, 0x3d, 0x5f, 0x92, 0x8b, 0x78, 0x81, 0xb0, 0x4a, 0x56, 0x6a,
	0xef, 0x6e, 0xd9, 0x3d, 0xac, 0x5a, 0xa5, 0xd2, 0x3d, 0x8f, 0xea, 0xdf, 0xf3, 0xd4, 0x89, 0xdf,
	0xa4, 0x3e, 0xb9, 0x4a, 0x75, 0xfc, 0xfa, 0xf1, 0xf9, 0x3e, 0x3b, 0xac, 0xd6, 0xcd, 0xcb, 0x79,
	0xa3, 0xfe, 0x02, 0xfe, 0x2a, 0x3b, 0x7b, 0x5c, 0xd2, 0xfc, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xa2, 0x36, 0x06, 0x84, 0x01, 0x00, 0x00,
}

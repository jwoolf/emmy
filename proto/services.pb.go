// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PseudonymSystemCA service

type PseudonymSystemCAClient interface {
	GenerateCertificate(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificateClient, error)
	GenerateCertificate_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificate_ECClient, error)
}

type pseudonymSystemCAClient struct {
	cc *grpc.ClientConn
}

func NewPseudonymSystemCAClient(cc *grpc.ClientConn) PseudonymSystemCAClient {
	return &pseudonymSystemCAClient{cc}
}

func (c *pseudonymSystemCAClient) GenerateCertificate(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystemCA_serviceDesc.Streams[0], c.cc, "/proto.PseudonymSystemCA/GenerateCertificate", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemCAGenerateCertificateClient{stream}
	return x, nil
}

type PseudonymSystemCA_GenerateCertificateClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemCAGenerateCertificateClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemCAGenerateCertificateClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificateClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemCAClient) GenerateCertificate_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificate_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystemCA_serviceDesc.Streams[1], c.cc, "/proto.PseudonymSystemCA/GenerateCertificate_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemCAGenerateCertificate_ECClient{stream}
	return x, nil
}

type PseudonymSystemCA_GenerateCertificate_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemCAGenerateCertificate_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemCAGenerateCertificate_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificate_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PseudonymSystemCA service

type PseudonymSystemCAServer interface {
	GenerateCertificate(PseudonymSystemCA_GenerateCertificateServer) error
	GenerateCertificate_EC(PseudonymSystemCA_GenerateCertificate_ECServer) error
}

func RegisterPseudonymSystemCAServer(s *grpc.Server, srv PseudonymSystemCAServer) {
	s.RegisterService(&_PseudonymSystemCA_serviceDesc, srv)
}

func _PseudonymSystemCA_GenerateCertificate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemCAServer).GenerateCertificate(&pseudonymSystemCAGenerateCertificateServer{stream})
}

type PseudonymSystemCA_GenerateCertificateServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemCAGenerateCertificateServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemCAGenerateCertificateServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificateServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystemCA_GenerateCertificate_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemCAServer).GenerateCertificate_EC(&pseudonymSystemCAGenerateCertificate_ECServer{stream})
}

type PseudonymSystemCA_GenerateCertificate_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemCAGenerateCertificate_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemCAGenerateCertificate_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificate_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PseudonymSystemCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PseudonymSystemCA",
	HandlerType: (*PseudonymSystemCAServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateCertificate",
			Handler:       _PseudonymSystemCA_GenerateCertificate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GenerateCertificate_EC",
			Handler:       _PseudonymSystemCA_GenerateCertificate_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// Client API for PseudonymSystem service

type PseudonymSystemClient interface {
	GenerateNym(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNymClient, error)
	GenerateNym_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNym_ECClient, error)
	ObtainCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredentialClient, error)
	ObtainCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredential_ECClient, error)
	TransferCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredentialClient, error)
	TransferCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredential_ECClient, error)
}

type pseudonymSystemClient struct {
	cc *grpc.ClientConn
}

func NewPseudonymSystemClient(cc *grpc.ClientConn) PseudonymSystemClient {
	return &pseudonymSystemClient{cc}
}

func (c *pseudonymSystemClient) GenerateNym(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNymClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[0], c.cc, "/proto.PseudonymSystem/GenerateNym", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemGenerateNymClient{stream}
	return x, nil
}

type PseudonymSystem_GenerateNymClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemGenerateNymClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemGenerateNymClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNymClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) GenerateNym_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNym_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[1], c.cc, "/proto.PseudonymSystem/GenerateNym_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemGenerateNym_ECClient{stream}
	return x, nil
}

type PseudonymSystem_GenerateNym_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemGenerateNym_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemGenerateNym_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNym_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) ObtainCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredentialClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[2], c.cc, "/proto.PseudonymSystem/ObtainCredential", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemObtainCredentialClient{stream}
	return x, nil
}

type PseudonymSystem_ObtainCredentialClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemObtainCredentialClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemObtainCredentialClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredentialClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) ObtainCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredential_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[3], c.cc, "/proto.PseudonymSystem/ObtainCredential_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemObtainCredential_ECClient{stream}
	return x, nil
}

type PseudonymSystem_ObtainCredential_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemObtainCredential_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemObtainCredential_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredential_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) TransferCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredentialClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[4], c.cc, "/proto.PseudonymSystem/TransferCredential", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemTransferCredentialClient{stream}
	return x, nil
}

type PseudonymSystem_TransferCredentialClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemTransferCredentialClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemTransferCredentialClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredentialClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) TransferCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredential_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[5], c.cc, "/proto.PseudonymSystem/TransferCredential_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemTransferCredential_ECClient{stream}
	return x, nil
}

type PseudonymSystem_TransferCredential_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemTransferCredential_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemTransferCredential_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredential_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PseudonymSystem service

type PseudonymSystemServer interface {
	GenerateNym(PseudonymSystem_GenerateNymServer) error
	GenerateNym_EC(PseudonymSystem_GenerateNym_ECServer) error
	ObtainCredential(PseudonymSystem_ObtainCredentialServer) error
	ObtainCredential_EC(PseudonymSystem_ObtainCredential_ECServer) error
	TransferCredential(PseudonymSystem_TransferCredentialServer) error
	TransferCredential_EC(PseudonymSystem_TransferCredential_ECServer) error
}

func RegisterPseudonymSystemServer(s *grpc.Server, srv PseudonymSystemServer) {
	s.RegisterService(&_PseudonymSystem_serviceDesc, srv)
}

func _PseudonymSystem_GenerateNym_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).GenerateNym(&pseudonymSystemGenerateNymServer{stream})
}

type PseudonymSystem_GenerateNymServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemGenerateNymServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemGenerateNymServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNymServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_GenerateNym_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).GenerateNym_EC(&pseudonymSystemGenerateNym_ECServer{stream})
}

type PseudonymSystem_GenerateNym_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemGenerateNym_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemGenerateNym_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNym_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_ObtainCredential_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).ObtainCredential(&pseudonymSystemObtainCredentialServer{stream})
}

type PseudonymSystem_ObtainCredentialServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemObtainCredentialServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemObtainCredentialServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredentialServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_ObtainCredential_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).ObtainCredential_EC(&pseudonymSystemObtainCredential_ECServer{stream})
}

type PseudonymSystem_ObtainCredential_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemObtainCredential_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemObtainCredential_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredential_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_TransferCredential_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).TransferCredential(&pseudonymSystemTransferCredentialServer{stream})
}

type PseudonymSystem_TransferCredentialServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemTransferCredentialServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemTransferCredentialServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredentialServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_TransferCredential_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).TransferCredential_EC(&pseudonymSystemTransferCredential_ECServer{stream})
}

type PseudonymSystem_TransferCredential_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemTransferCredential_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemTransferCredential_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredential_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PseudonymSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PseudonymSystem",
	HandlerType: (*PseudonymSystemServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateNym",
			Handler:       _PseudonymSystem_GenerateNym_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GenerateNym_EC",
			Handler:       _PseudonymSystem_GenerateNym_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ObtainCredential",
			Handler:       _PseudonymSystem_ObtainCredential_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ObtainCredential_EC",
			Handler:       _PseudonymSystem_ObtainCredential_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TransferCredential",
			Handler:       _PseudonymSystem_TransferCredential_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TransferCredential_EC",
			Handler:       _PseudonymSystem_TransferCredential_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// Client API for Info service

type InfoClient interface {
	GetServiceInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServiceInfo, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) GetServiceInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServiceInfo, error) {
	out := new(ServiceInfo)
	err := grpc.Invoke(ctx, "/proto.Info/GetServiceInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Info service

type InfoServer interface {
	GetServiceInfo(context.Context, *google_protobuf.Empty) (*ServiceInfo, error)
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Info/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).GetServiceInfo(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInfo",
			Handler:    _Info_GetServiceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

func init() { proto1.RegisterFile("services.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x7c,
	0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x61, 0x29, 0xe9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54,
	0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x22, 0x69, 0x34, 0x81,
	0x91, 0x4b, 0x30, 0xa0, 0x38, 0xb5, 0x34, 0x25, 0x3f, 0xaf, 0x32, 0x37, 0xb8, 0xb2, 0xb8, 0x24,
	0x35, 0xd7, 0xd9, 0x51, 0xc8, 0x9a, 0x4b, 0xd8, 0x3d, 0x35, 0x2f, 0xb5, 0x28, 0xb1, 0x24, 0xd5,
	0x39, 0xb5, 0xa8, 0x24, 0x33, 0x2d, 0x33, 0x39, 0xb1, 0x24, 0x55, 0x88, 0x0f, 0xa2, 0x49, 0xcf,
	0x17, 0x62, 0x81, 0x14, 0x1a, 0x5f, 0x89, 0x41, 0x83, 0xd1, 0x80, 0x51, 0xc8, 0x8e, 0x4b, 0x0c,
	0x8b, 0xe6, 0x78, 0x57, 0x67, 0xe2, 0xf4, 0x1b, 0x7d, 0x60, 0xe2, 0xe2, 0x47, 0x73, 0x92, 0x90,
	0x31, 0x17, 0x37, 0xcc, 0x4c, 0xbf, 0xca, 0x5c, 0x22, 0x1d, 0x62, 0xc6, 0xc5, 0x87, 0xa4, 0x89,
	0x68, 0x07, 0x08, 0x59, 0x70, 0x09, 0xf8, 0x27, 0x95, 0x24, 0x66, 0xe6, 0x39, 0x17, 0xa5, 0xa6,
	0xa4, 0xe6, 0x95, 0x64, 0x26, 0xe6, 0x10, 0xa9, 0xd3, 0x9a, 0x4b, 0x18, 0x5d, 0x27, 0xf1, 0xd6,
	0x5a, 0x71, 0x09, 0x85, 0x14, 0x25, 0xe6, 0x15, 0xa7, 0xa5, 0x16, 0x91, 0x6c, 0xb1, 0x2d, 0x97,
	0x28, 0xa6, 0x5e, 0xe2, 0x83, 0xdc, 0x8d, 0x8b, 0xc5, 0x33, 0x2f, 0x2d, 0x5f, 0xc8, 0x0e, 0x14,
	0x62, 0x25, 0xc1, 0x90, 0x64, 0x05, 0x16, 0x11, 0xd3, 0x83, 0xa4, 0x1e, 0x3d, 0x58, 0xea, 0xd1,
	0x73, 0x05, 0xa5, 0x1e, 0x29, 0x21, 0xa8, 0x39, 0x48, 0x6a, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x82,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x49, 0xd0, 0xbc, 0x9a, 0x02, 0x00, 0x00,
}

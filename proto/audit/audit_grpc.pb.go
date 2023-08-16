// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: audit/audit.proto

package audit

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuditClient is the client API for Audit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditClient interface {
	RawString(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error)
}

type auditClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditClient(cc grpc.ClientConnInterface) AuditClient {
	return &auditClient{cc}
}

func (c *auditClient) RawString(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/proto.Audit/RawString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServer is the server API for Audit service.
// All implementations must embed UnimplementedAuditServer
// for forward compatibility
type AuditServer interface {
	RawString(context.Context, *RawRequest) (*RawResponse, error)
	mustEmbedUnimplementedAuditServer()
}

// UnimplementedAuditServer must be embedded to have forward compatible implementations.
type UnimplementedAuditServer struct {
}

func (UnimplementedAuditServer) RawString(context.Context, *RawRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawString not implemented")
}
func (UnimplementedAuditServer) mustEmbedUnimplementedAuditServer() {}

// UnsafeAuditServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditServer will
// result in compilation errors.
type UnsafeAuditServer interface {
	mustEmbedUnimplementedAuditServer()
}

func RegisterAuditServer(s grpc.ServiceRegistrar, srv AuditServer) {
	s.RegisterService(&Audit_ServiceDesc, srv)
}

func _Audit_RawString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).RawString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Audit/RawString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).RawString(ctx, req.(*RawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Audit_ServiceDesc is the grpc.ServiceDesc for Audit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Audit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Audit",
	HandlerType: (*AuditServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RawString",
			Handler:    _Audit_RawString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit/audit.proto",
}
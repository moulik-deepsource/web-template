// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package phone

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

// PhoneServiceClient is the client API for PhoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhoneServiceClient interface {
	GetOneByID(ctx context.Context, in *GetOneByIDRequest, opts ...grpc.CallOption) (*GetOneByIDResponse, error)
	ListByCursor(ctx context.Context, in *ListByCursorRequest, opts ...grpc.CallOption) (*ListByCursorResponse, error)
	ListByPage(ctx context.Context, in *ListByPageRequest, opts ...grpc.CallOption) (*ListByPageResponse, error)
}

type phoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhoneServiceClient(cc grpc.ClientConnInterface) PhoneServiceClient {
	return &phoneServiceClient{cc}
}

func (c *phoneServiceClient) GetOneByID(ctx context.Context, in *GetOneByIDRequest, opts ...grpc.CallOption) (*GetOneByIDResponse, error) {
	out := new(GetOneByIDResponse)
	err := c.cc.Invoke(ctx, "/phone.PhoneService/GetOneByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneServiceClient) ListByCursor(ctx context.Context, in *ListByCursorRequest, opts ...grpc.CallOption) (*ListByCursorResponse, error) {
	out := new(ListByCursorResponse)
	err := c.cc.Invoke(ctx, "/phone.PhoneService/ListByCursor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneServiceClient) ListByPage(ctx context.Context, in *ListByPageRequest, opts ...grpc.CallOption) (*ListByPageResponse, error) {
	out := new(ListByPageResponse)
	err := c.cc.Invoke(ctx, "/phone.PhoneService/ListByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneServiceServer is the server API for PhoneService service.
// All implementations must embed UnimplementedPhoneServiceServer
// for forward compatibility
type PhoneServiceServer interface {
	GetOneByID(context.Context, *GetOneByIDRequest) (*GetOneByIDResponse, error)
	ListByCursor(context.Context, *ListByCursorRequest) (*ListByCursorResponse, error)
	ListByPage(context.Context, *ListByPageRequest) (*ListByPageResponse, error)
	mustEmbedUnimplementedPhoneServiceServer()
}

// UnimplementedPhoneServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPhoneServiceServer struct {
}

func (UnimplementedPhoneServiceServer) GetOneByID(context.Context, *GetOneByIDRequest) (*GetOneByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneByID not implemented")
}
func (UnimplementedPhoneServiceServer) ListByCursor(context.Context, *ListByCursorRequest) (*ListByCursorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByCursor not implemented")
}
func (UnimplementedPhoneServiceServer) ListByPage(context.Context, *ListByPageRequest) (*ListByPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByPage not implemented")
}
func (UnimplementedPhoneServiceServer) mustEmbedUnimplementedPhoneServiceServer() {}

// UnsafePhoneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhoneServiceServer will
// result in compilation errors.
type UnsafePhoneServiceServer interface {
	mustEmbedUnimplementedPhoneServiceServer()
}

func RegisterPhoneServiceServer(s grpc.ServiceRegistrar, srv PhoneServiceServer) {
	s.RegisterService(&PhoneService_ServiceDesc, srv)
}

func _PhoneService_GetOneByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneServiceServer).GetOneByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phone.PhoneService/GetOneByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneServiceServer).GetOneByID(ctx, req.(*GetOneByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneService_ListByCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByCursorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneServiceServer).ListByCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phone.PhoneService/ListByCursor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneServiceServer).ListByCursor(ctx, req.(*ListByCursorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneService_ListByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneServiceServer).ListByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phone.PhoneService/ListByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneServiceServer).ListByPage(ctx, req.(*ListByPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhoneService_ServiceDesc is the grpc.ServiceDesc for PhoneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhoneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phone.PhoneService",
	HandlerType: (*PhoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneByID",
			Handler:    _PhoneService_GetOneByID_Handler,
		},
		{
			MethodName: "ListByCursor",
			Handler:    _PhoneService_ListByCursor_Handler,
		},
		{
			MethodName: "ListByPage",
			Handler:    _PhoneService_ListByPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/phone/phone_service.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: github.com/trajectoryjp/trjx-vehicle-api/proto/mav_controller_outside/trjxgcs.proto

package trjxmavlink

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

// TrjxGcsServiceClient is the client API for TrjxGcsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrjxGcsServiceClient interface {
	// バージョン情報取得。認証不要。
	GetSeriviceAttribute(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SeriviceAttribute, error)
	// GCSからTRJXへの接続要求。機体ごとに行う。
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*GCToken, error)
	// 未実装。Connect()で得たaircraft_idとtokenでアクセスする。
	Communicate(ctx context.Context, opts ...grpc.CallOption) (TrjxGcsService_CommunicateClient, error)
	// 未実装。Connect()で得たaircraft_idとtokenでアクセスする。
	Request(ctx context.Context, in *TrjxVehicleRequest, opts ...grpc.CallOption) (TrjxGcsService_RequestClient, error)
	// 応答はkeepalive。Connect()で得たaircraft_idとtokenでアクセスする。
	// 1機体につき1ストリーム（機体ごとに本APIを呼ぶこと）
	Telemetry(ctx context.Context, opts ...grpc.CallOption) (TrjxGcsService_TelemetryClient, error)
	// 未実装。Connect()で得たaircraft_idとtokenでアクセスする。
	ReceiveCommand(ctx context.Context, opts ...grpc.CallOption) (TrjxGcsService_ReceiveCommandClient, error)
}

type trjxGcsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrjxGcsServiceClient(cc grpc.ClientConnInterface) TrjxGcsServiceClient {
	return &trjxGcsServiceClient{cc}
}

func (c *trjxGcsServiceClient) GetSeriviceAttribute(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SeriviceAttribute, error) {
	out := new(SeriviceAttribute)
	err := c.cc.Invoke(ctx, "/trjxmavlink.TrjxGcsService/GetSeriviceAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trjxGcsServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*GCToken, error) {
	out := new(GCToken)
	err := c.cc.Invoke(ctx, "/trjxmavlink.TrjxGcsService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trjxGcsServiceClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (TrjxGcsService_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrjxGcsService_ServiceDesc.Streams[0], "/trjxmavlink.TrjxGcsService/Communicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &trjxGcsServiceCommunicateClient{stream}
	return x, nil
}

type TrjxGcsService_CommunicateClient interface {
	Send(*TrjxVehicleMessage) error
	Recv() (*TrjxVehicleMessage, error)
	grpc.ClientStream
}

type trjxGcsServiceCommunicateClient struct {
	grpc.ClientStream
}

func (x *trjxGcsServiceCommunicateClient) Send(m *TrjxVehicleMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *trjxGcsServiceCommunicateClient) Recv() (*TrjxVehicleMessage, error) {
	m := new(TrjxVehicleMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trjxGcsServiceClient) Request(ctx context.Context, in *TrjxVehicleRequest, opts ...grpc.CallOption) (TrjxGcsService_RequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrjxGcsService_ServiceDesc.Streams[1], "/trjxmavlink.TrjxGcsService/Request", opts...)
	if err != nil {
		return nil, err
	}
	x := &trjxGcsServiceRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrjxGcsService_RequestClient interface {
	Recv() (*TrjxVehicleResponse, error)
	grpc.ClientStream
}

type trjxGcsServiceRequestClient struct {
	grpc.ClientStream
}

func (x *trjxGcsServiceRequestClient) Recv() (*TrjxVehicleResponse, error) {
	m := new(TrjxVehicleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trjxGcsServiceClient) Telemetry(ctx context.Context, opts ...grpc.CallOption) (TrjxGcsService_TelemetryClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrjxGcsService_ServiceDesc.Streams[2], "/trjxmavlink.TrjxGcsService/Telemetry", opts...)
	if err != nil {
		return nil, err
	}
	x := &trjxGcsServiceTelemetryClient{stream}
	return x, nil
}

type TrjxGcsService_TelemetryClient interface {
	Send(*TrjxVehicleTelemetry) error
	Recv() (*Empty, error)
	grpc.ClientStream
}

type trjxGcsServiceTelemetryClient struct {
	grpc.ClientStream
}

func (x *trjxGcsServiceTelemetryClient) Send(m *TrjxVehicleTelemetry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *trjxGcsServiceTelemetryClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trjxGcsServiceClient) ReceiveCommand(ctx context.Context, opts ...grpc.CallOption) (TrjxGcsService_ReceiveCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrjxGcsService_ServiceDesc.Streams[3], "/trjxmavlink.TrjxGcsService/ReceiveCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &trjxGcsServiceReceiveCommandClient{stream}
	return x, nil
}

type TrjxGcsService_ReceiveCommandClient interface {
	Send(*TrjxVehicleReceiveCommandHandling) error
	Recv() (*TrjxVehicleCommand, error)
	grpc.ClientStream
}

type trjxGcsServiceReceiveCommandClient struct {
	grpc.ClientStream
}

func (x *trjxGcsServiceReceiveCommandClient) Send(m *TrjxVehicleReceiveCommandHandling) error {
	return x.ClientStream.SendMsg(m)
}

func (x *trjxGcsServiceReceiveCommandClient) Recv() (*TrjxVehicleCommand, error) {
	m := new(TrjxVehicleCommand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrjxGcsServiceServer is the server API for TrjxGcsService service.
// All implementations must embed UnimplementedTrjxGcsServiceServer
// for forward compatibility
type TrjxGcsServiceServer interface {
	// バージョン情報取得。認証不要。
	GetSeriviceAttribute(context.Context, *Empty) (*SeriviceAttribute, error)
	// GCSからTRJXへの接続要求。機体ごとに行う。
	Connect(context.Context, *ConnectRequest) (*GCToken, error)
	// 未実装。Connect()で得たaircraft_idとtokenでアクセスする。
	Communicate(TrjxGcsService_CommunicateServer) error
	// 未実装。Connect()で得たaircraft_idとtokenでアクセスする。
	Request(*TrjxVehicleRequest, TrjxGcsService_RequestServer) error
	// 応答はkeepalive。Connect()で得たaircraft_idとtokenでアクセスする。
	// 1機体につき1ストリーム（機体ごとに本APIを呼ぶこと）
	Telemetry(TrjxGcsService_TelemetryServer) error
	// 未実装。Connect()で得たaircraft_idとtokenでアクセスする。
	ReceiveCommand(TrjxGcsService_ReceiveCommandServer) error
	mustEmbedUnimplementedTrjxGcsServiceServer()
}

// UnimplementedTrjxGcsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrjxGcsServiceServer struct {
}

func (UnimplementedTrjxGcsServiceServer) GetSeriviceAttribute(context.Context, *Empty) (*SeriviceAttribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeriviceAttribute not implemented")
}
func (UnimplementedTrjxGcsServiceServer) Connect(context.Context, *ConnectRequest) (*GCToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedTrjxGcsServiceServer) Communicate(TrjxGcsService_CommunicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Communicate not implemented")
}
func (UnimplementedTrjxGcsServiceServer) Request(*TrjxVehicleRequest, TrjxGcsService_RequestServer) error {
	return status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedTrjxGcsServiceServer) Telemetry(TrjxGcsService_TelemetryServer) error {
	return status.Errorf(codes.Unimplemented, "method Telemetry not implemented")
}
func (UnimplementedTrjxGcsServiceServer) ReceiveCommand(TrjxGcsService_ReceiveCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveCommand not implemented")
}
func (UnimplementedTrjxGcsServiceServer) mustEmbedUnimplementedTrjxGcsServiceServer() {}

// UnsafeTrjxGcsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrjxGcsServiceServer will
// result in compilation errors.
type UnsafeTrjxGcsServiceServer interface {
	mustEmbedUnimplementedTrjxGcsServiceServer()
}

func RegisterTrjxGcsServiceServer(s grpc.ServiceRegistrar, srv TrjxGcsServiceServer) {
	s.RegisterService(&TrjxGcsService_ServiceDesc, srv)
}

func _TrjxGcsService_GetSeriviceAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrjxGcsServiceServer).GetSeriviceAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trjxmavlink.TrjxGcsService/GetSeriviceAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrjxGcsServiceServer).GetSeriviceAttribute(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrjxGcsService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrjxGcsServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trjxmavlink.TrjxGcsService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrjxGcsServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrjxGcsService_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TrjxGcsServiceServer).Communicate(&trjxGcsServiceCommunicateServer{stream})
}

type TrjxGcsService_CommunicateServer interface {
	Send(*TrjxVehicleMessage) error
	Recv() (*TrjxVehicleMessage, error)
	grpc.ServerStream
}

type trjxGcsServiceCommunicateServer struct {
	grpc.ServerStream
}

func (x *trjxGcsServiceCommunicateServer) Send(m *TrjxVehicleMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *trjxGcsServiceCommunicateServer) Recv() (*TrjxVehicleMessage, error) {
	m := new(TrjxVehicleMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TrjxGcsService_Request_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TrjxVehicleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrjxGcsServiceServer).Request(m, &trjxGcsServiceRequestServer{stream})
}

type TrjxGcsService_RequestServer interface {
	Send(*TrjxVehicleResponse) error
	grpc.ServerStream
}

type trjxGcsServiceRequestServer struct {
	grpc.ServerStream
}

func (x *trjxGcsServiceRequestServer) Send(m *TrjxVehicleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TrjxGcsService_Telemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TrjxGcsServiceServer).Telemetry(&trjxGcsServiceTelemetryServer{stream})
}

type TrjxGcsService_TelemetryServer interface {
	Send(*Empty) error
	Recv() (*TrjxVehicleTelemetry, error)
	grpc.ServerStream
}

type trjxGcsServiceTelemetryServer struct {
	grpc.ServerStream
}

func (x *trjxGcsServiceTelemetryServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *trjxGcsServiceTelemetryServer) Recv() (*TrjxVehicleTelemetry, error) {
	m := new(TrjxVehicleTelemetry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TrjxGcsService_ReceiveCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TrjxGcsServiceServer).ReceiveCommand(&trjxGcsServiceReceiveCommandServer{stream})
}

type TrjxGcsService_ReceiveCommandServer interface {
	Send(*TrjxVehicleCommand) error
	Recv() (*TrjxVehicleReceiveCommandHandling, error)
	grpc.ServerStream
}

type trjxGcsServiceReceiveCommandServer struct {
	grpc.ServerStream
}

func (x *trjxGcsServiceReceiveCommandServer) Send(m *TrjxVehicleCommand) error {
	return x.ServerStream.SendMsg(m)
}

func (x *trjxGcsServiceReceiveCommandServer) Recv() (*TrjxVehicleReceiveCommandHandling, error) {
	m := new(TrjxVehicleReceiveCommandHandling)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrjxGcsService_ServiceDesc is the grpc.ServiceDesc for TrjxGcsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrjxGcsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trjxmavlink.TrjxGcsService",
	HandlerType: (*TrjxGcsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSeriviceAttribute",
			Handler:    _TrjxGcsService_GetSeriviceAttribute_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _TrjxGcsService_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Communicate",
			Handler:       _TrjxGcsService_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Request",
			Handler:       _TrjxGcsService_Request_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Telemetry",
			Handler:       _TrjxGcsService_Telemetry_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveCommand",
			Handler:       _TrjxGcsService_ReceiveCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/trajectoryjp/trjx-vehicle-api/proto/mav_controller_outside/trjxgcs.proto",
}

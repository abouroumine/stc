// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package db_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthenticationInfoClient is the client API for AuthenticationInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationInfoClient interface {
	CheckInfoDB(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*UserAuth, error)
	SignUp(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type authenticationInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationInfoClient(cc grpc.ClientConnInterface) AuthenticationInfoClient {
	return &authenticationInfoClient{cc}
}

func (c *authenticationInfoClient) CheckInfoDB(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*UserAuth, error) {
	out := new(UserAuth)
	err := c.cc.Invoke(ctx, "/space_traffic_control.AuthenticationInfo/checkInfoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationInfoClient) SignUp(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/space_traffic_control.AuthenticationInfo/signUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationInfoServer is the server API for AuthenticationInfo service.
// All implementations must embed UnimplementedAuthenticationInfoServer
// for forward compatibility
type AuthenticationInfoServer interface {
	CheckInfoDB(context.Context, *UserAuth) (*UserAuth, error)
	SignUp(context.Context, *UserAuth) (*wrapperspb.BoolValue, error)
	mustEmbedUnimplementedAuthenticationInfoServer()
}

// UnimplementedAuthenticationInfoServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationInfoServer struct {
}

func (UnimplementedAuthenticationInfoServer) CheckInfoDB(context.Context, *UserAuth) (*UserAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInfoDB not implemented")
}
func (UnimplementedAuthenticationInfoServer) SignUp(context.Context, *UserAuth) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthenticationInfoServer) mustEmbedUnimplementedAuthenticationInfoServer() {}

// UnsafeAuthenticationInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationInfoServer will
// result in compilation errors.
type UnsafeAuthenticationInfoServer interface {
	mustEmbedUnimplementedAuthenticationInfoServer()
}

func RegisterAuthenticationInfoServer(s grpc.ServiceRegistrar, srv AuthenticationInfoServer) {
	s.RegisterService(&AuthenticationInfo_ServiceDesc, srv)
}

func _AuthenticationInfo_CheckInfoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationInfoServer).CheckInfoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.AuthenticationInfo/checkInfoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationInfoServer).CheckInfoDB(ctx, req.(*UserAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationInfo_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationInfoServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.AuthenticationInfo/signUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationInfoServer).SignUp(ctx, req.(*UserAuth))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationInfo_ServiceDesc is the grpc.ServiceDesc for AuthenticationInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "space_traffic_control.AuthenticationInfo",
	HandlerType: (*AuthenticationInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkInfoDB",
			Handler:    _AuthenticationInfo_CheckInfoDB_Handler,
		},
		{
			MethodName: "signUp",
			Handler:    _AuthenticationInfo_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db_proto.proto",
}

// CCServiceClient is the client API for CCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CCServiceClient interface {
	StationRegister(ctx context.Context, in *Station, opts ...grpc.CallOption) (*Station, error)
	ShipCCInfo(ctx context.Context, in *wrapperspb.Int32Value, opts ...grpc.CallOption) (*Ship, error)
	AllStationsNoCondition(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Stations, error)
	AllStationsWithCondition(ctx context.Context, in *wrapperspb.FloatValue, opts ...grpc.CallOption) (*Stations, error)
	ShipRegister(ctx context.Context, in *wrapperspb.FloatValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AllShips(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Ships, error)
}

type cCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCCServiceClient(cc grpc.ClientConnInterface) CCServiceClient {
	return &cCServiceClient{cc}
}

func (c *cCServiceClient) StationRegister(ctx context.Context, in *Station, opts ...grpc.CallOption) (*Station, error) {
	out := new(Station)
	err := c.cc.Invoke(ctx, "/space_traffic_control.CCService/stationRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cCServiceClient) ShipCCInfo(ctx context.Context, in *wrapperspb.Int32Value, opts ...grpc.CallOption) (*Ship, error) {
	out := new(Ship)
	err := c.cc.Invoke(ctx, "/space_traffic_control.CCService/shipCCInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cCServiceClient) AllStationsNoCondition(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Stations, error) {
	out := new(Stations)
	err := c.cc.Invoke(ctx, "/space_traffic_control.CCService/allStationsNoCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cCServiceClient) AllStationsWithCondition(ctx context.Context, in *wrapperspb.FloatValue, opts ...grpc.CallOption) (*Stations, error) {
	out := new(Stations)
	err := c.cc.Invoke(ctx, "/space_traffic_control.CCService/allStationsWithCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cCServiceClient) ShipRegister(ctx context.Context, in *wrapperspb.FloatValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/space_traffic_control.CCService/shipRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cCServiceClient) AllShips(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Ships, error) {
	out := new(Ships)
	err := c.cc.Invoke(ctx, "/space_traffic_control.CCService/allShips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CCServiceServer is the server API for CCService service.
// All implementations must embed UnimplementedCCServiceServer
// for forward compatibility
type CCServiceServer interface {
	StationRegister(context.Context, *Station) (*Station, error)
	ShipCCInfo(context.Context, *wrapperspb.Int32Value) (*Ship, error)
	AllStationsNoCondition(context.Context, *emptypb.Empty) (*Stations, error)
	AllStationsWithCondition(context.Context, *wrapperspb.FloatValue) (*Stations, error)
	ShipRegister(context.Context, *wrapperspb.FloatValue) (*emptypb.Empty, error)
	AllShips(context.Context, *emptypb.Empty) (*Ships, error)
	mustEmbedUnimplementedCCServiceServer()
}

// UnimplementedCCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCCServiceServer struct {
}

func (UnimplementedCCServiceServer) StationRegister(context.Context, *Station) (*Station, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StationRegister not implemented")
}
func (UnimplementedCCServiceServer) ShipCCInfo(context.Context, *wrapperspb.Int32Value) (*Ship, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipCCInfo not implemented")
}
func (UnimplementedCCServiceServer) AllStationsNoCondition(context.Context, *emptypb.Empty) (*Stations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllStationsNoCondition not implemented")
}
func (UnimplementedCCServiceServer) AllStationsWithCondition(context.Context, *wrapperspb.FloatValue) (*Stations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllStationsWithCondition not implemented")
}
func (UnimplementedCCServiceServer) ShipRegister(context.Context, *wrapperspb.FloatValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipRegister not implemented")
}
func (UnimplementedCCServiceServer) AllShips(context.Context, *emptypb.Empty) (*Ships, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllShips not implemented")
}
func (UnimplementedCCServiceServer) mustEmbedUnimplementedCCServiceServer() {}

// UnsafeCCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CCServiceServer will
// result in compilation errors.
type UnsafeCCServiceServer interface {
	mustEmbedUnimplementedCCServiceServer()
}

func RegisterCCServiceServer(s grpc.ServiceRegistrar, srv CCServiceServer) {
	s.RegisterService(&CCService_ServiceDesc, srv)
}

func _CCService_StationRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Station)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCServiceServer).StationRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.CCService/stationRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCServiceServer).StationRegister(ctx, req.(*Station))
	}
	return interceptor(ctx, in, info, handler)
}

func _CCService_ShipCCInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCServiceServer).ShipCCInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.CCService/shipCCInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCServiceServer).ShipCCInfo(ctx, req.(*wrapperspb.Int32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _CCService_AllStationsNoCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCServiceServer).AllStationsNoCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.CCService/allStationsNoCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCServiceServer).AllStationsNoCondition(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CCService_AllStationsWithCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.FloatValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCServiceServer).AllStationsWithCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.CCService/allStationsWithCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCServiceServer).AllStationsWithCondition(ctx, req.(*wrapperspb.FloatValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _CCService_ShipRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.FloatValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCServiceServer).ShipRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.CCService/shipRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCServiceServer).ShipRegister(ctx, req.(*wrapperspb.FloatValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _CCService_AllShips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCServiceServer).AllShips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.CCService/allShips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCServiceServer).AllShips(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CCService_ServiceDesc is the grpc.ServiceDesc for CCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "space_traffic_control.CCService",
	HandlerType: (*CCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "stationRegister",
			Handler:    _CCService_StationRegister_Handler,
		},
		{
			MethodName: "shipCCInfo",
			Handler:    _CCService_ShipCCInfo_Handler,
		},
		{
			MethodName: "allStationsNoCondition",
			Handler:    _CCService_AllStationsNoCondition_Handler,
		},
		{
			MethodName: "allStationsWithCondition",
			Handler:    _CCService_AllStationsWithCondition_Handler,
		},
		{
			MethodName: "shipRegister",
			Handler:    _CCService_ShipRegister_Handler,
		},
		{
			MethodName: "allShips",
			Handler:    _CCService_AllShips_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db_proto.proto",
}

// ShippingStationClient is the client API for ShippingStation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingStationClient interface {
	StationInfo(ctx context.Context, in *wrapperspb.Int32Value, opts ...grpc.CallOption) (*Station, error)
	ShipInfo(ctx context.Context, in *wrapperspb.Int32Value, opts ...grpc.CallOption) (*Ship, error)
	UpdateTheLandData(ctx context.Context, in *UpdateLandData, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type shippingStationClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingStationClient(cc grpc.ClientConnInterface) ShippingStationClient {
	return &shippingStationClient{cc}
}

func (c *shippingStationClient) StationInfo(ctx context.Context, in *wrapperspb.Int32Value, opts ...grpc.CallOption) (*Station, error) {
	out := new(Station)
	err := c.cc.Invoke(ctx, "/space_traffic_control.ShippingStation/stationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingStationClient) ShipInfo(ctx context.Context, in *wrapperspb.Int32Value, opts ...grpc.CallOption) (*Ship, error) {
	out := new(Ship)
	err := c.cc.Invoke(ctx, "/space_traffic_control.ShippingStation/shipInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingStationClient) UpdateTheLandData(ctx context.Context, in *UpdateLandData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/space_traffic_control.ShippingStation/updateTheLandData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingStationServer is the server API for ShippingStation service.
// All implementations must embed UnimplementedShippingStationServer
// for forward compatibility
type ShippingStationServer interface {
	StationInfo(context.Context, *wrapperspb.Int32Value) (*Station, error)
	ShipInfo(context.Context, *wrapperspb.Int32Value) (*Ship, error)
	UpdateTheLandData(context.Context, *UpdateLandData) (*emptypb.Empty, error)
	mustEmbedUnimplementedShippingStationServer()
}

// UnimplementedShippingStationServer must be embedded to have forward compatible implementations.
type UnimplementedShippingStationServer struct {
}

func (UnimplementedShippingStationServer) StationInfo(context.Context, *wrapperspb.Int32Value) (*Station, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StationInfo not implemented")
}
func (UnimplementedShippingStationServer) ShipInfo(context.Context, *wrapperspb.Int32Value) (*Ship, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipInfo not implemented")
}
func (UnimplementedShippingStationServer) UpdateTheLandData(context.Context, *UpdateLandData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTheLandData not implemented")
}
func (UnimplementedShippingStationServer) mustEmbedUnimplementedShippingStationServer() {}

// UnsafeShippingStationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingStationServer will
// result in compilation errors.
type UnsafeShippingStationServer interface {
	mustEmbedUnimplementedShippingStationServer()
}

func RegisterShippingStationServer(s grpc.ServiceRegistrar, srv ShippingStationServer) {
	s.RegisterService(&ShippingStation_ServiceDesc, srv)
}

func _ShippingStation_StationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingStationServer).StationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.ShippingStation/stationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingStationServer).StationInfo(ctx, req.(*wrapperspb.Int32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingStation_ShipInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingStationServer).ShipInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.ShippingStation/shipInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingStationServer).ShipInfo(ctx, req.(*wrapperspb.Int32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingStation_UpdateTheLandData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLandData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingStationServer).UpdateTheLandData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space_traffic_control.ShippingStation/updateTheLandData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingStationServer).UpdateTheLandData(ctx, req.(*UpdateLandData))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingStation_ServiceDesc is the grpc.ServiceDesc for ShippingStation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingStation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "space_traffic_control.ShippingStation",
	HandlerType: (*ShippingStationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "stationInfo",
			Handler:    _ShippingStation_StationInfo_Handler,
		},
		{
			MethodName: "shipInfo",
			Handler:    _ShippingStation_ShipInfo_Handler,
		},
		{
			MethodName: "updateTheLandData",
			Handler:    _ShippingStation_UpdateTheLandData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db_proto.proto",
}

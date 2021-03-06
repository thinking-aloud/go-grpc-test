// Code generated by protoc-gen-go.
// source: routes.proto
// DO NOT EDIT!

/*
Package routes is a generated protocol buffer package.

It is generated from these files:
	routes.proto

It has these top-level messages:
	Point
	Rectangle
	Feature
	RouteNote
	RouteSummary
*/
package routes

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

type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Rectangle struct {
	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`
	Hi *Point `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

type Feature struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

type RouteNote struct {
	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RouteNote) Reset()                    { *m = RouteNote{} }
func (m *RouteNote) String() string            { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()               {}
func (*RouteNote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RouteSummary struct {
	PointCount   int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`
	Distance     int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`
	ElapsedTime  int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
}

func (m *RouteSummary) Reset()                    { *m = RouteSummary{} }
func (m *RouteSummary) String() string            { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()               {}
func (*RouteSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "routes.Point")
	proto.RegisterType((*Rectangle)(nil), "routes.Rectangle")
	proto.RegisterType((*Feature)(nil), "routes.Feature")
	proto.RegisterType((*RouteNote)(nil), "routes.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "routes.RouteSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteGuide service

type RouteGuideClient interface {
	// simple
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// server-side streaming
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// client-side streaming
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// bidirectional streaming
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := grpc.Invoke(ctx, "/routes.RouteGuide/GetFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[0], c.cc, "/routes.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[1], c.cc, "/routes.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[2], c.cc, "/routes.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RouteGuide service

type RouteGuideServer interface {
	// simple
	GetFeature(context.Context, *Point) (*Feature, error)
	// server-side streaming
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	// client-side streaming
	RecordRoute(RouteGuide_RecordRouteServer) error
	// bidirectional streaming
	RouteChat(RouteGuide_RouteChatServer) error
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routes.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routes.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routes.proto",
}

func init() { proto.RegisterFile("routes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0xed, 0xf6, 0xf6, 0x2b, 0x93, 0x94, 0xcb, 0x1d, 0xee, 0x43, 0x08, 0x8a, 0x1a, 0x5f, 0xea,
	0x4b, 0x29, 0x15, 0xea, 0xb3, 0x14, 0xac, 0x82, 0x48, 0x89, 0xbe, 0x97, 0x35, 0x19, 0xdb, 0x85,
	0x24, 0x5b, 0x92, 0xcd, 0x83, 0xbf, 0xc3, 0xdf, 0xe8, 0xff, 0x90, 0xec, 0x26, 0xb1, 0xb5, 0xc5,
	0xb7, 0xcc, 0x39, 0x73, 0xe6, 0xec, 0x9c, 0x09, 0x38, 0x99, 0x2c, 0x14, 0xe5, 0xe3, 0x6d, 0x26,
	0x95, 0xc4, 0x9e, 0xa9, 0xfc, 0x5b, 0xe8, 0x2e, 0xa5, 0x48, 0x15, 0x7a, 0x30, 0x88, 0xb9, 0x12,
	0xaa, 0x88, 0xc8, 0x65, 0xe7, 0x6c, 0xd4, 0x0d, 0x9a, 0x1a, 0x4f, 0xc0, 0x8a, 0x65, 0xba, 0x36,
	0x64, 0x5b, 0x93, 0xdf, 0x80, 0xff, 0x00, 0x56, 0x40, 0xa1, 0xe2, 0xe9, 0x3a, 0x26, 0x3c, 0x85,
	0x76, 0x2c, 0xf5, 0x00, 0x7b, 0x3a, 0x1c, 0x57, 0x96, 0xda, 0x21, 0x68, 0xc7, 0xb2, 0xa4, 0x37,
	0x42, 0x8f, 0x38, 0xa4, 0x37, 0xc2, 0xbf, 0x87, 0xfe, 0x1d, 0x71, 0x55, 0x64, 0x84, 0x08, 0x9d,
	0x94, 0x27, 0xe6, 0x2d, 0x56, 0xa0, 0xbf, 0xf1, 0x0a, 0x06, 0xb1, 0x0c, 0xb9, 0x12, 0x32, 0x3d,
	0x3e, 0xa3, 0xa1, 0xfd, 0x25, 0x58, 0x41, 0xc9, 0x3c, 0x49, 0xb5, 0xaf, 0x63, 0xbf, 0xea, 0xd0,
	0x85, 0x7e, 0x42, 0x79, 0xce, 0xd7, 0x66, 0x51, 0x2b, 0xa8, 0x4b, 0xff, 0x83, 0x81, 0xa3, 0x47,
	0x3e, 0x17, 0x49, 0xc2, 0xb3, 0x77, 0x3c, 0x03, 0x7b, 0x5b, 0xaa, 0x57, 0xa1, 0x2c, 0x52, 0x55,
	0x85, 0x06, 0x1a, 0x9a, 0x97, 0x08, 0x5e, 0xc2, 0xf0, 0xcd, 0x6c, 0x53, 0xb5, 0x98, 0xe8, 0x9c,
	0x0a, 0x34, 0x4d, 0x1e, 0x0c, 0x22, 0x91, 0x2b, 0x9e, 0x86, 0xe4, 0xfe, 0x31, 0xb9, 0xd7, 0x35,
	0x5e, 0x80, 0x43, 0x31, 0xdf, 0xe6, 0x14, 0xad, 0x94, 0x48, 0xc8, 0xed, 0x68, 0xde, 0xae, 0xb0,
	0x17, 0x91, 0xd0, 0xf4, 0x93, 0x01, 0xe8, 0x57, 0x2d, 0x0a, 0x11, 0x11, 0x8e, 0x01, 0x16, 0xa4,
	0xea, 0x0c, 0xf7, 0xb7, 0xf4, 0xfe, 0xd6, 0x65, 0xc5, 0xfb, 0x2d, 0x9c, 0x81, 0xf3, 0x28, 0xf2,
	0x5a, 0x90, 0xe3, 0xbf, 0xba, 0xa5, 0xb9, 0xe8, 0x11, 0xd5, 0x84, 0xe1, 0x0c, 0xec, 0x80, 0x42,
	0x99, 0x45, 0xda, 0xfb, 0xa7, 0xd1, 0xff, 0x66, 0xca, 0x4e, 0x5e, 0x7e, 0x6b, 0xc4, 0xf0, 0xa6,
	0x3a, 0xcb, 0x7c, 0xc3, 0xd5, 0x8e, 0x59, 0x7d, 0x29, 0xef, 0x10, 0x2a, 0x65, 0x13, 0xf6, 0xda,
	0xd3, 0xbf, 0xed, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x0a, 0x86, 0x5f, 0xc6, 0x02,
	0x00, 0x00,
}

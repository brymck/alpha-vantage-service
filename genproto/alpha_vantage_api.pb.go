// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpha_vantage_api.proto

package genproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetTimeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbd68805efa1c3c, []int{0}
}

func (m *GetTimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeRequest.Unmarshal(m, b)
}
func (m *GetTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeRequest.Marshal(b, m, deterministic)
}
func (m *GetTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeRequest.Merge(m, src)
}
func (m *GetTimeRequest) XXX_Size() int {
	return xxx_messageInfo_GetTimeRequest.Size(m)
}
func (m *GetTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeRequest proto.InternalMessageInfo

func (m *GetTimeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTimeResponse struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetTimeResponse) Reset()         { *m = GetTimeResponse{} }
func (m *GetTimeResponse) String() string { return proto.CompactTextString(m) }
func (*GetTimeResponse) ProtoMessage()    {}
func (*GetTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbd68805efa1c3c, []int{1}
}

func (m *GetTimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeResponse.Unmarshal(m, b)
}
func (m *GetTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeResponse.Marshal(b, m, deterministic)
}
func (m *GetTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeResponse.Merge(m, src)
}
func (m *GetTimeResponse) XXX_Size() int {
	return xxx_messageInfo_GetTimeResponse.Size(m)
}
func (m *GetTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeResponse proto.InternalMessageInfo

func (m *GetTimeResponse) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTimeRequest)(nil), "brymck.alpha_vantage.v1.GetTimeRequest")
	proto.RegisterType((*GetTimeResponse)(nil), "brymck.alpha_vantage.v1.GetTimeResponse")
}

func init() { proto.RegisterFile("alpha_vantage_api.proto", fileDescriptor_0dbd68805efa1c3c) }

var fileDescriptor_0dbd68805efa1c3c = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcc, 0x29, 0xc8,
	0x48, 0x8c, 0x2f, 0x4b, 0xcc, 0x2b, 0x49, 0x4c, 0x4f, 0x8d, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0x2a, 0xaa, 0xcc, 0x4d, 0xce, 0xd6, 0x43, 0x91, 0xd7, 0x2b,
	0x33, 0x94, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x2b, 0x4b, 0x2a, 0x4d, 0xd3,
	0x2f, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x80, 0xe8, 0x54, 0x52, 0xe0, 0xe2, 0x73,
	0x4f, 0x2d, 0x09, 0xc9, 0xcc, 0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x51, 0x72, 0xe4,
	0xe2, 0x87, 0xab, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xd2, 0xe3, 0x62, 0x01, 0x99, 0x03,
	0x56, 0xc4, 0x6d, 0x24, 0xa5, 0x07, 0xb1, 0x44, 0x0f, 0x66, 0x89, 0x5e, 0x08, 0xcc, 0x92, 0x20,
	0xb0, 0x3a, 0xa3, 0x7c, 0x2e, 0x7e, 0x47, 0x90, 0xcb, 0xc2, 0x20, 0x0e, 0x73, 0x0c, 0xf0, 0x14,
	0x8a, 0xe1, 0x62, 0x87, 0x9a, 0x2a, 0xa4, 0xae, 0x87, 0xc3, 0xf5, 0x7a, 0xa8, 0x2e, 0x93, 0xd2,
	0x20, 0xac, 0x10, 0xe2, 0x40, 0x27, 0xae, 0x28, 0x8e, 0xf4, 0xd4, 0x3c, 0x88, 0x7b, 0xd8, 0xc0,
	0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x59, 0x9c, 0x1a, 0x3d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlphaVantageAPIClient is the client API for AlphaVantageAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlphaVantageAPIClient interface {
	// Get the current time.
	GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (*GetTimeResponse, error)
}

type alphaVantageAPIClient struct {
	cc *grpc.ClientConn
}

func NewAlphaVantageAPIClient(cc *grpc.ClientConn) AlphaVantageAPIClient {
	return &alphaVantageAPIClient{cc}
}

func (c *alphaVantageAPIClient) GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (*GetTimeResponse, error) {
	out := new(GetTimeResponse)
	err := c.cc.Invoke(ctx, "/brymck.alpha_vantage.v1.AlphaVantageAPI/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlphaVantageAPIServer is the server API for AlphaVantageAPI service.
type AlphaVantageAPIServer interface {
	// Get the current time.
	GetTime(context.Context, *GetTimeRequest) (*GetTimeResponse, error)
}

// UnimplementedAlphaVantageAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAlphaVantageAPIServer struct {
}

func (*UnimplementedAlphaVantageAPIServer) GetTime(ctx context.Context, req *GetTimeRequest) (*GetTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}

func RegisterAlphaVantageAPIServer(s *grpc.Server, srv AlphaVantageAPIServer) {
	s.RegisterService(&_AlphaVantageAPI_serviceDesc, srv)
}

func _AlphaVantageAPI_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlphaVantageAPIServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.alpha_vantage.v1.AlphaVantageAPI/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlphaVantageAPIServer).GetTime(ctx, req.(*GetTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlphaVantageAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brymck.alpha_vantage.v1.AlphaVantageAPI",
	HandlerType: (*AlphaVantageAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _AlphaVantageAPI_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpha_vantage_api.proto",
}

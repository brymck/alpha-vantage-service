// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brymck/alpha_vantage/v1/alpha_vantage_api.proto

package genproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// A request for a quote
type GetQuoteRequest struct {
	// The public ticker symbol for the stock
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQuoteRequest) Reset()         { *m = GetQuoteRequest{} }
func (m *GetQuoteRequest) String() string { return proto.CompactTextString(m) }
func (*GetQuoteRequest) ProtoMessage()    {}
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76dd78e493e1430d, []int{0}
}

func (m *GetQuoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuoteRequest.Unmarshal(m, b)
}
func (m *GetQuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuoteRequest.Marshal(b, m, deterministic)
}
func (m *GetQuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuoteRequest.Merge(m, src)
}
func (m *GetQuoteRequest) XXX_Size() int {
	return xxx_messageInfo_GetQuoteRequest.Size(m)
}
func (m *GetQuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuoteRequest proto.InternalMessageInfo

func (m *GetQuoteRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

// A response to a quote request
type GetQuoteResponse struct {
	// The open price
	Open float64 `protobuf:"fixed64,1,opt,name=open,proto3" json:"open,omitempty"`
	// The high price
	High float64 `protobuf:"fixed64,2,opt,name=high,proto3" json:"high,omitempty"`
	// The low price
	Low float64 `protobuf:"fixed64,3,opt,name=low,proto3" json:"low,omitempty"`
	// The current price
	Price float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	// The daily volume
	Volume float64 `protobuf:"fixed64,5,opt,name=volume,proto3" json:"volume,omitempty"`
	// The latest trading day
	LatestTradingDay *Date `protobuf:"bytes,6,opt,name=latest_trading_day,json=latestTradingDay,proto3" json:"latest_trading_day,omitempty"`
	// The previous close
	PreviousClose        float64  `protobuf:"fixed64,7,opt,name=previous_close,json=previousClose,proto3" json:"previous_close,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQuoteResponse) Reset()         { *m = GetQuoteResponse{} }
func (m *GetQuoteResponse) String() string { return proto.CompactTextString(m) }
func (*GetQuoteResponse) ProtoMessage()    {}
func (*GetQuoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76dd78e493e1430d, []int{1}
}

func (m *GetQuoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuoteResponse.Unmarshal(m, b)
}
func (m *GetQuoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuoteResponse.Marshal(b, m, deterministic)
}
func (m *GetQuoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuoteResponse.Merge(m, src)
}
func (m *GetQuoteResponse) XXX_Size() int {
	return xxx_messageInfo_GetQuoteResponse.Size(m)
}
func (m *GetQuoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuoteResponse proto.InternalMessageInfo

func (m *GetQuoteResponse) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *GetQuoteResponse) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *GetQuoteResponse) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *GetQuoteResponse) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GetQuoteResponse) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *GetQuoteResponse) GetLatestTradingDay() *Date {
	if m != nil {
		return m.LatestTradingDay
	}
	return nil
}

func (m *GetQuoteResponse) GetPreviousClose() float64 {
	if m != nil {
		return m.PreviousClose
	}
	return 0
}

// A request for time series
type GetTimeSeriesRequest struct {
	// The public ticker symbol for the stock
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Whether to return the full time series for the stock. By default, only the latest 100 data points are retrieved.
	// When anbled, the full-length time series of 20+ years of historical data are retrieved. Leaving this option
	// false is recommended if you would like to reduce the data size of each API call.
	Full                 bool     `protobuf:"varint,2,opt,name=full,proto3" json:"full,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeSeriesRequest) Reset()         { *m = GetTimeSeriesRequest{} }
func (m *GetTimeSeriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeSeriesRequest) ProtoMessage()    {}
func (*GetTimeSeriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76dd78e493e1430d, []int{2}
}

func (m *GetTimeSeriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeSeriesRequest.Unmarshal(m, b)
}
func (m *GetTimeSeriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeSeriesRequest.Marshal(b, m, deterministic)
}
func (m *GetTimeSeriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeSeriesRequest.Merge(m, src)
}
func (m *GetTimeSeriesRequest) XXX_Size() int {
	return xxx_messageInfo_GetTimeSeriesRequest.Size(m)
}
func (m *GetTimeSeriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeSeriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeSeriesRequest proto.InternalMessageInfo

func (m *GetTimeSeriesRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetTimeSeriesRequest) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

// A response to a time series request
type GetTimeSeriesResponse struct {
	// A series of data points
	TimeSeries           []*TimeSeriesDataPoint `protobuf:"bytes,1,rep,name=time_series,json=timeSeries,proto3" json:"time_series,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetTimeSeriesResponse) Reset()         { *m = GetTimeSeriesResponse{} }
func (m *GetTimeSeriesResponse) String() string { return proto.CompactTextString(m) }
func (*GetTimeSeriesResponse) ProtoMessage()    {}
func (*GetTimeSeriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76dd78e493e1430d, []int{3}
}

func (m *GetTimeSeriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeSeriesResponse.Unmarshal(m, b)
}
func (m *GetTimeSeriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeSeriesResponse.Marshal(b, m, deterministic)
}
func (m *GetTimeSeriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeSeriesResponse.Merge(m, src)
}
func (m *GetTimeSeriesResponse) XXX_Size() int {
	return xxx_messageInfo_GetTimeSeriesResponse.Size(m)
}
func (m *GetTimeSeriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeSeriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeSeriesResponse proto.InternalMessageInfo

func (m *GetTimeSeriesResponse) GetTimeSeries() []*TimeSeriesDataPoint {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

// An observation in a time series
type TimeSeriesDataPoint struct {
	// The observation date
	Date *Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	// The open price
	Open float64 `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	// The high price
	High float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	// The low price
	Low float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	// The close price
	Close float64 `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	// The daily volume
	Volume               float64  `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSeriesDataPoint) Reset()         { *m = TimeSeriesDataPoint{} }
func (m *TimeSeriesDataPoint) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesDataPoint) ProtoMessage()    {}
func (*TimeSeriesDataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_76dd78e493e1430d, []int{4}
}

func (m *TimeSeriesDataPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesDataPoint.Unmarshal(m, b)
}
func (m *TimeSeriesDataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesDataPoint.Marshal(b, m, deterministic)
}
func (m *TimeSeriesDataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesDataPoint.Merge(m, src)
}
func (m *TimeSeriesDataPoint) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesDataPoint.Size(m)
}
func (m *TimeSeriesDataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesDataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesDataPoint proto.InternalMessageInfo

func (m *TimeSeriesDataPoint) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *TimeSeriesDataPoint) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *TimeSeriesDataPoint) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *TimeSeriesDataPoint) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *TimeSeriesDataPoint) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *TimeSeriesDataPoint) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

// A whole or partial calendar date, where the time of day and time zone are not significant
type Date struct {
	// Year of date
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month of year
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Day of month
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_76dd78e493e1430d, []int{5}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*GetQuoteRequest)(nil), "brymck.alpha_vantage.v1.GetQuoteRequest")
	proto.RegisterType((*GetQuoteResponse)(nil), "brymck.alpha_vantage.v1.GetQuoteResponse")
	proto.RegisterType((*GetTimeSeriesRequest)(nil), "brymck.alpha_vantage.v1.GetTimeSeriesRequest")
	proto.RegisterType((*GetTimeSeriesResponse)(nil), "brymck.alpha_vantage.v1.GetTimeSeriesResponse")
	proto.RegisterType((*TimeSeriesDataPoint)(nil), "brymck.alpha_vantage.v1.TimeSeriesDataPoint")
	proto.RegisterType((*Date)(nil), "brymck.alpha_vantage.v1.Date")
}

func init() {
	proto.RegisterFile("brymck/alpha_vantage/v1/alpha_vantage_api.proto", fileDescriptor_76dd78e493e1430d)
}

var fileDescriptor_76dd78e493e1430d = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x6d, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x6b, 0x27, 0x84, 0x89, 0x4a, 0xa3, 0xa5, 0x80, 0x55, 0x09, 0x29, 0xb2, 0x84, 0x94,
	0x4a, 0xe0, 0x28, 0xe1, 0x04, 0x0d, 0x91, 0x2a, 0x84, 0x90, 0x8a, 0xa9, 0xf8, 0xc1, 0x1f, 0x6b,
	0x93, 0x4c, 0x93, 0x15, 0xf6, 0xae, 0xf1, 0xae, 0x8d, 0x7c, 0x2b, 0xce, 0xc4, 0x01, 0x38, 0x03,
	0xda, 0x59, 0xf7, 0x23, 0x6d, 0x43, 0xfb, 0x2b, 0x33, 0x2f, 0x6f, 0x66, 0xf6, 0xcd, 0x1b, 0xc3,
	0x78, 0x51, 0x36, 0xf9, 0xf2, 0xc7, 0x98, 0x67, 0xc5, 0x86, 0xa7, 0x35, 0x97, 0x86, 0xaf, 0x71,
	0x5c, 0x4f, 0xb6, 0x81, 0x94, 0x17, 0x22, 0x2e, 0x4a, 0x65, 0x14, 0x7b, 0xe5, 0x0a, 0xe2, 0xad,
	0xff, 0xe3, 0x7a, 0x12, 0x1d, 0xc3, 0xc1, 0x29, 0x9a, 0x2f, 0x95, 0x32, 0x98, 0xe0, 0xcf, 0x0a,
	0xb5, 0x61, 0x2f, 0xa1, 0xab, 0x9b, 0x7c, 0xa1, 0xb2, 0xd0, 0x1b, 0x7a, 0xa3, 0xa7, 0x49, 0x9b,
	0x45, 0x7f, 0x3d, 0x18, 0x5c, 0x73, 0x75, 0xa1, 0xa4, 0x46, 0xc6, 0x20, 0x50, 0x05, 0x4a, 0xa2,
	0x7a, 0x09, 0xc5, 0x16, 0xdb, 0x88, 0xf5, 0x26, 0xdc, 0x73, 0x98, 0x8d, 0xd9, 0x00, 0xfc, 0x4c,
	0xfd, 0x0a, 0x7d, 0x82, 0x6c, 0xc8, 0x0e, 0xa1, 0x53, 0x94, 0x62, 0x89, 0x61, 0x40, 0x98, 0x4b,
	0xec, 0xf0, 0x5a, 0x65, 0x55, 0x8e, 0x61, 0x87, 0xe0, 0x36, 0x63, 0x9f, 0x80, 0x65, 0xdc, 0xa0,
	0x36, 0xa9, 0x29, 0xf9, 0x4a, 0xc8, 0x75, 0xba, 0xe2, 0x4d, 0xd8, 0x1d, 0x7a, 0xa3, 0xfe, 0xf4,
	0x75, 0xbc, 0x43, 0x5d, 0x3c, 0xe7, 0x06, 0x93, 0x81, 0x2b, 0x3c, 0x77, 0x75, 0x73, 0xde, 0xb0,
	0x37, 0xf0, 0xac, 0x28, 0xb1, 0x16, 0xaa, 0xd2, 0xe9, 0x32, 0x53, 0x1a, 0xc3, 0x27, 0x34, 0x6c,
	0xff, 0x12, 0xfd, 0x60, 0xc1, 0x68, 0x06, 0x87, 0xa7, 0x68, 0xce, 0x45, 0x8e, 0x5f, 0xb1, 0x14,
	0xa8, 0x1f, 0x58, 0x90, 0xd5, 0x7d, 0x51, 0x65, 0x19, 0xe9, 0xee, 0x25, 0x14, 0x47, 0x17, 0xf0,
	0xe2, 0x56, 0x8f, 0x76, 0x71, 0x9f, 0xa1, 0x6f, 0x44, 0x8e, 0xa9, 0x26, 0x38, 0xf4, 0x86, 0xfe,
	0xa8, 0x3f, 0x7d, 0xbb, 0x53, 0xc9, 0x75, 0x87, 0x39, 0x37, 0xfc, 0x4c, 0x09, 0x69, 0x12, 0x30,
	0x57, 0x60, 0xf4, 0xdb, 0x83, 0xe7, 0xf7, 0x70, 0xd8, 0x04, 0x82, 0x15, 0x37, 0x48, 0x2f, 0x7d,
	0x70, 0x53, 0x44, 0xbd, 0xb2, 0x74, 0xef, 0x1e, 0x4b, 0xfd, 0xbb, 0x96, 0x06, 0x5b, 0x96, 0xba,
	0x75, 0x3a, 0xef, 0x5c, 0x72, 0xc3, 0xd2, 0xee, 0x4d, 0x4b, 0xa3, 0x19, 0x04, 0xf3, 0x76, 0x5e,
	0x83, 0xbc, 0xa4, 0x27, 0x76, 0x12, 0x8a, 0x6d, 0xa7, 0x5c, 0x49, 0xe3, 0x6e, 0xa8, 0x93, 0xb8,
	0xc4, 0x4e, 0xb4, 0xae, 0xfb, 0x84, 0xd9, 0x70, 0xfa, 0xc7, 0x83, 0x83, 0x13, 0xab, 0xe5, 0x9b,
	0x93, 0x72, 0x72, 0xf6, 0x91, 0xa5, 0xd0, 0xbb, 0x3c, 0x53, 0x36, 0xda, 0x29, 0xf8, 0xd6, 0xd5,
	0x1f, 0x1d, 0x3f, 0x82, 0xd9, 0x5a, 0x27, 0x61, 0x7f, 0xcb, 0x53, 0xf6, 0xee, 0x7f, 0xb5, 0x77,
	0xee, 0xe7, 0x28, 0x7e, 0x2c, 0xdd, 0xcd, 0x9b, 0xc1, 0xf7, 0xde, 0x1a, 0x25, 0x7d, 0xc8, 0x8b,
	0x2e, 0xfd, 0xbc, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x87, 0x78, 0xe5, 0xf9, 0x02, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlphaVantageAPIClient is the client API for AlphaVantageAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlphaVantageAPIClient interface {
	// Get the latest quote for a security
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error)
	// Get the price time series for a security
	GetTimeSeries(ctx context.Context, in *GetTimeSeriesRequest, opts ...grpc.CallOption) (*GetTimeSeriesResponse, error)
}

type alphaVantageAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAlphaVantageAPIClient(cc grpc.ClientConnInterface) AlphaVantageAPIClient {
	return &alphaVantageAPIClient{cc}
}

func (c *alphaVantageAPIClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error) {
	out := new(GetQuoteResponse)
	err := c.cc.Invoke(ctx, "/brymck.alpha_vantage.v1.AlphaVantageAPI/GetQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alphaVantageAPIClient) GetTimeSeries(ctx context.Context, in *GetTimeSeriesRequest, opts ...grpc.CallOption) (*GetTimeSeriesResponse, error) {
	out := new(GetTimeSeriesResponse)
	err := c.cc.Invoke(ctx, "/brymck.alpha_vantage.v1.AlphaVantageAPI/GetTimeSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlphaVantageAPIServer is the server API for AlphaVantageAPI service.
type AlphaVantageAPIServer interface {
	// Get the latest quote for a security
	GetQuote(context.Context, *GetQuoteRequest) (*GetQuoteResponse, error)
	// Get the price time series for a security
	GetTimeSeries(context.Context, *GetTimeSeriesRequest) (*GetTimeSeriesResponse, error)
}

// UnimplementedAlphaVantageAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAlphaVantageAPIServer struct {
}

func (*UnimplementedAlphaVantageAPIServer) GetQuote(ctx context.Context, req *GetQuoteRequest) (*GetQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuote not implemented")
}
func (*UnimplementedAlphaVantageAPIServer) GetTimeSeries(ctx context.Context, req *GetTimeSeriesRequest) (*GetTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeries not implemented")
}

func RegisterAlphaVantageAPIServer(s *grpc.Server, srv AlphaVantageAPIServer) {
	s.RegisterService(&_AlphaVantageAPI_serviceDesc, srv)
}

func _AlphaVantageAPI_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlphaVantageAPIServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.alpha_vantage.v1.AlphaVantageAPI/GetQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlphaVantageAPIServer).GetQuote(ctx, req.(*GetQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlphaVantageAPI_GetTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlphaVantageAPIServer).GetTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.alpha_vantage.v1.AlphaVantageAPI/GetTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlphaVantageAPIServer).GetTimeSeries(ctx, req.(*GetTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlphaVantageAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brymck.alpha_vantage.v1.AlphaVantageAPI",
	HandlerType: (*AlphaVantageAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _AlphaVantageAPI_GetQuote_Handler,
		},
		{
			MethodName: "GetTimeSeries",
			Handler:    _AlphaVantageAPI_GetTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brymck/alpha_vantage/v1/alpha_vantage_api.proto",
}

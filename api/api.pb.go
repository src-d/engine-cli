// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	VersionRequest
	VersionResponse
	ParseRequest
	ParseResponse
	ListDriversRequest
	ListDriversResponse
	SQLRequest
	SQLResponse
	StartBblfshWebRequest
	StartBblfshWebResponse
	StartGitbaseWebRequest
	StartGitbaseWebResponse
	StopGitbaseWebRequest
	StopGitbaseWebResponse
	StopBblfshWebRequest
	StopBblfshWebResponse
*/
package api

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

type ParseRequest_Kind int32

const (
	ParseRequest_INVALID ParseRequest_Kind = 0
	ParseRequest_LANG    ParseRequest_Kind = 1
	ParseRequest_UAST    ParseRequest_Kind = 2
)

var ParseRequest_Kind_name = map[int32]string{
	0: "INVALID",
	1: "LANG",
	2: "UAST",
}
var ParseRequest_Kind_value = map[string]int32{
	"INVALID": 0,
	"LANG":    1,
	"UAST":    2,
}

func (x ParseRequest_Kind) String() string {
	return proto.EnumName(ParseRequest_Kind_name, int32(x))
}
func (ParseRequest_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ParseResponse_Kind int32

const (
	ParseResponse_INVALID ParseResponse_Kind = 0
	ParseResponse_LOG     ParseResponse_Kind = 1
	ParseResponse_FINAL   ParseResponse_Kind = 2
)

var ParseResponse_Kind_name = map[int32]string{
	0: "INVALID",
	1: "LOG",
	2: "FINAL",
}
var ParseResponse_Kind_value = map[string]int32{
	"INVALID": 0,
	"LOG":     1,
	"FINAL":   2,
}

func (x ParseResponse_Kind) String() string {
	return proto.EnumName(ParseResponse_Kind_name, int32(x))
}
func (ParseResponse_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type VersionResponse struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ParseRequest struct {
	Kind    ParseRequest_Kind `protobuf:"varint,1,opt,name=kind,enum=ParseRequest_Kind" json:"kind,omitempty"`
	Name    string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Content []byte            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// used for UAST and Native only
	Lang  string `protobuf:"bytes,4,opt,name=lang" json:"lang,omitempty"`
	Query string `protobuf:"bytes,5,opt,name=query" json:"query,omitempty"`
}

func (m *ParseRequest) Reset()                    { *m = ParseRequest{} }
func (m *ParseRequest) String() string            { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()               {}
func (*ParseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ParseRequest) GetKind() ParseRequest_Kind {
	if m != nil {
		return m.Kind
	}
	return ParseRequest_INVALID
}

func (m *ParseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ParseRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ParseRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *ParseRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type ParseResponse struct {
	Kind ParseResponse_Kind `protobuf:"varint,1,opt,name=kind,enum=ParseResponse_Kind" json:"kind,omitempty"`
	Lang string             `protobuf:"bytes,2,opt,name=lang" json:"lang,omitempty"`
	Uast [][]byte           `protobuf:"bytes,3,rep,name=uast,proto3" json:"uast,omitempty"`
	Log  string             `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
}

func (m *ParseResponse) Reset()                    { *m = ParseResponse{} }
func (m *ParseResponse) String() string            { return proto.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()               {}
func (*ParseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ParseResponse) GetKind() ParseResponse_Kind {
	if m != nil {
		return m.Kind
	}
	return ParseResponse_INVALID
}

func (m *ParseResponse) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *ParseResponse) GetUast() [][]byte {
	if m != nil {
		return m.Uast
	}
	return nil
}

func (m *ParseResponse) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type ListDriversRequest struct {
}

func (m *ListDriversRequest) Reset()                    { *m = ListDriversRequest{} }
func (m *ListDriversRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDriversRequest) ProtoMessage()               {}
func (*ListDriversRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListDriversResponse struct {
	Drivers []*ListDriversResponse_DriverInfo `protobuf:"bytes,1,rep,name=drivers" json:"drivers,omitempty"`
}

func (m *ListDriversResponse) Reset()                    { *m = ListDriversResponse{} }
func (m *ListDriversResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDriversResponse) ProtoMessage()               {}
func (*ListDriversResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListDriversResponse) GetDrivers() []*ListDriversResponse_DriverInfo {
	if m != nil {
		return m.Drivers
	}
	return nil
}

type ListDriversResponse_DriverInfo struct {
	Lang    string `protobuf:"bytes,1,opt,name=lang" json:"lang,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *ListDriversResponse_DriverInfo) Reset()         { *m = ListDriversResponse_DriverInfo{} }
func (m *ListDriversResponse_DriverInfo) String() string { return proto.CompactTextString(m) }
func (*ListDriversResponse_DriverInfo) ProtoMessage()    {}
func (*ListDriversResponse_DriverInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *ListDriversResponse_DriverInfo) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *ListDriversResponse_DriverInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type SQLRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *SQLRequest) Reset()                    { *m = SQLRequest{} }
func (m *SQLRequest) String() string            { return proto.CompactTextString(m) }
func (*SQLRequest) ProtoMessage()               {}
func (*SQLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SQLRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SQLResponse struct {
	Header *SQLResponse_Row   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Rows   []*SQLResponse_Row `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
}

func (m *SQLResponse) Reset()                    { *m = SQLResponse{} }
func (m *SQLResponse) String() string            { return proto.CompactTextString(m) }
func (*SQLResponse) ProtoMessage()               {}
func (*SQLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SQLResponse) GetHeader() *SQLResponse_Row {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SQLResponse) GetRows() []*SQLResponse_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type SQLResponse_Row struct {
	Cell []string `protobuf:"bytes,1,rep,name=cell" json:"cell,omitempty"`
}

func (m *SQLResponse_Row) Reset()                    { *m = SQLResponse_Row{} }
func (m *SQLResponse_Row) String() string            { return proto.CompactTextString(m) }
func (*SQLResponse_Row) ProtoMessage()               {}
func (*SQLResponse_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *SQLResponse_Row) GetCell() []string {
	if m != nil {
		return m.Cell
	}
	return nil
}

type StartBblfshWebRequest struct {
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
}

func (m *StartBblfshWebRequest) Reset()                    { *m = StartBblfshWebRequest{} }
func (m *StartBblfshWebRequest) String() string            { return proto.CompactTextString(m) }
func (*StartBblfshWebRequest) ProtoMessage()               {}
func (*StartBblfshWebRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StartBblfshWebRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type StartBblfshWebResponse struct {
}

func (m *StartBblfshWebResponse) Reset()                    { *m = StartBblfshWebResponse{} }
func (m *StartBblfshWebResponse) String() string            { return proto.CompactTextString(m) }
func (*StartBblfshWebResponse) ProtoMessage()               {}
func (*StartBblfshWebResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type StartGitbaseWebRequest struct {
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
}

func (m *StartGitbaseWebRequest) Reset()                    { *m = StartGitbaseWebRequest{} }
func (m *StartGitbaseWebRequest) String() string            { return proto.CompactTextString(m) }
func (*StartGitbaseWebRequest) ProtoMessage()               {}
func (*StartGitbaseWebRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StartGitbaseWebRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type StartGitbaseWebResponse struct {
}

func (m *StartGitbaseWebResponse) Reset()                    { *m = StartGitbaseWebResponse{} }
func (m *StartGitbaseWebResponse) String() string            { return proto.CompactTextString(m) }
func (*StartGitbaseWebResponse) ProtoMessage()               {}
func (*StartGitbaseWebResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type StopGitbaseWebRequest struct {
}

func (m *StopGitbaseWebRequest) Reset()                    { *m = StopGitbaseWebRequest{} }
func (m *StopGitbaseWebRequest) String() string            { return proto.CompactTextString(m) }
func (*StopGitbaseWebRequest) ProtoMessage()               {}
func (*StopGitbaseWebRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type StopGitbaseWebResponse struct {
}

func (m *StopGitbaseWebResponse) Reset()                    { *m = StopGitbaseWebResponse{} }
func (m *StopGitbaseWebResponse) String() string            { return proto.CompactTextString(m) }
func (*StopGitbaseWebResponse) ProtoMessage()               {}
func (*StopGitbaseWebResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type StopBblfshWebRequest struct {
}

func (m *StopBblfshWebRequest) Reset()                    { *m = StopBblfshWebRequest{} }
func (m *StopBblfshWebRequest) String() string            { return proto.CompactTextString(m) }
func (*StopBblfshWebRequest) ProtoMessage()               {}
func (*StopBblfshWebRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type StopBblfshWebResponse struct {
}

func (m *StopBblfshWebResponse) Reset()                    { *m = StopBblfshWebResponse{} }
func (m *StopBblfshWebResponse) String() string            { return proto.CompactTextString(m) }
func (*StopBblfshWebResponse) ProtoMessage()               {}
func (*StopBblfshWebResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*VersionRequest)(nil), "VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "VersionResponse")
	proto.RegisterType((*ParseRequest)(nil), "ParseRequest")
	proto.RegisterType((*ParseResponse)(nil), "ParseResponse")
	proto.RegisterType((*ListDriversRequest)(nil), "ListDriversRequest")
	proto.RegisterType((*ListDriversResponse)(nil), "ListDriversResponse")
	proto.RegisterType((*ListDriversResponse_DriverInfo)(nil), "ListDriversResponse.DriverInfo")
	proto.RegisterType((*SQLRequest)(nil), "SQLRequest")
	proto.RegisterType((*SQLResponse)(nil), "SQLResponse")
	proto.RegisterType((*SQLResponse_Row)(nil), "SQLResponse.Row")
	proto.RegisterType((*StartBblfshWebRequest)(nil), "StartBblfshWebRequest")
	proto.RegisterType((*StartBblfshWebResponse)(nil), "StartBblfshWebResponse")
	proto.RegisterType((*StartGitbaseWebRequest)(nil), "StartGitbaseWebRequest")
	proto.RegisterType((*StartGitbaseWebResponse)(nil), "StartGitbaseWebResponse")
	proto.RegisterType((*StopGitbaseWebRequest)(nil), "StopGitbaseWebRequest")
	proto.RegisterType((*StopGitbaseWebResponse)(nil), "StopGitbaseWebResponse")
	proto.RegisterType((*StopBblfshWebRequest)(nil), "StopBblfshWebRequest")
	proto.RegisterType((*StopBblfshWebResponse)(nil), "StopBblfshWebResponse")
	proto.RegisterEnum("ParseRequest_Kind", ParseRequest_Kind_name, ParseRequest_Kind_value)
	proto.RegisterEnum("ParseResponse_Kind", ParseResponse_Kind_name, ParseResponse_Kind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Engine service

type EngineClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// A single response with the parsing result.
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	// A stream of responses with logs and finally the parsing result.
	ParseWithLogs(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (Engine_ParseWithLogsClient, error)
	// Driver management.
	// List all drivers.
	ListDrivers(ctx context.Context, in *ListDriversRequest, opts ...grpc.CallOption) (*ListDriversResponse, error)
	// SQL stuff.
	SQL(ctx context.Context, in *SQLRequest, opts ...grpc.CallOption) (*SQLResponse, error)
	// Start a bblfsh playground.
	StartBblfshWeb(ctx context.Context, in *StartBblfshWebRequest, opts ...grpc.CallOption) (*StartBblfshWebResponse, error)
	// Start a gitbase playground.
	StartGitbaseWeb(ctx context.Context, in *StartGitbaseWebRequest, opts ...grpc.CallOption) (*StartGitbaseWebResponse, error)
	// Stop a bblfsh playground.
	StopBblfshWeb(ctx context.Context, in *StopBblfshWebRequest, opts ...grpc.CallOption) (*StopBblfshWebResponse, error)
	// Stop a gitbase playground.
	StopGitbaseWeb(ctx context.Context, in *StopGitbaseWebRequest, opts ...grpc.CallOption) (*StopGitbaseWebResponse, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/Engine/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := grpc.Invoke(ctx, "/Engine/Parse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ParseWithLogs(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (Engine_ParseWithLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Engine_serviceDesc.Streams[0], c.cc, "/Engine/ParseWithLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &engineParseWithLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Engine_ParseWithLogsClient interface {
	Recv() (*ParseResponse, error)
	grpc.ClientStream
}

type engineParseWithLogsClient struct {
	grpc.ClientStream
}

func (x *engineParseWithLogsClient) Recv() (*ParseResponse, error) {
	m := new(ParseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *engineClient) ListDrivers(ctx context.Context, in *ListDriversRequest, opts ...grpc.CallOption) (*ListDriversResponse, error) {
	out := new(ListDriversResponse)
	err := grpc.Invoke(ctx, "/Engine/ListDrivers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) SQL(ctx context.Context, in *SQLRequest, opts ...grpc.CallOption) (*SQLResponse, error) {
	out := new(SQLResponse)
	err := grpc.Invoke(ctx, "/Engine/SQL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) StartBblfshWeb(ctx context.Context, in *StartBblfshWebRequest, opts ...grpc.CallOption) (*StartBblfshWebResponse, error) {
	out := new(StartBblfshWebResponse)
	err := grpc.Invoke(ctx, "/Engine/StartBblfshWeb", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) StartGitbaseWeb(ctx context.Context, in *StartGitbaseWebRequest, opts ...grpc.CallOption) (*StartGitbaseWebResponse, error) {
	out := new(StartGitbaseWebResponse)
	err := grpc.Invoke(ctx, "/Engine/StartGitbaseWeb", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) StopBblfshWeb(ctx context.Context, in *StopBblfshWebRequest, opts ...grpc.CallOption) (*StopBblfshWebResponse, error) {
	out := new(StopBblfshWebResponse)
	err := grpc.Invoke(ctx, "/Engine/StopBblfshWeb", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) StopGitbaseWeb(ctx context.Context, in *StopGitbaseWebRequest, opts ...grpc.CallOption) (*StopGitbaseWebResponse, error) {
	out := new(StopGitbaseWebResponse)
	err := grpc.Invoke(ctx, "/Engine/StopGitbaseWeb", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Engine service

type EngineServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// A single response with the parsing result.
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	// A stream of responses with logs and finally the parsing result.
	ParseWithLogs(*ParseRequest, Engine_ParseWithLogsServer) error
	// Driver management.
	// List all drivers.
	ListDrivers(context.Context, *ListDriversRequest) (*ListDriversResponse, error)
	// SQL stuff.
	SQL(context.Context, *SQLRequest) (*SQLResponse, error)
	// Start a bblfsh playground.
	StartBblfshWeb(context.Context, *StartBblfshWebRequest) (*StartBblfshWebResponse, error)
	// Start a gitbase playground.
	StartGitbaseWeb(context.Context, *StartGitbaseWebRequest) (*StartGitbaseWebResponse, error)
	// Stop a bblfsh playground.
	StopBblfshWeb(context.Context, *StopBblfshWebRequest) (*StopBblfshWebResponse, error)
	// Stop a gitbase playground.
	StopGitbaseWeb(context.Context, *StopGitbaseWebRequest) (*StopGitbaseWebResponse, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ParseWithLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ParseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EngineServer).ParseWithLogs(m, &engineParseWithLogsServer{stream})
}

type Engine_ParseWithLogsServer interface {
	Send(*ParseResponse) error
	grpc.ServerStream
}

type engineParseWithLogsServer struct {
	grpc.ServerStream
}

func (x *engineParseWithLogsServer) Send(m *ParseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Engine_ListDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDriversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ListDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/ListDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ListDrivers(ctx, req.(*ListDriversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_SQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).SQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/SQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).SQL(ctx, req.(*SQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_StartBblfshWeb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBblfshWebRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).StartBblfshWeb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/StartBblfshWeb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).StartBblfshWeb(ctx, req.(*StartBblfshWebRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_StartGitbaseWeb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGitbaseWebRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).StartGitbaseWeb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/StartGitbaseWeb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).StartGitbaseWeb(ctx, req.(*StartGitbaseWebRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_StopBblfshWeb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopBblfshWebRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).StopBblfshWeb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/StopBblfshWeb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).StopBblfshWeb(ctx, req.(*StopBblfshWebRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_StopGitbaseWeb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopGitbaseWebRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).StopGitbaseWeb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/StopGitbaseWeb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).StopGitbaseWeb(ctx, req.(*StopGitbaseWebRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Engine_Version_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _Engine_Parse_Handler,
		},
		{
			MethodName: "ListDrivers",
			Handler:    _Engine_ListDrivers_Handler,
		},
		{
			MethodName: "SQL",
			Handler:    _Engine_SQL_Handler,
		},
		{
			MethodName: "StartBblfshWeb",
			Handler:    _Engine_StartBblfshWeb_Handler,
		},
		{
			MethodName: "StartGitbaseWeb",
			Handler:    _Engine_StartGitbaseWeb_Handler,
		},
		{
			MethodName: "StopBblfshWeb",
			Handler:    _Engine_StopBblfshWeb_Handler,
		},
		{
			MethodName: "StopGitbaseWeb",
			Handler:    _Engine_StopGitbaseWeb_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ParseWithLogs",
			Handler:       _Engine_ParseWithLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0xb5, 0x63, 0x87, 0xdc, 0x4c, 0x42, 0xb0, 0x86, 0x90, 0x18, 0xbf, 0x5c, 0xb4, 0xaa, 0x8a,
	0x25, 0xaa, 0x55, 0x95, 0x3e, 0x95, 0xa7, 0xa6, 0xa5, 0xa0, 0xa8, 0x16, 0x2d, 0x4e, 0x0b, 0xcf,
	0x0e, 0x59, 0xc0, 0x6a, 0xea, 0x0d, 0xb6, 0x29, 0xed, 0x2f, 0x54, 0xfd, 0x83, 0x7e, 0x48, 0x7f,
	0xaf, 0xda, 0xf5, 0x1a, 0xec, 0xd8, 0x11, 0x6f, 0xb3, 0x73, 0xce, 0xce, 0x9e, 0xf1, 0x99, 0x31,
	0xb4, 0x83, 0x65, 0x48, 0x97, 0x31, 0x4f, 0x39, 0xb1, 0xa0, 0x77, 0xce, 0xe2, 0x24, 0xe4, 0x91,
	0xcf, 0x6e, 0xef, 0x58, 0x92, 0x92, 0x03, 0xd8, 0x7a, 0xc8, 0x24, 0x4b, 0x1e, 0x25, 0x0c, 0x6d,
	0x68, 0x7d, 0xcf, 0x52, 0xb6, 0xbe, 0xa7, 0xbb, 0x6d, 0x3f, 0x3f, 0x92, 0xbf, 0x3a, 0x74, 0x3f,
	0x05, 0x71, 0xc2, 0xd4, 0x6d, 0x7c, 0x0e, 0xe6, 0xd7, 0x30, 0x9a, 0x4b, 0x5e, 0x6f, 0x84, 0xb4,
	0x08, 0xd2, 0x0f, 0x61, 0x34, 0xf7, 0x25, 0x8e, 0x08, 0x66, 0x14, 0x7c, 0x63, 0x76, 0x43, 0xd6,
	0x93, 0xb1, 0x78, 0xe6, 0x92, 0x47, 0x29, 0x8b, 0x52, 0xdb, 0xd8, 0xd3, 0xdd, 0xae, 0x9f, 0x1f,
	0x05, 0x7b, 0x11, 0x44, 0xd7, 0xb6, 0x99, 0xb1, 0x45, 0x8c, 0x7d, 0x68, 0xde, 0xde, 0xb1, 0xf8,
	0xa7, 0xdd, 0x94, 0xc9, 0xec, 0x40, 0xf6, 0xc1, 0x14, 0xaf, 0x60, 0x07, 0x5a, 0x93, 0xd3, 0xf3,
	0xb1, 0x37, 0x39, 0xb2, 0x34, 0xfc, 0x0f, 0x4c, 0x6f, 0x7c, 0x7a, 0x62, 0xe9, 0x22, 0xfa, 0x32,
	0x9e, 0x7e, 0xb6, 0x1a, 0xe4, 0x8f, 0x0e, 0x9b, 0x4a, 0x9c, 0xea, 0x72, 0xbf, 0x24, 0x7d, 0x9b,
	0x96, 0xd0, 0x15, 0xed, 0x52, 0x4d, 0xa3, 0xa0, 0x06, 0xc1, 0xbc, 0x0b, 0x12, 0x21, 0xdc, 0x70,
	0xbb, 0xbe, 0x8c, 0xd1, 0x02, 0x63, 0xc1, 0x73, 0xd1, 0x22, 0xac, 0x57, 0xd7, 0x02, 0xc3, 0xfb,
	0x28, 0xc4, 0xb5, 0xa1, 0x79, 0x3c, 0x39, 0x1d, 0x7b, 0x56, 0x83, 0xf4, 0x01, 0xbd, 0x30, 0x49,
	0x8f, 0xe2, 0x50, 0x7c, 0xe9, 0xdc, 0x9a, 0xdf, 0x3a, 0x6c, 0x97, 0xd2, 0x4a, 0xf9, 0x6b, 0x68,
	0xcd, 0xb3, 0x94, 0xad, 0xef, 0x19, 0x6e, 0x67, 0xf4, 0x3f, 0xad, 0xa1, 0xd1, 0xec, 0x3c, 0x89,
	0xae, 0xb8, 0x9f, 0xf3, 0x9d, 0x43, 0x80, 0xc7, 0xf4, 0x43, 0x67, 0x7a, 0xa1, 0xb3, 0x82, 0xf9,
	0x8d, 0xb2, 0xf9, 0x04, 0x60, 0x7a, 0xe6, 0xe5, 0xce, 0x3f, 0xf8, 0xa1, 0x17, 0xfd, 0xf8, 0x01,
	0x1d, 0xc9, 0x51, 0x4a, 0x5d, 0xd8, 0xb8, 0x61, 0xc1, 0x9c, 0xc5, 0x92, 0xd5, 0x19, 0x59, 0xb4,
	0x80, 0x52, 0x9f, 0xdf, 0xfb, 0x0a, 0xc7, 0x67, 0x60, 0xc6, 0xfc, 0x3e, 0xb1, 0x1b, 0xb2, 0xa1,
	0x2a, 0x4f, 0xa2, 0xce, 0x2e, 0x18, 0x3e, 0xbf, 0x17, 0xba, 0x2f, 0xd9, 0x62, 0x21, 0xbb, 0x6f,
	0xfb, 0x32, 0x26, 0x07, 0xb0, 0x33, 0x4d, 0x83, 0x38, 0x7d, 0x3b, 0x5b, 0x5c, 0x25, 0x37, 0x17,
	0x6c, 0x96, 0x0b, 0x45, 0x30, 0x97, 0x3c, 0x4e, 0xa5, 0x82, 0xa6, 0x2f, 0x63, 0x62, 0xc3, 0x60,
	0x95, 0x9c, 0xbd, 0x45, 0x5e, 0x28, 0xe4, 0x24, 0x4c, 0x67, 0x41, 0xc2, 0x9e, 0xa8, 0xb3, 0x0b,
	0xc3, 0x0a, 0x5b, 0x15, 0x1a, 0x0a, 0x3d, 0x7c, 0x59, 0xa9, 0x93, 0xbd, 0x5d, 0x06, 0xd4, 0x95,
	0x01, 0xf4, 0x05, 0xb2, 0xda, 0x41, 0x5e, 0xaa, 0x22, 0x76, 0xf4, 0xcb, 0x84, 0x8d, 0xf7, 0xd1,
	0x75, 0x18, 0x31, 0xa4, 0xd0, 0x52, 0x6b, 0x8c, 0x5b, 0xb4, 0xbc, 0xe2, 0x8e, 0x45, 0x57, 0x36,
	0x9c, 0x68, 0xe8, 0x42, 0x53, 0x0e, 0x3c, 0x6e, 0x96, 0x76, 0xd6, 0xe9, 0x95, 0xf7, 0x80, 0x68,
	0x38, 0x52, 0x8b, 0x73, 0x11, 0xa6, 0x37, 0x1e, 0xbf, 0x4e, 0x9e, 0xbc, 0xf1, 0x52, 0xc7, 0x43,
	0xe8, 0x14, 0x26, 0x12, 0xb7, 0x69, 0x75, 0xba, 0x9d, 0x7e, 0xdd, 0xd0, 0x12, 0x0d, 0x09, 0x18,
	0xd3, 0x33, 0x0f, 0x3b, 0xf4, 0x71, 0xd8, 0x9c, 0x6e, 0x71, 0x1e, 0x88, 0x86, 0xef, 0xa0, 0x57,
	0xf6, 0x0f, 0x07, 0xb4, 0xd6, 0x7d, 0x67, 0x48, 0xd7, 0x18, 0xad, 0xe1, 0x31, 0x6c, 0xad, 0x98,
	0x87, 0x8a, 0x5d, 0x31, 0xcd, 0xb1, 0xe9, 0x3a, 0x9f, 0x35, 0x7c, 0x03, 0x9b, 0x25, 0x7b, 0x70,
	0x87, 0xd6, 0xd9, 0xe8, 0x0c, 0x68, 0xad, 0x8b, 0x79, 0x3b, 0xc5, 0x91, 0xc0, 0x8c, 0x5b, 0xd5,
	0x31, 0xa4, 0x6b, 0x66, 0x47, 0x9b, 0x6d, 0xc8, 0x3f, 0xfc, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xee, 0x1e, 0x9f, 0xc2, 0xee, 0x05, 0x00, 0x00,
}

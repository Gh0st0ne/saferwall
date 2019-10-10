// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiav.mcafee.proto

package mcafee_api

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

// The scan file request message containing the file path to scan.
type ScanFileRequest struct {
	Filepath             string   `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanFileRequest) Reset()         { *m = ScanFileRequest{} }
func (m *ScanFileRequest) String() string { return proto.CompactTextString(m) }
func (*ScanFileRequest) ProtoMessage()    {}
func (*ScanFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa94ec8f462e068, []int{0}
}

func (m *ScanFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanFileRequest.Unmarshal(m, b)
}
func (m *ScanFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanFileRequest.Marshal(b, m, deterministic)
}
func (m *ScanFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanFileRequest.Merge(m, src)
}
func (m *ScanFileRequest) XXX_Size() int {
	return xxx_messageInfo_ScanFileRequest.Size(m)
}
func (m *ScanFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanFileRequest proto.InternalMessageInfo

func (m *ScanFileRequest) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

// The scan response message containing detection results of the AntiVirus.
type ScanResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Infected             bool     `protobuf:"varint,2,opt,name=infected,proto3" json:"infected,omitempty"`
	Update               int64    `protobuf:"varint,3,opt,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanResponse) Reset()         { *m = ScanResponse{} }
func (m *ScanResponse) String() string { return proto.CompactTextString(m) }
func (*ScanResponse) ProtoMessage()    {}
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa94ec8f462e068, []int{1}
}

func (m *ScanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResponse.Unmarshal(m, b)
}
func (m *ScanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResponse.Marshal(b, m, deterministic)
}
func (m *ScanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResponse.Merge(m, src)
}
func (m *ScanResponse) XXX_Size() int {
	return xxx_messageInfo_ScanResponse.Size(m)
}
func (m *ScanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResponse proto.InternalMessageInfo

func (m *ScanResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *ScanResponse) GetInfected() bool {
	if m != nil {
		return m.Infected
	}
	return false
}

func (m *ScanResponse) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

// The version request message ask for version.
type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa94ec8f462e068, []int{2}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

// The response message containing program/VPS version.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa94ec8f462e068, []int{3}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ScanFileRequest)(nil), "mcafee.api.ScanFileRequest")
	proto.RegisterType((*ScanResponse)(nil), "mcafee.api.ScanResponse")
	proto.RegisterType((*VersionRequest)(nil), "mcafee.api.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "mcafee.api.VersionResponse")
}

func init() { proto.RegisterFile("multiav.mcafee.proto", fileDescriptor_4fa94ec8f462e068) }

var fileDescriptor_4fa94ec8f462e068 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x8d, 0x0b, 0x6b, 0x1d, 0xd4, 0x95, 0x20, 0x12, 0xba, 0x97, 0x92, 0x53, 0x41, 0xec,
	0x41, 0x7f, 0x81, 0x08, 0xee, 0xc9, 0x4b, 0x04, 0x0f, 0xde, 0x62, 0x77, 0x8a, 0x81, 0x6e, 0x12,
	0xdb, 0xc9, 0xfe, 0x1b, 0xff, 0xab, 0x54, 0xa7, 0x55, 0x8b, 0xc7, 0x6f, 0xe6, 0x65, 0xde, 0x7b,
	0x81, 0x8b, 0x5d, 0x6a, 0xc9, 0xd9, 0x7d, 0xb5, 0xab, 0x6d, 0x83, 0x58, 0xc5, 0x2e, 0x50, 0x90,
	0xc0, 0x64, 0xa3, 0xd3, 0xd7, 0xb0, 0x7a, 0xaa, 0xad, 0x7f, 0x70, 0x2d, 0x1a, 0x7c, 0x4f, 0xd8,
	0x93, 0xcc, 0x21, 0x6b, 0x5c, 0x8b, 0xd1, 0xd2, 0x9b, 0x12, 0x85, 0x28, 0x8f, 0xcd, 0xc4, 0xfa,
	0x05, 0x4e, 0x06, 0xb9, 0xc1, 0x3e, 0x06, 0xdf, 0xa3, 0xbc, 0x84, 0x65, 0x48, 0x14, 0x13, 0xb1,
	0x92, 0x69, 0xb8, 0xe1, 0x7c, 0x83, 0x35, 0xe1, 0x56, 0x1d, 0x16, 0xa2, 0xcc, 0xcc, 0xc4, 0xc3,
	0x9b, 0x14, 0xb7, 0x96, 0x50, 0x2d, 0x0a, 0x51, 0x2e, 0x0c, 0x93, 0x3e, 0x87, 0xb3, 0x67, 0xec,
	0x7a, 0x17, 0x3c, 0x27, 0xd1, 0x57, 0xb0, 0x9a, 0x26, 0x6c, 0xa8, 0xe0, 0x68, 0xff, 0x3d, 0x62,
	0xc7, 0x11, 0x6f, 0x3e, 0x04, 0x9c, 0x3e, 0xd6, 0x77, 0x0d, 0xe2, 0x90, 0xd0, 0x63, 0x27, 0xef,
	0x21, 0x1b, 0xbb, 0xc9, 0x75, 0xf5, 0x53, 0xba, 0x9a, 0x35, 0xce, 0xd5, 0x7c, 0x39, 0xda, 0xe9,
	0x03, 0xb9, 0x01, 0xd8, 0x20, 0x71, 0x0c, 0x99, 0xff, 0x56, 0xfe, 0x4d, 0x9b, 0xaf, 0xff, 0xdd,
	0x8d, 0x87, 0x5e, 0x97, 0x5f, 0x9f, 0x7f, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x68, 0xc2, 0x1d,
	0xd5, 0x94, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// McAfeeScannerClient is the client API for McAfeeScanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type McAfeeScannerClient interface {
	// Scan a file
	ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	// Get program version
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type mcAfeeScannerClient struct {
	cc *grpc.ClientConn
}

func NewMcAfeeScannerClient(cc *grpc.ClientConn) McAfeeScannerClient {
	return &mcAfeeScannerClient{cc}
}

func (c *mcAfeeScannerClient) ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/mcafee.api.McAfeeScanner/ScanFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcAfeeScannerClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/mcafee.api.McAfeeScanner/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// McAfeeScannerServer is the server API for McAfeeScanner service.
type McAfeeScannerServer interface {
	// Scan a file
	ScanFile(context.Context, *ScanFileRequest) (*ScanResponse, error)
	// Get program version
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedMcAfeeScannerServer can be embedded to have forward compatible implementations.
type UnimplementedMcAfeeScannerServer struct {
}

func (*UnimplementedMcAfeeScannerServer) ScanFile(ctx context.Context, req *ScanFileRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanFile not implemented")
}
func (*UnimplementedMcAfeeScannerServer) GetVersion(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterMcAfeeScannerServer(s *grpc.Server, srv McAfeeScannerServer) {
	s.RegisterService(&_McAfeeScanner_serviceDesc, srv)
}

func _McAfeeScanner_ScanFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McAfeeScannerServer).ScanFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcafee.api.McAfeeScanner/ScanFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McAfeeScannerServer).ScanFile(ctx, req.(*ScanFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McAfeeScanner_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McAfeeScannerServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcafee.api.McAfeeScanner/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McAfeeScannerServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _McAfeeScanner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mcafee.api.McAfeeScanner",
	HandlerType: (*McAfeeScannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanFile",
			Handler:    _McAfeeScanner_ScanFile_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _McAfeeScanner_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiav.mcafee.proto",
}
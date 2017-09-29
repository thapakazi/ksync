// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/radar.proto

/*
Package proto_radar is a generated protocol buffer package.

It is generated from these files:
	proto/radar.proto

It has these top-level messages:
	ContainerPath
	Files
	File
*/
package proto_radar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type ContainerPath struct {
	ContainerId string `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	PathName    string `protobuf:"bytes,2,opt,name=path_name,json=pathName" json:"path_name,omitempty"`
}

func (m *ContainerPath) Reset()                    { *m = ContainerPath{} }
func (m *ContainerPath) String() string            { return proto.CompactTextString(m) }
func (*ContainerPath) ProtoMessage()               {}
func (*ContainerPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContainerPath) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *ContainerPath) GetPathName() string {
	if m != nil {
		return m.PathName
	}
	return ""
}

type Files struct {
	File []*File `protobuf:"bytes,1,rep,name=file" json:"file,omitempty"`
}

func (m *Files) Reset()                    { *m = Files{} }
func (m *Files) String() string            { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()               {}
func (*Files) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Files) GetFile() []*File {
	if m != nil {
		return m.File
	}
	return nil
}

type File struct {
	Path    string                     `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Size    int64                      `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Mode    uint32                     `protobuf:"varint,3,opt,name=mode" json:"mode,omitempty"`
	ModTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=mod_time,json=modTime" json:"mod_time,omitempty"`
	IsDir   bool                       `protobuf:"varint,5,opt,name=is_dir,json=isDir" json:"is_dir,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *File) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *File) GetModTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ModTime
	}
	return nil
}

func (m *File) GetIsDir() bool {
	if m != nil {
		return m.IsDir
	}
	return false
}

func init() {
	proto.RegisterType((*ContainerPath)(nil), "proto.radar.ContainerPath")
	proto.RegisterType((*Files)(nil), "proto.radar.Files")
	proto.RegisterType((*File)(nil), "proto.radar.File")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Radar service

type RadarClient interface {
	ListContainerFiles(ctx context.Context, in *ContainerPath, opts ...grpc.CallOption) (*Files, error)
}

type radarClient struct {
	cc *grpc.ClientConn
}

func NewRadarClient(cc *grpc.ClientConn) RadarClient {
	return &radarClient{cc}
}

func (c *radarClient) ListContainerFiles(ctx context.Context, in *ContainerPath, opts ...grpc.CallOption) (*Files, error) {
	out := new(Files)
	err := grpc.Invoke(ctx, "/proto.radar.Radar/ListContainerFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Radar service

type RadarServer interface {
	ListContainerFiles(context.Context, *ContainerPath) (*Files, error)
}

func RegisterRadarServer(s *grpc.Server, srv RadarServer) {
	s.RegisterService(&_Radar_serviceDesc, srv)
}

func _Radar_ListContainerFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadarServer).ListContainerFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.radar.Radar/ListContainerFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadarServer).ListContainerFiles(ctx, req.(*ContainerPath))
	}
	return interceptor(ctx, in, info, handler)
}

var _Radar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.radar.Radar",
	HandlerType: (*RadarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContainerFiles",
			Handler:    _Radar_ListContainerFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/radar.proto",
}

func init() { proto.RegisterFile("proto/radar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x8d, 0x6b, 0xe7, 0xf6, 0xea, 0x0e, 0x0b, 0x08, 0xa1, 0x1e, 0xac, 0x05, 0xa1, 0xa7,
	0x0c, 0x2a, 0x7e, 0x02, 0x65, 0x20, 0x88, 0x93, 0xe0, 0xbd, 0x64, 0x26, 0xdb, 0x02, 0x4d, 0x53,
	0x92, 0x78, 0xf1, 0x53, 0xf8, 0x91, 0x25, 0xa9, 0x1d, 0x0e, 0x4f, 0xfd, 0xe7, 0xf7, 0x1e, 0x7d,
	0xef, 0xf7, 0x60, 0xd9, 0x5b, 0xe3, 0xcd, 0xca, 0x72, 0xc1, 0x2d, 0x8d, 0x19, 0x67, 0xf1, 0x43,
	0x23, 0xca, 0x6f, 0xf6, 0xc6, 0xec, 0x5b, 0xb9, 0x8a, 0x6c, 0xfb, 0xb9, 0x5b, 0x79, 0xa5, 0xa5,
	0xf3, 0x5c, 0xf7, 0x43, 0x77, 0xb9, 0x81, 0xc5, 0xa3, 0xe9, 0x3c, 0x57, 0x9d, 0xb4, 0x6f, 0xdc,
	0x1f, 0xf0, 0x2d, 0x5c, 0x7e, 0x8c, 0xa0, 0x51, 0x82, 0xa0, 0x02, 0x55, 0x73, 0x96, 0x1d, 0xd9,
	0xb3, 0xc0, 0xd7, 0x30, 0xef, 0xb9, 0x3f, 0x34, 0x1d, 0xd7, 0x92, 0x9c, 0xc7, 0xfa, 0x2c, 0x80,
	0x57, 0xae, 0x65, 0x49, 0x21, 0x5d, 0xab, 0x56, 0x3a, 0x7c, 0x07, 0xc9, 0x4e, 0xb5, 0x92, 0xa0,
	0x62, 0x52, 0x65, 0xf5, 0x92, 0xfe, 0x59, 0x8b, 0x86, 0x0e, 0x16, 0xcb, 0xe5, 0x37, 0x82, 0x24,
	0x3c, 0x31, 0x86, 0x24, 0xfc, 0xe4, 0x77, 0x60, 0xcc, 0x81, 0x39, 0xf5, 0x35, 0x0c, 0x99, 0xb0,
	0x98, 0x03, 0xd3, 0x46, 0x48, 0x32, 0x29, 0x50, 0xb5, 0x60, 0x31, 0xe3, 0x07, 0x98, 0x69, 0x23,
	0x9a, 0x20, 0x47, 0x92, 0x02, 0x55, 0x59, 0x9d, 0xd3, 0xc1, 0x9c, 0x8e, 0xe6, 0xf4, 0x7d, 0x34,
	0x67, 0x17, 0xda, 0x88, 0xf0, 0xc2, 0x57, 0x30, 0x55, 0xae, 0x11, 0xca, 0x92, 0xb4, 0x40, 0xd5,
	0x8c, 0xa5, 0xca, 0x3d, 0x29, 0x5b, 0x6f, 0x20, 0x65, 0x61, 0x4d, 0xbc, 0x06, 0xfc, 0xa2, 0x9c,
	0x3f, 0x1e, 0x68, 0x10, 0xcb, 0x4f, 0x54, 0x4e, 0xae, 0x97, 0xe3, 0x7f, 0x9a, 0xae, 0x3c, 0xdb,
	0x4e, 0x23, 0xbc, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x98, 0xc2, 0xed, 0xc7, 0xae, 0x01, 0x00,
	0x00,
}

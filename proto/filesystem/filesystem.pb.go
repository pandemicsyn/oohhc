// Code generated by protoc-gen-go.
// source: filesystem.proto
// DO NOT EDIT!

/*
Package filesystem is a generated protocol buffer package.

It is generated from these files:
	filesystem.proto

It has these top-level messages:
	FileSystem
	IPAddress
	CreateFSRequest
	CreateFSResponse
	ListFSRequest
	ListFSResponse
	ShowFSRequest
	ShowFSResponse
	DeleteFSRequest
	DeleteFSResponse
	UpdateFSRequest
	UpdateFSResponse
	GrantAddrFSRequest
	GrantAddrFSResponse
	RevokeAddrFSRequest
	RevokeAddrFSResponse
	ValidateAddrFSRequest
	ValidateAddrFSResponse
*/
package filesystem

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
const _ = proto.ProtoPackageIsVersion1

// FileSystem Profile
type FileSystem struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	CreateDate int64  `protobuf:"varint,4,opt,name=createDate" json:"createDate,omitempty"`
	DeleteDate int64  `protobuf:"varint,5,opt,name=deleteDate" json:"deleteDate,omitempty"`
}

func (m *FileSystem) Reset()                    { *m = FileSystem{} }
func (m *FileSystem) String() string            { return proto.CompactTextString(m) }
func (*FileSystem) ProtoMessage()               {}
func (*FileSystem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// IP Address Profile
type IPAddress struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *IPAddress) Reset()                    { *m = IPAddress{} }
func (m *IPAddress) String() string            { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()               {}
func (*IPAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request to create a new filesystem
type CreateFSRequest struct {
	Acct    string `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys string `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Apikey  string `protobuf:"bytes,3,opt,name=apikey" json:"apikey,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
}

func (m *CreateFSRequest) Reset()                    { *m = CreateFSRequest{} }
func (m *CreateFSRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFSRequest) ProtoMessage()               {}
func (*CreateFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response from creating a new filesystem
type CreateFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CreateFSResponse) Reset()                    { *m = CreateFSResponse{} }
func (m *CreateFSResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateFSResponse) ProtoMessage()               {}
func (*CreateFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Request a list of all file systems for a
// given account
type ListFSRequest struct {
	Acct   string `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Apikey string `protobuf:"bytes,2,opt,name=apikey" json:"apikey,omitempty"`
}

func (m *ListFSRequest) Reset()                    { *m = ListFSRequest{} }
func (m *ListFSRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFSRequest) ProtoMessage()               {}
func (*ListFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Response from showing either an accounts specific
// file system or a list of all an accounts file systems.
type ListFSResponse struct {
	Filesystems *FileSystem `protobuf:"bytes,1,opt,name=filesystems" json:"filesystems,omitempty"`
	Status      string      `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ListFSResponse) Reset()                    { *m = ListFSResponse{} }
func (m *ListFSResponse) String() string            { return proto.CompactTextString(m) }
func (*ListFSResponse) ProtoMessage()               {}
func (*ListFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListFSResponse) GetFilesystems() *FileSystem {
	if m != nil {
		return m.Filesystems
	}
	return nil
}

// Request to show the specific details about a file system
type ShowFSRequest struct {
	Acct    string `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys string `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Apikey  string `protobuf:"bytes,3,opt,name=apikey" json:"apikey,omitempty"`
}

func (m *ShowFSRequest) Reset()                    { *m = ShowFSRequest{} }
func (m *ShowFSRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowFSRequest) ProtoMessage()               {}
func (*ShowFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Response from showing either an accounts specific
// file system or a list of all an accounts file systems.
type ShowFSResponse struct {
	Filesystems *FileSystem  `protobuf:"bytes,1,opt,name=filesystems" json:"filesystems,omitempty"`
	Ipaddresses []*IPAddress `protobuf:"bytes,2,rep,name=ipaddresses" json:"ipaddresses,omitempty"`
	Status      string       `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *ShowFSResponse) Reset()                    { *m = ShowFSResponse{} }
func (m *ShowFSResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowFSResponse) ProtoMessage()               {}
func (*ShowFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ShowFSResponse) GetFilesystems() *FileSystem {
	if m != nil {
		return m.Filesystems
	}
	return nil
}

func (m *ShowFSResponse) GetIpaddresses() []*IPAddress {
	if m != nil {
		return m.Ipaddresses
	}
	return nil
}

// Request to delete a specific file system
type DeleteFSRequest struct {
	Acct    string `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys string `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Apikey  string `protobuf:"bytes,3,opt,name=apikey" json:"apikey,omitempty"`
}

func (m *DeleteFSRequest) Reset()                    { *m = DeleteFSRequest{} }
func (m *DeleteFSRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteFSRequest) ProtoMessage()               {}
func (*DeleteFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Response from deleting a file system
type DeleteFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *DeleteFSResponse) Reset()                    { *m = DeleteFSResponse{} }
func (m *DeleteFSResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteFSResponse) ProtoMessage()               {}
func (*DeleteFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Request to update a specific file system's information
type UpdateFSRequest struct {
	Acct       string      `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys    string      `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Apikey     string      `protobuf:"bytes,3,opt,name=apikey" json:"apikey,omitempty"`
	Filesystem *FileSystem `protobuf:"bytes,4,opt,name=filesystem" json:"filesystem,omitempty"`
}

func (m *UpdateFSRequest) Reset()                    { *m = UpdateFSRequest{} }
func (m *UpdateFSRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateFSRequest) ProtoMessage()               {}
func (*UpdateFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateFSRequest) GetFilesystem() *FileSystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

// Response from an update operation
type UpdateFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *UpdateFSResponse) Reset()                    { *m = UpdateFSResponse{} }
func (m *UpdateFSResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateFSResponse) ProtoMessage()               {}
func (*UpdateFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// Request grant an ip address access to a file system
type GrantAddrFSRequest struct {
	Acct      string `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys   string `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Apikey    string `protobuf:"bytes,3,opt,name=apikey" json:"apikey,omitempty"`
	Ipaddress string `protobuf:"bytes,4,opt,name=ipaddress" json:"ipaddress,omitempty"`
}

func (m *GrantAddrFSRequest) Reset()                    { *m = GrantAddrFSRequest{} }
func (m *GrantAddrFSRequest) String() string            { return proto.CompactTextString(m) }
func (*GrantAddrFSRequest) ProtoMessage()               {}
func (*GrantAddrFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// Response from granting ip address access to a file system
type GrantAddrFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *GrantAddrFSResponse) Reset()                    { *m = GrantAddrFSResponse{} }
func (m *GrantAddrFSResponse) String() string            { return proto.CompactTextString(m) }
func (*GrantAddrFSResponse) ProtoMessage()               {}
func (*GrantAddrFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Request revoke an ip address access to a file system
type RevokeAddrFSRequest struct {
	Acct      string `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys   string `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Apikey    string `protobuf:"bytes,3,opt,name=apikey" json:"apikey,omitempty"`
	Ipaddress string `protobuf:"bytes,4,opt,name=ipaddress" json:"ipaddress,omitempty"`
}

func (m *RevokeAddrFSRequest) Reset()                    { *m = RevokeAddrFSRequest{} }
func (m *RevokeAddrFSRequest) String() string            { return proto.CompactTextString(m) }
func (*RevokeAddrFSRequest) ProtoMessage()               {}
func (*RevokeAddrFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// Response from revoking ip address access to a file system
type RevokeAddrFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *RevokeAddrFSResponse) Reset()                    { *m = RevokeAddrFSResponse{} }
func (m *RevokeAddrFSResponse) String() string            { return proto.CompactTextString(m) }
func (*RevokeAddrFSResponse) ProtoMessage()               {}
func (*RevokeAddrFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// Request to validate an addr has access to a file system
type ValidateAddrFSRequest struct {
	Acct    string     `protobuf:"bytes,1,opt,name=acct" json:"acct,omitempty"`
	Filesys string     `protobuf:"bytes,2,opt,name=filesys" json:"filesys,omitempty"`
	Addr    *IPAddress `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
}

func (m *ValidateAddrFSRequest) Reset()                    { *m = ValidateAddrFSRequest{} }
func (m *ValidateAddrFSRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateAddrFSRequest) ProtoMessage()               {}
func (*ValidateAddrFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ValidateAddrFSRequest) GetAddr() *IPAddress {
	if m != nil {
		return m.Addr
	}
	return nil
}

// Response from validating an addr has access to a file system
type ValidateAddrFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ValidateAddrFSResponse) Reset()                    { *m = ValidateAddrFSResponse{} }
func (m *ValidateAddrFSResponse) String() string            { return proto.CompactTextString(m) }
func (*ValidateAddrFSResponse) ProtoMessage()               {}
func (*ValidateAddrFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func init() {
	proto.RegisterType((*FileSystem)(nil), "filesystem.FileSystem")
	proto.RegisterType((*IPAddress)(nil), "filesystem.IPAddress")
	proto.RegisterType((*CreateFSRequest)(nil), "filesystem.CreateFSRequest")
	proto.RegisterType((*CreateFSResponse)(nil), "filesystem.CreateFSResponse")
	proto.RegisterType((*ListFSRequest)(nil), "filesystem.ListFSRequest")
	proto.RegisterType((*ListFSResponse)(nil), "filesystem.ListFSResponse")
	proto.RegisterType((*ShowFSRequest)(nil), "filesystem.ShowFSRequest")
	proto.RegisterType((*ShowFSResponse)(nil), "filesystem.ShowFSResponse")
	proto.RegisterType((*DeleteFSRequest)(nil), "filesystem.DeleteFSRequest")
	proto.RegisterType((*DeleteFSResponse)(nil), "filesystem.DeleteFSResponse")
	proto.RegisterType((*UpdateFSRequest)(nil), "filesystem.UpdateFSRequest")
	proto.RegisterType((*UpdateFSResponse)(nil), "filesystem.UpdateFSResponse")
	proto.RegisterType((*GrantAddrFSRequest)(nil), "filesystem.GrantAddrFSRequest")
	proto.RegisterType((*GrantAddrFSResponse)(nil), "filesystem.GrantAddrFSResponse")
	proto.RegisterType((*RevokeAddrFSRequest)(nil), "filesystem.RevokeAddrFSRequest")
	proto.RegisterType((*RevokeAddrFSResponse)(nil), "filesystem.RevokeAddrFSResponse")
	proto.RegisterType((*ValidateAddrFSRequest)(nil), "filesystem.ValidateAddrFSRequest")
	proto.RegisterType((*ValidateAddrFSResponse)(nil), "filesystem.ValidateAddrFSResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for FileSystemApi service

type FileSystemApiClient interface {
	CreateFS(ctx context.Context, in *CreateFSRequest, opts ...grpc.CallOption) (*CreateFSResponse, error)
	ListFS(ctx context.Context, in *ListFSRequest, opts ...grpc.CallOption) (*ListFSResponse, error)
	ShowFS(ctx context.Context, in *ShowFSRequest, opts ...grpc.CallOption) (*ShowFSResponse, error)
	DeleteFS(ctx context.Context, in *DeleteFSRequest, opts ...grpc.CallOption) (*DeleteFSResponse, error)
	UpdateFS(ctx context.Context, in *UpdateFSRequest, opts ...grpc.CallOption) (*UpdateFSResponse, error)
	GrantAddrFS(ctx context.Context, in *GrantAddrFSRequest, opts ...grpc.CallOption) (*GrantAddrFSResponse, error)
	RevokeAddrFS(ctx context.Context, in *RevokeAddrFSRequest, opts ...grpc.CallOption) (*RevokeAddrFSResponse, error)
	ValidiateAddrFS(ctx context.Context, in *ValidateAddrFSRequest, opts ...grpc.CallOption) (*ValidateAddrFSResponse, error)
}

type fileSystemApiClient struct {
	cc *grpc.ClientConn
}

func NewFileSystemApiClient(cc *grpc.ClientConn) FileSystemApiClient {
	return &fileSystemApiClient{cc}
}

func (c *fileSystemApiClient) CreateFS(ctx context.Context, in *CreateFSRequest, opts ...grpc.CallOption) (*CreateFSResponse, error) {
	out := new(CreateFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/CreateFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) ListFS(ctx context.Context, in *ListFSRequest, opts ...grpc.CallOption) (*ListFSResponse, error) {
	out := new(ListFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/ListFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) ShowFS(ctx context.Context, in *ShowFSRequest, opts ...grpc.CallOption) (*ShowFSResponse, error) {
	out := new(ShowFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/ShowFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) DeleteFS(ctx context.Context, in *DeleteFSRequest, opts ...grpc.CallOption) (*DeleteFSResponse, error) {
	out := new(DeleteFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/DeleteFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) UpdateFS(ctx context.Context, in *UpdateFSRequest, opts ...grpc.CallOption) (*UpdateFSResponse, error) {
	out := new(UpdateFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/UpdateFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) GrantAddrFS(ctx context.Context, in *GrantAddrFSRequest, opts ...grpc.CallOption) (*GrantAddrFSResponse, error) {
	out := new(GrantAddrFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/GrantAddrFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) RevokeAddrFS(ctx context.Context, in *RevokeAddrFSRequest, opts ...grpc.CallOption) (*RevokeAddrFSResponse, error) {
	out := new(RevokeAddrFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/RevokeAddrFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemApiClient) ValidiateAddrFS(ctx context.Context, in *ValidateAddrFSRequest, opts ...grpc.CallOption) (*ValidateAddrFSResponse, error) {
	out := new(ValidateAddrFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemApi/ValidiateAddrFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileSystemApi service

type FileSystemApiServer interface {
	CreateFS(context.Context, *CreateFSRequest) (*CreateFSResponse, error)
	ListFS(context.Context, *ListFSRequest) (*ListFSResponse, error)
	ShowFS(context.Context, *ShowFSRequest) (*ShowFSResponse, error)
	DeleteFS(context.Context, *DeleteFSRequest) (*DeleteFSResponse, error)
	UpdateFS(context.Context, *UpdateFSRequest) (*UpdateFSResponse, error)
	GrantAddrFS(context.Context, *GrantAddrFSRequest) (*GrantAddrFSResponse, error)
	RevokeAddrFS(context.Context, *RevokeAddrFSRequest) (*RevokeAddrFSResponse, error)
	ValidiateAddrFS(context.Context, *ValidateAddrFSRequest) (*ValidateAddrFSResponse, error)
}

func RegisterFileSystemApiServer(s *grpc.Server, srv FileSystemApiServer) {
	s.RegisterService(&_FileSystemApi_serviceDesc, srv)
}

func _FileSystemApi_CreateFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).CreateFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_ListFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).ListFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_ShowFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ShowFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).ShowFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_DeleteFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).DeleteFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_UpdateFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).UpdateFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_GrantAddrFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GrantAddrFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).GrantAddrFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_RevokeAddrFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RevokeAddrFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).RevokeAddrFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileSystemApi_ValidiateAddrFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ValidateAddrFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FileSystemApiServer).ValidiateAddrFS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _FileSystemApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filesystem.FileSystemApi",
	HandlerType: (*FileSystemApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFS",
			Handler:    _FileSystemApi_CreateFS_Handler,
		},
		{
			MethodName: "ListFS",
			Handler:    _FileSystemApi_ListFS_Handler,
		},
		{
			MethodName: "ShowFS",
			Handler:    _FileSystemApi_ShowFS_Handler,
		},
		{
			MethodName: "DeleteFS",
			Handler:    _FileSystemApi_DeleteFS_Handler,
		},
		{
			MethodName: "UpdateFS",
			Handler:    _FileSystemApi_UpdateFS_Handler,
		},
		{
			MethodName: "GrantAddrFS",
			Handler:    _FileSystemApi_GrantAddrFS_Handler,
		},
		{
			MethodName: "RevokeAddrFS",
			Handler:    _FileSystemApi_RevokeAddrFS_Handler,
		},
		{
			MethodName: "ValidiateAddrFS",
			Handler:    _FileSystemApi_ValidiateAddrFS_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x95, 0xe1, 0x6b, 0x13, 0x31,
	0x18, 0xc6, 0xdd, 0x5a, 0xab, 0x7d, 0xba, 0xf6, 0x66, 0xe6, 0x46, 0x77, 0x8a, 0xce, 0x88, 0x32,
	0x26, 0xee, 0xc3, 0xfc, 0x2e, 0x16, 0xc7, 0x86, 0xa0, 0x30, 0x57, 0x9c, 0xa0, 0x20, 0x9c, 0xbd,
	0x88, 0xc7, 0xba, 0xde, 0xd9, 0xa4, 0x4a, 0xff, 0x6e, 0xff, 0x01, 0x73, 0xc9, 0x5d, 0x93, 0x9c,
	0x17, 0xc1, 0xd3, 0x0f, 0x83, 0x25, 0x79, 0xf3, 0x7b, 0x9f, 0xbc, 0xef, 0xfb, 0x5c, 0xb1, 0xf9,
	0x25, 0x99, 0x32, 0xbe, 0xe4, 0x82, 0x5d, 0x1d, 0x66, 0xf3, 0x54, 0xa4, 0x04, 0x66, 0x87, 0x7e,
	0x02, 0x4e, 0xe4, 0x6a, 0xac, 0x56, 0x04, 0x58, 0x4f, 0xe2, 0xe1, 0xda, 0xde, 0xda, 0x7e, 0x97,
	0x6c, 0xa0, 0x3d, 0x8b, 0xae, 0xd8, 0x70, 0x5d, 0xad, 0x06, 0xe8, 0x70, 0x11, 0x89, 0x05, 0x1f,
	0xb6, 0xd4, 0x5a, 0x86, 0x4e, 0xe6, 0x2c, 0x12, 0xec, 0x58, 0xfe, 0x0d, 0xdb, 0x72, 0xaf, 0x95,
	0xef, 0xc5, 0x6c, 0xca, 0x8a, 0xbd, 0xeb, 0xf9, 0x1e, 0xdd, 0x45, 0xf7, 0xd5, 0xd9, 0x28, 0x8e,
	0xe7, 0x8c, 0xf3, 0x1c, 0x19, 0xc9, 0x7f, 0x75, 0x02, 0xfa, 0x16, 0xc1, 0x4b, 0x85, 0x38, 0x19,
	0x9f, 0xb3, 0x6f, 0x0b, 0xc6, 0x85, 0x0a, 0x98, 0x4c, 0x44, 0xa1, 0x20, 0xc0, 0x8d, 0x42, 0xa9,
	0x11, 0x11, 0x65, 0xc9, 0x25, 0x5b, 0x16, 0x22, 0x64, 0x40, 0xa4, 0xd1, 0x4a, 0x41, 0x97, 0x52,
	0x6c, 0x1a, 0x24, 0xcf, 0xd2, 0x19, 0x67, 0x96, 0x72, 0x9d, 0xf6, 0x29, 0xfa, 0xaf, 0x13, 0x2e,
	0x7c, 0x49, 0x4d, 0x0e, 0x95, 0x93, 0xbe, 0xc1, 0xa0, 0x0c, 0x2f, 0x80, 0x4f, 0xd0, 0x33, 0x05,
	0xd4, 0xd4, 0xde, 0xd1, 0xce, 0xa1, 0x55, 0x66, 0xab, 0xa2, 0x26, 0xbb, 0xc6, 0x3d, 0x47, 0x7f,
	0xfc, 0x35, 0xfd, 0xd1, 0xf4, 0xc9, 0x74, 0x89, 0x41, 0x79, 0xbf, 0x89, 0x9c, 0x03, 0xf4, 0x92,
	0xac, 0xa8, 0x19, 0xcb, 0x73, 0xb4, 0x64, 0xf0, 0xb6, 0x1d, 0x6c, 0xba, 0x55, 0x69, 0x39, 0x7d,
	0x81, 0xe0, 0x58, 0xb5, 0xb7, 0xb1, 0x78, 0xd9, 0x1e, 0x43, 0xf0, 0xb4, 0x67, 0x8a, 0xe0, 0x5d,
	0x16, 0xff, 0xcb, 0x54, 0x1c, 0xc0, 0x1a, 0x70, 0x35, 0x18, 0xde, 0x7a, 0xe4, 0x8a, 0x4c, 0x36,
	0x8f, 0xa2, 0x0b, 0x90, 0xd3, 0x79, 0x34, 0x13, 0x79, 0x5d, 0x1a, 0x8b, 0xba, 0x85, 0xee, 0xaa,
	0xf0, 0xc5, 0xb0, 0x3e, 0xc2, 0x96, 0xc3, 0xf5, 0xa4, 0x7f, 0x8f, 0xad, 0x73, 0xf6, 0x3d, 0xbd,
	0x64, 0xff, 0x3b, 0xff, 0x63, 0xdc, 0x76, 0xc1, 0x1e, 0x01, 0x1f, 0xb1, 0x7d, 0x11, 0x4d, 0x93,
	0xbc, 0x4a, 0x7f, 0x25, 0xe1, 0x61, 0xe1, 0xf6, 0x96, 0xea, 0x40, 0xfd, 0x90, 0xd1, 0x7d, 0xec,
	0x54, 0xe1, 0xf5, 0x32, 0x8e, 0x7e, 0xb6, 0xd1, 0x37, 0x9d, 0x1b, 0x65, 0x09, 0x39, 0xc5, 0xcd,
	0xd2, 0xed, 0xe4, 0x8e, 0x8d, 0xaf, 0x7c, 0x56, 0xc2, 0xbb, 0xf5, 0x87, 0x3a, 0x11, 0xbd, 0x46,
	0x46, 0xe8, 0x68, 0x8f, 0x93, 0x5d, 0x3b, 0xd2, 0xf9, 0x4c, 0x84, 0x61, 0xdd, 0x91, 0x8d, 0xd0,
	0xbe, 0x74, 0x11, 0x8e, 0xd7, 0x5d, 0x84, 0x6b, 0x63, 0x89, 0x90, 0xcf, 0x29, 0xdd, 0xe1, 0x3e,
	0xa7, 0xe2, 0x3a, 0xf7, 0x39, 0x55, 0x43, 0x69, 0x50, 0x39, 0xd4, 0x2e, 0xa8, 0x62, 0x2c, 0x17,
	0x54, 0xf5, 0x81, 0x04, 0x9d, 0xa1, 0x67, 0x4d, 0x28, 0xb9, 0x67, 0x87, 0xff, 0x6e, 0x89, 0xf0,
	0xbe, 0xf7, 0x7c, 0x45, 0x1c, 0x63, 0xc3, 0x9e, 0x39, 0xe2, 0x5c, 0xa9, 0x19, 0xf3, 0x70, 0xcf,
	0x1f, 0xb0, 0x82, 0x7e, 0x40, 0xa0, 0x66, 0x28, 0x59, 0x0d, 0x11, 0x79, 0x60, 0x5f, 0xab, 0x9d,
	0xde, 0x90, 0xfe, 0x29, 0xa4, 0x64, 0x7f, 0xee, 0xa8, 0x9f, 0xcc, 0x67, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xae, 0x09, 0xfd, 0xa7, 0x46, 0x07, 0x00, 0x00,
}
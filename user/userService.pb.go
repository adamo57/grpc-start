// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userService.proto

package user

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

type Username struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Username) Reset()         { *m = Username{} }
func (m *Username) String() string { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()    {}
func (*Username) Descriptor() ([]byte, []int) {
	return fileDescriptor_userService_fe1316d9b7634d36, []int{0}
}
func (m *Username) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Username.Unmarshal(m, b)
}
func (m *Username) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Username.Marshal(b, m, deterministic)
}
func (dst *Username) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Username.Merge(dst, src)
}
func (m *Username) XXX_Size() int {
	return xxx_messageInfo_Username.Size(m)
}
func (m *Username) XXX_DiscardUnknown() {
	xxx_messageInfo_Username.DiscardUnknown(m)
}

var xxx_messageInfo_Username proto.InternalMessageInfo

func (m *Username) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type DelUserMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelUserMessage) Reset()         { *m = DelUserMessage{} }
func (m *DelUserMessage) String() string { return proto.CompactTextString(m) }
func (*DelUserMessage) ProtoMessage()    {}
func (*DelUserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_userService_fe1316d9b7634d36, []int{1}
}
func (m *DelUserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelUserMessage.Unmarshal(m, b)
}
func (m *DelUserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelUserMessage.Marshal(b, m, deterministic)
}
func (dst *DelUserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelUserMessage.Merge(dst, src)
}
func (m *DelUserMessage) XXX_Size() int {
	return xxx_messageInfo_DelUserMessage.Size(m)
}
func (m *DelUserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DelUserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DelUserMessage proto.InternalMessageInfo

func (m *DelUserMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_userService_fe1316d9b7634d36, []int{2}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AddUserMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserMessage) Reset()         { *m = AddUserMessage{} }
func (m *AddUserMessage) String() string { return proto.CompactTextString(m) }
func (*AddUserMessage) ProtoMessage()    {}
func (*AddUserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_userService_fe1316d9b7634d36, []int{3}
}
func (m *AddUserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserMessage.Unmarshal(m, b)
}
func (m *AddUserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserMessage.Marshal(b, m, deterministic)
}
func (dst *AddUserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserMessage.Merge(dst, src)
}
func (m *AddUserMessage) XXX_Size() int {
	return xxx_messageInfo_AddUserMessage.Size(m)
}
func (m *AddUserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserMessage proto.InternalMessageInfo

func (m *AddUserMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Username)(nil), "user.Username")
	proto.RegisterType((*DelUserMessage)(nil), "user.DelUserMessage")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*AddUserMessage)(nil), "user.AddUserMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*AddUserMessage, error)
	DelUser(ctx context.Context, in *Username, opts ...grpc.CallOption) (*DelUserMessage, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*AddUserMessage, error) {
	out := new(AddUserMessage)
	err := c.cc.Invoke(ctx, "/user.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DelUser(ctx context.Context, in *Username, opts ...grpc.CallOption) (*DelUserMessage, error) {
	out := new(DelUserMessage)
	err := c.cc.Invoke(ctx, "/user.UserService/DelUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *UserInfo) (*AddUserMessage, error)
	DelUser(context.Context, *Username) (*DelUserMessage, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DelUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DelUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DelUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DelUser(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "DelUser",
			Handler:    _UserService_DelUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userService.proto",
}

func init() { proto.RegisterFile("userService.proto", fileDescriptor_userService_fe1316d9b7634d36) }

var fileDescriptor_userService_fe1316d9b7634d36 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2d, 0x4e, 0x2d,
	0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01,
	0x09, 0x29, 0xa9, 0x71, 0x71, 0x84, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x0a, 0x49, 0x71,
	0x71, 0x94, 0x42, 0xd9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x92, 0x16, 0x17,
	0x9f, 0x4b, 0x6a, 0x0e, 0x48, 0xa9, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x04, 0x17,
	0x7b, 0x2e, 0x84, 0x09, 0x55, 0x0c, 0xe3, 0x2a, 0xb9, 0x41, 0xcc, 0xf4, 0xcc, 0x4b, 0xcb, 0x17,
	0x12, 0xe2, 0x62, 0x41, 0x32, 0x0f, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0x06, 0xe9, 0x62, 0x52, 0x60,
	0xd4, 0x60, 0x0d, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x73, 0x13, 0x33, 0x73, 0x24, 0x98,
	0xc1, 0xca, 0x20, 0x1c, 0x90, 0x9d, 0x8e, 0x29, 0x29, 0x44, 0xd9, 0x69, 0x54, 0xcc, 0xc5, 0x1d,
	0x8a, 0xf0, 0xa2, 0x90, 0x21, 0x17, 0x3b, 0x54, 0xab, 0x10, 0x9f, 0x1e, 0xc8, 0x13, 0x7a, 0x30,
	0x17, 0x49, 0x89, 0x40, 0xf8, 0xa8, 0x26, 0x2b, 0x31, 0x80, 0xb4, 0x40, 0x7d, 0x88, 0xac, 0x05,
	0xe4, 0x60, 0x98, 0x16, 0xd4, 0x00, 0x50, 0x62, 0x70, 0x12, 0xe3, 0x62, 0xcf, 0xcc, 0x07, 0xcb,
	0x39, 0x71, 0x82, 0x64, 0x02, 0x40, 0x01, 0x1b, 0xc0, 0x98, 0xc4, 0x06, 0x0e, 0x61, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xff, 0x0a, 0xdc, 0x76, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow.proto

package workflow

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CreateRequest struct {
	Template             string   `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *CreateRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "workflow.Empty")
	proto.RegisterType((*CreateRequest)(nil), "workflow.CreateRequest")
}

func init() { proto.RegisterFile("workflow.proto", fileDescriptor_892c7f566756b0be) }

var fileDescriptor_892c7f566756b0be = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xcf, 0x2f, 0xca,
	0x4e, 0xcb, 0xc9, 0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xd8,
	0xb9, 0x58, 0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0x95, 0x9c, 0xb9, 0x78, 0x9d, 0x8b, 0x52, 0x13, 0x4b,
	0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0x52, 0x73, 0x0b,
	0x72, 0x12, 0x4b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x31, 0x2e,
	0xb6, 0x92, 0xc4, 0xa2, 0xf4, 0xd4, 0x12, 0x09, 0x26, 0xb0, 0x0c, 0x94, 0x67, 0xe4, 0xc6, 0xc5,
	0x11, 0x0e, 0x35, 0x59, 0xc8, 0x8a, 0x8b, 0x0f, 0x62, 0x60, 0x08, 0x4c, 0x97, 0xb8, 0x1e, 0xdc,
	0x19, 0x28, 0x56, 0x49, 0xf1, 0x23, 0x24, 0xc0, 0x8e, 0x49, 0x62, 0x03, 0x3b, 0xd3, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x79, 0xb5, 0x4b, 0xd5, 0xb8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowClient interface {
	CreateTemplate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
}

type workflowClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowClient(cc *grpc.ClientConn) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) CreateTemplate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/workflow.Workflow/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
type WorkflowServer interface {
	CreateTemplate(context.Context, *CreateRequest) (*Empty, error)
}

// UnimplementedWorkflowServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (*UnimplementedWorkflowServer) CreateTemplate(ctx context.Context, req *CreateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}

func RegisterWorkflowServer(s *grpc.Server, srv WorkflowServer) {
	s.RegisterService(&_Workflow_serviceDesc, srv)
}

func _Workflow_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.Workflow/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).CreateTemplate(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _Workflow_CreateTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow.proto",
}

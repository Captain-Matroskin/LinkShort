// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: link_share.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LinkShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkShort string `protobuf:"bytes,1,opt,name=link_short,json=linkShort,proto3" json:"link_short,omitempty"`
}

func (x *LinkShort) Reset() {
	*x = LinkShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkShort) ProtoMessage() {}

func (x *LinkShort) ProtoReflect() protoreflect.Message {
	mi := &file_link_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkShort.ProtoReflect.Descriptor instead.
func (*LinkShort) Descriptor() ([]byte, []int) {
	return file_link_share_proto_rawDescGZIP(), []int{0}
}

func (x *LinkShort) GetLinkShort() string {
	if x != nil {
		return x.LinkShort
	}
	return ""
}

type LinkFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkFull string `protobuf:"bytes,1,opt,name=link_full,json=linkFull,proto3" json:"link_full,omitempty"`
}

func (x *LinkFull) Reset() {
	*x = LinkFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkFull) ProtoMessage() {}

func (x *LinkFull) ProtoReflect() protoreflect.Message {
	mi := &file_link_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkFull.ProtoReflect.Descriptor instead.
func (*LinkFull) Descriptor() ([]byte, []int) {
	return file_link_share_proto_rawDescGZIP(), []int{1}
}

func (x *LinkFull) GetLinkFull() string {
	if x != nil {
		return x.LinkFull
	}
	return ""
}

type ResultLinkShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Body       *LinkShort `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Error      string     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResultLinkShort) Reset() {
	*x = ResultLinkShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultLinkShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultLinkShort) ProtoMessage() {}

func (x *ResultLinkShort) ProtoReflect() protoreflect.Message {
	mi := &file_link_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultLinkShort.ProtoReflect.Descriptor instead.
func (*ResultLinkShort) Descriptor() ([]byte, []int) {
	return file_link_share_proto_rawDescGZIP(), []int{2}
}

func (x *ResultLinkShort) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ResultLinkShort) GetBody() *LinkShort {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ResultLinkShort) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ResultLinkFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64     `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Body       *LinkFull `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Error      string    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResultLinkFull) Reset() {
	*x = ResultLinkFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultLinkFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultLinkFull) ProtoMessage() {}

func (x *ResultLinkFull) ProtoReflect() protoreflect.Message {
	mi := &file_link_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultLinkFull.ProtoReflect.Descriptor instead.
func (*ResultLinkFull) Descriptor() ([]byte, []int) {
	return file_link_share_proto_rawDescGZIP(), []int{3}
}

func (x *ResultLinkFull) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ResultLinkFull) GetBody() *LinkFull {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ResultLinkFull) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_link_share_proto protoreflect.FileDescriptor

var file_link_share_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x09, 0x4c, 0x69, 0x6e,
	0x6b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x6b,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x27, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x75, 0x6c,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x75, 0x6c, 0x6c, 0x22, 0x6e,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6c,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x75, 0x6c, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x75, 0x6c, 0x6c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x87, 0x01, 0x0a,
	0x10, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x46, 0x75, 0x6c, 0x6c, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x0a,
	0x0c, 0x54, 0x61, 0x6b, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x46, 0x75, 0x6c, 0x6c, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_share_proto_rawDescOnce sync.Once
	file_link_share_proto_rawDescData = file_link_share_proto_rawDesc
)

func file_link_share_proto_rawDescGZIP() []byte {
	file_link_share_proto_rawDescOnce.Do(func() {
		file_link_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_share_proto_rawDescData)
	})
	return file_link_share_proto_rawDescData
}

var file_link_share_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_link_share_proto_goTypes = []interface{}{
	(*LinkShort)(nil),       // 0: proto.LinkShort
	(*LinkFull)(nil),        // 1: proto.LinkFull
	(*ResultLinkShort)(nil), // 2: proto.ResultLinkShort
	(*ResultLinkFull)(nil),  // 3: proto.ResultLinkFull
}
var file_link_share_proto_depIdxs = []int32{
	0, // 0: proto.ResultLinkShort.body:type_name -> proto.LinkShort
	1, // 1: proto.ResultLinkFull.body:type_name -> proto.LinkFull
	1, // 2: proto.LinkShortService.CreateLinkShort:input_type -> proto.LinkFull
	0, // 3: proto.LinkShortService.TakeLinkFull:input_type -> proto.LinkShort
	2, // 4: proto.LinkShortService.CreateLinkShort:output_type -> proto.ResultLinkShort
	3, // 5: proto.LinkShortService.TakeLinkFull:output_type -> proto.ResultLinkFull
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_link_share_proto_init() }
func file_link_share_proto_init() {
	if File_link_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_link_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkShort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_link_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkFull); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_link_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultLinkShort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_link_share_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultLinkFull); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_link_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_share_proto_goTypes,
		DependencyIndexes: file_link_share_proto_depIdxs,
		MessageInfos:      file_link_share_proto_msgTypes,
	}.Build()
	File_link_share_proto = out.File
	file_link_share_proto_rawDesc = nil
	file_link_share_proto_goTypes = nil
	file_link_share_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LinkShortServiceClient is the client API for LinkShortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinkShortServiceClient interface {
	CreateLinkShort(ctx context.Context, in *LinkFull, opts ...grpc.CallOption) (*ResultLinkShort, error)
	TakeLinkFull(ctx context.Context, in *LinkShort, opts ...grpc.CallOption) (*ResultLinkFull, error)
}

type linkShortServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkShortServiceClient(cc grpc.ClientConnInterface) LinkShortServiceClient {
	return &linkShortServiceClient{cc}
}

func (c *linkShortServiceClient) CreateLinkShort(ctx context.Context, in *LinkFull, opts ...grpc.CallOption) (*ResultLinkShort, error) {
	out := new(ResultLinkShort)
	err := c.cc.Invoke(ctx, "/proto.LinkShortService/CreateLinkShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkShortServiceClient) TakeLinkFull(ctx context.Context, in *LinkShort, opts ...grpc.CallOption) (*ResultLinkFull, error) {
	out := new(ResultLinkFull)
	err := c.cc.Invoke(ctx, "/proto.LinkShortService/TakeLinkFull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkShortServiceServer is the server API for LinkShortService service.
type LinkShortServiceServer interface {
	CreateLinkShort(context.Context, *LinkFull) (*ResultLinkShort, error)
	TakeLinkFull(context.Context, *LinkShort) (*ResultLinkFull, error)
}

// UnimplementedLinkShortServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLinkShortServiceServer struct {
}

func (*UnimplementedLinkShortServiceServer) CreateLinkShort(context.Context, *LinkFull) (*ResultLinkShort, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLinkShort not implemented")
}
func (*UnimplementedLinkShortServiceServer) TakeLinkFull(context.Context, *LinkShort) (*ResultLinkFull, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeLinkFull not implemented")
}

func RegisterLinkShortServiceServer(s *grpc.Server, srv LinkShortServiceServer) {
	s.RegisterService(&_LinkShortService_serviceDesc, srv)
}

func _LinkShortService_CreateLinkShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkFull)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShortServiceServer).CreateLinkShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LinkShortService/CreateLinkShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShortServiceServer).CreateLinkShort(ctx, req.(*LinkFull))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkShortService_TakeLinkFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkShort)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShortServiceServer).TakeLinkFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LinkShortService/TakeLinkFull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShortServiceServer).TakeLinkFull(ctx, req.(*LinkShort))
	}
	return interceptor(ctx, in, info, handler)
}

var _LinkShortService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LinkShortService",
	HandlerType: (*LinkShortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLinkShort",
			Handler:    _LinkShortService_CreateLinkShort_Handler,
		},
		{
			MethodName: "TakeLinkFull",
			Handler:    _LinkShortService_TakeLinkFull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link_share.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: simple.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Simplerequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add1  int32  `protobuf:"varint,1,opt,name=add1,proto3" json:"add1,omitempty"`
	Add2  int32  `protobuf:"varint,2,opt,name=add2,proto3" json:"add2,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Simplerequest) Reset() {
	*x = Simplerequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simplerequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simplerequest) ProtoMessage() {}

func (x *Simplerequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simplerequest.ProtoReflect.Descriptor instead.
func (*Simplerequest) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{0}
}

func (x *Simplerequest) GetAdd1() int32 {
	if x != nil {
		return x.Add1
	}
	return 0
}

func (x *Simplerequest) GetAdd2() int32 {
	if x != nil {
		return x.Add2
	}
	return 0
}

func (x *Simplerequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Simpleresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum   int32  `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Simpleresponse) Reset() {
	*x = Simpleresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simpleresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simpleresponse) ProtoMessage() {}

func (x *Simpleresponse) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simpleresponse.ProtoReflect.Descriptor instead.
func (*Simpleresponse) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{1}
}

func (x *Simpleresponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Simpleresponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_simple_proto protoreflect.FileDescriptor

var file_simple_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x4d, 0x0a, 0x0d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x64, 0x64, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x64, 0x64, 0x32, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x40, 0x0a,
	0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simple_proto_rawDescOnce sync.Once
	file_simple_proto_rawDescData = file_simple_proto_rawDesc
)

func file_simple_proto_rawDescGZIP() []byte {
	file_simple_proto_rawDescOnce.Do(func() {
		file_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_simple_proto_rawDescData)
	})
	return file_simple_proto_rawDescData
}

var file_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_simple_proto_goTypes = []interface{}{
	(*Simplerequest)(nil),  // 0: main.simplerequest
	(*Simpleresponse)(nil), // 1: main.simpleresponse
}
var file_simple_proto_depIdxs = []int32{
	0, // 0: main.simple.getinfo:input_type -> main.simplerequest
	1, // 1: main.simple.getinfo:output_type -> main.simpleresponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_simple_proto_init() }
func file_simple_proto_init() {
	if File_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simplerequest); i {
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
		file_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simpleresponse); i {
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
			RawDescriptor: file_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simple_proto_goTypes,
		DependencyIndexes: file_simple_proto_depIdxs,
		MessageInfos:      file_simple_proto_msgTypes,
	}.Build()
	File_simple_proto = out.File
	file_simple_proto_rawDesc = nil
	file_simple_proto_goTypes = nil
	file_simple_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleClient interface {
	Getinfo(ctx context.Context, in *Simplerequest, opts ...grpc.CallOption) (*Simpleresponse, error)
}

type simpleClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleClient(cc grpc.ClientConnInterface) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) Getinfo(ctx context.Context, in *Simplerequest, opts ...grpc.CallOption) (*Simpleresponse, error) {
	out := new(Simpleresponse)
	err := c.cc.Invoke(ctx, "/main.simple/getinfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
type SimpleServer interface {
	Getinfo(context.Context, *Simplerequest) (*Simpleresponse, error)
}

// UnimplementedSimpleServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (*UnimplementedSimpleServer) Getinfo(context.Context, *Simplerequest) (*Simpleresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getinfo not implemented")
}

func RegisterSimpleServer(s *grpc.Server, srv SimpleServer) {
	s.RegisterService(&_Simple_serviceDesc, srv)
}

func _Simple_Getinfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simplerequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).Getinfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.simple/Getinfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).Getinfo(ctx, req.(*Simplerequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simple_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getinfo",
			Handler:    _Simple_Getinfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple.proto",
}

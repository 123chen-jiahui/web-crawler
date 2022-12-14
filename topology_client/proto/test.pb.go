// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: test.proto

package proto

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

type Destination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []int32 `protobuf:"varint,1,rep,packed,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *Destination) Reset() {
	*x = Destination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Destination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Destination) ProtoMessage() {}

func (x *Destination) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destination.ProtoReflect.Descriptor instead.
func (*Destination) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Destination) GetNodes() []int32 {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type TopologyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaveLocation string                 `protobuf:"bytes,1,opt,name=SaveLocation,proto3" json:"SaveLocation,omitempty"`
	Total        int32                  `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Contents     []string               `protobuf:"bytes,3,rep,name=Contents,proto3" json:"Contents,omitempty"`
	G            map[int32]*Destination `protobuf:"bytes,4,rep,name=G,proto3" json:"G,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TopologyRequest) Reset() {
	*x = TopologyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyRequest) ProtoMessage() {}

func (x *TopologyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyRequest.ProtoReflect.Descriptor instead.
func (*TopologyRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *TopologyRequest) GetSaveLocation() string {
	if x != nil {
		return x.SaveLocation
	}
	return ""
}

func (x *TopologyRequest) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TopologyRequest) GetContents() []string {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *TopologyRequest) GetG() map[int32]*Destination {
	if x != nil {
		return x.G
	}
	return nil
}

type TopologyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok       bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	FilePath string `protobuf:"bytes,2,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
}

func (x *TopologyReply) Reset() {
	*x = TopologyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyReply) ProtoMessage() {}

func (x *TopologyReply) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyReply.ProtoReflect.Descriptor instead.
func (*TopologyReply) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *TopologyReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *TopologyReply) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x61, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2b, 0x0a, 0x01, 0x47, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x01, 0x47, 0x1a,
	0x48, 0x0a, 0x06, 0x47, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x0d, 0x54, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x32, 0x42, 0x0a, 0x04, 0x44, 0x72, 0x61, 0x77, 0x12, 0x3a,
	0x0a, 0x08, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_proto_goTypes = []interface{}{
	(*Destination)(nil),     // 0: proto.Destination
	(*TopologyRequest)(nil), // 1: proto.TopologyRequest
	(*TopologyReply)(nil),   // 2: proto.TopologyReply
	nil,                     // 3: proto.TopologyRequest.GEntry
}
var file_test_proto_depIdxs = []int32{
	3, // 0: proto.TopologyRequest.G:type_name -> proto.TopologyRequest.GEntry
	0, // 1: proto.TopologyRequest.GEntry.value:type_name -> proto.Destination
	1, // 2: proto.Draw.Topology:input_type -> proto.TopologyRequest
	2, // 3: proto.Draw.Topology:output_type -> proto.TopologyReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Destination); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopologyRequest); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopologyReply); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DrawClient is the client API for Draw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DrawClient interface {
	Topology(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*TopologyReply, error)
}

type drawClient struct {
	cc grpc.ClientConnInterface
}

func NewDrawClient(cc grpc.ClientConnInterface) DrawClient {
	return &drawClient{cc}
}

func (c *drawClient) Topology(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*TopologyReply, error) {
	out := new(TopologyReply)
	err := c.cc.Invoke(ctx, "/proto.Draw/Topology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DrawServer is the server API for Draw service.
type DrawServer interface {
	Topology(context.Context, *TopologyRequest) (*TopologyReply, error)
}

// UnimplementedDrawServer can be embedded to have forward compatible implementations.
type UnimplementedDrawServer struct {
}

func (*UnimplementedDrawServer) Topology(context.Context, *TopologyRequest) (*TopologyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topology not implemented")
}

func RegisterDrawServer(s *grpc.Server, srv DrawServer) {
	s.RegisterService(&_Draw_serviceDesc, srv)
}

func _Draw_Topology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DrawServer).Topology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Draw/Topology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DrawServer).Topology(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Draw_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Draw",
	HandlerType: (*DrawServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topology",
			Handler:    _Draw_Topology_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

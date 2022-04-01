// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: gateway/middleware/cors/v1/cors.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Cors middleware config.
type Cors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default is true
	AllowCredentials bool     `protobuf:"varint,1,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
	AllowedOrigins   []string `protobuf:"bytes,2,rep,name=allowed_origins,json=allowedOrigins,proto3" json:"allowed_origins,omitempty"`
	// default is GET/POST
	AllowedMethods []string `protobuf:"bytes,3,rep,name=allowed_methods,json=allowedMethods,proto3" json:"allowed_methods,omitempty"`
	// default is Origin/Content-Type/Content-Length
	AllowedHeaders []string `protobuf:"bytes,4,rep,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	ExposedHeaders []string `protobuf:"bytes,5,rep,name=exposed_headers,json=exposedHeaders,proto3" json:"exposed_headers,omitempty"`
	// default is 10 minutes
	MaxAge *durationpb.Duration `protobuf:"bytes,6,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
}

func (x *Cors) Reset() {
	*x = Cors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_middleware_cors_v1_cors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cors) ProtoMessage() {}

func (x *Cors) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_middleware_cors_v1_cors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cors.ProtoReflect.Descriptor instead.
func (*Cors) Descriptor() ([]byte, []int) {
	return file_gateway_middleware_cors_v1_cors_proto_rawDescGZIP(), []int{0}
}

func (x *Cors) GetAllowCredentials() bool {
	if x != nil {
		return x.AllowCredentials
	}
	return false
}

func (x *Cors) GetAllowedOrigins() []string {
	if x != nil {
		return x.AllowedOrigins
	}
	return nil
}

func (x *Cors) GetAllowedMethods() []string {
	if x != nil {
		return x.AllowedMethods
	}
	return nil
}

func (x *Cors) GetAllowedHeaders() []string {
	if x != nil {
		return x.AllowedHeaders
	}
	return nil
}

func (x *Cors) GetExposedHeaders() []string {
	if x != nil {
		return x.ExposedHeaders
	}
	return nil
}

func (x *Cors) GetMaxAge() *durationpb.Duration {
	if x != nil {
		return x.MaxAge
	}
	return nil
}

var File_gateway_middleware_cors_v1_cors_proto protoreflect.FileDescriptor

var file_gateway_middleware_cors_v1_cors_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65,
	0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67,
	0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_middleware_cors_v1_cors_proto_rawDescOnce sync.Once
	file_gateway_middleware_cors_v1_cors_proto_rawDescData = file_gateway_middleware_cors_v1_cors_proto_rawDesc
)

func file_gateway_middleware_cors_v1_cors_proto_rawDescGZIP() []byte {
	file_gateway_middleware_cors_v1_cors_proto_rawDescOnce.Do(func() {
		file_gateway_middleware_cors_v1_cors_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_middleware_cors_v1_cors_proto_rawDescData)
	})
	return file_gateway_middleware_cors_v1_cors_proto_rawDescData
}

var file_gateway_middleware_cors_v1_cors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_middleware_cors_v1_cors_proto_goTypes = []interface{}{
	(*Cors)(nil),                // 0: gateway.middleware.cors.v1.Cors
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_gateway_middleware_cors_v1_cors_proto_depIdxs = []int32{
	1, // 0: gateway.middleware.cors.v1.Cors.max_age:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gateway_middleware_cors_v1_cors_proto_init() }
func file_gateway_middleware_cors_v1_cors_proto_init() {
	if File_gateway_middleware_cors_v1_cors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_middleware_cors_v1_cors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cors); i {
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
			RawDescriptor: file_gateway_middleware_cors_v1_cors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_middleware_cors_v1_cors_proto_goTypes,
		DependencyIndexes: file_gateway_middleware_cors_v1_cors_proto_depIdxs,
		MessageInfos:      file_gateway_middleware_cors_v1_cors_proto_msgTypes,
	}.Build()
	File_gateway_middleware_cors_v1_cors_proto = out.File
	file_gateway_middleware_cors_v1_cors_proto_rawDesc = nil
	file_gateway_middleware_cors_v1_cors_proto_goTypes = nil
	file_gateway_middleware_cors_v1_cors_proto_depIdxs = nil
}

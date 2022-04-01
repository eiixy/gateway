// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: gateway/middleware/circuitbreaker/v1/circuitbreaker.proto

package v1

import (
	v1 "github.com/go-kratos/gateway/api/gateway/config/v1"
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

// CircuitBreaker middleware config.
type CircuitBreaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Trigger:
	//	*CircuitBreaker_SuccessRatio
	//	*CircuitBreaker_Ratio
	Trigger isCircuitBreaker_Trigger `protobuf_oneof:"trigger"`
	// Types that are assignable to Action:
	//	*CircuitBreaker_ResponseData
	//	*CircuitBreaker_BackupService
	Action          isCircuitBreaker_Action `protobuf_oneof:"action"`
	AssertCondtions []*v1.Condition         `protobuf:"bytes,5,rep,name=assert_condtions,json=assertCondtions,proto3" json:"assert_condtions,omitempty"`
}

func (x *CircuitBreaker) Reset() {
	*x = CircuitBreaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircuitBreaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircuitBreaker) ProtoMessage() {}

func (x *CircuitBreaker) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircuitBreaker.ProtoReflect.Descriptor instead.
func (*CircuitBreaker) Descriptor() ([]byte, []int) {
	return file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescGZIP(), []int{0}
}

func (m *CircuitBreaker) GetTrigger() isCircuitBreaker_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (x *CircuitBreaker) GetSuccessRatio() *SuccessRatio {
	if x, ok := x.GetTrigger().(*CircuitBreaker_SuccessRatio); ok {
		return x.SuccessRatio
	}
	return nil
}

func (x *CircuitBreaker) GetRatio() int64 {
	if x, ok := x.GetTrigger().(*CircuitBreaker_Ratio); ok {
		return x.Ratio
	}
	return 0
}

func (m *CircuitBreaker) GetAction() isCircuitBreaker_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *CircuitBreaker) GetResponseData() *ResponseData {
	if x, ok := x.GetAction().(*CircuitBreaker_ResponseData); ok {
		return x.ResponseData
	}
	return nil
}

func (x *CircuitBreaker) GetBackupService() *BackupService {
	if x, ok := x.GetAction().(*CircuitBreaker_BackupService); ok {
		return x.BackupService
	}
	return nil
}

func (x *CircuitBreaker) GetAssertCondtions() []*v1.Condition {
	if x != nil {
		return x.AssertCondtions
	}
	return nil
}

type isCircuitBreaker_Trigger interface {
	isCircuitBreaker_Trigger()
}

type CircuitBreaker_SuccessRatio struct {
	SuccessRatio *SuccessRatio `protobuf:"bytes,1,opt,name=success_ratio,json=successRatio,proto3,oneof"`
}

type CircuitBreaker_Ratio struct {
	Ratio int64 `protobuf:"varint,2,opt,name=ratio,proto3,oneof"`
}

func (*CircuitBreaker_SuccessRatio) isCircuitBreaker_Trigger() {}

func (*CircuitBreaker_Ratio) isCircuitBreaker_Trigger() {}

type isCircuitBreaker_Action interface {
	isCircuitBreaker_Action()
}

type CircuitBreaker_ResponseData struct {
	ResponseData *ResponseData `protobuf:"bytes,3,opt,name=response_data,json=responseData,proto3,oneof"`
}

type CircuitBreaker_BackupService struct {
	BackupService *BackupService `protobuf:"bytes,4,opt,name=backup_service,json=backupService,proto3,oneof"`
}

func (*CircuitBreaker_ResponseData) isCircuitBreaker_Action() {}

func (*CircuitBreaker_BackupService) isCircuitBreaker_Action() {}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32     `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Header     []*Header `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty"`
	Body       []byte    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseData) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ResponseData) GetHeader() []*Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ResponseData) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type BackupService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint *v1.Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *BackupService) Reset() {
	*x = BackupService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupService) ProtoMessage() {}

func (x *BackupService) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupService.ProtoReflect.Descriptor instead.
func (*BackupService) Descriptor() ([]byte, []int) {
	return file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescGZIP(), []int{3}
}

func (x *BackupService) GetEndpoint() *v1.Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

type SuccessRatio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success float64              `protobuf:"fixed64,1,opt,name=success,proto3" json:"success,omitempty"`
	Request int64                `protobuf:"varint,2,opt,name=request,proto3" json:"request,omitempty"`
	Bucket  int64                `protobuf:"varint,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Window  *durationpb.Duration `protobuf:"bytes,4,opt,name=window,proto3" json:"window,omitempty"`
}

func (x *SuccessRatio) Reset() {
	*x = SuccessRatio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessRatio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessRatio) ProtoMessage() {}

func (x *SuccessRatio) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessRatio.ProtoReflect.Descriptor instead.
func (*SuccessRatio) Descriptor() ([]byte, []int) {
	return file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescGZIP(), []int{4}
}

func (x *SuccessRatio) GetSuccess() float64 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *SuccessRatio) GetRequest() int64 {
	if x != nil {
		return x.Request
	}
	return 0
}

func (x *SuccessRatio) GetBucket() int64 {
	if x != nil {
		return x.Bucket
	}
	return 0
}

func (x *SuccessRatio) GetWindow() *durationpb.Duration {
	if x != nil {
		return x.Window
	}
	return nil
}

var File_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto protoreflect.FileDescriptor

var file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x0e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x48, 0x00, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x16, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x48, 0x01, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x10, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x30, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x48, 0x0a,
	0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x63,
	0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescOnce sync.Once
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescData = file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDesc
)

func file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescGZIP() []byte {
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescOnce.Do(func() {
		file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescData)
	})
	return file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDescData
}

var file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_goTypes = []interface{}{
	(*CircuitBreaker)(nil),      // 0: gateway.middleware.circuitbreaker.v1.CircuitBreaker
	(*Header)(nil),              // 1: gateway.middleware.circuitbreaker.v1.Header
	(*ResponseData)(nil),        // 2: gateway.middleware.circuitbreaker.v1.ResponseData
	(*BackupService)(nil),       // 3: gateway.middleware.circuitbreaker.v1.BackupService
	(*SuccessRatio)(nil),        // 4: gateway.middleware.circuitbreaker.v1.SuccessRatio
	(*v1.Condition)(nil),        // 5: gateway.config.v1.Condition
	(*v1.Endpoint)(nil),         // 6: gateway.config.v1.Endpoint
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
}
var file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_depIdxs = []int32{
	4, // 0: gateway.middleware.circuitbreaker.v1.CircuitBreaker.success_ratio:type_name -> gateway.middleware.circuitbreaker.v1.SuccessRatio
	2, // 1: gateway.middleware.circuitbreaker.v1.CircuitBreaker.response_data:type_name -> gateway.middleware.circuitbreaker.v1.ResponseData
	3, // 2: gateway.middleware.circuitbreaker.v1.CircuitBreaker.backup_service:type_name -> gateway.middleware.circuitbreaker.v1.BackupService
	5, // 3: gateway.middleware.circuitbreaker.v1.CircuitBreaker.assert_condtions:type_name -> gateway.config.v1.Condition
	1, // 4: gateway.middleware.circuitbreaker.v1.ResponseData.header:type_name -> gateway.middleware.circuitbreaker.v1.Header
	6, // 5: gateway.middleware.circuitbreaker.v1.BackupService.endpoint:type_name -> gateway.config.v1.Endpoint
	7, // 6: gateway.middleware.circuitbreaker.v1.SuccessRatio.window:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_init() }
func file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_init() {
	if File_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CircuitBreaker); i {
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
		file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
		file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupService); i {
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
		file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessRatio); i {
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
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CircuitBreaker_SuccessRatio)(nil),
		(*CircuitBreaker_Ratio)(nil),
		(*CircuitBreaker_ResponseData)(nil),
		(*CircuitBreaker_BackupService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_goTypes,
		DependencyIndexes: file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_depIdxs,
		MessageInfos:      file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_msgTypes,
	}.Build()
	File_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto = out.File
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_rawDesc = nil
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_goTypes = nil
	file_gateway_middleware_circuitbreaker_v1_circuitbreaker_proto_depIdxs = nil
}

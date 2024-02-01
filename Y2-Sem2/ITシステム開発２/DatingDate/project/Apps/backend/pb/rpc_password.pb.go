// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: rpc_password.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InputPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *InputPasswordRequest) Reset() {
	*x = InputPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPasswordRequest) ProtoMessage() {}

func (x *InputPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPasswordRequest.ProtoReflect.Descriptor instead.
func (*InputPasswordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_password_proto_rawDescGZIP(), []int{0}
}

func (x *InputPasswordRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *InputPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type InputPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message           string                 `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=PasswordChangedAt,proto3" json:"PasswordChangedAt,omitempty"`
}

func (x *InputPasswordResponse) Reset() {
	*x = InputPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPasswordResponse) ProtoMessage() {}

func (x *InputPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPasswordResponse.ProtoReflect.Descriptor instead.
func (*InputPasswordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_password_proto_rawDescGZIP(), []int{1}
}

func (x *InputPasswordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InputPasswordResponse) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_password_proto_rawDescGZIP(), []int{2}
}

func (x *ResetPasswordRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *ResetPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResetPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message           string                 `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=PasswordChangedAt,proto3" json:"PasswordChangedAt,omitempty"`
}

func (x *ResetPasswordResponse) Reset() {
	*x = ResetPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_password_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordResponse) ProtoMessage() {}

func (x *ResetPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_password_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordResponse.ProtoReflect.Descriptor instead.
func (*ResetPasswordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_password_proto_rawDescGZIP(), []int{3}
}

func (x *ResetPasswordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResetPasswordResponse) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

var File_rpc_password_proto protoreflect.FileDescriptor

var file_rpc_password_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x14, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7b, 0x0a, 0x15, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x48,
	0x0a, 0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x22, 0x50, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7b, 0x0a, 0x15, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x48, 0x0a,
	0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_password_proto_rawDescOnce sync.Once
	file_rpc_password_proto_rawDescData = file_rpc_password_proto_rawDesc
)

func file_rpc_password_proto_rawDescGZIP() []byte {
	file_rpc_password_proto_rawDescOnce.Do(func() {
		file_rpc_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_password_proto_rawDescData)
	})
	return file_rpc_password_proto_rawDescData
}

var file_rpc_password_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_password_proto_goTypes = []interface{}{
	(*InputPasswordRequest)(nil),  // 0: pb.InputPasswordRequest
	(*InputPasswordResponse)(nil), // 1: pb.InputPasswordResponse
	(*ResetPasswordRequest)(nil),  // 2: pb.ResetPasswordRequest
	(*ResetPasswordResponse)(nil), // 3: pb.ResetPasswordResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_rpc_password_proto_depIdxs = []int32{
	4, // 0: pb.InputPasswordResponse.PasswordChangedAt:type_name -> google.protobuf.Timestamp
	4, // 1: pb.ResetPasswordResponse.PasswordChangedAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_password_proto_init() }
func file_rpc_password_proto_init() {
	if File_rpc_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputPasswordRequest); i {
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
		file_rpc_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputPasswordResponse); i {
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
		file_rpc_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_rpc_password_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordResponse); i {
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
			RawDescriptor: file_rpc_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_password_proto_goTypes,
		DependencyIndexes: file_rpc_password_proto_depIdxs,
		MessageInfos:      file_rpc_password_proto_msgTypes,
	}.Build()
	File_rpc_password_proto = out.File
	file_rpc_password_proto_rawDesc = nil
	file_rpc_password_proto_goTypes = nil
	file_rpc_password_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: file.proto

package file

import (
	context "context"
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

type NewMultiUploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash         string `protobuf:"bytes,1,opt,name=FileHash,proto3" json:"FileHash,omitempty"`                  // 文件内容的唯一哈希值
	ChunkTotalNumber int64  `protobuf:"varint,2,opt,name=ChunkTotalNumber,proto3" json:"ChunkTotalNumber,omitempty"` // 分块总数
	FileSize         int64  `protobuf:"varint,3,opt,name=FileSize,proto3" json:"FileSize,omitempty"`                 // 分块大小
	FileName         string `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`                  // 文件名
	UserID           int64  `protobuf:"varint,5,opt,name=UserID,proto3" json:"UserID,omitempty"`                     // 用户ID
	FileType         int32  `protobuf:"varint,6,opt,name=FileType,proto3" json:"FileType,omitempty"`                 // 文件类型，文件属性比如图片，视频等
}

func (x *NewMultiUploadReq) Reset() {
	*x = NewMultiUploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMultiUploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMultiUploadReq) ProtoMessage() {}

func (x *NewMultiUploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMultiUploadReq.ProtoReflect.Descriptor instead.
func (*NewMultiUploadReq) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *NewMultiUploadReq) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *NewMultiUploadReq) GetChunkTotalNumber() int64 {
	if x != nil {
		return x.ChunkTotalNumber
	}
	return 0
}

func (x *NewMultiUploadReq) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *NewMultiUploadReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *NewMultiUploadReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *NewMultiUploadReq) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

type NewMultiUploadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewMultiUploadResp) Reset() {
	*x = NewMultiUploadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMultiUploadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMultiUploadResp) ProtoMessage() {}

func (x *NewMultiUploadResp) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMultiUploadResp.ProtoReflect.Descriptor instead.
func (*NewMultiUploadResp) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

type GetMultiUploadUriReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash  string `protobuf:"bytes,1,opt,name=FileHash,proto3" json:"FileHash,omitempty"`
	UserID    int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`       // 用户ID
	ChunkID   int64  `protobuf:"varint,3,opt,name=ChunkID,proto3" json:"ChunkID,omitempty"`     // 分块ID
	ChunkSize int64  `protobuf:"varint,4,opt,name=ChunkSize,proto3" json:"ChunkSize,omitempty"` // 分块大小
}

func (x *GetMultiUploadUriReq) Reset() {
	*x = GetMultiUploadUriReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultiUploadUriReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultiUploadUriReq) ProtoMessage() {}

func (x *GetMultiUploadUriReq) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultiUploadUriReq.ProtoReflect.Descriptor instead.
func (*GetMultiUploadUriReq) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *GetMultiUploadUriReq) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *GetMultiUploadUriReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetMultiUploadUriReq) GetChunkID() int64 {
	if x != nil {
		return x.ChunkID
	}
	return 0
}

func (x *GetMultiUploadUriReq) GetChunkSize() int64 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

type GetMultiUploadUriResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=Uri,proto3" json:"Uri,omitempty"`
}

func (x *GetMultiUploadUriResp) Reset() {
	*x = GetMultiUploadUriResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultiUploadUriResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultiUploadUriResp) ProtoMessage() {}

func (x *GetMultiUploadUriResp) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultiUploadUriResp.ProtoReflect.Descriptor instead.
func (*GetMultiUploadUriResp) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *GetMultiUploadUriResp) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type CompleteMultipartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash string `protobuf:"bytes,1,opt,name=FileHash,proto3" json:"FileHash,omitempty"`
	UserID   int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *CompleteMultipartReq) Reset() {
	*x = CompleteMultipartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteMultipartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteMultipartReq) ProtoMessage() {}

func (x *CompleteMultipartReq) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteMultipartReq.ProtoReflect.Descriptor instead.
func (*CompleteMultipartReq) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *CompleteMultipartReq) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *CompleteMultipartReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type CompleteMultipartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteMultipartResp) Reset() {
	*x = CompleteMultipartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteMultipartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteMultipartResp) ProtoMessage() {}

func (x *CompleteMultipartResp) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteMultipartResp.ProtoReflect.Descriptor instead.
func (*CompleteMultipartResp) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{5}
}

type GetSuccessChunksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash string `protobuf:"bytes,1,opt,name=FileHash,proto3" json:"FileHash,omitempty"`
	UserID   int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetSuccessChunksReq) Reset() {
	*x = GetSuccessChunksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuccessChunksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuccessChunksReq) ProtoMessage() {}

func (x *GetSuccessChunksReq) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuccessChunksReq.ProtoReflect.Descriptor instead.
func (*GetSuccessChunksReq) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{6}
}

func (x *GetSuccessChunksReq) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *GetSuccessChunksReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetSuccessChunksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUpload bool   `protobuf:"varint,1,opt,name=IsUpload,proto3" json:"IsUpload,omitempty"` // 是否文件完整存在minio
	IsRecord bool   `protobuf:"varint,2,opt,name=IsRecord,proto3" json:"IsRecord,omitempty"` // 是否已经记录在数据库
	Chunks   string `protobuf:"bytes,3,opt,name=Chunks,proto3" json:"Chunks,omitempty"`
}

func (x *GetSuccessChunksResp) Reset() {
	*x = GetSuccessChunksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuccessChunksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuccessChunksResp) ProtoMessage() {}

func (x *GetSuccessChunksResp) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuccessChunksResp.ProtoReflect.Descriptor instead.
func (*GetSuccessChunksResp) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{7}
}

func (x *GetSuccessChunksResp) GetIsUpload() bool {
	if x != nil {
		return x.IsUpload
	}
	return false
}

func (x *GetSuccessChunksResp) GetIsRecord() bool {
	if x != nil {
		return x.IsRecord
	}
	return false
}

func (x *GetSuccessChunksResp) GetChunks() string {
	if x != nil {
		return x.Chunks
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x4e, 0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x69, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x72, 0x69, 0x22, 0x4a, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x17,
	0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x32, 0xc1, 0x02, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x4e, 0x65,
	0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x77,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x72, 0x69, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x69, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x69, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x25,
	0x5a, 0x23, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_file_proto_goTypes = []interface{}{
	(*NewMultiUploadReq)(nil),     // 0: file.NewMultiUploadReq
	(*NewMultiUploadResp)(nil),    // 1: file.NewMultiUploadResp
	(*GetMultiUploadUriReq)(nil),  // 2: file.GetMultiUploadUriReq
	(*GetMultiUploadUriResp)(nil), // 3: file.GetMultiUploadUriResp
	(*CompleteMultipartReq)(nil),  // 4: file.CompleteMultipartReq
	(*CompleteMultipartResp)(nil), // 5: file.CompleteMultipartResp
	(*GetSuccessChunksReq)(nil),   // 6: file.GetSuccessChunksReq
	(*GetSuccessChunksResp)(nil),  // 7: file.GetSuccessChunksResp
}
var file_file_proto_depIdxs = []int32{
	0, // 0: file.FileService.NewMultiUpload:input_type -> file.NewMultiUploadReq
	2, // 1: file.FileService.GetMultiUploadUri:input_type -> file.GetMultiUploadUriReq
	4, // 2: file.FileService.CompleteMultipart:input_type -> file.CompleteMultipartReq
	6, // 3: file.FileService.GetSuccessChunks:input_type -> file.GetSuccessChunksReq
	1, // 4: file.FileService.NewMultiUpload:output_type -> file.NewMultiUploadResp
	3, // 5: file.FileService.GetMultiUploadUri:output_type -> file.GetMultiUploadUriResp
	5, // 6: file.FileService.CompleteMultipart:output_type -> file.CompleteMultipartResp
	7, // 7: file.FileService.GetSuccessChunks:output_type -> file.GetSuccessChunksResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMultiUploadReq); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMultiUploadResp); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultiUploadUriReq); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultiUploadUriResp); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteMultipartReq); i {
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
		file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteMultipartResp); i {
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
		file_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuccessChunksReq); i {
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
		file_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuccessChunksResp); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type FileService interface {
	NewMultiUpload(ctx context.Context, req *NewMultiUploadReq) (res *NewMultiUploadResp, err error)
	GetMultiUploadUri(ctx context.Context, req *GetMultiUploadUriReq) (res *GetMultiUploadUriResp, err error)
	CompleteMultipart(ctx context.Context, req *CompleteMultipartReq) (res *CompleteMultipartResp, err error)
	GetSuccessChunks(ctx context.Context, req *GetSuccessChunksReq) (res *GetSuccessChunksResp, err error)
}

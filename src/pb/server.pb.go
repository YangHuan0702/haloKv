// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: src/pb/server.proto

package pb

import (
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

type RequestVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         *int32 `protobuf:"varint,1,opt,name=Term,proto3,oneof" json:"Term,omitempty"`
	CandidateId  *int32 `protobuf:"varint,2,opt,name=CandidateId,proto3,oneof" json:"CandidateId,omitempty"`
	PrevLogIndex *int32 `protobuf:"varint,3,opt,name=PrevLogIndex,proto3,oneof" json:"PrevLogIndex,omitempty"`
	PrevLogTerm  *int32 `protobuf:"varint,4,opt,name=PrevLogTerm,proto3,oneof" json:"PrevLogTerm,omitempty"`
	For          *int32 `protobuf:"varint,5,opt,name=For,proto3,oneof" json:"For,omitempty"`
}

func (x *RequestVote) Reset() {
	*x = RequestVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVote) ProtoMessage() {}

func (x *RequestVote) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVote.ProtoReflect.Descriptor instead.
func (*RequestVote) Descriptor() ([]byte, []int) {
	return file_src_pb_server_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVote) GetTerm() int32 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *RequestVote) GetCandidateId() int32 {
	if x != nil && x.CandidateId != nil {
		return *x.CandidateId
	}
	return 0
}

func (x *RequestVote) GetPrevLogIndex() int32 {
	if x != nil && x.PrevLogIndex != nil {
		return *x.PrevLogIndex
	}
	return 0
}

func (x *RequestVote) GetPrevLogTerm() int32 {
	if x != nil && x.PrevLogTerm != nil {
		return *x.PrevLogTerm
	}
	return 0
}

func (x *RequestVote) GetFor() int32 {
	if x != nil && x.For != nil {
		return *x.For
	}
	return 0
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,opt,name=Key,proto3,oneof" json:"Key,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=Value,proto3,oneof" json:"Value,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_src_pb_server_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *Log) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type RaftLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term       *int32 `protobuf:"varint,1,opt,name=Term,proto3,oneof" json:"Term,omitempty"`
	LogEntries *Log   `protobuf:"bytes,2,opt,name=LogEntries,proto3,oneof" json:"LogEntries,omitempty"`
}

func (x *RaftLog) Reset() {
	*x = RaftLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftLog) ProtoMessage() {}

func (x *RaftLog) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftLog.ProtoReflect.Descriptor instead.
func (*RaftLog) Descriptor() ([]byte, []int) {
	return file_src_pb_server_proto_rawDescGZIP(), []int{2}
}

func (x *RaftLog) GetTerm() int32 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *RaftLog) GetLogEntries() *Log {
	if x != nil {
		return x.LogEntries
	}
	return nil
}

type ResponseVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        *int32 `protobuf:"varint,1,opt,name=Term,proto3,oneof" json:"Term,omitempty"`
	VoteGranted *bool  `protobuf:"varint,2,opt,name=VoteGranted,proto3,oneof" json:"VoteGranted,omitempty"`
}

func (x *ResponseVote) Reset() {
	*x = ResponseVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVote) ProtoMessage() {}

func (x *ResponseVote) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseVote.ProtoReflect.Descriptor instead.
func (*ResponseVote) Descriptor() ([]byte, []int) {
	return file_src_pb_server_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseVote) GetTerm() int32 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *ResponseVote) GetVoteGranted() bool {
	if x != nil && x.VoteGranted != nil {
		return *x.VoteGranted
	}
	return false
}

type RequestAppendEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         *int32     `protobuf:"varint,1,opt,name=Term,proto3,oneof" json:"Term,omitempty"`
	LeaderId     *int32     `protobuf:"varint,2,opt,name=LeaderId,proto3,oneof" json:"LeaderId,omitempty"`
	PrevLogIndex *int32     `protobuf:"varint,3,opt,name=PrevLogIndex,proto3,oneof" json:"PrevLogIndex,omitempty"`
	PrevLogTerm  *int32     `protobuf:"varint,4,opt,name=PrevLogTerm,proto3,oneof" json:"PrevLogTerm,omitempty"`
	LeaderCommit *int32     `protobuf:"varint,5,opt,name=LeaderCommit,proto3,oneof" json:"LeaderCommit,omitempty"`
	For          *int32     `protobuf:"varint,6,opt,name=For,proto3,oneof" json:"For,omitempty"`
	Logs         []*RaftLog `protobuf:"bytes,7,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (x *RequestAppendEntries) Reset() {
	*x = RequestAppendEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAppendEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAppendEntries) ProtoMessage() {}

func (x *RequestAppendEntries) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAppendEntries.ProtoReflect.Descriptor instead.
func (*RequestAppendEntries) Descriptor() ([]byte, []int) {
	return file_src_pb_server_proto_rawDescGZIP(), []int{4}
}

func (x *RequestAppendEntries) GetTerm() int32 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *RequestAppendEntries) GetLeaderId() int32 {
	if x != nil && x.LeaderId != nil {
		return *x.LeaderId
	}
	return 0
}

func (x *RequestAppendEntries) GetPrevLogIndex() int32 {
	if x != nil && x.PrevLogIndex != nil {
		return *x.PrevLogIndex
	}
	return 0
}

func (x *RequestAppendEntries) GetPrevLogTerm() int32 {
	if x != nil && x.PrevLogTerm != nil {
		return *x.PrevLogTerm
	}
	return 0
}

func (x *RequestAppendEntries) GetLeaderCommit() int32 {
	if x != nil && x.LeaderCommit != nil {
		return *x.LeaderCommit
	}
	return 0
}

func (x *RequestAppendEntries) GetFor() int32 {
	if x != nil && x.For != nil {
		return *x.For
	}
	return 0
}

func (x *RequestAppendEntries) GetLogs() []*RaftLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type ResponseAppendEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term       *int32 `protobuf:"varint,1,opt,name=Term,proto3,oneof" json:"Term,omitempty"`
	Success    *bool  `protobuf:"varint,2,opt,name=Success,proto3,oneof" json:"Success,omitempty"`
	FirstIndex *int32 `protobuf:"varint,3,opt,name=FirstIndex,proto3,oneof" json:"FirstIndex,omitempty"`
	FirstTerm  *int32 `protobuf:"varint,4,opt,name=FirstTerm,proto3,oneof" json:"FirstTerm,omitempty"`
}

func (x *ResponseAppendEntries) Reset() {
	*x = ResponseAppendEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAppendEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAppendEntries) ProtoMessage() {}

func (x *ResponseAppendEntries) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAppendEntries.ProtoReflect.Descriptor instead.
func (*ResponseAppendEntries) Descriptor() ([]byte, []int) {
	return file_src_pb_server_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseAppendEntries) GetTerm() int32 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *ResponseAppendEntries) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *ResponseAppendEntries) GetFirstIndex() int32 {
	if x != nil && x.FirstIndex != nil {
		return *x.FirstIndex
	}
	return 0
}

func (x *ResponseAppendEntries) GetFirstTerm() int32 {
	if x != nil && x.FirstTerm != nil {
		return *x.FirstTerm
	}
	return 0
}

var File_src_pb_server_proto protoreflect.FileDescriptor

var file_src_pb_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0c, 0x50,
	0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65,
	0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x46, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x04, 0x52, 0x03, 0x46, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x54, 0x65, 0x72, 0x6d, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x50, 0x72, 0x65, 0x76, 0x4c,
	0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x46, 0x6f, 0x72, 0x22, 0x49,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x15, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x4b, 0x65, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x07, 0x52, 0x61, 0x66,
	0x74, 0x4c, 0x6f, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0a, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x48, 0x01, 0x52, 0x0a, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x65, 0x72,
	0x6d, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x67, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x56, 0x6f, 0x74,
	0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01,
	0x52, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x65, 0x72, 0x6d, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x56, 0x6f,
	0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0xce, 0x02, 0x0a, 0x14, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x08, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x02, 0x52, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67,
	0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0b, 0x50, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x04, 0x52, 0x0c, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x46, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x05, 0x52, 0x03, 0x46, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x61, 0x66,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54,
	0x65, 0x72, 0x6d, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72,
	0x6d, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x46, 0x6f, 0x72, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01,
	0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x02, 0x52, 0x0a, 0x46, 0x69, 0x72, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x54, 0x65, 0x72,
	0x6d, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x65, 0x72, 0x6d, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x32, 0x79, 0x0a, 0x07, 0x52, 0x61, 0x66, 0x74, 0x52, 0x70,
	0x63, 0x12, 0x2c, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x1a, 0x0d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x15, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x68, 0x61, 0x6c, 0x6f, 0x4b, 0x76, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_pb_server_proto_rawDescOnce sync.Once
	file_src_pb_server_proto_rawDescData = file_src_pb_server_proto_rawDesc
)

func file_src_pb_server_proto_rawDescGZIP() []byte {
	file_src_pb_server_proto_rawDescOnce.Do(func() {
		file_src_pb_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_pb_server_proto_rawDescData)
	})
	return file_src_pb_server_proto_rawDescData
}

var file_src_pb_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_pb_server_proto_goTypes = []interface{}{
	(*RequestVote)(nil),           // 0: RequestVote
	(*Log)(nil),                   // 1: Log
	(*RaftLog)(nil),               // 2: RaftLog
	(*ResponseVote)(nil),          // 3: ResponseVote
	(*RequestAppendEntries)(nil),  // 4: RequestAppendEntries
	(*ResponseAppendEntries)(nil), // 5: ResponseAppendEntries
}
var file_src_pb_server_proto_depIdxs = []int32{
	1, // 0: RaftLog.LogEntries:type_name -> Log
	2, // 1: RequestAppendEntries.Logs:type_name -> RaftLog
	0, // 2: RaftRpc.VoteRequest:input_type -> RequestVote
	4, // 3: RaftRpc.AppendEntries:input_type -> RequestAppendEntries
	3, // 4: RaftRpc.VoteRequest:output_type -> ResponseVote
	5, // 5: RaftRpc.AppendEntries:output_type -> ResponseAppendEntries
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_src_pb_server_proto_init() }
func file_src_pb_server_proto_init() {
	if File_src_pb_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_pb_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVote); i {
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
		file_src_pb_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_src_pb_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftLog); i {
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
		file_src_pb_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseVote); i {
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
		file_src_pb_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAppendEntries); i {
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
		file_src_pb_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAppendEntries); i {
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
	file_src_pb_server_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_src_pb_server_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_src_pb_server_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_src_pb_server_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_src_pb_server_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_src_pb_server_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_pb_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_pb_server_proto_goTypes,
		DependencyIndexes: file_src_pb_server_proto_depIdxs,
		MessageInfos:      file_src_pb_server_proto_msgTypes,
	}.Build()
	File_src_pb_server_proto = out.File
	file_src_pb_server_proto_rawDesc = nil
	file_src_pb_server_proto_goTypes = nil
	file_src_pb_server_proto_depIdxs = nil
}

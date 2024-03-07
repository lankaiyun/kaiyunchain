// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: kaiyunchain.proto

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

// ----------- GetAllBlock -------------
type GetAllBlockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin string `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetAllBlockReq) Reset() {
	*x = GetAllBlockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaiyunchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBlockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBlockReq) ProtoMessage() {}

func (x *GetAllBlockReq) ProtoReflect() protoreflect.Message {
	mi := &file_kaiyunchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBlockReq.ProtoReflect.Descriptor instead.
func (*GetAllBlockReq) Descriptor() ([]byte, []int) {
	return file_kaiyunchain_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllBlockReq) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *GetAllBlockReq) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type GetAllBlockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*GetAllBlockResp_Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *GetAllBlockResp) Reset() {
	*x = GetAllBlockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaiyunchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBlockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBlockResp) ProtoMessage() {}

func (x *GetAllBlockResp) ProtoReflect() protoreflect.Message {
	mi := &file_kaiyunchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBlockResp.ProtoReflect.Descriptor instead.
func (*GetAllBlockResp) Descriptor() ([]byte, []int) {
	return file_kaiyunchain_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllBlockResp) GetBlocks() []*GetAllBlockResp_Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

// ------------- GetAllTx --------------
type GetAllTxReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin string `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetAllTxReq) Reset() {
	*x = GetAllTxReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaiyunchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTxReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTxReq) ProtoMessage() {}

func (x *GetAllTxReq) ProtoReflect() protoreflect.Message {
	mi := &file_kaiyunchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTxReq.ProtoReflect.Descriptor instead.
func (*GetAllTxReq) Descriptor() ([]byte, []int) {
	return file_kaiyunchain_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllTxReq) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *GetAllTxReq) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type GetAllTxResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs []*GetAllTxResp_Tx `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *GetAllTxResp) Reset() {
	*x = GetAllTxResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaiyunchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTxResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTxResp) ProtoMessage() {}

func (x *GetAllTxResp) ProtoReflect() protoreflect.Message {
	mi := &file_kaiyunchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTxResp.ProtoReflect.Descriptor instead.
func (*GetAllTxResp) Descriptor() ([]byte, []int) {
	return file_kaiyunchain_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTxResp) GetTxs() []*GetAllTxResp_Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

type GetAllBlockResp_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height   string `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
	Time     string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Coinbase string `protobuf:"bytes,3,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Txs      string `protobuf:"bytes,4,opt,name=txs,proto3" json:"txs,omitempty"`
	Reward   string `protobuf:"bytes,5,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *GetAllBlockResp_Block) Reset() {
	*x = GetAllBlockResp_Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaiyunchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBlockResp_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBlockResp_Block) ProtoMessage() {}

func (x *GetAllBlockResp_Block) ProtoReflect() protoreflect.Message {
	mi := &file_kaiyunchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBlockResp_Block.ProtoReflect.Descriptor instead.
func (*GetAllBlockResp_Block) Descriptor() ([]byte, []int) {
	return file_kaiyunchain_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAllBlockResp_Block) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *GetAllBlockResp_Block) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *GetAllBlockResp_Block) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *GetAllBlockResp_Block) GetTxs() string {
	if x != nil {
		return x.Txs
	}
	return ""
}

func (x *GetAllBlockResp_Block) GetReward() string {
	if x != nil {
		return x.Reward
	}
	return ""
}

type GetAllTxResp_Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash      string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	From        string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To          string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value       string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Time        string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	BelongBlock string `protobuf:"bytes,6,opt,name=belongBlock,proto3" json:"belongBlock,omitempty"`
}

func (x *GetAllTxResp_Tx) Reset() {
	*x = GetAllTxResp_Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaiyunchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTxResp_Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTxResp_Tx) ProtoMessage() {}

func (x *GetAllTxResp_Tx) ProtoReflect() protoreflect.Message {
	mi := &file_kaiyunchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTxResp_Tx.ProtoReflect.Descriptor instead.
func (*GetAllTxResp_Tx) Descriptor() ([]byte, []int) {
	return file_kaiyunchain_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetAllTxResp_Tx) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *GetAllTxResp_Tx) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetAllTxResp_Tx) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *GetAllTxResp_Tx) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetAllTxResp_Tx) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *GetAllTxResp_Tx) GetBelongBlock() string {
	if x != nil {
		return x.BelongBlock
	}
	return ""
}

var File_kaiyunchain_proto protoreflect.FileDescriptor

var file_kaiyunchain_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6b, 0x61, 0x69, 0x79, 0x75, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xbc, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2e, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x1a, 0x79, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x78, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x35, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x78, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78, 0x73, 0x1a, 0x8c, 0x01, 0x0a, 0x02, 0x54, 0x78, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x65, 0x6c, 0x6f,
	0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x64, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x32,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x78, 0x12, 0x0c,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kaiyunchain_proto_rawDescOnce sync.Once
	file_kaiyunchain_proto_rawDescData = file_kaiyunchain_proto_rawDesc
)

func file_kaiyunchain_proto_rawDescGZIP() []byte {
	file_kaiyunchain_proto_rawDescOnce.Do(func() {
		file_kaiyunchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_kaiyunchain_proto_rawDescData)
	})
	return file_kaiyunchain_proto_rawDescData
}

var file_kaiyunchain_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kaiyunchain_proto_goTypes = []interface{}{
	(*GetAllBlockReq)(nil),        // 0: GetAllBlockReq
	(*GetAllBlockResp)(nil),       // 1: GetAllBlockResp
	(*GetAllTxReq)(nil),           // 2: GetAllTxReq
	(*GetAllTxResp)(nil),          // 3: GetAllTxResp
	(*GetAllBlockResp_Block)(nil), // 4: GetAllBlockResp.Block
	(*GetAllTxResp_Tx)(nil),       // 5: GetAllTxResp.Tx
}
var file_kaiyunchain_proto_depIdxs = []int32{
	4, // 0: GetAllBlockResp.blocks:type_name -> GetAllBlockResp.Block
	5, // 1: GetAllTxResp.txs:type_name -> GetAllTxResp.Tx
	0, // 2: Rpc.GetAllBlock:input_type -> GetAllBlockReq
	2, // 3: Rpc.GetAllTx:input_type -> GetAllTxReq
	1, // 4: Rpc.GetAllBlock:output_type -> GetAllBlockResp
	3, // 5: Rpc.GetAllTx:output_type -> GetAllTxResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kaiyunchain_proto_init() }
func file_kaiyunchain_proto_init() {
	if File_kaiyunchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kaiyunchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBlockReq); i {
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
		file_kaiyunchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBlockResp); i {
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
		file_kaiyunchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTxReq); i {
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
		file_kaiyunchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTxResp); i {
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
		file_kaiyunchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBlockResp_Block); i {
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
		file_kaiyunchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTxResp_Tx); i {
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
			RawDescriptor: file_kaiyunchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kaiyunchain_proto_goTypes,
		DependencyIndexes: file_kaiyunchain_proto_depIdxs,
		MessageInfos:      file_kaiyunchain_proto_msgTypes,
	}.Build()
	File_kaiyunchain_proto = out.File
	file_kaiyunchain_proto_rawDesc = nil
	file_kaiyunchain_proto_goTypes = nil
	file_kaiyunchain_proto_depIdxs = nil
}

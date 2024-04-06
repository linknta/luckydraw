// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/luckydraw.proto

package luckydraw

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

type RollRquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RollRquest) Reset() {
	*x = RollRquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_luckydraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollRquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollRquest) ProtoMessage() {}

func (x *RollRquest) ProtoReflect() protoreflect.Message {
	mi := &file_api_luckydraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollRquest.ProtoReflect.Descriptor instead.
func (*RollRquest) Descriptor() ([]byte, []int) {
	return file_api_luckydraw_proto_rawDescGZIP(), []int{0}
}

type RollReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	State  *State  `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *RollReply) Reset() {
	*x = RollReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_luckydraw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollReply) ProtoMessage() {}

func (x *RollReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_luckydraw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollReply.ProtoReflect.Descriptor instead.
func (*RollReply) Descriptor() ([]byte, []int) {
	return file_api_luckydraw_proto_rawDescGZIP(), []int{1}
}

func (x *RollReply) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *RollReply) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

// model
type Cell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row    int32  `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Reel   int32  `protobuf:"varint,2,opt,name=reel,proto3" json:"reel,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *Cell) Reset() {
	*x = Cell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_luckydraw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_api_luckydraw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_api_luckydraw_proto_rawDescGZIP(), []int{2}
}

func (x *Cell) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *Cell) GetReel() int32 {
	if x != nil {
		return x.Reel
	}
	return 0
}

func (x *Cell) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type Win struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Coin   int64  `protobuf:"varint,2,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (x *Win) Reset() {
	*x = Win{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_luckydraw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Win) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Win) ProtoMessage() {}

func (x *Win) ProtoReflect() protoreflect.Message {
	mi := &file_api_luckydraw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Win.ProtoReflect.Descriptor instead.
func (*Win) Descriptor() ([]byte, []int) {
	return file_api_luckydraw_proto_rawDescGZIP(), []int{3}
}

func (x *Win) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Win) GetCoin() int64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []*Cell `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
	Win   *Win    `protobuf:"bytes,2,opt,name=win,proto3" json:"win,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_luckydraw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_luckydraw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_luckydraw_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

func (x *Result) GetWin() *Win {
	if x != nil {
		return x.Win
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCoin      int64   `protobuf:"varint,1,opt,name=total_coin,json=totalCoin,proto3" json:"total_coin,omitempty"`
	RemainingSpins int64   `protobuf:"varint,2,opt,name=remaining_spins,json=remainingSpins,proto3" json:"remaining_spins,omitempty"`
	LastResult     *Result `protobuf:"bytes,3,opt,name=last_result,json=lastResult,proto3" json:"last_result,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_luckydraw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_api_luckydraw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_api_luckydraw_proto_rawDescGZIP(), []int{5}
}

func (x *State) GetTotalCoin() int64 {
	if x != nil {
		return x.TotalCoin
	}
	return 0
}

func (x *State) GetRemainingSpins() int64 {
	if x != nil {
		return x.RemainingSpins
	}
	return 0
}

func (x *State) GetLastResult() *Result {
	if x != nil {
		return x.LastResult
	}
	return nil
}

var File_api_luckydraw_proto protoreflect.FileDescriptor

var file_api_luckydraw_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77,
	0x2e, 0x76, 0x31, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x64, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x75,
	0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f,
	0x77, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x72, 0x65, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x31, 0x0a,
	0x03, 0x57, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e,
	0x22, 0x57, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x65,
	0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x75, 0x63, 0x6b,
	0x79, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63,
	0x65, 0x6c, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x69, 0x6e, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x70, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x6d,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x69, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x46, 0x0a, 0x09, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x12,
	0x39, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64,
	0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x6e, 0x74, 0x61,
	0x2f, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x64, 0x72, 0x61, 0x77, 0x3b, 0x6c, 0x75, 0x63, 0x6b,
	0x79, 0x64, 0x72, 0x61, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_luckydraw_proto_rawDescOnce sync.Once
	file_api_luckydraw_proto_rawDescData = file_api_luckydraw_proto_rawDesc
)

func file_api_luckydraw_proto_rawDescGZIP() []byte {
	file_api_luckydraw_proto_rawDescOnce.Do(func() {
		file_api_luckydraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_luckydraw_proto_rawDescData)
	})
	return file_api_luckydraw_proto_rawDescData
}

var file_api_luckydraw_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_luckydraw_proto_goTypes = []interface{}{
	(*RollRquest)(nil), // 0: luckydraw.v1.RollRquest
	(*RollReply)(nil),  // 1: luckydraw.v1.RollReply
	(*Cell)(nil),       // 2: luckydraw.v1.Cell
	(*Win)(nil),        // 3: luckydraw.v1.Win
	(*Result)(nil),     // 4: luckydraw.v1.Result
	(*State)(nil),      // 5: luckydraw.v1.State
}
var file_api_luckydraw_proto_depIdxs = []int32{
	4, // 0: luckydraw.v1.RollReply.result:type_name -> luckydraw.v1.Result
	5, // 1: luckydraw.v1.RollReply.state:type_name -> luckydraw.v1.State
	2, // 2: luckydraw.v1.Result.cells:type_name -> luckydraw.v1.Cell
	3, // 3: luckydraw.v1.Result.win:type_name -> luckydraw.v1.Win
	4, // 4: luckydraw.v1.State.last_result:type_name -> luckydraw.v1.Result
	0, // 5: luckydraw.v1.Luckydraw.Roll:input_type -> luckydraw.v1.RollRquest
	1, // 6: luckydraw.v1.Luckydraw.Roll:output_type -> luckydraw.v1.RollReply
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_luckydraw_proto_init() }
func file_api_luckydraw_proto_init() {
	if File_api_luckydraw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_luckydraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollRquest); i {
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
		file_api_luckydraw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollReply); i {
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
		file_api_luckydraw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cell); i {
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
		file_api_luckydraw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Win); i {
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
		file_api_luckydraw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_api_luckydraw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
			RawDescriptor: file_api_luckydraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_luckydraw_proto_goTypes,
		DependencyIndexes: file_api_luckydraw_proto_depIdxs,
		MessageInfos:      file_api_luckydraw_proto_msgTypes,
	}.Build()
	File_api_luckydraw_proto = out.File
	file_api_luckydraw_proto_rawDesc = nil
	file_api_luckydraw_proto_goTypes = nil
	file_api_luckydraw_proto_depIdxs = nil
}

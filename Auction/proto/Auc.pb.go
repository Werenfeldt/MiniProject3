// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: Auction/proto/Auc.proto

package Auction

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

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bid int32 `protobuf:"varint,2,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auction_proto_Auc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Auction_proto_Auc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_Auction_proto_Auc_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BidRequest) GetBid() int32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Mssg     string `protobuf:"bytes,2,opt,name=mssg,proto3" json:"mssg,omitempty"`
	AucDone  bool   `protobuf:"varint,3,opt,name=aucDone,proto3" json:"aucDone,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auction_proto_Auc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Auction_proto_Auc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_Auction_proto_Auc_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *BidResponse) GetMssg() string {
	if x != nil {
		return x.Mssg
	}
	return ""
}

func (x *BidResponse) GetAucDone() bool {
	if x != nil {
		return x.AucDone
	}
	return false
}

type GetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GetResult string `protobuf:"bytes,2,opt,name=getResult,proto3" json:"getResult,omitempty"`
}

func (x *GetResult) Reset() {
	*x = GetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auction_proto_Auc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResult) ProtoMessage() {}

func (x *GetResult) ProtoReflect() protoreflect.Message {
	mi := &file_Auction_proto_Auc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResult.ProtoReflect.Descriptor instead.
func (*GetResult) Descriptor() ([]byte, []int) {
	return file_Auction_proto_Auc_proto_rawDescGZIP(), []int{2}
}

func (x *GetResult) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetResult) GetGetResult() string {
	if x != nil {
		return x.GetResult
	}
	return ""
}

type HighestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HighestResult) Reset() {
	*x = HighestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auction_proto_Auc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HighestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighestResult) ProtoMessage() {}

func (x *HighestResult) ProtoReflect() protoreflect.Message {
	mi := &file_Auction_proto_Auc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighestResult.ProtoReflect.Descriptor instead.
func (*HighestResult) Descriptor() ([]byte, []int) {
	return file_Auction_proto_Auc_proto_rawDescGZIP(), []int{3}
}

func (x *HighestResult) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type DoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *DoneRequest) Reset() {
	*x = DoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auction_proto_Auc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneRequest) ProtoMessage() {}

func (x *DoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Auction_proto_Auc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneRequest.ProtoReflect.Descriptor instead.
func (*DoneRequest) Descriptor() ([]byte, []int) {
	return file_Auction_proto_Auc_proto_rawDescGZIP(), []int{4}
}

func (x *DoneRequest) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auction_proto_Auc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_Auction_proto_Auc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_Auction_proto_Auc_proto_rawDescGZIP(), []int{5}
}

var File_Auction_proto_Auc_proto protoreflect.FileDescriptor

var file_Auction_proto_Auc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x41, 0x75, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x2e, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x69, 0x64, 0x22,
	0x57, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x73, 0x73, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x75, 0x63, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x61, 0x75, 0x63, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x27, 0x0a, 0x0d, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x0b,
	0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0x8d, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x03,
	0x42, 0x69, 0x64, 0x12, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x48, 0x69, 0x67, 0x68, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65,
	0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x34, 0x5a, 0x32, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x65, 0x72, 0x65, 0x6e, 0x66, 0x65, 0x6c, 0x64,
	0x74, 0x2f, 0x4d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x33, 0x2f, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Auction_proto_Auc_proto_rawDescOnce sync.Once
	file_Auction_proto_Auc_proto_rawDescData = file_Auction_proto_Auc_proto_rawDesc
)

func file_Auction_proto_Auc_proto_rawDescGZIP() []byte {
	file_Auction_proto_Auc_proto_rawDescOnce.Do(func() {
		file_Auction_proto_Auc_proto_rawDescData = protoimpl.X.CompressGZIP(file_Auction_proto_Auc_proto_rawDescData)
	})
	return file_Auction_proto_Auc_proto_rawDescData
}

var file_Auction_proto_Auc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Auction_proto_Auc_proto_goTypes = []interface{}{
	(*BidRequest)(nil),    // 0: main.BidRequest
	(*BidResponse)(nil),   // 1: main.BidResponse
	(*GetResult)(nil),     // 2: main.GetResult
	(*HighestResult)(nil), // 3: main.HighestResult
	(*DoneRequest)(nil),   // 4: main.DoneRequest
	(*Empty)(nil),         // 5: main.Empty
}
var file_Auction_proto_Auc_proto_depIdxs = []int32{
	0, // 0: main.Auction.Bid:input_type -> main.BidRequest
	2, // 1: main.Auction.Result:input_type -> main.GetResult
	4, // 2: main.Auction.Done:input_type -> main.DoneRequest
	1, // 3: main.Auction.Bid:output_type -> main.BidResponse
	3, // 4: main.Auction.Result:output_type -> main.HighestResult
	5, // 5: main.Auction.Done:output_type -> main.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Auction_proto_Auc_proto_init() }
func file_Auction_proto_Auc_proto_init() {
	if File_Auction_proto_Auc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Auction_proto_Auc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_Auction_proto_Auc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidResponse); i {
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
		file_Auction_proto_Auc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResult); i {
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
		file_Auction_proto_Auc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HighestResult); i {
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
		file_Auction_proto_Auc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneRequest); i {
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
		file_Auction_proto_Auc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_Auction_proto_Auc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Auction_proto_Auc_proto_goTypes,
		DependencyIndexes: file_Auction_proto_Auc_proto_depIdxs,
		MessageInfos:      file_Auction_proto_Auc_proto_msgTypes,
	}.Build()
	File_Auction_proto_Auc_proto = out.File
	file_Auction_proto_Auc_proto_rawDesc = nil
	file_Auction_proto_Auc_proto_goTypes = nil
	file_Auction_proto_Auc_proto_depIdxs = nil
}
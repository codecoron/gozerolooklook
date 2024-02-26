// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: shop.proto

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

type Goods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GoodsId              string  `protobuf:"bytes,2,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	CategoryId           int64   `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	PrecouponPrice       float64 `protobuf:"fixed64,4,opt,name=precoupon_price,json=precouponPrice,proto3" json:"precoupon_price,omitempty"`
	AftercouponPrice     float64 `protobuf:"fixed64,5,opt,name=aftercoupon_price,json=aftercouponPrice,proto3" json:"aftercoupon_price,omitempty"`
	GoodsDesc            string  `protobuf:"bytes,6,opt,name=goods_desc,json=goodsDesc,proto3" json:"goods_desc,omitempty"`
	WishPoints           int64   `protobuf:"varint,7,opt,name=wish_points,json=wishPoints,proto3" json:"wish_points,omitempty"`
	CouponStartTime      int64   `protobuf:"varint,8,opt,name=coupon_start_time,json=couponStartTime,proto3" json:"coupon_start_time,omitempty"`
	CouponEndTime        int64   `protobuf:"varint,9,opt,name=coupon_end_time,json=couponEndTime,proto3" json:"coupon_end_time,omitempty"`
	CouponDiscount       int64   `protobuf:"varint,10,opt,name=coupon_discount,json=couponDiscount,proto3" json:"coupon_discount,omitempty"`
	CouponRemainQuantity int64   `protobuf:"varint,11,opt,name=coupon_remain_quantity,json=couponRemainQuantity,proto3" json:"coupon_remain_quantity,omitempty"`
}

func (x *Goods) Reset() {
	*x = Goods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goods) ProtoMessage() {}

func (x *Goods) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goods.ProtoReflect.Descriptor instead.
func (*Goods) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{0}
}

func (x *Goods) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Goods) GetGoodsId() string {
	if x != nil {
		return x.GoodsId
	}
	return ""
}

func (x *Goods) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Goods) GetPrecouponPrice() float64 {
	if x != nil {
		return x.PrecouponPrice
	}
	return 0
}

func (x *Goods) GetAftercouponPrice() float64 {
	if x != nil {
		return x.AftercouponPrice
	}
	return 0
}

func (x *Goods) GetGoodsDesc() string {
	if x != nil {
		return x.GoodsDesc
	}
	return ""
}

func (x *Goods) GetWishPoints() int64 {
	if x != nil {
		return x.WishPoints
	}
	return 0
}

func (x *Goods) GetCouponStartTime() int64 {
	if x != nil {
		return x.CouponStartTime
	}
	return 0
}

func (x *Goods) GetCouponEndTime() int64 {
	if x != nil {
		return x.CouponEndTime
	}
	return 0
}

func (x *Goods) GetCouponDiscount() int64 {
	if x != nil {
		return x.CouponDiscount
	}
	return 0
}

func (x *Goods) GetCouponRemainQuantity() int64 {
	if x != nil {
		return x.CouponRemainQuantity
	}
	return 0
}

type GoodsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GoodsReq) Reset() {
	*x = GoodsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsReq) ProtoMessage() {}

func (x *GoodsReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsReq.ProtoReflect.Descriptor instead.
func (*GoodsReq) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoodsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods *Goods `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
}

func (x *GoodsResp) Reset() {
	*x = GoodsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsResp) ProtoMessage() {}

func (x *GoodsResp) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsResp.ProtoReflect.Descriptor instead.
func (*GoodsResp) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsResp) GetGoods() *Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

var File_shop_proto protoreflect.FileDescriptor

var file_shop_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x22, 0x9c, 0x03, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72,
	0x65, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x66, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x73, 0x63, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x16, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x30, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x05,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x32, 0x3b, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x33, 0x0a, 0x0c, 0x67, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_proto_rawDescOnce sync.Once
	file_shop_proto_rawDescData = file_shop_proto_rawDesc
)

func file_shop_proto_rawDescGZIP() []byte {
	file_shop_proto_rawDescOnce.Do(func() {
		file_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_proto_rawDescData)
	})
	return file_shop_proto_rawDescData
}

var file_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shop_proto_goTypes = []interface{}{
	(*Goods)(nil),     // 0: stream.Goods
	(*GoodsReq)(nil),  // 1: stream.GoodsReq
	(*GoodsResp)(nil), // 2: stream.GoodsResp
}
var file_shop_proto_depIdxs = []int32{
	0, // 0: stream.GoodsResp.goods:type_name -> stream.Goods
	1, // 1: stream.Shop.getGoodsById:input_type -> stream.GoodsReq
	2, // 2: stream.Shop.getGoodsById:output_type -> stream.GoodsResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shop_proto_init() }
func file_shop_proto_init() {
	if File_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goods); i {
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
		file_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsReq); i {
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
		file_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsResp); i {
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
			RawDescriptor: file_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_proto_goTypes,
		DependencyIndexes: file_shop_proto_depIdxs,
		MessageInfos:      file_shop_proto_msgTypes,
	}.Build()
	File_shop_proto = out.File
	file_shop_proto_rawDesc = nil
	file_shop_proto_goTypes = nil
	file_shop_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: scale.proto

package render

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ChartScaleKind contains available scale kinds.
type ChartScale_ChartScaleKind int32

const (
	ChartScale_UNSPECIFIED_SCALE ChartScale_ChartScaleKind = 0
	ChartScale_LINEAR            ChartScale_ChartScaleKind = 1
	ChartScale_BAND              ChartScale_ChartScaleKind = 2
)

// Enum value maps for ChartScale_ChartScaleKind.
var (
	ChartScale_ChartScaleKind_name = map[int32]string{
		0: "UNSPECIFIED_SCALE",
		1: "LINEAR",
		2: "BAND",
	}
	ChartScale_ChartScaleKind_value = map[string]int32{
		"UNSPECIFIED_SCALE": 0,
		"LINEAR":            1,
		"BAND":              2,
	}
)

func (x ChartScale_ChartScaleKind) Enum() *ChartScale_ChartScaleKind {
	p := new(ChartScale_ChartScaleKind)
	*p = x
	return p
}

func (x ChartScale_ChartScaleKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartScale_ChartScaleKind) Descriptor() protoreflect.EnumDescriptor {
	return file_scale_proto_enumTypes[0].Descriptor()
}

func (ChartScale_ChartScaleKind) Type() protoreflect.EnumType {
	return &file_scale_proto_enumTypes[0]
}

func (x ChartScale_ChartScaleKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartScale_ChartScaleKind.Descriptor instead.
func (ChartScale_ChartScaleKind) EnumDescriptor() ([]byte, []int) {
	return file_scale_proto_rawDescGZIP(), []int{0, 0}
}

// ChartScale represents options to configure chart scale.
type ChartScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One of the available scale kinds.
	Kind ChartScale_ChartScaleKind `protobuf:"varint,1,opt,name=kind,proto3,enum=render.ChartScale_ChartScaleKind" json:"kind,omitempty"`
	// Start of the scale range.
	RangeStart *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=range_start,json=rangeStart,proto3" json:"range_start,omitempty"`
	// End of the scale range.
	RangeEnd *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	// Scale domain with one of available kind.
	//
	// Types that are assignable to Domain:
	//	*ChartScale_DomainNumeric
	//	*ChartScale_DomainCategories
	Domain isChartScale_Domain `protobuf_oneof:"domain"`
	// Does this scale needs an offset from the start and end of an axis.
	// This is usually need for an area or line views.
	NoBoundariesOffset bool `protobuf:"varint,7,opt,name=no_boundaries_offset,json=noBoundariesOffset,proto3" json:"no_boundaries_offset,omitempty"`
	// Inner padding for categories.
	InnerPadding *wrapperspb.FloatValue `protobuf:"bytes,8,opt,name=inner_padding,json=innerPadding,proto3" json:"inner_padding,omitempty"`
	// Outer padding for categories.
	OuterPadding *wrapperspb.FloatValue `protobuf:"bytes,9,opt,name=outer_padding,json=outerPadding,proto3" json:"outer_padding,omitempty"`
}

func (x *ChartScale) Reset() {
	*x = ChartScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scale_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartScale) ProtoMessage() {}

func (x *ChartScale) ProtoReflect() protoreflect.Message {
	mi := &file_scale_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartScale.ProtoReflect.Descriptor instead.
func (*ChartScale) Descriptor() ([]byte, []int) {
	return file_scale_proto_rawDescGZIP(), []int{0}
}

func (x *ChartScale) GetKind() ChartScale_ChartScaleKind {
	if x != nil {
		return x.Kind
	}
	return ChartScale_UNSPECIFIED_SCALE
}

func (x *ChartScale) GetRangeStart() *wrapperspb.Int32Value {
	if x != nil {
		return x.RangeStart
	}
	return nil
}

func (x *ChartScale) GetRangeEnd() *wrapperspb.Int32Value {
	if x != nil {
		return x.RangeEnd
	}
	return nil
}

func (m *ChartScale) GetDomain() isChartScale_Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (x *ChartScale) GetDomainNumeric() *DomainNumeric {
	if x, ok := x.GetDomain().(*ChartScale_DomainNumeric); ok {
		return x.DomainNumeric
	}
	return nil
}

func (x *ChartScale) GetDomainCategories() *DomainCategories {
	if x, ok := x.GetDomain().(*ChartScale_DomainCategories); ok {
		return x.DomainCategories
	}
	return nil
}

func (x *ChartScale) GetNoBoundariesOffset() bool {
	if x != nil {
		return x.NoBoundariesOffset
	}
	return false
}

func (x *ChartScale) GetInnerPadding() *wrapperspb.FloatValue {
	if x != nil {
		return x.InnerPadding
	}
	return nil
}

func (x *ChartScale) GetOuterPadding() *wrapperspb.FloatValue {
	if x != nil {
		return x.OuterPadding
	}
	return nil
}

type isChartScale_Domain interface {
	isChartScale_Domain()
}

type ChartScale_DomainNumeric struct {
	// Numeric scale domain.
	DomainNumeric *DomainNumeric `protobuf:"bytes,4,opt,name=domain_numeric,json=domainNumeric,proto3,oneof"`
}

type ChartScale_DomainCategories struct {
	// String scale domain categories.
	DomainCategories *DomainCategories `protobuf:"bytes,5,opt,name=domain_categories,json=domainCategories,proto3,oneof"`
}

func (*ChartScale_DomainNumeric) isChartScale_Domain() {}

func (*ChartScale_DomainCategories) isChartScale_Domain() {}

// DomainNumeric represents numeric scale domain.
type DomainNumeric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start of the numeric scale domain.
	Start float32 `protobuf:"fixed32,1,opt,name=start,proto3" json:"start,omitempty"`
	// End of the numeric scale domain.
	End float32 `protobuf:"fixed32,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DomainNumeric) Reset() {
	*x = DomainNumeric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scale_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainNumeric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainNumeric) ProtoMessage() {}

func (x *DomainNumeric) ProtoReflect() protoreflect.Message {
	mi := &file_scale_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainNumeric.ProtoReflect.Descriptor instead.
func (*DomainNumeric) Descriptor() ([]byte, []int) {
	return file_scale_proto_rawDescGZIP(), []int{1}
}

func (x *DomainNumeric) GetStart() float32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DomainNumeric) GetEnd() float32 {
	if x != nil {
		return x.End
	}
	return 0
}

// DomainCategories represents string categorical scale domain.
type DomainCategories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []string `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *DomainCategories) Reset() {
	*x = DomainCategories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scale_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCategories) ProtoMessage() {}

func (x *DomainCategories) ProtoReflect() protoreflect.Message {
	mi := &file_scale_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCategories.ProtoReflect.Descriptor instead.
func (*DomainCategories) Descriptor() ([]byte, []int) {
	return file_scale_proto_rawDescGZIP(), []int{2}
}

func (x *DomainCategories) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_scale_proto protoreflect.FileDescriptor

var file_scale_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x04, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53,
	0x63, 0x61, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x45, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x12, 0x47, 0x0a, 0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x10, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x6e, 0x6f, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6e, 0x6f, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x40,
	0x0a, 0x0d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x40, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x22, 0x3d, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c,
	0x49, 0x4e, 0x45, 0x41, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x4e, 0x44, 0x10,
	0x02, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x37, 0x0a, 0x0d, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0x32, 0x0a, 0x10, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6d, 0x70, 0x69, 0x64, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x2f, 0x6c, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x30, 0x3b, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scale_proto_rawDescOnce sync.Once
	file_scale_proto_rawDescData = file_scale_proto_rawDesc
)

func file_scale_proto_rawDescGZIP() []byte {
	file_scale_proto_rawDescOnce.Do(func() {
		file_scale_proto_rawDescData = protoimpl.X.CompressGZIP(file_scale_proto_rawDescData)
	})
	return file_scale_proto_rawDescData
}

var file_scale_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scale_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scale_proto_goTypes = []interface{}{
	(ChartScale_ChartScaleKind)(0), // 0: render.ChartScale.ChartScaleKind
	(*ChartScale)(nil),             // 1: render.ChartScale
	(*DomainNumeric)(nil),          // 2: render.DomainNumeric
	(*DomainCategories)(nil),       // 3: render.DomainCategories
	(*wrapperspb.Int32Value)(nil),  // 4: google.protobuf.Int32Value
	(*wrapperspb.FloatValue)(nil),  // 5: google.protobuf.FloatValue
}
var file_scale_proto_depIdxs = []int32{
	0, // 0: render.ChartScale.kind:type_name -> render.ChartScale.ChartScaleKind
	4, // 1: render.ChartScale.range_start:type_name -> google.protobuf.Int32Value
	4, // 2: render.ChartScale.range_end:type_name -> google.protobuf.Int32Value
	2, // 3: render.ChartScale.domain_numeric:type_name -> render.DomainNumeric
	3, // 4: render.ChartScale.domain_categories:type_name -> render.DomainCategories
	5, // 5: render.ChartScale.inner_padding:type_name -> google.protobuf.FloatValue
	5, // 6: render.ChartScale.outer_padding:type_name -> google.protobuf.FloatValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_scale_proto_init() }
func file_scale_proto_init() {
	if File_scale_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scale_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartScale); i {
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
		file_scale_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainNumeric); i {
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
		file_scale_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCategories); i {
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
	file_scale_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ChartScale_DomainNumeric)(nil),
		(*ChartScale_DomainCategories)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scale_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scale_proto_goTypes,
		DependencyIndexes: file_scale_proto_depIdxs,
		EnumInfos:         file_scale_proto_enumTypes,
		MessageInfos:      file_scale_proto_msgTypes,
	}.Build()
	File_scale_proto = out.File
	file_scale_proto_rawDesc = nil
	file_scale_proto_goTypes = nil
	file_scale_proto_depIdxs = nil
}

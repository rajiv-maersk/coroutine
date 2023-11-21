// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: coroutine/v1/type.proto

package coroutinev1

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

type Kind int32

const (
	Kind_KIND_UNSPECIFIED    Kind = 0
	Kind_KIND_NIL            Kind = 1
	Kind_KIND_BOOL           Kind = 2
	Kind_KIND_INT            Kind = 3
	Kind_KIND_INT8           Kind = 4
	Kind_KIND_INT16          Kind = 5
	Kind_KIND_INT32          Kind = 6
	Kind_KIND_INT64          Kind = 7
	Kind_KIND_UINT           Kind = 8
	Kind_KIND_UINT8          Kind = 9
	Kind_KIND_UINT16         Kind = 10
	Kind_KIND_UINT32         Kind = 11
	Kind_KIND_UINT64         Kind = 12
	Kind_KIND_UINTPTR        Kind = 13
	Kind_KIND_FLOAT32        Kind = 14
	Kind_KIND_FLOAT64        Kind = 15
	Kind_KIND_COMPLEX64      Kind = 16
	Kind_KIND_COMPLEX128     Kind = 17
	Kind_KIND_ARRAY          Kind = 18
	Kind_KIND_CHAN           Kind = 19
	Kind_KIND_FUNC           Kind = 20
	Kind_KIND_INTERFACE      Kind = 21
	Kind_KIND_MAP            Kind = 22
	Kind_KIND_POINTER        Kind = 23
	Kind_KIND_SLICE          Kind = 24
	Kind_KIND_STRING         Kind = 25
	Kind_KIND_STRUCT         Kind = 26
	Kind_KIND_UNSAFE_POINTER Kind = 27
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0:  "KIND_UNSPECIFIED",
		1:  "KIND_NIL",
		2:  "KIND_BOOL",
		3:  "KIND_INT",
		4:  "KIND_INT8",
		5:  "KIND_INT16",
		6:  "KIND_INT32",
		7:  "KIND_INT64",
		8:  "KIND_UINT",
		9:  "KIND_UINT8",
		10: "KIND_UINT16",
		11: "KIND_UINT32",
		12: "KIND_UINT64",
		13: "KIND_UINTPTR",
		14: "KIND_FLOAT32",
		15: "KIND_FLOAT64",
		16: "KIND_COMPLEX64",
		17: "KIND_COMPLEX128",
		18: "KIND_ARRAY",
		19: "KIND_CHAN",
		20: "KIND_FUNC",
		21: "KIND_INTERFACE",
		22: "KIND_MAP",
		23: "KIND_POINTER",
		24: "KIND_SLICE",
		25: "KIND_STRING",
		26: "KIND_STRUCT",
		27: "KIND_UNSAFE_POINTER",
	}
	Kind_value = map[string]int32{
		"KIND_UNSPECIFIED":    0,
		"KIND_NIL":            1,
		"KIND_BOOL":           2,
		"KIND_INT":            3,
		"KIND_INT8":           4,
		"KIND_INT16":          5,
		"KIND_INT32":          6,
		"KIND_INT64":          7,
		"KIND_UINT":           8,
		"KIND_UINT8":          9,
		"KIND_UINT16":         10,
		"KIND_UINT32":         11,
		"KIND_UINT64":         12,
		"KIND_UINTPTR":        13,
		"KIND_FLOAT32":        14,
		"KIND_FLOAT64":        15,
		"KIND_COMPLEX64":      16,
		"KIND_COMPLEX128":     17,
		"KIND_ARRAY":          18,
		"KIND_CHAN":           19,
		"KIND_FUNC":           20,
		"KIND_INTERFACE":      21,
		"KIND_MAP":            22,
		"KIND_POINTER":        23,
		"KIND_SLICE":          24,
		"KIND_STRING":         25,
		"KIND_STRUCT":         26,
		"KIND_UNSAFE_POINTER": 27,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_coroutine_v1_type_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_coroutine_v1_type_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_coroutine_v1_type_proto_rawDescGZIP(), []int{0}
}

type ChanDir int32

const (
	ChanDir_CHAN_DIR_UNSPECIFIED ChanDir = 0
	ChanDir_CHAN_DIR_RECV        ChanDir = 1
	ChanDir_CHAN_DIR_SEND        ChanDir = 2
	ChanDir_CHAN_DIR_BOTH        ChanDir = 3
)

// Enum value maps for ChanDir.
var (
	ChanDir_name = map[int32]string{
		0: "CHAN_DIR_UNSPECIFIED",
		1: "CHAN_DIR_RECV",
		2: "CHAN_DIR_SEND",
		3: "CHAN_DIR_BOTH",
	}
	ChanDir_value = map[string]int32{
		"CHAN_DIR_UNSPECIFIED": 0,
		"CHAN_DIR_RECV":        1,
		"CHAN_DIR_SEND":        2,
		"CHAN_DIR_BOTH":        3,
	}
)

func (x ChanDir) Enum() *ChanDir {
	p := new(ChanDir)
	*p = x
	return p
}

func (x ChanDir) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChanDir) Descriptor() protoreflect.EnumDescriptor {
	return file_coroutine_v1_type_proto_enumTypes[1].Descriptor()
}

func (ChanDir) Type() protoreflect.EnumType {
	return &file_coroutine_v1_type_proto_enumTypes[1]
}

func (x ChanDir) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChanDir.Descriptor instead.
func (ChanDir) EnumDescriptor() ([]byte, []int) {
	return file_coroutine_v1_type_proto_rawDescGZIP(), []int{1}
}

// Type is a data type.
//
// Types may reference other types internally (including themselves).
// To encode the graph of types in a program, a Type is expected to be
// stored alongside all other Types in an array. The index of each Type
// in the array becomes its unique identifier, and internal references
// to other types are simply an int32 index.
type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the string ID of the name of the type within the package it was
	// defined.
	Name int32 `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
	// Package is the string ID the name of the package that defines the type.
	Package int32 `protobuf:"varint,2,opt,name=package,proto3" json:"package,omitempty"`
	// Kind is the underlying type.
	Kind Kind `protobuf:"varint,3,opt,name=kind,proto3,enum=coroutine.v1.Kind" json:"kind,omitempty"`
	// Elem is the type of an array, slice, pointer, chan, or map's element.
	Elem int32 `protobuf:"varint,4,opt,name=elem,proto3" json:"elem,omitempty"`
	// Key is the key type for map types.
	Key int32 `protobuf:"varint,5,opt,name=key,proto3" json:"key,omitempty"`
	// Fields is the set of fields defined in a struct type.
	Fields []*Field `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	// Params are the types of params for a function type.
	Params []int32 `protobuf:"varint,7,rep,packed,name=params,proto3" json:"params,omitempty"`
	// Results are the types of results for a function type.
	Results []int32 `protobuf:"varint,8,rep,packed,name=results,proto3" json:"results,omitempty"`
	// Length is the length of an array type.
	Length int64 `protobuf:"varint,9,opt,name=length,proto3" json:"length,omitempty"`
	// MemoryOffset is an optional field that encodes the type's location
	// in memory.
	MemoryOffset uint64 `protobuf:"varint,10,opt,name=memory_offset,json=memoryOffset,proto3" json:"memory_offset,omitempty"`
	// ChanDir is the direction of a channel type.
	ChanDir ChanDir `protobuf:"varint,11,opt,name=chan_dir,json=chanDir,proto3,enum=coroutine.v1.ChanDir" json:"chan_dir,omitempty"`
	// Variadic is true for function types with a variadic argument.
	Variadic bool `protobuf:"varint,12,opt,name=variadic,proto3" json:"variadic,omitempty"`
	// Custom is true if the type is a custom opaque type.
	Custom bool `protobuf:"varint,13,opt,name=custom,proto3" json:"custom,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coroutine_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_coroutine_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_coroutine_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *Type) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *Type) GetPackage() int32 {
	if x != nil {
		return x.Package
	}
	return 0
}

func (x *Type) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_KIND_UNSPECIFIED
}

func (x *Type) GetElem() int32 {
	if x != nil {
		return x.Elem
	}
	return 0
}

func (x *Type) GetKey() int32 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *Type) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Type) GetParams() []int32 {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Type) GetResults() []int32 {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *Type) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Type) GetMemoryOffset() uint64 {
	if x != nil {
		return x.MemoryOffset
	}
	return 0
}

func (x *Type) GetChanDir() ChanDir {
	if x != nil {
		return x.ChanDir
	}
	return ChanDir_CHAN_DIR_UNSPECIFIED
}

func (x *Type) GetVariadic() bool {
	if x != nil {
		return x.Variadic
	}
	return false
}

func (x *Type) GetCustom() bool {
	if x != nil {
		return x.Custom
	}
	return false
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the string ID of the name of the field.
	Name int32 `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
	// Package is the string ID of package path that qualifies a lower case
	// (unexported) field name. It is empty for exported field names.
	Package int32 `protobuf:"varint,2,opt,name=package,proto3" json:"package,omitempty"`
	// Type is the type of the field.
	Type int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	// Offset is the offset of the field within its struct, in bytes.
	Offset uint64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	// Index is a sequence used by (reflect.Type).FieldByIndex.
	Index []int32 `protobuf:"varint,5,rep,packed,name=index,proto3" json:"index,omitempty"`
	// Anonymous indicates whether the field is an embedded field (with a name
	// derived from its type).
	Anonymous bool `protobuf:"varint,6,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	// Tag contains field metadata.
	Tag string `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coroutine_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_coroutine_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_coroutine_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *Field) GetPackage() int32 {
	if x != nil {
		return x.Package
	}
	return 0
}

func (x *Field) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Field) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Field) GetIndex() []int32 {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *Field) GetAnonymous() bool {
	if x != nil {
		return x.Anonymous
	}
	return false
}

func (x *Field) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_coroutine_v1_type_proto protoreflect.FileDescriptor

var file_coroutine_v1_type_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x84, 0x03, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x6e, 0x64,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e,
	0x5f, 0x64, 0x69, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x44, 0x69,
	0x72, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x44, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x64, 0x69, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x64, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x22, 0xa7,
	0x01, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x2a, 0xe4, 0x03, 0x0a, 0x04, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x4e, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x42, 0x4f,
	0x4f, 0x4c, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x38, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10,
	0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10,
	0x07, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x08,
	0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x09,
	0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10,
	0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32,
	0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36,
	0x34, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54,
	0x50, 0x54, 0x52, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x46, 0x4c,
	0x4f, 0x41, 0x54, 0x33, 0x32, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x36, 0x34, 0x10, 0x10, 0x12, 0x13, 0x0a,
	0x0f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x31, 0x32, 0x38,
	0x10, 0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59,
	0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x10,
	0x13, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x10, 0x14,
	0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x46, 0x41,
	0x43, 0x45, 0x10, 0x15, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4d, 0x41, 0x50,
	0x10, 0x16, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x10, 0x17, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x4c, 0x49,
	0x43, 0x45, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x10, 0x19, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x54,
	0x52, 0x55, 0x43, 0x54, 0x10, 0x1a, 0x12, 0x17, 0x0a, 0x13, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55,
	0x4e, 0x53, 0x41, 0x46, 0x45, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x1b, 0x2a,
	0x5c, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x44, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x48,
	0x41, 0x4e, 0x5f, 0x44, 0x49, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x4e, 0x5f, 0x44, 0x49, 0x52,
	0x5f, 0x52, 0x45, 0x43, 0x56, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x4e, 0x5f,
	0x44, 0x49, 0x52, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48,
	0x41, 0x4e, 0x5f, 0x44, 0x49, 0x52, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x03, 0x42, 0xb8, 0x01,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa,
	0x02, 0x0c, 0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0c, 0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18,
	0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x6f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coroutine_v1_type_proto_rawDescOnce sync.Once
	file_coroutine_v1_type_proto_rawDescData = file_coroutine_v1_type_proto_rawDesc
)

func file_coroutine_v1_type_proto_rawDescGZIP() []byte {
	file_coroutine_v1_type_proto_rawDescOnce.Do(func() {
		file_coroutine_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_coroutine_v1_type_proto_rawDescData)
	})
	return file_coroutine_v1_type_proto_rawDescData
}

var file_coroutine_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_coroutine_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coroutine_v1_type_proto_goTypes = []interface{}{
	(Kind)(0),     // 0: coroutine.v1.Kind
	(ChanDir)(0),  // 1: coroutine.v1.ChanDir
	(*Type)(nil),  // 2: coroutine.v1.Type
	(*Field)(nil), // 3: coroutine.v1.Field
}
var file_coroutine_v1_type_proto_depIdxs = []int32{
	0, // 0: coroutine.v1.Type.kind:type_name -> coroutine.v1.Kind
	3, // 1: coroutine.v1.Type.fields:type_name -> coroutine.v1.Field
	1, // 2: coroutine.v1.Type.chan_dir:type_name -> coroutine.v1.ChanDir
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coroutine_v1_type_proto_init() }
func file_coroutine_v1_type_proto_init() {
	if File_coroutine_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coroutine_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_coroutine_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
			RawDescriptor: file_coroutine_v1_type_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coroutine_v1_type_proto_goTypes,
		DependencyIndexes: file_coroutine_v1_type_proto_depIdxs,
		EnumInfos:         file_coroutine_v1_type_proto_enumTypes,
		MessageInfos:      file_coroutine_v1_type_proto_msgTypes,
	}.Build()
	File_coroutine_v1_type_proto = out.File
	file_coroutine_v1_type_proto_rawDesc = nil
	file_coroutine_v1_type_proto_goTypes = nil
	file_coroutine_v1_type_proto_depIdxs = nil
}

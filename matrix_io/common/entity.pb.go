// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package matrix_io_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An array value
type ArrayValue struct {
	// Values in the array.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{0}
}

func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (m *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(m, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// Container of the value of an object message
type Value struct {
	// Must have a value set
	//
	// Types that are valid to be assigned to ValueType:
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_EntityValue
	//	*Value_ArrayValue
	//	*Value_TimestampValue
	//	*Value_StringValue
	//	*Value_BlobValue
	//	*Value_SizeValue
	ValueType            isValue_ValueType `protobuf_oneof:"value_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_EntityValue struct {
	EntityValue *Entity `protobuf:"bytes,6,opt,name=entity_value,json=entityValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BlobValue struct {
	BlobValue []byte `protobuf:"bytes,18,opt,name=blob_value,json=blobValue,proto3,oneof"`
}

type Value_SizeValue struct {
	SizeValue *Value_Size `protobuf:"bytes,19,opt,name=size_value,json=sizeValue,proto3,oneof"`
}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_EntityValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BlobValue) isValue_ValueType() {}

func (*Value_SizeValue) isValue_ValueType() {}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetEntityValue() *Entity {
	if x, ok := m.GetValueType().(*Value_EntityValue); ok {
		return x.EntityValue
	}
	return nil
}

func (m *Value) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (m *Value) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBlobValue() []byte {
	if x, ok := m.GetValueType().(*Value_BlobValue); ok {
		return x.BlobValue
	}
	return nil
}

func (m *Value) GetSizeValue() *Value_Size {
	if x, ok := m.GetValueType().(*Value_SizeValue); ok {
		return x.SizeValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_EntityValue)(nil),
		(*Value_ArrayValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_SizeValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_IntegerValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_EntityValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EntityValue); err != nil {
			return err
		}
	case *Value_ArrayValue:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ArrayValue); err != nil {
			return err
		}
	case *Value_TimestampValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case *Value_StringValue:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BlobValue:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BlobValue)
	case *Value_SizeValue:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SizeValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.ValueType has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // value_type.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_BooleanValue{x != 0}
		return true, err
	case 2: // value_type.integer_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_IntegerValue{int64(x)}
		return true, err
	case 3: // value_type.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ValueType = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 6: // value_type.entity_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Entity)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_EntityValue{msg}
		return true, err
	case 9: // value_type.array_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayValue)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_ArrayValue{msg}
		return true, err
	case 10: // value_type.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_TimestampValue{msg}
		return true, err
	case 17: // value_type.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueType = &Value_StringValue{x}
		return true, err
	case 18: // value_type.blob_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ValueType = &Value_BlobValue{x}
		return true, err
	case 19: // value_type.size_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value_Size)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_SizeValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_BooleanValue:
		n += 1 // tag and wire
		n += 1
	case *Value_IntegerValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_EntityValue:
		s := proto.Size(x.EntityValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ArrayValue:
		s := proto.Size(x.ArrayValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_StringValue:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BlobValue:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BlobValue)))
		n += len(x.BlobValue)
	case *Value_SizeValue:
		s := proto.Size(x.SizeValue)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Value_Size struct {
	Width                int64    `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               int64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value_Size) Reset()         { *m = Value_Size{} }
func (m *Value_Size) String() string { return proto.CompactTextString(m) }
func (*Value_Size) ProtoMessage()    {}
func (*Value_Size) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{1, 0}
}

func (m *Value_Size) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value_Size.Unmarshal(m, b)
}
func (m *Value_Size) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value_Size.Marshal(b, m, deterministic)
}
func (m *Value_Size) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value_Size.Merge(m, src)
}
func (m *Value_Size) XXX_Size() int {
	return xxx_messageInfo_Value_Size.Size(m)
}
func (m *Value_Size) XXX_DiscardUnknown() {
	xxx_messageInfo_Value_Size.DiscardUnknown(m)
}

var xxx_messageInfo_Value_Size proto.InternalMessageInfo

func (m *Value_Size) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Value_Size) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// Object message that can contain object structures defined at
// runtime
type Entity struct {
	// Collection of named object fields and their values
	// The name must not contain more than 60 characters
	// the name cannot be empty
	Properties           map[string]*Value `protobuf:"bytes,1,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{2}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*ArrayValue)(nil), "matrix_io.common.ArrayValue")
	proto.RegisterType((*Value)(nil), "matrix_io.common.Value")
	proto.RegisterType((*Value_Size)(nil), "matrix_io.common.Value.Size")
	proto.RegisterType((*Entity)(nil), "matrix_io.common.Entity")
	proto.RegisterMapType((map[string]*Value)(nil), "matrix_io.common.Entity.PropertiesEntry")
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor_cf50d946d740d100) }

var fileDescriptor_cf50d946d740d100 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xe3, 0x66, 0x8d, 0xc8, 0x49, 0xa0, 0x23, 0x43, 0x50, 0x55, 0x48, 0x8b, 0x8a, 0x90,
	0x72, 0x83, 0x2b, 0x0d, 0x2e, 0x10, 0xd2, 0x40, 0xeb, 0xa8, 0x14, 0x2e, 0x26, 0x2a, 0x83, 0x76,
	0x5b, 0x25, 0xcc, 0xa4, 0x16, 0x49, 0x1c, 0xb9, 0x2e, 0x90, 0x3d, 0x02, 0x8f, 0xc1, 0xe5, 0x9e,
	0x12, 0xf9, 0x4f, 0x52, 0xb4, 0xa9, 0xbb, 0x8b, 0x3f, 0xff, 0xbe, 0x93, 0x73, 0xbe, 0x63, 0x08,
	0x69, 0x2d, 0x99, 0x6c, 0x71, 0x23, 0xb8, 0xe4, 0xd1, 0x61, 0x95, 0x49, 0xc1, 0x7e, 0xaf, 0x18,
	0xc7, 0xdf, 0x78, 0x55, 0xf1, 0x7a, 0x72, 0x5c, 0x70, 0x5e, 0x94, 0x74, 0xa6, 0xef, 0xf3, 0xed,
	0xf7, 0x99, 0x64, 0x15, 0xdd, 0xc8, 0xac, 0x6a, 0x8c, 0x65, 0x7a, 0x0a, 0x70, 0x26, 0x44, 0xd6,
	0x5e, 0x66, 0xe5, 0x96, 0x46, 0x33, 0xf0, 0x7e, 0xaa, 0x8f, 0xcd, 0x18, 0xc5, 0x6e, 0x12, 0x9c,
	0x3c, 0xc3, 0xb7, 0x2b, 0x62, 0x0d, 0x12, 0x8b, 0x4d, 0xff, 0x1c, 0xc0, 0xd0, 0x58, 0x5f, 0xc2,
	0xc3, 0x9c, 0xf3, 0x92, 0x66, 0xf5, 0x4a, 0xdf, 0x8d, 0x51, 0x8c, 0x92, 0x07, 0xa9, 0x43, 0x42,
	0x2b, 0xf7, 0x18, 0xab, 0x25, 0x2d, 0xa8, 0xb0, 0xd8, 0x20, 0x46, 0x89, 0xab, 0x30, 0x2b, 0x1b,
	0xec, 0x05, 0x84, 0x57, 0x7c, 0x9b, 0x97, 0xd4, 0x52, 0x6e, 0x8c, 0x12, 0x94, 0x3a, 0x24, 0x30,
	0xaa, 0x81, 0x4e, 0xbb, 0xf1, 0x2d, 0xe4, 0xc5, 0x28, 0x09, 0x4e, 0xc6, 0x77, 0x7b, 0x5e, 0x68,
	0x4a, 0xd9, 0x0d, 0x6f, 0xec, 0x1f, 0x20, 0xc8, 0xd4, 0xe8, 0xd6, 0xed, 0x6b, 0xf7, 0xf3, 0xbb,
	0xee, 0x5d, 0x3e, 0xa9, 0x43, 0x20, 0xdb, 0xa5, 0xb5, 0x80, 0x51, 0x1f, 0xa7, 0x2d, 0x02, 0xba,
	0xc8, 0x04, 0x9b, 0xd8, 0x71, 0x17, 0x3b, 0xfe, 0xda, 0x71, 0xa9, 0x43, 0x1e, 0xf5, 0xa6, 0x7e,
	0xd6, 0x8d, 0x14, 0xac, 0x2e, 0x6c, 0x8d, 0xc7, 0x31, 0x4a, 0x7c, 0xd5, 0xac, 0x51, 0x0d, 0x74,
	0x0c, 0x90, 0x97, 0x3c, 0xb7, 0x48, 0x14, 0xa3, 0x24, 0x4c, 0x1d, 0xe2, 0x2b, 0xad, 0x0b, 0x03,
	0x36, 0xec, 0xba, 0xcb, 0xeb, 0x68, 0xdf, 0x30, 0x1a, 0xc6, 0x5f, 0xd8, 0xb5, 0x1a, 0xc6, 0x57,
	0x0e, 0xad, 0x4c, 0xde, 0xc0, 0x81, 0x12, 0xa3, 0x27, 0x30, 0xfc, 0xc5, 0xae, 0xe4, 0x5a, 0xaf,
	0xcf, 0x25, 0xe6, 0x10, 0x3d, 0x05, 0x6f, 0x4d, 0x59, 0xb1, 0x96, 0x66, 0x5d, 0xc4, 0x9e, 0xe6,
	0x21, 0x80, 0xfe, 0xdf, 0x4a, 0xb6, 0x0d, 0x9d, 0xde, 0x20, 0xf0, 0x4c, 0xd4, 0x51, 0x0a, 0xd0,
	0x08, 0xde, 0x50, 0x21, 0x59, 0xff, 0x98, 0x92, 0x7d, 0x8b, 0xc1, 0xcb, 0x1e, 0x5d, 0xd4, 0x52,
	0xb4, 0xe4, 0x3f, 0xef, 0xe4, 0x12, 0x46, 0xb7, 0xae, 0xa3, 0x43, 0x70, 0x7f, 0xd0, 0x56, 0x77,
	0xe8, 0x13, 0xf5, 0x19, 0xbd, 0x82, 0xe1, 0xee, 0x35, 0xdd, 0xf3, 0x6c, 0x0d, 0xf5, 0x6e, 0xf0,
	0x16, 0xcd, 0xdf, 0xc3, 0x11, 0xaf, 0xa9, 0x05, 0x7b, 0x6e, 0x1e, 0x98, 0x96, 0x96, 0x6a, 0x71,
	0x4b, 0xf4, 0x77, 0x30, 0x3c, 0xfb, 0x78, 0x71, 0xbe, 0xb8, 0x19, 0x8c, 0x2e, 0x34, 0xf8, 0xe9,
	0x33, 0x3e, 0xd7, 0x60, 0xee, 0xe9, 0xdd, 0xbe, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x60, 0xec,
	0x2f, 0xfe, 0x82, 0x03, 0x00, 0x00,
}

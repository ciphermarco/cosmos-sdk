// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/orm/v1/orm.proto

package ormv1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TableDescriptor describes an ORM table.
type TableDescriptor struct {
	// primary_key defines the primary key for the table.
	PrimaryKey *PrimaryKeyDescriptor `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	// index defines one or more secondary indexes.
	Index []*SecondaryIndexDescriptor `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
	// id is a non-zero integer ID that must be unique within the
	// tables and singletons in this file. It may be deprecated in the future when this
	// can be auto-generated.
	Id uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *TableDescriptor) Reset()         { *m = TableDescriptor{} }
func (m *TableDescriptor) String() string { return proto.CompactTextString(m) }
func (*TableDescriptor) ProtoMessage()    {}
func (*TableDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_552545c04761d6eb, []int{0}
}
func (m *TableDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TableDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TableDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableDescriptor.Merge(m, src)
}
func (m *TableDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *TableDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TableDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TableDescriptor proto.InternalMessageInfo

func (m *TableDescriptor) GetPrimaryKey() *PrimaryKeyDescriptor {
	if m != nil {
		return m.PrimaryKey
	}
	return nil
}

func (m *TableDescriptor) GetIndex() []*SecondaryIndexDescriptor {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *TableDescriptor) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// PrimaryKeyDescriptor describes a table primary key.
type PrimaryKeyDescriptor struct {
	// fields is a comma-separated list of fields in the primary key. Spaces are
	// not allowed. Supported field types, their encodings, and any applicable constraints
	// are described below.
	//   - uint32 are encoded as 2,3,4 or 5 bytes using a compact encoding that
	//     is suitable for sorted iteration (not varint encoding). This type is
	//     well-suited for small integers.
	//   - uint64 are encoded as 2,4,6 or 9 bytes using a compact encoding that
	//     is suitable for sorted iteration (not varint encoding). This type is
	//     well-suited for small integers such as auto-incrementing sequences.
	//   - fixed32, fixed64 are encoded as big-endian fixed width bytes and support
	//     sorted iteration. These types are well-suited for encoding fixed with
	//     decimals as integers.
	//   - string's are encoded as raw bytes in terminal key segments and null-terminated
	//     in non-terminal segments. Null characters are thus forbidden in strings.
	//     string fields support sorted iteration.
	//   - bytes are encoded as raw bytes in terminal segments and length-prefixed
	//     with a 32-bit unsigned varint in non-terminal segments.
	//   - int32, sint32, int64, sint64, sfixed32, sfixed64 are encoded as fixed width bytes with
	//     an encoding that enables sorted iteration.
	//   - google.protobuf.Timestamp is encoded such that values with only seconds occupy 6 bytes,
	//     values including nanos occupy 9 bytes, and nil values occupy 1 byte. When iterating, nil
	//     values will always be ordered last. Seconds and nanos values must conform to the officially
	//     specified ranges of 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z and 0 to 999,999,999 respectively.
	//   - google.protobuf.Duration is encoded as 12 bytes using an encoding that enables sorted iteration.
	//   - enum fields are encoded using varint encoding and do not support sorted
	//     iteration.
	//   - bool fields are encoded as a single byte 0 or 1.
	//
	// All other fields types are unsupported in keys including repeated and
	// oneof fields.
	//
	// Primary keys are prefixed by the varint encoded table id and the byte 0x0
	// plus any additional prefix specified by the schema.
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	// auto_increment specifies that the primary key is generated by an
	// auto-incrementing integer. If this is set to true fields must only
	// contain one field of that is of type uint64.
	AutoIncrement bool `protobuf:"varint,2,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
}

func (m *PrimaryKeyDescriptor) Reset()         { *m = PrimaryKeyDescriptor{} }
func (m *PrimaryKeyDescriptor) String() string { return proto.CompactTextString(m) }
func (*PrimaryKeyDescriptor) ProtoMessage()    {}
func (*PrimaryKeyDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_552545c04761d6eb, []int{1}
}
func (m *PrimaryKeyDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimaryKeyDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimaryKeyDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimaryKeyDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimaryKeyDescriptor.Merge(m, src)
}
func (m *PrimaryKeyDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *PrimaryKeyDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimaryKeyDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_PrimaryKeyDescriptor proto.InternalMessageInfo

func (m *PrimaryKeyDescriptor) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

func (m *PrimaryKeyDescriptor) GetAutoIncrement() bool {
	if m != nil {
		return m.AutoIncrement
	}
	return false
}

// PrimaryKeyDescriptor describes a table secondary index.
type SecondaryIndexDescriptor struct {
	// fields is a comma-separated list of fields in the index. The supported
	// field types are the same as those for PrimaryKeyDescriptor.fields.
	// Index keys are prefixed by the varint encoded table id and the varint
	// encoded index id plus any additional prefix specified by the schema.
	//
	// In addition the field segments, non-unique index keys are suffixed with
	// any additional primary key fields not present in the index fields so that the
	// primary key can be reconstructed. Unique indexes instead of being suffixed
	// store the remaining primary key fields in the value..
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	// id is a non-zero integer ID that must be unique within the indexes for this
	// table and less than 32768. It may be deprecated in the future when this can
	// be auto-generated.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// unique specifies that this an unique index.
	Unique bool `protobuf:"varint,3,opt,name=unique,proto3" json:"unique,omitempty"`
}

func (m *SecondaryIndexDescriptor) Reset()         { *m = SecondaryIndexDescriptor{} }
func (m *SecondaryIndexDescriptor) String() string { return proto.CompactTextString(m) }
func (*SecondaryIndexDescriptor) ProtoMessage()    {}
func (*SecondaryIndexDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_552545c04761d6eb, []int{2}
}
func (m *SecondaryIndexDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecondaryIndexDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecondaryIndexDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecondaryIndexDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondaryIndexDescriptor.Merge(m, src)
}
func (m *SecondaryIndexDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *SecondaryIndexDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondaryIndexDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SecondaryIndexDescriptor proto.InternalMessageInfo

func (m *SecondaryIndexDescriptor) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

func (m *SecondaryIndexDescriptor) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SecondaryIndexDescriptor) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

// TableDescriptor describes an ORM singleton table which has at most one instance.
type SingletonDescriptor struct {
	// id is a non-zero integer ID that must be unique within the
	// tables and singletons in this file. It may be deprecated in the future when this
	// can be auto-generated.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *SingletonDescriptor) Reset()         { *m = SingletonDescriptor{} }
func (m *SingletonDescriptor) String() string { return proto.CompactTextString(m) }
func (*SingletonDescriptor) ProtoMessage()    {}
func (*SingletonDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_552545c04761d6eb, []int{3}
}
func (m *SingletonDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingletonDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingletonDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SingletonDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingletonDescriptor.Merge(m, src)
}
func (m *SingletonDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *SingletonDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SingletonDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SingletonDescriptor proto.InternalMessageInfo

func (m *SingletonDescriptor) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

var E_Table = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*TableDescriptor)(nil),
	Field:         104503790,
	Name:          "cosmos.orm.v1.table",
	Tag:           "bytes,104503790,opt,name=table",
	Filename:      "cosmos/orm/v1/orm.proto",
}

var E_Singleton = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*SingletonDescriptor)(nil),
	Field:         104503791,
	Name:          "cosmos.orm.v1.singleton",
	Tag:           "bytes,104503791,opt,name=singleton",
	Filename:      "cosmos/orm/v1/orm.proto",
}

func init() {
	proto.RegisterType((*TableDescriptor)(nil), "cosmos.orm.v1.TableDescriptor")
	proto.RegisterType((*PrimaryKeyDescriptor)(nil), "cosmos.orm.v1.PrimaryKeyDescriptor")
	proto.RegisterType((*SecondaryIndexDescriptor)(nil), "cosmos.orm.v1.SecondaryIndexDescriptor")
	proto.RegisterType((*SingletonDescriptor)(nil), "cosmos.orm.v1.SingletonDescriptor")
	proto.RegisterExtension(E_Table)
	proto.RegisterExtension(E_Singleton)
}

func init() { proto.RegisterFile("cosmos/orm/v1/orm.proto", fileDescriptor_552545c04761d6eb) }

var fileDescriptor_552545c04761d6eb = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x6b, 0x57, 0x9b, 0x3a, 0x57, 0xdd, 0x44, 0x40, 0x10, 0x38, 0x84, 0x2a, 0x30, 0xd1,
	0x93, 0xa3, 0x8c, 0x5b, 0x10, 0x97, 0x6e, 0x12, 0x9a, 0xd0, 0x94, 0x2a, 0x83, 0x6a, 0x9a, 0x2a,
	0x4d, 0x6e, 0xec, 0x55, 0xd6, 0xea, 0x38, 0xd8, 0x6e, 0x45, 0xdf, 0x82, 0x67, 0xe0, 0xc0, 0x81,
	0x1b, 0xe2, 0x25, 0x10, 0xa7, 0x1d, 0x39, 0xa2, 0xf6, 0xb6, 0x03, 0xbc, 0x02, 0x4a, 0xbc, 0x6c,
	0x6b, 0x61, 0xe2, 0x14, 0xfd, 0xff, 0x9f, 0xbf, 0x5f, 0xbe, 0xf8, 0x0b, 0x7a, 0x90, 0x4a, 0x2d,
	0xa4, 0x0e, 0xa4, 0x12, 0xc1, 0x34, 0x2c, 0x1e, 0x38, 0x57, 0xd2, 0x48, 0xa7, 0x65, 0x05, 0x5c,
	0x6c, 0xa6, 0xe1, 0xa3, 0xf6, 0x48, 0xca, 0xd1, 0x98, 0x05, 0xa5, 0x38, 0x9c, 0x9c, 0x06, 0x94,
	0xe9, 0x54, 0xf1, 0xdc, 0x48, 0x65, 0x0d, 0xfe, 0x27, 0x80, 0xb6, 0xde, 0x90, 0xe1, 0x98, 0xed,
	0x5d, 0x29, 0xce, 0x1e, 0x6a, 0xe6, 0x8a, 0x0b, 0xa2, 0x66, 0x27, 0x67, 0x6c, 0xe6, 0x82, 0x36,
	0xe8, 0x34, 0x77, 0x9e, 0xe0, 0x25, 0x34, 0xee, 0xd9, 0x13, 0xaf, 0xd9, 0xec, 0xda, 0x99, 0xa0,
	0xfc, 0x6a, 0xeb, 0xbc, 0x44, 0x6b, 0x3c, 0xa3, 0xec, 0xbd, 0x0b, 0xdb, 0xf5, 0x4e, 0x73, 0xe7,
	0xd9, 0x8a, 0xff, 0x90, 0xa5, 0x32, 0xa3, 0x44, 0xcd, 0xf6, 0x8b, 0x43, 0x37, 0x18, 0xd6, 0xe5,
	0x6c, 0x22, 0xc8, 0xa9, 0x5b, 0x6f, 0x83, 0x4e, 0x2b, 0x81, 0x9c, 0xfa, 0x6f, 0xd1, 0xbd, 0x7f,
	0xbd, 0xd2, 0xb9, 0x8f, 0xd6, 0x4f, 0x39, 0x1b, 0x53, 0x5d, 0xe6, 0xdc, 0x48, 0x2e, 0x27, 0x67,
	0x1b, 0x6d, 0x92, 0x89, 0x91, 0x27, 0x3c, 0x4b, 0x15, 0x13, 0x2c, 0x33, 0x2e, 0x6c, 0x83, 0x4e,
	0x23, 0x69, 0x15, 0xdb, 0xfd, 0x6a, 0xe9, 0x1f, 0x23, 0xf7, 0xb6, 0x24, 0xb7, 0xa2, 0x6d, 0x34,
	0x58, 0x45, 0x2b, 0xce, 0x4d, 0x32, 0xfe, 0x6e, 0xc2, 0xca, 0xb8, 0x8d, 0xe4, 0x72, 0xf2, 0xb7,
	0xd1, 0xdd, 0x43, 0x9e, 0x8d, 0xc6, 0xcc, 0xc8, 0xec, 0x06, 0xd6, 0xda, 0x41, 0x65, 0x8f, 0x8e,
	0xd0, 0x9a, 0x29, 0x1a, 0x70, 0x1e, 0x63, 0x5b, 0x17, 0xae, 0xea, 0xc2, 0x07, 0x4c, 0x6b, 0x32,
	0x62, 0x71, 0x6e, 0xb8, 0xcc, 0xb4, 0xfb, 0xeb, 0xeb, 0x45, 0x58, 0x96, 0xe1, 0xad, 0x5c, 0xe6,
	0x4a, 0x83, 0x89, 0x05, 0x46, 0x14, 0x6d, 0xe8, 0x2a, 0xc0, 0xff, 0xe9, 0xbf, 0x2b, 0xba, 0xbf,
	0x5a, 0xd5, 0xdf, 0x1f, 0x91, 0x5c, 0x83, 0xbb, 0x5f, 0xc0, 0xb7, 0xb9, 0x07, 0xce, 0xe7, 0x1e,
	0xf8, 0x39, 0xf7, 0xc0, 0x87, 0x85, 0x57, 0x3b, 0x5f, 0x78, 0xb5, 0x1f, 0x0b, 0xaf, 0x86, 0xee,
	0xa4, 0x52, 0x2c, 0xc3, 0xba, 0x8d, 0x58, 0x89, 0x5e, 0x91, 0xa0, 0x07, 0x8e, 0x9f, 0x5a, 0x49,
	0xd3, 0x33, 0xcc, 0x65, 0x40, 0x72, 0x1e, 0x2c, 0xfd, 0xd7, 0x2f, 0xa4, 0x12, 0xd3, 0xf0, 0x23,
	0xac, 0xef, 0xc6, 0x47, 0x9f, 0x61, 0x6b, 0xd7, 0x72, 0x62, 0x25, 0x70, 0x3f, 0xfc, 0x5e, 0xcd,
	0x83, 0x58, 0x89, 0x41, 0x3f, 0x9c, 0xc3, 0x87, 0x4b, 0xf3, 0xe0, 0x55, 0xaf, 0x7b, 0xc0, 0x0c,
	0xa1, 0xc4, 0x90, 0x0b, 0xb8, 0x65, 0xb5, 0x28, 0x8a, 0x95, 0x88, 0xa2, 0x7e, 0x38, 0x5c, 0x2f,
	0x2f, 0xe1, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xc2, 0xdc, 0x47, 0x49, 0x03, 0x00,
	0x00,
}

func (m *TableDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TableDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintOrm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		for iNdEx := len(m.Index) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Index[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOrm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PrimaryKey != nil {
		{
			size, err := m.PrimaryKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PrimaryKeyDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimaryKeyDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimaryKeyDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AutoIncrement {
		i--
		if m.AutoIncrement {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Fields) > 0 {
		i -= len(m.Fields)
		copy(dAtA[i:], m.Fields)
		i = encodeVarintOrm(dAtA, i, uint64(len(m.Fields)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SecondaryIndexDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecondaryIndexDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecondaryIndexDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Unique {
		i--
		if m.Unique {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintOrm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Fields) > 0 {
		i -= len(m.Fields)
		copy(dAtA[i:], m.Fields)
		i = encodeVarintOrm(dAtA, i, uint64(len(m.Fields)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SingletonDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingletonDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SingletonDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintOrm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrm(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TableDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PrimaryKey != nil {
		l = m.PrimaryKey.Size()
		n += 1 + l + sovOrm(uint64(l))
	}
	if len(m.Index) > 0 {
		for _, e := range m.Index {
			l = e.Size()
			n += 1 + l + sovOrm(uint64(l))
		}
	}
	if m.Id != 0 {
		n += 1 + sovOrm(uint64(m.Id))
	}
	return n
}

func (m *PrimaryKeyDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fields)
	if l > 0 {
		n += 1 + l + sovOrm(uint64(l))
	}
	if m.AutoIncrement {
		n += 2
	}
	return n
}

func (m *SecondaryIndexDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fields)
	if l > 0 {
		n += 1 + l + sovOrm(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovOrm(uint64(m.Id))
	}
	if m.Unique {
		n += 2
	}
	return n
}

func (m *SingletonDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovOrm(uint64(m.Id))
	}
	return n
}

func sovOrm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrm(x uint64) (n int) {
	return sovOrm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrimaryKey == nil {
				m.PrimaryKey = &PrimaryKeyDescriptor{}
			}
			if err := m.PrimaryKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = append(m.Index, &SecondaryIndexDescriptor{})
			if err := m.Index[len(m.Index)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PrimaryKeyDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrimaryKeyDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimaryKeyDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoIncrement", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoIncrement = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOrm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SecondaryIndexDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SecondaryIndexDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecondaryIndexDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Unique = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOrm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SingletonDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SingletonDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingletonDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOrm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrm
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOrm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrm = fmt.Errorf("proto: unexpected end of group")
)
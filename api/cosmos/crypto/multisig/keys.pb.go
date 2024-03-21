// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crypto/multisig/keys.proto

package multisig

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// LegacyAminoPubKey specifies a public key type
// which nests multiple public keys and a threshold,
// it uses legacy amino address rules.
type LegacyAminoPubKey struct {
	Threshold uint32       `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	PubKeys   []*anypb.Any `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
}

func (m *LegacyAminoPubKey) Reset()         { *m = LegacyAminoPubKey{} }
func (m *LegacyAminoPubKey) String() string { return proto.CompactTextString(m) }
func (*LegacyAminoPubKey) ProtoMessage()    {}
func (*LegacyAminoPubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b57537e097d47d, []int{0}
}
func (m *LegacyAminoPubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyAminoPubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyAminoPubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyAminoPubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyAminoPubKey.Merge(m, src)
}
func (m *LegacyAminoPubKey) XXX_Size() int {
	return m.Size()
}
func (m *LegacyAminoPubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyAminoPubKey.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyAminoPubKey proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LegacyAminoPubKey)(nil), "cosmos.crypto.multisig.LegacyAminoPubKey")
}

func init() { proto.RegisterFile("cosmos/crypto/multisig/keys.proto", fileDescriptor_46b57537e097d47d) }

var fileDescriptor_46b57537e097d47d = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0xcf, 0x2d, 0xcd, 0x29, 0xc9, 0x2c,
	0xce, 0x4c, 0xd7, 0xcf, 0x4e, 0xad, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x83,
	0x28, 0xd1, 0x83, 0x28, 0xd1, 0x83, 0x29, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd1,
	0x07, 0xb1, 0x20, 0xaa, 0xa5, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4,
	0xd2, 0x34, 0xfd, 0xc4, 0xbc, 0x4a, 0xa8, 0x94, 0x60, 0x62, 0x6e, 0x66, 0x5e, 0xbe, 0x3e, 0x98,
	0x84, 0x08, 0x29, 0x1d, 0x66, 0xe4, 0x12, 0xf4, 0x49, 0x4d, 0x4f, 0x4c, 0xae, 0x74, 0x04, 0x89,
	0x06, 0x94, 0x26, 0x79, 0xa7, 0x56, 0x0a, 0xc9, 0x70, 0x71, 0x96, 0x64, 0x14, 0xa5, 0x16, 0x67,
	0xe4, 0xe7, 0xa4, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x21, 0x04, 0x84, 0xfc, 0xb8, 0xb8,
	0x0b, 0x4a, 0x93, 0x72, 0x32, 0x93, 0xe3, 0x41, 0x8e, 0x94, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36,
	0x12, 0xd1, 0x83, 0xd8, 0xab, 0x07, 0xb3, 0x57, 0xcf, 0x31, 0xaf, 0xd2, 0x49, 0xfc, 0xd1, 0x3d,
	0x79, 0x76, 0x88, 0xa1, 0xc5, 0x8b, 0x9e, 0x6f, 0xd0, 0x62, 0x2f, 0x28, 0x4d, 0x02, 0x69, 0x0a,
	0xe2, 0x82, 0x98, 0x00, 0x12, 0xb7, 0x72, 0xe8, 0x58, 0x20, 0xcf, 0xd0, 0xf5, 0x7c, 0x83, 0x96,
	0x52, 0x49, 0x6a, 0x5e, 0x4a, 0x6a, 0x51, 0x6e, 0x66, 0x5e, 0x89, 0x3e, 0x44, 0x93, 0x2f, 0xd4,
	0xaf, 0x21, 0x30, 0xcb, 0x27, 0x3d, 0xdf, 0xa0, 0x25, 0x00, 0x77, 0x4a, 0x7c, 0x71, 0x49, 0x51,
	0x66, 0x5e, 0xba, 0xd3, 0x23, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0,
	0x92, 0x4a, 0xce, 0xcf, 0xd5, 0xc3, 0x1e, 0x80, 0x4e, 0x9c, 0x20, 0xeb, 0x03, 0x40, 0xee, 0x0d,
	0x60, 0x8c, 0x52, 0x87, 0x28, 0x2a, 0x4e, 0xc9, 0xd6, 0xcb, 0xcc, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4,
	0xc7, 0x1e, 0x33, 0x8b, 0x98, 0x98, 0x9d, 0x9d, 0x7d, 0x57, 0x31, 0x89, 0x39, 0x43, 0x0c, 0x75,
	0x86, 0x18, 0x0a, 0x73, 0xe9, 0x29, 0x98, 0x44, 0x0c, 0x44, 0x22, 0x06, 0x26, 0xf1, 0x88, 0x49,
	0x09, 0xbb, 0x44, 0x8c, 0x7b, 0x80, 0x93, 0x6f, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x49, 0xe2, 0x2b,
	0x26, 0x09, 0x88, 0x22, 0x2b, 0x2b, 0x88, 0x2a, 0x2b, 0x2b, 0x98, 0xb2, 0x24, 0x36, 0x70, 0xc8,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x7d, 0x61, 0x5a, 0x32, 0x02, 0x00, 0x00,
}

func (m *LegacyAminoPubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyAminoPubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyAminoPubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PubKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintKeys(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintKeys(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LegacyAminoPubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Threshold != 0 {
		n += 1 + sovKeys(uint64(m.Threshold))
	}
	if len(m.PubKeys) > 0 {
		for _, e := range m.PubKeys {
			l = e.Size()
			n += 1 + l + sovKeys(uint64(l))
		}
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyAminoPubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: LegacyAminoPubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyAminoPubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, &anypb.Any{})
			if err := m.PubKeys[len(m.PubKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeys
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
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeys = fmt.Errorf("proto: unexpected end of group")
)
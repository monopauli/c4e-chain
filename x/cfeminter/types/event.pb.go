// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfeminter/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Mint struct {
	BondedRatio string `protobuf:"bytes,1,opt,name=bonded_ratio,json=bondedRatio,proto3" json:"bonded_ratio,omitempty"`
	Inflation   string `protobuf:"bytes,2,opt,name=inflation,proto3" json:"inflation,omitempty"`
	Amount      string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Mint) Reset()         { *m = Mint{} }
func (m *Mint) String() string { return proto.CompactTextString(m) }
func (*Mint) ProtoMessage()    {}
func (*Mint) Descriptor() ([]byte, []int) {
	return fileDescriptor_b98ea54e6d90ae85, []int{0}
}
func (m *Mint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mint.Merge(m, src)
}
func (m *Mint) XXX_Size() int {
	return m.Size()
}
func (m *Mint) XXX_DiscardUnknown() {
	xxx_messageInfo_Mint.DiscardUnknown(m)
}

var xxx_messageInfo_Mint proto.InternalMessageInfo

func (m *Mint) GetBondedRatio() string {
	if m != nil {
		return m.BondedRatio
	}
	return ""
}

func (m *Mint) GetInflation() string {
	if m != nil {
		return m.Inflation
	}
	return ""
}

func (m *Mint) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*Mint)(nil), "chain4energy.c4echain.cfeminter.Mint")
}

func init() { proto.RegisterFile("cfeminter/event.proto", fileDescriptor_b98ea54e6d90ae85) }

var fileDescriptor_b98ea54e6d90ae85 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x40, 0x95, 0x6a, 0x98, 0x22, 0x40, 0x51, 0x85, 0x5c, 0x60, 0x62, 0x21, 0x1e,
	0x9a, 0x27, 0x60, 0xef, 0xd2, 0x91, 0xa5, 0x4a, 0xdc, 0x8b, 0x6b, 0xa9, 0xbe, 0x8b, 0xdc, 0x0b,
	0xa2, 0x6f, 0xc1, 0x63, 0x31, 0x76, 0x64, 0x44, 0xc9, 0x8b, 0xa0, 0x38, 0x88, 0xb2, 0xfd, 0xff,
	0xf7, 0xff, 0x27, 0xdd, 0x9d, 0xbc, 0x31, 0x35, 0x78, 0x87, 0x0c, 0x41, 0xc3, 0x1b, 0x20, 0xe7,
	0x4d, 0x20, 0xa6, 0x74, 0x6e, 0xb6, 0xa5, 0xc3, 0x02, 0x10, 0x82, 0x3d, 0xe4, 0xa6, 0x80, 0xe8,
	0xf3, 0xbf, 0xf2, 0xec, 0xda, 0x92, 0xa5, 0xd8, 0xd5, 0x83, 0x1a, 0xc7, 0x66, 0x73, 0x4b, 0x64,
	0x77, 0xa0, 0xa3, 0xab, 0xda, 0x5a, 0xb3, 0xf3, 0xb0, 0xe7, 0xd2, 0x37, 0x63, 0xe1, 0x71, 0x2d,
	0x2f, 0x96, 0x0e, 0x39, 0x7d, 0x90, 0x57, 0x15, 0xe1, 0x06, 0x36, 0xeb, 0x50, 0xb2, 0xa3, 0x4c,
	0xdc, 0x8b, 0xa7, 0xe9, 0xea, 0x72, 0x64, 0xab, 0x01, 0xa5, 0x77, 0x72, 0xea, 0xb0, 0xde, 0x0d,
	0x1a, 0xb3, 0xb3, 0x98, 0x9f, 0x40, 0x7a, 0x2b, 0x27, 0xa5, 0xa7, 0x16, 0x39, 0x3b, 0x8f, 0xd1,
	0xaf, 0x7b, 0x59, 0x7e, 0x76, 0x4a, 0x1c, 0x3b, 0x25, 0xbe, 0x3b, 0x25, 0x3e, 0x7a, 0x95, 0x1c,
	0x7b, 0x95, 0x7c, 0xf5, 0x2a, 0x79, 0x5d, 0x58, 0xc7, 0xdb, 0xb6, 0xca, 0x0d, 0x79, 0xfd, 0xff,
	0x3a, 0x6d, 0x0a, 0x78, 0x8e, 0x40, 0xbf, 0xeb, 0xd3, 0x37, 0xf8, 0xd0, 0xc0, 0xbe, 0x9a, 0xc4,
	0xb5, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x9f, 0x63, 0x2d, 0x27, 0x01, 0x00, 0x00,
}

func (m *Mint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Inflation) > 0 {
		i -= len(m.Inflation)
		copy(dAtA[i:], m.Inflation)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Inflation)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BondedRatio) > 0 {
		i -= len(m.BondedRatio)
		copy(dAtA[i:], m.BondedRatio)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.BondedRatio)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BondedRatio)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Inflation)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: Mint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondedRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondedRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inflation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)

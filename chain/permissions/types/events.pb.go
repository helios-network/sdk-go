// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helios/permissions/v1beta1/events.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type EventSetVoucher struct {
	Addr    string     `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Voucher types.Coin `protobuf:"bytes,2,opt,name=voucher,proto3" json:"voucher"`
}

func (m *EventSetVoucher) Reset()         { *m = EventSetVoucher{} }
func (m *EventSetVoucher) String() string { return proto.CompactTextString(m) }
func (*EventSetVoucher) ProtoMessage()    {}
func (*EventSetVoucher) Descriptor() ([]byte, []int) {
	return fileDescriptor_705c3e21b20426fa, []int{0}
}
func (m *EventSetVoucher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetVoucher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetVoucher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetVoucher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetVoucher.Merge(m, src)
}
func (m *EventSetVoucher) XXX_Size() int {
	return m.Size()
}
func (m *EventSetVoucher) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetVoucher.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetVoucher proto.InternalMessageInfo

func (m *EventSetVoucher) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *EventSetVoucher) GetVoucher() types.Coin {
	if m != nil {
		return m.Voucher
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*EventSetVoucher)(nil), "helios.permissions.v1beta1.EventSetVoucher")
}

func init() {
	proto.RegisterFile("helios/permissions/v1beta1/events.proto", fileDescriptor_705c3e21b20426fa)
}

var fileDescriptor_705c3e21b20426fa = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xb3, 0x72, 0x28, 0xc6, 0x42, 0x08, 0x16, 0xe7, 0x81, 0xeb, 0x61, 0x75, 0x08, 0xee,
	0x72, 0x5a, 0xd9, 0x9e, 0x58, 0x08, 0x36, 0x9e, 0x60, 0x61, 0x65, 0xb2, 0x19, 0x92, 0xf5, 0xcc,
	0x4e, 0xd8, 0xd9, 0x04, 0x7c, 0x0b, 0x1f, 0xeb, 0xca, 0x2b, 0xad, 0x44, 0x92, 0x17, 0x91, 0x24,
	0x26, 0x9c, 0xd8, 0xcd, 0xf0, 0x7d, 0x3b, 0x3b, 0xf3, 0xfb, 0xe7, 0xda, 0xbc, 0x82, 0x72, 0xba,
	0x04, 0x99, 0x83, 0xcd, 0x34, 0x91, 0x46, 0x43, 0xb2, 0x9c, 0x47, 0xe0, 0xc2, 0xb9, 0x84, 0x12,
	0x8c, 0x23, 0x91, 0x5b, 0x74, 0x18, 0x9c, 0x0c, 0xae, 0xd8, 0x72, 0xc5, 0xaf, 0x3b, 0x39, 0x4a,
	0x30, 0xc1, 0xd6, 0x94, 0x4d, 0xd5, 0x3d, 0x9a, 0x70, 0x85, 0x94, 0x21, 0xc9, 0x28, 0x24, 0x18,
	0xc6, 0x2a, 0xd4, 0xe6, 0x1f, 0x37, 0xab, 0x81, 0x37, 0x4d, 0xc7, 0xcf, 0x5e, 0xfc, 0xc3, 0xdb,
	0x66, 0x89, 0x47, 0x70, 0x4f, 0x58, 0xa8, 0x14, 0x6c, 0x10, 0xf8, 0xa3, 0x30, 0x8e, 0xed, 0x98,
	0x4d, 0xd9, 0x6c, 0x7f, 0xd9, 0xd6, 0xc1, 0xb5, 0xbf, 0x57, 0x76, 0x78, 0xbc, 0x33, 0x65, 0xb3,
	0x83, 0xcb, 0x63, 0xd1, 0x0d, 0x16, 0xcd, 0xc7, 0xfd, 0x8e, 0xe2, 0x06, 0xb5, 0x59, 0x8c, 0xd6,
	0x5f, 0xa7, 0xde, 0xb2, 0xf7, 0x17, 0xab, 0x75, 0xc5, 0xd9, 0xa6, 0xe2, 0xec, 0xbb, 0xe2, 0xec,
	0xa3, 0xe6, 0xde, 0xa6, 0xe6, 0xde, 0x67, 0xcd, 0xbd, 0xe7, 0x87, 0x44, 0xbb, 0xb4, 0x88, 0x84,
	0xc2, 0x4c, 0xde, 0xf5, 0xb7, 0xdf, 0x87, 0x11, 0xc9, 0x21, 0x89, 0x0b, 0x85, 0x16, 0xb6, 0xdb,
	0x34, 0xd4, 0x46, 0x66, 0x18, 0x17, 0x6f, 0x40, 0x7f, 0x22, 0x75, 0xef, 0x39, 0x50, 0xb4, 0xdb,
	0x5e, 0x75, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0x74, 0x79, 0xe1, 0x33, 0x78, 0x01, 0x00, 0x00,
}

func (m *EventSetVoucher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetVoucher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetVoucher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Voucher.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventSetVoucher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Voucher.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventSetVoucher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetVoucher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetVoucher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voucher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Voucher.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)

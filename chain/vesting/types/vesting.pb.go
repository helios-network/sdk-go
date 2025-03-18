// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helios/vesting/v2/vesting.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	github_com_cosmos_cosmos_sdk_x_auth_vesting_types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClawbackVestingAccount implements the VestingAccount interface. It provides
// an account that can hold contributions subject to "lockup" (like a
// PeriodicVestingAccount), or vesting which is subject to clawback
// of unvested tokens, or a combination (tokens vest, but are still locked).
type ClawbackVestingAccount struct {
	// base_vesting_account implements the VestingAccount interface. It contains
	// all the necessary fields needed for any vesting account implementation
	*types.BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	// funder_address specifies the account which can perform clawback
	FunderAddress string `protobuf:"bytes,2,opt,name=funder_address,json=funderAddress,proto3" json:"funder_address,omitempty"`
	// start_time defines the time at which the vesting period begins
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// lockup_periods defines the unlocking schedule relative to the start_time
	LockupPeriods github_com_cosmos_cosmos_sdk_x_auth_vesting_types.Periods `protobuf:"bytes,4,rep,name=lockup_periods,json=lockupPeriods,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/auth/vesting/types.Periods" json:"lockup_periods"`
	// vesting_periods defines the vesting schedule relative to the start_time
	VestingPeriods github_com_cosmos_cosmos_sdk_x_auth_vesting_types.Periods `protobuf:"bytes,5,rep,name=vesting_periods,json=vestingPeriods,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/auth/vesting/types.Periods" json:"vesting_periods"`
}

func (m *ClawbackVestingAccount) Reset()      { *m = ClawbackVestingAccount{} }
func (*ClawbackVestingAccount) ProtoMessage() {}
func (*ClawbackVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_08cc7bcd9d50f5ea, []int{0}
}
func (m *ClawbackVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClawbackVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClawbackVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClawbackVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClawbackVestingAccount.Merge(m, src)
}
func (m *ClawbackVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ClawbackVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ClawbackVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ClawbackVestingAccount proto.InternalMessageInfo

// ClawbackProposal is a gov Content type to clawback funds
// from a vesting account that has this functionality enabled.
type ClawbackProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// address is the vesting account address
	// to clawback the funds from
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// destination_address is the address that will receive
	// the clawbacked funds from the given vesting account. When
	// empty, proposal will return the coins to the vesting
	// account funder.
	DestinationAddress string `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
}

func (m *ClawbackProposal) Reset()         { *m = ClawbackProposal{} }
func (m *ClawbackProposal) String() string { return proto.CompactTextString(m) }
func (*ClawbackProposal) ProtoMessage()    {}
func (*ClawbackProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_08cc7bcd9d50f5ea, []int{1}
}
func (m *ClawbackProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClawbackProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClawbackProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClawbackProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClawbackProposal.Merge(m, src)
}
func (m *ClawbackProposal) XXX_Size() int {
	return m.Size()
}
func (m *ClawbackProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ClawbackProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ClawbackProposal proto.InternalMessageInfo

func (m *ClawbackProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ClawbackProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ClawbackProposal) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClawbackProposal) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*ClawbackVestingAccount)(nil), "helios.vesting.v2.ClawbackVestingAccount")
	proto.RegisterType((*ClawbackProposal)(nil), "helios.vesting.v2.ClawbackProposal")
}

func init() { proto.RegisterFile("helios/vesting/v2/vesting.proto", fileDescriptor_08cc7bcd9d50f5ea) }

var fileDescriptor_08cc7bcd9d50f5ea = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x8d, 0x69, 0x07, 0xaa, 0xab, 0x16, 0x66, 0x2a, 0x54, 0xf5, 0x90, 0x54, 0x13, 0x48, 0xa5,
	0xd2, 0x62, 0xad, 0x48, 0x48, 0xec, 0xb6, 0x20, 0x21, 0x8e, 0x53, 0x85, 0x38, 0x70, 0xa9, 0x9c,
	0xc4, 0x4b, 0xad, 0xa6, 0x71, 0x14, 0x3b, 0x65, 0xfc, 0x03, 0x84, 0x84, 0xd8, 0x11, 0x89, 0x4b,
	0x8f, 0x88, 0xd3, 0xfe, 0x05, 0x3b, 0xf6, 0xc8, 0x69, 0x43, 0xed, 0x61, 0xfc, 0x0c, 0x14, 0x3b,
	0x1e, 0x99, 0x80, 0xeb, 0x2e, 0xc9, 0xe7, 0xf7, 0xbd, 0x3c, 0xbf, 0xcf, 0x7e, 0x81, 0xce, 0x94,
	0xc6, 0x8c, 0x0b, 0xbc, 0xa0, 0x42, 0xb2, 0x24, 0xc2, 0x8b, 0x91, 0x29, 0xdd, 0x34, 0xe3, 0x92,
	0xa3, 0x6d, 0x4d, 0x70, 0x0d, 0xba, 0x18, 0xf5, 0xb6, 0xc9, 0x9c, 0x25, 0x1c, 0xab, 0xa7, 0x66,
	0xf5, 0x1e, 0x06, 0x5c, 0xcc, 0xab, 0x32, 0x7b, 0x3e, 0x95, 0x64, 0xef, 0xba, 0x56, 0xaf, 0x13,
	0xf1, 0x88, 0xab, 0x12, 0x17, 0x55, 0x89, 0x3a, 0x11, 0xe7, 0x51, 0x4c, 0xb1, 0x5a, 0xf9, 0xf9,
	0x11, 0x96, 0x6c, 0x4e, 0x85, 0x24, 0xf3, 0x54, 0x13, 0x76, 0xbe, 0xd7, 0xe1, 0x83, 0xe7, 0x31,
	0x79, 0xeb, 0x93, 0x60, 0xf6, 0x5a, 0x0b, 0x1e, 0x04, 0x01, 0xcf, 0x13, 0x89, 0x7c, 0xd8, 0xf1,
	0x89, 0xa0, 0x93, 0x72, 0x9f, 0x09, 0xd1, 0x78, 0x17, 0xf4, 0xc1, 0xa0, 0x39, 0x1a, 0xba, 0xda,
	0xd6, 0x1f, 0xf3, 0xda, 0x96, 0xeb, 0x11, 0x41, 0xaf, 0x2b, 0x79, 0xf5, 0xd5, 0xb9, 0x03, 0xc6,
	0xc8, 0xff, 0xab, 0x83, 0x1e, 0xc1, 0xf6, 0x51, 0x9e, 0x84, 0x34, 0x9b, 0x90, 0x30, 0xcc, 0xa8,
	0x10, 0xdd, 0x5b, 0x7d, 0x30, 0x68, 0x8c, 0x5b, 0x1a, 0x3d, 0xd0, 0x20, 0x7a, 0x09, 0xa1, 0x90,
	0x24, 0x93, 0x93, 0xc2, 0x7e, 0xb7, 0xa6, 0x0c, 0xf4, 0x5c, 0x3d, 0x9b, 0x6b, 0x66, 0x73, 0x5f,
	0x99, 0xd9, 0xbc, 0xd6, 0xd9, 0xb9, 0x63, 0x9d, 0x5c, 0x38, 0xe0, 0xeb, 0xe5, 0xe9, 0x10, 0x8c,
	0x1b, 0xea, 0xe3, 0xa2, 0x8d, 0x3e, 0x02, 0xd8, 0x8e, 0x79, 0x30, 0xcb, 0xd3, 0x49, 0x4a, 0x33,
	0xc6, 0x43, 0xd1, 0xad, 0xf7, 0x6b, 0x83, 0xe6, 0xc8, 0xfe, 0xdf, 0x3c, 0x87, 0x8a, 0xe6, 0xbd,
	0x28, 0x24, 0xbf, 0x5d, 0x38, 0xcf, 0x22, 0x26, 0xa7, 0xb9, 0xef, 0x06, 0x7c, 0x8e, 0xcb, 0x8b,
	0xd1, 0xaf, 0x5d, 0x11, 0xce, 0xf0, 0x31, 0x26, 0xb9, 0x9c, 0x5e, 0x5d, 0x95, 0x7c, 0x97, 0x52,
	0x51, 0x2a, 0x08, 0xed, 0xa5, 0xa5, 0x77, 0x2f, 0x31, 0xf4, 0x09, 0xc0, 0xbb, 0xe6, 0x80, 0x8d,
	0xa1, 0xad, 0x1b, 0x35, 0xd4, 0x2e, 0x7b, 0x25, 0xb8, 0xff, 0xf4, 0xfd, 0xd2, 0xb1, 0x3e, 0x2f,
	0x1d, 0xeb, 0xc3, 0xe5, 0xe9, 0xf0, 0x31, 0x5d, 0x14, 0x82, 0xc7, 0xd5, 0x0c, 0xff, 0x3b, 0x2e,
	0x3b, 0x5f, 0x00, 0xbc, 0x67, 0x5a, 0x87, 0x19, 0x4f, 0xb9, 0x20, 0x31, 0xea, 0xc0, 0x2d, 0xc9,
	0x64, 0x4c, 0x55, 0x68, 0x1a, 0x63, 0xbd, 0x40, 0x7d, 0xd8, 0x0c, 0xa9, 0x08, 0x32, 0x96, 0x4a,
	0xc6, 0x93, 0xf2, 0xca, 0xab, 0x10, 0xea, 0xc2, 0x3b, 0x26, 0x10, 0x35, 0xd5, 0x35, 0x4b, 0x84,
	0xe1, 0xfd, 0x50, 0x6d, 0x4c, 0x0a, 0xe2, 0x55, 0x6c, 0xea, 0x8a, 0x85, 0x2a, 0xad, 0x32, 0x3b,
	0xfb, 0xf5, 0x5f, 0x4b, 0xc7, 0xf2, 0xbc, 0xb3, 0xb5, 0x0d, 0x56, 0x6b, 0x1b, 0xfc, 0x5c, 0xdb,
	0xe0, 0x64, 0x63, 0x5b, 0xab, 0x8d, 0x6d, 0xfd, 0xd8, 0xd8, 0xd6, 0x9b, 0x81, 0xfe, 0x09, 0x77,
	0x03, 0x9e, 0x51, 0x6c, 0xea, 0x29, 0x61, 0x49, 0x65, 0x6a, 0x75, 0x6c, 0xfe, 0x6d, 0x95, 0xb4,
	0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x3e, 0xa5, 0xf3, 0xd8, 0x03, 0x00, 0x00,
}

func (m *ClawbackVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClawbackVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClawbackVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LockupPeriods) > 0 {
		for iNdEx := len(m.LockupPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockupPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintVesting(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.FunderAddress) > 0 {
		i -= len(m.FunderAddress)
		copy(dAtA[i:], m.FunderAddress)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.FunderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClawbackProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClawbackProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClawbackProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClawbackVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.FunderAddress)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovVesting(uint64(l))
	if len(m.LockupPeriods) > 0 {
		for _, e := range m.LockupPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func (m *ClawbackProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClawbackVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ClawbackVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClawbackVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &types.BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockupPeriods = append(m.LockupPeriods, types.Period{})
			if err := m.LockupPeriods[len(m.LockupPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, types.Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *ClawbackProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ClawbackProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClawbackProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)

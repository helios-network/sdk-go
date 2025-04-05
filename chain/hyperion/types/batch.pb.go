// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helios/hyperion/v1/batch.proto

package types

import (
	fmt "fmt"
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

// OutgoingTxBatch represents a batch of transactions going from Hyperion to ETH
type OutgoingTxBatch struct {
	HyperionId    uint64                `protobuf:"varint,1,opt,name=hyperion_id,json=hyperionId,proto3" json:"hyperion_id,omitempty"`
	BatchNonce    uint64                `protobuf:"varint,2,opt,name=batch_nonce,json=batchNonce,proto3" json:"batch_nonce,omitempty"`
	BatchTimeout  uint64                `protobuf:"varint,3,opt,name=batch_timeout,json=batchTimeout,proto3" json:"batch_timeout,omitempty"`
	Transactions  []*OutgoingTransferTx `protobuf:"bytes,4,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TokenContract string                `protobuf:"bytes,5,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
	Block         uint64                `protobuf:"varint,6,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *OutgoingTxBatch) Reset()         { *m = OutgoingTxBatch{} }
func (m *OutgoingTxBatch) String() string { return proto.CompactTextString(m) }
func (*OutgoingTxBatch) ProtoMessage()    {}
func (*OutgoingTxBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_7440bfae5ce81cd7, []int{0}
}
func (m *OutgoingTxBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTxBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTxBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTxBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTxBatch.Merge(m, src)
}
func (m *OutgoingTxBatch) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTxBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTxBatch.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTxBatch proto.InternalMessageInfo

func (m *OutgoingTxBatch) GetHyperionId() uint64 {
	if m != nil {
		return m.HyperionId
	}
	return 0
}

func (m *OutgoingTxBatch) GetBatchNonce() uint64 {
	if m != nil {
		return m.BatchNonce
	}
	return 0
}

func (m *OutgoingTxBatch) GetBatchTimeout() uint64 {
	if m != nil {
		return m.BatchTimeout
	}
	return 0
}

func (m *OutgoingTxBatch) GetTransactions() []*OutgoingTransferTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *OutgoingTxBatch) GetTokenContract() string {
	if m != nil {
		return m.TokenContract
	}
	return ""
}

func (m *OutgoingTxBatch) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

// OutgoingTransferTx represents an individual send from Hyperion to ETH
type OutgoingTransferTx struct {
	HyperionId  uint64      `protobuf:"varint,1,opt,name=hyperion_id,json=hyperionId,proto3" json:"hyperion_id,omitempty"`
	Id          uint64      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Sender      string      `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddress string      `protobuf:"bytes,4,opt,name=dest_address,json=destAddress,proto3" json:"dest_address,omitempty"`
	Erc20Token  *ERC20Token `protobuf:"bytes,5,opt,name=erc20_token,json=erc20Token,proto3" json:"erc20_token,omitempty"`
	Erc20Fee    *ERC20Token `protobuf:"bytes,6,opt,name=erc20_fee,json=erc20Fee,proto3" json:"erc20_fee,omitempty"`
	TxTimeout   uint64      `protobuf:"varint,7,opt,name=tx_timeout,json=txTimeout,proto3" json:"tx_timeout,omitempty"`
	TxHash      string      `protobuf:"bytes,8,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (m *OutgoingTransferTx) Reset()         { *m = OutgoingTransferTx{} }
func (m *OutgoingTransferTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTransferTx) ProtoMessage()    {}
func (*OutgoingTransferTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_7440bfae5ce81cd7, []int{1}
}
func (m *OutgoingTransferTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTransferTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTransferTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTransferTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTransferTx.Merge(m, src)
}
func (m *OutgoingTransferTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTransferTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTransferTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTransferTx proto.InternalMessageInfo

func (m *OutgoingTransferTx) GetHyperionId() uint64 {
	if m != nil {
		return m.HyperionId
	}
	return 0
}

func (m *OutgoingTransferTx) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OutgoingTransferTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTransferTx) GetDestAddress() string {
	if m != nil {
		return m.DestAddress
	}
	return ""
}

func (m *OutgoingTransferTx) GetErc20Token() *ERC20Token {
	if m != nil {
		return m.Erc20Token
	}
	return nil
}

func (m *OutgoingTransferTx) GetErc20Fee() *ERC20Token {
	if m != nil {
		return m.Erc20Fee
	}
	return nil
}

func (m *OutgoingTransferTx) GetTxTimeout() uint64 {
	if m != nil {
		return m.TxTimeout
	}
	return 0
}

func (m *OutgoingTransferTx) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func init() {
	proto.RegisterType((*OutgoingTxBatch)(nil), "helios.hyperion.v1.OutgoingTxBatch")
	proto.RegisterType((*OutgoingTransferTx)(nil), "helios.hyperion.v1.OutgoingTransferTx")
}

func init() { proto.RegisterFile("helios/hyperion/v1/batch.proto", fileDescriptor_7440bfae5ce81cd7) }

var fileDescriptor_7440bfae5ce81cd7 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x37, 0x69, 0xbb, 0xed, 0x4e, 0xb6, 0x45, 0xb2, 0x10, 0x44, 0x48, 0x84, 0xa5, 0xfc,
	0xd1, 0x72, 0x20, 0x69, 0x97, 0x23, 0x07, 0x44, 0x57, 0x20, 0xe0, 0x00, 0x52, 0x94, 0x13, 0x97,
	0xc8, 0x9b, 0x4c, 0x1b, 0xab, 0xc5, 0x5e, 0xd9, 0xd3, 0x2a, 0x3d, 0xf2, 0x06, 0x3c, 0x0d, 0xcf,
	0xc0, 0xb1, 0x47, 0x8e, 0x68, 0xf7, 0x45, 0x90, 0xed, 0x4d, 0x11, 0x6a, 0xa5, 0xbd, 0xd9, 0xbf,
	0xf9, 0x3e, 0xcf, 0x7c, 0xb6, 0x21, 0x69, 0xf0, 0x4c, 0x28, 0x93, 0x35, 0x97, 0x73, 0xd4, 0x42,
	0xc9, 0xec, 0xe2, 0x30, 0x9b, 0x71, 0xaa, 0x9a, 0x74, 0xae, 0x15, 0x29, 0xc6, 0x7c, 0x3d, 0xed,
	0xea, 0xe9, 0xc5, 0xe1, 0x83, 0xa7, 0xb7, 0x78, 0x38, 0x11, 0x1a, 0xe2, 0x64, 0x25, 0xce, 0xb9,
	0xff, 0x3d, 0x84, 0x3b, 0x5f, 0xce, 0xe9, 0x44, 0x09, 0x79, 0x52, 0xb4, 0x47, 0xf6, 0x4c, 0xf6,
	0x08, 0xa2, 0xce, 0x54, 0x8a, 0x3a, 0x0e, 0x46, 0xc1, 0x78, 0x33, 0x87, 0x0e, 0x7d, 0xac, 0xad,
	0xc0, 0x75, 0x2f, 0xa5, 0x92, 0x15, 0xc6, 0xa1, 0x17, 0x38, 0xf4, 0xd9, 0x12, 0xf6, 0x04, 0x76,
	0xbd, 0x80, 0xc4, 0x37, 0x54, 0xe7, 0x14, 0x6f, 0x38, 0xc9, 0xd0, 0xc1, 0xc2, 0x33, 0xf6, 0x09,
	0x86, 0xa4, 0xb9, 0x34, 0xbc, 0xb2, 0xf3, 0x98, 0x78, 0x73, 0xb4, 0x31, 0x8e, 0x26, 0xcf, 0xd3,
	0x9b, 0x59, 0xd2, 0xeb, 0x09, 0xad, 0xfe, 0x18, 0x75, 0xd1, 0xe6, 0xff, 0x79, 0xd9, 0x33, 0xd8,
	0x23, 0x75, 0x8a, 0xb2, 0xac, 0x94, 0x24, 0xcd, 0x2b, 0x8a, 0xb7, 0x46, 0xc1, 0x78, 0x90, 0xef,
	0x3a, 0x3a, 0x5d, 0x41, 0x76, 0x17, 0xb6, 0x66, 0x67, 0xaa, 0x3a, 0x8d, 0xfb, 0x6e, 0x1e, 0xbf,
	0xd9, 0xff, 0x19, 0x02, 0xbb, 0xd9, 0x61, 0xfd, 0x35, 0xec, 0x41, 0x28, 0xea, 0x55, 0xfa, 0x50,
	0xd4, 0xec, 0x1e, 0xf4, 0x0d, 0xca, 0x1a, 0xb5, 0x8b, 0x3b, 0xc8, 0x57, 0x3b, 0xf6, 0x18, 0x86,
	0x35, 0x1a, 0x2a, 0x79, 0x5d, 0x6b, 0x34, 0x36, 0xa8, 0xad, 0x46, 0x96, 0xbd, 0xf5, 0x88, 0xbd,
	0x81, 0x08, 0x75, 0x35, 0x39, 0x28, 0xdd, 0xbc, 0x6e, 0xf8, 0x68, 0x92, 0xdc, 0x76, 0x15, 0xef,
	0xf2, 0xe9, 0xe4, 0xa0, 0xb0, 0xaa, 0x1c, 0x9c, 0xc5, 0xad, 0xd9, 0x6b, 0x18, 0xf8, 0x03, 0x8e,
	0x11, 0x5d, 0xba, 0xf5, 0xf6, 0x1d, 0x67, 0x78, 0x8f, 0xc8, 0x1e, 0x02, 0x50, 0x7b, 0xfd, 0x56,
	0xdb, 0x2e, 0xd0, 0x80, 0xda, 0xee, 0xa1, 0xee, 0xc3, 0x36, 0xb5, 0x65, 0xc3, 0x4d, 0x13, 0xef,
	0xf8, 0x60, 0xd4, 0x7e, 0xe0, 0xa6, 0x39, 0x9a, 0xfe, 0x5a, 0x24, 0xc1, 0xd5, 0x22, 0x09, 0xfe,
	0x2c, 0x92, 0xe0, 0xc7, 0x32, 0xe9, 0x5d, 0x2d, 0x93, 0xde, 0xef, 0x65, 0xd2, 0xfb, 0xfa, 0xc2,
	0xb7, 0x7e, 0x59, 0x29, 0x8d, 0x59, 0xb7, 0x6e, 0xb8, 0x90, 0x59, 0xfb, 0xef, 0x43, 0xd2, 0xe5,
	0x1c, 0xcd, 0xac, 0xef, 0x3e, 0xe2, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x38, 0x85, 0x78,
	0x1d, 0xe4, 0x02, 0x00, 0x00,
}

func (m *OutgoingTxBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTxBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTxBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TokenContract) > 0 {
		i -= len(m.TokenContract)
		copy(dAtA[i:], m.TokenContract)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.TokenContract)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BatchTimeout != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.BatchTimeout))
		i--
		dAtA[i] = 0x18
	}
	if m.BatchNonce != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.BatchNonce))
		i--
		dAtA[i] = 0x10
	}
	if m.HyperionId != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.HyperionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutgoingTransferTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTransferTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTransferTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x42
	}
	if m.TxTimeout != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.TxTimeout))
		i--
		dAtA[i] = 0x38
	}
	if m.Erc20Fee != nil {
		{
			size, err := m.Erc20Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Erc20Token != nil {
		{
			size, err := m.Erc20Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestAddress) > 0 {
		i -= len(m.DestAddress)
		copy(dAtA[i:], m.DestAddress)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.DestAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.HyperionId != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.HyperionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTxBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HyperionId != 0 {
		n += 1 + sovBatch(uint64(m.HyperionId))
	}
	if m.BatchNonce != 0 {
		n += 1 + sovBatch(uint64(m.BatchNonce))
	}
	if m.BatchTimeout != 0 {
		n += 1 + sovBatch(uint64(m.BatchTimeout))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovBatch(uint64(l))
		}
	}
	l = len(m.TokenContract)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovBatch(uint64(m.Block))
	}
	return n
}

func (m *OutgoingTransferTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HyperionId != 0 {
		n += 1 + sovBatch(uint64(m.HyperionId))
	}
	if m.Id != 0 {
		n += 1 + sovBatch(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.DestAddress)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Erc20Token != nil {
		l = m.Erc20Token.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Erc20Fee != nil {
		l = m.Erc20Fee.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.TxTimeout != 0 {
		n += 1 + sovBatch(uint64(m.TxTimeout))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	return n
}

func sovBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBatch(x uint64) (n int) {
	return sovBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTxBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: OutgoingTxBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTxBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HyperionId", wireType)
			}
			m.HyperionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HyperionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchNonce", wireType)
			}
			m.BatchNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchTimeout", wireType)
			}
			m.BatchTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &OutgoingTransferTx{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBatch
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
func (m *OutgoingTransferTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: OutgoingTransferTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTransferTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HyperionId", wireType)
			}
			m.HyperionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HyperionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Token == nil {
				m.Erc20Token = &ERC20Token{}
			}
			if err := m.Erc20Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Fee == nil {
				m.Erc20Fee = &ERC20Token{}
			}
			if err := m.Erc20Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxTimeout", wireType)
			}
			m.TxTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBatch
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
func skipBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
				return 0, ErrInvalidLengthBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBatch = fmt.Errorf("proto: unexpected end of group")
)

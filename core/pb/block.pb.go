// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	BlockHeader
	Block
	DownloadParentBlock
	Transaction
	TransactionHashTarget
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	Hash                 []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash           []byte `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Coinbase             []byte `protobuf:"bytes,3,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Reward               []byte `protobuf:"bytes,4,opt,name=reward,proto3" json:"reward,omitempty"`
	Supply               []byte `protobuf:"bytes,5,opt,name=supply,proto3" json:"supply,omitempty"`
	Timestamp            int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChainId              uint32 `protobuf:"varint,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Alg                  uint32 `protobuf:"varint,11,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign                 []byte `protobuf:"bytes,12,opt,name=sign,proto3" json:"sign,omitempty"`
	AccStateRoot         []byte `protobuf:"bytes,21,opt,name=acc_state_root,json=accStateRoot,proto3" json:"acc_state_root,omitempty"`
	DataStateRoot        []byte `protobuf:"bytes,22,opt,name=data_state_root,json=dataStateRoot,proto3" json:"data_state_root,omitempty"`
	DposRoot             []byte `protobuf:"bytes,23,opt,name=dpos_root,json=dposRoot,proto3" json:"dpos_root,omitempty"`
	UsageRoot            []byte `protobuf:"bytes,24,opt,name=usage_root,json=usageRoot,proto3" json:"usage_root,omitempty"`
	ReservationQueueHash []byte `protobuf:"bytes,25,opt,name=reservation_queue_hash,json=reservationQueueHash,proto3" json:"reservation_queue_hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetReward() []byte {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *BlockHeader) GetSupply() []byte {
	if m != nil {
		return m.Supply
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockHeader) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetAccStateRoot() []byte {
	if m != nil {
		return m.AccStateRoot
	}
	return nil
}

func (m *BlockHeader) GetDataStateRoot() []byte {
	if m != nil {
		return m.DataStateRoot
	}
	return nil
}

func (m *BlockHeader) GetDposRoot() []byte {
	if m != nil {
		return m.DposRoot
	}
	return nil
}

func (m *BlockHeader) GetUsageRoot() []byte {
	if m != nil {
		return m.UsageRoot
	}
	return nil
}

func (m *BlockHeader) GetReservationQueueHash() []byte {
	if m != nil {
		return m.ReservationQueueHash
	}
	return nil
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Height       uint64         `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type DownloadParentBlock struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Sign []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *DownloadParentBlock) Reset()                    { *m = DownloadParentBlock{} }
func (m *DownloadParentBlock) String() string            { return proto.CompactTextString(m) }
func (*DownloadParentBlock) ProtoMessage()               {}
func (*DownloadParentBlock) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *DownloadParentBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DownloadParentBlock) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	TxType    string `protobuf:"bytes,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	From      []byte `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce     uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Payload   []byte `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
	Alg       uint32 `protobuf:"varint,21,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,22,opt,name=sign,proto3" json:"sign,omitempty"`
	PayerSign []byte `protobuf:"bytes,23,opt,name=payerSign,proto3" json:"payerSign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{3} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Transaction) GetPayerSign() []byte {
	if m != nil {
		return m.PayerSign
	}
	return nil
}

type TransactionHashTarget struct {
	TxType    string `protobuf:"bytes,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	From      []byte `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce     uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Payload   []byte `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionHashTarget) Reset()                    { *m = TransactionHashTarget{} }
func (m *TransactionHashTarget) String() string            { return proto.CompactTextString(m) }
func (*TransactionHashTarget) ProtoMessage()               {}
func (*TransactionHashTarget) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{4} }

func (m *TransactionHashTarget) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *TransactionHashTarget) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransactionHashTarget) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransactionHashTarget) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TransactionHashTarget) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TransactionHashTarget) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransactionHashTarget) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *TransactionHashTarget) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
	proto.RegisterType((*DownloadParentBlock)(nil), "corepb.DownloadParentBlock")
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*TransactionHashTarget)(nil), "corepb.TransactionHashTarget")
}

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0x4e, 0x0b, 0x14, 0x38, 0x65, 0x57, 0x33, 0x0b, 0xec, 0xac, 0xae, 0x91, 0x10, 0x63, 0x48,
	0x4c, 0xb8, 0x50, 0x13, 0xaf, 0xbc, 0x31, 0x5e, 0xac, 0x77, 0xda, 0xe5, 0xbe, 0x39, 0xb4, 0x23,
	0x6d, 0x84, 0xce, 0x38, 0x33, 0xec, 0x2e, 0x0f, 0xe0, 0x83, 0xf8, 0x00, 0x3e, 0x8d, 0x2f, 0x64,
	0xe6, 0x4c, 0x91, 0x12, 0x37, 0xf1, 0x7a, 0xef, 0xe6, 0xfb, 0x39, 0x0c, 0x73, 0xbe, 0x2f, 0x85,
	0x78, 0xb9, 0x96, 0xd9, 0xb7, 0xb9, 0xd2, 0xd2, 0x4a, 0x16, 0x65, 0x52, 0x0b, 0xb5, 0x9c, 0xfe,
	0x6a, 0x41, 0xfc, 0xc1, 0xf1, 0x57, 0x02, 0x73, 0xa1, 0x19, 0x83, 0x76, 0x81, 0xa6, 0xe0, 0xc1,
	0x24, 0x98, 0x0d, 0x12, 0x3a, 0xb3, 0xe7, 0x10, 0x2b, 0xd4, 0xa2, 0xb2, 0x29, 0x49, 0x21, 0x49,
	0xe0, 0xa9, 0x2b, 0x67, 0x78, 0x02, 0xbd, 0x4c, 0x96, 0xd5, 0x12, 0x8d, 0xe0, 0x2d, 0x52, 0xff,
	0x62, 0x36, 0x86, 0x48, 0x8b, 0x5b, 0xd4, 0x39, 0x6f, 0x93, 0x52, 0x23, 0xc7, 0x9b, 0xad, 0x52,
	0xeb, 0x1d, 0xef, 0x78, 0xde, 0x23, 0x76, 0x09, 0x7d, 0x5b, 0x6e, 0x84, 0xb1, 0xb8, 0x51, 0x3c,
	0x9a, 0x04, 0xb3, 0x56, 0x72, 0x20, 0xd8, 0x05, 0xf4, 0xb2, 0x02, 0xcb, 0x2a, 0x2d, 0x73, 0xde,
	0x9d, 0x04, 0xb3, 0x93, 0xa4, 0x4b, 0xf8, 0x53, 0xce, 0x1e, 0x43, 0x0b, 0xd7, 0x2b, 0x1e, 0x13,
	0xeb, 0x8e, 0xee, 0x2d, 0xa6, 0x5c, 0x55, 0x7c, 0xe0, 0xdf, 0xe2, 0xce, 0xec, 0x05, 0x9c, 0x62,
	0x96, 0xa5, 0xc6, 0xa2, 0x15, 0xa9, 0x96, 0xd2, 0xf2, 0x11, 0xa9, 0x03, 0xcc, 0xb2, 0x6b, 0x47,
	0x26, 0x52, 0x5a, 0xf6, 0x12, 0x1e, 0xe5, 0x68, 0xb1, 0x69, 0x1b, 0x93, 0xed, 0xc4, 0xd1, 0x07,
	0xdf, 0x53, 0xe8, 0xe7, 0x4a, 0x1a, 0xef, 0x38, 0xf7, 0x2f, 0x77, 0x04, 0x89, 0xcf, 0x00, 0xb6,
	0x06, 0x57, 0xf5, 0x3c, 0x27, 0xb5, 0x4f, 0x0c, 0xc9, 0x6f, 0x61, 0xac, 0x85, 0x11, 0xfa, 0x06,
	0x6d, 0x29, 0xab, 0xf4, 0xfb, 0x56, 0x6c, 0x85, 0x5f, 0xf0, 0x05, 0x59, 0x87, 0x0d, 0xf5, 0x8b,
	0x13, 0xdd, 0xaa, 0xa7, 0x3f, 0x02, 0xe8, 0x50, 0x5e, 0xec, 0x15, 0x44, 0x05, 0x65, 0x46, 0x59,
	0xc5, 0xaf, 0xcf, 0xe6, 0x3e, 0xd2, 0x79, 0x23, 0xce, 0xa4, 0xb6, 0xb0, 0x77, 0x30, 0xb0, 0x1a,
	0x2b, 0x83, 0x99, 0xfb, 0x39, 0xc3, 0xc3, 0x49, 0xab, 0x39, 0xb2, 0x38, 0x68, 0xc9, 0x91, 0xd1,
	0xc5, 0x54, 0x88, 0x72, 0x55, 0x58, 0x0a, 0xb6, 0x9d, 0xd4, 0x68, 0xfa, 0x1e, 0xce, 0x3e, 0xca,
	0xdb, 0x6a, 0x2d, 0x31, 0xff, 0x4c, 0x45, 0xf0, 0x7f, 0xea, 0xbe, 0xfa, 0xec, 0x63, 0x08, 0x0f,
	0x31, 0x4c, 0x7f, 0x86, 0x10, 0x37, 0x2e, 0xbd, 0x77, 0xee, 0x1c, 0xba, 0xf6, 0x2e, 0xb5, 0x3b,
	0x25, 0x68, 0xb4, 0x9f, 0x44, 0xf6, 0x6e, 0xb1, 0x53, 0xc2, 0x99, 0xbf, 0x6a, 0xb9, 0xa9, 0xab,
	0x46, 0x67, 0x76, 0x0a, 0xa1, 0x95, 0x75, 0xc5, 0x42, 0x2b, 0xd9, 0x10, 0x3a, 0x37, 0xb8, 0xde,
	0x8a, 0xba, 0x5d, 0x1e, 0xfc, 0xa7, 0x5c, 0x43, 0xe8, 0x54, 0xb2, 0xca, 0x04, 0x35, 0xab, 0x9d,
	0x78, 0x70, 0x54, 0xb9, 0xde, 0x71, 0xe5, 0x38, 0x74, 0x15, 0xee, 0xdc, 0x0e, 0x38, 0xd0, 0x35,
	0x7b, 0xb8, 0x2f, 0xe3, 0xe8, 0xdf, 0x32, 0x8e, 0x1b, 0x65, 0xbc, 0x84, 0xbe, 0xc2, 0x9d, 0xd0,
	0xd7, 0x4e, 0xf0, 0xf5, 0x39, 0x10, 0xd3, 0xdf, 0x01, 0x8c, 0x1a, 0x3b, 0x72, 0xf1, 0x2f, 0x50,
	0xaf, 0x84, 0x7d, 0xc8, 0x9b, 0x59, 0x46, 0xf4, 0xfd, 0x79, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0x72, 0xd2, 0xac, 0x8e, 0x04, 0x00, 0x00,
}

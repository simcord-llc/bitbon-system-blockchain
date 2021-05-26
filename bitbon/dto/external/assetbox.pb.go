// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assetbox.proto

package external

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AssetboxBalance struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetboxBalance) Reset()         { *m = AssetboxBalance{} }
func (m *AssetboxBalance) String() string { return proto.CompactTextString(m) }
func (*AssetboxBalance) ProtoMessage()    {}
func (*AssetboxBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_661719b864488047, []int{0}
}

func (m *AssetboxBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetboxBalance.Unmarshal(m, b)
}
func (m *AssetboxBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetboxBalance.Marshal(b, m, deterministic)
}
func (m *AssetboxBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetboxBalance.Merge(m, src)
}
func (m *AssetboxBalance) XXX_Size() int {
	return xxx_messageInfo_AssetboxBalance.Size(m)
}
func (m *AssetboxBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetboxBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AssetboxBalance proto.InternalMessageInfo

func (m *AssetboxBalance) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *AssetboxBalance) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type AssetboxCryptoData struct {
	Wallet               []byte   `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Passphrase           []byte   `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetboxCryptoData) Reset()         { *m = AssetboxCryptoData{} }
func (m *AssetboxCryptoData) String() string { return proto.CompactTextString(m) }
func (*AssetboxCryptoData) ProtoMessage()    {}
func (*AssetboxCryptoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_661719b864488047, []int{1}
}

func (m *AssetboxCryptoData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetboxCryptoData.Unmarshal(m, b)
}
func (m *AssetboxCryptoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetboxCryptoData.Marshal(b, m, deterministic)
}
func (m *AssetboxCryptoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetboxCryptoData.Merge(m, src)
}
func (m *AssetboxCryptoData) XXX_Size() int {
	return xxx_messageInfo_AssetboxCryptoData.Size(m)
}
func (m *AssetboxCryptoData) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetboxCryptoData.DiscardUnknown(m)
}

var xxx_messageInfo_AssetboxCryptoData proto.InternalMessageInfo

func (m *AssetboxCryptoData) GetWallet() []byte {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func (m *AssetboxCryptoData) GetPassphrase() []byte {
	if m != nil {
		return m.Passphrase
	}
	return nil
}

func init() {
	proto.RegisterType((*AssetboxBalance)(nil), "dto.AssetboxBalance")
	proto.RegisterType((*AssetboxCryptoData)(nil), "dto.AssetboxCryptoData")
}

func init() { proto.RegisterFile("assetbox.proto", fileDescriptor_661719b864488047) }

var fileDescriptor_661719b864488047 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0x0b, 0xc2, 0x40,
	0x0c, 0x46, 0xa9, 0x42, 0xd1, 0x50, 0x14, 0x6e, 0x90, 0xe2, 0x20, 0xd2, 0xc9, 0xa9, 0x8b, 0xa3,
	0x93, 0x55, 0x37, 0xa7, 0x8e, 0x6e, 0x69, 0x0d, 0x38, 0x9c, 0x97, 0x23, 0x97, 0x62, 0xfb, 0xef,
	0xc5, 0xa3, 0x8a, 0xe3, 0x7b, 0x09, 0x8f, 0x0f, 0x16, 0x18, 0x02, 0x69, 0xc3, 0x7d, 0xe9, 0x85,
	0x95, 0xcd, 0xf4, 0xae, 0x5c, 0x5c, 0x60, 0x79, 0x1c, 0x75, 0x85, 0x16, 0x5d, 0x4b, 0x66, 0x0d,
	0xb3, 0xb6, 0x13, 0x21, 0xd7, 0x0e, 0x79, 0xb2, 0x4d, 0x76, 0xf3, 0xfa, 0xc7, 0x66, 0x05, 0x29,
	0x3e, 0xb9, 0x73, 0x9a, 0x4f, 0xe2, 0x65, 0xa4, 0xe2, 0x0a, 0xe6, 0x9b, 0x39, 0xc9, 0xe0, 0x95,
	0xcf, 0xa8, 0xf8, 0xf9, 0x7e, 0xa1, 0xb5, 0xa4, 0xb1, 0x93, 0xd5, 0x23, 0x99, 0x0d, 0x80, 0xc7,
	0x10, 0xfc, 0x43, 0x30, 0x50, 0x2c, 0x65, 0xf5, 0x9f, 0xa9, 0xb2, 0x1b, 0x94, 0x07, 0xea, 0x95,
	0xc4, 0xa1, 0x6d, 0xd2, 0x38, 0x77, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x63, 0x74, 0x75, 0xf5,
	0xc0, 0x00, 0x00, 0x00,
}
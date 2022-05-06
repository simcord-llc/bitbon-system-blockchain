// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assetbox-delete.proto

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

type DeleteAssetboxRequest struct {
	// routing key: r.assetbox.SENDER.DeleteAssetboxRequest
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string              `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	CryptoData           *AssetboxCryptoData `protobuf:"bytes,3,opt,name=crypto_data,json=cryptoData,proto3" json:"crypto_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeleteAssetboxRequest) Reset()         { *m = DeleteAssetboxRequest{} }
func (m *DeleteAssetboxRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAssetboxRequest) ProtoMessage()    {}
func (*DeleteAssetboxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d62a22667de34299, []int{0}
}

func (m *DeleteAssetboxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAssetboxRequest.Unmarshal(m, b)
}
func (m *DeleteAssetboxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAssetboxRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAssetboxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAssetboxRequest.Merge(m, src)
}
func (m *DeleteAssetboxRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAssetboxRequest.Size(m)
}
func (m *DeleteAssetboxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAssetboxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAssetboxRequest proto.InternalMessageInfo

func (m *DeleteAssetboxRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteAssetboxRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DeleteAssetboxRequest) GetCryptoData() *AssetboxCryptoData {
	if m != nil {
		return m.CryptoData
	}
	return nil
}

type DeleteAssetboxResponse struct {
	// r.assetbox.bitbon-node.DeleteAssetboxResponse
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAssetboxResponse) Reset()         { *m = DeleteAssetboxResponse{} }
func (m *DeleteAssetboxResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAssetboxResponse) ProtoMessage()    {}
func (*DeleteAssetboxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d62a22667de34299, []int{1}
}

func (m *DeleteAssetboxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAssetboxResponse.Unmarshal(m, b)
}
func (m *DeleteAssetboxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAssetboxResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAssetboxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAssetboxResponse.Merge(m, src)
}
func (m *DeleteAssetboxResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAssetboxResponse.Size(m)
}
func (m *DeleteAssetboxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAssetboxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAssetboxResponse proto.InternalMessageInfo

func (m *DeleteAssetboxResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteAssetboxResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteAssetboxRequest)(nil), "dto.DeleteAssetboxRequest")
	proto.RegisterType((*DeleteAssetboxResponse)(nil), "dto.DeleteAssetboxResponse")
}

func init() { proto.RegisterFile("assetbox-delete.proto", fileDescriptor_d62a22667de34299) }

var fileDescriptor_d62a22667de34299 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x4b, 0xc6, 0x30,
	0x10, 0xc5, 0x69, 0x3f, 0x54, 0xbc, 0xc8, 0x37, 0x04, 0xaa, 0xa1, 0x53, 0xe9, 0xd4, 0xc5, 0x0c,
	0x75, 0x11, 0x9c, 0xd4, 0xba, 0x38, 0x66, 0x74, 0x91, 0xb4, 0x77, 0x43, 0xa1, 0xf4, 0x6a, 0x72,
	0x42, 0xc5, 0x7f, 0x5e, 0x4c, 0xad, 0x83, 0x6e, 0x79, 0xbf, 0x97, 0x7b, 0x8f, 0x07, 0x85, 0x8f,
	0x91, 0xa4, 0xe7, 0xf5, 0x1a, 0x69, 0x22, 0x21, 0xbb, 0x04, 0x16, 0xd6, 0x07, 0x14, 0x2e, 0x15,
	0x85, 0xc0, 0x61, 0x23, 0xe5, 0x71, 0xff, 0xb8, 0xe9, 0xfa, 0x13, 0x8a, 0x2e, 0x5d, 0xdc, 0xff,
	0x70, 0x47, 0x6f, 0xef, 0x14, 0x45, 0x1f, 0x21, 0x1f, 0xd1, 0x64, 0x55, 0xd6, 0x9c, 0xbb, 0x7c,
	0x44, 0x6d, 0xe0, 0xcc, 0x23, 0x06, 0x8a, 0xd1, 0xe4, 0x09, 0xee, 0x52, 0xdf, 0x82, 0x1a, 0xc2,
	0xc7, 0x22, 0xfc, 0x8a, 0x5e, 0xbc, 0x39, 0x54, 0x59, 0xa3, 0xda, 0x2b, 0x8b, 0xc2, 0x76, 0x0f,
	0x7d, 0x4c, 0x7e, 0xe7, 0xc5, 0x3b, 0x18, 0x7e, 0xdf, 0xf5, 0x33, 0x5c, 0xfe, 0x2d, 0x8f, 0x0b,
	0xcf, 0x91, 0xfe, 0xb5, 0x57, 0x70, 0x92, 0x56, 0xa4, 0x6e, 0xd5, 0x42, 0x4a, 0x7f, 0xfa, 0x26,
	0x6e, 0x33, 0x1e, 0x2e, 0x5e, 0xc0, 0xde, 0xd1, 0x2a, 0x14, 0x66, 0x3f, 0xf5, 0xa7, 0x69, 0xdd,
	0xcd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x87, 0xa2, 0x30, 0x18, 0x01, 0x00, 0x00,
}

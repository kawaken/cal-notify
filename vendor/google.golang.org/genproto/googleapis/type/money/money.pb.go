// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/type/money/money.proto
// DO NOT EDIT!

/*
Package google_type is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/type/money/money.proto

It has these top-level messages:
	Money
*/
package google_type // import "google.golang.org/genproto/googleapis/type/money"

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents an amount of money with its currency type.
type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int32 `protobuf:"varint,3,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Money)(nil), "google.type.Money")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/type/money/money.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xeb, 0x97, 0x54,
	0x16, 0xa4, 0xea, 0xe7, 0xe6, 0xe7, 0xa5, 0x56, 0x42, 0x48, 0x3d, 0xb0, 0x0a, 0x21, 0x6e, 0xa8,
	0x6e, 0x90, 0xb4, 0x52, 0x04, 0x17, 0xab, 0x2f, 0x48, 0x4e, 0x48, 0x99, 0x8b, 0x37, 0xb9, 0xb4,
	0xa8, 0x28, 0x35, 0x2f, 0xb9, 0x32, 0x3e, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x88, 0x07, 0x26, 0xe8, 0x9c, 0x9f, 0x92, 0x2a, 0x24, 0xc2, 0xc5, 0x5a, 0x9a, 0x97, 0x59,
	0x52, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe1, 0x80, 0x44, 0xf3, 0x12, 0xf3, 0xf2,
	0x8b, 0x25, 0x98, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x27, 0x55, 0x2e, 0xfe, 0xe4, 0xfc,
	0x5c, 0x3d, 0x24, 0xcb, 0x9c, 0xb8, 0xc0, 0x56, 0x05, 0x80, 0x5c, 0x11, 0xc0, 0xb8, 0x88, 0x89,
	0xd9, 0x3d, 0x24, 0x20, 0x89, 0x0d, 0xec, 0x28, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10,
	0xfa, 0x29, 0x0d, 0xd4, 0x00, 0x00, 0x00,
}

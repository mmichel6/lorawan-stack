// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/_api.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

/*
The Things Network v3 API
*/

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("lorawan-stack/api/_api.proto", fileDescriptor__api_51166803f86d0f98) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/_api.proto", fileDescriptor__api_51166803f86d0f98)
}

var fileDescriptor__api_51166803f86d0f98 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xce, 0xa1, 0x4e, 0x03, 0x41,
	0x10, 0x06, 0xe0, 0x19, 0x83, 0x40, 0x20, 0xd0, 0xe4, 0x7f, 0x02, 0xd8, 0x15, 0x7d, 0x03, 0x1e,
	0x03, 0x43, 0xb6, 0xa4, 0xb9, 0x5e, 0x8e, 0xec, 0x6e, 0xae, 0x13, 0x6a, 0x2b, 0x2b, 0x91, 0x48,
	0x82, 0xaa, 0xac, 0xac, 0xac, 0xac, 0x3c, 0x79, 0xf2, 0x76, 0xc6, 0x9c, 0x3c, 0x79, 0x92, 0x84,
	0x60, 0xea, 0x3f, 0xf1, 0xdd, 0x3e, 0xbc, 0xa7, 0x36, 0x6c, 0x43, 0x7c, 0xda, 0x48, 0x78, 0x6b,
	0x7c, 0xc8, 0xb5, 0x7f, 0x0d, 0xb9, 0x76, 0xb9, 0x4d, 0x92, 0xee, 0xef, 0x44, 0xa2, 0xfb, 0x17,
	0xee, 0x63, 0xf1, 0xfc, 0xc3, 0x97, 0x02, 0xee, 0x0a, 0xb8, 0x2f, 0xa0, 0xa1, 0x80, 0xc6, 0x02,
	0x9a, 0x0a, 0x68, 0x2e, 0xe0, 0x9d, 0x82, 0xf7, 0x0a, 0x3a, 0x28, 0xf8, 0xa8, 0xa0, 0x93, 0x82,
	0xce, 0x0a, 0xba, 0x28, 0xb8, 0x53, 0x70, 0xaf, 0xa0, 0x41, 0xc1, 0xa3, 0x82, 0x26, 0x05, 0xcf,
	0x0a, 0xda, 0x19, 0x68, 0x6f, 0xe0, 0x4f, 0x03, 0x7d, 0x19, 0xf8, 0xdb, 0x40, 0x07, 0x03, 0x1d,
	0x0d, 0x7c, 0x32, 0xf0, 0xd9, 0xc0, 0x2f, 0x8f, 0x55, 0x72, 0xb2, 0x5e, 0xc9, 0xba, 0x8e, 0xd5,
	0xc6, 0xc5, 0x95, 0x6c, 0x53, 0xdb, 0xf8, 0xeb, 0x79, 0x6e, 0x2a, 0x2f, 0x12, 0xf3, 0x72, 0x79,
	0xf3, 0x77, 0x5f, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x81, 0x44, 0x5f, 0xdb, 0x00, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ketty.proto

/*
Package ketty is a generated protocol buffer package.

It is generated from these files:
	ketty.proto

It has these top-level messages:
*/
package ketty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_UseKettyHttpExtend = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10001,
	Name:          "use_ketty_http_extend",
	Tag:           "varint,10001,opt,name=use_ketty_http_extend,json=useKettyHttpExtend",
	Filename:      "ketty.proto",
}

var E_MultiTransport = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10004,
	Name:          "multi_transport",
	Tag:           "varint,10004,opt,name=multi_transport,json=multiTransport",
	Filename:      "ketty.proto",
}

var E_Marshal = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         10002,
	Name:          "marshal",
	Tag:           "bytes,10002,opt,name=marshal",
	Filename:      "ketty.proto",
}

var E_Transport = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         10003,
	Name:          "transport",
	Tag:           "bytes,10003,opt,name=transport",
	Filename:      "ketty.proto",
}

var E_JsonHyaline = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10005,
	Name:          "json_hyaline",
	Tag:           "varint,10005,opt,name=json_hyaline,json=jsonHyaline",
	Filename:      "ketty.proto",
}

var E_JsonAllowOmitemptyRepeated = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20001,
	Name:          "json_allow_omitempty_repeated",
	Tag:           "varint,20001,opt,name=json_allow_omitempty_repeated,json=jsonAllowOmitemptyRepeated",
	Filename:      "ketty.proto",
}

var E_JsonAllowOmitempty = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         30001,
	Name:          "json_allow_omitempty",
	Tag:           "varint,30001,opt,name=json_allow_omitempty,json=jsonAllowOmitempty",
	Filename:      "ketty.proto",
}

func init() {
	proto.RegisterExtension(E_UseKettyHttpExtend)
	proto.RegisterExtension(E_MultiTransport)
	proto.RegisterExtension(E_Marshal)
	proto.RegisterExtension(E_Transport)
	proto.RegisterExtension(E_JsonHyaline)
	proto.RegisterExtension(E_JsonAllowOmitemptyRepeated)
	proto.RegisterExtension(E_JsonAllowOmitempty)
}

func init() { proto.RegisterFile("ketty.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xcf, 0xcf, 0x4a, 0xfc, 0x30,
	0x10, 0xc0, 0x71, 0xf6, 0xf2, 0xfb, 0xb9, 0x59, 0x51, 0x28, 0x0a, 0xb2, 0xb8, 0xb8, 0x47, 0x4f,
	0xdd, 0x83, 0xb7, 0x1c, 0x04, 0x11, 0x65, 0x45, 0x74, 0xb1, 0x7a, 0x0f, 0x59, 0x3b, 0xb6, 0xd1,
	0xb4, 0x13, 0x92, 0x29, 0xda, 0xc7, 0xf0, 0xdf, 0xcd, 0x8b, 0x8f, 0xe0, 0x6b, 0xf9, 0x14, 0xd2,
	0xc4, 0xea, 0x61, 0x17, 0x8a, 0xc7, 0x40, 0xbe, 0x9f, 0x99, 0x61, 0x83, 0x3b, 0x20, 0xaa, 0x63,
	0x63, 0x91, 0x70, 0x38, 0xce, 0x10, 0x33, 0x0d, 0x13, 0xff, 0x9a, 0x57, 0x37, 0x93, 0x14, 0xdc,
	0xb5, 0x55, 0x86, 0xd0, 0x86, 0x1f, 0xfc, 0x92, 0x6d, 0x56, 0x0e, 0x84, 0x8f, 0x44, 0x4e, 0x64,
	0x04, 0x3c, 0x10, 0x94, 0x69, 0xb4, 0x13, 0x87, 0x36, 0x6e, 0xdb, 0xf8, 0x0c, 0x9c, 0x93, 0x19,
	0xcc, 0x0c, 0x29, 0x2c, 0xdd, 0xd6, 0xe3, 0xf9, 0xb8, 0xb7, 0xbb, 0x92, 0x44, 0x95, 0x83, 0xd3,
	0xa6, 0x9e, 0x12, 0x99, 0x23, 0xdf, 0xf2, 0x13, 0xb6, 0x5e, 0x54, 0x9a, 0x94, 0x20, 0x2b, 0x4b,
	0x67, 0xd0, 0x52, 0x37, 0xf7, 0x12, 0xb8, 0x35, 0x1f, 0x5e, 0xb5, 0x1d, 0xe7, 0xec, 0x7f, 0x21,
	0xad, 0xcb, 0xa5, 0xee, 0x26, 0x9e, 0x1a, 0xa2, 0x9f, 0xb4, 0x01, 0xdf, 0x67, 0xfd, 0x3f, 0x2c,
	0xf0, 0x1c, 0xea, 0xdf, 0x84, 0x1f, 0xb2, 0xd5, 0x5b, 0x87, 0xa5, 0xc8, 0x6b, 0xa9, 0x55, 0x09,
	0xdd, 0xc4, 0x6b, 0xb8, 0x61, 0xd0, 0x54, 0xd3, 0x10, 0x71, 0xc9, 0x46, 0x1e, 0x91, 0x5a, 0xe3,
	0xbd, 0xc0, 0x42, 0x11, 0x14, 0x86, 0x6a, 0x61, 0xc1, 0x80, 0x24, 0x48, 0xa3, 0xed, 0x05, 0xf5,
	0x58, 0xe9, 0x1f, 0xf2, 0xfd, 0xad, 0xe7, 0xcd, 0x61, 0x83, 0x1c, 0x34, 0xc6, 0xac, 0x25, 0x92,
	0x6f, 0x81, 0x5f, 0xb0, 0x8d, 0x65, 0x23, 0xa2, 0xd1, 0x12, 0x19, 0x74, 0xda, 0xd2, 0x1f, 0x9f,
	0x81, 0x8e, 0x16, 0xe9, 0xf9, 0x3f, 0x9f, 0xec, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x80,
	0xaf, 0x19, 0x4e, 0x02, 0x00, 0x00,
}

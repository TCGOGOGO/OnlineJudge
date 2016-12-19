// Code generated by protoc-gen-go.
// source: api/ojinfo.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OJInfo struct {
	OjId       int64  `protobuf:"varint,1,opt,name=oj_id,json=ojId" json:"oj_id,omitempty"`
	OjName     string `protobuf:"bytes,2,opt,name=oj_name,json=ojName" json:"oj_name,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Int64Io    string `protobuf:"bytes,4,opt,name=int64io" json:"int64io,omitempty"`
	Javaclass  string `protobuf:"bytes,5,opt,name=javaclass" json:"javaclass,omitempty"`
	Status     string `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	StatusInfo string `protobuf:"bytes,7,opt,name=status_info,json=statusInfo" json:"status_info,omitempty"`
	Lastcheck  string `protobuf:"bytes,8,opt,name=lastcheck" json:"lastcheck,omitempty"`
}

func (m *OJInfo) Reset()                    { *m = OJInfo{} }
func (m *OJInfo) String() string            { return proto.CompactTextString(m) }
func (*OJInfo) ProtoMessage()               {}
func (*OJInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *OJInfo) GetOjId() int64 {
	if m != nil {
		return m.OjId
	}
	return 0
}

func (m *OJInfo) GetOjName() string {
	if m != nil {
		return m.OjName
	}
	return ""
}

func (m *OJInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OJInfo) GetInt64Io() string {
	if m != nil {
		return m.Int64Io
	}
	return ""
}

func (m *OJInfo) GetJavaclass() string {
	if m != nil {
		return m.Javaclass
	}
	return ""
}

func (m *OJInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OJInfo) GetStatusInfo() string {
	if m != nil {
		return m.StatusInfo
	}
	return ""
}

func (m *OJInfo) GetLastcheck() string {
	if m != nil {
		return m.Lastcheck
	}
	return ""
}

func init() {
	proto.RegisterType((*OJInfo)(nil), "api.OJInfo")
}

func init() { proto.RegisterFile("api/ojinfo.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8f, 0x41, 0x4e, 0x85, 0x30,
	0x10, 0x86, 0x53, 0x79, 0xaf, 0xc8, 0xb8, 0x31, 0x35, 0xd1, 0x59, 0x90, 0x48, 0x5c, 0xb1, 0xd2,
	0x85, 0xc6, 0x3b, 0xe0, 0x42, 0x13, 0x2e, 0x40, 0x46, 0x28, 0xb1, 0x15, 0x3a, 0x84, 0x56, 0xae,
	0xeb, 0x55, 0x0c, 0xad, 0xfa, 0x76, 0xf3, 0x7f, 0x5f, 0x32, 0xf3, 0x0f, 0x5c, 0xd2, 0x62, 0x1e,
	0xd8, 0x1a, 0x37, 0xf2, 0xfd, 0xb2, 0x72, 0x60, 0x95, 0xd1, 0x62, 0xee, 0xbe, 0x05, 0xc8, 0xb7,
	0x97, 0xc6, 0x8d, 0xac, 0xae, 0xe0, 0xc8, 0xb6, 0x33, 0x03, 0x8a, 0x4a, 0xd4, 0x59, 0x7b, 0x60,
	0xdb, 0x0c, 0xea, 0x06, 0x72, 0xb6, 0x9d, 0xa3, 0x59, 0xe3, 0x59, 0x25, 0xea, 0xa2, 0x95, 0x6c,
	0x5f, 0x69, 0xd6, 0x0a, 0x21, 0xdf, 0xf4, 0xea, 0x0d, 0x3b, 0xcc, 0xa2, 0xf8, 0x8b, 0xbb, 0x31,
	0x2e, 0x3c, 0x3f, 0x19, 0xc6, 0x43, 0x32, 0xbf, 0x51, 0x95, 0x50, 0x58, 0xda, 0xa8, 0x9f, 0xc8,
	0x7b, 0x3c, 0x46, 0x77, 0x02, 0xea, 0x1a, 0xa4, 0x0f, 0x14, 0xbe, 0x3c, 0xca, 0x74, 0x29, 0x25,
	0x75, 0x0b, 0x17, 0x69, 0xea, 0xf6, 0xf2, 0x98, 0x47, 0x09, 0x09, 0xc5, 0xe2, 0x25, 0x14, 0x13,
	0xf9, 0xd0, 0x7f, 0xe8, 0xfe, 0x13, 0xcf, 0xd3, 0xda, 0x7f, 0xf0, 0x2e, 0xe3, 0xb7, 0x8f, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x8b, 0xb4, 0x9c, 0x01, 0x01, 0x00, 0x00,
}

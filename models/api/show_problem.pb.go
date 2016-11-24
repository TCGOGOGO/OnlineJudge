// Code generated by protoc-gen-go.
// source: api/show_problem.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/show_problem.proto
	api/submit.proto
	api/test.proto

It has these top-level messages:
	ShowProblemRequest
	ShowProblemResponse
	SubmitRequest
	SubmitResponse
	Test
*/
package api

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

type ShowProblemRequest_Element int32

const (
	ShowProblemRequest_TITLE         ShowProblemRequest_Element = 0
	ShowProblemRequest_TIME_LIMIT    ShowProblemRequest_Element = 1
	ShowProblemRequest_MEMORY_LIMIT  ShowProblemRequest_Element = 2
	ShowProblemRequest_DESCRIPTION   ShowProblemRequest_Element = 3
	ShowProblemRequest_INPUT         ShowProblemRequest_Element = 4
	ShowProblemRequest_OUTPUT        ShowProblemRequest_Element = 5
	ShowProblemRequest_SAMPLE_INPUT  ShowProblemRequest_Element = 6
	ShowProblemRequest_SAMPLE_OUTPUT ShowProblemRequest_Element = 7
	ShowProblemRequest_SOURCE        ShowProblemRequest_Element = 8
	ShowProblemRequest_HINT          ShowProblemRequest_Element = 9
)

var ShowProblemRequest_Element_name = map[int32]string{
	0: "TITLE",
	1: "TIME_LIMIT",
	2: "MEMORY_LIMIT",
	3: "DESCRIPTION",
	4: "INPUT",
	5: "OUTPUT",
	6: "SAMPLE_INPUT",
	7: "SAMPLE_OUTPUT",
	8: "SOURCE",
	9: "HINT",
}
var ShowProblemRequest_Element_value = map[string]int32{
	"TITLE":         0,
	"TIME_LIMIT":    1,
	"MEMORY_LIMIT":  2,
	"DESCRIPTION":   3,
	"INPUT":         4,
	"OUTPUT":        5,
	"SAMPLE_INPUT":  6,
	"SAMPLE_OUTPUT": 7,
	"SOURCE":        8,
	"HINT":          9,
}

func (x ShowProblemRequest_Element) String() string {
	return proto.EnumName(ShowProblemRequest_Element_name, int32(x))
}
func (ShowProblemRequest_Element) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type ShowProblemRequest struct {
	ContestId         uint64                       `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid        string                       `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	IfNeedAllElements bool                         `protobuf:"varint,3,opt,name=if_need_all_elements,json=ifNeedAllElements" json:"if_need_all_elements,omitempty"`
	ElementsNeeded    []ShowProblemRequest_Element `protobuf:"varint,4,rep,packed,name=elements_needed,json=elementsNeeded,enum=api.ShowProblemRequest_Element" json:"elements_needed,omitempty"`
}

func (m *ShowProblemRequest) Reset()                    { *m = ShowProblemRequest{} }
func (m *ShowProblemRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemRequest) ProtoMessage()               {}
func (*ShowProblemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ShowProblemRequest) GetContestId() uint64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ShowProblemRequest) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

func (m *ShowProblemRequest) GetIfNeedAllElements() bool {
	if m != nil {
		return m.IfNeedAllElements
	}
	return false
}

func (m *ShowProblemRequest) GetElementsNeeded() []ShowProblemRequest_Element {
	if m != nil {
		return m.ElementsNeeded
	}
	return nil
}

// response
type ShowProblemResponse struct {
	ContestId    uint64                          `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid   uint64                          `protobuf:"varint,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Title        string                          `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	TimeLimit    int32                           `protobuf:"varint,4,opt,name=time_limit,json=timeLimit" json:"time_limit,omitempty"`
	MemoryLimit  int32                           `protobuf:"varint,5,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	Input        string                          `protobuf:"bytes,6,opt,name=input" json:"input,omitempty"`
	Output       string                          `protobuf:"bytes,7,opt,name=output" json:"output,omitempty"`
	SampleInput  string                          `protobuf:"bytes,8,opt,name=sample_input,json=sampleInput" json:"sample_input,omitempty"`
	SampleOutput string                          `protobuf:"bytes,9,opt,name=sample_output,json=sampleOutput" json:"sample_output,omitempty"`
	Source       string                          `protobuf:"bytes,19,opt,name=source" json:"source,omitempty"`
	Hint         string                          `protobuf:"bytes,11,opt,name=hint" json:"hint,omitempty"`
	Languages    []*ShowProblemResponse_Language `protobuf:"bytes,12,rep,name=languages" json:"languages,omitempty"`
}

func (m *ShowProblemResponse) Reset()                    { *m = ShowProblemResponse{} }
func (m *ShowProblemResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemResponse) ProtoMessage()               {}
func (*ShowProblemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ShowProblemResponse) GetContestId() uint64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ShowProblemResponse) GetProblemSid() uint64 {
	if m != nil {
		return m.ProblemSid
	}
	return 0
}

func (m *ShowProblemResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ShowProblemResponse) GetTimeLimit() int32 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *ShowProblemResponse) GetMemoryLimit() int32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *ShowProblemResponse) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ShowProblemResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *ShowProblemResponse) GetSampleInput() string {
	if m != nil {
		return m.SampleInput
	}
	return ""
}

func (m *ShowProblemResponse) GetSampleOutput() string {
	if m != nil {
		return m.SampleOutput
	}
	return ""
}

func (m *ShowProblemResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ShowProblemResponse) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *ShowProblemResponse) GetLanguages() []*ShowProblemResponse_Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

type ShowProblemResponse_Language struct {
	Language string `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Code     int32  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *ShowProblemResponse_Language) Reset()                    { *m = ShowProblemResponse_Language{} }
func (m *ShowProblemResponse_Language) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemResponse_Language) ProtoMessage()               {}
func (*ShowProblemResponse_Language) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ShowProblemResponse_Language) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *ShowProblemResponse_Language) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*ShowProblemRequest)(nil), "api.ShowProblemRequest")
	proto.RegisterType((*ShowProblemResponse)(nil), "api.ShowProblemResponse")
	proto.RegisterType((*ShowProblemResponse_Language)(nil), "api.ShowProblemResponse.Language")
	proto.RegisterEnum("api.ShowProblemRequest_Element", ShowProblemRequest_Element_name, ShowProblemRequest_Element_value)
}

func init() { proto.RegisterFile("api/show_problem.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xcb, 0x6e, 0x9b, 0x4c,
	0x14, 0xc7, 0x3f, 0xcc, 0xc5, 0x70, 0x70, 0x1c, 0x32, 0xb1, 0x2c, 0x14, 0xe9, 0x53, 0x88, 0xbb,
	0x61, 0xe5, 0x48, 0xe9, 0xae, 0x9b, 0x2a, 0x4a, 0x91, 0x82, 0xe4, 0x9b, 0xc6, 0x78, 0xd1, 0x15,
	0x22, 0x66, 0x12, 0x8f, 0x04, 0x0c, 0x35, 0x83, 0xa2, 0x3e, 0x4a, 0xdf, 0xa9, 0x6f, 0xd2, 0x97,
	0xa8, 0xe6, 0xe2, 0x56, 0x51, 0x76, 0xdd, 0xcd, 0xf9, 0xfd, 0x2f, 0xc0, 0x01, 0x60, 0x5a, 0xb4,
	0xf4, 0xb6, 0x3b, 0xb0, 0xd7, 0xbc, 0x3d, 0xb2, 0xa7, 0x8a, 0xd4, 0xf3, 0xf6, 0xc8, 0x38, 0x43,
	0x66, 0xd1, 0xd2, 0xd9, 0xaf, 0x01, 0xa0, 0xed, 0x81, 0xbd, 0x6e, 0x94, 0x84, 0xc9, 0xb7, 0x9e,
	0x74, 0x1c, 0xfd, 0x0f, 0xb0, 0x67, 0x0d, 0x27, 0x1d, 0xcf, 0x69, 0x19, 0x1a, 0x91, 0x11, 0x5b,
	0xd8, 0xd3, 0x24, 0x2d, 0xd1, 0x35, 0xf8, 0xba, 0x2b, 0xef, 0x68, 0x19, 0x0e, 0x22, 0x23, 0xf6,
	0x30, 0x68, 0xb4, 0xa5, 0x25, 0xba, 0x85, 0x09, 0x7d, 0xce, 0x1b, 0x42, 0xca, 0xbc, 0xa8, 0xaa,
	0x9c, 0x54, 0xa4, 0x26, 0x0d, 0xef, 0x42, 0x33, 0x32, 0x62, 0x17, 0x5f, 0xd0, 0xe7, 0x15, 0x21,
	0xe5, 0x7d, 0x55, 0x25, 0x5a, 0x40, 0x8f, 0x70, 0x7e, 0x32, 0xc9, 0x18, 0x29, 0x43, 0x2b, 0x32,
	0xe3, 0xf1, 0xdd, 0xf5, 0xbc, 0x68, 0xe9, 0xfc, 0xfd, 0x2d, 0xce, 0x75, 0x14, 0x8f, 0x4f, 0xb9,
	0x95, 0x8c, 0xcd, 0x7e, 0x18, 0x30, 0xd4, 0x1a, 0xf2, 0xc0, 0xce, 0xd2, 0x6c, 0x91, 0x04, 0xff,
	0xa1, 0x31, 0x40, 0x96, 0x2e, 0x93, 0x7c, 0x91, 0x2e, 0xd3, 0x2c, 0x30, 0x50, 0x00, 0xa3, 0x65,
	0xb2, 0x5c, 0xe3, 0xaf, 0x9a, 0x0c, 0xd0, 0x39, 0xf8, 0x5f, 0x92, 0xed, 0x03, 0x4e, 0x37, 0x59,
	0xba, 0x5e, 0x05, 0xa6, 0x48, 0xa7, 0xab, 0xcd, 0x2e, 0x0b, 0x2c, 0x04, 0xe0, 0xac, 0x77, 0x99,
	0x38, 0xdb, 0x22, 0xb9, 0xbd, 0x5f, 0x6e, 0x16, 0x49, 0xae, 0x54, 0x07, 0x5d, 0xc0, 0x99, 0x26,
	0xda, 0x34, 0x14, 0x81, 0xed, 0x7a, 0x87, 0x1f, 0x92, 0xc0, 0x45, 0x2e, 0x58, 0x8f, 0xe9, 0x2a,
	0x0b, 0xbc, 0xd9, 0x4f, 0x13, 0x2e, 0xdf, 0x3c, 0x4a, 0xd7, 0xb2, 0xa6, 0x23, 0xff, 0xb0, 0x6e,
	0xeb, 0xcd, 0xba, 0x27, 0x60, 0x73, 0xca, 0x2b, 0x22, 0xf7, 0xeb, 0x61, 0x35, 0x88, 0x56, 0x4e,
	0x6b, 0x92, 0x57, 0xb4, 0xa6, 0x3c, 0xb4, 0x22, 0x23, 0xb6, 0xb1, 0x27, 0xc8, 0x42, 0x00, 0x74,
	0x03, 0xa3, 0x9a, 0xd4, 0xec, 0xf8, 0x5d, 0x1b, 0x6c, 0x69, 0xf0, 0x15, 0x53, 0x96, 0x09, 0xd8,
	0xb4, 0x69, 0x7b, 0x1e, 0x3a, 0xaa, 0x57, 0x0e, 0x68, 0x0a, 0x0e, 0xeb, 0xb9, 0xc0, 0x43, 0x89,
	0xf5, 0x24, 0x0a, 0xbb, 0xa2, 0x6e, 0x2b, 0x92, 0xab, 0x90, 0x2b, 0x55, 0x5f, 0xb1, 0x54, 0x46,
	0x3f, 0xc0, 0x99, 0xb6, 0xe8, 0x06, 0x4f, 0x7a, 0x74, 0x6e, 0xad, 0x7a, 0xa6, 0xe0, 0x74, 0xac,
	0x3f, 0xee, 0x49, 0x78, 0xa9, 0xfa, 0xd5, 0x84, 0x10, 0x58, 0x07, 0xda, 0xf0, 0xd0, 0x97, 0x54,
	0x9e, 0xd1, 0x67, 0xf0, 0xaa, 0xa2, 0x79, 0xe9, 0x8b, 0x17, 0xd2, 0x85, 0xa3, 0xc8, 0x8c, 0xfd,
	0xbb, 0x9b, 0xf7, 0x5f, 0x8c, 0x5a, 0xf3, 0x7c, 0xa1, 0x9d, 0xf8, 0x6f, 0xe6, 0xea, 0x13, 0xb8,
	0x27, 0x8c, 0xae, 0xc0, 0x3d, 0x09, 0xf2, 0x25, 0x78, 0xf8, 0xcf, 0x2c, 0x2e, 0xbe, 0x67, 0x25,
	0x91, 0xcb, 0xb7, 0xb1, 0x3c, 0x3f, 0x39, 0xf2, 0x47, 0xfa, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0x9c, 0x90, 0x0d, 0x62, 0x03, 0x00, 0x00,
}

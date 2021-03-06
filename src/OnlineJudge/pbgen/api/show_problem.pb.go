// Code generated by protoc-gen-go.
// source: api/show_problem.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShowProblemRequest struct {
	ContestId  uint64 `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
}

func (m *ShowProblemRequest) Reset()                    { *m = ShowProblemRequest{} }
func (m *ShowProblemRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemRequest) ProtoMessage()               {}
func (*ShowProblemRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

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

type Problem struct {
	Title        string           `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Limits       []*Problem_Limit `protobuf:"bytes,2,rep,name=limits" json:"limits,omitempty"`
	Description  string           `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Input        string           `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	Output       string           `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	SampleInput  string           `protobuf:"bytes,6,opt,name=sample_input,json=sampleInput" json:"sample_input,omitempty"`
	SampleOutput string           `protobuf:"bytes,7,opt,name=sample_output,json=sampleOutput" json:"sample_output,omitempty"`
	Source       string           `protobuf:"bytes,8,opt,name=source" json:"source,omitempty"`
	Hint         string           `protobuf:"bytes,9,opt,name=hint" json:"hint,omitempty"`
	Hide         bool             `protobuf:"varint,10,opt,name=hide" json:"hide,omitempty"`
}

func (m *Problem) Reset()                    { *m = Problem{} }
func (m *Problem) String() string            { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()               {}
func (*Problem) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *Problem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Problem) GetLimits() []*Problem_Limit {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *Problem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Problem) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *Problem) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Problem) GetSampleInput() string {
	if m != nil {
		return m.SampleInput
	}
	return ""
}

func (m *Problem) GetSampleOutput() string {
	if m != nil {
		return m.SampleOutput
	}
	return ""
}

func (m *Problem) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Problem) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *Problem) GetHide() bool {
	if m != nil {
		return m.Hide
	}
	return false
}

type Problem_Limit struct {
	Language      string `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	TimeLimit     int32  `protobuf:"varint,2,opt,name=time_limit,json=timeLimit" json:"time_limit,omitempty"`
	CaseTimeLimit int32  `protobuf:"varint,3,opt,name=case_time_limit,json=caseTimeLimit" json:"case_time_limit,omitempty"`
	MemoryLimit   int32  `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
}

func (m *Problem_Limit) Reset()                    { *m = Problem_Limit{} }
func (m *Problem_Limit) String() string            { return proto.CompactTextString(m) }
func (*Problem_Limit) ProtoMessage()               {}
func (*Problem_Limit) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1, 0} }

func (m *Problem_Limit) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Problem_Limit) GetTimeLimit() int32 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *Problem_Limit) GetCaseTimeLimit() int32 {
	if m != nil {
		return m.CaseTimeLimit
	}
	return 0
}

func (m *Problem_Limit) GetMemoryLimit() int32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

// response
type ShowProblemResponse struct {
	ContestId  uint64      `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid string      `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Problem    *Problem    `protobuf:"bytes,3,opt,name=problem" json:"problem,omitempty"`
	Languages  []*Language `protobuf:"bytes,4,rep,name=languages" json:"languages,omitempty"`
	Error      *Error      `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
}

func (m *ShowProblemResponse) Reset()                    { *m = ShowProblemResponse{} }
func (m *ShowProblemResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemResponse) ProtoMessage()               {}
func (*ShowProblemResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *ShowProblemResponse) GetContestId() uint64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ShowProblemResponse) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

func (m *ShowProblemResponse) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

func (m *ShowProblemResponse) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *ShowProblemResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowProblemRequest)(nil), "api.ShowProblemRequest")
	proto.RegisterType((*Problem)(nil), "api.Problem")
	proto.RegisterType((*Problem_Limit)(nil), "api.Problem.Limit")
	proto.RegisterType((*ShowProblemResponse)(nil), "api.ShowProblemResponse")
}

func init() { proto.RegisterFile("api/show_problem.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x5f, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xf1, 0xfa, 0x4f, 0xe2, 0x71, 0x42, 0xcb, 0xb4, 0x04, 0x11, 0x28, 0xf5, 0xa6, 0xb0,
	0x84, 0x16, 0x5c, 0x48, 0xcf, 0xd0, 0x87, 0x85, 0x85, 0x16, 0xed, 0xbe, 0x1b, 0xaf, 0x2d, 0x36,
	0x02, 0xdb, 0x52, 0x2d, 0x99, 0xa5, 0x87, 0xe8, 0x69, 0x7a, 0x8d, 0x1e, 0xaa, 0x68, 0xa4, 0xa4,
	0xe9, 0xf3, 0xbe, 0x79, 0x7e, 0xf3, 0x7d, 0x9f, 0x2c, 0xcd, 0xc0, 0xa6, 0xd1, 0xf2, 0xb3, 0x39,
	0xaa, 0xe7, 0x5a, 0x4f, 0xea, 0xb1, 0x17, 0x43, 0xa5, 0x27, 0x65, 0x15, 0xc6, 0x8d, 0x96, 0xdb,
	0xd7, 0xae, 0xd9, 0xaa, 0x61, 0x50, 0xa3, 0xc7, 0x5b, 0x74, 0xa4, 0x6f, 0xc6, 0xa7, 0xb9, 0x79,
	0x12, 0x9e, 0xed, 0x1e, 0x00, 0xef, 0x8f, 0xea, 0xf9, 0xbb, 0xf7, 0x73, 0xf1, 0x63, 0x16, 0xc6,
	0xe2, 0x3b, 0x80, 0x56, 0x8d, 0x56, 0x18, 0x5b, 0xcb, 0x8e, 0x45, 0x65, 0xb4, 0x4f, 0x78, 0x1e,
	0xc8, 0x6d, 0x87, 0xef, 0xa1, 0x08, 0x07, 0xd6, 0x46, 0x76, 0xec, 0xaa, 0x8c, 0xf6, 0x39, 0x87,
	0x80, 0xee, 0x65, 0xb7, 0xfb, 0x1d, 0xc3, 0x22, 0x44, 0xe2, 0x5b, 0x48, 0xad, 0xb4, 0xbd, 0xa0,
	0x98, 0x9c, 0xfb, 0x02, 0x3f, 0x42, 0xd6, 0xcb, 0x41, 0x5a, 0xc3, 0xae, 0xca, 0x78, 0x5f, 0x1c,
	0xb0, 0x6a, 0xb4, 0xac, 0x82, 0xa7, 0xba, 0x73, 0x2d, 0x1e, 0x14, 0x58, 0x42, 0xd1, 0x09, 0xd3,
	0x4e, 0x52, 0x5b, 0xa9, 0x46, 0x16, 0x53, 0xce, 0x25, 0x72, 0x67, 0xc8, 0x51, 0xcf, 0x96, 0x25,
	0xfe, 0x0c, 0x2a, 0x70, 0x03, 0x99, 0x9a, 0xad, 0xc3, 0x29, 0xe1, 0x50, 0xe1, 0x35, 0xac, 0x4c,
	0x33, 0xe8, 0x5e, 0xd4, 0xde, 0x94, 0xf9, 0x40, 0xcf, 0x6e, 0xc9, 0xfa, 0x01, 0xd6, 0x41, 0x12,
	0x12, 0x16, 0xa4, 0x09, 0xbe, 0x6f, 0x3e, 0x67, 0x03, 0x99, 0x51, 0xf3, 0xd4, 0x0a, 0xb6, 0xf4,
	0xf9, 0xbe, 0x42, 0x84, 0xe4, 0x28, 0x47, 0xcb, 0x72, 0xa2, 0xf4, 0xed, 0x59, 0x27, 0x18, 0x94,
	0xd1, 0x7e, 0xc9, 0xe9, 0x7b, 0xfb, 0x2b, 0x82, 0x94, 0x6e, 0x8a, 0x5b, 0x58, 0x9e, 0xe6, 0x12,
	0x9e, 0xe9, 0x5c, 0xbb, 0x59, 0x58, 0x39, 0x88, 0x9a, 0x1e, 0x83, 0xde, 0x3a, 0xe5, 0xb9, 0x23,
	0xde, 0x7a, 0x03, 0xaf, 0xda, 0xc6, 0x88, 0xfa, 0x42, 0x13, 0x93, 0x66, 0xed, 0xf0, 0xc3, 0x59,
	0x77, 0x0d, 0xab, 0x41, 0x0c, 0x6a, 0xfa, 0x19, 0x44, 0x09, 0x89, 0x0a, 0xcf, 0x48, 0xb2, 0xfb,
	0x13, 0xc1, 0x9b, 0xff, 0x96, 0xc1, 0x68, 0x35, 0x1a, 0xf1, 0xd2, 0x6d, 0xc0, 0x1b, 0x58, 0x84,
	0x8a, 0x7e, 0xad, 0x38, 0xac, 0x2e, 0x87, 0xcd, 0x4f, 0x4d, 0xfc, 0x04, 0xf9, 0xe9, 0xd6, 0x86,
	0x25, 0xb4, 0x16, 0x6b, 0x52, 0xde, 0x05, 0xca, 0xff, 0xf5, 0xb1, 0x84, 0x54, 0x4c, 0x93, 0x9a,
	0x68, 0xb6, 0xc5, 0x01, 0x48, 0xf8, 0xd5, 0x11, 0xee, 0x1b, 0x8f, 0x19, 0x6d, 0xf8, 0x97, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xcd, 0xbe, 0x99, 0x26, 0x03, 0x00, 0x00,
}

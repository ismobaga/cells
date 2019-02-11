// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

/*
Package log is a generated protocol buffer package.

It is generated from these files:
	log.proto

It has these top-level messages:
	RecorderPutResponse
	Log
	LogMessage
	ListLogRequest
	ListLogResponse
	DeleteLogsResponse
	TimeRangeResponse
	TimeRangeResult
	TimeRangeRequest
	TimeRangeCursor
*/
package log

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

// Relative links types.
// Note that First is time.Now() and last time.Unix(0).
// We added an unused NONE enum with value 0 to workaround 0 issues between JSON and proto3.
type RelType int32

const (
	RelType_NONE  RelType = 0
	RelType_FIRST RelType = 1
	RelType_PREV  RelType = 2
	RelType_NEXT  RelType = 3
	RelType_LAST  RelType = 4
)

var RelType_name = map[int32]string{
	0: "NONE",
	1: "FIRST",
	2: "PREV",
	3: "NEXT",
	4: "LAST",
}
var RelType_value = map[string]int32{
	"NONE":  0,
	"FIRST": 1,
	"PREV":  2,
	"NEXT":  3,
	"LAST":  4,
}

func (x RelType) String() string {
	return proto.EnumName(RelType_name, int32(x))
}
func (RelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Output Format
type ListLogRequest_LogFormat int32

const (
	ListLogRequest_JSON ListLogRequest_LogFormat = 0
	ListLogRequest_CSV  ListLogRequest_LogFormat = 1
	ListLogRequest_XLSX ListLogRequest_LogFormat = 2
)

var ListLogRequest_LogFormat_name = map[int32]string{
	0: "JSON",
	1: "CSV",
	2: "XLSX",
}
var ListLogRequest_LogFormat_value = map[string]int32{
	"JSON": 0,
	"CSV":  1,
	"XLSX": 2,
}

func (x ListLogRequest_LogFormat) String() string {
	return proto.EnumName(ListLogRequest_LogFormat_name, int32(x))
}
func (ListLogRequest_LogFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type RecorderPutResponse struct {
}

func (m *RecorderPutResponse) Reset()                    { *m = RecorderPutResponse{} }
func (m *RecorderPutResponse) String() string            { return proto.CompactTextString(m) }
func (*RecorderPutResponse) ProtoMessage()               {}
func (*RecorderPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Log is a generic message format used by the sync service
// to publish log messages to the various log repositories (typically, bleve).
type Log struct {
	Message map[string]string `protobuf:"bytes,1,rep,name=message" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Log) GetMessage() map[string]string {
	if m != nil {
		return m.Message
	}
	return nil
}

// LogMessage is the format used to transmit log messages to clients via the REST API.
type LogMessage struct {
	// Generic zap fields
	Ts     int32  `protobuf:"varint,1,opt,name=Ts" json:"Ts,omitempty"`
	Level  string `protobuf:"bytes,2,opt,name=Level" json:"Level,omitempty"`
	Logger string `protobuf:"bytes,3,opt,name=Logger" json:"Logger,omitempty"`
	Msg    string `protobuf:"bytes,4,opt,name=Msg" json:"Msg,omitempty"`
	// Pydio specific
	MsgId string `protobuf:"bytes,5,opt,name=MsgId" json:"MsgId,omitempty"`
	// User Info
	UserName  string   `protobuf:"bytes,6,opt,name=UserName" json:"UserName,omitempty"`
	UserUuid  string   `protobuf:"bytes,7,opt,name=UserUuid" json:"UserUuid,omitempty"`
	GroupPath string   `protobuf:"bytes,8,opt,name=GroupPath" json:"GroupPath,omitempty"`
	Profile   string   `protobuf:"bytes,16,opt,name=Profile" json:"Profile,omitempty"`
	RoleUuids []string `protobuf:"bytes,9,rep,name=RoleUuids" json:"RoleUuids,omitempty"`
	// Client info
	RemoteAddress string `protobuf:"bytes,10,opt,name=RemoteAddress" json:"RemoteAddress,omitempty"`
	UserAgent     string `protobuf:"bytes,11,opt,name=UserAgent" json:"UserAgent,omitempty"`
	HttpProtocol  string `protobuf:"bytes,12,opt,name=HttpProtocol" json:"HttpProtocol,omitempty"`
	// Tree Info
	NodeUuid string `protobuf:"bytes,13,opt,name=NodeUuid" json:"NodeUuid,omitempty"`
	NodePath string `protobuf:"bytes,14,opt,name=NodePath" json:"NodePath,omitempty"`
	WsUuid   string `protobuf:"bytes,15,opt,name=WsUuid" json:"WsUuid,omitempty"`
	WsScope  string `protobuf:"bytes,17,opt,name=WsScope" json:"WsScope,omitempty"`
	// Span Info
	SpanUuid       string `protobuf:"bytes,18,opt,name=SpanUuid" json:"SpanUuid,omitempty"`
	SpanParentUuid string `protobuf:"bytes,19,opt,name=SpanParentUuid" json:"SpanParentUuid,omitempty"`
	SpanRootUuid   string `protobuf:"bytes,20,opt,name=SpanRootUuid" json:"SpanRootUuid,omitempty"`
	// High Level Operation Info
	OperationUuid  string `protobuf:"bytes,21,opt,name=OperationUuid" json:"OperationUuid,omitempty"`
	OperationLabel string `protobuf:"bytes,22,opt,name=OperationLabel" json:"OperationLabel,omitempty"`
}

func (m *LogMessage) Reset()                    { *m = LogMessage{} }
func (m *LogMessage) String() string            { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()               {}
func (*LogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogMessage) GetTs() int32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *LogMessage) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LogMessage) GetLogger() string {
	if m != nil {
		return m.Logger
	}
	return ""
}

func (m *LogMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LogMessage) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *LogMessage) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LogMessage) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *LogMessage) GetGroupPath() string {
	if m != nil {
		return m.GroupPath
	}
	return ""
}

func (m *LogMessage) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *LogMessage) GetRoleUuids() []string {
	if m != nil {
		return m.RoleUuids
	}
	return nil
}

func (m *LogMessage) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func (m *LogMessage) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *LogMessage) GetHttpProtocol() string {
	if m != nil {
		return m.HttpProtocol
	}
	return ""
}

func (m *LogMessage) GetNodeUuid() string {
	if m != nil {
		return m.NodeUuid
	}
	return ""
}

func (m *LogMessage) GetNodePath() string {
	if m != nil {
		return m.NodePath
	}
	return ""
}

func (m *LogMessage) GetWsUuid() string {
	if m != nil {
		return m.WsUuid
	}
	return ""
}

func (m *LogMessage) GetWsScope() string {
	if m != nil {
		return m.WsScope
	}
	return ""
}

func (m *LogMessage) GetSpanUuid() string {
	if m != nil {
		return m.SpanUuid
	}
	return ""
}

func (m *LogMessage) GetSpanParentUuid() string {
	if m != nil {
		return m.SpanParentUuid
	}
	return ""
}

func (m *LogMessage) GetSpanRootUuid() string {
	if m != nil {
		return m.SpanRootUuid
	}
	return ""
}

func (m *LogMessage) GetOperationUuid() string {
	if m != nil {
		return m.OperationUuid
	}
	return ""
}

func (m *LogMessage) GetOperationLabel() string {
	if m != nil {
		return m.OperationLabel
	}
	return ""
}

// ListLogRequest launches a parameterised query in the log repository and streams the results.
type ListLogRequest struct {
	// Bleve-type Query stsring
	Query string `protobuf:"bytes,1,opt,name=Query" json:"Query,omitempty"`
	// Start at page
	Page int32 `protobuf:"varint,2,opt,name=Page" json:"Page,omitempty"`
	// Number of results
	Size   int32                    `protobuf:"varint,3,opt,name=Size" json:"Size,omitempty"`
	Format ListLogRequest_LogFormat `protobuf:"varint,4,opt,name=Format,enum=log.ListLogRequest_LogFormat" json:"Format,omitempty"`
}

func (m *ListLogRequest) Reset()                    { *m = ListLogRequest{} }
func (m *ListLogRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogRequest) ProtoMessage()               {}
func (*ListLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListLogRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ListLogRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListLogRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListLogRequest) GetFormat() ListLogRequest_LogFormat {
	if m != nil {
		return m.Format
	}
	return ListLogRequest_JSON
}

type ListLogResponse struct {
	LogMessage *LogMessage `protobuf:"bytes,1,opt,name=LogMessage" json:"LogMessage,omitempty"`
}

func (m *ListLogResponse) Reset()                    { *m = ListLogResponse{} }
func (m *ListLogResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogResponse) ProtoMessage()               {}
func (*ListLogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListLogResponse) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

type DeleteLogsResponse struct {
	Deleted int64 `protobuf:"varint,1,opt,name=Deleted" json:"Deleted,omitempty"`
}

func (m *DeleteLogsResponse) Reset()                    { *m = DeleteLogsResponse{} }
func (m *DeleteLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogsResponse) ProtoMessage()               {}
func (*DeleteLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteLogsResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// TimeRangeResponse contains either one aggregated result of a facetted request
// OR a time range cursor.
type TimeRangeResponse struct {
	TimeRangeResult *TimeRangeResult `protobuf:"bytes,1,opt,name=TimeRangeResult" json:"TimeRangeResult,omitempty"`
	TimeRangeCursor *TimeRangeCursor `protobuf:"bytes,2,opt,name=TimeRangeCursor" json:"TimeRangeCursor,omitempty"`
}

func (m *TimeRangeResponse) Reset()                    { *m = TimeRangeResponse{} }
func (m *TimeRangeResponse) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeResponse) ProtoMessage()               {}
func (*TimeRangeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TimeRangeResponse) GetTimeRangeResult() *TimeRangeResult {
	if m != nil {
		return m.TimeRangeResult
	}
	return nil
}

func (m *TimeRangeResponse) GetTimeRangeCursor() *TimeRangeCursor {
	if m != nil {
		return m.TimeRangeCursor
	}
	return nil
}

// TimeRangeResult represents one point of a graph.
type TimeRangeResult struct {
	// a label for this time range
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// begin timestamp
	Start int32 `protobuf:"varint,2,opt,name=Start" json:"Start,omitempty"`
	// end timestamp
	End int32 `protobuf:"varint,3,opt,name=End" json:"End,omitempty"`
	// nb of occurrences found within this range
	Count int32 `protobuf:"varint,4,opt,name=Count" json:"Count,omitempty"`
	// a score between 1 and 100 that gives the relevance of this result:
	// if End > now, we ponderate the returned count with the duration of the last time range
	// for instance for a hour range if now is 6PM, last count will be
	// multiplied by 4/3 and have a relevance of 75.
	// Relevance will be almost always equals to 100
	Relevance int32 `protobuf:"varint,5,opt,name=Relevance" json:"Relevance,omitempty"`
}

func (m *TimeRangeResult) Reset()                    { *m = TimeRangeResult{} }
func (m *TimeRangeResult) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeResult) ProtoMessage()               {}
func (*TimeRangeResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TimeRangeResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeRangeResult) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimeRangeResult) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *TimeRangeResult) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TimeRangeResult) GetRelevance() int32 {
	if m != nil {
		return m.Relevance
	}
	return 0
}

// TimeRangeRequest contains the parameter to configure the query to
// retrieve the number of audit events of this type for a given time range
// defined by last timestamp and a range type.
type TimeRangeRequest struct {
	// Type of event we are auditing
	MsgId string `protobuf:"bytes,1,opt,name=MsgId" json:"MsgId,omitempty"`
	// Known types: H, D, W, M or Y
	TimeRangeType string `protobuf:"bytes,2,opt,name=TimeRangeType" json:"TimeRangeType,omitempty"`
	// Upper bound for our request
	RefTime int32 `protobuf:"varint,3,opt,name=RefTime" json:"RefTime,omitempty"`
}

func (m *TimeRangeRequest) Reset()                    { *m = TimeRangeRequest{} }
func (m *TimeRangeRequest) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeRequest) ProtoMessage()               {}
func (*TimeRangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TimeRangeRequest) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *TimeRangeRequest) GetTimeRangeType() string {
	if m != nil {
		return m.TimeRangeType
	}
	return ""
}

func (m *TimeRangeRequest) GetRefTime() int32 {
	if m != nil {
		return m.RefTime
	}
	return 0
}

// Ease implementation of data navigation for a chart.
type TimeRangeCursor struct {
	Rel     RelType `protobuf:"varint,1,opt,name=Rel,enum=log.RelType" json:"Rel,omitempty"`
	RefTime int32   `protobuf:"varint,2,opt,name=RefTime" json:"RefTime,omitempty"`
	Count   int32   `protobuf:"varint,3,opt,name=Count" json:"Count,omitempty"`
}

func (m *TimeRangeCursor) Reset()                    { *m = TimeRangeCursor{} }
func (m *TimeRangeCursor) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeCursor) ProtoMessage()               {}
func (*TimeRangeCursor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TimeRangeCursor) GetRel() RelType {
	if m != nil {
		return m.Rel
	}
	return RelType_NONE
}

func (m *TimeRangeCursor) GetRefTime() int32 {
	if m != nil {
		return m.RefTime
	}
	return 0
}

func (m *TimeRangeCursor) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*RecorderPutResponse)(nil), "log.RecorderPutResponse")
	proto.RegisterType((*Log)(nil), "log.Log")
	proto.RegisterType((*LogMessage)(nil), "log.LogMessage")
	proto.RegisterType((*ListLogRequest)(nil), "log.ListLogRequest")
	proto.RegisterType((*ListLogResponse)(nil), "log.ListLogResponse")
	proto.RegisterType((*DeleteLogsResponse)(nil), "log.DeleteLogsResponse")
	proto.RegisterType((*TimeRangeResponse)(nil), "log.TimeRangeResponse")
	proto.RegisterType((*TimeRangeResult)(nil), "log.TimeRangeResult")
	proto.RegisterType((*TimeRangeRequest)(nil), "log.TimeRangeRequest")
	proto.RegisterType((*TimeRangeCursor)(nil), "log.TimeRangeCursor")
	proto.RegisterEnum("log.RelType", RelType_name, RelType_value)
	proto.RegisterEnum("log.ListLogRequest_LogFormat", ListLogRequest_LogFormat_name, ListLogRequest_LogFormat_value)
}

func init() { proto.RegisterFile("log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 884 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x35, 0x45, 0xeb, 0x6b, 0x6c, 0xcb, 0xcc, 0xfa, 0xa3, 0x0b, 0xa3, 0x2d, 0x0c, 0x22, 0x28,
	0x8c, 0x1e, 0x9c, 0x40, 0x45, 0x81, 0x36, 0x08, 0x0a, 0xb8, 0xae, 0xd2, 0xa6, 0xa0, 0x65, 0x75,
	0xa5, 0x24, 0xbe, 0x32, 0xd6, 0x84, 0x11, 0x4a, 0x69, 0x55, 0xee, 0xd2, 0x80, 0x7b, 0xeb, 0xad,
	0xf7, 0xfe, 0x90, 0xfe, 0xc2, 0x02, 0xc5, 0xcc, 0x92, 0x94, 0x28, 0xe5, 0xb6, 0xf3, 0x66, 0xde,
	0xcc, 0xce, 0xce, 0x1b, 0x12, 0xba, 0xa9, 0x4e, 0x2e, 0x97, 0x99, 0xb6, 0x5a, 0xf8, 0xa9, 0x4e,
	0xc2, 0x13, 0x38, 0x52, 0x78, 0xaf, 0xb3, 0x29, 0x66, 0xa3, 0xdc, 0x2a, 0x34, 0x4b, 0xbd, 0x30,
	0x18, 0x66, 0xe0, 0x47, 0x3a, 0x11, 0xcf, 0xa0, 0x3d, 0x47, 0x63, 0xe2, 0x04, 0xa5, 0x77, 0xee,
	0x5f, 0xec, 0xf5, 0x4f, 0x2e, 0x89, 0x1f, 0xe9, 0xe4, 0xf2, 0xc6, 0xe1, 0x83, 0x85, 0xcd, 0x1e,
	0x55, 0x19, 0x75, 0xf6, 0x02, 0xf6, 0xd7, 0x1d, 0x22, 0x00, 0xff, 0x77, 0x7c, 0x94, 0xde, 0xb9,
	0x77, 0xd1, 0x55, 0x74, 0x14, 0xc7, 0xd0, 0x7c, 0x88, 0xd3, 0x1c, 0x65, 0x83, 0x31, 0x67, 0xbc,
	0x68, 0x7c, 0xe7, 0x85, 0x7f, 0x37, 0x01, 0x22, 0x9d, 0x14, 0x7c, 0xd1, 0x83, 0xc6, 0xc4, 0x30,
	0xb3, 0xa9, 0x1a, 0x13, 0x43, 0xc4, 0x08, 0x1f, 0x30, 0x2d, 0x89, 0x6c, 0x88, 0x53, 0x68, 0x45,
	0x3a, 0x49, 0x30, 0x93, 0x3e, 0xc3, 0x85, 0x45, 0x85, 0x6f, 0x4c, 0x22, 0x77, 0x5d, 0xe1, 0x1b,
	0x93, 0x10, 0xff, 0xc6, 0x24, 0xaf, 0xa7, 0xb2, 0xe9, 0xf8, 0x6c, 0x88, 0x33, 0xe8, 0xbc, 0x31,
	0x98, 0x0d, 0xe3, 0x39, 0xca, 0x16, 0x3b, 0x2a, 0xbb, 0xf4, 0xbd, 0xc9, 0x67, 0x53, 0xd9, 0x5e,
	0xf9, 0xc8, 0x16, 0x9f, 0x43, 0xf7, 0xe7, 0x4c, 0xe7, 0xcb, 0x51, 0x6c, 0x3f, 0xca, 0x0e, 0x3b,
	0x57, 0x80, 0x90, 0xd0, 0x1e, 0x65, 0xfa, 0xc3, 0x2c, 0x45, 0x19, 0xb0, 0xaf, 0x34, 0x89, 0xa7,
	0x74, 0x8a, 0x94, 0xc3, 0xc8, 0xee, 0xb9, 0x4f, 0xbc, 0x0a, 0x10, 0x4f, 0xe1, 0x40, 0xe1, 0x5c,
	0x5b, 0xbc, 0x9a, 0x4e, 0x33, 0x34, 0x46, 0x02, 0xb3, 0xeb, 0x20, 0xe5, 0xa0, 0x7b, 0x5c, 0x25,
	0xb8, 0xb0, 0x72, 0xcf, 0xd5, 0xae, 0x00, 0x11, 0xc2, 0xfe, 0x2f, 0xd6, 0x2e, 0x47, 0x34, 0xe3,
	0x7b, 0x9d, 0xca, 0x7d, 0x0e, 0xa8, 0x61, 0xd4, 0xd9, 0x50, 0x4f, 0xb9, 0xa8, 0x3c, 0x70, 0x9d,
	0x95, 0x76, 0xe9, 0xe3, 0xc6, 0x7a, 0x2b, 0x1f, 0xf7, 0x75, 0x0a, 0xad, 0x77, 0x86, 0x59, 0x87,
	0xee, 0xb5, 0x9d, 0x45, 0xfd, 0xbe, 0x33, 0xe3, 0x7b, 0xbd, 0x44, 0xf9, 0xc4, 0xf5, 0x5b, 0x98,
	0x94, 0x6d, 0xbc, 0x8c, 0x17, 0xcc, 0x11, 0x2e, 0x5b, 0x69, 0x8b, 0xaf, 0xa0, 0x47, 0xe7, 0x51,
	0x9c, 0xe1, 0xc2, 0x72, 0xc4, 0x11, 0x47, 0x6c, 0xa0, 0xd4, 0x11, 0x21, 0x4a, 0x6b, 0x17, 0x75,
	0xec, 0x3a, 0x5a, 0xc7, 0xe8, 0xe5, 0x6e, 0x97, 0x98, 0xc5, 0x76, 0xa6, 0x5d, 0xb1, 0x13, 0xf7,
	0x72, 0x35, 0x90, 0x2a, 0x56, 0x40, 0x14, 0xbf, 0xc7, 0x54, 0x9e, 0xba, 0x8a, 0x75, 0x34, 0xfc,
	0xd7, 0x83, 0x5e, 0x34, 0x33, 0x36, 0xd2, 0x89, 0xc2, 0x3f, 0x72, 0x34, 0x96, 0xe4, 0xf3, 0x5b,
	0x8e, 0x59, 0xa9, 0x65, 0x67, 0x08, 0x01, 0xbb, 0x23, 0xda, 0x8e, 0x06, 0xcb, 0x94, 0xcf, 0x84,
	0x8d, 0x67, 0x7f, 0x22, 0x0b, 0xb2, 0xa9, 0xf8, 0x2c, 0xbe, 0x85, 0xd6, 0x2b, 0x9d, 0xcd, 0x63,
	0xcb, 0x8a, 0xec, 0xf5, 0xbf, 0x70, 0x7b, 0x54, 0x2b, 0x41, 0x6b, 0xe5, 0x82, 0x54, 0x11, 0x1c,
	0x5e, 0x40, 0xb7, 0x02, 0x45, 0x07, 0x76, 0x7f, 0x1d, 0xdf, 0x0e, 0x83, 0x1d, 0xd1, 0x06, 0xff,
	0x7a, 0xfc, 0x36, 0xf0, 0x08, 0xba, 0x8b, 0xc6, 0x77, 0x41, 0x23, 0xfc, 0x11, 0x0e, 0xab, 0x6c,
	0x6e, 0x87, 0xc5, 0xb3, 0xf5, 0x75, 0xe2, 0x6b, 0xef, 0xf5, 0x0f, 0xcb, 0xfd, 0x2d, 0x60, 0xb5,
	0x16, 0x12, 0x5e, 0x82, 0xf8, 0x09, 0x53, 0xb4, 0x18, 0xe9, 0xc4, 0x54, 0x69, 0x24, 0xb4, 0x1d,
	0x3a, 0xe5, 0x1c, 0xbe, 0x2a, 0xcd, 0xf0, 0x1f, 0x0f, 0x9e, 0x4c, 0x66, 0x73, 0x54, 0xf1, 0x22,
	0xc1, 0x2a, 0xfe, 0x07, 0x38, 0x5c, 0x07, 0xf3, 0xd4, 0x16, 0xb5, 0x8f, 0xb9, 0xf6, 0x86, 0x4f,
	0x6d, 0x06, 0xd7, 0xf8, 0xd7, 0x79, 0x66, 0x74, 0xc6, 0xaf, 0xbb, 0xc5, 0x77, 0x3e, 0xb5, 0x19,
	0x1c, 0xfe, 0xe5, 0x6d, 0x5d, 0x80, 0x46, 0xc2, 0x1b, 0xee, 0x66, 0xc7, 0x67, 0x1a, 0xe8, 0xd8,
	0xc6, 0x99, 0x2d, 0x66, 0xe7, 0x0c, 0xfa, 0x6e, 0x0c, 0x16, 0xd3, 0x62, 0x76, 0x74, 0xa4, 0xb8,
	0x6b, 0x9d, 0x2f, 0xdc, 0xe4, 0x9a, 0xca, 0x19, 0xbc, 0xc7, 0x98, 0xe2, 0x43, 0xbc, 0xb8, 0x47,
	0xfe, 0xa2, 0x34, 0xd5, 0x0a, 0x08, 0x3f, 0x42, 0xb0, 0x76, 0x85, 0x4a, 0x40, 0xee, 0xfb, 0xe3,
	0xad, 0x7f, 0x7f, 0x9e, 0xc2, 0x41, 0x15, 0x39, 0x79, 0x5c, 0x96, 0x9f, 0xc5, 0x3a, 0x48, 0x33,
	0x50, 0xf8, 0x81, 0xb0, 0xe2, 0x66, 0xa5, 0x19, 0xc6, 0x5b, 0xaf, 0x25, 0xbe, 0x04, 0x5f, 0x61,
	0xca, 0x65, 0x7a, 0xfd, 0x7d, 0x7e, 0x34, 0x85, 0x29, 0xe5, 0x51, 0xe4, 0x58, 0x4f, 0xd6, 0xa8,
	0x25, 0x5b, 0xb5, 0xea, 0xaf, 0xb5, 0xfa, 0xf5, 0x4b, 0x8a, 0x67, 0x3e, 0xe9, 0x6d, 0x78, 0x3b,
	0x1c, 0x04, 0x3b, 0xa2, 0x0b, 0xcd, 0x57, 0xaf, 0xd5, 0x78, 0xe2, 0x44, 0x38, 0x52, 0x83, 0xb7,
	0x41, 0x83, 0xdd, 0x83, 0xbb, 0x49, 0xe0, 0xd3, 0x29, 0xba, 0x1a, 0x4f, 0x82, 0xdd, 0xfe, 0x7f,
	0x1e, 0xec, 0xb1, 0x2a, 0xdd, 0x4f, 0x46, 0x3c, 0x87, 0xd6, 0x28, 0x27, 0x9d, 0x8a, 0x4e, 0xa9,
	0xc5, 0x33, 0x59, 0x5c, 0x72, 0xfb, 0x3f, 0xb4, 0x73, 0xe1, 0x89, 0xef, 0xa1, 0x53, 0x48, 0xdb,
	0x88, 0xa3, 0x4f, 0xec, 0xcd, 0xd9, 0x71, 0x1d, 0x2c, 0xa9, 0xcf, 0x3d, 0xf1, 0x12, 0x60, 0xa5,
	0xe8, 0x4f, 0x93, 0x3f, 0x63, 0x70, 0x5b, 0xf7, 0xe1, 0x8e, 0xb8, 0x86, 0xde, 0x55, 0x92, 0x64,
	0x98, 0xc4, 0x16, 0xa7, 0x9c, 0xe1, 0x64, 0x53, 0xc2, 0x2e, 0xc7, 0xe9, 0x96, 0xb2, 0xab, 0x2b,
	0xbc, 0x6f, 0xf1, 0xcf, 0xf6, 0x9b, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x86, 0x1a, 0x1d, 0xf0,
	0x79, 0x07, 0x00, 0x00,
}

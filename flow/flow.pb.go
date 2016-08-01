// Code generated by protoc-gen-go.
// source: flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	flow/flow.proto
	flow/set.proto
	flow/request.proto

It has these top-level messages:
	FlowEndpointStatistics
	FlowEndpointsStatistics
	FlowStatistics
	Flow
	FlowSet
	TermStringFilter
	TermInt64Filter
	NeStringFilter
	NeInt64Filter
	GtInt64Filter
	LtInt64Filter
	GteInt64Filter
	LteInt64Filter
	Filter
	BoolFilter
	FlowSearchQuery
	FlowSearchReply
*/
package flow

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

type FlowEndpointLayer int32

const (
	FlowEndpointLayer_LINK      FlowEndpointLayer = 0
	FlowEndpointLayer_NETWORK   FlowEndpointLayer = 1
	FlowEndpointLayer_TRANSPORT FlowEndpointLayer = 2
)

var FlowEndpointLayer_name = map[int32]string{
	0: "LINK",
	1: "NETWORK",
	2: "TRANSPORT",
}
var FlowEndpointLayer_value = map[string]int32{
	"LINK":      0,
	"NETWORK":   1,
	"TRANSPORT": 2,
}

func (x FlowEndpointLayer) String() string {
	return proto.EnumName(FlowEndpointLayer_name, int32(x))
}
func (FlowEndpointLayer) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FlowEndpointType int32

const (
	FlowEndpointType_ETHERNET FlowEndpointType = 0
	FlowEndpointType_IPV4     FlowEndpointType = 1
	FlowEndpointType_TCPPORT  FlowEndpointType = 2
	FlowEndpointType_UDPPORT  FlowEndpointType = 3
	FlowEndpointType_SCTPPORT FlowEndpointType = 4
	FlowEndpointType_IPV6     FlowEndpointType = 5
)

var FlowEndpointType_name = map[int32]string{
	0: "ETHERNET",
	1: "IPV4",
	2: "TCPPORT",
	3: "UDPPORT",
	4: "SCTPPORT",
	5: "IPV6",
}
var FlowEndpointType_value = map[string]int32{
	"ETHERNET": 0,
	"IPV4":     1,
	"TCPPORT":  2,
	"UDPPORT":  3,
	"SCTPPORT": 4,
	"IPV6":     5,
}

func (x FlowEndpointType) String() string {
	return proto.EnumName(FlowEndpointType_name, int32(x))
}
func (FlowEndpointType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FlowEndpointStatistics struct {
	Value   string `protobuf:"bytes,2,opt,name=Value,json=value" json:"Value,omitempty"`
	Packets uint64 `protobuf:"varint,5,opt,name=Packets,json=packets" json:"Packets,omitempty"`
	Bytes   uint64 `protobuf:"varint,6,opt,name=Bytes,json=bytes" json:"Bytes,omitempty"`
}

func (m *FlowEndpointStatistics) Reset()                    { *m = FlowEndpointStatistics{} }
func (m *FlowEndpointStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowEndpointStatistics) ProtoMessage()               {}
func (*FlowEndpointStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FlowEndpointsStatistics struct {
	Type FlowEndpointType        `protobuf:"varint,1,opt,name=Type,json=type,enum=flow.FlowEndpointType" json:"Type,omitempty"`
	Hash []byte                  `protobuf:"bytes,2,opt,name=Hash,json=hash,proto3" json:"Hash,omitempty"`
	AB   *FlowEndpointStatistics `protobuf:"bytes,3,opt,name=AB,json=aB" json:"AB,omitempty"`
	BA   *FlowEndpointStatistics `protobuf:"bytes,4,opt,name=BA,json=bA" json:"BA,omitempty"`
}

func (m *FlowEndpointsStatistics) Reset()                    { *m = FlowEndpointsStatistics{} }
func (m *FlowEndpointsStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowEndpointsStatistics) ProtoMessage()               {}
func (*FlowEndpointsStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlowEndpointsStatistics) GetAB() *FlowEndpointStatistics {
	if m != nil {
		return m.AB
	}
	return nil
}

func (m *FlowEndpointsStatistics) GetBA() *FlowEndpointStatistics {
	if m != nil {
		return m.BA
	}
	return nil
}

type FlowStatistics struct {
	Start     int64                      `protobuf:"varint,1,opt,name=Start,json=start" json:"Start,omitempty"`
	Last      int64                      `protobuf:"varint,2,opt,name=Last,json=last" json:"Last,omitempty"`
	Endpoints []*FlowEndpointsStatistics `protobuf:"bytes,3,rep,name=Endpoints,json=endpoints" json:"Endpoints,omitempty"`
}

func (m *FlowStatistics) Reset()                    { *m = FlowStatistics{} }
func (m *FlowStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowStatistics) ProtoMessage()               {}
func (*FlowStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FlowStatistics) GetEndpoints() []*FlowEndpointsStatistics {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Flow struct {
	// Flow Universally Unique IDentifier
	//
	// flow.UUID is unique in the universe, as it should be used as a key of an
	// hashtable. By design 2 different flows, their UUID are always different.
	// flow.UUID can be used as Database Index.
	UUID       string `protobuf:"bytes,1,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	LayersPath string `protobuf:"bytes,2,opt,name=LayersPath,json=layersPath" json:"LayersPath,omitempty"`
	// Data Flow info
	Statistics *FlowStatistics `protobuf:"bytes,3,opt,name=Statistics,json=statistics" json:"Statistics,omitempty"`
	// Flow Tracking IDentifier, from 1st packet bytes
	//
	// flow.TrackingID could be used to identify an unique flow whatever it has
	// been captured on the infrastructure. flow.TrackingID is calculated from
	// the bytes of the first packet of his session.
	// flow.TrackingID can be used as a Tag.
	TrackingID string `protobuf:"bytes,5,opt,name=TrackingID,json=trackingID" json:"TrackingID,omitempty"`
	// Topology info
	ProbeNodeUUID string `protobuf:"bytes,11,opt,name=ProbeNodeUUID,json=probeNodeUUID" json:"ProbeNodeUUID,omitempty"`
	IfSrcNodeUUID string `protobuf:"bytes,14,opt,name=IfSrcNodeUUID,json=ifSrcNodeUUID" json:"IfSrcNodeUUID,omitempty"`
	IfDstNodeUUID string `protobuf:"bytes,19,opt,name=IfDstNodeUUID,json=ifDstNodeUUID" json:"IfDstNodeUUID,omitempty"`
}

func (m *Flow) Reset()                    { *m = Flow{} }
func (m *Flow) String() string            { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()               {}
func (*Flow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Flow) GetStatistics() *FlowStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FlowEndpointStatistics)(nil), "flow.FlowEndpointStatistics")
	proto.RegisterType((*FlowEndpointsStatistics)(nil), "flow.FlowEndpointsStatistics")
	proto.RegisterType((*FlowStatistics)(nil), "flow.FlowStatistics")
	proto.RegisterType((*Flow)(nil), "flow.Flow")
	proto.RegisterEnum("flow.FlowEndpointLayer", FlowEndpointLayer_name, FlowEndpointLayer_value)
	proto.RegisterEnum("flow.FlowEndpointType", FlowEndpointType_name, FlowEndpointType_value)
}

func init() { proto.RegisterFile("flow/flow.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xa6, 0x89, 0xb3, 0x2e, 0xaf, 0x6b, 0x09, 0x66, 0x1a, 0x39, 0x00, 0x42, 0x15, 0x07, 0x54,
	0xa1, 0x21, 0x8d, 0x09, 0x09, 0x71, 0x6a, 0xd7, 0xa2, 0x55, 0x9b, 0xba, 0xc8, 0x75, 0xc7, 0x05,
	0x21, 0xb9, 0x9d, 0x47, 0x23, 0xaa, 0x26, 0x8a, 0x3d, 0xa6, 0x5e, 0xf9, 0x4f, 0xfc, 0x3f, 0x9e,
	0x9d, 0x2e, 0x31, 0xf4, 0xc2, 0x25, 0x7d, 0xdf, 0x97, 0xef, 0x7b, 0xdf, 0xf3, 0x73, 0x03, 0x8f,
	0x6f, 0x57, 0xd9, 0xfd, 0x3b, 0xf3, 0x38, 0xce, 0x8b, 0x4c, 0x67, 0x94, 0x98, 0xba, 0xfb, 0x0d,
	0x8e, 0x3e, 0xe3, 0xef, 0x68, 0x7d, 0x93, 0x67, 0xe9, 0x5a, 0x4f, 0xb5, 0xd0, 0xa9, 0xd2, 0xe9,
	0x42, 0xd1, 0x43, 0x08, 0xae, 0xc5, 0xea, 0x4e, 0xc6, 0xde, 0xab, 0xc6, 0x9b, 0x90, 0x05, 0x3f,
	0x0d, 0xa0, 0x31, 0x34, 0x13, 0xb1, 0xf8, 0x21, 0xb5, 0x8a, 0x03, 0xe4, 0x09, 0x6b, 0xe6, 0x25,
	0x34, 0xfa, 0xc1, 0x46, 0x4b, 0x15, 0xef, 0x59, 0x3e, 0x98, 0x1b, 0xd0, 0xfd, 0xdd, 0x80, 0x67,
	0x6e, 0x80, 0x72, 0x12, 0x7a, 0x40, 0xf8, 0x26, 0x97, 0x71, 0x03, 0x0d, 0x9d, 0x93, 0xa3, 0x63,
	0x3b, 0x9c, 0x2b, 0x36, 0x6f, 0x19, 0xd1, 0xf8, 0xa4, 0x14, 0xc8, 0xb9, 0x50, 0x4b, 0x3b, 0xcc,
	0x01, 0x23, 0x4b, 0xac, 0xe9, 0x5b, 0xf0, 0xfa, 0x83, 0xd8, 0x47, 0xa6, 0x75, 0xf2, 0x7c, 0xd7,
	0x5d, 0x27, 0x31, 0x4f, 0x0c, 0x8c, 0x7a, 0xd0, 0x8f, 0xc9, 0xff, 0xa8, 0xe7, 0xfd, 0xee, 0x3d,
	0x74, 0xcc, 0xdb, 0xbf, 0xf7, 0x81, 0xa8, 0xd0, 0x76, 0x5c, 0x9f, 0x05, 0xca, 0x00, 0x33, 0xd7,
	0xa5, 0x50, 0xda, 0xce, 0xe5, 0x33, 0xb2, 0xc2, 0x9a, 0x7e, 0x82, 0xb0, 0x3a, 0x2e, 0x8e, 0xe7,
	0x63, 0xe0, 0x8b, 0xdd, 0x40, 0x67, 0x13, 0x2c, 0x94, 0x0f, 0x64, 0xf7, 0x97, 0x07, 0xc4, 0xc8,
	0x4c, 0xe7, 0xd9, 0x6c, 0x3c, 0xb4, 0x71, 0x21, 0x23, 0x77, 0x58, 0xd3, 0x97, 0x00, 0x97, 0x62,
	0x23, 0x0b, 0x95, 0x08, 0xbd, 0xdc, 0x5e, 0x0c, 0xac, 0x2a, 0x86, 0x9e, 0x02, 0xd4, 0x5d, 0xb7,
	0x9b, 0x39, 0xac, 0xa3, 0x9d, 0x44, 0x50, 0xf5, 0xc9, 0xb0, 0x2b, 0x2f, 0xf0, 0x16, 0xd3, 0xf5,
	0x77, 0xcc, 0x0b, 0xca, 0xae, 0xba, 0x62, 0xe8, 0x6b, 0x68, 0x27, 0x45, 0x36, 0x97, 0x93, 0xec,
	0x46, 0xda, 0x91, 0x5a, 0x56, 0xd2, 0xce, 0x5d, 0xd2, 0xa8, 0xc6, 0xb7, 0xd3, 0x62, 0x51, 0xa9,
	0x3a, 0xa5, 0x2a, 0x75, 0xc9, 0x52, 0x35, 0x54, 0xba, 0x52, 0x3d, 0x7d, 0x50, 0x39, 0x64, 0xef,
	0x23, 0x3c, 0x71, 0x57, 0x65, 0xcf, 0x4c, 0xf7, 0x71, 0xd5, 0xe3, 0xc9, 0x45, 0xf4, 0x88, 0xb6,
	0xa0, 0x39, 0x19, 0xf1, 0x2f, 0x57, 0xec, 0x22, 0x6a, 0xd0, 0x36, 0x84, 0x9c, 0xf5, 0x27, 0xd3,
	0xe4, 0x8a, 0xf1, 0xc8, 0xeb, 0x7d, 0x85, 0xe8, 0xdf, 0xbf, 0x10, 0x3d, 0x80, 0xfd, 0x11, 0x3f,
	0x1f, 0x31, 0x34, 0xa1, 0x1b, 0xfb, 0x8c, 0x93, 0xeb, 0x53, 0xb4, 0x62, 0x1f, 0x7e, 0x96, 0x94,
	0x46, 0x03, 0x66, 0xc3, 0x12, 0xf8, 0xc6, 0x31, 0x3d, 0xe3, 0x25, 0x22, 0x5b, 0xc7, 0x87, 0x28,
	0x98, 0xef, 0xd9, 0x6f, 0xe7, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x93, 0x46, 0x22,
	0x4e, 0x03, 0x00, 0x00,
}

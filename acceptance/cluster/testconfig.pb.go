// Code generated by protoc-gen-gogo.
// source: cockroach/acceptance/cluster/testconfig.proto
// DO NOT EDIT!

/*
	Package cluster is a generated protocol buffer package.

	It is generated from these files:
		cockroach/acceptance/cluster/testconfig.proto

	It has these top-level messages:
		StoreConfig
		NodeConfig
		TestConfig
*/
package cluster

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import time "time"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// StoreConfig holds the configuration of a collection of similar stores.
type StoreConfig struct {
	Count     int32 `protobuf:"varint,1,opt,name=count" json:"count"`
	MaxRanges int32 `protobuf:"varint,2,opt,name=max_ranges" json:"max_ranges"`
}

func (m *StoreConfig) Reset()                    { *m = StoreConfig{} }
func (m *StoreConfig) String() string            { return proto.CompactTextString(m) }
func (*StoreConfig) ProtoMessage()               {}
func (*StoreConfig) Descriptor() ([]byte, []int) { return fileDescriptorTestconfig, []int{0} }

// NodeConfig holds the configuration of a collection of similar nodes.
type NodeConfig struct {
	Count  int32         `protobuf:"varint,1,opt,name=count" json:"count"`
	Stores []StoreConfig `protobuf:"bytes,2,rep,name=stores" json:"stores"`
}

func (m *NodeConfig) Reset()                    { *m = NodeConfig{} }
func (m *NodeConfig) String() string            { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()               {}
func (*NodeConfig) Descriptor() ([]byte, []int) { return fileDescriptorTestconfig, []int{1} }

type TestConfig struct {
	Name  string       `protobuf:"bytes,1,opt,name=name" json:"name"`
	Nodes []NodeConfig `protobuf:"bytes,2,rep,name=nodes" json:"nodes"`
	// Duration is the total time that the test should run for. Important for
	// tests such as TestPut that will run indefinitely.
	Duration time.Duration `protobuf:"varint,3,opt,name=duration,casttype=time.Duration" json:"duration"`
	// Stall is the duration after which if no forward progress is made in a test,
	// consider the test stalled.
	Stall time.Duration `protobuf:"varint,4,opt,name=stall,casttype=time.Duration" json:"stall"`
}

func (m *TestConfig) Reset()                    { *m = TestConfig{} }
func (m *TestConfig) String() string            { return proto.CompactTextString(m) }
func (*TestConfig) ProtoMessage()               {}
func (*TestConfig) Descriptor() ([]byte, []int) { return fileDescriptorTestconfig, []int{2} }

func init() {
	proto.RegisterType((*StoreConfig)(nil), "cockroach.acceptance.cluster.StoreConfig")
	proto.RegisterType((*NodeConfig)(nil), "cockroach.acceptance.cluster.NodeConfig")
	proto.RegisterType((*TestConfig)(nil), "cockroach.acceptance.cluster.TestConfig")
}
func (m *StoreConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintTestconfig(data, i, uint64(m.Count))
	data[i] = 0x10
	i++
	i = encodeVarintTestconfig(data, i, uint64(m.MaxRanges))
	return i, nil
}

func (m *NodeConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NodeConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintTestconfig(data, i, uint64(m.Count))
	if len(m.Stores) > 0 {
		for _, msg := range m.Stores {
			data[i] = 0x12
			i++
			i = encodeVarintTestconfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TestConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TestConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintTestconfig(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			data[i] = 0x12
			i++
			i = encodeVarintTestconfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x18
	i++
	i = encodeVarintTestconfig(data, i, uint64(m.Duration))
	data[i] = 0x20
	i++
	i = encodeVarintTestconfig(data, i, uint64(m.Stall))
	return i, nil
}

func encodeFixed64Testconfig(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Testconfig(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTestconfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StoreConfig) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTestconfig(uint64(m.Count))
	n += 1 + sovTestconfig(uint64(m.MaxRanges))
	return n
}

func (m *NodeConfig) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTestconfig(uint64(m.Count))
	if len(m.Stores) > 0 {
		for _, e := range m.Stores {
			l = e.Size()
			n += 1 + l + sovTestconfig(uint64(l))
		}
	}
	return n
}

func (m *TestConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovTestconfig(uint64(l))
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovTestconfig(uint64(l))
		}
	}
	n += 1 + sovTestconfig(uint64(m.Duration))
	n += 1 + sovTestconfig(uint64(m.Stall))
	return n
}

func sovTestconfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTestconfig(x uint64) (n int) {
	return sovTestconfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRanges", wireType)
			}
			m.MaxRanges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MaxRanges |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestconfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestconfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stores = append(m.Stores, StoreConfig{})
			if err := m.Stores[len(m.Stores)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestconfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestconfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, NodeConfig{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Duration |= (time.Duration(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stall", wireType)
			}
			m.Stall = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Stall |= (time.Duration(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestconfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestconfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTestconfig(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTestconfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTestconfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTestconfig(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTestconfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestconfig   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorTestconfig = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xce, 0x4f, 0xce,
	0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x4f, 0x4c, 0x4e, 0x4e, 0x2d, 0x28, 0x49, 0xcc, 0x4b, 0x4e,
	0xd5, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x2f, 0x49, 0x2d, 0x2e, 0x49, 0xce, 0xcf,
	0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0x2b, 0xd7, 0x43, 0x28,
	0xd7, 0x83, 0x2a, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd4, 0x07, 0xb1, 0x20, 0x7a,
	0x94, 0x34, 0xb9, 0xb8, 0x83, 0x4b, 0xf2, 0x8b, 0x52, 0x9d, 0xc1, 0x06, 0x09, 0x49, 0x71, 0xb1,
	0x26, 0xe7, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x3a, 0xb1, 0x9c, 0xb8, 0x27,
	0xcf, 0x10, 0x04, 0x11, 0x52, 0x2a, 0xe4, 0xe2, 0xf2, 0xcb, 0x4f, 0x21, 0x42, 0xa5, 0x90, 0x3b,
	0x17, 0x5b, 0x31, 0xc8, 0xd0, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x4d, 0x3d, 0x7c,
	0x2e, 0xd3, 0x43, 0x72, 0x00, 0xd4, 0x1c, 0xa8, 0x76, 0xa5, 0x33, 0x8c, 0x5c, 0x5c, 0x21, 0x40,
	0x6f, 0x42, 0xed, 0x94, 0xe0, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x05, 0x5b, 0xc9, 0x09, 0x55, 0x0a,
	0x16, 0x11, 0x72, 0xe1, 0x62, 0xcd, 0x03, 0xba, 0x0d, 0x66, 0xa1, 0x06, 0x7e, 0x0b, 0x11, 0xde,
	0x80, 0xb9, 0x1b, 0xac, 0x59, 0xc8, 0x90, 0x8b, 0x23, 0xa5, 0xb4, 0x28, 0xb1, 0x24, 0x33, 0x3f,
	0x4f, 0x82, 0x19, 0x68, 0x07, 0xb3, 0x93, 0x28, 0x48, 0xfa, 0xd7, 0x3d, 0x79, 0xde, 0x92, 0xcc,
	0xdc, 0x54, 0x3d, 0x17, 0xa8, 0x64, 0x10, 0x5c, 0x99, 0x90, 0x36, 0x17, 0x6b, 0x71, 0x49, 0x62,
	0x4e, 0x8e, 0x04, 0x0b, 0x3e, 0xf5, 0x10, 0x35, 0x4e, 0x8a, 0x27, 0x1e, 0xca, 0x31, 0x9c, 0x78,
	0x24, 0xc7, 0x78, 0x01, 0x88, 0x6f, 0x00, 0xf1, 0x03, 0x20, 0x9e, 0xf0, 0x58, 0x8e, 0x21, 0x8a,
	0x1d, 0xea, 0xb4, 0x08, 0x06, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x0d, 0x15, 0x65, 0xf5,
	0x01, 0x00, 0x00,
}

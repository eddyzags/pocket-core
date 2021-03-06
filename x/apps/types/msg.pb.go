// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/apps/msg.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_pokt_network_pocket_core_types "github.com/pokt-network/pocket-core/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgProtoStake struct {
	PubKey []byte                                           `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pubkey" yaml:"pubkey"`
	Chains []string                                         `protobuf:"bytes,2,rep,name=chains,proto3" json:"chains" yaml:"chains"`
	Value  github_com_pokt_network_pocket_core_types.BigInt `protobuf:"bytes,3,opt,name=value,proto3,customtype=github.com/pokt-network/pocket-core/types.BigInt" json:"value" yaml:"value"`
}

func (m *MsgProtoStake) Reset()         { *m = MsgProtoStake{} }
func (m *MsgProtoStake) String() string { return proto.CompactTextString(m) }
func (*MsgProtoStake) ProtoMessage()    {}
func (*MsgProtoStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd58e5eb64f87460, []int{0}
}
func (m *MsgProtoStake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProtoStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProtoStake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProtoStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProtoStake.Merge(m, src)
}
func (m *MsgProtoStake) XXX_Size() int {
	return m.Size()
}
func (m *MsgProtoStake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProtoStake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProtoStake proto.InternalMessageInfo

func (*MsgProtoStake) XXX_MessageName() string {
	return "x.apps.MsgProtoStake"
}

type MsgBeginUnstake struct {
	Address github_com_pokt_network_pocket_core_types.Address `protobuf:"bytes,1,opt,name=Address,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"application_address" yaml:"application_address"`
}

func (m *MsgBeginUnstake) Reset()         { *m = MsgBeginUnstake{} }
func (m *MsgBeginUnstake) String() string { return proto.CompactTextString(m) }
func (*MsgBeginUnstake) ProtoMessage()    {}
func (*MsgBeginUnstake) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd58e5eb64f87460, []int{1}
}
func (m *MsgBeginUnstake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBeginUnstake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBeginUnstake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBeginUnstake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBeginUnstake.Merge(m, src)
}
func (m *MsgBeginUnstake) XXX_Size() int {
	return m.Size()
}
func (m *MsgBeginUnstake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBeginUnstake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBeginUnstake proto.InternalMessageInfo

func (*MsgBeginUnstake) XXX_MessageName() string {
	return "x.apps.MsgBeginUnstake"
}

type MsgUnjail struct {
	AppAddr github_com_pokt_network_pocket_core_types.Address `protobuf:"bytes,1,opt,name=AppAddr,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address" yaml:"address"`
}

func (m *MsgUnjail) Reset()         { *m = MsgUnjail{} }
func (m *MsgUnjail) String() string { return proto.CompactTextString(m) }
func (*MsgUnjail) ProtoMessage()    {}
func (*MsgUnjail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd58e5eb64f87460, []int{2}
}
func (m *MsgUnjail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnjail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnjail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnjail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjail.Merge(m, src)
}
func (m *MsgUnjail) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnjail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjail proto.InternalMessageInfo

func (*MsgUnjail) XXX_MessageName() string {
	return "x.apps.MsgUnjail"
}
func init() {
	proto.RegisterType((*MsgProtoStake)(nil), "x.apps.MsgProtoStake")
	proto.RegisterType((*MsgBeginUnstake)(nil), "x.apps.MsgBeginUnstake")
	proto.RegisterType((*MsgUnjail)(nil), "x.apps.MsgUnjail")
}

func init() { proto.RegisterFile("x/apps/msg.proto", fileDescriptor_fd58e5eb64f87460) }

var fileDescriptor_fd58e5eb64f87460 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x8d, 0x39, 0x91, 0xa8, 0xd6, 0x1d, 0x9c, 0x02, 0x43, 0x75, 0x48, 0x71, 0x95, 0xa9, 0xcb,
	0x25, 0xa0, 0x63, 0xba, 0xed, 0xb2, 0x01, 0xaa, 0x04, 0x41, 0xb7, 0xb0, 0x9c, 0x9c, 0x9c, 0xe5,
	0xa6, 0x49, 0x63, 0x2b, 0x76, 0xa0, 0xd9, 0x19, 0x2a, 0xb1, 0x30, 0x32, 0x56, 0x8c, 0xfc, 0x92,
	0x8e, 0x1d, 0x11, 0x83, 0x85, 0xda, 0x05, 0x65, 0xac, 0xc4, 0xc2, 0x84, 0x12, 0x07, 0x75, 0xa0,
	0x43, 0xa5, 0xdb, 0xfc, 0xde, 0xf3, 0xe7, 0xf7, 0x9e, 0xfc, 0xc1, 0xd3, 0x99, 0x8f, 0x39, 0x17,
	0xfe, 0x54, 0x50, 0x8f, 0x17, 0x4c, 0x32, 0xdb, 0x9c, 0x79, 0x0d, 0x73, 0xf6, 0x98, 0x32, 0xca,
	0x5a, 0xca, 0x6f, 0x4e, 0x5a, 0x75, 0x7f, 0x03, 0x78, 0x32, 0x12, 0xf4, 0x75, 0x03, 0xde, 0x4a,
	0x9c, 0x12, 0xfb, 0x39, 0xb4, 0x78, 0x19, 0xdd, 0xa4, 0xa4, 0xea, 0x83, 0x01, 0x18, 0x1e, 0x07,
	0x4f, 0x6a, 0x85, 0x4c, 0x5e, 0x46, 0x29, 0xa9, 0xb6, 0x0a, 0x9d, 0x54, 0x78, 0x9a, 0x5d, 0xba,
	0x1a, 0xbb, 0x61, 0x23, 0xbc, 0x22, 0x95, 0x7d, 0x01, 0xcd, 0x78, 0x8c, 0x93, 0x5c, 0xf4, 0xef,
	0x0d, 0x8e, 0x86, 0x3d, 0x3d, 0xa4, 0x99, 0xdd, 0x90, 0xc6, 0x6e, 0xd8, 0x09, 0x36, 0x85, 0xf7,
	0xdf, 0xe3, 0xac, 0x24, 0xfd, 0xa3, 0x01, 0x18, 0xf6, 0x82, 0x37, 0x4b, 0x85, 0x8c, 0x1f, 0x0a,
	0x3d, 0xa5, 0x89, 0x1c, 0x97, 0x91, 0x17, 0xb3, 0xa9, 0xcf, 0x59, 0x2a, 0xcf, 0x73, 0x22, 0x3f,
	0xb0, 0x22, 0xf5, 0x39, 0x8b, 0x53, 0x22, 0xcf, 0x63, 0x56, 0x10, 0x5f, 0x56, 0x9c, 0x08, 0x2f,
	0x48, 0xe8, 0x8b, 0x5c, 0xd6, 0x0a, 0xe9, 0x87, 0xb6, 0x0a, 0x1d, 0x6b, 0xab, 0x16, 0xba, 0xa1,
	0xa6, 0x2f, 0x4f, 0xe7, 0x0b, 0x64, 0x7c, 0x59, 0x20, 0xf0, 0x6b, 0x81, 0xc0, 0xfc, 0x2b, 0x02,
	0xee, 0x37, 0x00, 0x1f, 0x8e, 0x04, 0x0d, 0x08, 0x4d, 0xf2, 0xeb, 0x5c, 0xb4, 0xcd, 0x3f, 0x02,
	0x68, 0x5d, 0xdd, 0xde, 0x16, 0x44, 0x88, 0xae, 0xfa, 0xa4, 0x56, 0xe8, 0x11, 0xe6, 0x3c, 0x4b,
	0x62, 0x2c, 0x13, 0x96, 0xdf, 0x60, 0x2d, 0x6f, 0x15, 0x3a, 0xd3, 0x3e, 0x7b, 0x44, 0xf7, 0x8f,
	0x42, 0xcf, 0x0e, 0xaf, 0xd0, 0x39, 0x86, 0xff, 0xac, 0xf7, 0x84, 0xfd, 0x04, 0x60, 0x6f, 0x24,
	0xe8, 0x75, 0x3e, 0xc1, 0x49, 0x66, 0x67, 0xd0, 0xba, 0xe2, 0xbc, 0xb9, 0xdd, 0xa5, 0x0c, 0x6b,
	0x85, 0xac, 0x5d, 0xb2, 0x07, 0x5d, 0xb2, 0x3b, 0xa6, 0xd1, 0x16, 0xff, 0xa7, 0x09, 0x5e, 0x2e,
	0xd7, 0x0e, 0x58, 0xad, 0x1d, 0xf0, 0x73, 0xed, 0x80, 0xcf, 0x1b, 0xc7, 0x58, 0x6d, 0x1c, 0xe3,
	0xfb, 0xc6, 0x31, 0xde, 0x1d, 0xf4, 0x71, 0xdd, 0x82, 0xb6, 0x76, 0x91, 0xd9, 0x6e, 0xe1, 0xc5,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x7e, 0x3d, 0x78, 0xb7, 0x02, 0x00, 0x00,
}

func (this *MsgProtoStake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgProtoStake)
	if !ok {
		that2, ok := that.(MsgProtoStake)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.PubKey, that1.PubKey) {
		return false
	}
	if len(this.Chains) != len(that1.Chains) {
		return false
	}
	for i := range this.Chains {
		if this.Chains[i] != that1.Chains[i] {
			return false
		}
	}
	if !this.Value.Equal(that1.Value) {
		return false
	}
	return true
}
func (this *MsgBeginUnstake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgBeginUnstake)
	if !ok {
		that2, ok := that.(MsgBeginUnstake)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	return true
}
func (this *MsgUnjail) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnjail)
	if !ok {
		that2, ok := that.(MsgUnjail)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.AppAddr, that1.AppAddr) {
		return false
	}
	return true
}
func (m *MsgProtoStake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProtoStake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProtoStake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintMsg(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBeginUnstake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBeginUnstake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBeginUnstake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnjail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnjail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnjail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppAddr) > 0 {
		i -= len(m.AppAddr)
		copy(dAtA[i:], m.AppAddr)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.AppAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgProtoStake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	l = m.Value.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgBeginUnstake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgUnjail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppAddr)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProtoStake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProtoStake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProtoStake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgBeginUnstake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgBeginUnstake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBeginUnstake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUnjail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnjail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnjail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAddr = append(m.AppAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.AppAddr == nil {
				m.AppAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)

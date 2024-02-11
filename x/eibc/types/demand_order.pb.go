// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/eibc/demand_order.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/dymensionxyz/dymension/v3/x/common/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type DemandOrder struct {
	// id is a hash of the form generated by GetRollappPacketKey,
	// e.g status/rollappid/packetProofHeight/packetDestinationChannel-PacketSequence which gurantees uniqueness
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// tracking_packet_key is the key of the packet that is being tracked.
	// This key can change depends on the packet status.
	TrackingPacketKey    string                                   `protobuf:"bytes,2,opt,name=tracking_packet_key,json=trackingPacketKey,proto3" json:"tracking_packet_key,omitempty"`
	Price                github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	Fee                  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee"`
	Recipient            string                                   `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	IsFullfilled         bool                                     `protobuf:"varint,6,opt,name=is_fullfilled,json=isFullfilled,proto3" json:"is_fullfilled,omitempty"`
	TrackingPacketStatus types1.Status                            `protobuf:"varint,8,opt,name=tracking_packet_status,json=trackingPacketStatus,proto3,enum=dymensionxyz.dymension.common.Status" json:"tracking_packet_status,omitempty"`
}

func (m *DemandOrder) Reset()         { *m = DemandOrder{} }
func (m *DemandOrder) String() string { return proto.CompactTextString(m) }
func (*DemandOrder) ProtoMessage()    {}
func (*DemandOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_3808f42eed32f331, []int{0}
}
func (m *DemandOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemandOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemandOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemandOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemandOrder.Merge(m, src)
}
func (m *DemandOrder) XXX_Size() int {
	return m.Size()
}
func (m *DemandOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_DemandOrder.DiscardUnknown(m)
}

var xxx_messageInfo_DemandOrder proto.InternalMessageInfo

func (m *DemandOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DemandOrder) GetTrackingPacketKey() string {
	if m != nil {
		return m.TrackingPacketKey
	}
	return ""
}

func (m *DemandOrder) GetPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *DemandOrder) GetFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *DemandOrder) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *DemandOrder) GetIsFullfilled() bool {
	if m != nil {
		return m.IsFullfilled
	}
	return false
}

func (m *DemandOrder) GetTrackingPacketStatus() types1.Status {
	if m != nil {
		return m.TrackingPacketStatus
	}
	return types1.Status_PENDING
}

func init() {
	proto.RegisterType((*DemandOrder)(nil), "dymensionxyz.dymension.eibc.DemandOrder")
}

func init() { proto.RegisterFile("dymension/eibc/demand_order.proto", fileDescriptor_3808f42eed32f331) }

var fileDescriptor_3808f42eed32f331 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x4e, 0x1a, 0x76, 0xb5, 0xeb, 0x85, 0x95, 0x30, 0x2b, 0x14, 0x16, 0xc8, 0x06, 0x10, 0x52,
	0x2e, 0xd8, 0x74, 0xfb, 0x06, 0x05, 0x71, 0xe9, 0x01, 0x14, 0x6e, 0x20, 0x14, 0x25, 0xf6, 0x34,
	0x58, 0xf9, 0x71, 0x14, 0xbb, 0x55, 0xc3, 0x53, 0xf0, 0x1c, 0x5c, 0x78, 0x8d, 0x1e, 0x7b, 0xe4,
	0x04, 0xa8, 0x7d, 0x11, 0x14, 0x27, 0xa4, 0x05, 0x89, 0x1b, 0x27, 0x7b, 0xe6, 0xfb, 0x66, 0xbe,
	0xf9, 0x43, 0x8f, 0x78, 0x53, 0x40, 0xa9, 0x84, 0x2c, 0x29, 0x88, 0x84, 0x51, 0x0e, 0x45, 0x5c,
	0xf2, 0x48, 0xd6, 0x1c, 0x6a, 0x52, 0xd5, 0x52, 0x4b, 0x7c, 0x7f, 0xa0, 0xac, 0x9a, 0x4f, 0x64,
	0x30, 0x48, 0xcb, 0xbf, 0xbc, 0x48, 0x65, 0x2a, 0x0d, 0x8f, 0xb6, 0xbf, 0x2e, 0xe4, 0xf2, 0xe1,
	0x3e, 0x2b, 0x93, 0x45, 0x31, 0x3c, 0x3d, 0xec, 0x31, 0xa9, 0x0a, 0xa9, 0x68, 0x12, 0x2b, 0xa0,
	0xcb, 0x71, 0x02, 0x3a, 0x1e, 0x53, 0x26, 0x45, 0x8f, 0x3f, 0xfe, 0xea, 0xa0, 0xb3, 0x97, 0xa6,
	0x90, 0xd7, 0x6d, 0x1d, 0xf8, 0x1c, 0x8d, 0x04, 0x77, 0x6d, 0xdf, 0x0e, 0x4e, 0xc3, 0x91, 0xe0,
	0x98, 0xa0, 0x3b, 0xba, 0x8e, 0x59, 0x26, 0xca, 0x34, 0xaa, 0x62, 0x96, 0x81, 0x8e, 0x32, 0x68,
	0xdc, 0x91, 0x21, 0xdc, 0xfe, 0x0d, 0xbd, 0x31, 0xc8, 0x0c, 0x1a, 0x1c, 0xa3, 0xa3, 0xaa, 0x16,
	0x0c, 0x5c, 0xc7, 0x77, 0x82, 0xb3, 0xeb, 0x7b, 0xa4, 0xd3, 0x27, 0xad, 0x3e, 0xe9, 0xf5, 0xc9,
	0x0b, 0x29, 0xca, 0xe9, 0xf3, 0xf5, 0xf7, 0x2b, 0xeb, 0xcb, 0x8f, 0xab, 0x20, 0x15, 0xfa, 0xe3,
	0x22, 0x21, 0x4c, 0x16, 0xb4, 0x2f, 0xb6, 0x7b, 0x9e, 0x29, 0x9e, 0x51, 0xdd, 0x54, 0xa0, 0x4c,
	0x80, 0x0a, 0xbb, 0xcc, 0xf8, 0x03, 0x72, 0xe6, 0x00, 0xee, 0x8d, 0xff, 0x2f, 0xd0, 0xe6, 0xc5,
	0x0f, 0xd0, 0x69, 0x0d, 0x4c, 0x54, 0x02, 0x4a, 0xed, 0x1e, 0x99, 0x3e, 0xf7, 0x0e, 0xfc, 0x04,
	0xdd, 0x12, 0x2a, 0x9a, 0x2f, 0xf2, 0x7c, 0x2e, 0xf2, 0x1c, 0xb8, 0x7b, 0xec, 0xdb, 0xc1, 0x49,
	0x78, 0x53, 0xa8, 0x57, 0x83, 0x0f, 0xbf, 0x47, 0x77, 0xff, 0x1e, 0x9a, 0xd2, 0xb1, 0x5e, 0x28,
	0xf7, 0xc4, 0xb7, 0x83, 0xf3, 0xeb, 0xa7, 0xe4, 0x1f, 0x7b, 0xee, 0x57, 0xf7, 0xd6, 0x90, 0xc3,
	0x8b, 0x3f, 0xc7, 0xdb, 0x79, 0xa7, 0xb3, 0xf5, 0xd6, 0xb3, 0x37, 0x5b, 0xcf, 0xfe, 0xb9, 0xf5,
	0xec, 0xcf, 0x3b, 0xcf, 0xda, 0xec, 0x3c, 0xeb, 0xdb, 0xce, 0xb3, 0xde, 0x8d, 0x0f, 0x1a, 0x3d,
	0x14, 0xd8, 0x1b, 0x74, 0x39, 0xa1, 0xab, 0xee, 0xfa, 0x4c, 0xdf, 0xc9, 0xb1, 0xb9, 0x82, 0xc9,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x27, 0xe9, 0x94, 0x9c, 0x02, 0x00, 0x00,
}

func (m *DemandOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemandOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemandOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TrackingPacketStatus != 0 {
		i = encodeVarintDemandOrder(dAtA, i, uint64(m.TrackingPacketStatus))
		i--
		dAtA[i] = 0x40
	}
	if m.IsFullfilled {
		i--
		if m.IsFullfilled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Fee) > 0 {
		for iNdEx := len(m.Fee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDemandOrder(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDemandOrder(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TrackingPacketKey) > 0 {
		i -= len(m.TrackingPacketKey)
		copy(dAtA[i:], m.TrackingPacketKey)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.TrackingPacketKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDemandOrder(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDemandOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovDemandOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DemandOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	l = len(m.TrackingPacketKey)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovDemandOrder(uint64(l))
		}
	}
	if len(m.Fee) > 0 {
		for _, e := range m.Fee {
			l = e.Size()
			n += 1 + l + sovDemandOrder(uint64(l))
		}
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovDemandOrder(uint64(l))
	}
	if m.IsFullfilled {
		n += 2
	}
	if m.TrackingPacketStatus != 0 {
		n += 1 + sovDemandOrder(uint64(m.TrackingPacketStatus))
	}
	return n
}

func sovDemandOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDemandOrder(x uint64) (n int) {
	return sovDemandOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DemandOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemandOrder
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
			return fmt.Errorf("proto: DemandOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemandOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackingPacketKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackingPacketKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee, types.Coin{})
			if err := m.Fee[len(m.Fee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
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
				return ErrInvalidLengthDemandOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemandOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFullfilled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsFullfilled = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackingPacketStatus", wireType)
			}
			m.TrackingPacketStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemandOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrackingPacketStatus |= types1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDemandOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDemandOrder
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
func skipDemandOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDemandOrder
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
					return 0, ErrIntOverflowDemandOrder
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
					return 0, ErrIntOverflowDemandOrder
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
				return 0, ErrInvalidLengthDemandOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDemandOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDemandOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDemandOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDemandOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDemandOrder = fmt.Errorf("proto: unexpected end of group")
)
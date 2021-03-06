// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tictactoe1/Game.proto

package types

import (
	fmt "fmt"
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

type OpenGame struct {
	Initiator string `protobuf:"bytes,1,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *OpenGame) Reset()         { *m = OpenGame{} }
func (m *OpenGame) String() string { return proto.CompactTextString(m) }
func (*OpenGame) ProtoMessage()    {}
func (*OpenGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f8e8722aac4998, []int{0}
}
func (m *OpenGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OpenGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OpenGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OpenGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenGame.Merge(m, src)
}
func (m *OpenGame) XXX_Size() int {
	return m.Size()
}
func (m *OpenGame) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenGame.DiscardUnknown(m)
}

var xxx_messageInfo_OpenGame proto.InternalMessageInfo

func (m *OpenGame) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *OpenGame) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type CurrGame struct {
	X     string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	O     string `protobuf:"bytes,2,opt,name=o,proto3" json:"o,omitempty"`
	Uuid  string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Board []byte `protobuf:"bytes,4,opt,name=board,proto3" json:"board,omitempty"`
	Turn  string `protobuf:"bytes,5,opt,name=turn,proto3" json:"turn,omitempty"`
}

func (m *CurrGame) Reset()         { *m = CurrGame{} }
func (m *CurrGame) String() string { return proto.CompactTextString(m) }
func (*CurrGame) ProtoMessage()    {}
func (*CurrGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f8e8722aac4998, []int{1}
}
func (m *CurrGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrGame.Merge(m, src)
}
func (m *CurrGame) XXX_Size() int {
	return m.Size()
}
func (m *CurrGame) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrGame.DiscardUnknown(m)
}

var xxx_messageInfo_CurrGame proto.InternalMessageInfo

func (m *CurrGame) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *CurrGame) GetO() string {
	if m != nil {
		return m.O
	}
	return ""
}

func (m *CurrGame) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CurrGame) GetBoard() []byte {
	if m != nil {
		return m.Board
	}
	return nil
}

func (m *CurrGame) GetTurn() string {
	if m != nil {
		return m.Turn
	}
	return ""
}

type DoneGame struct {
	Initiator  string `protobuf:"bytes,1,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Challenger string `protobuf:"bytes,2,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Uuid       string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Board      []byte `protobuf:"bytes,4,opt,name=board,proto3" json:"board,omitempty"`
	Winner     string `protobuf:"bytes,5,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (m *DoneGame) Reset()         { *m = DoneGame{} }
func (m *DoneGame) String() string { return proto.CompactTextString(m) }
func (*DoneGame) ProtoMessage()    {}
func (*DoneGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f8e8722aac4998, []int{2}
}
func (m *DoneGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DoneGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DoneGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DoneGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoneGame.Merge(m, src)
}
func (m *DoneGame) XXX_Size() int {
	return m.Size()
}
func (m *DoneGame) XXX_DiscardUnknown() {
	xxx_messageInfo_DoneGame.DiscardUnknown(m)
}

var xxx_messageInfo_DoneGame proto.InternalMessageInfo

func (m *DoneGame) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *DoneGame) GetChallenger() string {
	if m != nil {
		return m.Challenger
	}
	return ""
}

func (m *DoneGame) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *DoneGame) GetBoard() []byte {
	if m != nil {
		return m.Board
	}
	return nil
}

func (m *DoneGame) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func init() {
	proto.RegisterType((*OpenGame)(nil), "avendauz.tictactoe1.tictactoe1.OpenGame")
	proto.RegisterType((*CurrGame)(nil), "avendauz.tictactoe1.tictactoe1.CurrGame")
	proto.RegisterType((*DoneGame)(nil), "avendauz.tictactoe1.tictactoe1.DoneGame")
}

func init() { proto.RegisterFile("tictactoe1/Game.proto", fileDescriptor_a4f8e8722aac4998) }

var fileDescriptor_a4f8e8722aac4998 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xe3, 0xff, 0xbf, 0xad, 0xd2, 0x53, 0xa7, 0x08, 0x50, 0x06, 0x64, 0x55, 0x99, 0xba,
	0x24, 0x51, 0x05, 0x23, 0x13, 0x20, 0xb1, 0x81, 0xd4, 0x91, 0xcd, 0x71, 0xac, 0xd6, 0x52, 0xeb,
	0x8b, 0xcc, 0x19, 0x02, 0x1f, 0x81, 0x89, 0x8f, 0xc5, 0xd8, 0x91, 0x11, 0x25, 0x5f, 0x04, 0x25,
	0x0d, 0x24, 0x23, 0x6c, 0xef, 0xdd, 0xbd, 0x9f, 0xfc, 0xe4, 0x83, 0x63, 0xd2, 0x92, 0x84, 0x24,
	0x54, 0xcb, 0xf4, 0x46, 0xec, 0x54, 0x52, 0x58, 0x24, 0x0c, 0xb8, 0x78, 0x54, 0x26, 0x17, 0xee,
	0x25, 0xe9, 0xf7, 0x03, 0x19, 0x5d, 0x80, 0x7f, 0x57, 0x28, 0xd3, 0x10, 0xc1, 0x29, 0x4c, 0xb5,
	0xd1, 0xa4, 0x05, 0xa1, 0x0d, 0xd9, 0x9c, 0x2d, 0xa6, 0xab, 0x7e, 0x10, 0x04, 0x30, 0x72, 0x4e,
	0xe7, 0xe1, 0xbf, 0x76, 0xd1, 0xea, 0x28, 0x07, 0xff, 0xca, 0x59, 0xdb, 0xd2, 0x33, 0x60, 0x65,
	0x47, 0xb1, 0xb2, 0x71, 0xd8, 0x45, 0x19, 0xfe, 0xb0, 0xff, 0x7b, 0x36, 0x38, 0x82, 0x71, 0x86,
	0xc2, 0xe6, 0xe1, 0x68, 0xce, 0x16, 0xb3, 0xd5, 0xc1, 0x34, 0x49, 0x72, 0xd6, 0x84, 0xe3, 0x43,
	0xb2, 0xd1, 0xd1, 0x2b, 0x03, 0xff, 0x1a, 0x8d, 0xfa, 0x45, 0x49, 0x0e, 0x20, 0x37, 0x62, 0xbb,
	0x55, 0x66, 0xad, 0x6c, 0xf7, 0xfe, 0x60, 0xf2, 0x87, 0x22, 0x27, 0x30, 0x79, 0xd2, 0xc6, 0x28,
	0xdb, 0x55, 0xe9, 0xdc, 0xe5, 0xed, 0x7b, 0xc5, 0xd9, 0xbe, 0xe2, 0xec, 0xb3, 0xe2, 0xec, 0xad,
	0xe6, 0xde, 0xbe, 0xe6, 0xde, 0x47, 0xcd, 0xbd, 0xfb, 0xf3, 0xb5, 0xa6, 0x8d, 0xcb, 0x12, 0x89,
	0xbb, 0xf4, 0xfb, 0xd7, 0x53, 0xd2, 0x32, 0x26, 0x21, 0x63, 0x42, 0x15, 0x2f, 0xd3, 0x32, 0x1d,
	0x5c, 0x89, 0x9e, 0x0b, 0xf5, 0x90, 0x4d, 0xda, 0x3b, 0x9d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x6d, 0x09, 0xfe, 0xb1, 0xc0, 0x01, 0x00, 0x00,
}

func (m *OpenGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OpenGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OpenGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Initiator) > 0 {
		i -= len(m.Initiator)
		copy(dAtA[i:], m.Initiator)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Initiator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CurrGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Turn) > 0 {
		i -= len(m.Turn)
		copy(dAtA[i:], m.Turn)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Turn)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.O) > 0 {
		i -= len(m.O)
		copy(dAtA[i:], m.O)
		i = encodeVarintGame(dAtA, i, uint64(len(m.O)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarintGame(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DoneGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DoneGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DoneGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Challenger) > 0 {
		i -= len(m.Challenger)
		copy(dAtA[i:], m.Challenger)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Challenger)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Initiator) > 0 {
		i -= len(m.Initiator)
		copy(dAtA[i:], m.Initiator)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Initiator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OpenGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Initiator)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *CurrGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.O)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Turn)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *DoneGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Initiator)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Challenger)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func sovGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OpenGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: OpenGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OpenGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Initiator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Initiator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *CurrGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: CurrGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field O", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.O = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = append(m.Board[:0], dAtA[iNdEx:postIndex]...)
			if m.Board == nil {
				m.Board = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Turn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *DoneGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: DoneGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DoneGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Initiator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Initiator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Challenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Challenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = append(m.Board[:0], dAtA[iNdEx:postIndex]...)
			if m.Board == nil {
				m.Board = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGame = fmt.Errorf("proto: unexpected end of group")
)

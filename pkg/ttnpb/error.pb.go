// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go.thethings.network/lorawan-stack/api/error.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ErrorDetails struct {
	Namespace            string        `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MessageFormat        string        `protobuf:"bytes,3,opt,name=message_format,json=messageFormat,proto3" json:"message_format,omitempty"`
	Attributes           *types.Struct `protobuf:"bytes,4,opt,name=attributes" json:"attributes,omitempty"`
	CorrelationID        string        `protobuf:"bytes,5,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ErrorDetails) Reset()      { *m = ErrorDetails{} }
func (*ErrorDetails) ProtoMessage() {}
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_4f37e99190d558bd, []int{0}
}
func (m *ErrorDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ErrorDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetails.Merge(dst, src)
}
func (m *ErrorDetails) XXX_Size() int {
	return m.Size()
}
func (m *ErrorDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetails proto.InternalMessageInfo

func (m *ErrorDetails) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ErrorDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ErrorDetails) GetMessageFormat() string {
	if m != nil {
		return m.MessageFormat
	}
	return ""
}

func (m *ErrorDetails) GetAttributes() *types.Struct {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ErrorDetails) GetCorrelationID() string {
	if m != nil {
		return m.CorrelationID
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorDetails)(nil), "ttn.lorawan.v3.ErrorDetails")
	golang_proto.RegisterType((*ErrorDetails)(nil), "ttn.lorawan.v3.ErrorDetails")
}
func (this *ErrorDetails) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ErrorDetails)
	if !ok {
		that2, ok := that.(ErrorDetails)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ErrorDetails")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ErrorDetails but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ErrorDetails but is not nil && this == nil")
	}
	if this.Namespace != that1.Namespace {
		return fmt.Errorf("Namespace this(%v) Not Equal that(%v)", this.Namespace, that1.Namespace)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.MessageFormat != that1.MessageFormat {
		return fmt.Errorf("MessageFormat this(%v) Not Equal that(%v)", this.MessageFormat, that1.MessageFormat)
	}
	if !this.Attributes.Equal(that1.Attributes) {
		return fmt.Errorf("Attributes this(%v) Not Equal that(%v)", this.Attributes, that1.Attributes)
	}
	if this.CorrelationID != that1.CorrelationID {
		return fmt.Errorf("CorrelationID this(%v) Not Equal that(%v)", this.CorrelationID, that1.CorrelationID)
	}
	return nil
}
func (this *ErrorDetails) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ErrorDetails)
	if !ok {
		that2, ok := that.(ErrorDetails)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.MessageFormat != that1.MessageFormat {
		return false
	}
	if !this.Attributes.Equal(that1.Attributes) {
		return false
	}
	if this.CorrelationID != that1.CorrelationID {
		return false
	}
	return true
}
func (m *ErrorDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.MessageFormat) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.MessageFormat)))
		i += copy(dAtA[i:], m.MessageFormat)
	}
	if m.Attributes != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintError(dAtA, i, uint64(m.Attributes.Size()))
		n1, err := m.Attributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.CorrelationID) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.CorrelationID)))
		i += copy(dAtA[i:], m.CorrelationID)
	}
	return i, nil
}

func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedErrorDetails(r randyError, easy bool) *ErrorDetails {
	this := &ErrorDetails{}
	this.Namespace = randStringError(r)
	this.Name = randStringError(r)
	this.MessageFormat = randStringError(r)
	if r.Intn(10) != 0 {
		this.Attributes = types.NewPopulatedStruct(r, easy)
	}
	this.CorrelationID = randStringError(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyError interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneError(r randyError) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringError(r randyError) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneError(r)
	}
	return string(tmps)
}
func randUnrecognizedError(r randyError, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldError(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldError(dAtA []byte, r randyError, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateError(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateError(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateError(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ErrorDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.MessageFormat)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.Attributes != nil {
		l = m.Attributes.Size()
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.CorrelationID)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	return n
}

func sovError(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozError(x uint64) (n int) {
	return sovError((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ErrorDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ErrorDetails{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`MessageFormat:` + fmt.Sprintf("%v", this.MessageFormat) + `,`,
		`Attributes:` + strings.Replace(fmt.Sprintf("%v", this.Attributes), "Struct", "types.Struct", 1) + `,`,
		`CorrelationID:` + fmt.Sprintf("%v", this.CorrelationID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringError(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ErrorDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrorDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageFormat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageFormat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = &types.Struct{}
			}
			if err := m.Attributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
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
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowError
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthError
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowError
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipError(dAtA[start:])
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
	ErrInvalidLengthError = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("go.thethings.network/lorawan-stack/api/error.proto", fileDescriptor_error_4f37e99190d558bd)
}
func init() {
	golang_proto.RegisterFile("go.thethings.network/lorawan-stack/api/error.proto", fileDescriptor_error_4f37e99190d558bd)
}

var fileDescriptor_error_4f37e99190d558bd = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4c, 0x14, 0x41,
	0x14, 0x86, 0xe7, 0x29, 0x9a, 0x30, 0x7a, 0x97, 0xb8, 0x8d, 0x1b, 0x42, 0x9e, 0xc4, 0xc4, 0x84,
	0x42, 0x66, 0x13, 0x28, 0xb4, 0x46, 0x34, 0xa1, 0x3d, 0x3b, 0x1b, 0x32, 0xbb, 0x0c, 0x73, 0x9b,
	0xdb, 0xdd, 0xd9, 0xcc, 0xbc, 0x95, 0x96, 0x92, 0xd2, 0xd2, 0xd2, 0x58, 0x51, 0x52, 0x52, 0x52,
	0x52, 0x52, 0x12, 0x0b, 0xc3, 0xce, 0x34, 0x94, 0x94, 0x94, 0xe6, 0xe6, 0xce, 0xdc, 0x95, 0x74,
	0xef, 0xff, 0xff, 0x6f, 0x66, 0xfe, 0xc9, 0xe3, 0xdb, 0xda, 0x08, 0x1a, 0x2b, 0x1a, 0x97, 0x8d,
	0x76, 0xa2, 0x51, 0x74, 0x6c, 0xec, 0x24, 0xab, 0x8c, 0x95, 0xc7, 0xb2, 0xd9, 0x72, 0x24, 0x8b,
	0x49, 0x26, 0xdb, 0x32, 0x53, 0xd6, 0x1a, 0x2b, 0x5a, 0x6b, 0xc8, 0x24, 0x43, 0xa2, 0x46, 0xcc,
	0x11, 0xf1, 0x7d, 0x67, 0x6d, 0x4b, 0x97, 0x34, 0xee, 0x72, 0x51, 0x98, 0x3a, 0xd3, 0x46, 0x9b,
	0x2c, 0x62, 0x79, 0x77, 0x14, 0x55, 0x14, 0x71, 0x9a, 0x1d, 0x5f, 0x5b, 0xd7, 0xc6, 0xe8, 0x4a,
	0x2d, 0x28, 0x47, 0xb6, 0x2b, 0x68, 0x96, 0xbe, 0xfd, 0x03, 0xfc, 0xe5, 0xe7, 0xe9, 0x63, 0x7b,
	0x8a, 0x64, 0x59, 0xb9, 0x64, 0x9d, 0xaf, 0x36, 0xb2, 0x56, 0xae, 0x95, 0x85, 0x4a, 0x61, 0x03,
	0x36, 0x57, 0x47, 0x0b, 0x23, 0x49, 0xf8, 0xca, 0x54, 0xa4, 0x4f, 0x62, 0x10, 0xe7, 0xe4, 0x1d,
	0x1f, 0xd6, 0xca, 0x39, 0xa9, 0xd5, 0xc1, 0x91, 0xb1, 0xb5, 0xa4, 0xf4, 0x69, 0x4c, 0x07, 0x73,
	0xf7, 0x4b, 0x34, 0x93, 0x0f, 0x9c, 0x4b, 0x22, 0x5b, 0xe6, 0x1d, 0x29, 0x97, 0xae, 0x6c, 0xc0,
	0xe6, 0x8b, 0xed, 0xd7, 0x62, 0x56, 0x4e, 0xfc, 0x2f, 0x27, 0xbe, 0xc6, 0x72, 0xa3, 0x25, 0x34,
	0xf9, 0xc8, 0x87, 0x85, 0xb1, 0x56, 0x55, 0x92, 0x4a, 0xd3, 0x1c, 0x94, 0x87, 0xe9, 0xb3, 0xe9,
	0xfd, 0xbb, 0xaf, 0xfc, 0xdf, 0x37, 0x83, 0x4f, 0x8b, 0x64, 0x7f, 0x6f, 0x34, 0x58, 0x02, 0xf7,
	0x0f, 0x77, 0x7f, 0xc3, 0x55, 0x8f, 0x70, 0xdd, 0x23, 0xdc, 0xf4, 0xc8, 0x6e, 0x7b, 0x84, 0xbb,
	0x1e, 0xd9, 0x7d, 0x8f, 0xec, 0xa1, 0x47, 0x38, 0xf1, 0x08, 0xa7, 0x1e, 0xd9, 0x99, 0x47, 0x38,
	0xf7, 0xc8, 0x2e, 0x3c, 0xb2, 0x4b, 0x8f, 0x70, 0xe5, 0x11, 0xae, 0x3d, 0xc2, 0x8d, 0x47, 0x76,
	0xeb, 0x11, 0xee, 0x3c, 0xb2, 0x7b, 0x8f, 0xf0, 0xe0, 0x91, 0x9d, 0x04, 0x64, 0xa7, 0x01, 0xe1,
	0x47, 0x40, 0xf6, 0x33, 0x20, 0xfc, 0x0a, 0xc8, 0xce, 0x02, 0xb2, 0xf3, 0x80, 0x70, 0x11, 0x10,
	0x2e, 0x03, 0xc2, 0xb7, 0xf7, 0x8f, 0x58, 0x72, 0x3b, 0xd1, 0x19, 0x51, 0xd3, 0xe6, 0xf9, 0xf3,
	0xf8, 0xf7, 0x9d, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0xba, 0x93, 0xc3, 0x1b, 0x02, 0x00,
	0x00,
}

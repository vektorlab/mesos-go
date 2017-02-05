// Code generated by protoc-gen-gogo.
// source: allocator/allocator.proto
// DO NOT EDIT!

/*
	Package allocator is a generated protocol buffer package.

	It is generated from these files:
		allocator/allocator.proto

	It has these top-level messages:
		InverseOfferStatus
*/
package allocator

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos "github.com/mesos/mesos-go"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InverseOfferStatus_Status int32

const (
	// We have not received a response yet. This is the default state before
	// receiving a response.
	UNKNOWN InverseOfferStatus_Status = 1
	// The framework is ok with the inverse offer. This means it will not
	// violate any SLAs and will attempt to evacuate any tasks running on the
	// agent. If the tasks are not evacuated by the framework, the operator can
	// manually shut down the slave knowing that the framework will not have
	// violated its SLAs.
	ACCEPT InverseOfferStatus_Status = 2
	// The framework wants to block the maintenance operation from happening. An
	// example would be that it cannot meet its SLA by losing resources.
	DECLINE InverseOfferStatus_Status = 3
)

var InverseOfferStatus_Status_name = map[int32]string{
	1: "UNKNOWN",
	2: "ACCEPT",
	3: "DECLINE",
}
var InverseOfferStatus_Status_value = map[string]int32{
	"UNKNOWN": 1,
	"ACCEPT":  2,
	"DECLINE": 3,
}

func (x InverseOfferStatus_Status) Enum() *InverseOfferStatus_Status {
	p := new(InverseOfferStatus_Status)
	*p = x
	return p
}
func (x InverseOfferStatus_Status) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(InverseOfferStatus_Status_name, int32(x))
}
func (x *InverseOfferStatus_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InverseOfferStatus_Status_value, data, "InverseOfferStatus_Status")
	if err != nil {
		return err
	}
	*x = InverseOfferStatus_Status(value)
	return nil
}

// *
// Describes the status of an inverse offer.
//
// This is a protobuf so as to be able to share the status to inverse offers
// through endpoints such as the maintenance status endpoint.
type InverseOfferStatus struct {
	Status      *InverseOfferStatus_Status `protobuf:"varint,1,req,name=status,enum=mesos.allocator.InverseOfferStatus_Status" json:"status,omitempty"`
	FrameworkId *mesos.FrameworkID         `protobuf:"bytes,2,req,name=framework_id" json:"framework_id,omitempty"`
	// Time, since the epoch, when this status was last updated.
	Timestamp *mesos.TimeInfo `protobuf:"bytes,3,req,name=timestamp" json:"timestamp,omitempty"`
}

func (m *InverseOfferStatus) Reset()      { *m = InverseOfferStatus{} }
func (*InverseOfferStatus) ProtoMessage() {}

func (m *InverseOfferStatus) GetStatus() InverseOfferStatus_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return UNKNOWN
}

func (m *InverseOfferStatus) GetFrameworkId() *mesos.FrameworkID {
	if m != nil {
		return m.FrameworkId
	}
	return nil
}

func (m *InverseOfferStatus) GetTimestamp() *mesos.TimeInfo {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterEnum("mesos.allocator.InverseOfferStatus_Status", InverseOfferStatus_Status_name, InverseOfferStatus_Status_value)
}
func (x InverseOfferStatus_Status) String() string {
	s, ok := InverseOfferStatus_Status_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *InverseOfferStatus) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*InverseOfferStatus)
	if !ok {
		return fmt.Errorf("that is not of type *InverseOfferStatus")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *InverseOfferStatus but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *InverseOfferStatusbut is not nil && this == nil")
	}
	if this.Status != nil && that1.Status != nil {
		if *this.Status != *that1.Status {
			return fmt.Errorf("Status this(%v) Not Equal that(%v)", *this.Status, *that1.Status)
		}
	} else if this.Status != nil {
		return fmt.Errorf("this.Status == nil && that.Status != nil")
	} else if that1.Status != nil {
		return fmt.Errorf("Status this(%v) Not Equal that(%v)", this.Status, that1.Status)
	}
	if !this.FrameworkId.Equal(that1.FrameworkId) {
		return fmt.Errorf("FrameworkId this(%v) Not Equal that(%v)", this.FrameworkId, that1.FrameworkId)
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", this.Timestamp, that1.Timestamp)
	}
	return nil
}
func (this *InverseOfferStatus) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InverseOfferStatus)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Status != nil && that1.Status != nil {
		if *this.Status != *that1.Status {
			return false
		}
	} else if this.Status != nil {
		return false
	} else if that1.Status != nil {
		return false
	}
	if !this.FrameworkId.Equal(that1.FrameworkId) {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	return true
}
func (this *InverseOfferStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&allocator.InverseOfferStatus{")
	if this.Status != nil {
		s = append(s, "Status: "+valueToGoStringAllocator(this.Status, "allocator.InverseOfferStatus_Status")+",\n")
	}
	if this.FrameworkId != nil {
		s = append(s, "FrameworkId: "+fmt.Sprintf("%#v", this.FrameworkId)+",\n")
	}
	if this.Timestamp != nil {
		s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAllocator(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringAllocator(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *InverseOfferStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InverseOfferStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("status")
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintAllocator(data, i, uint64(*m.Status))
	}
	if m.FrameworkId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("framework_id")
	} else {
		data[i] = 0x12
		i++
		i = encodeVarintAllocator(data, i, uint64(m.FrameworkId.Size()))
		n1, err := m.FrameworkId.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Timestamp == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("timestamp")
	} else {
		data[i] = 0x1a
		i++
		i = encodeVarintAllocator(data, i, uint64(m.Timestamp.Size()))
		n2, err := m.Timestamp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Allocator(data []byte, offset int, v uint64) int {
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
func encodeFixed32Allocator(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAllocator(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedInverseOfferStatus(r randyAllocator, easy bool) *InverseOfferStatus {
	this := &InverseOfferStatus{}
	v1 := InverseOfferStatus_Status([]int32{1, 2, 3}[r.Intn(3)])
	this.Status = &v1
	this.FrameworkId = mesos.NewPopulatedFrameworkID(r, easy)
	this.Timestamp = mesos.NewPopulatedTimeInfo(r, easy)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAllocator interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAllocator(r randyAllocator) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAllocator(r randyAllocator) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneAllocator(r)
	}
	return string(tmps)
}
func randUnrecognizedAllocator(r randyAllocator, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldAllocator(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldAllocator(data []byte, r randyAllocator, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateAllocator(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateAllocator(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateAllocator(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateAllocator(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateAllocator(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateAllocator(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateAllocator(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *InverseOfferStatus) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		n += 1 + sovAllocator(uint64(*m.Status))
	}
	if m.FrameworkId != nil {
		l = m.FrameworkId.Size()
		n += 1 + l + sovAllocator(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovAllocator(uint64(l))
	}
	return n
}

func sovAllocator(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAllocator(x uint64) (n int) {
	return sovAllocator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *InverseOfferStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InverseOfferStatus{`,
		`Status:` + valueToStringAllocator(this.Status) + `,`,
		`FrameworkId:` + strings.Replace(fmt.Sprintf("%v", this.FrameworkId), "FrameworkID", "mesos.FrameworkID", 1) + `,`,
		`Timestamp:` + strings.Replace(fmt.Sprintf("%v", this.Timestamp), "TimeInfo", "mesos.TimeInfo", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAllocator(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *InverseOfferStatus) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocator
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
			return fmt.Errorf("proto: InverseOfferStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InverseOfferStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v InverseOfferStatus_Status
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (InverseOfferStatus_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameworkId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocator
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
				return ErrInvalidLengthAllocator
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FrameworkId == nil {
				m.FrameworkId = &mesos.FrameworkID{}
			}
			if err := m.FrameworkId.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocator
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
				return ErrInvalidLengthAllocator
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &mesos.TimeInfo{}
			}
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipAllocator(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAllocator
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("status")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("framework_id")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("timestamp")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAllocator(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllocator
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
					return 0, ErrIntOverflowAllocator
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
					return 0, ErrIntOverflowAllocator
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
				return 0, ErrInvalidLengthAllocator
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAllocator
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
				next, err := skipAllocator(data[start:])
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
	ErrInvalidLengthAllocator = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllocator   = fmt.Errorf("proto: integer overflow")
)

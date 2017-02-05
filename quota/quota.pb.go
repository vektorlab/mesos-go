// Code generated by protoc-gen-gogo.
// source: quota/quota.proto
// DO NOT EDIT!

/*
	Package quota is a generated protocol buffer package.

	It is generated from these files:
		quota/quota.proto

	It has these top-level messages:
		QuotaInfo
		QuotaRequest
		QuotaStatus
*/
package quota

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos "github.com/mesos/mesos-go"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TODO(joerg84): Add limits, i.e. upper bounds of resources that a
// role is allowed to use.
type QuotaInfo struct {
	// Quota is granted per role and not per framework, similar to
	// dynamic reservations.
	Role *string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// Principal which set the quota. Currently only operators can set quotas.
	Principal *string `protobuf:"bytes,2,opt,name=principal" json:"principal,omitempty"`
	// The guarantee that these resources are allocatable for the above role.
	// NOTE: `guarantee.role` should not specify any role except '*',
	// because quota does not reserve specific resources.
	Guarantee []*mesos.Resource `protobuf:"bytes,3,rep,name=guarantee" json:"guarantee,omitempty"`
}

func (m *QuotaInfo) Reset()      { *m = QuotaInfo{} }
func (*QuotaInfo) ProtoMessage() {}

func (m *QuotaInfo) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

func (m *QuotaInfo) GetPrincipal() string {
	if m != nil && m.Principal != nil {
		return *m.Principal
	}
	return ""
}

func (m *QuotaInfo) GetGuarantee() []*mesos.Resource {
	if m != nil {
		return m.Guarantee
	}
	return nil
}

// *
// `QuotaRequest` provides a schema for set quota JSON requests.
type QuotaRequest struct {
	// Disables the capacity heuristic check if set to `true`.
	Force *bool `protobuf:"varint,1,opt,name=force,def=0" json:"force,omitempty"`
	// The role for which to set quota.
	Role *string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	// The requested guarantee that these resources will be allocatable for
	// the above role.
	Guarantee []*mesos.Resource `protobuf:"bytes,3,rep,name=guarantee" json:"guarantee,omitempty"`
}

func (m *QuotaRequest) Reset()      { *m = QuotaRequest{} }
func (*QuotaRequest) ProtoMessage() {}

const Default_QuotaRequest_Force bool = false

func (m *QuotaRequest) GetForce() bool {
	if m != nil && m.Force != nil {
		return *m.Force
	}
	return Default_QuotaRequest_Force
}

func (m *QuotaRequest) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

func (m *QuotaRequest) GetGuarantee() []*mesos.Resource {
	if m != nil {
		return m.Guarantee
	}
	return nil
}

// *
// `QuotaStatus` describes the internal representation for the /quota/status
// response.
type QuotaStatus struct {
	// Quotas which are currently set, i.e. known to the master.
	Infos []*QuotaInfo `protobuf:"bytes,1,rep,name=infos" json:"infos,omitempty"`
}

func (m *QuotaStatus) Reset()      { *m = QuotaStatus{} }
func (*QuotaStatus) ProtoMessage() {}

func (m *QuotaStatus) GetInfos() []*QuotaInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (this *QuotaInfo) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*QuotaInfo)
	if !ok {
		return fmt.Errorf("that is not of type *QuotaInfo")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *QuotaInfo but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *QuotaInfobut is not nil && this == nil")
	}
	if this.Role != nil && that1.Role != nil {
		if *this.Role != *that1.Role {
			return fmt.Errorf("Role this(%v) Not Equal that(%v)", *this.Role, *that1.Role)
		}
	} else if this.Role != nil {
		return fmt.Errorf("this.Role == nil && that.Role != nil")
	} else if that1.Role != nil {
		return fmt.Errorf("Role this(%v) Not Equal that(%v)", this.Role, that1.Role)
	}
	if this.Principal != nil && that1.Principal != nil {
		if *this.Principal != *that1.Principal {
			return fmt.Errorf("Principal this(%v) Not Equal that(%v)", *this.Principal, *that1.Principal)
		}
	} else if this.Principal != nil {
		return fmt.Errorf("this.Principal == nil && that.Principal != nil")
	} else if that1.Principal != nil {
		return fmt.Errorf("Principal this(%v) Not Equal that(%v)", this.Principal, that1.Principal)
	}
	if len(this.Guarantee) != len(that1.Guarantee) {
		return fmt.Errorf("Guarantee this(%v) Not Equal that(%v)", len(this.Guarantee), len(that1.Guarantee))
	}
	for i := range this.Guarantee {
		if !this.Guarantee[i].Equal(that1.Guarantee[i]) {
			return fmt.Errorf("Guarantee this[%v](%v) Not Equal that[%v](%v)", i, this.Guarantee[i], i, that1.Guarantee[i])
		}
	}
	return nil
}
func (this *QuotaInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*QuotaInfo)
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
	if this.Role != nil && that1.Role != nil {
		if *this.Role != *that1.Role {
			return false
		}
	} else if this.Role != nil {
		return false
	} else if that1.Role != nil {
		return false
	}
	if this.Principal != nil && that1.Principal != nil {
		if *this.Principal != *that1.Principal {
			return false
		}
	} else if this.Principal != nil {
		return false
	} else if that1.Principal != nil {
		return false
	}
	if len(this.Guarantee) != len(that1.Guarantee) {
		return false
	}
	for i := range this.Guarantee {
		if !this.Guarantee[i].Equal(that1.Guarantee[i]) {
			return false
		}
	}
	return true
}
func (this *QuotaRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*QuotaRequest)
	if !ok {
		return fmt.Errorf("that is not of type *QuotaRequest")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *QuotaRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *QuotaRequestbut is not nil && this == nil")
	}
	if this.Force != nil && that1.Force != nil {
		if *this.Force != *that1.Force {
			return fmt.Errorf("Force this(%v) Not Equal that(%v)", *this.Force, *that1.Force)
		}
	} else if this.Force != nil {
		return fmt.Errorf("this.Force == nil && that.Force != nil")
	} else if that1.Force != nil {
		return fmt.Errorf("Force this(%v) Not Equal that(%v)", this.Force, that1.Force)
	}
	if this.Role != nil && that1.Role != nil {
		if *this.Role != *that1.Role {
			return fmt.Errorf("Role this(%v) Not Equal that(%v)", *this.Role, *that1.Role)
		}
	} else if this.Role != nil {
		return fmt.Errorf("this.Role == nil && that.Role != nil")
	} else if that1.Role != nil {
		return fmt.Errorf("Role this(%v) Not Equal that(%v)", this.Role, that1.Role)
	}
	if len(this.Guarantee) != len(that1.Guarantee) {
		return fmt.Errorf("Guarantee this(%v) Not Equal that(%v)", len(this.Guarantee), len(that1.Guarantee))
	}
	for i := range this.Guarantee {
		if !this.Guarantee[i].Equal(that1.Guarantee[i]) {
			return fmt.Errorf("Guarantee this[%v](%v) Not Equal that[%v](%v)", i, this.Guarantee[i], i, that1.Guarantee[i])
		}
	}
	return nil
}
func (this *QuotaRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*QuotaRequest)
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
	if this.Force != nil && that1.Force != nil {
		if *this.Force != *that1.Force {
			return false
		}
	} else if this.Force != nil {
		return false
	} else if that1.Force != nil {
		return false
	}
	if this.Role != nil && that1.Role != nil {
		if *this.Role != *that1.Role {
			return false
		}
	} else if this.Role != nil {
		return false
	} else if that1.Role != nil {
		return false
	}
	if len(this.Guarantee) != len(that1.Guarantee) {
		return false
	}
	for i := range this.Guarantee {
		if !this.Guarantee[i].Equal(that1.Guarantee[i]) {
			return false
		}
	}
	return true
}
func (this *QuotaStatus) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*QuotaStatus)
	if !ok {
		return fmt.Errorf("that is not of type *QuotaStatus")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *QuotaStatus but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *QuotaStatusbut is not nil && this == nil")
	}
	if len(this.Infos) != len(that1.Infos) {
		return fmt.Errorf("Infos this(%v) Not Equal that(%v)", len(this.Infos), len(that1.Infos))
	}
	for i := range this.Infos {
		if !this.Infos[i].Equal(that1.Infos[i]) {
			return fmt.Errorf("Infos this[%v](%v) Not Equal that[%v](%v)", i, this.Infos[i], i, that1.Infos[i])
		}
	}
	return nil
}
func (this *QuotaStatus) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*QuotaStatus)
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
	if len(this.Infos) != len(that1.Infos) {
		return false
	}
	for i := range this.Infos {
		if !this.Infos[i].Equal(that1.Infos[i]) {
			return false
		}
	}
	return true
}
func (this *QuotaInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&quota.QuotaInfo{")
	if this.Role != nil {
		s = append(s, "Role: "+valueToGoStringQuota(this.Role, "string")+",\n")
	}
	if this.Principal != nil {
		s = append(s, "Principal: "+valueToGoStringQuota(this.Principal, "string")+",\n")
	}
	if this.Guarantee != nil {
		s = append(s, "Guarantee: "+fmt.Sprintf("%#v", this.Guarantee)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QuotaRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&quota.QuotaRequest{")
	if this.Force != nil {
		s = append(s, "Force: "+valueToGoStringQuota(this.Force, "bool")+",\n")
	}
	if this.Role != nil {
		s = append(s, "Role: "+valueToGoStringQuota(this.Role, "string")+",\n")
	}
	if this.Guarantee != nil {
		s = append(s, "Guarantee: "+fmt.Sprintf("%#v", this.Guarantee)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QuotaStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&quota.QuotaStatus{")
	if this.Infos != nil {
		s = append(s, "Infos: "+fmt.Sprintf("%#v", this.Infos)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQuota(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringQuota(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
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
func (m *QuotaInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *QuotaInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Role != nil {
		data[i] = 0xa
		i++
		i = encodeVarintQuota(data, i, uint64(len(*m.Role)))
		i += copy(data[i:], *m.Role)
	}
	if m.Principal != nil {
		data[i] = 0x12
		i++
		i = encodeVarintQuota(data, i, uint64(len(*m.Principal)))
		i += copy(data[i:], *m.Principal)
	}
	if len(m.Guarantee) > 0 {
		for _, msg := range m.Guarantee {
			data[i] = 0x1a
			i++
			i = encodeVarintQuota(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *QuotaRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *QuotaRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Force != nil {
		data[i] = 0x8
		i++
		if *m.Force {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Role != nil {
		data[i] = 0x12
		i++
		i = encodeVarintQuota(data, i, uint64(len(*m.Role)))
		i += copy(data[i:], *m.Role)
	}
	if len(m.Guarantee) > 0 {
		for _, msg := range m.Guarantee {
			data[i] = 0x1a
			i++
			i = encodeVarintQuota(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *QuotaStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *QuotaStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Infos) > 0 {
		for _, msg := range m.Infos {
			data[i] = 0xa
			i++
			i = encodeVarintQuota(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Quota(data []byte, offset int, v uint64) int {
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
func encodeFixed32Quota(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintQuota(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedQuotaInfo(r randyQuota, easy bool) *QuotaInfo {
	this := &QuotaInfo{}
	if r.Intn(10) != 0 {
		v1 := randStringQuota(r)
		this.Role = &v1
	}
	if r.Intn(10) != 0 {
		v2 := randStringQuota(r)
		this.Principal = &v2
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.Guarantee = make([]*mesos.Resource, v3)
		for i := 0; i < v3; i++ {
			this.Guarantee[i] = mesos.NewPopulatedResource(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedQuotaRequest(r randyQuota, easy bool) *QuotaRequest {
	this := &QuotaRequest{}
	if r.Intn(10) != 0 {
		v4 := bool(bool(r.Intn(2) == 0))
		this.Force = &v4
	}
	if r.Intn(10) != 0 {
		v5 := randStringQuota(r)
		this.Role = &v5
	}
	if r.Intn(10) != 0 {
		v6 := r.Intn(10)
		this.Guarantee = make([]*mesos.Resource, v6)
		for i := 0; i < v6; i++ {
			this.Guarantee[i] = mesos.NewPopulatedResource(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedQuotaStatus(r randyQuota, easy bool) *QuotaStatus {
	this := &QuotaStatus{}
	if r.Intn(10) != 0 {
		v7 := r.Intn(10)
		this.Infos = make([]*QuotaInfo, v7)
		for i := 0; i < v7; i++ {
			this.Infos[i] = NewPopulatedQuotaInfo(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyQuota interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneQuota(r randyQuota) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringQuota(r randyQuota) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneQuota(r)
	}
	return string(tmps)
}
func randUnrecognizedQuota(r randyQuota, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldQuota(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldQuota(data []byte, r randyQuota, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateQuota(data, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		data = encodeVarintPopulateQuota(data, uint64(v9))
	case 1:
		data = encodeVarintPopulateQuota(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateQuota(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateQuota(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateQuota(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateQuota(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *QuotaInfo) Size() (n int) {
	var l int
	_ = l
	if m.Role != nil {
		l = len(*m.Role)
		n += 1 + l + sovQuota(uint64(l))
	}
	if m.Principal != nil {
		l = len(*m.Principal)
		n += 1 + l + sovQuota(uint64(l))
	}
	if len(m.Guarantee) > 0 {
		for _, e := range m.Guarantee {
			l = e.Size()
			n += 1 + l + sovQuota(uint64(l))
		}
	}
	return n
}

func (m *QuotaRequest) Size() (n int) {
	var l int
	_ = l
	if m.Force != nil {
		n += 2
	}
	if m.Role != nil {
		l = len(*m.Role)
		n += 1 + l + sovQuota(uint64(l))
	}
	if len(m.Guarantee) > 0 {
		for _, e := range m.Guarantee {
			l = e.Size()
			n += 1 + l + sovQuota(uint64(l))
		}
	}
	return n
}

func (m *QuotaStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.Infos) > 0 {
		for _, e := range m.Infos {
			l = e.Size()
			n += 1 + l + sovQuota(uint64(l))
		}
	}
	return n
}

func sovQuota(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuota(x uint64) (n int) {
	return sovQuota(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QuotaInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaInfo{`,
		`Role:` + valueToStringQuota(this.Role) + `,`,
		`Principal:` + valueToStringQuota(this.Principal) + `,`,
		`Guarantee:` + strings.Replace(fmt.Sprintf("%v", this.Guarantee), "Resource", "mesos.Resource", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaRequest{`,
		`Force:` + valueToStringQuota(this.Force) + `,`,
		`Role:` + valueToStringQuota(this.Role) + `,`,
		`Guarantee:` + strings.Replace(fmt.Sprintf("%v", this.Guarantee), "Resource", "mesos.Resource", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaStatus{`,
		`Infos:` + strings.Replace(fmt.Sprintf("%v", this.Infos), "QuotaInfo", "QuotaInfo", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQuota(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QuotaInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: QuotaInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Role = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Principal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Principal = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guarantee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guarantee = append(m.Guarantee, &mesos.Resource{})
			if err := m.Guarantee[len(m.Guarantee)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
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
func (m *QuotaRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: QuotaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Force", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Force = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Role = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guarantee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guarantee = append(m.Guarantee, &mesos.Resource{})
			if err := m.Guarantee[len(m.Guarantee)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
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
func (m *QuotaStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: QuotaStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Infos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Infos = append(m.Infos, &QuotaInfo{})
			if err := m.Infos[len(m.Infos)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
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
func skipQuota(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
				return 0, ErrInvalidLengthQuota
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuota
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
				next, err := skipQuota(data[start:])
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
	ErrInvalidLengthQuota = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuota   = fmt.Errorf("proto: integer overflow")
)

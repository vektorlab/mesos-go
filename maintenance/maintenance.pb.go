// Code generated by protoc-gen-gogo.
// source: maintenance/maintenance.proto
// DO NOT EDIT!

/*
	Package maintenance is a generated protocol buffer package.

	It is generated from these files:
		maintenance/maintenance.proto

	It has these top-level messages:
		Window
		Schedule
		ClusterStatus
*/
package maintenance

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos "github.com/mesos/mesos-go"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
import mesos_allocator "github.com/mesos/mesos-go/allocator"

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

// *
// A set of machines scheduled to go into maintenance
// in the same `unavailability`.
type Window struct {
	// Machines affected by this maintenance window.
	MachineIds []*mesos.MachineID `protobuf:"bytes,1,rep,name=machine_ids" json:"machine_ids,omitempty"`
	// Interval during which this set of machines is expected to be down.
	Unavailability *mesos.Unavailability `protobuf:"bytes,2,req,name=unavailability" json:"unavailability,omitempty"`
}

func (m *Window) Reset()      { *m = Window{} }
func (*Window) ProtoMessage() {}

func (m *Window) GetMachineIds() []*mesos.MachineID {
	if m != nil {
		return m.MachineIds
	}
	return nil
}

func (m *Window) GetUnavailability() *mesos.Unavailability {
	if m != nil {
		return m.Unavailability
	}
	return nil
}

// *
// A list of maintenance windows.
// For example, this may represent a rolling restart of agents.
type Schedule struct {
	Windows []*Window `protobuf:"bytes,1,rep,name=windows" json:"windows,omitempty"`
}

func (m *Schedule) Reset()      { *m = Schedule{} }
func (*Schedule) ProtoMessage() {}

func (m *Schedule) GetWindows() []*Window {
	if m != nil {
		return m.Windows
	}
	return nil
}

// *
// Represents the maintenance status of each machine in the cluster.
// The lists correspond to the `MachineInfo.Mode` enumeration.
type ClusterStatus struct {
	DrainingMachines []*ClusterStatus_DrainingMachine `protobuf:"bytes,1,rep,name=draining_machines" json:"draining_machines,omitempty"`
	DownMachines     []*mesos.MachineID               `protobuf:"bytes,2,rep,name=down_machines" json:"down_machines,omitempty"`
}

func (m *ClusterStatus) Reset()      { *m = ClusterStatus{} }
func (*ClusterStatus) ProtoMessage() {}

func (m *ClusterStatus) GetDrainingMachines() []*ClusterStatus_DrainingMachine {
	if m != nil {
		return m.DrainingMachines
	}
	return nil
}

func (m *ClusterStatus) GetDownMachines() []*mesos.MachineID {
	if m != nil {
		return m.DownMachines
	}
	return nil
}

type ClusterStatus_DrainingMachine struct {
	Id *mesos.MachineID `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// A list of the most recent responses to inverse offers from frameworks
	// running on this draining machine.
	Statuses []*mesos_allocator.InverseOfferStatus `protobuf:"bytes,2,rep,name=statuses" json:"statuses,omitempty"`
}

func (m *ClusterStatus_DrainingMachine) Reset()      { *m = ClusterStatus_DrainingMachine{} }
func (*ClusterStatus_DrainingMachine) ProtoMessage() {}

func (m *ClusterStatus_DrainingMachine) GetId() *mesos.MachineID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ClusterStatus_DrainingMachine) GetStatuses() []*mesos_allocator.InverseOfferStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (this *Window) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Window)
	if !ok {
		return fmt.Errorf("that is not of type *Window")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Window but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Windowbut is not nil && this == nil")
	}
	if len(this.MachineIds) != len(that1.MachineIds) {
		return fmt.Errorf("MachineIds this(%v) Not Equal that(%v)", len(this.MachineIds), len(that1.MachineIds))
	}
	for i := range this.MachineIds {
		if !this.MachineIds[i].Equal(that1.MachineIds[i]) {
			return fmt.Errorf("MachineIds this[%v](%v) Not Equal that[%v](%v)", i, this.MachineIds[i], i, that1.MachineIds[i])
		}
	}
	if !this.Unavailability.Equal(that1.Unavailability) {
		return fmt.Errorf("Unavailability this(%v) Not Equal that(%v)", this.Unavailability, that1.Unavailability)
	}
	return nil
}
func (this *Window) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Window)
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
	if len(this.MachineIds) != len(that1.MachineIds) {
		return false
	}
	for i := range this.MachineIds {
		if !this.MachineIds[i].Equal(that1.MachineIds[i]) {
			return false
		}
	}
	if !this.Unavailability.Equal(that1.Unavailability) {
		return false
	}
	return true
}
func (this *Schedule) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Schedule)
	if !ok {
		return fmt.Errorf("that is not of type *Schedule")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Schedule but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Schedulebut is not nil && this == nil")
	}
	if len(this.Windows) != len(that1.Windows) {
		return fmt.Errorf("Windows this(%v) Not Equal that(%v)", len(this.Windows), len(that1.Windows))
	}
	for i := range this.Windows {
		if !this.Windows[i].Equal(that1.Windows[i]) {
			return fmt.Errorf("Windows this[%v](%v) Not Equal that[%v](%v)", i, this.Windows[i], i, that1.Windows[i])
		}
	}
	return nil
}
func (this *Schedule) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Schedule)
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
	if len(this.Windows) != len(that1.Windows) {
		return false
	}
	for i := range this.Windows {
		if !this.Windows[i].Equal(that1.Windows[i]) {
			return false
		}
	}
	return true
}
func (this *ClusterStatus) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClusterStatus)
	if !ok {
		return fmt.Errorf("that is not of type *ClusterStatus")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClusterStatus but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClusterStatusbut is not nil && this == nil")
	}
	if len(this.DrainingMachines) != len(that1.DrainingMachines) {
		return fmt.Errorf("DrainingMachines this(%v) Not Equal that(%v)", len(this.DrainingMachines), len(that1.DrainingMachines))
	}
	for i := range this.DrainingMachines {
		if !this.DrainingMachines[i].Equal(that1.DrainingMachines[i]) {
			return fmt.Errorf("DrainingMachines this[%v](%v) Not Equal that[%v](%v)", i, this.DrainingMachines[i], i, that1.DrainingMachines[i])
		}
	}
	if len(this.DownMachines) != len(that1.DownMachines) {
		return fmt.Errorf("DownMachines this(%v) Not Equal that(%v)", len(this.DownMachines), len(that1.DownMachines))
	}
	for i := range this.DownMachines {
		if !this.DownMachines[i].Equal(that1.DownMachines[i]) {
			return fmt.Errorf("DownMachines this[%v](%v) Not Equal that[%v](%v)", i, this.DownMachines[i], i, that1.DownMachines[i])
		}
	}
	return nil
}
func (this *ClusterStatus) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClusterStatus)
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
	if len(this.DrainingMachines) != len(that1.DrainingMachines) {
		return false
	}
	for i := range this.DrainingMachines {
		if !this.DrainingMachines[i].Equal(that1.DrainingMachines[i]) {
			return false
		}
	}
	if len(this.DownMachines) != len(that1.DownMachines) {
		return false
	}
	for i := range this.DownMachines {
		if !this.DownMachines[i].Equal(that1.DownMachines[i]) {
			return false
		}
	}
	return true
}
func (this *ClusterStatus_DrainingMachine) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClusterStatus_DrainingMachine)
	if !ok {
		return fmt.Errorf("that is not of type *ClusterStatus_DrainingMachine")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClusterStatus_DrainingMachine but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClusterStatus_DrainingMachinebut is not nil && this == nil")
	}
	if !this.Id.Equal(that1.Id) {
		return fmt.Errorf("Id this(%v) Not Equal that(%v)", this.Id, that1.Id)
	}
	if len(this.Statuses) != len(that1.Statuses) {
		return fmt.Errorf("Statuses this(%v) Not Equal that(%v)", len(this.Statuses), len(that1.Statuses))
	}
	for i := range this.Statuses {
		if !this.Statuses[i].Equal(that1.Statuses[i]) {
			return fmt.Errorf("Statuses this[%v](%v) Not Equal that[%v](%v)", i, this.Statuses[i], i, that1.Statuses[i])
		}
	}
	return nil
}
func (this *ClusterStatus_DrainingMachine) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClusterStatus_DrainingMachine)
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
	if !this.Id.Equal(that1.Id) {
		return false
	}
	if len(this.Statuses) != len(that1.Statuses) {
		return false
	}
	for i := range this.Statuses {
		if !this.Statuses[i].Equal(that1.Statuses[i]) {
			return false
		}
	}
	return true
}
func (this *Window) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&maintenance.Window{")
	if this.MachineIds != nil {
		s = append(s, "MachineIds: "+fmt.Sprintf("%#v", this.MachineIds)+",\n")
	}
	if this.Unavailability != nil {
		s = append(s, "Unavailability: "+fmt.Sprintf("%#v", this.Unavailability)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Schedule) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&maintenance.Schedule{")
	if this.Windows != nil {
		s = append(s, "Windows: "+fmt.Sprintf("%#v", this.Windows)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClusterStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&maintenance.ClusterStatus{")
	if this.DrainingMachines != nil {
		s = append(s, "DrainingMachines: "+fmt.Sprintf("%#v", this.DrainingMachines)+",\n")
	}
	if this.DownMachines != nil {
		s = append(s, "DownMachines: "+fmt.Sprintf("%#v", this.DownMachines)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClusterStatus_DrainingMachine) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&maintenance.ClusterStatus_DrainingMachine{")
	if this.Id != nil {
		s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	}
	if this.Statuses != nil {
		s = append(s, "Statuses: "+fmt.Sprintf("%#v", this.Statuses)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMaintenance(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringMaintenance(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
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
func (m *Window) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Window) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MachineIds) > 0 {
		for _, msg := range m.MachineIds {
			data[i] = 0xa
			i++
			i = encodeVarintMaintenance(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Unavailability == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("unavailability")
	} else {
		data[i] = 0x12
		i++
		i = encodeVarintMaintenance(data, i, uint64(m.Unavailability.Size()))
		n1, err := m.Unavailability.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Schedule) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Schedule) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Windows) > 0 {
		for _, msg := range m.Windows {
			data[i] = 0xa
			i++
			i = encodeVarintMaintenance(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClusterStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DrainingMachines) > 0 {
		for _, msg := range m.DrainingMachines {
			data[i] = 0xa
			i++
			i = encodeVarintMaintenance(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DownMachines) > 0 {
		for _, msg := range m.DownMachines {
			data[i] = 0x12
			i++
			i = encodeVarintMaintenance(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterStatus_DrainingMachine) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClusterStatus_DrainingMachine) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("id")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMaintenance(data, i, uint64(m.Id.Size()))
		n2, err := m.Id.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Statuses) > 0 {
		for _, msg := range m.Statuses {
			data[i] = 0x12
			i++
			i = encodeVarintMaintenance(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Maintenance(data []byte, offset int, v uint64) int {
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
func encodeFixed32Maintenance(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMaintenance(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedWindow(r randyMaintenance, easy bool) *Window {
	this := &Window{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.MachineIds = make([]*mesos.MachineID, v1)
		for i := 0; i < v1; i++ {
			this.MachineIds[i] = mesos.NewPopulatedMachineID(r, easy)
		}
	}
	this.Unavailability = mesos.NewPopulatedUnavailability(r, easy)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSchedule(r randyMaintenance, easy bool) *Schedule {
	this := &Schedule{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Windows = make([]*Window, v2)
		for i := 0; i < v2; i++ {
			this.Windows[i] = NewPopulatedWindow(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClusterStatus(r randyMaintenance, easy bool) *ClusterStatus {
	this := &ClusterStatus{}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.DrainingMachines = make([]*ClusterStatus_DrainingMachine, v3)
		for i := 0; i < v3; i++ {
			this.DrainingMachines[i] = NewPopulatedClusterStatus_DrainingMachine(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v4 := r.Intn(10)
		this.DownMachines = make([]*mesos.MachineID, v4)
		for i := 0; i < v4; i++ {
			this.DownMachines[i] = mesos.NewPopulatedMachineID(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClusterStatus_DrainingMachine(r randyMaintenance, easy bool) *ClusterStatus_DrainingMachine {
	this := &ClusterStatus_DrainingMachine{}
	this.Id = mesos.NewPopulatedMachineID(r, easy)
	if r.Intn(10) != 0 {
		v5 := r.Intn(10)
		this.Statuses = make([]*mesos_allocator.InverseOfferStatus, v5)
		for i := 0; i < v5; i++ {
			this.Statuses[i] = mesos_allocator.NewPopulatedInverseOfferStatus(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMaintenance interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMaintenance(r randyMaintenance) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMaintenance(r randyMaintenance) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneMaintenance(r)
	}
	return string(tmps)
}
func randUnrecognizedMaintenance(r randyMaintenance, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMaintenance(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMaintenance(data []byte, r randyMaintenance, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMaintenance(data, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		data = encodeVarintPopulateMaintenance(data, uint64(v7))
	case 1:
		data = encodeVarintPopulateMaintenance(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMaintenance(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMaintenance(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMaintenance(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMaintenance(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Window) Size() (n int) {
	var l int
	_ = l
	if len(m.MachineIds) > 0 {
		for _, e := range m.MachineIds {
			l = e.Size()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	if m.Unavailability != nil {
		l = m.Unavailability.Size()
		n += 1 + l + sovMaintenance(uint64(l))
	}
	return n
}

func (m *Schedule) Size() (n int) {
	var l int
	_ = l
	if len(m.Windows) > 0 {
		for _, e := range m.Windows {
			l = e.Size()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	return n
}

func (m *ClusterStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.DrainingMachines) > 0 {
		for _, e := range m.DrainingMachines {
			l = e.Size()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	if len(m.DownMachines) > 0 {
		for _, e := range m.DownMachines {
			l = e.Size()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	return n
}

func (m *ClusterStatus_DrainingMachine) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovMaintenance(uint64(l))
	}
	if len(m.Statuses) > 0 {
		for _, e := range m.Statuses {
			l = e.Size()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	return n
}

func sovMaintenance(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMaintenance(x uint64) (n int) {
	return sovMaintenance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Window) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Window{`,
		`MachineIds:` + strings.Replace(fmt.Sprintf("%v", this.MachineIds), "MachineID", "mesos.MachineID", 1) + `,`,
		`Unavailability:` + strings.Replace(fmt.Sprintf("%v", this.Unavailability), "Unavailability", "mesos.Unavailability", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Schedule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Schedule{`,
		`Windows:` + strings.Replace(fmt.Sprintf("%v", this.Windows), "Window", "Window", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterStatus{`,
		`DrainingMachines:` + strings.Replace(fmt.Sprintf("%v", this.DrainingMachines), "ClusterStatus_DrainingMachine", "ClusterStatus_DrainingMachine", 1) + `,`,
		`DownMachines:` + strings.Replace(fmt.Sprintf("%v", this.DownMachines), "MachineID", "mesos.MachineID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterStatus_DrainingMachine) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterStatus_DrainingMachine{`,
		`Id:` + strings.Replace(fmt.Sprintf("%v", this.Id), "MachineID", "mesos.MachineID", 1) + `,`,
		`Statuses:` + strings.Replace(fmt.Sprintf("%v", this.Statuses), "InverseOfferStatus", "mesos_allocator.InverseOfferStatus", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMaintenance(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Window) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
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
			return fmt.Errorf("proto: Window: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Window: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineIds = append(m.MachineIds, &mesos.MachineID{})
			if err := m.MachineIds[len(m.MachineIds)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unavailability", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Unavailability == nil {
				m.Unavailability = &mesos.Unavailability{}
			}
			if err := m.Unavailability.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("unavailability")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Schedule) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
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
			return fmt.Errorf("proto: Schedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Windows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Windows = append(m.Windows, &Window{})
			if err := m.Windows[len(m.Windows)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
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
func (m *ClusterStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
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
			return fmt.Errorf("proto: ClusterStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrainingMachines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DrainingMachines = append(m.DrainingMachines, &ClusterStatus_DrainingMachine{})
			if err := m.DrainingMachines[len(m.DrainingMachines)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownMachines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownMachines = append(m.DownMachines, &mesos.MachineID{})
			if err := m.DownMachines[len(m.DownMachines)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
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
func (m *ClusterStatus_DrainingMachine) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
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
			return fmt.Errorf("proto: DrainingMachine: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DrainingMachine: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &mesos.MachineID{}
			}
			if err := m.Id.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
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
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statuses = append(m.Statuses, &mesos_allocator.InverseOfferStatus{})
			if err := m.Statuses[len(m.Statuses)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMaintenance(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMaintenance
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
					return 0, ErrIntOverflowMaintenance
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
					return 0, ErrIntOverflowMaintenance
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
				return 0, ErrInvalidLengthMaintenance
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMaintenance
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
				next, err := skipMaintenance(data[start:])
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
	ErrInvalidLengthMaintenance = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMaintenance   = fmt.Errorf("proto: integer overflow")
)

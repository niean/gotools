package nmap

import (
	"testing"
)

type DataStruct struct {
	Name string
	Age  int
}

func NewDataStruct(n string, a int) *DataStruct {
	return &DataStruct{Name: n, Age: a}
}

func TestSafeMap(t *testing.T) {
	sm := NewSafeMap()
	d1 := NewDataStruct("d1", 1)
	d2 := NewDataStruct("d2", 2)
	d3 := NewDataStruct("d3", 3)
	// Put & Get & Size
	sm.Put(d1.Name, d1)
	sm.Put(d2.Name, d2)
	sm.Put(d3.Name, d3)
	sm.Put(d1.Name, d1)
	d, found := sm.Get(d1.Name)
	if !(found && d.(*DataStruct).Name == d1.Name && sm.Size() == 3) {
		t.Error("error, Put & Get & Size")
	}
	// Clear
	sm.Clear()
	if !(sm.Size() == 0) {
		t.Error("error, Clear")
	}
	// Remove
	sm.Put(d1.Name, d1)
	sm.Put(d2.Name, d2)
	sm.Remove(d1.Name)
	d, found = sm.Get(d1.Name)
	if !(!found && sm.Size() == 1) {
		t.Error("error, Remove")
	}
	// ContainsKey
	sm.Clear()
	sm.Put(d1.Name, d1)
	sm.Put(d2.Name, d2)
	if !(sm.ContainsKey(d1.Name) && !sm.ContainsKey(d3.Name)) {
		t.Error("error, ContainsKey")
	}
	// IsEmpty
	sm.Clear()
	if !sm.IsEmpty() {
		t.Error("error, IsEmpty")
	}
	sm.Put(d1.Name, d1)
	if sm.IsEmpty() {
		t.Error("error, IsEmpty")
	}
}

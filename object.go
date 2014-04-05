package oden

import (
	"github.com/willf/bitset"
	"reflect"
)

type Object struct {
	data map[uint]IComponent
	bits bitset.BitSet
}

func NewObject() *Object {
	return &Object{
		data: make(map[uint]IComponent),
	}
}

// Add data container to this object with its index to use
// later for system processing
func (this *Object) AddComponent(data IComponent) {
	this.data[gDataManager.Get(data)] = data
	this.bits.Set(gDataManager.Get(data))
}

func (this *Object) Bits() *bitset.BitSet {
	return &this.bits
}

// Return data by its type name (the most inefficent)
func (this *Object) ComponentByName(name string) IComponent {
	for _, data := range this.data {
		if reflect.TypeOf(data).String() == name {
			return data
		}
	}
	return nil
}

// Return data by its type (more efficent)
func (this *Object) ComponentByType(data IComponent) IComponent {
	datatype := reflect.TypeOf(data)
	for _, data := range this.data {
		if datatype == reflect.TypeOf(data) {
			return data
		}
	}

	return nil
}

// Return data by its data index (the most efficent)
func (this *Object) ComponentByIndex(index uint) IComponent {
	return this.data[index]
}

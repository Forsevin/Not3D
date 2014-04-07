package oden

import (
	"github.com/willf/bitset"
	"reflect"
)

type Object struct {
	data map[uint]IComponent
	bits bitset.BitSet
	name string
}

func NewObject() *Object {
	return &Object{
		data: make(map[uint]IComponent),
		name: "NewObject",
	}
}

func (this *Object) Name() string {
	return this.name
}

func (this *Object) SetName(name string) {
	this.name = name
}

// Add data container to this object with its index to use
// later for system processing
func (this *Object) AddComponent(component IComponent) IComponent {
	this.data[gDataManager.Get(component)] = component
	this.bits.Set(gDataManager.Get(component))
	return component
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

// Simple return
func (this *Object) Component(component IComponent) IComponent {
	return this.data[gDataManager.Get(component)]
}

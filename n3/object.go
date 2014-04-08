package n3

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

func (object *Object) Name() string {
	return object.name
}

func (object *Object) SetName(name string) {
	object.name = name
}

// Add data container to object object with its index to use
// later for system processing
func (object *Object) AddComponent(component IComponent) IComponent {
	object.data[gDataManager.Get(component)] = component
	object.bits.Set(gDataManager.Get(component))
	return component
}

func (object *Object) Bits() *bitset.BitSet {
	return &object.bits
}

// Return data by its type name (the most inefficent)
func (object *Object) ComponentByName(name string) IComponent {
	for _, data := range object.data {
		if reflect.TypeOf(data).String() == name {
			return data
		}
	}
	return nil
}

// Return data by its type (more efficent)
func (object *Object) ComponentByType(data IComponent) IComponent {
	datatype := reflect.TypeOf(data)
	for _, data := range object.data {
		if datatype == reflect.TypeOf(data) {
			return data
		}
	}

	return nil
}

// Return data by its data index (the most efficent)
func (object *Object) ComponentByIndex(index uint) IComponent {
	return object.data[index]
}

// Simple return
func (object *Object) Component(component IComponent) IComponent {
	return object.data[gDataManager.Get(component)]
}

package n3

import (
	"reflect"

	"github.com/willf/bitset"
)

// Object holds states and provides way to modify it.
type Object struct {
	data map[uint]IComponent
	bits bitset.BitSet
	name string
}

// NewObject returns a new Object
func NewObject() *Object {
	return &Object{
		data: make(map[uint]IComponent),
		name: "NewObject",
	}
}

// Name returns the object's name
func (object *Object) Name() string {
	return object.name
}

// SetName sets the object's name
func (object *Object) SetName(name string) {
	object.name = name
}

// AddComponent adds a data container to an object with its index to use later for system processing
func (object *Object) AddComponent(component IComponent) IComponent {
	object.data[gBits.Get(component)] = component
	object.bits.Set(gBits.Get(component))
	return component
}

// Bits returns the objects internal bitset
func (object *Object) Bits() *bitset.BitSet {
	return &object.bits
}

// ComponentByName returns a Component by its type name; this is very inefficent
func (object *Object) ComponentByName(name string) IComponent {
	for _, data := range object.data {
		if reflect.TypeOf(data).String() == name {
			return data
		}
	}
	return nil
}

// ComponentByType returns a Component by its type; this is more efficient than ComponentByName
func (object *Object) ComponentByType(data IComponent) IComponent {
	datatype := reflect.TypeOf(data)
	for _, data := range object.data {
		if datatype == reflect.TypeOf(data) {
			return data
		}
	}

	return nil
}

// ComponentByIndex returns a Component by its data index; this is the most efficient way of getting a component
func (object *Object) ComponentByIndex(index uint) IComponent {
	return object.data[index]
}

// Component simply returns the providec component.
func (object *Object) Component(component IComponent) IComponent {
	return object.data[gBits.Get(component)]
}

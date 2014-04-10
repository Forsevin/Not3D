package n3

import "reflect"

// Bitmanager holds unique indexes for each component
type BitManager struct {
	uindex  uint
	indexes map[reflect.Type]uint
}

// NewBitManager returns a new BitManager
func NewBitManager() *BitManager {
	return &BitManager{
		uindex:  1,
		indexes: make(map[reflect.Type]uint),
	}
}

// Get returns the index for the provided IComponent
func (d *BitManager) Get(data IComponent) uint {
	if _, ok := d.indexes[reflect.TypeOf(data)]; !ok {
		d.indexes[reflect.TypeOf(data)] = d.uindex
	}

	d.uindex++
	return d.indexes[reflect.TypeOf(data)]
}

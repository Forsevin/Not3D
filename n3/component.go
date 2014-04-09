package n3

import "reflect"

// DataManager is ..
// TODO(j6n) describe what it does
type DataManager struct {
	uindex  uint
	indexes map[reflect.Type]uint
}

// NewDataManager returns a new DataManager
func NewDataManager() *DataManager {
	return &DataManager{
		uindex:  1,
		indexes: make(map[reflect.Type]uint),
	}
}

// Get returns the index for the provided IComponent
func (d *DataManager) Get(data IComponent) uint {
	if _, ok := d.indexes[reflect.TypeOf(data)]; !ok {
		d.indexes[reflect.TypeOf(data)] = d.uindex
	}

	d.uindex++
	return d.indexes[reflect.TypeOf(data)]
}

// IComponent represents an abstract component
// TODO(j6n) describe what it does
type IComponent interface {
	Index() uint
	SetIndex(index uint)
}

// Component is an concrete implementation of IComponent
// TODO(j6n) describe what it does
type Component struct {
	index uint
}

// Index returns the Component's current index
func (c *Component) Index() uint {
	// if index is zero when can presume it haven't been set and get a new one
	if c.index < 1 {
		c.index = gDataManager.Get(c)
	}

	return c.index
}

// SetIndex sets the Component's index.
func (c *Component) SetIndex(index uint) {
	c.index = index
}

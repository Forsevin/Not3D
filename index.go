// Handle component bits
package n3

import (
	"reflect"
)

func newCIndex() *cindex {
	return &cindex{
		indexes: make(map[reflect.Type]uint),
		index:   0,
	}
}

// store indexes for components
type cindex struct {
	indexes map[reflect.Type]uint
	index   uint
}

// return index for component
func (b *cindex) get(component interface{}) uint {
	t := reflect.TypeOf(component)
	if _, v := b.indexes[t]; !v {
		b.indexes[t] = b.index
		b.index++
	}
	return b.indexes[t]
}

package oden

import "reflect"

type DataManager struct {
	uindex  uint
	indexes map[reflect.Type]uint
}

func NewDataManager() *DataManager {
	return &DataManager{
		uindex:  1,
		indexes: make(map[reflect.Type]uint),
	}
}

func (this *DataManager) Get(data IComponent) uint {
	if _, ok := this.indexes[reflect.TypeOf(data)]; !ok {
		this.indexes[reflect.TypeOf(data)] = this.uindex
	}

	this.uindex += 1
	return this.indexes[reflect.TypeOf(data)]
}

type IComponent interface {
	Index() uint
	SetIndex(index uint)
}

type Component struct {
	index uint
}

func (this *Component) Index() uint {
	// if index is zero when can presume it haven't been set and get a new one
	if this.index < 1 {
		this.index = gDataManager.Get(this)
	}

	return this.index
}

func (this *Component) SetIndex(index uint) {
	this.index = index
}

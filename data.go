package oden

import "reflect"

type DataManager struct {
	uindex  uint
	indexes map[reflect.Type]uint
}

func NewDataManager() *DataManager {
	return &DataManager{
		uindex:  0,
		indexes: make(map[reflect.Type]uint),
	}
}

func (this *DataManager) Get(data IData) uint {
	if _, ok := this.indexes[reflect.TypeOf(data)]; !ok {
		this.indexes[reflect.TypeOf(data)] = this.uindex
	}

	this.uindex += 1
	return this.indexes[reflect.TypeOf(data)]
}

type IData interface {
	Index() uint
	SetIndex(index uint)
}

type Data struct {
	index uint
}

func (this *Data) Index() uint {
	return this.index
}

func (this *Data) SetIndex(index uint) {
	this.index = index
}

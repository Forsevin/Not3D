package oden

import (
	"github.com/willf/bitset"
)

type ISystem interface {
	Initialize()
	ProcessObject(object *Object)
	Check(object *Object)
	SetDataInterest(index uint)
	Process()
	RemoveObjects()
	SetBase(base *Base)
}

type System struct {
	activeObjects []*Object
	aspect        bitset.BitSet
	ProcessFunc   func(object *Object)
	base          *Base
}

func (this *System) Process() {
	if this.ProcessFunc == nil {
		panic("ProcessFunc not set, you may not have inititalized a system or not set the function at all")
	}
	for _, object := range this.activeObjects {
		this.ProcessFunc(object)
	}
}

func (this *System) SetDataInterest(index uint) {
	this.aspect.Set(index)
}

func (this *System) SetBase(base *Base) {

}

// Check if this object is of interest, if it is it will be added to active
// array of objects to be processed
func (this *System) Check(object *Object) {
	objectBits := object.Bits()
	var interested bool = true

	for i, v := this.aspect.NextSet(0); v != false; i, v = this.aspect.NextSet(i) {
		if !objectBits.Test(i) {
			interested = false
		}
		i += 1
	}

	if interested {
		this.activeObjects = append(this.activeObjects, object)
	}
}

// Remove all objects
func (this *System) RemoveObjects() {
	this.activeObjects = nil
}

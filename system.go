package oden

import (
	"fmt"
	"github.com/willf/bitset"
	"reflect"
)

// System interface, all systems to be processed need to implement these methods
// $ Initialize - Called before used
// $ ProcessObject - Each object of interest will be processed by this method
// $ Check - Check a object if it possesses the data in interest
// $ SetDataInterest - Register a data type of interest for this system
// $ Process - Calls ProcessObject for each object
// $ RemoveObject - Just remove all objects from this system
// $ SetBase - Set the base for this system (not really used, may be removed of no use are found)
type ISystem interface {
	Initialize()
	ProcessObject(object *Object)
	Check(object *Object)
	SetComponentInterest(component IComponent)
	Process()
	RemoveObjects()
	SetBase(base *Base)
}

// The base system (need to be "inherited")
// $ activeObjects - Objects of interest found by the check method
// $ aspect - A bitset of indexes for data checked against object bits
// $ ProcessFunc - Because Golang got no overloading the method need to be specificed manually (proboably in the initialized method)
// $ base - If something need to be gotten from the base (currently not in use, see last comment of ISystem)
type System struct {
	activeObjects []*Object
	aspect        bitset.BitSet
	ProcessFunc   func(object *Object)
	base          *Base
}

func (this *System) Process() {
	if this.ProcessFunc == nil {
		fmt.Println("For type:", reflect.TypeOf(this))
		panic("ProcessFunc not set, you may not have inititalized a system or not set the function at all")
	}
	for _, object := range this.activeObjects {
		this.ProcessFunc(object)
	}
}

func (this *System) SetComponentInterest(component IComponent) {
	this.aspect.Set(gDataManager.Get(component))
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

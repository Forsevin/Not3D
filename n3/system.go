package n3

import (
	"fmt"
	"github.com/willf/bitset"
	"reflect"
)

type ISystem interface {
	// Called before use
	Initialize()
	// Every object of interest get passed to this function
	ProcessObject(object *Object)
	// Check if object has bits of interest
	Check(object *Object)
	// Set a component as of interest
	AddComponent(component IComponent)
	// Calls ProcessObject for each object
	Process()
	// Simply removes all objects from activeObjects
	RemoveObjects()
	// Sets base
	SetBase(base *Base)
	// Called before processing
	Begin()
	// Called after processing
	End()
}

type System struct {
	// Objects put there by Check that will be processed
	activeObjects []*Object
	// Bits for components of interest
	aspect bitset.BitSet
	// Because go doesn't have overloading we'll have to set our process method manually
	ProcessFunc func(object *Object)
	// If something has to be retrieved from base
	base *Base
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

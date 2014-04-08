package n3

import (
	"fmt"
	"github.com/willf/bitset"
	"reflect"
)

type ISystem interface {
	// Called before use
	Initialize()
	// Every object of interest get passed to system function
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

func (system *System) Process() {
	if system.ProcessFunc == nil {
		fmt.Println("For type:", reflect.TypeOf(system))
		panic("ProcessFunc not set, you may not have inititalized a system or not set the function at all")
	}

	for _, object := range system.activeObjects {
		system.ProcessFunc(object)
	}
}

func (system *System) AddComponent(component IComponent) {
	system.aspect.Set(gDataManager.Get(component))
}

func (system *System) SetBase(base *Base) {

}

// Check if system object is of interest, if it is it will be added to active
// array of objects to be processed
func (system *System) Check(object *Object) {
	objectBits := object.Bits()
	var interested bool = true

	for i, v := system.aspect.NextSet(0); v != false; i, v = system.aspect.NextSet(i) {
		if !objectBits.Test(i) {
			interested = false
		}
		i += 1
	}
	if interested {
		system.activeObjects = append(system.activeObjects, object)
	}
}

// Remove all objects
func (system *System) RemoveObjects() {
	system.activeObjects = nil
}

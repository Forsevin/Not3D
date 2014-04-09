package n3

import (
	"fmt"
	"reflect"

	"github.com/willf/bitset"
)

// ISystem is an abstract System
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

// System implements the ISystem
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

// Process calls the process function on each object owned by the system
func (s *System) Process() {
	if s.ProcessFunc == nil {
		fmt.Println("For type:", reflect.TypeOf(s))
		panic("ProcessFunc not set, you may not have inititalized a system or not set the function at all")
	}

	for _, object := range s.activeObjects {
		s.ProcessFunc(object)
	}
}

// AddComponent adds the compoenet to the system's aspect
func (s *System) AddComponent(component IComponent) {
	s.aspect.Set(gBits.Get(component))
}

// SetBase sets the system's base
func (s *System) SetBase(base *Base) {
	//  TODO: implement
}

// Check if system object is of interest, if it is it will be added to active array of objects to be processed
func (s *System) Check(object *Object) {
	objectBits := object.Bits()
	interested := true

	for i, v := s.aspect.NextSet(0); v != false; i, v = s.aspect.NextSet(i) {
		if !objectBits.Test(i) {
			interested = false
		}
		i++
	}
	if interested {
		s.activeObjects = append(s.activeObjects, object)
	}
}

// RemoveObjects removes all of the objects in the system
func (s *System) RemoveObjects() {
	s.activeObjects = nil
}

package n3

import "fmt"

// Scene contains the data, upon selected  Base
// will take data from the scene and send to systems

type Scene struct {
	objects []*Object
}

func NewScene() *Scene {
	return &Scene{}
}

// Add an object to this scene
func (this *Scene) AddObject(object *Object) {
	this.objects = append(this.objects, object)
}

// Get a object by its name, pretty costly so use it wisely
func (this *Scene) ObjectByName(name string) *Object {
	for _, obj := range this.objects {
		if obj.Name() == name {
			return obj
		}
	}
	return nil
}

// Return all objects in this scene
func (this *Scene) Objects() []*Object {
	return this.objects
}

func (this *Scene) Debug() {
	for k, v := range this.objects {
		fmt.Println(k, v)
	}
}

package oden

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

// Return all objects in this scene
func (this *Scene) Objects() []*Object {
	return this.objects
}

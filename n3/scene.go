package n3

import "fmt"

// Scene contains the data, upon selected Base will take data from the scene and send to systems
type Scene struct {
	objects []*Object
}

// NewScene returns a NewScene
func NewScene() *Scene {
	return &Scene{}
}

// AddObject adds an object to scene
func (scene *Scene) AddObject(object *Object) {
	scene.objects = append(scene.objects, object)
}

// ObjectByName returns an object by its name, pretty costly so use it wisely
func (scene *Scene) ObjectByName(name string) *Object {
	for _, obj := range scene.objects {
		if obj.Name() == name {
			return obj
		}
	}
	return nil
}

// Objects returns all objects in scene
func (scene *Scene) Objects() []*Object {
	return scene.objects
}

// Debug prints all of the index, object mappings in the scene.
func (scene *Scene) Debug() {
	for k, v := range scene.objects {
		fmt.Println(k, v)
	}
}

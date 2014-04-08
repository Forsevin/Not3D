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

// Add an object to scene scene
func (scene *Scene) AddObject(object *Object) {
	scene.objects = append(scene.objects, object)
}

// Get a object by its name, pretty costly so use it wisely
func (scene *Scene) ObjectByName(name string) *Object {
	for _, obj := range scene.objects {
		if obj.Name() == name {
			return obj
		}
	}
	return nil
}

// Return all objects in scene scene
func (scene *Scene) Objects() []*Object {
	return scene.objects
}

func (scene *Scene) Debug() {
	for k, v := range scene.objects {
		fmt.Println(k, v)
	}
}

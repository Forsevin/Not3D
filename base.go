package oden

import "fmt"

// Having the data manager a global variable will make it a bit easier
// i.e we wont need to pass it arounds to every system or Data containers
var (
	dataManager *DataManager
)

// The interaction between user defined actions
// should be between this and the custom application

// $ activeScene - The currently in use
// $ scenes - Avaible scenes
// $ dataManager - Use this to initialize data containers
type Base struct {
	activeScene   *Scene
	scenes        map[string]*Scene
	dataManger    *DataManager
	globalSystems []ISystem
}

func NewBase() *Base {
	var base Base

	base.dataManger = NewDataManager()
	base.scenes = make(map[string]*Scene)

	dataManager = base.DataManager()

	// Base systems
	// The most important system, handles both the rendering of objects and
	// windows managments, surfaces etc.
	base.AddGlobalSystem(NewRenderSystem()).Initialize()
	// Allow objects to be manipulated by scripts (using Otto javascript implementation)
	base.AddGlobalSystem(NewScriptSystem()).Initialize()

	// Set base scene
	scene := NewScene()
	base.AddScene("main", scene)
	base.SetActiveScene("main")

	return &base

}

// Return the data manager used to intitialize
// data containers before use
func (this *Base) DataManager() *DataManager {
	return this.dataManger
}

func (this *Base) Process() {
	this.UpdateSystemObjectPossesions()
	for _, system := range this.globalSystems {
		system.Process()
	}
}

// Create a new object THIS CAN PROBOABLY BE DISCARGED (unless objects will need some additional data in the future)
// Note, this will not add this object to the scene
// To add a object to a scene use the scene instance
func (this *Base) CreateObject() *Object {
	return NewObject()
}

// Set/switch the active scene
func (this *Base) SetActiveScene(identifier string) {
	this.activeScene = this.scenes[identifier]
}

// Add a new scene
func (this *Base) AddScene(identifier string, scene *Scene) {
	this.scenes[identifier] = scene
}

// Get scene by identifier
func (this *Base) GetScene(identifier string) *Scene {
	return this.scenes[identifier]
}

// Get currently active scene
func (this *Base) GetActiveScene() *Scene {
	return this.activeScene
}

// Add a system to the current scene
func (this *Base) AddSystem(system ISystem) {

}

// Add a global system, this system will be proccesed in all scenes
func (this *Base) AddGlobalSystem(system ISystem) ISystem {
	this.globalSystems = append(this.globalSystems, system)
	return system
}

// Check objects towards systems
func (this *Base) UpdateSystemObjectPossesions() {
	for _, system := range this.globalSystems {
		for _, object := range this.activeScene.Objects() {
			system.RemoveObjects()
			system.Check(object)
		}
	}
}

// Just write what the fuck is in this scene
func (this *Base) DebugScene() {
	fmt.Println(this.activeScene.Objects())
	for _, object := range this.activeScene.Objects() {
		fmt.Println(object.Bits())
	}
}

package oden

import "fmt"

// Having the data manager a global variable will make it a bit easier
// we wont need to pass it arounds to every system or Data containers (it could be retrieved from base but fuck that)
var (
	dataManager *DataManager
)

// $ activeScene - The currently in use
// $ scenes - Avaible scenes
// $ dataManager - Use this to initialize data containers
// $ globalSystems - Systems that will be processed in every scene
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

// DataManager hold unique ids for each data
func (this *Base) DataManager() *DataManager {
	return this.dataManger
}

func (this *Base) Process() {
	// UpdateSystemObjectPossesions should only be called when needed (now it doesn't)
	// e.g when adding a new object, updating a system etc.
	this.UpdateSystemObjectPossesions()

	for _, system := range this.globalSystems {
		system.Process()
	}
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

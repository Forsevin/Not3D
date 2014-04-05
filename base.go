package oden

import (
	"github.com/jackyb/go-sdl2/sdl"
	"log"
	"os"
)

// Having the data manager a global variable will make it a bit easier
// we wont need to pass it arounds to every system or Data containers (it could be retrieved from base but fuck that)
var (
	gLogger      *log.Logger = log.New(os.Stdout, "[oden] ", 0)
	gDataManager *DataManager
)

// $ activeScene - The currently in use
// $ scenes - Avaible scenes
// $ globalSystems - Systems that will be processed in every scene
// $ Renderer - Renders shit 'n stuff
// $ quit - run while not quitting
type Base struct {
	activeScene   *Scene
	scenes        map[string]*Scene
	globalSystems []ISystem
	assets        *Assets
	renderer      *Renderer
	quit          bool
}

func New() *Base {
	var base Base

	gDataManager = NewDataManager()
	base.scenes = make(map[string]*Scene)

	base.SetRenderer(NewRenderer())
	base.assets = NewAssets()

	// Base systems
	// The most important system, handles both the rendering of objects and
	// windows managments, surfaces etc.
	base.AddGlobalSystem(NewRenderSystem(base.renderer, base.Assets())).Initialize()
	// Allow objects to be manipulated by scripts (using Otto javascript implementation)
	base.AddGlobalSystem(NewScriptSystem()).Initialize()

	// Set base scene
	scene := NewScene()
	base.AddScene("main", scene)
	base.SetActiveScene("main")

	return &base

}

func (this *Base) Process() {
	// UpdateSystemObjectPossesions should only be called when needed (now it doesn't)
	// e.g when adding a new object, updating a system etc.
	this.UpdateSystemObjectPossesions()

	for _, system := range this.globalSystems {
		system.Process()
	}
}

// Use default delta delayment
func (this *Base) DeltaSleep() {
	sdl.Delay(16)
}

// Set the window title
func (this *Base) SetWindowTitle(title string) {
	this.renderer.SetWindowTitle(title)
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
func (this *Base) Scene(identifier string) *Scene {
	return this.scenes[identifier]
}

// Get currently active scene
func (this *Base) ActiveScene() *Scene {
	return this.activeScene
}

func (this *Base) DeleteScene(identifier string) {
	delete(this.scenes, identifier)
}

func (this *Base) SetAssets(assets *Assets) {
	this.assets = assets
}

func (this *Base) Assets() *Assets {
	return this.assets
}

// Add a global system, this system will be proccesed in all scenes
func (this *Base) AddGlobalSystem(system ISystem) ISystem {
	this.globalSystems = append(this.globalSystems, system)
	return system
}

// Start the game loop
func (this *Base) Loop() {
	var event sdl.Event

	// Check for events in interest
	for this.quit != true {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				this.quit = true
			}
		}

		this.Process()
		sdl.Delay(160)
	}
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

func (this *Base) Log(msg string) {
	gLogger.Print(msg)
}

func (this *Base) SetQuit(quit bool) {
	this.quit = quit
}

func (this *Base) Quit() bool {
	return this.quit
}

func (this *Base) SetRenderer(renderer *Renderer) {
	this.renderer = renderer
}

func (this *Base) Error() string {
	return sdl.GetError().Error()
}

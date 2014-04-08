package oden

import (
	"github.com/jackyb/go-sdl2/sdl"
	"log"
	"os"
	//"time"
)

// Some global variables put here for convinience
var (
	gLogger      *log.Logger = log.New(os.Stdout, "[oden] ", 0)
	gDataManager *DataManager
)

type Graphics struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewGraphics() *Graphics {
	window := sdl.CreateWindow("Oden", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	return &Graphics{
		window:   window,
		renderer: sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED),
	}
}

type Base struct {
	// The scene currently in use, objects will be added to this
	activeScene *Scene
	// Avaible scenes switchable by SetActiveScene
	scenes map[string]*Scene
	// Systems that will be processed in every scene
	globalSystems []ISystem
	// Asset mananger
	assets *Assets
	// Handles windows and other SDL related variables
	graphics *Graphics
	// Handles input
	input *Input
	// Manange objects to be set and retrieved later
	prefabs *PrefabFactory
	// Loops while false
	quit bool
}

func New() *Base {
	var base Base

	// Set up our window
	base.SetGraphics(NewGraphics())

	gDataManager = NewDataManager()
	base.scenes = make(map[string]*Scene)

	base.assets = NewAssets(base.Graphics())
	base.SetInput(NewInput())
	// Base systems
	// The most important system, handles both the rendering of objects and
	// windows managments, surfaces etc.
	base.AddGlobalSystem(NewRenderSystem(base.Graphics())).Initialize()
	// Allow objects to be manipulated by scripts (using Otto javascript implementation)
	// We'll also create a Application Interface for it so it can work with our engine
	api := NewApi(&base)
	base.AddGlobalSystem(NewScriptSystem(api)).Initialize()

	base.prefabs = NewPrefabFactory()

	// Set base scene with a camera
	camera := base.CreateObject(0, 0)
	camera.AddComponent(NewCameraComponent())
	scene := NewScene()
	scene.AddObject(camera)
	base.AddScene("main", scene)
	base.SetActiveScene("main")

	return &base

}

func (this *Base) Process() {
	// UpdateSystemObjectPossesions should only be called when needed (now it doesn't)
	// e.g when adding a new object, updating a system etc.
	this.UpdateSystemObjectPossesions()

	for _, system := range this.globalSystems {
		system.Begin()
		system.Process()
		system.End()
	}
}

// Use default delta delayment
func (this *Base) DeltaSleep() {
	sdl.Delay(16)
}

// Create a new object (note: this doesn't add it to the scene)
func (this *Base) CreateObject(x, y int32) *Object {
	object := NewObject()
	return this.InitializeObject(object, x, y)
}

func (this *Base) InitializeObject(object *Object, x, y int32) *Object {
	object.AddComponent(NewTransformComponent(x, y))
	return object
}

// Set the window title
func (this *Base) SetWindowTitle(title string) {
	this.graphics.window.SetTitle(title)
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
	for this.quit == false {
		if this.input.Process() == true {
			this.SetQuit(true)
		}

		this.Process()

		sdl.Delay(16)
	}
}

// Check objects towards systems
func (this *Base) UpdateSystemObjectPossesions() {
	for _, system := range this.globalSystems {
		system.RemoveObjects()
		for _, object := range this.activeScene.Objects() {
			system.Check(object)
		}
	}
}

func (this *Base) Log(msg string) {
	gLogger.Print(msg)
}

func (this *Base) SetGraphics(graphics *Graphics) {
	this.graphics = graphics
}

func (this *Base) Graphics() *Graphics {
	return this.graphics
}

func (this *Base) SetQuit(quit bool) {
	this.quit = quit
}

func (this *Base) SetPrefabFactory(factory *PrefabFactory) {
	this.prefabs = factory
}

func (this *Base) Prefabs() *PrefabFactory {
	return this.prefabs
}

func (this *Base) Input() *Input {
	return this.input
}

func (this *Base) SetInput(input *Input) {
	this.input = input
}

func (this *Base) Quit() bool {
	return this.quit
}

func (this *Base) SDLLog(msg string) {

}

func (this *Base) Error() string {
	return sdl.GetError().Error()
}

package n3

import (
	"log"
	"os"

	"github.com/jackyb/go-sdl2/sdl"
	//"time"
)

// Some global variables put here for convinience
var (
	gLogger = log.New(os.Stdout, "[oden] ", 0)
	gBits   = NewBitManager()
)

// Graphics is a SDL context, a window and a renderer
type graphics struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

// NewGraphics returns a new Graphics context, with some defaults set.
func newGraphics() *graphics {
	window := sdl.CreateWindow("n3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	return &graphics{
		window:   window,
		renderer: sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED),
	}
}

// Base is a collection of systems, acting as the root system
type Base struct {
	// The scene currently in use, objects will be added to base
	activeScene *Scene
	// Avaible scenes switchable by SetActiveScene
	scenes map[string]*Scene
	// Systems that will be processed in every scene
	globalSystems []ISystem
	// Asset mananger
	assets *Assets
	// Handles windows and other SDL related variables
	graphics *graphics
	// Handles input
	input *Input
	// Manange objects to be set and retrieved later
	prefabs *PrefabFactory
	// Loops while false
	quit bool
}

// New returns a new Base with its components already built
func New() *Base {
	base := &Base{}

	base.graphics = newGraphics()
	base.scenes = make(map[string]*Scene)
	base.assets = NewAssets(base.graphics)

	base.SetInput(NewInput())

	base.AddGlobalSystem(NewRenderSystem(base.graphics)).Initialize()
	base.AddGlobalSystem(NewScriptSystem(NewAPI(base))).Initialize()

	base.prefabs = NewPrefabFactory()

	scene := NewScene()
	base.AddScene("main", scene)
	base.SetActiveScene("main")

	return base
}

// Process goes through each system and calls the batching/processing functions.
func (b *Base) Process() {
	// TODO only call UpdateSystemObjectPossession when needed
	b.UpdateSystemObjectPossesions()

	for _, system := range b.globalSystems {
		system.Begin()
		system.Process()
		system.End()
	}
}

// DeltaSleep tells SDL to wait
func (b *Base) DeltaSleep() {
	sdl.Delay(16)
}

// CreateObject creates and initializes a new Object at the given coords.
func (b *Base) CreateObject() *Object {
	object := NewObject()
	return b.InitializeObject(object, 0, 0)
}

// InitializeObject puts the object at the given coords.
func (b *Base) InitializeObject(object *Object, x, y int32) *Object {
	object.AddComponent(NewTransformComponent(x, y))
	return object
}

// SetWindowTitle sets the title of the window
func (b *Base) SetWindowTitle(title string) {
	b.graphics.window.SetTitle(title)
}

// SetActiveScene makes the active scene the one represented by the identifier
func (b *Base) SetActiveScene(identifier string) {
	b.activeScene = b.scenes[identifier]
}

// AddScene maps a scene to an indentifier
func (b *Base) AddScene(identifier string, scene *Scene) {
	b.scenes[identifier] = scene
}

// Scene takes an identifier and returns the associated scene
func (b *Base) Scene(identifier string) *Scene {
	return b.scenes[identifier]
}

// ActiveScene returns the current active scene
func (b *Base) ActiveScene() *Scene {
	return b.activeScene
}

// DeleteScene takes an id string and deletes the associated string
func (b *Base) DeleteScene(identifier string) {
	delete(b.scenes, identifier)
}

// SetAssets takes some assets and uses them in the base
func (b *Base) SetAssets(assets *Assets) {
	b.assets = assets
}

// Assets returns the currently held assets
func (b *Base) Assets() *Assets {
	return b.assets
}

// AddGlobalSystem adds a system to the base component
func (b *Base) AddGlobalSystem(system ISystem) ISystem {
	b.globalSystems = append(b.globalSystems, system)
	return system
}

// Loop is the main loop for the Base System, it continues until SetQuit is called with 'true'
func (b *Base) Loop() {
	for b.quit == false {
		if b.input.Process() == true {
			b.SetQuit(true)
		}

		b.Process()

		sdl.Delay(16)
	}
}

// UpdateSystemObjectPossesions removes all of the systems objects
// Then it'll check to see if they want any of the active scene's objects
func (b *Base) UpdateSystemObjectPossesions() {
	for _, system := range b.globalSystems {
		system.RemoveObjects()
		for _, object := range b.activeScene.Objects() {
			system.Check(object)
		}
	}
}

// Log delegates to the logger
func (b *Base) Log(msg string) {
	gLogger.Print(msg)
}

// Quit returns whether the base has quit its main loop.
func (b *Base) Quit() bool {
	return b.quit
}

// SetQuit sets the flag for the base to quit
func (b *Base) SetQuit(quit bool) {
	b.quit = quit
}

// SetPrefabFactory sets the base's PrefabFactory
func (b *Base) SetPrefabFactory(factory *PrefabFactory) {
	b.prefabs = factory
}

// Prefabs returns the base's PrefabFactory
func (b *Base) Prefabs() *PrefabFactory {
	return b.prefabs
}

// Input returns the base's input
func (b *Base) Input() *Input {
	return b.input
}

// SetInput sets the base's input
func (b *Base) SetInput(input *Input) {
	b.input = input
}

// SDLLog ...
func (b *Base) SDLLog(msg string) {
	// TODO: implement
}

func (b *Base) Error() string {
	return sdl.GetError().Error()
}

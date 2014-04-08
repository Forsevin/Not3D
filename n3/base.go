package n3

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
	// The scene currently in use, objects will be added to base
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

	base.SetGraphics(NewGraphics())

	gDataManager = NewDataManager()
	base.scenes = make(map[string]*Scene)
	base.assets = NewAssets(base.Graphics())
	base.SetInput(NewInput())
	base.AddGlobalSystem(NewRenderSystem(base.Graphics())).Initialize()
	api := NewApi(&base)
	base.AddGlobalSystem(NewScriptSystem(api)).Initialize()
	base.prefabs = NewPrefabFactory()

	camera := base.CreateObject(0, 0)
	camera.AddComponent(NewCameraComponent())

	scene := NewScene()
	scene.AddObject(camera)
	base.AddScene("main", scene)
	base.SetActiveScene("main")

	return &base

}

func (base *Base) Process() {
	// TODO only call UpdateSystemObjectPossession when needed
	base.UpdateSystemObjectPossesions()

	for _, system := range base.globalSystems {
		system.Begin()
		system.Process()
		system.End()
	}
}

func (base *Base) DeltaSleep() {
	sdl.Delay(16)
}

func (base *Base) CreateObject(x, y int32) *Object {
	object := NewObject()
	return base.InitializeObject(object, x, y)
}

func (base *Base) InitializeObject(object *Object, x, y int32) *Object {
	object.AddComponent(NewTransformComponent(x, y))
	return object
}

func (base *Base) SetWindowTitle(title string) {
	base.graphics.window.SetTitle(title)
}

func (base *Base) SetActiveScene(identifier string) {
	base.activeScene = base.scenes[identifier]
}

func (base *Base) AddScene(identifier string, scene *Scene) {
	base.scenes[identifier] = scene
}

func (base *Base) Scene(identifier string) *Scene {
	return base.scenes[identifier]
}

func (base *Base) ActiveScene() *Scene {
	return base.activeScene
}

func (base *Base) DeleteScene(identifier string) {
	delete(base.scenes, identifier)
}

func (base *Base) SetAssets(assets *Assets) {
	base.assets = assets
}

func (base *Base) Assets() *Assets {
	return base.assets
}

func (base *Base) AddGlobalSystem(system ISystem) ISystem {
	base.globalSystems = append(base.globalSystems, system)
	return system
}

func (base *Base) Loop() {
	for base.quit == false {
		if base.input.Process() == true {
			base.SetQuit(true)
		}

		base.Process()

		sdl.Delay(16)
	}
}

func (base *Base) UpdateSystemObjectPossesions() {
	for _, system := range base.globalSystems {
		system.RemoveObjects()
		for _, object := range base.activeScene.Objects() {
			system.Check(object)
		}
	}
}

func (base *Base) Log(msg string) {
	gLogger.Print(msg)
}

func (base *Base) SetGraphics(graphics *Graphics) {
	base.graphics = graphics
}

func (base *Base) Graphics() *Graphics {
	return base.graphics
}

func (base *Base) SetQuit(quit bool) {
	base.quit = quit
}

func (base *Base) SetPrefabFactory(factory *PrefabFactory) {
	base.prefabs = factory
}

func (base *Base) Prefabs() *PrefabFactory {
	return base.prefabs
}

func (base *Base) Input() *Input {
	return base.input
}

func (base *Base) SetInput(input *Input) {
	base.input = input
}

func (base *Base) Quit() bool {
	return base.quit
}

func (base *Base) SDLLog(msg string) {

}

func (base *Base) Error() string {
	return sdl.GetError().Error()
}

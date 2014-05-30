/* Core.go
 * Wraps up the functionality
 * Example:
 *
 * e := Gowerk.New()
 * e.Run()
 */
package n3

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
)

var (
	LOG     = true
	GUIText = &TextComponent{}
	Aspect  = &AspectComponent{}
	Sprite  = &SpriteComponent{}
)

// Core represents the top-level application of Gowerk
type Core struct {
	assets   *Assets
	platform *Platform
	logger   *log.Logger
	scene    *Scene
	index    *cindex
	scenes   map[string]*Scene
	systems  []SystemInterface
	quit     bool
	sceneStr string
}

// New creates a standard Core and return its pointer
func New() *Core {
	c := &Core{
		platform: NewPlatform(),
		logger:   log.New(os.Stdout, "Gowerk", -1),
		quit:     false,
		index:    newCIndex(),
		scenes:   make(map[string]*Scene),
		sceneStr: "main",
	}
	c.scene = newScene(c.index)
	c.assets = newAssets(c.platform)
	return c
}

// Inititalize core systems such as graphics, physcics etc.
func (c *Core) UseCoreSystems() {
	c.AddSystem(
		NewRenderSystem(c.platform),
	)
	c.AddSystem(
		NewFontSystem(),
	)
	c.AddSystem(
		NewScriptSystem(),
	)
}

func (c *Core) NewScene() *Scene {
	return newScene(c.index)
}

// Start processing system in our infinite loop (start the application)
func (c *Core) Run() {
	// start in fullscreen
	//c.platform.window.FullScreen()

	// make sure the standard font has been added
	if c.Assets().GetFont("std") == nil {
		c.logger.Fatalln("Standard font 'std' has not been loaded")
	}

	for c.quit != true {
		c.quit = c.platform.input.Process()
		c.platform.renderer.renderer.Clear()
		c.update()
		for _, system := range c.systems {
			system.Begin()
			system.Process()
			system.End()
		}
		c.platform.renderer.renderer.Present()
		sdl.Delay(16)
	}
}

// check entities towards the active systems
func (c *Core) update() {
	for _, s := range c.systems {
		s.clear()
		for _, e := range c.scene.entities.entities {
			s.check(e)
		}
	}
}

func (c *Core) SwitchScene(name string) {
	if c.scenes[name] == nil {
		c.logger.Println("Attempted switch to non-existing scene", name)
		return
	}
	c.sceneStr = name
	c.scene = c.scenes[name]
}

func (c *Core) GetScene(name string) *Scene {
	if c.scenes[name] == nil {
		c.Log().Println("GetScene", name, ", scene doesn't exist")
		return c.scene
	}

	return c.scenes[name]
}

func (c *Core) AddScene(name string, scene *Scene) {
	c.scenes[name] = scene
}

// Add a system
func (c *Core) AddSystem(system SystemInterface) {
	system.setCore(c)
	system.Init()
	c.systems = append(c.systems, system)
}

func (c *Core) Quit() {
	c.quit = true
}

// return this cores entities
func (c *Core) Entities() *Entities {
	return c.scene.entities
}

func (c *Core) Assets() *Assets {
	return c.assets
}

func (c *Core) Renderer() *Renderer {
	return c.platform.renderer
}

func (c *Core) Platform() *Platform {
	return c.platform
}

func (c *Core) Input() *Input {
	return c.platform.input
}

func (c *Core) Log() *log.Logger {
	return c.logger
}

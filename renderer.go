package oden

import "github.com/jackyb/go-sdl2/sdl"

type Renderer struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

type Texture2D struct {
}

func NewRenderer() *Renderer {
	window := sdl.CreateWindow("oden application", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if window == nil {
		gLogger.Fatalf("Could not create window:", sdl.GetError())
	}

	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		gLogger.Fatalf("Could not create renderer:", sdl.GetError())
	}

	return &Renderer{
		window:   window,
		renderer: renderer,
	}
}

func (this *Renderer) SetWindowTitle(title string) {
	this.window.SetTitle(title)
}

func (this *Renderer) DrawSprite(tex Texture2D) {

}

// For earily testing, might still be usable
func (this *Renderer) DrawRect(x, y, w, h int32) {
	rect := sdl.Rect{x, y, w, h}
	this.renderer.SetDrawColor(0, 0, 0, 0)
	this.renderer.Clear()
	this.renderer.SetDrawColor(255, 0, 255, 255)
	this.renderer.FillRect(&rect)
	this.renderer.Present()
}

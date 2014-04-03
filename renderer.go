package oden

import "github.com/jackyb/go-sdl2/sdl"

// Handles everything related to graphics

type Renderer struct {
	window *sdl.Window
}

func NewRenderer() *Renderer {
	return &Renderer{
		window: sdl.CreateWindow("OdenEngine (Testing)", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN),
	}
}

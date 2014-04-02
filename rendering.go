package oden

import "fmt"

// Aspect contains data for X and Y cordinates and Width and Height for rendering
type AspectData struct {
	Data
	x, y, w, h int
}

func NewAspectData(x, y int) *AspectData {
	return &AspectData{
		x: x,
		y: y,
	}
}

// Just contain some general data from the render system
type RenderData struct {
	Data
	visible bool
}

func NewRenderData() *RenderData {
	return &RenderData{}
}

type RenderSystem struct {
	System
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

func (this *RenderSystem) Initialize() {
	// Create windows and what not
	this.setup()

	this.ProcessFunc = this.ProcessObject

	// Interested in objects with aspect and render data
	this.SetDataInterest(gDataManager.Get(new(RenderData)))
	this.SetDataInterest(gDataManager.Get(new(AspectData)))
}

func (this *RenderSystem) ProcessObject(object *Object) {
	fmt.Println("Rendered something")
}

// Create a window, surfaces etc. What is necessary for rendering
func (this *RenderSystem) setup() {

}

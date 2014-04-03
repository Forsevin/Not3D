package oden

import (
	"fmt"
	//"github.com/jackyb/go-sdl2/sdl"
)

type RenderSystem struct {
	System
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

func (this *RenderSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject

	// Interested in objects with aspect and render data
	this.SetDataInterest(gDataManager.Get(new(RenderData)))
	this.SetDataInterest(gDataManager.Get(new(AspectData)))
}

func (this *RenderSystem) ProcessObject(object *Object) {
	fmt.Println("Rendered something")
}

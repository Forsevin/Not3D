package oden

type RenderSystem struct {
	System
	renderer *Renderer
}

func NewRenderSystem(renderer *Renderer) *RenderSystem {
	return &RenderSystem{
		renderer: renderer,
	}
}

func (this *RenderSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject

	// Interested in objects with aspect and render data
	this.SetDataInterest(gDataManager.Get(new(RenderData)))
	this.SetDataInterest(gDataManager.Get(new(AspectData)))
}

func (this *RenderSystem) ProcessObject(object *Object) {
	aspectdata := object.DataByIndex(gDataManager.Get(new(AspectData))).(*AspectData)
	this.renderer.DrawRect(aspectdata.x, aspectdata.y, aspectdata.w, aspectdata.h)
	aspectdata.x += 1
}

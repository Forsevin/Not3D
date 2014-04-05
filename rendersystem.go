package oden

type RenderSystem struct {
	System
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

func (this *RenderSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject
}

func (this *RenderSystem) ProcessObject(object *Object) {

}

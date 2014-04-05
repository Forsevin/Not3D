package oden

type RenderSystem struct {
	System
	renderer *Renderer
	assets   *Assets
}

func NewRenderSystem(renderer *Renderer, assets *Assets) *RenderSystem {
	return &RenderSystem{
		renderer: renderer,
		assets:   assets,
	}
}

func (this *RenderSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject

	// Interested in objects with aspect and render data
	this.SetDataInterest(gDataManager.Get(new(SpriteData)))
	this.SetDataInterest(gDataManager.Get(new(AspectData)))
}

func (this *RenderSystem) ProcessObject(object *Object) {
	aspectdata := object.DataByIndex(gDataManager.Get(new(AspectData))).(*AspectData)
	spritedata := object.DataByIndex(gDataManager.Get(new(SpriteData))).(*SpriteData)
	this.renderer.DrawSprite(aspectdata.x, aspectdata.y, aspectdata.w, aspectdata.h, this.assets.ImageAsset(spritedata.Asset))
	aspectdata.x += 1
}

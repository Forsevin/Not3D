package oden

type RenderSystem struct {
	System
	//Maybe this should be moved somewhere else?
	spriteBatch *SpriteBatch
}

func NewRenderSystem(graphics *Graphics) *RenderSystem {
	return &RenderSystem{
		spriteBatch: NewSpriteBatch(graphics),
	}
}

func (this *RenderSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject
	this.SetComponentInterest(new(SpriteComponent))
	this.SetComponentInterest(new(TransformComponent))
}

func (this *RenderSystem) ProcessObject(object *Object) {
	sprite := object.Component(new(SpriteComponent)).(*SpriteComponent)
	this.spriteBatch.Begin()
	this.spriteBatch.Draw(&sprite.texture, 0, 0, 0, 0)
	this.spriteBatch.End()
}

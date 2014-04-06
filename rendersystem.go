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
	transform := object.Component(new(TransformComponent)).(*TransformComponent)

	this.spriteBatch.Begin()
	this.spriteBatch.Draw(&sprite.Texture, transform.X, transform.Y, 10, 15)
	this.spriteBatch.End()
}

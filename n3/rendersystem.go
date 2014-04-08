package n3

type RenderSystem struct {
	System
	//Maybe rendersystem should be moved somewhere else?
	spriteBatch *SpriteBatch
}

func NewRenderSystem(graphics *Graphics) *RenderSystem {
	return &RenderSystem{
		spriteBatch: NewSpriteBatch(graphics),
	}
}

func (rendersystem *RenderSystem) Begin() {
	rendersystem.spriteBatch.Begin()
}

func (rendersystem *RenderSystem) End() {
	rendersystem.spriteBatch.End()
}

func (rendersystem *RenderSystem) Initialize() {
	rendersystem.ProcessFunc = rendersystem.ProcessObject
	rendersystem.AddComponent(new(SpriteComponent))
	rendersystem.AddComponent(new(TransformComponent))
}

func (rendersystem *RenderSystem) ProcessObject(object *Object) {
	sprite := object.Component(new(SpriteComponent)).(*SpriteComponent)
	transform := object.Component(new(TransformComponent)).(*TransformComponent)
	rendersystem.spriteBatch.Draw(&sprite.Texture, transform.X, transform.Y, 63, 87)
}

package n3

// RenderSystem is a system that renders to the screen
type RenderSystem struct {
	System
	//Maybe rendersystem should be moved somewhere else?
	spriteBatch *SpriteBatch
}

// NewRenderSystem returns a new RenderSystem with the graphics system set.
func NewRenderSystem(graphics_ *graphics) *RenderSystem {
	return &RenderSystem{
		spriteBatch: NewSpriteBatch(graphics_),
	}
}

// Begin the batch process
func (r *RenderSystem) Begin() {
	r.spriteBatch.Begin()
}

// End the batch process
func (r *RenderSystem) End() {
	r.spriteBatch.End()
}

// Initialize the rendersystem with some defaults
func (r *RenderSystem) Initialize() {
	r.ProcessFunc = r.ProcessObject
	r.AddComponent(new(SpriteComponent))
	r.AddComponent(new(TransformComponent))
}

// ProcessObject takes an object and adds it to the batching process
func (r *RenderSystem) ProcessObject(object *Object) {
	sprite := object.Component(new(SpriteComponent)).(*SpriteComponent)
	transform := object.Component(new(TransformComponent)).(*TransformComponent)
	r.spriteBatch.Draw(&sprite.Texture, transform.X, transform.Y, 63, 87)
}

/* factory.go
 * Provide common entities with components
 * Example:
 * Core().Factory().CreateDebugInfo().PutInWorld()
 */
package n3

type Factory struct {
	entities *Entities
}

func NewFactory(entities *Entities) *Factory {
	return &Factory{
		entities: entities,
	}
}

// Creates a debug box with general world and memory usage information
func (f *Factory) DebugScript() *Entity {
	entity := f.entities.Create()

	// standard components
	entity.AddComponent(NewAspectComponent(0, 0, LAYER_TOPGUI))
	entity.AddComponent(NewTextComponent("DebugBoxSTD", "arial"))
	entity.AddComponent(NewSpriteComponent(nil))

	// Set layer

	// scripts
	entity.AddComponent(NewScriptComponent(new(DebugBox)))

	return entity
}

func (f *Factory) RenderableObject() (*Entity, *AspectComponent, *SpriteComponent) {
	entity := f.entities.Create()

	sprite := NewSpriteComponent(nil)
	aspect := NewAspectComponent(0, 0, LAYER_FRONT)

	entity.AddComponent(aspect)
	entity.AddComponent(sprite)

	aspect.layer = 0

	return entity, aspect, sprite
}

func (f *Factory) KeyScript() *Entity {
	return f.entities.Create().AddComponent(NewScriptComponent(new(KeyScript)))
}

func (f *Factory) GUIText() (*Entity, *AspectComponent, *TextComponent) {
	entity := f.entities.Create()

	sprite := NewSpriteComponent(nil)
	aspect := NewAspectComponent(0, 0, LAYER_GUI)
	text := NewTextComponent("GUIText", "std")

	entity.AddComponent(sprite)
	entity.AddComponent(aspect)
	entity.AddComponent(text)

	return entity, aspect, text
}

/* Handles rendering of text */
package n3

type TextComponent struct {
	str  string
	font string
}

func NewTextComponent(str string, font string) *TextComponent {
	return &TextComponent{
		str:  str,
		font: font,
	}
}

func (t *TextComponent) SetString(str string) {
	t.str = str
}

// Handles rendering of fonts
type FontSystem struct {
	EntitySystem
}

func NewFontSystem() *FontSystem {
	return &FontSystem{}
}

func (f *FontSystem) Init() {
	f.Component(new(TextComponent))
	f.Component(new(SpriteComponent))
}

func (f *FontSystem) Begin() {

}

func (f *FontSystem) Process() {
	for _, entity := range f.entities {
		sprite := entity.GetComponent(new(SpriteComponent)).(*SpriteComponent)
		text := entity.GetComponent(new(TextComponent)).(*TextComponent)
		font := f.Core().Assets().GetFont("std")
		surface := font.RenderText_Blended(text.str, COLOR_RED)
		texture := f.Core().platform.renderer.renderer.CreateTextureFromSurface(surface)
		surface.Free()
		sprite.tex.Destroy()
		sprite.tex = texture
	}

}

func (f *FontSystem) End() {

}

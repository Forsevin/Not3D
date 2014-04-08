package n3

import "github.com/jackyb/go-sdl2/sdl"

type Texture2D struct {
	Texture *sdl.Texture
}

type SpriteBatch struct {
	graphics *Graphics
}

func NewSpriteBatch(graphics *Graphics) *SpriteBatch {
	return &SpriteBatch{
		graphics: graphics,
	}
}

// Get ready for a new draw (clear screen etc)
func (spritebatch *SpriteBatch) Begin() {
	spritebatch.graphics.renderer.Clear()
}

func (spritebatch *SpriteBatch) Draw(texture *Texture2D, x, y, w, h int32) {
	src := sdl.Rect{0, 0, 512, 512}
	dst := sdl.Rect{x, y, w, h}

	spritebatch.graphics.renderer.Copy(texture.Texture, &src, &dst)
}

// Draw our shit
func (spritebatch *SpriteBatch) End() {
	spritebatch.graphics.renderer.Present()
}

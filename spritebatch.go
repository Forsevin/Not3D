package oden

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
func (this *SpriteBatch) Begin() {
	this.graphics.renderer.Clear()
}

func (this *SpriteBatch) Draw(texture *Texture2D, x, y, w, h int32) {
	src := sdl.Rect{0, 0, 512, 512}
	dst := sdl.Rect{x, y, w, h}

	this.graphics.renderer.Copy(texture.Texture, &src, &dst)
}

// Draw our shit
func (this *SpriteBatch) End() {
	this.graphics.renderer.Present()
}

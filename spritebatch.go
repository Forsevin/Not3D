package oden

import "github.com/jackyb/go-sdl2/sdl"

type Texture2D struct {
}

type SpriteBatch struct {
	renderer *sdl.Renderer
}

// Get ready for a new draw (clear screen etc)
func (this *SpriteBatch) Begin() {

}

func (this *SpriteBatch) Draw(texture *Texture2D, x, y, w, h int32) {

}

// Draw our shit
func (this *SpriteBatch) End() {

}

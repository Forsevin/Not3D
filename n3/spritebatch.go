package n3

import "github.com/jackyb/go-sdl2/sdl"

// Texture2D is a wrapper around an sdl.Texture
type Texture2D struct {
	Texture *sdl.Texture
}

// SpriteBatch is a system for batched requests dealing with sprites
type SpriteBatch struct {
	graphics_ *graphics
}

// NewSpriteBatch returns a new SpriteBatch with the graphics set.
func NewSpriteBatch(graphics_ *graphics) *SpriteBatch {
	return &SpriteBatch{
		graphics_: graphics_,
	}
}

// Begin gets ready for a new draw (clear screen etc)
func (s *SpriteBatch) Begin() {
	s.graphics_.renderer.Clear()
}

// Draw the provided texture at the coordinates
func (s *SpriteBatch) Draw(texture *Texture2D, x, y, w, h int32) {
	src := sdl.Rect{0, 0, 512, 512}
	dst := sdl.Rect{x, y, w, h}

	s.graphics_.renderer.Copy(texture.Texture, &src, &dst)
}

// End finishes the batching and does the actual drawing.
func (s *SpriteBatch) End() {
	s.graphics_.renderer.Present()
}

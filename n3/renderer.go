package n3

import "github.com/jackyb/go-sdl2/sdl"

// Texture2D is a wrapper around an sdl.Texture
type Texture2D struct {
	Texture *sdl.Texture
}

type Vector2D struct {
	x, y int32
}

// SpriteBatch is a system for batched requests dealing with sprites
type Renderer struct {
	graphics *graphics
}

// NewSpriteBatch returns a new SpriteBatch with the graphics set.
func NewRenderer(graphics *graphics) *Renderer {
	return &Renderer{
		graphics: graphics,
	}
}

// Render present
func (r *Renderer) Render() {
	r.graphics.renderer.Present()
}

// Draw a texture to the screen
func (r *Renderer) DrawTex(texture *Texture2D, x, y int32, w, h int) {
	src := sdl.Rect{0, 0, 512, 512}
	dst := sdl.Rect{x, y, 10, 10}

	r.graphics.renderer.Copy(texture.Texture, &src, &dst)
}

// Clear screen
func (r *Renderer) Clear() {
	r.graphics.renderer.Clear()
}

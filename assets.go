// Handle loading and removal of assets such as textures and sounds
package n3

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

// Represent the top level asset structure
type Assets struct {
	platform *Platform
	images   map[string]*sdl.Texture
	fonts    map[string]*ttf.Font
}

// Returns a pointer to a new assets
func newAssets(p *Platform) *Assets {
	return &Assets{
		images:   make(map[string]*sdl.Texture),
		fonts:    make(map[string]*ttf.Font),
		platform: p,
	}
}

// Load a image file f of filetype t with name n
func (a *Assets) Load(t string, f string, n string) error {
	switch t {
	case "img":
		surface := img.Load(f)
		if surface == nil {
			return sdl.GetError()
		}
		texture := a.platform.renderer.renderer.CreateTextureFromSurface(surface)
		if texture == nil {
			return sdl.GetError()
		}
		a.images[n] = texture
	case "font":
		font, err := ttf.OpenFont(f, 12)
		if err != nil {
			return err
		}
		a.fonts[n] = font
	}

	return nil
}

func (a *Assets) GetFont(n string) *ttf.Font {
	return a.fonts[n]
}

func (a *Assets) RemoveRes(n string) {
	a.images[n].Destroy()
	delete(a.images, n)
}

func (a *Assets) ClearAll() {

}

func (a *Assets) GetImage(n string) (*sdl.Texture, error) {
	return a.images[n], nil
}

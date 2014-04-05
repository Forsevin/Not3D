package oden

import "github.com/jackyb/go-sdl2/sdl"

// Handle loading and storing of assets
// Until I'll figure out wether sdl_image actually works and how it works
// we'll have to use BMPs' :(

type Assets struct {
	imageAssets map[string]*sdl.Texture
	soundAssets map[string]int
	graphics    *Graphics
}

func NewAssets(graphics *Graphics) *Assets {
	return &Assets{
		imageAssets: make(map[string]*sdl.Texture),
		graphics:    graphics,
	}
}

func (this *Assets) LoadImageAsset(file string) {
	img := sdl.LoadBMP(file)
	if img == nil {
		// Maybe we'll load some nifty image here instead
		return
	}
	this.imageAssets[file] = this.graphics.renderer.CreateTextureFromSurface(img)
}

func (this *Assets) ImageAsset(name string) *sdl.Texture {
	return this.imageAssets[name]
}

func (this *Assets) Set(object *Object, asset string) {
	sprite := object.Component(new(SpriteComponent)).(*SpriteComponent)
	sprite.texture.texture = this.ImageAsset(asset)
}

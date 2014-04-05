package oden

import "github.com/jackyb/go-sdl2/sdl"

// Handle loading and storing of assets
// Until I'll figure out wether sdl_image actually works and how it works
// we'll have to use BMPs' :(

type Assets struct {
	imageAssets map[string]*sdl.Surface
	soundAssets map[string]int
}

func NewAssets() *Assets {
	return &Assets{
		imageAssets: make(map[string]*sdl.Surface),
	}
}

func (this *Assets) LoadImageAsset(file string) {
	img := sdl.LoadBMP(file)
	if img == nil {
		// Maybe we'll load some nifty image here instead
		return
	}
	this.imageAssets[file] = img
}

func (this *Assets) ImageAsset(name string) *sdl.Surface {
	return this.imageAssets[name]
}

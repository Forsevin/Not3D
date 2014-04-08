package n3

import (
	"github.com/jackyb/go-sdl2/sdl"
	"io/ioutil"
)

var (
	scriptDir string = "assets/scripts/"
	imageDir  string = "assets/images/"
)

// Handle loading and storing of assets
// Until I'll figure out wether sdl_image actually works and how it works we'll have to use BMPs' :(

type Assets struct {
	// For hardware accelerated textures
	graphics     *Graphics
	imageAssets  map[string]*sdl.Texture
	soundAssets  map[string]int
	scriptAssets map[string]string
}

func NewAssets(graphics *Graphics) *Assets {
	return &Assets{
		imageAssets:  make(map[string]*sdl.Texture),
		scriptAssets: make(map[string]string),
		graphics:     graphics,
	}
}

func (this *Assets) LoadImageAsset(file string) {
	img := sdl.LoadBMP(imageDir + file)
	if img == nil {
		gLogger.Fatalln("Couldn't load image asset:", sdl.GetError())
		return
	}
	this.imageAssets[file] = this.graphics.renderer.CreateTextureFromSurface(img)
}

func (this *Assets) ImageAsset(name string) *sdl.Texture {
	return this.imageAssets[name]
}

func (this *Assets) LoadScriptAsset(file string) {
	raw, err := ioutil.ReadFile(scriptDir + file)
	if err != nil {
		gLogger.Fatalln("Couldn't load script asset:", err)
	}
	this.scriptAssets[file] = string(raw)
}

func (this *Assets) ScriptAsset(script string) string {
	return this.scriptAssets[script]
}
package n3

import (
	"io/ioutil"

	"github.com/jackyb/go-sdl2/sdl"
)

// Assets are a collection of assets to be used in the engine
type Assets struct {
	graphics     *Graphics
	imageAssets  map[string]*sdl.Texture
	soundAssets  map[string]int
	scriptAssets map[string]string
	scriptDir    string
	imageDir     string
}

// NewAssets takes a pointer to graphics object and returns an Assets pointer
func NewAssets(graphics *Graphics) *Assets {
	return &Assets{
		imageAssets:  make(map[string]*sdl.Texture),
		scriptAssets: make(map[string]string),
		graphics:     graphics,
		imageDir:     "assets/images/",
		scriptDir:    "assets/scripts/",
	}
}

// LoadImageAsset loads a file at the provided path, adding it to the assets
func (assets *Assets) LoadImageAsset(file string) {
	img := sdl.LoadBMP(assets.imageDir + file)
	if img == nil {
		gLogger.Fatalln("Couldn't load image asset:", sdl.GetError())
		return
	}
	assets.imageAssets[file] = assets.graphics.renderer.CreateTextureFromSurface(img)
}

// ImageAsset returns the image with the provided name
func (assets *Assets) ImageAsset(name string) *sdl.Texture {
	return assets.imageAssets[name]
}

// LoadScriptAsset loads the script at the provided path, adding it to the assets
func (assets *Assets) LoadScriptAsset(file string) {
	raw, err := ioutil.ReadFile(assets.scriptDir + file)
	if err != nil {
		gLogger.Fatalln("Couldn't load script asset:", err)
	}
	assets.scriptAssets[file] = string(raw)
}

// ScriptAsset returns the script with the provided name
func (assets *Assets) ScriptAsset(script string) string {
	return assets.scriptAssets[script]
}

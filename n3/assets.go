package n3

import (
	"github.com/jackyb/go-sdl2/sdl"
	"io/ioutil"
)

type Assets struct {
	graphics     *Graphics
	imageAssets  map[string]*sdl.Texture
	soundAssets  map[string]int
	scriptAssets map[string]string
	scriptDir    string
	imageDir     string
}

func NewAssets(graphics *Graphics) *Assets {
	return &Assets{
		imageAssets:  make(map[string]*sdl.Texture),
		scriptAssets: make(map[string]string),
		graphics:     graphics,
		imageDir:     "assets/images/",
		scriptDir:    "assets/scripts/",
	}
}

func (assets *Assets) LoadImageAsset(file string) {
	img := sdl.LoadBMP(assets.imageDir + file)
	if img == nil {
		gLogger.Fatalln("Couldn't load image asset:", sdl.GetError())
		return
	}
	assets.imageAssets[file] = assets.graphics.renderer.CreateTextureFromSurface(img)
}

func (assets *Assets) ImageAsset(name string) *sdl.Texture {
	return assets.imageAssets[name]
}

func (assets *Assets) LoadScriptAsset(file string) {
	raw, err := ioutil.ReadFile(assets.scriptDir + file)
	if err != nil {
		gLogger.Fatalln("Couldn't load script asset:", err)
	}
	assets.scriptAssets[file] = string(raw)
}

func (assets *Assets) ScriptAsset(script string) string {
	return assets.scriptAssets[script]
}

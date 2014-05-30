package n3

import (
	"github.com/veandco/go-sdl2/sdl"
	"sort"
)

const (
	LAYER_BACKGROUND = 0
	LAYER_FRONT      = 1
	LAYER_TOPFRONT   = 2
	LAYER_GUI        = 3
	LAYER_TOPGUI     = 4
)

type SpriteComponent struct {
	tex            *sdl.Texture
	scaleW, scaleH int
}

func NewSpriteComponent(tex *sdl.Texture) *SpriteComponent {
	return &SpriteComponent{
		tex: tex,
	}
}

func (s *SpriteComponent) SetTexture(tex *sdl.Texture) {
	s.tex = tex
}

// Render System work upon the AspectComponent
type AspectComponent struct {
	x, y  int32
	layer int
}

func NewAspectComponent(x, y int32, layer int) *AspectComponent {
	return &AspectComponent{
		x:     x,
		y:     y,
		layer: layer,
	}
}

func (a *AspectComponent) Intersect(x, y int32) bool {
	return false
}

func (a *AspectComponent) Move(x, y int32) {
	a.x += x
	a.y += y
}

func (a *AspectComponent) Set(x, y int32) {
	a.x = x
	a.y = y
}

// required for sorting by struct field
type renderItem struct {
	layer  int
	aspect *AspectComponent
	tex    *sdl.Texture
}

type ByLayer []renderItem

func (a ByLayer) Len() int           { return len(a) }
func (a ByLayer) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLayer) Less(i, j int) bool { return a[i].layer < a[j].layer }

// Top level render systems implements systemInterface
type RenderSystem struct {
	EntitySystem
	platform   *Platform
	renderList []renderItem
}

func NewRenderSystem(p *Platform) *RenderSystem {
	return &RenderSystem{
		platform: p,
	}
}

func (r *RenderSystem) Init() {
	r.Component(new(AspectComponent))
	r.Component(new(SpriteComponent))
}

func (r *RenderSystem) Begin() {
	if r.platform.renderer.backgroundTex != nil {
		w, h := r.platform.window.window.GetSize()
		r.platform.renderer.renderer.Copy(r.platform.renderer.backgroundTex, &sdl.Rect{0, 0, int32(w), int32(h)}, nil)
	}
}

func (r *RenderSystem) Process() {
	for _, entity := range r.entities {
		aspect := entity.GetComponent(new(AspectComponent)).(*AspectComponent)
		sprite := entity.GetComponent(new(SpriteComponent)).(*SpriteComponent)
		r.renderList = append(r.renderList, renderItem{
			aspect.layer,
			aspect,
			sprite.tex,
		})
	}

	sort.Sort(ByLayer(r.renderList))
	for _, item := range r.renderList {
		var w, h int
		sdl.QueryTexture(item.tex, nil, nil, &w, &h)

		dst := sdl.Rect{item.aspect.x, item.aspect.y, int32(w), int32(h)}
		r.platform.renderer.renderer.Copy(item.tex, nil, &dst)
	}
}

func (r *RenderSystem) End() {
	r.renderList = nil
}

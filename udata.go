package oden

// Some default data often used
type ScriptData struct {
	Data
	scripts []string
}

// Aspect contains data for X and Y cordinates and Width and Height for rendering
type AspectData struct {
	Data
	x, y, w, h int32
}

func NewAspectData(x, y, w, h int32) *AspectData {
	return &AspectData{
		x: x,
		y: y,
		w: w,
		h: h,
	}
}

// Just contain some general data from the render system
type RenderData struct {
	Data
	visible bool
}

func NewRenderData() *RenderData {
	return &RenderData{}
}

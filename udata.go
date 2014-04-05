package oden

// Data for X, Y cordinates, width and height and angle for rotation
type TransformComponent struct {
	Component
	x, y  int32
	w, h  int32
	angle int32
}

func NewTransformComponent(x, y int32) *TransformComponent {
	return &TransformComponent{
		x:     x,
		y:     y,
		angle: 0,
	}
}

// The camera used by rendersystem
type CameraComponent struct {
}

func NewCameraComponent() *CameraComponent {
	return &CameraComponent{}
}

// Data for textures used for rendering
type SpriteComponent struct {
	texture Texture2D
}

func NewSpriteComponent() *SpriteComponent {
	return &SpriteComponent{}
}

// Data for the Gel script system
type GelScriptComponent struct {
}

func NewGelScriptComponent() *GelScriptComponent {
	return &GelScriptComponent{}
}

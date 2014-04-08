package n3

import "github.com/robertkrimen/otto"

// TransformComponent contains data for X, Y cordinates, width and height and angle for rotation
type TransformComponent struct {
	Component
	X, Y  int32
	W, H  int32
	Angle int32
}

// NewTransformComponent returns a new TransformComponent with its X,Y set
func NewTransformComponent(x, y int32) *TransformComponent {
	return &TransformComponent{
		X:     x,
		Y:     y,
		Angle: 0,
	}
}

// CameraComponent is the camera used by RenderSystem
type CameraComponent struct {
	Component
}

// NewCameraComponent returns a new CameraComponent
func NewCameraComponent() *CameraComponent {
	return &CameraComponent{}
}

// SpriteComponent contains data for textures used for rendering
type SpriteComponent struct {
	Component
	Texture Texture2D
}

// NewSpriteComponent returns a new SpriteComponent
func NewSpriteComponent() *SpriteComponent {
	return &SpriteComponent{}
}

// SpriteSetComponent is a set of string to textures
type SpriteSetComponent struct {
	Set map[string]Texture2D
}

// ScriptComponent contains data for the Gel script system
type ScriptComponent struct {
	Component
	// Script loaded either manually or by base.Assets
	Src string
	// Every script have its own runtime (?)
	runtime *otto.Otto
	object  *otto.Object
}

// NewScriptComponent returns a new ScriptComponent
func NewScriptComponent() *ScriptComponent {
	return &ScriptComponent{}
}

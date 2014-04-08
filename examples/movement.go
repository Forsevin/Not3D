package main

import "github.com/j6n/Not3D/n3"

func main() {
	engine := n3.New()

	// Load our resources
	engine.Assets().LoadImageAsset("player.bmp")
	engine.Assets().LoadScriptAsset("player.js")

	// Create player
	player := engine.CreateObject(0, 0)
	sprite := player.AddComponent(n3.NewSpriteComponent()).(*n3.SpriteComponent)
	script := player.AddComponent(n3.NewScriptComponent()).(*n3.ScriptComponent)

	// Set script and sprite
	script.Src = engine.Assets().ScriptAsset("player.js")
	sprite.Texture.Texture = engine.Assets().ImageAsset("player.bmp")

	// Add player to scene
	engine.ActiveScene().AddObject(player)

	engine.Loop()
}

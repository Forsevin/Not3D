package n3

import "github.com/jackyb/go-sdl2/sdl"

// Handles input

type Input struct {
	// Save states for all keys here
	keyStates map[string]bool
}

func NewInput() *Input {
	return &Input{
		keyStates: make(map[string]bool),
	}
}

// Process the input,
func (this *Input) Process() bool {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.KeyDownEvent:
			switch t.Keysym.Sym {
			case sdl.K_w:
				this.SetKeyDown("w")
			case sdl.K_a:
				this.SetKeyDown("a")
			case sdl.K_s:
				this.SetKeyDown("s")
			case sdl.K_d:
				this.SetKeyDown("d")
			}
		case *sdl.KeyUpEvent:
			switch t.Keysym.Sym {
			case sdl.K_w:
				this.SetKeyUp("w")
			case sdl.K_a:
				this.SetKeyUp("a")
			case sdl.K_s:
				this.SetKeyUp("s")
			case sdl.K_d:
				this.SetKeyUp("d")
			}
		}

	}

	return false
}

func (this *Input) KeyDown(key string) bool {
	return this.keyStates[key]
}

func (this *Input) SetKeyDown(key string) {
	this.keyStates[key] = true
}

func (this *Input) SetKeyUp(key string) {
	this.keyStates[key] = false
}

package oden

import "github.com/jackyb/go-sdl2/sdl"

// Handles input

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

// Process the input,
func (this *Input) Process() bool {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true

		}
	}

	return false
}

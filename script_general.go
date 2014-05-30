package n3

import (
	"runtime"
	"strconv"
)

// Most of these scripts are for testing capabilities as of now!

// Handle key events such as quitting at ESC etc
type KeyScript struct {
	Script
}

func (s *KeyScript) Init() {
}

func (s *KeyScript) Update() {
	if s.Engine().Input().KeyDown("ESCAPE") {
		s.Engine().Quit()
	}

}

// A box with some useful data
type DebugBox struct {
	Script
	text *TextComponent
}

func (d *DebugBox) Init() {
	d.text = d.Entity().GetComponent(new(TextComponent)).(*TextComponent)
}

func (d *DebugBox) Update() {
	entities := "Entities: " + strconv.Itoa(d.Engine().scene.entities.Len())
	components := "Components: ~" + strconv.Itoa(int(d.Engine().scene.components.indexes.index))
	scene := "Scene: \"" + d.Engine().sceneStr + "\""
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	memoryUsage := "GC: " + strconv.Itoa(int(memory.HeapInuse))
	d.text.SetString(scene + " - " + entities + " - " + components + " - " + memoryUsage)
}

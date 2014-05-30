//  system
package n3

import "github.com/Forsevin/n3/utils"

// interface for systems
type SystemInterface interface {
	Init()
	Begin()
	Process()
	End()

	Entities() []*Entity
	clear()
	setCore(c *Core)
	Component(component interface{})
	check(entity *Entity)
}

// basic entity processing system
type EntitySystem struct {
	entities []*Entity
	core     *Core
	bits     utils.BitSet
}

// Add a component
func (e *EntitySystem) Component(component interface{}) {
	e.bits.Set(e.Core().index.get(component))
}

// Check if component is of interest, if add it to entities
func (s *EntitySystem) check(entity *Entity) {
	objectBits := entity.bits
	interested := true

	for i, v := s.bits.NextSet(0); v != false; i, v = s.bits.NextSet(i) {
		if !objectBits.Test(i) {
			interested = false
		}
		i++
	}
	if interested {
		s.entities = append(s.entities, entity)
	}
}

func (e *EntitySystem) clear() {
	e.entities = nil
}

func (e *EntitySystem) setCore(c *Core) {
	e.core = c
}

func (e *EntitySystem) Entities() []*Entity {
	return e.entities
}

func (e *EntitySystem) Core() *Core {
	return e.core
}

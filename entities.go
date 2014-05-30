/* entities.go
 * Mananges entities e.g deleting and adding new entities to our world
 */
package n3

// Top level entity mananger
type Entities struct {
	entities    map[int]*Entity
	names       map[string]int
	cm          *componentMananger
	uniqueIndex *uniqueIndex
}

// return a pointer to a new Entities
func NewEntities(cm *componentMananger) *Entities {
	return &Entities{
		entities: make(map[int]*Entity),
		names:    make(map[string]int),
		cm:       cm,
		uniqueIndex: &uniqueIndex{
			uindex: 0,
		},
	}
}

// Add a entity for processing
func (e *Entities) Add(entity *Entity) {
	e.entities[entity.index] = entity
}

func (e *Entities) Create() *Entity {
	ent := newEntity(e.cm)
	ent.index = e.uniqueIndex.Get()
	e.entities[ent.index] = ent

	return ent

}

func (e *Entities) Len() int {
	return len(e.entities)
}

// Remove a entity from processing
func (e *Entities) Remove(entity *Entity) {
	e.cm.components[entity.index] = nil
	delete(e.entities, entity.index)
}

// Generates unique indexes for each entity
type uniqueIndex struct {
	uindex int
}

func (u *uniqueIndex) Get() int {
	u.uindex++
	return u.uindex
}

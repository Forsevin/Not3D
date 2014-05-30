/* entity .go
 */
package n3

import (
	"github.com/Forsevin/n3/utils"
)

type Entity struct {
	uniqueId int64
	index    int
	name     string
	bits     utils.BitSet
	cm       *componentMananger
	em       *Entities
}

func newEntity(cm *componentMananger) *Entity {
	return &Entity{
		cm: cm,
	}
}

func (e *Entity) AddComponent(component interface{}) *Entity {
	e.cm.add(e, component)
	return e
}

func (e *Entity) GetComponent(component interface{}) interface{} {
	c := e.cm.get(e, component)
	return c
}

func (e *Entity) SetName(name string) {
	e.name = name
}

func (e *Entity) Destroy() {

}

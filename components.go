/* components.go
 * Manange components for entities
 */
package n3

type componentMananger struct {
	indexes    *cindex
	components [][]interface{}
}

func newComponentMananger(indexes *cindex) *componentMananger {
	c := &componentMananger{
		indexes: indexes,
	}

	c.components = make([][]interface{}, 64)
	for i := range c.components {
		c.components[i] = make([]interface{}, 64)
	}

	return c
}

func (c *componentMananger) get(entity *Entity, component interface{}) interface{} {
	cmpt := c.components[entity.index][c.indexes.get(component)]
	if cmpt == nil {
		panic("Trying to retrieve non-existing component of entity " + entity.name)
	}
	return cmpt
}

func (c *componentMananger) add(entity *Entity, component interface{}) {
	c.components[entity.index][c.indexes.get(component)] = component
	entity.bits.Set(c.indexes.get(component))
}

package n3

type Scene struct {
	components *componentMananger
	entities   *Entities
	factory    *Factory
}

func newScene(indexes *cindex) *Scene {
	components := newComponentMananger(indexes)
	entities := NewEntities(components)

	return &Scene{
		components: components,
		entities:   entities,
		factory:    NewFactory(entities),
	}
}

func (s *Scene) Entities() *Entities {
	return s.entities
}

func (s *Scene) Factory() *Factory {
	return s.factory
}

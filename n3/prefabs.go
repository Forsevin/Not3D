package n3

// PrefabFactory is a factory that can cache prefab objects
type PrefabFactory struct {
	prefabs map[string]*Object
}

// NewPrefabFactory returns a new PrefabFactory
func NewPrefabFactory() *PrefabFactory {
	return &PrefabFactory{
		prefabs: make(map[string]*Object),
	}
}

// Prefab returns the prefab for the provided string
func (p *PrefabFactory) Prefab(prefab string) *Object {
	return p.prefabs[prefab]
}

// NewPrefab binds a name to an object, returning it
func (p *PrefabFactory) NewPrefab(name string, object *Object) *Object {
	p.prefabs[name] = object
	return object
}

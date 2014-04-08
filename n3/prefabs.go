package n3

type PrefabFactory struct {
	prefabs map[string]*Object
}

func NewPrefabFactory() *PrefabFactory {
	return &PrefabFactory{
		prefabs: make(map[string]*Object),
	}
}

func (prefabs *PrefabFactory) Prefab(prefab string) *Object {
	return prefabs.prefabs[prefab]
}

func (prefabs *PrefabFactory) NewPrefab(name string, object *Object) *Object {
	prefabs.prefabs[name] = object
	return object
}

package oden

type PrefabFactory struct {
	prefabs map[string]*Object
}

func NewPrefabFactory() *PrefabFactory {
	return &PrefabFactory{
		prefabs: make(map[string]*Object),
	}
}

func (this *PrefabFactory) Prefab(prefab string) *Object {
	return this.prefabs[prefab]
}

func (this *PrefabFactory) NewPrefab(name string, object *Object) *Object {
	this.prefabs[name] = object
	return object
}

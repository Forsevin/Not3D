package oden

type ScriptSystem struct {
	System
}

func NewScriptSystem() *ScriptSystem {
	return &ScriptSystem{}
}

func (this *ScriptSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject

}

// Generally, only the update method in source is called
func (this *ScriptSystem) ProcessObject(object *Object) {

}

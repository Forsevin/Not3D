package oden

import "github.com/robertkrimen/otto"

type ScriptSystem struct {
	System
	runtime *otto.Otto
	api     *Api
}

func NewScriptSystem(api *Api) *ScriptSystem {
	return &ScriptSystem{
		runtime: otto.New(),
		api:     api,
	}
}

func (this *ScriptSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject
	this.api.InitializeRuntime(this.runtime)
	this.SetComponentInterest(new(ScriptComponent))

}

// Generally, only the update method in source is called
func (this *ScriptSystem) ProcessObject(object *Object) {
	script := object.Component(new(ScriptComponent)).(*ScriptComponent)
	// For now we'll create a new runtime for each script
	this.runtime.Run(script.Src)
}

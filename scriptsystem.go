package oden

import "github.com/robertkrimen/otto"

type ScriptSystem struct {
	System
	api *Api
}

func NewScriptSystem(api *Api) *ScriptSystem {
	return &ScriptSystem{
		api: api,
	}
}

func (this *ScriptSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject
	this.SetComponentInterest(new(ScriptComponent))
}

func (this *ScriptSystem) Begin() {

}

func (this *ScriptSystem) End() {

}

// Generally, only the update method in source is called
func (this *ScriptSystem) ProcessObject(object *Object) {
	script := object.Component(new(ScriptComponent)).(*ScriptComponent)
	// For now we'll create a new runtime for each script
	if script.runtime == nil {
		// If we haven't ran this script earlier we'll create a new runtime for it and set up its interface
		script.runtime = otto.New()
		this.api.InitializeRuntime(script.runtime, object)
		script.runtime.Run(script.Src)

		// Call initialize (we assume this is the first time it's called)
		_, err := script.runtime.Call("Initialize", nil)
		if err != nil {
			gLogger.Fatalln("Otto runtime error:", err)
		}
	}

	_, err := script.runtime.Call("Update", nil)
	if err != nil {
		gLogger.Fatalln("Otto runtime error:", err)
	}
}

package n3

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

func (scriptsystem *ScriptSystem) Initialize() {
	scriptsystem.ProcessFunc = scriptsystem.ProcessObject
	scriptsystem.AddComponent(new(ScriptComponent))
}

func (scriptsystem *ScriptSystem) Begin() {

}

func (scriptsystem *ScriptSystem) End() {

}

// Generally, only the update method in source is called
func (scriptsystem *ScriptSystem) ProcessObject(object *Object) {
	script := object.Component(new(ScriptComponent)).(*ScriptComponent)
	// For now we'll create a new runtime for each script
	if script.runtime == nil {
		// If we haven't ran scriptsystem script earlier we'll create a new runtime for it and set up its interface
		script.runtime = otto.New()
		scriptsystem.api.InitializeRuntime(script.runtime, object)
		script.runtime.Run(script.Src)

		// Call initialize (we assume scriptsystem is the first time it's called)
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

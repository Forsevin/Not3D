package n3

import "github.com/robertkrimen/otto"

// ScriptSystem represents a system for scripting.
type ScriptSystem struct {
	System
	api *API
}

// NewScriptSystem returns a new ScriptSytem with the API set.
func NewScriptSystem(api *API) *ScriptSystem {
	return &ScriptSystem{
		api: api,
	}
}

// Initialize a new ScriptSystem
func (s *ScriptSystem) Initialize() {
	s.ProcessFunc = s.ProcessObject
	s.AddComponent(new(ScriptComponent))
}

// Begin the batch process
func (s *ScriptSystem) Begin() {
	// TODO: implement
}

// End the batch process
func (s *ScriptSystem) End() {
	// TODO: implement
}

// ProcessObject takes an object and tries to run it in an otto context.
// If the context doesn't exist, it'll be created for now.
// Generally, only the update method in source is called
func (s *ScriptSystem) ProcessObject(object *Object) {
	script := object.Component(new(ScriptComponent)).(*ScriptComponent)
	// For now we'll create a new runtime for each script
	if script.runtime == nil {
		// If we haven't ran scriptsystem script earlier we'll create a new runtime for it and set up its interface
		script.runtime = otto.New()
		s.api.MustInitializeRuntime(script.runtime, object)

		// TODO: check to see if we can run the script
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

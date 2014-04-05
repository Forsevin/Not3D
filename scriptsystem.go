package oden

import "github.com/robertkrimen/otto"
import "os"

type ScriptSystem struct {
	System
	runtime *otto.Otto
}

func NewScriptSystem() *ScriptSystem {
	return &ScriptSystem{
		runtime: otto.New(),
	}
}

func (this *ScriptSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject
	this.SetComponentInterest(new(ScriptComponent))

	this.runtime.Set("quit", func(call otto.FunctionCall) otto.Value {
		os.Exit(0)
		return otto.UndefinedValue()
	})

}

// Generally, only the update method in source is called
func (this *ScriptSystem) ProcessObject(object *Object) {
	script := object.Component(new(ScriptComponent)).(*ScriptComponent)
	this.runtime.Run(script.Src)
}

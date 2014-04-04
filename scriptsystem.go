package oden

import (
	"fmt"
	//"github.com/robertkrimen/otto"
)

type ScriptSystem struct {
	System
	//runtime *otto.Otto
}

func NewScriptSystem() *ScriptSystem {
	return &ScriptSystem{
	//runtime: otto.New(),
	}
}

func (this *ScriptSystem) Initialize() {
	this.ProcessFunc = this.ProcessObject
	this.SetDataInterest(gDataManager.Get(new(ScriptData)))
}

// Generally, only the update method in source is called
func (this *ScriptSystem) ProcessObject(object *Object) {
	scriptData := object.DataByType(new(ScriptData)).(*ScriptData)

	// Call the Update() method
	fmt.Println("Ran", scriptData.scripts)
}

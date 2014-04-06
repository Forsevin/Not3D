package oden

import "github.com/robertkrimen/otto"

// Interface for Otto

type Api struct {
	base *Base
}

// The api will need a lot of variables
func NewApi(base *Base) *Api {
	return &Api{
		base: base,
	}
}

func (this *Api) InitializeRuntime(runtime *otto.Otto) {
	runtime.Set("Quit", this.QuitGame)
	runtime.Set("SwitchScene", this.SetScene)
	runtime.Log("Log", this.Log)
}

// Switch scene
func (this *Api) SetScene(call otto.FunctionCall) otto.Value {
	scene, _ := call.Argument(0).ToString()

	gLogger.Printf("Switch scene", scene)

	return otto.NullValue()
}

// Get scene
func (this *Api) Scene(call otto.FunctionCall) otto.Value {
	return otto.NullValue()
}

// Quit the game
func (this *Api) QuitGame(call otto.FunctionCall) otto.Value {
	this.base.quit = true
	return otto.NullValue()
}

// Check if a key is down
func (this *Api) KeyDown(call otto.FunctionCall) otto.Value {
	return otto.NullValue()
}

// Check if a key is up
func (this *Api) KeyUp(call otto.FunctionCall) otto.Value {
	return otto.NullValue()
}

// Delete an object by name
func (this *Api) DeleteObject(call otto.FunctionCall) otto.Value {
	return otto.NullValue()
}

func (this *Api) Log(call otto.FunctionCall) otto.Value {
	msg, _ := call.Argument(0).ToString()
	gLogger.Print(msg)
}

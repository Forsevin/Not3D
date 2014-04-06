package oden

import "github.com/robertkrimen/otto"

// Interface for Otto

type Api struct {
	base  *Base
	input *Input
}

func NewApi(base *Base) *Api {
	return &Api{
		base: base,
	}
}

// Inititalize a runtime, setting classes avaible in scripting etc
func (this *Api) InitializeRuntime(runtime *otto.Otto, object *Object) {
	runtime.Set("engine", NewEngineInterface(this.base))
	runtime.Set("object", NewObjectInterface(object))
}

//// ENGINE INTERFACE
// General engine functions
type EngineInterface struct {
	base *Base
}

func NewEngineInterface(base *Base) *EngineInterface {
	return &EngineInterface{
		base: base,
	}
}

// Set base.quit to true
func (this *EngineInterface) Quit(call otto.FunctionCall) otto.Value {
	this.base.SetQuit(true)

	return otto.NullValue()
}

// Sets the active scene
func (this *EngineInterface) SetActiveScene(call otto.FunctionCall) otto.Value {
	scene, err := call.Argument(0).ToString()
	if err != nil {
		gLogger.Fatalln(err)
	}

	// TODO check if scene exists first
	this.base.SetActiveScene(scene)

	return otto.NullValue()
}

//// TRANSFORM INTERFACE
// Used to manipulate objects, for ObjectInterface

//// OBJECT INTERFACE
// Interface for objects (If you write a new component you want avaible for scripting you need to add a function for it here)
type ObjectInterface struct {
	object *Object
}

func NewObjectInterface(object *Object) *ObjectInterface {
	return &ObjectInterface{
		object: object,
	}
}

func (this *ObjectInterface) SetX(call otto.FunctionCall) otto.Value {
	x, err := call.Argument(0).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}
	transform := this.object.Component(new(TransformComponent)).(*TransformComponent)
	transform.X = int32(x)

	return otto.NullValue()
}

func (this *ObjectInterface) GetX(call otto.FunctionCall) otto.Value {
	transform := this.object.Component(new(TransformComponent)).(*TransformComponent)
	// This object doesn't have a transform component
	if transform == nil {
		return otto.NullValue()
	}

	r, _ := otto.ToValue(int64(transform.X))
	return r
}

func (this *ObjectInterface) SetY(call otto.FunctionCall) otto.Value {
	y, err := call.Argument(0).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}
	transform := this.object.Component(new(TransformComponent)).(*TransformComponent)
	transform.Y = int32(y)

	return otto.NullValue()
}

func (this *ObjectInterface) GetY(call otto.FunctionCall) otto.Value {
	transform := this.object.Component(new(TransformComponent)).(*TransformComponent)
	// This object doesn't have a transform component
	if transform == nil {
		return otto.NullValue()
	}

	r, _ := otto.ToValue(int64(transform.Y))
	return r
}

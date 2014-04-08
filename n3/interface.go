package n3

import "github.com/robertkrimen/otto"

// Interface for Otto

type Api struct {
	base *Base
}

func NewApi(base *Base) *Api {
	return &Api{
		base: base,
	}
}

// Inititalize a runtime, setting classes avaible in scripting etc
func (api *Api) InitializeRuntime(runtime *otto.Otto, object *Object) {
	runtime.Set("engine", NewEngineInterface(api.base))
	runtime.Set("object", NewObjectInterface(object))
	runtime.Set("input", NewInputInterface(api.base.Input()))
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
func (api *EngineInterface) Quit(call otto.FunctionCall) otto.Value {
	api.base.SetQuit(true)

	return otto.NullValue()
}

// Spawn a object
// @arg1 name of prefab
// @arg2 x cordinates
// @arg3 y cordinates
func (api *EngineInterface) SpawnPrefab(call otto.FunctionCall) otto.Value {
	prefabName, err := call.Argument(0).ToString()
	if err != nil {
		gLogger.Fatalln(err)
	}

	x, err := call.Argument(1).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}

	y, err := call.Argument(2).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}

	prefab := api.base.Prefabs().Prefab(prefabName)
	if prefab == nil {
		gLogger.Fatalln("Prefab", prefabName, "doesn't exit")
	}
	// api is a bit dirty don't you think?
	cords := prefab.Component(new(TransformComponent)).(*TransformComponent)
	cords.X = int32(x)
	cords.Y = int32(y)

	return otto.NullValue()
}

// Print something
func (api *EngineInterface) Print(call otto.FunctionCall) otto.Value {
	msg, _ := call.Argument(0).ToString()
	gLogger.Println(msg)
	return otto.NullValue()
}

// Sets the active scene
func (api *EngineInterface) SetActiveScene(call otto.FunctionCall) otto.Value {
	scene, err := call.Argument(0).ToString()
	if err != nil {
		gLogger.Fatalln(err)
	}

	// TODO check if scene exists first
	api.base.SetActiveScene(scene)

	return otto.NullValue()
}

//// INPUT INTERFACE
// For getting input
type InputInterface struct {
	input *Input
}

func NewInputInterface(input *Input) *InputInterface {
	return &InputInterface{
		input: input,
	}
}

func (inputInterface *InputInterface) KeyDown(call otto.FunctionCall) otto.Value {
	key, err := call.Argument(0).ToString()
	if err != nil {
		gLogger.Fatalln(err)
	}
	r, err := otto.ToValue(inputInterface.input.KeyDown(key))
	if err != nil {
		gLogger.Fatalln(err)
	}
	return r

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

func (objectInterface *ObjectInterface) SetX(call otto.FunctionCall) otto.Value {
	x, err := call.Argument(0).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}
	transform := objectInterface.object.Component(new(TransformComponent)).(*TransformComponent)
	transform.X = int32(x)

	return otto.NullValue()
}

func (objectInterface *ObjectInterface) GetX(call otto.FunctionCall) otto.Value {
	transform := objectInterface.object.Component(new(TransformComponent)).(*TransformComponent)
	// api object doesn't have a transform component
	if transform == nil {
		return otto.NullValue()
	}

	r, _ := otto.ToValue(int64(transform.X))
	return r
}

func (objectInterface *ObjectInterface) SetY(call otto.FunctionCall) otto.Value {
	y, err := call.Argument(0).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}
	transform := objectInterface.object.Component(new(TransformComponent)).(*TransformComponent)
	transform.Y = int32(y)

	return otto.NullValue()
}

func (objectInterface *ObjectInterface) GetY(call otto.FunctionCall) otto.Value {
	transform := objectInterface.object.Component(new(TransformComponent)).(*TransformComponent)

	if transform == nil {
		return otto.NullValue()
	}

	r, _ := otto.ToValue(int64(transform.Y))
	return r
}

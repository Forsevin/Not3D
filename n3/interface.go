package n3

import "github.com/robertkrimen/otto"

// Interface for Otto

// API ...
type API struct {
	base *Base
}

// NewAPI returns an API with the base set
func NewAPI(base *Base) *API {
	return &API{
		base: base,
	}
}

// InitializeRuntime initializes a runtime, setting classes avaible in scripting etc
func (api *API) InitializeRuntime(runtime *otto.Otto, object *Object) {
	// BUG(j6n): check and return these errors
	runtime.Set("engine", NewEngineInterface(api.base))
	runtime.Set("object", NewObjectInterface(object))
	runtime.Set("input", NewInputInterface(api.base.Input()))
}

// EngineInterface contains general engine functions
type EngineInterface struct {
	base *Base
}

// NewEngineInterface creates a new EngineInterface with the base set
func NewEngineInterface(base *Base) *EngineInterface {
	return &EngineInterface{
		base: base,
	}
}

// Quit sets base.quit to true
func (api *EngineInterface) Quit(call otto.FunctionCall) otto.Value {
	api.base.SetQuit(true)

	return otto.NullValue()
}

// SpawnPrefab is an otto function that takes three arguments, a string and two integer. This spawns an object.
// The string is the name of the prefab, while the int arguments are X and Y coords, respectively.
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

// Print is an otto function that takes a javascript string and prints it
func (api *EngineInterface) Print(call otto.FunctionCall) otto.Value {
	msg, _ := call.Argument(0).ToString()
	gLogger.Println(msg)
	return otto.NullValue()
}

// SetActiveScene is an otto function that takes a javascript string and sets the active scene to it
func (api *EngineInterface) SetActiveScene(call otto.FunctionCall) otto.Value {
	scene, err := call.Argument(0).ToString()
	if err != nil {
		gLogger.Fatalln(err)
	}

	// TODO check if scene exists first
	api.base.SetActiveScene(scene)

	return otto.NullValue()
}

// InputInterface is the interface to the input system
type InputInterface struct {
	input *Input
}

// NewInputInterface returns a new InputInterface with the input system set.
func NewInputInterface(input *Input) *InputInterface {
	return &InputInterface{
		input: input,
	}
}

// KeyDown is an otto function that takes a string, which is the key.
// It then tries to set the keydown state in the Input system, returning the bool result.
func (i *InputInterface) KeyDown(call otto.FunctionCall) otto.Value {
	key, err := call.Argument(0).ToString()
	if err != nil {
		gLogger.Fatalln(err)
	}
	r, err := otto.ToValue(i.input.KeyDown(key))
	if err != nil {
		gLogger.Fatalln(err)
	}

	return r
}

//// TRANSFORM INTERFACE
// Used to manipulate objects, for ObjectInterface

// ObjectInterface is an interface for objects.
// If you write a new component you want avaible for scripting you need to add a function for it here
type ObjectInterface struct {
	object *Object
}

// NewObjectInterface returns a new ObjectInterface based on object
func NewObjectInterface(object *Object) *ObjectInterface {
	return &ObjectInterface{
		object: object,
	}
}

// SetX is an otto function that takes an integer and sets the object's X coord
func (o *ObjectInterface) SetX(call otto.FunctionCall) otto.Value {
	x, err := call.Argument(0).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}
	transform := o.object.Component(new(TransformComponent)).(*TransformComponent)
	transform.X = int32(x)

	return otto.NullValue()
}

// GetX is an otto function that returns the object's X coord
func (o *ObjectInterface) GetX(call otto.FunctionCall) otto.Value {
	transform := o.object.Component(new(TransformComponent)).(*TransformComponent)
	// api object doesn't have a transform component
	if transform == nil {
		return otto.NullValue()
	}

	r, _ := otto.ToValue(int64(transform.X))
	return r
}

// SetY is an otto function that takes an integer and sets the object's Y coord
func (o *ObjectInterface) SetY(call otto.FunctionCall) otto.Value {
	y, err := call.Argument(0).ToInteger()
	if err != nil {
		gLogger.Fatalln(err)
	}
	transform := o.object.Component(new(TransformComponent)).(*TransformComponent)
	transform.Y = int32(y)

	return otto.NullValue()
}

// GetY is an otto function that returns the object's Y coord
func (o *ObjectInterface) GetY(call otto.FunctionCall) otto.Value {
	transform := o.object.Component(new(TransformComponent)).(*TransformComponent)

	if transform == nil {
		return otto.NullValue()
	}

	r, _ := otto.ToValue(int64(transform.Y))
	return r
}

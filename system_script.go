package n3

type Script struct {
	entity       *Entity
	core         *Core
	init         bool
	enabled      bool
	onMouseDown  func()
	onMouseLeave func()
}

func (s *Script) getMouseDown() *func() {
	return &s.onMouseDown
}

func (s *Script) getInit() bool {
	return s.init
}

func (s *Script) setInit(v bool) {
	s.init = v
}

func (s *Script) Entity() *Entity {
	return s.entity
}

func (s *Script) Disable() {
	s.enabled = false
}

func (s *Script) Enable() {
	s.enabled = true
}

func (s *Script) Engine() *Core {
	return s.core
}

func (s *Script) setEntity(entity *Entity) {
	s.entity = entity
}

func (s *Script) setEngine(engine *Core) {
	s.core = engine
}

// holds an collection of entities
type ScriptComponent struct {
	scripts []scriptInterface
}

func NewScriptComponent(scripts ...scriptInterface) *ScriptComponent {
	s := &ScriptComponent{}

	for _, script := range scripts {
		script.setInit(true)
		s.scripts = append(s.scripts, script)
	}

	return s
}

type scriptInterface interface {
	Init()
	Update()
	Entity() *Entity
	Engine() *Core
	setEntity(entity *Entity)
	setEngine(engine *Core)
	setInit(v bool)
	getInit() bool
	Enable()
	Disable()
}

type ScriptSystem struct {
	EntitySystem
}

func NewScriptSystem() *ScriptSystem {
	return &ScriptSystem{}
}

func (s *ScriptSystem) Init() {
	s.Component(new(ScriptComponent))
}

func (s *ScriptSystem) Begin() {
}

func (s *ScriptSystem) Process() {
	for _, entity := range s.entities {
		scripts := entity.GetComponent(new(ScriptComponent)).(*ScriptComponent)
		for _, script := range scripts.scripts {
			script.setEntity(entity)
		}
		for _, script := range scripts.scripts {
			if script.getInit() == true {
				script.Init()
				script.setEngine(s.Core())
				script.setInit(false)
			}
			script.Update()
		}
	}
}

func (s *ScriptSystem) End() {
}

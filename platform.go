/* Platform.go
 * Platform specific events and handlers e.g for windows, rendering and input
 * example:
 *
 * if platform.Input.Get("w") {
 *	fmt.Println("W is down")
 * }
 */
package n3

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/sdl_ttf"

// basic colors
var (
	COLOR_RED   = sdl.Color{255, 0, 0, 0}
	COLOR_GREEN = sdl.Color{0, 255, 0, 0}
	COLOR_BLUE  = sdl.Color{0, 0, 255, 0}
)

// wraps around the platform structures
type Platform struct {
	input    *Input
	window   *Window
	renderer *Renderer
}

// NewPlatform returns a new platform
func NewPlatform() *Platform {
	p := &Platform{
		input: &Input{
			keyStates: make(map[string]bool),
		},
		window: &Window{
			window: sdl.CreateWindow("Application", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE),
		},
	}
	p.renderer = &Renderer{
		renderer: sdl.CreateRenderer(p.window.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC),
	}

	ttf.Init()

	return p
}

// Handle window
type Window struct {
	window *sdl.Window
}

// Set window to fullscreen
func (w *Window) FullScreen() {
	w.window.SetFullscreen(1)
}

// Wrapper around sdl.Renderer
type Renderer struct {
	renderer      *sdl.Renderer
	backgroundTex *sdl.Texture
}

func (r *Renderer) SetBackground(tex *sdl.Texture) {
	r.backgroundTex = tex
}

// Input handles key presses
type Input struct {
	// Save states for all keys here
	keyStates map[string]bool
}

// NewInput returns a new input
func NewInput() *Input {
	return &Input{
		keyStates: make(map[string]bool),
	}
}

// Process the input,
func (input *Input) Process() bool {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.KeyDownEvent:
			if key, ok := sdlKeyMap[t.Keysym.Sym]; ok {
				input.SetKeyDown(key)
			}
		case *sdl.KeyUpEvent:
			if key, ok := sdlKeyMap[t.Keysym.Sym]; ok {
				input.SetKeyUp(key)
			}
		}
	}

	return false
}

// KeyDown returns whether the key is in the down state
func (input *Input) KeyDown(key string) bool {
	return input.keyStates[key]
}

func (input *Input) PointerToKey(key string) *bool {
	return nil
}

// SetKeyDown sets the key state to down
func (input *Input) SetKeyDown(key string) {
	input.keyStates[key] = true
}

func (input *Input) GetMouseCords() (x, y int32) {
	return 0, 0
}

// SetKeyUp sets the key state to up
func (input *Input) SetKeyUp(key string) {
	input.keyStates[key] = false
}

// this is a mapping of SDL keys to strings
var sdlKeyMap = map[sdl.Keycode]string{
	sdl.K_BACKSPACE:    "BACKSPACE",
	sdl.K_TAB:          "TAB",
	sdl.K_CLEAR:        "CLEAR",
	sdl.K_RETURN:       "RETURN",
	sdl.K_PAUSE:        "PAUSE",
	sdl.K_ESCAPE:       "ESCAPE",
	sdl.K_SPACE:        "SPACE",
	sdl.K_EXCLAIM:      "EXCLAIM",
	sdl.K_QUOTEDBL:     "QUOTEDBL",
	sdl.K_HASH:         "HASH",
	sdl.K_DOLLAR:       "DOLLAR",
	sdl.K_AMPERSAND:    "AMPERSAND",
	sdl.K_QUOTE:        "QUOTE",
	sdl.K_LEFTPAREN:    "LEFTPAREN",
	sdl.K_RIGHTPAREN:   "RIGHTPAREN",
	sdl.K_ASTERISK:     "ASTERISK",
	sdl.K_PLUS:         "PLUS",
	sdl.K_COMMA:        "COMMA",
	sdl.K_MINUS:        "MINUS",
	sdl.K_PERIOD:       "PERIOD",
	sdl.K_SLASH:        "SLASH",
	sdl.K_0:            "0",
	sdl.K_1:            "1",
	sdl.K_2:            "2",
	sdl.K_3:            "3",
	sdl.K_4:            "4",
	sdl.K_5:            "5",
	sdl.K_6:            "6",
	sdl.K_7:            "7",
	sdl.K_8:            "8",
	sdl.K_9:            "9",
	sdl.K_COLON:        "COLON",
	sdl.K_SEMICOLON:    "SEMICOLON",
	sdl.K_LESS:         "LESS",
	sdl.K_EQUALS:       "EQUALS",
	sdl.K_GREATER:      "GREATER",
	sdl.K_QUESTION:     "QUESTION",
	sdl.K_AT:           "AT",
	sdl.K_LEFTBRACKET:  "LEFTBRACKET",
	sdl.K_BACKSLASH:    "BACKSLASH",
	sdl.K_RIGHTBRACKET: "RIGHTBRACKET",
	sdl.K_CARET:        "CARET",
	sdl.K_UNDERSCORE:   "UNDERSCORE",
	sdl.K_BACKQUOTE:    "BACKQUOTE",
	sdl.K_a:            "a",
	sdl.K_b:            "b",
	sdl.K_c:            "c",
	sdl.K_d:            "d",
	sdl.K_e:            "e",
	sdl.K_f:            "f",
	sdl.K_g:            "g",
	sdl.K_h:            "h",
	sdl.K_i:            "i",
	sdl.K_j:            "j",
	sdl.K_k:            "k",
	sdl.K_l:            "l",
	sdl.K_m:            "m",
	sdl.K_n:            "n",
	sdl.K_o:            "o",
	sdl.K_p:            "p",
	sdl.K_q:            "q",
	sdl.K_r:            "r",
	sdl.K_s:            "s",
	sdl.K_t:            "t",
	sdl.K_u:            "u",
	sdl.K_v:            "v",
	sdl.K_w:            "w",
	sdl.K_x:            "x",
	sdl.K_y:            "y",
	sdl.K_z:            "z",
	sdl.K_DELETE:       "DELETE",
	sdl.K_KP_PERIOD:    "KP_PERIOD",
	sdl.K_KP_DIVIDE:    "KP_DIVIDE",
	sdl.K_KP_MULTIPLY:  "KP_MULTIPLY",
	sdl.K_KP_MINUS:     "KP_MINUS",
	sdl.K_KP_PLUS:      "KP_PLUS",
	sdl.K_KP_ENTER:     "KP_ENTER",
	sdl.K_KP_EQUALS:    "KP_EQUALS",
	sdl.K_UP:           "UP",
	sdl.K_DOWN:         "DOWN",
	sdl.K_RIGHT:        "RIGHT",
	sdl.K_LEFT:         "LEFT",
	sdl.K_INSERT:       "INSERT",
	sdl.K_HOME:         "HOME",
	sdl.K_END:          "END",
	sdl.K_PAGEUP:       "PAGEUP",
	sdl.K_PAGEDOWN:     "PAGEDOWN",
	sdl.K_F1:           "F1",
	sdl.K_F2:           "F2",
	sdl.K_F3:           "F3",
	sdl.K_F4:           "F4",
	sdl.K_F5:           "F5",
	sdl.K_F6:           "F6",
	sdl.K_F7:           "F7",
	sdl.K_F8:           "F8",
	sdl.K_F9:           "F9",
	sdl.K_F10:          "F10",
	sdl.K_F11:          "F11",
	sdl.K_F12:          "F12",
	sdl.K_F13:          "F13",
	sdl.K_F14:          "F14",
	sdl.K_F15:          "F15",
	sdl.K_CAPSLOCK:     "CAPSLOCK",
	sdl.K_RSHIFT:       "RSHIFT",
	sdl.K_LSHIFT:       "LSHIFT",
	sdl.K_RCTRL:        "RCTRL",
	sdl.K_LCTRL:        "LCTRL",
	sdl.K_RALT:         "RALT",
	sdl.K_LALT:         "LALT",
	sdl.K_MODE:         "MODE",
	sdl.K_HELP:         "HELP",
	sdl.K_SYSREQ:       "SYSREQ",
	sdl.K_MENU:         "MENU",
	sdl.K_POWER:        "POWER",
}

package n3

import "github.com/jackyb/go-sdl2/sdl"

// Handles input

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
			switch t.Keysym.Sym {
			case sdl.K_BACKSPACE:
				input.SetKeyDown("BACKSPACE")
			case sdl.K_TAB:
				input.SetKeyDown("TAB")
			case sdl.K_CLEAR:
				input.SetKeyDown("CLEAR")
			case sdl.K_RETURN:
				input.SetKeyDown("RETURN")
			case sdl.K_PAUSE:
				input.SetKeyDown("PAUSE")
			case sdl.K_ESCAPE:
				input.SetKeyDown("ESCAPE")
			case sdl.K_SPACE:
				input.SetKeyDown("SPACE")
			case sdl.K_EXCLAIM:
				input.SetKeyDown("EXCLAIM")
			case sdl.K_QUOTEDBL:
				input.SetKeyDown("QUOTEDBL")
			case sdl.K_HASH:
				input.SetKeyDown("HASH")
			case sdl.K_DOLLAR:
				input.SetKeyDown("DOLLAR")
			case sdl.K_AMPERSAND:
				input.SetKeyDown("AMPERSAND")
			case sdl.K_QUOTE:
				input.SetKeyDown("QUOTE")
			case sdl.K_LEFTPAREN:
				input.SetKeyDown("LEFTPAREN")
			case sdl.K_RIGHTPAREN:
				input.SetKeyDown("RIGHTPAREN")
			case sdl.K_ASTERISK:
				input.SetKeyDown("ASTERISK")
			case sdl.K_PLUS:
				input.SetKeyDown("PLUS")
			case sdl.K_COMMA:
				input.SetKeyDown("COMMA")
			case sdl.K_MINUS:
				input.SetKeyDown("MINUS")
			case sdl.K_PERIOD:
				input.SetKeyDown("PERIOD")
			case sdl.K_SLASH:
				input.SetKeyDown("SLASH")
			case sdl.K_0:
				input.SetKeyDown("0")
			case sdl.K_1:
				input.SetKeyDown("1")
			case sdl.K_2:
				input.SetKeyDown("2")
			case sdl.K_3:
				input.SetKeyDown("3")
			case sdl.K_4:
				input.SetKeyDown("4")
			case sdl.K_5:
				input.SetKeyDown("5")
			case sdl.K_6:
				input.SetKeyDown("6")
			case sdl.K_7:
				input.SetKeyDown("7")
			case sdl.K_8:
				input.SetKeyDown("8")
			case sdl.K_9:
				input.SetKeyDown("9")
			case sdl.K_COLON:
				input.SetKeyDown("COLON")
			case sdl.K_SEMICOLON:
				input.SetKeyDown("SEMICOLON")
			case sdl.K_LESS:
				input.SetKeyDown("LESS")
			case sdl.K_EQUALS:
				input.SetKeyDown("EQUALS")
			case sdl.K_GREATER:
				input.SetKeyDown("GREATER")
			case sdl.K_QUESTION:
				input.SetKeyDown("QUESTION")
			case sdl.K_AT:
				input.SetKeyDown("AT")
			case sdl.K_LEFTBRACKET:
				input.SetKeyDown("LEFTBRACKET")
			case sdl.K_BACKSLASH:
				input.SetKeyDown("BACKSLASH")
			case sdl.K_RIGHTBRACKET:
				input.SetKeyDown("RIGHTBRACKET")
			case sdl.K_CARET:
				input.SetKeyDown("CARET")
			case sdl.K_UNDERSCORE:
				input.SetKeyDown("UNDERSCORE")
			case sdl.K_BACKQUOTE:
				input.SetKeyDown("BACKQUOTE")
			case sdl.K_a:
				input.SetKeyDown("a")
			case sdl.K_b:
				input.SetKeyDown("b")
			case sdl.K_c:
				input.SetKeyDown("c")
			case sdl.K_d:
				input.SetKeyDown("d")
			case sdl.K_e:
				input.SetKeyDown("e")
			case sdl.K_f:
				input.SetKeyDown("f")
			case sdl.K_g:
				input.SetKeyDown("g")
			case sdl.K_h:
				input.SetKeyDown("h")
			case sdl.K_i:
				input.SetKeyDown("i")
			case sdl.K_j:
				input.SetKeyDown("j")
			case sdl.K_k:
				input.SetKeyDown("k")
			case sdl.K_l:
				input.SetKeyDown("l")
			case sdl.K_m:
				input.SetKeyDown("m")
			case sdl.K_n:
				input.SetKeyDown("n")
			case sdl.K_o:
				input.SetKeyDown("o")
			case sdl.K_p:
				input.SetKeyDown("p")
			case sdl.K_q:
				input.SetKeyDown("q")
			case sdl.K_r:
				input.SetKeyDown("r")
			case sdl.K_s:
				input.SetKeyDown("s")
			case sdl.K_t:
				input.SetKeyDown("t")
			case sdl.K_u:
				input.SetKeyDown("u")
			case sdl.K_v:
				input.SetKeyDown("v")
			case sdl.K_w:
				input.SetKeyDown("w")
			case sdl.K_x:
				input.SetKeyDown("x")
			case sdl.K_y:
				input.SetKeyDown("y")
			case sdl.K_z:
				input.SetKeyDown("z")
			case sdl.K_DELETE:
				input.SetKeyDown("DELETE")
			case sdl.K_KP_PERIOD:
				input.SetKeyDown("KP_PERIOD")
			case sdl.K_KP_DIVIDE:
				input.SetKeyDown("KP_DIVIDE")
			case sdl.K_KP_MULTIPLY:
				input.SetKeyDown("KP_MULTIPLY")
			case sdl.K_KP_MINUS:
				input.SetKeyDown("KP_MINUS")
			case sdl.K_KP_PLUS:
				input.SetKeyDown("KP_PLUS")
			case sdl.K_KP_ENTER:
				input.SetKeyDown("KP_ENTER")
			case sdl.K_KP_EQUALS:
				input.SetKeyDown("KP_EQUALS")
			case sdl.K_UP:
				input.SetKeyDown("UP")
			case sdl.K_DOWN:
				input.SetKeyDown("DOWN")
			case sdl.K_RIGHT:
				input.SetKeyDown("RIGHT")
			case sdl.K_LEFT:
				input.SetKeyDown("LEFT")
			case sdl.K_INSERT:
				input.SetKeyDown("INSERT")
			case sdl.K_HOME:
				input.SetKeyDown("HOME")
			case sdl.K_END:
				input.SetKeyDown("END")
			case sdl.K_PAGEUP:
				input.SetKeyDown("PAGEUP")
			case sdl.K_PAGEDOWN:
				input.SetKeyDown("PAGEDOWN")
			case sdl.K_F1:
				input.SetKeyDown("F1")
			case sdl.K_F2:
				input.SetKeyDown("F2")
			case sdl.K_F3:
				input.SetKeyDown("F3")
			case sdl.K_F4:
				input.SetKeyDown("F4")
			case sdl.K_F5:
				input.SetKeyDown("F5")
			case sdl.K_F6:
				input.SetKeyDown("F6")
			case sdl.K_F7:
				input.SetKeyDown("F7")
			case sdl.K_F8:
				input.SetKeyDown("F8")
			case sdl.K_F9:
				input.SetKeyDown("F9")
			case sdl.K_F10:
				input.SetKeyDown("F10")
			case sdl.K_F11:
				input.SetKeyDown("F11")
			case sdl.K_F12:
				input.SetKeyDown("F12")
			case sdl.K_F13:
				input.SetKeyDown("F13")
			case sdl.K_F14:
				input.SetKeyDown("F14")
			case sdl.K_F15:
				input.SetKeyDown("F15")
			case sdl.K_CAPSLOCK:
				input.SetKeyDown("CAPSLOCK")
			case sdl.K_RSHIFT:
				input.SetKeyDown("RSHIFT")
			case sdl.K_LSHIFT:
				input.SetKeyDown("LSHIFT")
			case sdl.K_RCTRL:
				input.SetKeyDown("RCTRL")
			case sdl.K_LCTRL:
				input.SetKeyDown("LCTRL")
			case sdl.K_RALT:
				input.SetKeyDown("RALT")
			case sdl.K_LALT:
				input.SetKeyDown("LALT")
			case sdl.K_MODE:
				input.SetKeyDown("MODE")
			case sdl.K_HELP:
				input.SetKeyDown("HELP")
			case sdl.K_SYSREQ:
				input.SetKeyDown("SYSREQ")
			case sdl.K_MENU:
				input.SetKeyDown("MENU")
			case sdl.K_POWER:
				input.SetKeyDown("POWER")
			}
		case *sdl.KeyUpEvent:
			switch t.Keysym.Sym {
			case sdl.K_BACKSPACE:
				input.SetKeyUp("BACKSPACE")
			case sdl.K_TAB:
				input.SetKeyUp("TAB")
			case sdl.K_CLEAR:
				input.SetKeyUp("CLEAR")
			case sdl.K_RETURN:
				input.SetKeyUp("RETURN")
			case sdl.K_PAUSE:
				input.SetKeyUp("PAUSE")
			case sdl.K_ESCAPE:
				input.SetKeyUp("ESCAPE")
			case sdl.K_SPACE:
				input.SetKeyUp("SPACE")
			case sdl.K_EXCLAIM:
				input.SetKeyUp("EXCLAIM")
			case sdl.K_QUOTEDBL:
				input.SetKeyUp("QUOTEDBL")
			case sdl.K_HASH:
				input.SetKeyUp("HASH")
			case sdl.K_DOLLAR:
				input.SetKeyUp("DOLLAR")
			case sdl.K_AMPERSAND:
				input.SetKeyUp("AMPERSAND")
			case sdl.K_QUOTE:
				input.SetKeyUp("QUOTE")
			case sdl.K_LEFTPAREN:
				input.SetKeyUp("LEFTPAREN")
			case sdl.K_RIGHTPAREN:
				input.SetKeyUp("RIGHTPAREN")
			case sdl.K_ASTERISK:
				input.SetKeyUp("ASTERISK")
			case sdl.K_PLUS:
				input.SetKeyUp("PLUS")
			case sdl.K_COMMA:
				input.SetKeyUp("COMMA")
			case sdl.K_MINUS:
				input.SetKeyUp("MINUS")
			case sdl.K_PERIOD:
				input.SetKeyUp("PERIOD")
			case sdl.K_SLASH:
				input.SetKeyUp("SLASH")
			case sdl.K_0:
				input.SetKeyUp("0")
			case sdl.K_1:
				input.SetKeyUp("1")
			case sdl.K_2:
				input.SetKeyUp("2")
			case sdl.K_3:
				input.SetKeyUp("3")
			case sdl.K_4:
				input.SetKeyUp("4")
			case sdl.K_5:
				input.SetKeyUp("5")
			case sdl.K_6:
				input.SetKeyUp("6")
			case sdl.K_7:
				input.SetKeyUp("7")
			case sdl.K_8:
				input.SetKeyUp("8")
			case sdl.K_9:
				input.SetKeyUp("9")
			case sdl.K_COLON:
				input.SetKeyUp("COLON")
			case sdl.K_SEMICOLON:
				input.SetKeyUp("SEMICOLON")
			case sdl.K_LESS:
				input.SetKeyUp("LESS")
			case sdl.K_EQUALS:
				input.SetKeyUp("EQUALS")
			case sdl.K_GREATER:
				input.SetKeyUp("GREATER")
			case sdl.K_QUESTION:
				input.SetKeyUp("QUESTION")
			case sdl.K_AT:
				input.SetKeyUp("AT")
			case sdl.K_LEFTBRACKET:
				input.SetKeyUp("LEFTBRACKET")
			case sdl.K_BACKSLASH:
				input.SetKeyUp("BACKSLASH")
			case sdl.K_RIGHTBRACKET:
				input.SetKeyUp("RIGHTBRACKET")
			case sdl.K_CARET:
				input.SetKeyUp("CARET")
			case sdl.K_UNDERSCORE:
				input.SetKeyUp("UNDERSCORE")
			case sdl.K_BACKQUOTE:
				input.SetKeyUp("BACKQUOTE")
			case sdl.K_a:
				input.SetKeyUp("a")
			case sdl.K_b:
				input.SetKeyUp("b")
			case sdl.K_c:
				input.SetKeyUp("c")
			case sdl.K_d:
				input.SetKeyUp("d")
			case sdl.K_e:
				input.SetKeyUp("e")
			case sdl.K_f:
				input.SetKeyUp("f")
			case sdl.K_g:
				input.SetKeyUp("g")
			case sdl.K_h:
				input.SetKeyUp("h")
			case sdl.K_i:
				input.SetKeyUp("i")
			case sdl.K_j:
				input.SetKeyUp("j")
			case sdl.K_k:
				input.SetKeyUp("k")
			case sdl.K_l:
				input.SetKeyUp("l")
			case sdl.K_m:
				input.SetKeyUp("m")
			case sdl.K_n:
				input.SetKeyUp("n")
			case sdl.K_o:
				input.SetKeyUp("o")
			case sdl.K_p:
				input.SetKeyUp("p")
			case sdl.K_q:
				input.SetKeyUp("q")
			case sdl.K_r:
				input.SetKeyUp("r")
			case sdl.K_s:
				input.SetKeyUp("s")
			case sdl.K_t:
				input.SetKeyUp("t")
			case sdl.K_u:
				input.SetKeyUp("u")
			case sdl.K_v:
				input.SetKeyUp("v")
			case sdl.K_w:
				input.SetKeyUp("w")
			case sdl.K_x:
				input.SetKeyUp("x")
			case sdl.K_y:
				input.SetKeyUp("y")
			case sdl.K_z:
				input.SetKeyUp("z")
			case sdl.K_DELETE:
				input.SetKeyUp("DELETE")
			case sdl.K_KP_PERIOD:
				input.SetKeyUp("KP_PERIOD")
			case sdl.K_KP_DIVIDE:
				input.SetKeyUp("KP_DIVIDE")
			case sdl.K_KP_MULTIPLY:
				input.SetKeyUp("KP_MULTIPLY")
			case sdl.K_KP_MINUS:
				input.SetKeyUp("KP_MINUS")
			case sdl.K_KP_PLUS:
				input.SetKeyUp("KP_PLUS")
			case sdl.K_KP_ENTER:
				input.SetKeyUp("KP_ENTER")
			case sdl.K_KP_EQUALS:
				input.SetKeyUp("KP_EQUALS")
			case sdl.K_UP:
				input.SetKeyUp("UP")
			case sdl.K_DOWN:
				input.SetKeyUp("DOWN")
			case sdl.K_RIGHT:
				input.SetKeyUp("RIGHT")
			case sdl.K_LEFT:
				input.SetKeyUp("LEFT")
			case sdl.K_INSERT:
				input.SetKeyUp("INSERT")
			case sdl.K_HOME:
				input.SetKeyUp("HOME")
			case sdl.K_END:
				input.SetKeyUp("END")
			case sdl.K_PAGEUP:
				input.SetKeyUp("PAGEUP")
			case sdl.K_PAGEDOWN:
				input.SetKeyUp("PAGEDOWN")
			case sdl.K_F1:
				input.SetKeyUp("F1")
			case sdl.K_F2:
				input.SetKeyUp("F2")
			case sdl.K_F3:
				input.SetKeyUp("F3")
			case sdl.K_F4:
				input.SetKeyUp("F4")
			case sdl.K_F5:
				input.SetKeyUp("F5")
			case sdl.K_F6:
				input.SetKeyUp("F6")
			case sdl.K_F7:
				input.SetKeyUp("F7")
			case sdl.K_F8:
				input.SetKeyUp("F8")
			case sdl.K_F9:
				input.SetKeyUp("F9")
			case sdl.K_F10:
				input.SetKeyUp("F10")
			case sdl.K_F11:
				input.SetKeyUp("F11")
			case sdl.K_F12:
				input.SetKeyUp("F12")
			case sdl.K_F13:
				input.SetKeyUp("F13")
			case sdl.K_F14:
				input.SetKeyUp("F14")
			case sdl.K_F15:
				input.SetKeyUp("F15")
			case sdl.K_CAPSLOCK:
				input.SetKeyUp("CAPSLOCK")
			case sdl.K_RSHIFT:
				input.SetKeyUp("RSHIFT")
			case sdl.K_LSHIFT:
				input.SetKeyUp("LSHIFT")
			case sdl.K_RCTRL:
				input.SetKeyUp("RCTRL")
			case sdl.K_LCTRL:
				input.SetKeyUp("LCTRL")
			case sdl.K_RALT:
				input.SetKeyUp("RALT")
			case sdl.K_LALT:
				input.SetKeyUp("LALT")
			case sdl.K_MODE:
				input.SetKeyUp("MODE")
			case sdl.K_HELP:
				input.SetKeyUp("HELP")
			case sdl.K_SYSREQ:
				input.SetKeyUp("SYSREQ")
			case sdl.K_MENU:
				input.SetKeyUp("MENU")
			case sdl.K_POWER:
				input.SetKeyUp("POWER")
			}
		}

	}

	return false
}

// KeyDown returns whether the key is in the down state
func (input *Input) KeyDown(key string) bool {
	return input.keyStates[key]
}

// SetKeyDown sets the key state to down
func (input *Input) SetKeyDown(key string) {
	input.keyStates[key] = true
}

// SetKeyUp sets the key state to up
func (input *Input) SetKeyUp(key string) {
	input.keyStates[key] = false
}

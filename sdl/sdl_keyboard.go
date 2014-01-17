package sdl

// #include <SDL2/SDL_keyboard.h>
import "C"
import "unsafe"
import "reflect"

type Keysym struct {
	Scancode Scancode
	Sym Keycode
	Mod uint16
	Unicode uint32
}

func GetKeyboardFocus() *Window {
	return (*Window) (unsafe.Pointer(C.SDL_GetKeyboardFocus()))
}

func GetKeyboardState() []uint8 {
	var numkeys C.int
	start := C.SDL_GetKeyboardState(&numkeys)
	sh := reflect.SliceHeader{}
	sh.Len = int(numkeys)
	sh.Cap = int(numkeys)
	sh.Data = uintptr(unsafe.Pointer(start))
	return *(*[]uint8)(unsafe.Pointer(&sh))
}

func SetModState(modstate Keymod) {
	_modstate := (C.SDL_Keymod) (modstate)
	C.SDL_SetModState(_modstate)
}

func GetKeyFromScancode(scancode Scancode) Keycode {
	_scancode := (C.SDL_Scancode) (scancode)
	return (Keycode) (C.SDL_GetKeyFromScancode(_scancode))
}

func GetScancodeFromKey(key Keycode) Scancode {
	_key := (C.SDL_Keycode) (key)
	return (Scancode) (C.SDL_GetScancodeFromKey(_key))
}

func GetScancodeName(scancode Scancode) string {
	_scancode := (C.SDL_Scancode) (scancode)
	return (C.GoString) (C.SDL_GetScancodeName(_scancode))
}

func GetScancodeFromName(name string) Scancode {
	_name := (C.CString) (name)
	return (Scancode) (C.SDL_GetScancodeFromName(_name))
}

func GetKeyName(key Keycode) string {
	_key := (C.SDL_Keycode) (key)
	return (C.GoString) (C.SDL_GetKeyName(_key))
}

func GetKeyFromName(name string) Keycode {
	_name := (C.CString) (name)
	return (Keycode) (C.SDL_GetKeyFromName(_name))
}

func StartTextInput() {
	C.SDL_StartTextInput()
}

func IsTextInputActive() bool {
	return C.SDL_IsTextInputActive() > 0
}

func StopTextInput() {
	C.SDL_StopTextInput()
}

func SetTextInputRect(rect *Rect) {
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	C.SDL_SetTextInputRect(_rect)
}

func HasScreenKeyboardSupport() bool {
	return C.SDL_HasScreenKeyboardSupport() > 0
}

func IsScreenKeyboardShown(window *Window) bool {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return C.SDL_IsScreenKeyboardShown(_window) > 0
}

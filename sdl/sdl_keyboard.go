package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"
import "reflect"

type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      uint16
	Unicode  uint32
}

func GetKeyboardFocus() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetKeyboardFocus()))
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

func SetModState(mod Keymod) {
	C.SDL_SetModState(mod.c())
}

func GetKeyFromScancode(code Scancode) Keycode {
	return (Keycode)(C.SDL_GetKeyFromScancode(code.c()))
}

func GetScancodeFromKey(code Keycode) Scancode {
	return (Scancode)(C.SDL_GetScancodeFromKey(code.c()))
}

func GetScancodeName(code Scancode) string {
	return (C.GoString)(C.SDL_GetScancodeName(code.c()))
}

func GetScancodeFromName(name string) Scancode {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (Scancode)(C.SDL_GetScancodeFromName(_name))
}

func GetKeyName(code Keycode) string {
	return (C.GoString)(C.SDL_GetKeyName(code.c()))
}

func GetKeyFromName(name string) Keycode {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (Keycode)(C.SDL_GetKeyFromName(_name))
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
	C.SDL_SetTextInputRect(rect.cptr())
}

func HasScreenKeyboardSupport() bool {
	return C.SDL_HasScreenKeyboardSupport() > 0
}

func IsScreenKeyboardShown(window *Window) bool {
	return C.SDL_IsScreenKeyboardShown(window.cptr()) > 0
}

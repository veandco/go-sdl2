package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"
import "reflect"

// Keysym (https://wiki.libsdl.org/SDL_Keysym)
type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      uint16
	Unicode  uint32
}

// GetKeyboardFocus (https://wiki.libsdl.org/SDL_GetKeyboardFocus)
func GetKeyboardFocus() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetKeyboardFocus()))
}

// GetKeyboardState (https://wiki.libsdl.org/SDL_GetKeyboardState)
func GetKeyboardState() []uint8 {
	var numkeys C.int
	start := C.SDL_GetKeyboardState(&numkeys)
	sh := reflect.SliceHeader{}
	sh.Len = int(numkeys)
	sh.Cap = int(numkeys)
	sh.Data = uintptr(unsafe.Pointer(start))
	return *(*[]uint8)(unsafe.Pointer(&sh))
}

// GetModState (https://wiki.libsdl.org/SDL_GetModState)
func GetModState() Keymod {
	return (Keymod)(C.SDL_GetModState())
}

// SetModState (https://wiki.libsdl.org/SDL_SetModState)
func SetModState(mod Keymod) {
	C.SDL_SetModState(mod.c())
}

// GetKeyFromScancode (https://wiki.libsdl.org/SDL_GetKeyFromScancode)
func GetKeyFromScancode(code Scancode) Keycode {
	return (Keycode)(C.SDL_GetKeyFromScancode(code.c()))
}

// GetScancodeFromKey (https://wiki.libsdl.org/SDL_GetScancodeFromKey)
func GetScancodeFromKey(code Keycode) Scancode {
	return (Scancode)(C.SDL_GetScancodeFromKey(code.c()))
}

// GetScancodeName (https://wiki.libsdl.org/SDL_GetScancodeName)
func GetScancodeName(code Scancode) string {
	return (C.GoString)(C.SDL_GetScancodeName(code.c()))
}

// GetScancodeFromName (https://wiki.libsdl.org/SDL_GetScancodeFromName)
func GetScancodeFromName(name string) Scancode {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (Scancode)(C.SDL_GetScancodeFromName(_name))
}

// GetKeyName (https://wiki.libsdl.org/SDL_GetKeyName)
func GetKeyName(code Keycode) string {
	return (C.GoString)(C.SDL_GetKeyName(code.c()))
}

// GetKeyFromName (https://wiki.libsdl.org/SDL_GetKeyFromName)
func GetKeyFromName(name string) Keycode {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (Keycode)(C.SDL_GetKeyFromName(_name))
}

// StartTextInput (https://wiki.libsdl.org/SDL_StartTextInput)
func StartTextInput() {
	C.SDL_StartTextInput()
}

// IsTextInputActive (https://wiki.libsdl.org/SDL_IsTextInputActive)
func IsTextInputActive() bool {
	return C.SDL_IsTextInputActive() > 0
}

// StopTextInput (https://wiki.libsdl.org/SDL_StopTextInput)
func StopTextInput() {
	C.SDL_StopTextInput()
}

// SetTextInputRect (https://wiki.libsdl.org/SDL_SetTextInputRect)
func SetTextInputRect(rect *Rect) {
	C.SDL_SetTextInputRect(rect.cptr())
}

// HasScreenKeyboardSupport (https://wiki.libsdl.org/SDL_HasScreenKeyboardSupport)
func HasScreenKeyboardSupport() bool {
	return C.SDL_HasScreenKeyboardSupport() > 0
}

// IsScreenKeyboardShown (https://wiki.libsdl.org/SDL_IsScreenKeyboardShown)
func IsScreenKeyboardShown(window *Window) bool {
	return C.SDL_IsScreenKeyboardShown(window.cptr()) > 0
}

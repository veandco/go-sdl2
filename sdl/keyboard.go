package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,22))

#if defined(WARN_OUTDATED)
#pragma message("SDL_IsTextInputShown is not supported before SDL 2.0.22")
#pragma message("SDL_ClearComposition is not supported before SDL 2.0.22")
#endif

static inline SDL_bool SDL_IsTextInputShown(void)
{
	return SDL_FALSE;
}

static inline void SDL_ClearComposition(void)
{
}

#endif
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// Keysym contains key information used in key events.
// (https://wiki.libsdl.org/SDL_Keysym)
type Keysym struct {
	Scancode Scancode // SDL physical key code
	Sym      Keycode  // SDL virtual key code
	Mod      uint16   // current key modifiers
	unused   uint32   // unused
}

// GetKeyboardFocus returns the window which currently has keyboard focus.
// (https://wiki.libsdl.org/SDL_GetKeyboardFocus)
func GetKeyboardFocus() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetKeyboardFocus()))
}

// GetKeyboardState returns a snapshot of the current state of the keyboard.
// (https://wiki.libsdl.org/SDL_GetKeyboardState)
func GetKeyboardState() []uint8 {
	var numkeys C.int
	start := C.SDL_GetKeyboardState(&numkeys)
	sh := reflect.SliceHeader{}
	sh.Len = int(numkeys)
	sh.Cap = int(numkeys)
	sh.Data = uintptr(unsafe.Pointer(start))
	return *(*[]uint8)(unsafe.Pointer(&sh))
}

// GetModState returns the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL_GetModState)
func GetModState() Keymod {
	return (Keymod)(C.SDL_GetModState())
}

// SetModState sets the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL_SetModState)
func SetModState(mod Keymod) {
	C.SDL_SetModState(mod.c())
}

// GetKeyFromScancode returns the key code corresponding to the given scancode according to the current keyboard layout.
// (https://wiki.libsdl.org/SDL_GetKeyFromScancode)
func GetKeyFromScancode(code Scancode) Keycode {
	return (Keycode)(C.SDL_GetKeyFromScancode(code.c()))
}

// GetScancodeFromKey returns the scancode corresponding to the given key code according to the current keyboard layout.
// (https://wiki.libsdl.org/SDL_GetScancodeFromKey)
func GetScancodeFromKey(code Keycode) Scancode {
	return (Scancode)(C.SDL_GetScancodeFromKey(code.c()))
}

// GetScancodeName returns a human-readable name for a scancode
// (https://wiki.libsdl.org/SDL_GetScancodeName)
func GetScancodeName(code Scancode) string {
	return (C.GoString)(C.SDL_GetScancodeName(code.c()))
}

// GetScancodeFromName returns a scancode from a human-readable name.
// (https://wiki.libsdl.org/SDL_GetScancodeFromName)
func GetScancodeFromName(name string) Scancode {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (Scancode)(C.SDL_GetScancodeFromName(_name))
}

// GetKeyName returns a human-readable name for a key.
// (https://wiki.libsdl.org/SDL_GetKeyName)
func GetKeyName(code Keycode) string {
	return (C.GoString)(C.SDL_GetKeyName(code.c()))
}

// GetKeyFromName returns a key code from a human-readable name.
// (https://wiki.libsdl.org/SDL_GetKeyFromName)
func GetKeyFromName(name string) Keycode {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (Keycode)(C.SDL_GetKeyFromName(_name))
}

// StartTextInput starts accepting Unicode text input events.
// (https://wiki.libsdl.org/SDL_StartTextInput)
func StartTextInput() {
	C.SDL_StartTextInput()
}

// IsTextInputActive checks whether or not Unicode text input events are enabled.
// (https://wiki.libsdl.org/SDL_IsTextInputActive)
func IsTextInputActive() bool {
	return C.SDL_IsTextInputActive() > 0
}

// StopTextInput stops receiving any text input events.
// (https://wiki.libsdl.org/SDL_StopTextInput)
func StopTextInput() {
	C.SDL_StopTextInput()
}

// SetTextInputRect sets the rectangle used to type Unicode text inputs.
// (https://wiki.libsdl.org/SDL_SetTextInputRect)
func SetTextInputRect(rect *Rect) {
	C.SDL_SetTextInputRect(rect.cptr())
}

// HasScreenKeyboardSupport reports whether the platform has some screen keyboard support.
// (https://wiki.libsdl.org/SDL_HasScreenKeyboardSupport)
func HasScreenKeyboardSupport() bool {
	return C.SDL_HasScreenKeyboardSupport() > 0
}

// IsScreenKeyboardShown reports whether the screen keyboard is shown for given window.
// (https://wiki.libsdl.org/SDL_IsScreenKeyboardShown)
func IsScreenKeyboardShown(window *Window) bool {
	return C.SDL_IsScreenKeyboardShown(window.cptr()) > 0
}

// IsTextInputShown returns if an IME Composite or Candidate window is currently shown.
// (https://wiki.libsdl.org/SDL_IsTextInputShown)
func IsTextInputShown() bool {
	return C.SDL_IsTextInputShown() > 0
}

// ClearComposition dismisses the composition window/IME without disabling the subsystem.
// (https://wiki.libsdl.org/SDL_ClearComposition)
func ClearComposition() {
	C.SDL_ClearComposition()
}

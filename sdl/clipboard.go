package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// SetClipboardText (https://wiki.libsdl.org/SDL_SetClipboardText)
func SetClipboardText(text string) int {
	_text := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(_text))
	return int(C.SDL_SetClipboardText(_text))
}

// GetClipboardText (https://wiki.libsdl.org/SDL_GetClipboardText)
func GetClipboardText() string {
	return C.GoString(C.SDL_GetClipboardText())
}

// HasClipboardText (https://wiki.libsdl.org/SDL_HasClipboardText)
func HasClipboardText() bool {
	return C.SDL_HasClipboardText() > 0
}

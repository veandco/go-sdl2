package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// SetClipboardText (https://wiki.libsdl.org/SDL_SetClipboardText)
func SetClipboardText(text string) error {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	if C.SDL_SetClipboardText(_text) < 0 {
		return GetError()
	}
	return nil
}

// GetClipboardText (https://wiki.libsdl.org/SDL_GetClipboardText)
func GetClipboardText() (string, error) {
	text := C.SDL_GetClipboardText()
	if text == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(text))
	_text := C.GoString(text)
	return _text, nil
}

// HasClipboardText (https://wiki.libsdl.org/SDL_HasClipboardText)
func HasClipboardText() bool {
	return C.SDL_HasClipboardText() > 0
}

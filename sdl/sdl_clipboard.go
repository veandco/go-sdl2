package sdl

// #include <SDL2/SDL_stdinc.h>
// #include <SDL2/SDL_clipboard.h>
import "C"
import "unsafe"

func SetClipboardText(text string) int {
	_text := (C.CString) (text)
	defer C.SDL_free(unsafe.Pointer(_text))
	return (int) (C.SDL_SetClipboardText(_text))
}

func GetClipboardText() string {
	return (C.GoString) (C.SDL_GetClipboardText())
}

func HasClipboardText() bool {
	return C.SDL_HasClipboardText() > 0
}

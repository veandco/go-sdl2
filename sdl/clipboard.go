package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

func SetClipboardText(text string) int {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	return int(C.SDL_SetClipboardText(_text))
}

func GetClipboardText() string {
	return C.GoString(C.SDL_GetClipboardText())
}

func HasClipboardText() bool {
	return C.SDL_HasClipboardText() > 0
}

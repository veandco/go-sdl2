package sdl

// #include <SDL2/SDL.h>
import "C"

func GetError() string {
	return (string) (C.GoString(C.SDL_GetError()))
}

func ClearError() {
	C.SDL_ClearError()
}

func Error(code int) {
	_code := (C.SDL_errorcode) (C.int(code))
	C.SDL_Error(_code)
}

func OutOfMemory() {
	Error(ENOMEM)
}

func Unsupported() {
	Error(UNSUPPORTED)
}

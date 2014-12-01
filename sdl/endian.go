package sdl

// #include "sdl.h"
import "C"

const (
	LIL_ENDIAN = C.SDL_LIL_ENDIAN
	BIG_ENDIAN = C.SDL_BIG_ENDIAN
)

package sdl

// #include "sdl_wrapper.h"
import "C"

const (
	BYTEORDER  = C.SDL_BYTEORDER
	LIL_ENDIAN = C.SDL_LIL_ENDIAN
	BIG_ENDIAN = C.SDL_BIG_ENDIAN
)

package sdl

// #include <SDL.h>
import "C"

const (
	LIL_ENDIAN = C.SDL_LIL_ENDIAN
	BIG_ENDIAN = C.SDL_BIG_ENDIAN
)

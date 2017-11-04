package sdl

// #include "sdl_wrapper.h"
import "C"

// Endian-specific values.
// (https://wiki.libsdl.org/CategoryEndian)
const (
	BYTEORDER  = C.SDL_BYTEORDER  // macro that corresponds to the byte order used by the processor type it was compiled for
	LIL_ENDIAN = C.SDL_LIL_ENDIAN // byte order is 1234, where the least significant byte is stored first
	BIG_ENDIAN = C.SDL_BIG_ENDIAN // byte order is 4321, where the most significant byte is stored first
)

package sdl

// #include <SDL2/SDL.h>
// #include "misc.h"
import "C"

const CACHELINE_SIZE = 128

func GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

func HasRDTSC() bool {
	return C.SDL_HasRDTSC() > 0
}

func HasAltiVec() bool {
	return C.SDL_HasAltiVec() > 0
}

func HasMMX() bool {
	return C.SDL_HasMMX() > 0
}

func Has3DNow() bool {
	return C.SDL_Has3DNow() > 0
}

func HasSSE() bool {
	return C.SDL_HasSSE() > 0
}

func HasSSE2() bool {
	return C.SDL_HasSSE2() > 0
}

func HasSSE3() bool {
	return C.SDL_HasSSE3() > 0
}

func HasSSE41() bool {
	return C.SDL_HasSSE41() > 0
}

func HasSSE42() bool {
	return C.SDL_HasSSE42() > 0
}

func GetSystemRAM() int {
	return int(C._SDL_GetSystemRAM())
}

func HasAVX() bool {
	return C._SDL_HasAVX() > 0
}

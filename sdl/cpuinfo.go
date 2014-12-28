package sdl

/*
#include "sdl_wrapper.h"

static inline int _SDL_GetSystemRAM()
{
#if (SDL_VERSION_ATLEAST(2,0,1))
    return SDL_GetSystemRAM();
#else
    return 0;
#endif
}

static inline SDL_bool _SDL_HasAVX()
{
#if (SDL_VERSION_ATLEAST(2,0,2))
    return SDL_HasAVX();
#else
    return SDL_FALSE;
#endif
}

*/
import "C"

const CACHELINE_SIZE = C.SDL_CACHELINE_SIZE

// GetCPUCount (https://wiki.libsdl.org/SDL_GetCPUCount)
func GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

// GetCPUCacheLineSize (https://wiki.libsdl.org/SDL_GetCPUCacheLineSize)
func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

// HasRDTSC (https://wiki.libsdl.org/SDL_HasRDTSC)
func HasRDTSC() bool {
	return C.SDL_HasRDTSC() > 0
}

// HasAltiVec (https://wiki.libsdl.org/SDL_HasAltiVec)
func HasAltiVec() bool {
	return C.SDL_HasAltiVec() > 0
}

// HasMMX (https://wiki.libsdl.org/SDL_HasMMX)
func HasMMX() bool {
	return C.SDL_HasMMX() > 0
}

// Has3DNow (https://wiki.libsdl.org/SDL_Has3DNow)
func Has3DNow() bool {
	return C.SDL_Has3DNow() > 0
}

// HasSSE (https://wiki.libsdl.org/SDL_HasSSE)
func HasSSE() bool {
	return C.SDL_HasSSE() > 0
}

// HasSSE2 (https://wiki.libsdl.org/SDL_HasSSE2)
func HasSSE2() bool {
	return C.SDL_HasSSE2() > 0
}

// HasSSE3 (https://wiki.libsdl.org/SDL_HasSSE3)
func HasSSE3() bool {
	return C.SDL_HasSSE3() > 0
}

// HasSSE41 (https://wiki.libsdl.org/SDL_HasSSE41)
func HasSSE41() bool {
	return C.SDL_HasSSE41() > 0
}

// HasSSE42 (https://wiki.libsdl.org/SDL_HasSSE42)
func HasSSE42() bool {
	return C.SDL_HasSSE42() > 0
}

// GetSystemRAM (https://wiki.libsdl.org/SDL_GetSystemRAM)
func GetSystemRAM() int {
	return int(C._SDL_GetSystemRAM())
}

func HasAVX() bool {
	return C._SDL_HasAVX() > 0
}

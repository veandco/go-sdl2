package sdl

/*
#include "sdl.h"

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

package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,1))
#pragma message("SDL_GetSystemRAM is not supported before SDL 2.0.1")
static inline int SDL_GetSystemRAM()
{
	return -1;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,2))
#pragma message("SDL_HasAVX is not supported before SDL 2.0.2")
static inline SDL_bool SDL_HasAVX()
{
	return SDL_FALSE;
}
#endif

*/
import "C"

// CACHELINE_SIZE is a cacheline size used for padding.
const CACHELINE_SIZE = C.SDL_CACHELINE_SIZE

// GetCPUCount returns the number of CPU cores available.
// (https://wiki.libsdl.org/SDL_GetCPUCount)
func GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

// GetCPUCacheLineSize returns the L1 cache line size of the CPU.
// (https://wiki.libsdl.org/SDL_GetCPUCacheLineSize)
func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

// HasRDTSC reports whether the CPU has the RDTSC instruction.
// (https://wiki.libsdl.org/SDL_HasRDTSC)
func HasRDTSC() bool {
	return C.SDL_HasRDTSC() > 0
}

// HasAltiVec reports whether the CPU has AltiVec features.
// (https://wiki.libsdl.org/SDL_HasAltiVec)
func HasAltiVec() bool {
	return C.SDL_HasAltiVec() > 0
}

// HasMMX reports whether the CPU has MMX features.
// (https://wiki.libsdl.org/SDL_HasMMX)
func HasMMX() bool {
	return C.SDL_HasMMX() > 0
}

// Has3DNow reports whether the CPU has 3DNow! features.
// (https://wiki.libsdl.org/SDL_Has3DNow)
func Has3DNow() bool {
	return C.SDL_Has3DNow() > 0
}

// HasSSE reports whether the CPU has SSE features.
// (https://wiki.libsdl.org/SDL_HasSSE)
func HasSSE() bool {
	return C.SDL_HasSSE() > 0
}

// HasSSE2 reports whether the CPU has SSE2 features.
// (https://wiki.libsdl.org/SDL_HasSSE2)
func HasSSE2() bool {
	return C.SDL_HasSSE2() > 0
}

// HasSSE3 reports whether the CPU has SSE3 features.
// (https://wiki.libsdl.org/SDL_HasSSE3)
func HasSSE3() bool {
	return C.SDL_HasSSE3() > 0
}

// HasSSE41 reports whether the CPU has SSE4.1 features.
// (https://wiki.libsdl.org/SDL_HasSSE41)
func HasSSE41() bool {
	return C.SDL_HasSSE41() > 0
}

// HasSSE42 reports whether the CPU has SSE4.2 features.
// (https://wiki.libsdl.org/SDL_HasSSE42)
func HasSSE42() bool {
	return C.SDL_HasSSE42() > 0
}

// GetSystemRAM returns the amount of RAM configured in the system.
// (https://wiki.libsdl.org/SDL_GetSystemRAM)
func GetSystemRAM() int {
	return int(C.SDL_GetSystemRAM())
}

// HasAVX reports whether the CPU has AVX features.
// (https://wiki.libsdl.org/SDL_HasAVX)
func HasAVX() bool {
	return C.SDL_HasAVX() > 0
}

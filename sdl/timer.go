package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetTicks64 is not supported before SDL 2.0.18")
#endif

static inline Uint64 SDLCALL SDL_GetTicks64(void)
{
	return 0;
}
#endif
*/
import "C"

// GetTicks returns the number of milliseconds since the SDL library initialization.
//
// Deprecated: This function is not recommended as of SDL 2.0.18; use GetTicks64()
// instead, where the value doesn't wrap every ~49 days. There are places in
// SDL where we provide a 32-bit timestamp that can not change without
// breaking binary compatibility, though, so this function isn't officially
// deprecated.
//
// (https://wiki.libsdl.org/SDL_GetTicks)
func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// GetTicks64 returns the number of milliseconds since the SDL library initialization.
// (https://wiki.libsdl.org/SDL_GetTicks64)
func GetTicks64() uint64 {
	return uint64(C.SDL_GetTicks64())
}

// GetPerformanceCounter returns the current value of the high resolution counter.
// (https://wiki.libsdl.org/SDL_GetPerformanceCounter)
func GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

// GetPerformanceFrequency returns the count per second of the high resolution counter.
// (https://wiki.libsdl.org/SDL_GetPerformanceFrequency)
func GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}

// Delay waits a specified number of milliseconds before returning.
// (https://wiki.libsdl.org/SDL_Delay)
func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

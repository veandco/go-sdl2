package sdl

// #include "sdl_wrapper.h"
import "C"

// GetTicks returns the number of milliseconds since the SDL library initialization.
// (https://wiki.libsdl.org/SDL_GetTicks)
func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
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

package sdl

// #include "sdl_wrapper.h"
import "C"

// GetTicks (https://wiki.libsdl.org/SDL_GetTicks)
func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// GetPerformanceCounter (https://wiki.libsdl.org/SDL_GetPerformanceCounter)
func GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

// GetPerformanceFrequency (https://wiki.libsdl.org/SDL_GetPerformanceFrequency)
func GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}

// Delay (https://wiki.libsdl.org/SDL_Delay)
func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

package sdl

// #include <SDL.h>
import "C"

func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

func GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

func GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}

func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

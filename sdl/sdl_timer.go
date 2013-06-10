package sdl

// #include <SDL2/SDL_timer.h>
import "C"

func Delay(ms uint32) {
	_ms := (C.Uint32) (ms)
	C.SDL_Delay(C.Uint32(_ms));
}

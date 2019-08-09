package sdl

// #include "sdl_wrapper.h"
import "C"

// CurrentThreadID gets the thread identifier for the current thread.
// (https://wiki.libsdl.org/SDL_ThreadID)
func CurrentThreadID() uint {
	return uint(C.SDL_ThreadID())
}

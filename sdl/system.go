package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,9))

#if defined(WARN_OUTDATED)
#pragma message("SDL_IsTablet is not supported before SDL 2.0.9")
#endif

static inline SDL_bool SDL_IsTablet()
{
	return SDL_FALSE;
}

#endif
*/
import "C"

// IsTablet returns true if the current device is a tablet
// TODO: (https://wiki.libsdl.org/SDL_IsTablet)
func IsTablet() bool {
	return C.SDL_IsTablet() == C.SDL_TRUE
}

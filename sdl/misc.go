package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_OpenURL is not supported before SDL 2.0.14")
#endif

static int SDL_OpenURL(const char *url)
{
	return -1;
}

#endif
*/
import "C"
import "unsafe"

// OpenURL opens an URL / URI in the browser or other
// (https://wiki.libsdl.org/SDL_OpenURL)
func OpenURL(url string) error {
	_url := C.CString(url)
	defer C.free(unsafe.Pointer(_url))
	return errorFromInt(int(C.SDL_OpenURL(_url)))
}

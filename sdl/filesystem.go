package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,1))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetBasePath is not supported before SDL 2.0.1")
#endif

static inline char* SDL_GetBasePath()
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_GetPrefPath is not supported before SDL 2.0.1")
#endif

static inline char* SDL_GetPrefPath(const char *org, const char *app)
{
	return NULL;
}
#endif

*/
import "C"
import "unsafe"

// GetBasePath returns the directory where the application was run from. This is where the application data directory is.
// (https://wiki.libsdl.org/SDL_GetBasePath)
func GetBasePath() string {
	_val := C.SDL_GetBasePath()
	defer C.SDL_free(unsafe.Pointer(_val))
	return C.GoString(_val)
}

// GetPrefPath returns the "pref dir". This is meant to be where the application can write personal files (Preferences and save games, etc.) that are specific to the application. This directory is unique per user and per application.
// (https://wiki.libsdl.org/SDL_GetPrefPath)
func GetPrefPath(org, app string) string {
	_org := C.CString(org)
	_app := C.CString(app)
	defer C.free(unsafe.Pointer(_org))
	defer C.free(unsafe.Pointer(_app))
	_val := C.SDL_GetPrefPath(_org, _app)
	defer C.SDL_free(unsafe.Pointer(_val))
	return C.GoString(_val)
}

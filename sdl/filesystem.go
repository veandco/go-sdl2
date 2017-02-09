package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,1))
#pragma message("SDL_GetBasePath is not supported before SDL 2.0.1")
static inline char* SDL_GetBasePath()
{
	return NULL;
}

#pragma message("SDL_GetPrefPath is not supported before SDL 2.0.1")
static inline char* SDL_GetPrefPath(const char *org, const char *app)
{
	return NULL;
}
#endif

*/
import "C"
import "unsafe"

// GetBasePath (https://wiki.libsdl.org/SDL_GetBasePath)
func GetBasePath() string {
	_val := C.SDL_GetBasePath()
	defer C.SDL_free(unsafe.Pointer(_val))
	return C.GoString(_val)
}

// GetPrefPath (https://wiki.libsdl.org/SDL_GetPrefPath)
func GetPrefPath(org, app string) string {
	_org := C.CString(org)
	_app := C.CString(app)
	defer C.free(unsafe.Pointer(_org))
	defer C.free(unsafe.Pointer(_app))
	_val := C.SDL_GetPrefPath(_org, _app)
	defer C.SDL_free(unsafe.Pointer(_val))
	return C.GoString(_val)
}

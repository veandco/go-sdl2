package sdl

/*
#include <SDL2/SDL.h>

static inline char* _SDL_GetBasePath() {
#if  (SDL_MAJOR_VERSION >= 2 && SDL_MINOR_VERSION >= 0 && SDL_PATCHLEVEL >= 1)
    return SDL_GetBasePath();
#else
    return NULL;
#endif
}

static inline char* _SDL_GetPrefPath(const char *org, const char *app)
{
#if  (SDL_MAJOR_VERSION >= 2 && SDL_MINOR_VERSION >= 0 && SDL_PATCHLEVEL >= 1)
    return SDL_GetPrefPath(org, app);
#else
    return NULL;
#endif
}

*/
import "C"
import "unsafe"

func GetBasePath() string {
	_val := C._SDL_GetBasePath()
	defer C.SDL_free(unsafe.Pointer(_val))
	return C.GoString(_val)
}

func GetPrefPath(org, app string) string {
	_org := C.CString(org)
	_app := C.CString(app)
	defer C.free(unsafe.Pointer(_org))
	defer C.free(unsafe.Pointer(_app))
	_val := C._SDL_GetPrefPath(_org, _app)
	defer C.SDL_free(unsafe.Pointer(_val))
	return C.GoString(_val)
}

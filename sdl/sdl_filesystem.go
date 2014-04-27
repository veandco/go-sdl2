package sdl

// #include "misc.h"
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

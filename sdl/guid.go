package sdl

/*
#include "sdl_wrapper.h"

#if !SDL_VERSION_ATLEAST(2,24,0)
#if defined(WARN_OUTDATED)
#pragma message("SDL_GUIDToString is not supported before SDL 2.24.0")
#pragma message("SDL_GUIDFromString is not supported before SDL 2.24.0")
#endif

typedef struct {
    Uint8 data[16];
} SDL_GUID;

static inline void SDL_GUIDToString(SDL_GUID guid, char *pszGUID, int cbGUID)
{
	return;
}

static inline SDL_GUID SDL_GUIDFromString(const char *pchGUID)
{
	SDL_GUID guid;
	return guid;
}
#endif
*/
import "C"
import (
	"unsafe"
)

type GUID C.SDL_GUID

// ToString returns an ASCII string representation for a given GUID.
func (guid GUID) ToString() (ascii string) {
	_cap := C.size_t(33)
	_buf := (*C.char)(C.SDL_malloc(_cap))
	defer C.SDL_free(unsafe.Pointer(_buf))
	C.SDL_GUIDToString(C.SDL_GUID(guid), _buf, C.int(_cap))
	return C.GoString(_buf)
}

// GUIDFromString converts a GUID string into a GUID structure.
func GUIDFromString(ascii string) (guid GUID) {
	_ascii := C.CString(ascii)
	return GUID(C.SDL_GUIDFromString(_ascii))
}

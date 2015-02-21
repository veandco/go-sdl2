package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

const (
	MAJOR_VERSION = C.SDL_MAJOR_VERSION
	MINOR_VERSION = C.SDL_MINOR_VERSION
	PATCHLEVEL    = C.SDL_PATCHLEVEL
)

// Version (https://wiki.libsdl.org/SDL_Version)
type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}
type cVersion C.SDL_version

func (v *Version) cptr() *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}

// VERSION (https://wiki.libsdl.org/SDL_VERSION)
func VERSION(version *Version) {
	version.Major = MAJOR_VERSION
	version.Minor = MINOR_VERSION
	version.Patch = PATCHLEVEL
}

// VERSIONNUM (https://wiki.libsdl.org/SDL_VERSIONNUM)
func VERSIONNUM(x, y, z int) int {
	return (x*1000 + y*100 + z)
}

// COMPILEDVERSION (https://wiki.libsdl.org/SDL_COMPILEDVERSION)
func COMPILEDVERSION() int {
	return VERSIONNUM(MAJOR_VERSION, MINOR_VERSION, PATCHLEVEL)
}

// VERSION_ATLEAST (https://wiki.libsdl.org/SDL_VERSION_ATLEAST)
func VERSION_ATLEAST(x, y, z int) bool {
	return COMPILEDVERSION() >= VERSIONNUM(x, y, z)
}

// GetVersion (https://wiki.libsdl.org/SDL_GetVersion)
func GetVersion(v *Version) {
	C.SDL_GetVersion(v.cptr())
}

// GetRevision (https://wiki.libsdl.org/SDL_GetRevision)
func GetRevision() string {
	return (string)(C.GoString(C.SDL_GetRevision()))
}

// GetRevisionNumber (https://wiki.libsdl.org/SDL_GetRevisionNumber)
func GetRevisionNumber() int {
	return (int)(C.SDL_GetRevisionNumber())
}

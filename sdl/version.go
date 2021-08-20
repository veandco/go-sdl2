package sdl

/*
#include "sdl_wrapper.h"

#if SDL_VERSION_ATLEAST(2,0,16)

static inline int GetRevisionNumber(void)
{
	return 0;
}

#else

static inline int GetRevisionNumber(void)
{
	return SDL_GetRevisionNumber();
}

#endif
*/
import "C"
import "unsafe"

// The version of SDL in use.
const (
	MAJOR_VERSION = C.SDL_MAJOR_VERSION // major version
	MINOR_VERSION = C.SDL_MINOR_VERSION // minor version
	PATCHLEVEL    = C.SDL_PATCHLEVEL    // update version (patchlevel)
)

// Version contains information about the version of SDL in use.
// (https://wiki.libsdl.org/SDL_version)
type Version struct {
	Major uint8 // major version
	Minor uint8 // minor version
	Patch uint8 // update version (patchlevel)
}
type cVersion C.SDL_version

func (v *Version) cptr() *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}

// VERSION fills the selected struct with the version of SDL in use.
// (https://wiki.libsdl.org/SDL_VERSION)
func VERSION(v *Version) {
	v.Major = MAJOR_VERSION
	v.Minor = MINOR_VERSION
	v.Patch = PATCHLEVEL
}

// VERSIONNUM converts separate version components into a single numeric value.
// (https://wiki.libsdl.org/SDL_VERSIONNUM)
func VERSIONNUM(x, y, z int) int {
	return (x*1000 + y*100 + z)
}

// COMPILEDVERSION returns the SDL version number that you compiled against.
// (https://wiki.libsdl.org/SDL_COMPILEDVERSION)
func COMPILEDVERSION() int {
	return VERSIONNUM(MAJOR_VERSION, MINOR_VERSION, PATCHLEVEL)
}

// VERSION_ATLEAST reports whether the SDL version compiled against is at least as new as the specified version.
// (https://wiki.libsdl.org/SDL_VERSION_ATLEAST)
func VERSION_ATLEAST(x, y, z int) bool {
	return COMPILEDVERSION() >= VERSIONNUM(x, y, z)
}

// GetVersion returns the version of SDL that is linked against your program.
// (https://wiki.libsdl.org/SDL_GetVersion)
func GetVersion(v *Version) {
	C.SDL_GetVersion(v.cptr())
}

// GetRevision returns the code revision of SDL that is linked against your program.
// (https://wiki.libsdl.org/SDL_GetRevision)
func GetRevision() string {
	return (string)(C.GoString(C.SDL_GetRevision()))
}

// Deprecated: GetRevisionNumber is deprecated in SDL2 2.0.16 and will return 0. Users should use GetRevision instead.
func GetRevisionNumber() int {
	return (int)(C.GetRevisionNumber())
}

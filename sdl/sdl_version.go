package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	MAJOR_VERSION = C.SDL_MAJOR_VERSION
	MINOR_VERSION = C.SDL_MINOR_VERSION
	PATCHLEVEL    = C.SDL_PATCHLEVEL
)

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func (v *Version) cptr() *C.SDL_version {
    return (*C.SDL_version)(unsafe.Pointer(v))
}

func VERSION(version *Version) {
	version.Major = C.SDL_MAJOR_VERSION
	version.Minor = C.SDL_MINOR_VERSION
	version.Patch = C.SDL_PATCHLEVEL
}

func VERSIONNUM(x, y, z int) int {
	return (x*1000 + y*100 + z)
}

func COMPILEDVERSION() int {
	return VERSIONNUM(MAJOR_VERSION, MINOR_VERSION, PATCHLEVEL)
}

func VERSION_ATLEAST(x, y, z int) bool {
	return COMPILEDVERSION() >= VERSIONNUM(x, y, z)
}

func GetVersion(v *Version) {
	C.SDL_GetVersion(v.cptr())
}

func GetRevision() string {
	return (string)(C.GoString(C.SDL_GetRevision()))
}

func GetRevisionNumber() int {
	return (int)(C.SDL_GetRevisionNumber())
}

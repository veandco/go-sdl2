package sdl

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
import "C"
import "unsafe"

func Init(flags uint32) int {
	_flags := (C.Uint32) (flags)
	return (int) (C.SDL_Init(_flags))
}

func Quit() {
	C.SDL_Quit()
}

func InitSubSystem(flags uint32) int {
	_flags := (C.Uint32) (flags)
	return (int) (C.SDL_InitSubSystem(_flags))
}

func QuitSubSystem(flags uint32) {
	_flags := (C.Uint32) (flags)
	C.SDL_QuitSubSystem(_flags)
}

func WasInit(flags uint32) uint32 {
	_flags := (C.Uint32) (flags)
	return (uint32) (C.SDL_WasInit(_flags))
}

func GetVersion(v *Version) {
	version := (*C.SDL_version) (unsafe.Pointer(v))
	C.SDL_GetVersion(version)
}

func GetRevision() string {
	return (string) (C.GoString(C.SDL_GetRevision()))
}

func GetRevisionNumber() int {
	return (int) (C.SDL_GetRevisionNumber())
}

func GetPlatform() string {
	return (string) (C.GoString(C.SDL_GetPlatform()))
}

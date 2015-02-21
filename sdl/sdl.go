package sdl

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "sdl_wrapper.h"
import "C"

const (
	INIT_TIMER          = 0x00000001
	INIT_AUDIO          = 0x00000010
	INIT_VIDEO          = 0x00000020
	INIT_JOYSTICK       = 0x00000200
	INIT_HAPTIC         = 0x00001000
	INIT_GAMECONTROLLER = 0x00002000
	INIT_NOPARACHUTE    = 0x00100000
	INIT_EVERYTHING     = INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_JOYSTICK |
		INIT_HAPTIC | INIT_GAMECONTROLLER
)

const (
	RELEASED = 0
	PRESSED  = 1
)

// Init (https://wiki.libsdl.org/SDL_Init)
func Init(flags uint32) error {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// Quit (https://wiki.libsdl.org/SDL_Quit)
func Quit() {
	C.SDL_Quit()

	eventFilterCache = nil
	for k, _ := range eventWatchesCache {
		delete(eventWatchesCache, k)
	}
}

// InitSubSystem (https://wiki.libsdl.org/SDL_InitSubSystem)
func InitSubSystem(flags uint32) error {
	if C.SDL_InitSubSystem(C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// QuitSubSystem (https://wiki.libsdl.org/SDL_QuitSubSystem)
func QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

// WasInit (https://wiki.libsdl.org/SDL_WasInit)
func WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(C.Uint32(flags)))
}

// GetPlatform (https://wiki.libsdl.org/SDL_GetPlatform)
func GetPlatform() string {
	return string(C.GoString(C.SDL_GetPlatform()))
}

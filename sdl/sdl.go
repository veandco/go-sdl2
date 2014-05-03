package sdl

// #cgo windows LDFLAGS: -lSDL2
// #cgo darwin LDFLAGS: -framework SDL2
// #cgo linux freebsd pkg-config: sdl2
// #include <SDL2/SDL.h>
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

func Init(flags uint32) int {
	return int(C.SDL_Init(C.Uint32(flags)))
}

func Quit() {
	C.SDL_Quit()
}

func InitSubSystem(flags uint32) int {
	return int(C.SDL_InitSubSystem(C.Uint32(flags)))
}

func QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

func WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(C.Uint32(flags)))
}

func GetPlatform() string {
	return string(C.GoString(C.SDL_GetPlatform()))
}

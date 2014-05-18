package sdl

// #include <SDL2/SDL.h>
import "C"

const (
	POWERSTATE_UNKNOWN = iota
	POWERSTATE_ON_BATTERY
	POWERSTATE_NO_BATTERY
	POWERSTATE_CHARGING
	POWERSTATE_CHARGED
)

type PowerState C.SDL_PowerState

func GetPowerInfo() (int, int, int) {
	_secs := C.int(0)
	_percent := C.int(0)
	_state := C.SDL_GetPowerInfo(&_secs, &_percent)
	return (int)(_state), (int)(_secs), (int)(_percent)
}

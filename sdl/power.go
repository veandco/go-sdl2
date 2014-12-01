package sdl

// #include "sdl.h"
import "C"

const (
	POWERSTATE_UNKNOWN = C.SDL_POWERSTATE_UNKNOWN
	POWERSTATE_ON_BATTERY = C.SDL_POWERSTATE_ON_BATTERY
	POWERSTATE_NO_BATTERY = C.SDL_POWERSTATE_NO_BATTERY
	POWERSTATE_CHARGING = C.SDL_POWERSTATE_CHARGING
	POWERSTATE_CHARGED = C.SDL_POWERSTATE_CHARGED
)

type PowerState C.SDL_PowerState

func GetPowerInfo() (int, int, int) {
	_secs := C.int(0)
	_percent := C.int(0)
	_state := C.SDL_GetPowerInfo(&_secs, &_percent)
	return (int)(_state), (int)(_secs), (int)(_percent)
}

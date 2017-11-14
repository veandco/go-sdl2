package sdl

// #include "sdl_wrapper.h"
import "C"

// An enumeration of the basic state of the system's power supply.
// (https://wiki.libsdl.org/SDL_PowerState)
const (
	POWERSTATE_UNKNOWN    = C.SDL_POWERSTATE_UNKNOWN    // cannot determine power status
	POWERSTATE_ON_BATTERY = C.SDL_POWERSTATE_ON_BATTERY // not plugged in, running on the battery
	POWERSTATE_NO_BATTERY = C.SDL_POWERSTATE_NO_BATTERY // plugged in, no battery available
	POWERSTATE_CHARGING   = C.SDL_POWERSTATE_CHARGING   // plugged in, charging battery
	POWERSTATE_CHARGED    = C.SDL_POWERSTATE_CHARGED    // plugged in, battery charged
)

// PowerState is the basic state for the system's power supply.
// (https://wiki.libsdl.org/SDL_PowerState)
type PowerState C.SDL_PowerState

// GetPowerInfo returns the current power supply details.
// (https://wiki.libsdl.org/SDL_GetPowerInfo)
func GetPowerInfo() (int, int, int) {
	_secs := C.int(0)
	_percent := C.int(0)
	_state := C.SDL_GetPowerInfo(&_secs, &_percent)
	return (int)(_state), (int)(_secs), (int)(_percent)
}

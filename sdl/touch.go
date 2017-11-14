package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// TOUCH_MOUSEID is the device ID for mouse events simulated with touch input
const TOUCH_MOUSEID = C.SDL_TOUCH_MOUSEID

// TouchID is the ID of a touch device.
type TouchID C.SDL_TouchID

// FingerID is a finger id.
type FingerID C.SDL_FingerID

// Finger contains touch information.
type Finger struct {
	ID       FingerID // the finger id
	X        float32  // the x-axis location of the touch event, normalized (0...1)
	Y        float32  // the y-axis location of the touch event, normalized (0...1)
	Pressure float32  // the quantity of pressure applied, normalized (0...1)
}

func (t TouchID) c() C.SDL_TouchID {
	return C.SDL_TouchID(t)
}

// GetNumTouchDevices returns the number of registered touch devices.
// (https://wiki.libsdl.org/SDL_GetNumTouchDevices)
func GetNumTouchDevices() int {
	return int(C.SDL_GetNumTouchDevices())
}

// GetTouchDevice returns the touch ID with the given index.
// (https://wiki.libsdl.org/SDL_GetTouchDevice)
func GetTouchDevice(index int) TouchID {
	return TouchID(C.SDL_GetTouchDevice(C.int(index)))
}

// GetNumTouchFingers returns the number of active fingers for a given touch device.
// (https://wiki.libsdl.org/SDL_GetNumTouchFingers)
func GetNumTouchFingers(t TouchID) int {
	return int(C.SDL_GetNumTouchFingers(t.c()))
}

// GetTouchFinger returns the finger object for specified touch device ID and finger index.
// (https://wiki.libsdl.org/SDL_GetTouchFinger)
func GetTouchFinger(t TouchID, index int) *Finger {
	return (*Finger)(unsafe.Pointer(C.SDL_GetTouchFinger(t.c(), C.int(index))))
}

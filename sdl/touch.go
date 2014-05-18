package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const TOUCH_MOUSEID = C.SDL_TOUCH_MOUSEID

type TouchID C.SDL_TouchID
type FingerID C.SDL_FingerID
type Finger struct {
	Id       FingerID
	X        float32
	Y        float32
	Pressure float32
}

func (t TouchID) c() C.SDL_TouchID {
    return C.SDL_TouchID(t)
}

func GetNumTouchDevices() int {
	return int(C.SDL_GetNumTouchDevices())
}

func GetTouchDevice(index int) TouchID {
	return TouchID(C.SDL_GetTouchDevice(C.int(index)))
}

func GetNumTouchFingers(t TouchID) int {
	return int(C.SDL_GetNumTouchFingers(t.c()))
}

func GetTouchFinger(t TouchID, index int) *Finger {
	return (*Finger)(unsafe.Pointer(C.SDL_GetTouchFinger(t.c(), C.int(index))))
}

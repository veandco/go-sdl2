package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const TOUCH_MOUSEID = C.SDL_TOUCH_MOUSEID

type TouchID C.SDL_TouchID
type FingerID C.SDL_FingerID
type Finger struct {
	Id FingerID
	X float32
	Y float32
	Pressure float32
}

func GetNumTouchDevices() int {
	return (int) (C.SDL_GetNumTouchDevices())
}

func GetTouchDevice(index int) TouchID {
	_index := (C.int) (index)
	return (TouchID) (C.SDL_GetTouchDevice(_index))
}

func GetNumTouchFingers(touchId TouchID) int {
	_touchId := (C.SDL_TouchID) (touchId)
	return (int) (C.SDL_GetNumTouchFingers(_touchId))
}

func GetTouchFinger(touchId TouchID, index int) *Finger {
	_touchId := (C.SDL_TouchID) (touchId)
	_index := (C.int) (index)
	return (*Finger) (unsafe.Pointer(C.SDL_GetTouchFinger(_touchId, _index)))
}

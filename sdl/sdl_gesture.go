package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

type GestureID C.SDL_GestureID

func RecordGesture(touchId TouchID) int {
	_touchId := (C.SDL_TouchID) (touchId)
	return (int) (C.SDL_RecordGesture(_touchId))
}

func SaveAllDollarTemplates(src *RWops) int {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (int) (C.SDL_SaveAllDollarTemplates(_src))
}

func SaveDollarTemplate(gestureId GestureID, src *RWops) int {
	_gestureId := (C.SDL_GestureID) (gestureId)
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (int) (C.SDL_SaveDollarTemplate(_gestureId, _src))
}

func LoadDollarTemplates(touchId TouchID, src *RWops) int {
	_touchId := (C.SDL_TouchID) (touchId)
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (int) (C.SDL_LoadDollarTemplates(_touchId, _src))
}

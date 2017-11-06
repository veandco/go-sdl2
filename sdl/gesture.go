package sdl

// #include "sdl_wrapper.h"
import "C"

// GestureID is the unique id of the closest gesture to the performed stroke.
type GestureID C.SDL_GestureID

func (g GestureID) c() C.SDL_GestureID {
	return C.SDL_GestureID(g)
}

// RecordGesture begins recording a gesture on a specified touch device or all touch devices.
// (https://wiki.libsdl.org/SDL_RecordGesture)
func RecordGesture(t TouchID) int {
	return int(C.SDL_RecordGesture(t.c()))
}

// SaveAllDollarTemplates saves all currently loaded Dollar Gesture templates.
// (https://wiki.libsdl.org/SDL_SaveAllDollarTemplates)
func SaveAllDollarTemplates(src *RWops) int {
	return int(C.SDL_SaveAllDollarTemplates(src.cptr()))
}

// SaveDollarTemplate saves a currently loaded Dollar Gesture template.
// (https://wiki.libsdl.org/SDL_SaveDollarTemplate)
func SaveDollarTemplate(g GestureID, src *RWops) int {
	return int(C.SDL_SaveDollarTemplate(g.c(), src.cptr()))
}

// LoadDollarTemplates loads Dollar Gesture templates from a file.
// (https://wiki.libsdl.org/SDL_LoadDollarTemplates)
func LoadDollarTemplates(t TouchID, src *RWops) int {
	return int(C.SDL_LoadDollarTemplates(t.c(), src.cptr()))
}

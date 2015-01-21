package sdl

// #include "sdl_wrapper.h"
import "C"

type GestureID C.SDL_GestureID

func (g GestureID) c() C.SDL_GestureID {
	return C.SDL_GestureID(g)
}

// RecordGesture (https://wiki.libsdl.org/SDL_RecordGesture)
func RecordGesture(t TouchID) int {
	return int(C.SDL_RecordGesture(t.c()))
}

// SaveAllDollarTemplates (https://wiki.libsdl.org/SDL_SaveAllDollarTemplates)
func SaveAllDollarTemplates(src *RWops) int {
	return int(C.SDL_SaveAllDollarTemplates(src.cptr()))
}

// SaveDollarTemplate (https://wiki.libsdl.org/SDL_SaveDollarTemplate)
func SaveDollarTemplate(g GestureID, src *RWops) int {
	return int(C.SDL_SaveDollarTemplate(g.c(), src.cptr()))
}

// LoadDollarTemplates (https://wiki.libsdl.org/SDL_LoadDollarTemplates)
func LoadDollarTemplates(t TouchID, src *RWops) int {
	return int(C.SDL_LoadDollarTemplates(t.c(), src.cptr()))
}

package sdl

// #include <SDL2/SDL.h>
import "C"

type GestureID C.SDL_GestureID

func (g GestureID) c() C.SDL_GestureID {
    return C.SDL_GestureID(g)
}

func RecordGesture(t TouchID) int {
	return int(C.SDL_RecordGesture(t.c()))
}

func SaveAllDollarTemplates(src *RWops) int {
	return int(C.SDL_SaveAllDollarTemplates(src.cptr()))
}

func SaveDollarTemplate(g GestureID, src *RWops) int {
	return int(C.SDL_SaveDollarTemplate(g.c(), src.cptr()))
}

func LoadDollarTemplates(t TouchID, src *RWops) int {
	return int(C.SDL_LoadDollarTemplates(t.c(), src.cptr()))
}

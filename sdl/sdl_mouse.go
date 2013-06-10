package sdl

/*
#include <SDL2/SDL_mouse.h>
*/
import "C"
import "unsafe"

/* SystemCursor */
const (
	SYSTEM_CURSOR_ARROW = iota
	SYSTEM_CURSOR_IBEAM
	SYSTEM_CURSOR_WAIT
	SYSTEM_CURSOR_CROSSHAIR
	SYSTEM_CURSOR_WAITARROW
	SYSTEM_CURSOR_SIZENWSE
	SYSTEM_CURSOR_SIZENESW
	SYSTEM_CURSOR_SIZEWE
	SYSTEM_CURSOR_SIZENS
	SYSTEM_CURSOR_SIZEALL
	SYSTEM_CURSOR_NO
	SYSTEM_CURSOR_HAND
	NUM_SYSTEM_CURSORS
)

type Cursor C.SDL_Cursor

func GetMouseFocus() *Window {
	return (*Window) (unsafe.Pointer(C.SDL_GetMouseFocus()))
}

func GetMouseState(x, y *int) uint32 {
	_x := (*C.int) (unsafe.Pointer(x))
	_y := (*C.int) (unsafe.Pointer(y))
	return (uint32) (C.SDL_GetMouseState(_x, _y))
}

func GetRelativeMouseState(x, y *int) uint32 {
	_x := (*C.int) (unsafe.Pointer(x))
	_y := (*C.int) (unsafe.Pointer(y))
	return (uint32) (C.SDL_GetRelativeMouseState(_x, _y))
}

func (window *Window) WarpMouseInWindow(x, y int) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_x := (C.int) (x)
	_y := (C.int) (y)
	C.SDL_WarpMouseInWindow(_window, _x, _y)
}

func SetRelativeMouseMode(enabled bool) int {
	_enabled := (C.SDL_bool) (btoi(enabled))
	return (int) (C.SDL_SetRelativeMouseMode(_enabled))
}

func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() > 0
}

func CreateCursor(data, mask *uint8, w, h, hot_x, hot_y int) *Cursor {
	_data := (*C.Uint8) (unsafe.Pointer(data))
	_mask := (*C.Uint8) (unsafe.Pointer(mask))
	_w := (C.int) (w)
	_h := (C.int) (h)
	_hot_x := (C.int) (hot_x)
	_hot_y := (C.int) (hot_y)
	return (*Cursor) (C.SDL_CreateCursor(_data, _mask, _w, _h, _hot_x, _hot_y))
}

func CreateColorCursor(surface *Surface, hot_x, hot_y int) *Cursor {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_hot_x := (C.int) (hot_x)
	_hot_y := (C.int) (hot_y)
	return (*Cursor) (C.SDL_CreateColorCursor(_surface, _hot_x, _hot_y))
}

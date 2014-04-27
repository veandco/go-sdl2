package sdl

// #include <SDL2/SDL.h>
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

const (
	BUTTON_LEFT   = 1
	BUTTON_MIDDLE = 2
	BUTTON_RIGHT  = 3
	BUTTON_X1     = 4
	BUTTON_X2     = 5
)

type Cursor C.SDL_Cursor
type SystemCursor C.SDL_SystemCursor

func GetMouseFocus() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetMouseFocus()))
}

func GetMouseState() (x, y int, state uint32) {
	var _x, _y C.int
	_state := uint32(C.SDL_GetMouseState(&_x, &_y))
	return int(_x), int(_y), _state
}

func GetRelativeMouseState() (x, y int, state uint32) {
	var _x, _y C.int
	_state := uint32(C.SDL_GetRelativeMouseState(&_x, &_y))
	return int(_x), int(_y), _state
}

func (window *Window) WarpMouseInWindow(x, y int) {
	_window := (*C.SDL_Window)(unsafe.Pointer(window))
	_x := (C.int)(x)
	_y := (C.int)(y)
	C.SDL_WarpMouseInWindow(_window, _x, _y)
}

func SetRelativeMouseMode(enabled bool) int {
	_enabled := (C.SDL_bool)(Btoi(enabled))
	return (int)(C.SDL_SetRelativeMouseMode(_enabled))
}

func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() > 0
}

func CreateCursor(data, mask *uint8, w, h, hot_x, hot_y int) *Cursor {
	_data := (*C.Uint8)(unsafe.Pointer(data))
	_mask := (*C.Uint8)(unsafe.Pointer(mask))
	_w := (C.int)(w)
	_h := (C.int)(h)
	_hot_x := (C.int)(hot_x)
	_hot_y := (C.int)(hot_y)
	return (*Cursor)(C.SDL_CreateCursor(_data, _mask, _w, _h, _hot_x, _hot_y))
}

func CreateColorCursor(surface *Surface, hot_x, hot_y int) *Cursor {
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_hot_x := (C.int)(hot_x)
	_hot_y := (C.int)(hot_y)
	return (*Cursor)(C.SDL_CreateColorCursor(_surface, _hot_x, _hot_y))
}

func CreateSystemCursor(id SystemCursor) *Cursor {
	_id := (C.SDL_SystemCursor)(id)
	return (*Cursor)(C.SDL_CreateSystemCursor(_id))
}

func SetCursor(cursor *Cursor) {
	_cursor := (*C.SDL_Cursor)(cursor)
	C.SDL_SetCursor(_cursor)
}

func GetCursor() *Cursor {
	return (*Cursor)(C.SDL_GetCursor())
}

func GetDefaultCursor() *Cursor {
	return (*Cursor)(C.SDL_GetDefaultCursor())
}

func FreeCursor(cursor *Cursor) {
	_cursor := (*C.SDL_Cursor)(cursor)
	C.SDL_FreeCursor(_cursor)
}

func ShowCursor(toggle int) int {
	_toggle := (C.int)(toggle)
	return (int)(C.SDL_ShowCursor(_toggle))
}

func Button(flag uint32) uint32 {
	return 1 << (flag - 1)
}

func ButtonLMask() uint32 {
	return Button(BUTTON_LEFT)
}

func ButtonMMask() uint32 {
	return Button(BUTTON_MIDDLE)
}

func ButtonRMask() uint32 {
	return Button(BUTTON_RIGHT)
}

func ButtonX1Mask() uint32 {
	return Button(BUTTON_X1)
}

func ButtonX2Mask() uint32 {
	return Button(BUTTON_X2)
}

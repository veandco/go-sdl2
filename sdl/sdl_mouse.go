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

func (c *Cursor) cptr() *C.SDL_Cursor {
    return (*C.SDL_Cursor)(unsafe.Pointer(c))
}

func (c SystemCursor) c() C.SDL_SystemCursor {
    return C.SDL_SystemCursor(c)
}

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
	C.SDL_WarpMouseInWindow(window.cptr(), C.int(x), C.int(y))
}

func SetRelativeMouseMode(enabled bool) int {
	return int(C.SDL_SetRelativeMouseMode(C.SDL_bool(Btoi(enabled))))
}

func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() > 0
}

func CreateCursor(data, mask *uint8, w, h, hotX, hotY int) *Cursor {
	_data := (*C.Uint8)(unsafe.Pointer(data))
	_mask := (*C.Uint8)(unsafe.Pointer(mask))
	return (*Cursor)(C.SDL_CreateCursor(_data, _mask, C.int(w), C.int(h), C.int(hotX), C.int(hotY)))
}

func CreateColorCursor(surface *Surface, hotX, hotY int) *Cursor {
	return (*Cursor)(C.SDL_CreateColorCursor(surface.cptr(), C.int(hotX), C.int(hotY)))
}

func CreateSystemCursor(id SystemCursor) *Cursor {
	return (*Cursor)(C.SDL_CreateSystemCursor(id.c()))
}

func SetCursor(cursor *Cursor) {
	C.SDL_SetCursor(cursor.cptr())
}

func GetCursor() *Cursor {
	return (*Cursor)(C.SDL_GetCursor())
}

func GetDefaultCursor() *Cursor {
	return (*Cursor)(C.SDL_GetDefaultCursor())
}

func FreeCursor(cursor *Cursor) {
	C.SDL_FreeCursor(cursor.cptr())
}

func ShowCursor(toggle int) int {
	return int(C.SDL_ShowCursor(C.int(toggle)))
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

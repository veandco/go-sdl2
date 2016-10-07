package sdl

/*
#include "sdl_wrapper.h"

#if defined(__WIN32)
#include <SDL2/SDL_syswm.h>
#else
#include <SDL_syswm.h>
#endif

#if !(SDL_VERSION_ATLEAST(2,0,4))
#pragma message("SDL_CaptureMouse is not supported before SDL 2.0.4")
static int SDL_CaptureMouse(SDL_bool enabled)
{
	return -1;
}
#endif
*/
import "C"
import "unsafe"

/* SystemCursor */
const (
	SYSTEM_CURSOR_ARROW     = C.SDL_SYSTEM_CURSOR_ARROW
	SYSTEM_CURSOR_IBEAM     = C.SDL_SYSTEM_CURSOR_IBEAM
	SYSTEM_CURSOR_WAIT      = C.SDL_SYSTEM_CURSOR_WAIT
	SYSTEM_CURSOR_CROSSHAIR = C.SDL_SYSTEM_CURSOR_CROSSHAIR
	SYSTEM_CURSOR_WAITARROW = C.SDL_SYSTEM_CURSOR_WAITARROW
	SYSTEM_CURSOR_SIZENWSE  = C.SDL_SYSTEM_CURSOR_SIZENWSE
	SYSTEM_CURSOR_SIZENESW  = C.SDL_SYSTEM_CURSOR_SIZENESW
	SYSTEM_CURSOR_SIZEWE    = C.SDL_SYSTEM_CURSOR_SIZEWE
	SYSTEM_CURSOR_SIZENS    = C.SDL_SYSTEM_CURSOR_SIZENS
	SYSTEM_CURSOR_SIZEALL   = C.SDL_SYSTEM_CURSOR_SIZEALL
	SYSTEM_CURSOR_NO        = C.SDL_SYSTEM_CURSOR_NO
	SYSTEM_CURSOR_HAND      = C.SDL_SYSTEM_CURSOR_HAND
	NUM_SYSTEM_CURSORS      = C.SDL_NUM_SYSTEM_CURSORS
)

const (
	BUTTON_LEFT   = C.SDL_BUTTON_LEFT
	BUTTON_MIDDLE = C.SDL_BUTTON_MIDDLE
	BUTTON_RIGHT  = C.SDL_BUTTON_RIGHT
	BUTTON_X1     = C.SDL_BUTTON_X1
	BUTTON_X2     = C.SDL_BUTTON_X2
)

type Cursor C.SDL_Cursor
type SystemCursor C.SDL_SystemCursor

func (c *Cursor) cptr() *C.SDL_Cursor {
	return (*C.SDL_Cursor)(unsafe.Pointer(c))
}

func (c SystemCursor) c() C.SDL_SystemCursor {
	return C.SDL_SystemCursor(c)
}

// GetMouseFocus (https://wiki.libsdl.org/SDL_GetMouseFocus)
func GetMouseFocus() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetMouseFocus()))
}

// GetMouseState (https://wiki.libsdl.org/SDL_GetMouseState)
func GetMouseState() (x, y int, state uint32) {
	var _x, _y C.int
	_state := uint32(C.SDL_GetMouseState(&_x, &_y))
	return int(_x), int(_y), _state
}

// GetRelativeMouseState (https://wiki.libsdl.org/SDL_GetRelativeMouseState)
func GetRelativeMouseState() (x, y int, state uint32) {
	var _x, _y C.int
	_state := uint32(C.SDL_GetRelativeMouseState(&_x, &_y))
	return int(_x), int(_y), _state
}

// Window (https://wiki.libsdl.org/SDL_WarpMouseInWindow)
func (window *Window) WarpMouseInWindow(x, y int) {
	C.SDL_WarpMouseInWindow(window.cptr(), C.int(x), C.int(y))
}

// SetRelativeMouseMode (https://wiki.libsdl.org/SDL_SetRelativeMouseMode)
func SetRelativeMouseMode(enabled bool) int {
	return int(C.SDL_SetRelativeMouseMode(C.SDL_bool(Btoi(enabled))))
}

// GetRelativeMouseMode (https://wiki.libsdl.org/SDL_GetRelativeMouseMode)
func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() > 0
}

// CreateCursor (https://wiki.libsdl.org/SDL_CreateCursor)
func CreateCursor(data, mask *uint8, w, h, hotX, hotY int) *Cursor {
	_data := (*C.Uint8)(unsafe.Pointer(data))
	_mask := (*C.Uint8)(unsafe.Pointer(mask))
	return (*Cursor)(C.SDL_CreateCursor(_data, _mask, C.int(w), C.int(h), C.int(hotX), C.int(hotY)))
}

// CreateColorCursor (https://wiki.libsdl.org/SDL_CreateColorCursor)
func CreateColorCursor(surface *Surface, hotX, hotY int) *Cursor {
	return (*Cursor)(C.SDL_CreateColorCursor(surface.cptr(), C.int(hotX), C.int(hotY)))
}

// CreateSystemCursor (https://wiki.libsdl.org/SDL_CreateSystemCursor)
func CreateSystemCursor(id SystemCursor) *Cursor {
	return (*Cursor)(C.SDL_CreateSystemCursor(id.c()))
}

// SetCursor (https://wiki.libsdl.org/SDL_SetCursor)
func SetCursor(cursor *Cursor) {
	C.SDL_SetCursor(cursor.cptr())
}

// GetCursor (https://wiki.libsdl.org/SDL_GetCursor)
func GetCursor() *Cursor {
	return (*Cursor)(C.SDL_GetCursor())
}

// GetDefaultCursor (https://wiki.libsdl.org/SDL_GetDefaultCursor)
func GetDefaultCursor() *Cursor {
	return (*Cursor)(C.SDL_GetDefaultCursor())
}

// FreeCursor (https://wiki.libsdl.org/SDL_FreeCursor)
func FreeCursor(cursor *Cursor) {
	C.SDL_FreeCursor(cursor.cptr())
}

// ShowCursor (https://wiki.libsdl.org/SDL_ShowCursor)
func ShowCursor(toggle int) int {
	return int(C.SDL_ShowCursor(C.int(toggle)))
}

// CaptureMouse (https://wiki.libsdl.org/SDL_CaptureMouse)
func CaptureMouse(toggle bool) error {
	var ierr C.int
	if toggle {
		ierr = C.SDL_CaptureMouse(C.SDL_TRUE)
	} else {
		ierr = C.SDL_CaptureMouse(C.SDL_FALSE)
	}
	if ierr != 0 {
		return GetError()
	}
	return nil
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

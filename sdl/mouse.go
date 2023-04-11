package sdl

/*
#include "sdl_wrapper.h"

#if defined(__WIN32)
#include <SDL2/SDL_syswm.h>
#else
#include <SDL_syswm.h>
#endif

#if !(SDL_VERSION_ATLEAST(2,0,4))

#if defined(WARN_OUTDATED)
#pragma message("SDL_CaptureMouse is not supported before SDL 2.0.4")
#endif

static int SDL_CaptureMouse(SDL_bool enabled)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_MOUSEWHEEL_NORMAL is not supported before SDL 2.0.4")
#endif

#define SDL_MOUSEWHEEL_NORMAL (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_MOUSEWHEEL_FLIPPED is not supported before SDL 2.0.4")
#endif

#define SDL_MOUSEWHEEL_FLIPPED (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WarpMouseGlobal is not supported before SDL 2.0.4")
#endif

static int SDL_WarpMouseGlobal(int x, int y)
{
	return -1;
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetGlobalMouseState is not supported before SDL 2.0.4")
#endif

static Uint32 SDL_GetGlobalMouseState(int *x, int *y)
{
	return 0;
}

#endif
*/
import "C"
import "unsafe"

// Cursor types for CreateSystemCursor()
// (https://wiki.libsdl.org/SDL_SystemCursor)
type SystemCursor C.SDL_SystemCursor

const (
	SYSTEM_CURSOR_ARROW     SystemCursor = C.SDL_SYSTEM_CURSOR_ARROW     // arrow
	SYSTEM_CURSOR_IBEAM     SystemCursor = C.SDL_SYSTEM_CURSOR_IBEAM     // i-beam
	SYSTEM_CURSOR_WAIT      SystemCursor = C.SDL_SYSTEM_CURSOR_WAIT      // wait
	SYSTEM_CURSOR_CROSSHAIR SystemCursor = C.SDL_SYSTEM_CURSOR_CROSSHAIR // crosshair
	SYSTEM_CURSOR_WAITARROW SystemCursor = C.SDL_SYSTEM_CURSOR_WAITARROW // small wait cursor (or wait if not available)
	SYSTEM_CURSOR_SIZENWSE  SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENWSE  // double arrow pointing northwest and southeast
	SYSTEM_CURSOR_SIZENESW  SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENESW  // double arrow pointing northeast and southwest
	SYSTEM_CURSOR_SIZEWE    SystemCursor = C.SDL_SYSTEM_CURSOR_SIZEWE    // double arrow pointing west and east
	SYSTEM_CURSOR_SIZENS    SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENS    // double arrow pointing north and south
	SYSTEM_CURSOR_SIZEALL   SystemCursor = C.SDL_SYSTEM_CURSOR_SIZEALL   // four pointed arrow pointing north, south, east, and west
	SYSTEM_CURSOR_NO        SystemCursor = C.SDL_SYSTEM_CURSOR_NO        // slashed circle or crossbones
	SYSTEM_CURSOR_HAND      SystemCursor = C.SDL_SYSTEM_CURSOR_HAND      // hand
	NUM_SYSTEM_CURSORS      SystemCursor = C.SDL_NUM_SYSTEM_CURSORS      // (only for bounding internal arrays)
)

// Scroll direction types for the Scroll event
const (
	MOUSEWHEEL_NORMAL  = C.SDL_MOUSEWHEEL_NORMAL  // the scroll direction is normal
	MOUSEWHEEL_FLIPPED = C.SDL_MOUSEWHEEL_FLIPPED // the scroll direction is flipped / natural
)

// Mouse buttons
type Button uint32

const (
	ButtonLeft   Button = C.SDL_BUTTON_LEFT   // left mouse button
	ButtonMiddle Button = C.SDL_BUTTON_MIDDLE // middle mouse button
	ButtonRight  Button = C.SDL_BUTTON_RIGHT  // right mouse button
	ButtonX1     Button = C.SDL_BUTTON_X1     // x1 mouse button
	ButtonX2     Button = C.SDL_BUTTON_X2     // x2 mouse button
)

func (b Button) Mask() ButtonMask {
	return ButtonMask(1 << (b - 1))
}

// Used as a mask when testing buttons in buttonstate.
type ButtonMask uint32

const (
	ButtonLMask  ButtonMask = 1 << (ButtonLeft - 1)   // left mouse button mask
	ButtonMMask  ButtonMask = 1 << (ButtonMiddle - 1) // middle mouse button mask
	ButtonRMask  ButtonMask = 1 << (ButtonRight - 1)  // right mouse button mask
	ButtonX1Mask ButtonMask = 1 << (ButtonX1 - 1)     // x1 mouse button mask
	ButtonX2Mask ButtonMask = 1 << (ButtonX2 - 1)     // x2 mouse button mask
)

func (m ButtonMask) Has(b Button) bool {
	return m&b.Mask() != 0
}

// Cursor is a custom cursor created by CreateCursor() or CreateColorCursor().
type Cursor C.SDL_Cursor

func (c *Cursor) cptr() *C.SDL_Cursor {
	return (*C.SDL_Cursor)(unsafe.Pointer(c))
}

func (c SystemCursor) c() C.SDL_SystemCursor {
	return C.SDL_SystemCursor(c)
}

// GetMouseFocus returns the window which currently has mouse focus.
// (https://wiki.libsdl.org/SDL_GetMouseFocus)
func GetMouseFocus() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetMouseFocus()))
}

// GetGlobalMouseState returns the current state of the mouse.
// (https://wiki.libsdl.org/SDL_GetGlobalMouseState)
func GetGlobalMouseState() (x, y int32, state ButtonMask) {
	var _x, _y C.int
	_state := ButtonMask(C.SDL_GetGlobalMouseState(&_x, &_y))
	return int32(_x), int32(_y), _state
}

// GetMouseState returns the current state of the mouse.
// (https://wiki.libsdl.org/SDL_GetMouseState)
func GetMouseState() (x, y int32, state ButtonMask) {
	var _x, _y C.int
	_state := ButtonMask(C.SDL_GetMouseState(&_x, &_y))
	return int32(_x), int32(_y), _state
}

// GetRelativeMouseState returns the relative state of the mouse.
// (https://wiki.libsdl.org/SDL_GetRelativeMouseState)
func GetRelativeMouseState() (x, y int32, state ButtonMask) {
	var _x, _y C.int
	_state := ButtonMask(C.SDL_GetRelativeMouseState(&_x, &_y))
	return int32(_x), int32(_y), _state
}

// WarpMouseInWindow moves the mouse to the given position within the window.
// (https://wiki.libsdl.org/SDL_WarpMouseInWindow)
func (window *Window) WarpMouseInWindow(x, y int32) {
	C.SDL_WarpMouseInWindow(window.cptr(), C.int(x), C.int(y))
}

// SetRelativeMouseMode sets relative mouse mode.
// (https://wiki.libsdl.org/SDL_SetRelativeMouseMode)
func SetRelativeMouseMode(enabled bool) int {
	return int(C.SDL_SetRelativeMouseMode(C.SDL_bool(Btoi(enabled))))
}

// GetRelativeMouseMode reports where relative mouse mode is enabled.
// (https://wiki.libsdl.org/SDL_GetRelativeMouseMode)
func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() > 0
}

// CreateCursor creates a cursor using the specified bitmap data and mask (in MSB format).
// (https://wiki.libsdl.org/SDL_CreateCursor)
func CreateCursor(data, mask *uint8, w, h, hotX, hotY int32) *Cursor {
	_data := (*C.Uint8)(unsafe.Pointer(data))
	_mask := (*C.Uint8)(unsafe.Pointer(mask))
	return (*Cursor)(C.SDL_CreateCursor(_data, _mask, C.int(w), C.int(h), C.int(hotX), C.int(hotY)))
}

// CreateColorCursor creates a color cursor.
// (https://wiki.libsdl.org/SDL_CreateColorCursor)
func CreateColorCursor(surface *Surface, hotX, hotY int32) *Cursor {
	return (*Cursor)(C.SDL_CreateColorCursor(surface.cptr(), C.int(hotX), C.int(hotY)))
}

// CreateSystemCursor creates a system cursor.
// (https://wiki.libsdl.org/SDL_CreateSystemCursor)
func CreateSystemCursor(id SystemCursor) *Cursor {
	return (*Cursor)(C.SDL_CreateSystemCursor(id.c()))
}

// SetCursor sets the active cursor.
// (https://wiki.libsdl.org/SDL_SetCursor)
func SetCursor(cursor *Cursor) {
	C.SDL_SetCursor(cursor.cptr())
}

// GetCursor returns the active cursor.
// (https://wiki.libsdl.org/SDL_GetCursor)
func GetCursor() *Cursor {
	return (*Cursor)(C.SDL_GetCursor())
}

// GetDefaultCursor returns the default cursor.
// (https://wiki.libsdl.org/SDL_GetDefaultCursor)
func GetDefaultCursor() *Cursor {
	return (*Cursor)(C.SDL_GetDefaultCursor())
}

// FreeCursor frees a cursor created with CreateCursor(), CreateColorCursor() or CreateSystemCursor().
// (https://wiki.libsdl.org/SDL_FreeCursor)
func FreeCursor(cursor *Cursor) {
	C.SDL_FreeCursor(cursor.cptr())
}

// ShowCursor toggles whether or not the cursor is shown.
// (https://wiki.libsdl.org/SDL_ShowCursor)
func ShowCursor(toggle EventStateConstant) (int, error) {
	i := int(C.SDL_ShowCursor(C.int(toggle)))
	return i, errorFromInt(i)
}

// CaptureMouse captures the mouse and tracks input outside an SDL window.
// (https://wiki.libsdl.org/SDL_CaptureMouse)
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

// WarpMouseGlobal moves the mouse to the given position in global screen space.
// (https://wiki.libsdl.org/SDL_WarpMouseGlobal)
func WarpMouseGlobal(x, y int32) error {
	i := int(C.SDL_WarpMouseGlobal(C.int(x), C.int(y)))
	return errorFromInt(i)
}

package sdl

/*
#include "sdl_wrapper.h"

// until SDL 2.0.6 SDL_shape.h was not included in SDL.h
#if !(SDL_VERSION_ATLEAST(2,0,6))
#if defined(_WIN32)
	#include <SDL2/SDL_shape.h>
#else
	#include <SDL_shape.h>
#endif
#endif
*/
import "C"
import "unsafe"

const (
	NONSHAPEABLE_WINDOW    = C.SDL_NONSHAPEABLE_WINDOW
	INVALID_SHAPE_ARGUMENT = C.SDL_INVALID_SHAPE_ARGUMENT
	WINDOW_LACKS_SHAPE     = C.SDL_WINDOW_LACKS_SHAPE
)

type WindowShapeModeKind C.WindowShapeMode

const (
	ShapeModeDefaultKind              WindowShapeModeKind = C.ShapeModeDefault              // the default mode, a binarized alpha cutoff of 1
	ShapeModeBinarizeAlphaKind        WindowShapeModeKind = C.ShapeModeBinarizeAlpha        // a binarized alpha cutoff with a given integer value
	ShapeModeReverseBinarizeAlphaKind WindowShapeModeKind = C.ShapeModeReverseBinarizeAlpha // a binarized alpha cutoff with a given integer value, but with the opposite comparison
	ShapeModeColorKeyKind             WindowShapeModeKind = C.ShapeModeColorKey             // a color key is applied
)

func SHAPEMODEALPHA(mode WindowShapeModeKind) bool {
	return (mode == ShapeModeDefaultKind ||
		mode == ShapeModeBinarizeAlphaKind ||
		mode == ShapeModeReverseBinarizeAlphaKind)
}

// A union containing parameters for shaped windows
// uint8 or Color
type cWindowShapeParams C.SDL_WindowShapeParams

type cWindowShapeMode struct {
	mode       WindowShapeModeKind
	parameters cWindowShapeParams
}

type WindowShapeMode interface {
	Mode() WindowShapeModeKind
	cWSM() cWindowShapeMode
}

type ShapeModeDefault struct{}

func (smd ShapeModeDefault) Mode() WindowShapeModeKind {
	return ShapeModeDefaultKind
}

func (smd ShapeModeDefault) cWSM() cWindowShapeMode {
	return cWindowShapeMode{
		mode:       ShapeModeDefaultKind,
		parameters: [4]uint8{1, 0, 0, 0},
	}
}

type ShapeModeBinarizeAlpha struct {
	Cutoff uint8
}

func (smba ShapeModeBinarizeAlpha) Mode() WindowShapeModeKind {
	return ShapeModeBinarizeAlphaKind
}

func (smba ShapeModeBinarizeAlpha) cWSM() cWindowShapeMode {
	return cWindowShapeMode{
		mode:       ShapeModeBinarizeAlphaKind,
		parameters: [4]uint8{smba.Cutoff, 0, 0, 0},
	}
}

type ShapeModeReverseBinarizeAlpha struct {
	Cutoff uint8
}

func (smba ShapeModeReverseBinarizeAlpha) Mode() WindowShapeModeKind {
	return ShapeModeReverseBinarizeAlphaKind
}

func (smba ShapeModeReverseBinarizeAlpha) cWSM() cWindowShapeMode {
	return cWindowShapeMode{
		mode:       ShapeModeReverseBinarizeAlphaKind,
		parameters: [4]uint8{smba.Cutoff, 0, 0, 0},
	}
}

type ShapeModeColorKey struct {
	Color Color
}

func (smck ShapeModeColorKey) Mode() WindowShapeModeKind {
	return ShapeModeColorKeyKind
}

func (smck ShapeModeColorKey) cWSM() cWindowShapeMode {
	return cWindowShapeMode{
		mode:       ShapeModeReverseBinarizeAlphaKind,
		parameters: [4]uint8{smck.Color.R, smck.Color.G, smck.Color.B, smck.Color.A},
	}
}

func (cwsm cWindowShapeMode) goWSM() WindowShapeMode {
	switch cwsm.mode {
	case ShapeModeDefaultKind:
		return ShapeModeDefault{}

	case ShapeModeBinarizeAlphaKind:
		return ShapeModeBinarizeAlpha{
			Cutoff: ([4]uint8)(cwsm.parameters)[0],
		}

	case ShapeModeReverseBinarizeAlphaKind:
		return ShapeModeReverseBinarizeAlpha{
			Cutoff: ([4]uint8)(cwsm.parameters)[0],
		}

	case ShapeModeColorKeyKind:
		return ShapeModeColorKey{
			Color: Color{
				R: ([4]uint8)(cwsm.parameters)[0],
				G: ([4]uint8)(cwsm.parameters)[1],
				B: ([4]uint8)(cwsm.parameters)[2],
				A: ([4]uint8)(cwsm.parameters)[3],
			},
		}

	default:
		panic("Unknown WindowShapeModeKind")

	}
}

func (wsm *cWindowShapeMode) cptr() *C.SDL_WindowShapeMode {
	return (*C.SDL_WindowShapeMode)(unsafe.Pointer(wsm))
}

// CreateShapedWindow creates a window that can be shaped with the specified position, dimensions, and flags
// (https://wiki.libsdl.org/SDL_CreateShapedWindow)
func CreateShapedWindow(title string, x, y, w, h uint32, flags uint32) (*Window, error) {
	var _window = C.SDL_CreateShapedWindow(C.CString(title), C.uint(x), C.uint(y), C.uint(w), C.uint(h), C.Uint32(flags))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// IsShapedWindow returns whether the given window is a shaped window.
// (https://wiki.libsdl.org/SDL_IsShapedWindow)
func (window *Window) IsShaped() bool {
	return (C.SDL_IsShapedWindow(window.cptr()) & 1) == 1
}

// SetShape sets the shape and parameters of a shaped window
// (https://wiki.libsdl.org/SDL_SetWindowShape)
func (window *Window) SetShape(shape *Surface, shape_mode WindowShapeMode) int32 {
	if shape_mode == nil {
		panic("shape_mode can not be nil")
	}
	var _cWSM cWindowShapeMode = shape_mode.cWSM()

	return (int32)(C.SDL_SetWindowShape(window.cptr(), shape.cptr(), _cWSM.cptr()))
}

// GetShapeMode gets the shape parameters of a shaped window
// (https://wiki.libsdl.org/SDL_GetShapedWindowMode)
func (window *Window) GetShapeMode() (WindowShapeMode, int32) {
	var _cWSM cWindowShapeMode
	var _resInt32 = (int32)(C.SDL_GetShapedWindowMode(window.cptr(), _cWSM.cptr()))
	if _resInt32 != 0 {
		return nil, _resInt32
	}
	return _cWSM.goWSM(), _resInt32
}

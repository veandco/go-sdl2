// Package gfx is an add-on for SDL2 that provides drawing of graphics primitives, rotozoomer, MMX image filters and framerate control.
package gfx

//#include <stdlib.h>
//#include "sdl_gfx_wrapper.h"
import "C"
import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// FPS definitions.
const (
	FPS_UPPER_LIMIT = 200 // highest possible rate supported by framerate controller in Hz (1/s)
	FPS_LOWER_LIMIT = 1   // lowest possible rate supported by framerate controller in Hz (1/s)
	FPS_DEFAULT     = 30  // default rate of framerate controller in Hz (1/s)
)

// FPSmanager contains the state and timing information of the framerate controller.
type FPSmanager struct {
	FrameCount uint32
	RateTicks  float32
	BaseTicks  uint32
	LastTicks  uint32
	Rate       uint32
}

// Rotozoom smoothing settings.
const (
	SMOOTHING_OFF = 0 // disable anti-aliasing (no smoothing)
	SMOOTHING_ON  = 1 // enable anti-aliasing (smoothing)
)

func min(a ...int) int {
	b := a[0]

	for i := 1; i < len(a); i++ {
		if b > a[i] {
			b = a[i]
		}
	}

	return b
}

func gfxColor(color sdl.Color) uint32 {
	return (uint32(color.A) << 24) | (uint32(color.B) << 16) | (uint32(color.G) << 8) | uint32(color.R)
}

func (manager *FPSmanager) cptr() *C.FPSmanager {
	return (*C.FPSmanager)(unsafe.Pointer(manager))
}

// InitFramerate initializes the framerate manager.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__framerate_8h.html#a843f0672446aff01ef03bbcd977fbedf)
func InitFramerate(manager *FPSmanager) {
	C.SDL_initFramerate(manager.cptr())
}

// SetFramerate sets the framerate in Hz.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__framerate_8h.html#aa2f7f11e60d81489392707faf07c5ac5)
func SetFramerate(manager *FPSmanager, rate uint32) bool {
	_rate := C.Uint32(rate)
	return C.SDL_setFramerate(manager.cptr(), _rate) == 0
}

// GetFramerate returns the current target framerate in Hz.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__framerate_8h.html#aebe43457dbb9fbfec6de18e7adf49e21)
func GetFramerate(manager *FPSmanager) (int, bool) {
	fps := int(C.SDL_getFramerate(manager.cptr()))
	return fps, fps >= 0
}

// GetFramecount returns the current framecount.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__framerate_8h.html#a41de3f516b9633f102d7ba2e85d2bb98)
func GetFramecount(manager *FPSmanager) (int, bool) {
	count := int(C.SDL_getFramecount(manager.cptr()))
	return count, count >= 0
}

// FramerateDelay delays execution to maintain a constant framerate and calculate fps.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__framerate_8h.html#afc3b999d59d913771cd2588299096274)
func FramerateDelay(manager *FPSmanager) uint32 {
	return uint32(C.SDL_framerateDelay(manager.cptr()))
}

// PixelColor pixel draws with blending enabled if a<255.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a2bdb83ac58c2e091ef26f720bdeb66bd)
func PixelColor(renderer *sdl.Renderer, x, y int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.pixelColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _color) == 0
}

// PixelRGBA pixel draws with blending enabled if a<255.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a6f4a05ba92bf39420280ee8ccf961e77)
func PixelRGBA(renderer *sdl.Renderer, x, y int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.pixelRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _r, _g, _b, _a) == 0
}

// HlineColor draws horizontal line with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a0d1977c09c0bcfe77b4e1a792ce1e79a)
func HlineColor(renderer *sdl.Renderer, x1, x2, y int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.hlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _color) == 0
}

// HlineRGBA draws horizontal line with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#ace5a96b77edbccba3170852b093503aa)
func HlineRGBA(renderer *sdl.Renderer, x1, x2, y int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.hlineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _r, _g, _b, _a) == 0
}

// VlineColor draws vertical line with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#af8cf6591ddc81242369e50cb3e159dfe)
func VlineColor(renderer *sdl.Renderer, x, y1, y2 int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.vlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _color) == 0
}

// VlineRGBA draws vertical line with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a60d71d2bc6e450d8063256ebc37f21f5)
func VlineRGBA(renderer *sdl.Renderer, x, y1, y2 int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.vlineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _r, _g, _b, _a) == 0
}

// RectangleColor draws rectangle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a53b996645c0acdf5cb74c969503e791e)
func RectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.rectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// RectangleRGBA draws rectangle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#aa926924c3650d10d6a20cd7e4036a9a4)
func RectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.rectangleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

// RoundedRectangleColor draws rounded-corner rectangle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a314698defb7da3f3415ddc5468315d83)
func RoundedRectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.roundedRectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _color) == 0
}

// RoundedRectangleRGBA draws rounded-corner rectangle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a2b54ba16d7e7243fb1921a286156cad9)
func RoundedRectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.roundedRectangleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _r, _g, _b, _a) == 0
}

// BoxColor draws box (filled rectangle) with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a5799b0d1e99252b19d9acc77e7a94541)
func BoxColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.boxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// BoxRGBA draws box (filled rectangle) with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a903bb79c5a03077047ff8820df263bd8)
func BoxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.boxRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

// RoundedBoxColor draws rounded-corner box (filled rectangle) with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a13ffcecd9c186d2522ae1e66bdedf8d7)
func RoundedBoxColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.roundedBoxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _color) == 0
}

// RoundedBoxRGBA draws rounded-corner box (filled rectangle) with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a3db064ca0f03a2c46852661494ff7a65)
func RoundedBoxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.roundedBoxRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _r, _g, _b, _a) == 0
}

// LineColor draws line with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a84f917eb19ea41b586f1afa9eb4b552c)
func LineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.lineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// LineRGBA draws line with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#ab56ec3fb82b59f2ab1d1877b6adb3b82)
func LineRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.lineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

// AALineColor draws anti-aliased line with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a86ab777f53655a509a3ac3c008941920)
func AALineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.aalineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// AALineRGBA draws anti-aliased line with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#ad7074bc1af414cea003712621b1c7d86)
func AALineRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aalineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

// ThickLineColor draws a thick line with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a613c16293c3571b695cc0c60358bb862)
func ThickLineColor(renderer *sdl.Renderer, x1, y1, x2, y2, width int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_width := C.Uint8(width)
	_color := C.Uint32(gfxColor(color))
	return C.thickLineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _width, _color) == 0
}

// ThickLineRGBA draws a thick line with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a247136a562abec2649718d38f5819b44)
func ThickLineRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, width int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_width := C.Uint8(width)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.thickLineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _width, _r, _g, _b, _a) == 0
}

// CircleColor draws circle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#af68b52c59d2a7a6b7cc817158941ac54)
func CircleColor(renderer *sdl.Renderer, x, y, rad int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.circleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

// CircleRGBA draws circle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a1daab46976bf585477be1cfcef2fe1ad)
func CircleRGBA(renderer *sdl.Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.circleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

// ArcColor draws arc with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a493fe37b39da582668431ab7e1e3dbf0)
func ArcColor(renderer *sdl.Renderer, x, y, rad, start, end int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.arcColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

// ArcRGBA draws arc with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#acf0091a17501f375f2b55032134b3017)
func ArcRGBA(renderer *sdl.Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.arcRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

// AACircleColor draws anti-aliased circle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a46b037505b91133ae3bd18556092a632)
func AACircleColor(renderer *sdl.Renderer, x, y, rad int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.aacircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

// AACircleRGBA draws anti-aliased circle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#aa0b0f764826715353e9ca0fe937d0f0f)
func AACircleRGBA(renderer *sdl.Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aacircleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

// FilledCircleColor draws filled circle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#ab973fd9868527fb13078c356d4a9c6f7)
func FilledCircleColor(renderer *sdl.Renderer, x, y, rad int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.filledCircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

// FilledCircleRGBA draws filled circle with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a666bd764e2fe962656e5829d0aad5ba6)
func FilledCircleRGBA(renderer *sdl.Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledCircleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

// EllipseColor draws ellipse with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a9ac841634751689ddb1c26babe10f3f6)
func EllipseColor(renderer *sdl.Renderer, x, y, rx, ry int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.ellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

// EllipseRGBA draws ellipse with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a3ed26f8b2a25cb94a6412aa772f533aa)
func EllipseRGBA(renderer *sdl.Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.ellipseRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

// AAEllipseColor draws anti-aliased ellipse with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#af676a520e9ea0deabe711da842bc1e55)
func AAEllipseColor(renderer *sdl.Renderer, x, y, rx, ry int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.aaellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

// AAEllipseRGBA draws anti-aliased ellipse with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a12e9ff795a5b9996f07f5b0bc4f60f81)
func AAEllipseRGBA(renderer *sdl.Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aaellipseRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

// FilledEllipseColor draws filled ellipse with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#aa3119e2dd056874c0ca92ded41fea43c)
func FilledEllipseColor(renderer *sdl.Renderer, x, y, rx, ry int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.filledEllipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

// FilledEllipseRGBA draws filled ellipse with blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a5240918c243c3e60dd8ae1cef50dd529)
func FilledEllipseRGBA(renderer *sdl.Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledEllipseRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

// PieColor draws pie (outline) with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a9f9070739e72ea242fd48239a4ff48e4)
func PieColor(renderer *sdl.Renderer, x, y, rad, start, end int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.pieColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

// PieRGBA draws pie (outline) with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a3b8baf3eecea4738ac2a1c28af4bfb41)
func PieRGBA(renderer *sdl.Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.pieRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

// FilledPieColor draws filled pie with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a86f189dabaa2a26115ecb819ad7da8e5)
func FilledPieColor(renderer *sdl.Renderer, x, y, rad, start, end int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.filledPieColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

// FilledPieRGBA draws filled pie with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#adad0f423ae093b5c3cad51a438954a50)
func FilledPieRGBA(renderer *sdl.Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledPieRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

// TrigonColor draws trigon (triangle outline) with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#af4fc0fdb33e57047d9347d3fe7b1bdbd)
func TrigonColor(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(gfxColor(color))
	return C.trigonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
}

// TrigonRGBA draws trigon (triangle outline) with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a152662f6985587d137837086aaa95311)
func TrigonRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.trigonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _r, _g, _b, _a) == 0
}

// FilledTrigonColor draws filled trigon (triangle) with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a57fd7355780a114b9d72b1e9591f955d)
func FilledTrigonColor(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(gfxColor(color))
	return C.filledTrigonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
}

// FilledTrigonRGBA draws filled trigon (triangle) with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a273cf4a88abf6c6a5e019b2c58ee2423)
func FilledTrigonRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledTrigonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _r, _g, _b, _a) == 0
}

// PolygonColor draws polygon with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a3577e5aee374d5176a3587343e11c4aa)
func PolygonColor(renderer *sdl.Renderer, vx, vy []int16, color sdl.Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_color := C.Uint32(gfxColor(color))
	return C.polygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

// PolygonRGBA draws polygon with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a69baa68215840bb54ddd0281e6ad63a0)
func PolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.polygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

// AAPolygonColor draws anti-aliased polygon with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a54c28d1e469b52483e952258c2f6f69b)
func AAPolygonColor(renderer *sdl.Renderer, vx, vy []int16, color sdl.Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_color := C.Uint32(gfxColor(color))
	return C.aapolygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

// AAPolygonRGBA draws anti-aliased polygon with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a6a18d7838305c636b8462e0ba587d859)
func AAPolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aapolygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

// FilledPolygonColor draws filled polygon with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a97308e4f19363baee0e5bdecee3265f1)
func FilledPolygonColor(renderer *sdl.Renderer, vx, vy []int16, color sdl.Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_color := C.Uint32(gfxColor(color))
	return C.filledPolygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

// FilledPolygonRGBA draws filled polygon with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#ab80fe08abcf78a6ce8f037f3fdc15625)
func FilledPolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledPolygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

// TexturedPolygon draws a polygon filled with the given texture.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a3b4c81592098da506f534d4a81b266f2)
func TexturedPolygon(renderer *sdl.Renderer, vx, vy []int16, surface *sdl.Surface, textureDX, textureDY int) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_textureDX := C.int(textureDX)
	_textureDY := C.int(textureDY)
	return C.texturedPolygon((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _surface, _textureDX, _textureDY) == 0
}

// BezierColor draws a bezier curve with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a74d59d17ba21ce5d8c84697989456bfb)
func BezierColor(renderer *sdl.Renderer, vx, vy []int16, s int, color sdl.Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_s := C.int(s)
	_color := C.Uint32(gfxColor(color))
	return C.bezierColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _s, _color) == 0
}

// BezierRGBA draws a bezier curve with alpha blending.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a6cb082e6eb4253d591927c6bf7eba06f)
func BezierRGBA(renderer *sdl.Renderer, vx, vy []int16, s int, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_s := C.int(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.bezierRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _s, _r, _g, _b, _a) == 0
}

// SetFont sets or resets the current global font data. The font data array is organized in follows: [fontdata] = [character 0][character 1]...[character 255] where [character n] = [byte 1 row 1][byte 2 row 1]...[byte {pitch} row 1][byte 1 row 2] ...[byte {pitch} row height] where [byte n] = [bit 0]...[bit 7] where [bit n] = [0 for transparent pixel|1 for colored pixel].
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#adcbe6b74a9e0fef4165bc7bdb73294ec)
func SetFont(fontdata []byte, cw, ch uint32) {
	_fontdata := unsafe.Pointer(nil)
	if fontdata != nil {
		_fontdata = unsafe.Pointer(&fontdata[0])
	}
	_cw := C.Uint32(cw)
	_ch := C.Uint32(ch)
	C.gfxPrimitivesSetFont(_fontdata, _cw, _ch)
}

// SetFontRotation sets current global font character rotation steps. Default is 0 (no rotation). 1 = 90deg clockwise. 2 = 180deg clockwise. 3 = 270deg clockwise. Changing the rotation, will reset the character cache.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a2349bb02995a364f6b1069c46a09b7ba)
func SetFontRotation(rotation uint32) {
	_rotation := C.Uint32(rotation)
	C.gfxPrimitivesSetFontRotation(_rotation)
}

// CharacterColor draws a character of the currently set font. On first call for a particular character and color combination, the function needs to generate the character surface (slower). Subsequent calls blit a cached surface (fast). Uses alpha blending if A<255 in color.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#aaeed3ccb288e032856cff488bdba381d)
func CharacterColor(renderer *sdl.Renderer, x, y int32, c byte, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_color := C.Uint32(gfxColor(color))
	return C.characterColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _color) == 0
}

// CharacterRGBA draws a character of the currently set font.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#abff6cc81b9a35e9380e0b02da905714f)
func CharacterRGBA(renderer *sdl.Renderer, x, y int32, c, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.characterRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _r, _g, _b, _a) == 0
}

// StringColor draws a string in the currently set font. The spacing between consequtive characters in the string is the fixed number of pixels of the character width of the current global font.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a329e9c81b07e5af08d73e3b3d6078bf0)
func StringColor(renderer *sdl.Renderer, x, y int32, s string, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_color := C.Uint32(gfxColor(color))
	return C.stringColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _color) == 0
}

// StringRGBA draws a string in the currently set font.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__gfx_primitives_8h.html#a744afe14139c1c553204a47cc284e478)
func StringRGBA(renderer *sdl.Renderer, x, y int32, s string, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.stringRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _r, _g, _b, _a) == 0
}

// ImageFilterMMXdetect reports whether MMX check for filter functions is enabled.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#ae3f67cbe712f604b16b6de3f4bfbf31c)
func ImageFilterMMXdetect() bool {
	return C.SDL_imageFilterMMXdetect() > 0
}

// ImageFilterMMXoff disables MMX check for filter functions and force to use non-MMX C based code.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#afc46d09d46b1302becfc170214dee0c0)
func ImageFilterMMXoff() {
	C.SDL_imageFilterMMXoff()
}

// ImageFilterMMXon enables MMX check for filter functions and use MMX code if available.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a0b1d8468dc6e6304b62276acbb7336f6)
func ImageFilterMMXon() {
	C.SDL_imageFilterMMXon()
}

// ImageFilterAdd filters using Add: D = saturation255(S1 + S2).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a1e4de9be8feb43595719fd0494601952)
func ImageFilterAdd(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterAdd(_src1, _src2, _dest, _len) == 0
}

// ImageFilterMean filters using Mean: D = S1/2 + S2/2.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a6012332e1b5c33fad53d71c7848db823)
func ImageFilterMean(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMean(_src1, _src2, _dest, _len) == 0
}

// ImageFilterSub filters using Sub: D = saturation0(S1 - S2).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a720893e0f6512aee4dd3875b9c9607b5)
func ImageFilterSub(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterSub(_src1, _src2, _dest, _len) == 0
}

// ImageFilterAbsDiff filters using AbsDiff: D = | S1 - S2 |.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#aaa9e8718bcba856ddee135385ebdec26)
func ImageFilterAbsDiff(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterAbsDiff(_src1, _src2, _dest, _len) == 0
}

// ImageFilterMult filters using Mult: D = saturation255(S1 * S2).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a1966f22bee81045917e776fd64821051)
func ImageFilterMult(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMult(_src1, _src2, _dest, _len) == 0
}

// ImageFilterMultNor filters using MultNor: D = S1 * S2.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#aff3256626208bfc490268cf07e8a29af)
func ImageFilterMultNor(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultNor(_src1, _src2, _dest, _len) == 0
}

// ImageFilterMultDivby2 filter using MultDivby2: D = saturation255(S1/2 * S2)
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a462b662e34e0ea7f1da83fb493f9d9f5)
func ImageFilterMultDivby2(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultDivby2(_src1, _src2, _dest, _len) == 0
}

// ImageFilterMultDivby4 filters using MultDivby4: D = saturation255(S1/2 * S2/2).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a52e4de0e4818b4256c189f35e68e1242)
func ImageFilterMultDivby4(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultDivby4(_src1, _src2, _dest, _len) == 0
}

// ImageFilterBitAnd filters using BitAnd: D = S1 & S2.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a47a7564f857e42dcc2e3b5f8cd2943a9)
func ImageFilterBitAnd(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitAnd(_src1, _src2, _dest, _len) == 0
}

// ImageFilterBitOr filters using BitOr: D = S1 | S2.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a7c6288c51dcf074b4ba8f1bf0c349f02)
func ImageFilterBitOr(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitOr(_src1, _src2, _dest, _len) == 0
}

// ImageFilterDiv filters using Div: D = S1 / S2.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a2944f525acc587ca8d701fbdf1a49c36)
func ImageFilterDiv(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterDiv(_src1, _src2, _dest, _len) == 0
}

// ImageFilterBitNegation filters using BitNegation: D = !S.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#ac11af558f478ec72eb2b61e8bdf43225)
func ImageFilterBitNegation(src1, dest []byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitNegation(_src1, _dest, _len) == 0
}

// ImageFilterAddByte filters using AddByte: D = saturation255(S + C).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#ad00178c9482a9959023a6bec03c8dba5)
func ImageFilterAddByte(src1, dest []byte, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uchar(c)
	return C.SDL_imageFilterAddByte(_src1, _dest, _len, _c) == 0
}

// ImageFilterAddUint filters using AddUint: D = saturation255((S[i] + Cs[i % 4]), Cs=Swap32((uint)C).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#ab7d7f266f047a63755a2341cdfe018e9)
func ImageFilterAddUint(src1, dest []byte, c uint) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uint(c)
	return C.SDL_imageFilterAddUint(_src1, _dest, _len, _c) == 0
}

// ImageFilterSubByte filters using SubByte: D = saturation0(S - C).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a5899423c538fa35660ded0f5945c014f)
func ImageFilterSubByte(src1, dest []byte, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uchar(c)
	return C.SDL_imageFilterSubByte(_src1, _dest, _len, _c) == 0
}

// ImageFilterSubUint filters using SubUint: D = saturation0(S[i] - Cs[i % 4]), Cs=Swap32((uint)C).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a8532da4511ef9657c8688f66e6309118)
func ImageFilterSubUint(src1, dest []byte, c uint) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uint(c)
	return C.SDL_imageFilterSubUint(_src1, _dest, _len, _c) == 0
}

// ImageFilterShiftRight filters using ShiftRight: D = saturation0(S >> N).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a1ed688eb128d71af36386e9853d001a9)
func ImageFilterShiftRight(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftRight(_src1, _dest, _len, _n) == 0
}

// ImageFilterShiftRightUint filters using ShiftRightUint: D = saturation0((uint)S[i] >> N).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a2e5ec075145b34c5ea797ffa70891e53)
func ImageFilterShiftRightUint(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftRightUint(_src1, _dest, _len, _n) == 0
}

// ImageFilterMultByByte filters using MultByByte: D = saturation255(S * C).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#aef668f157cc152554872ccac491ee2f7)
func ImageFilterMultByByte(src1, dest []byte, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uchar(c)
	return C.SDL_imageFilterMultByByte(_src1, _dest, _len, _c) == 0
}

// ImageFilterShiftRightAndMultByByte filters using ShiftRightAndMultByByte: D = saturation255((S >> N) * C).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#ad8d11768b921ba823d412166903340b8)
func ImageFilterShiftRightAndMultByByte(src1, dest []byte, n, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	_c := C.uchar(c)
	return C.SDL_imageFilterShiftRightAndMultByByte(_src1, _dest, _len, _n, _c) == 0
}

// ImageFilterShiftLeftByte filters using ShiftLeftByte: D = (S << N).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a656657c3f31effa01163532fd96b3011)
func ImageFilterShiftLeftByte(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeftByte(_src1, _dest, _len, _n) == 0
}

// ImageFilterShiftLeftUint filters using ShiftLeftUint: D = ((uint)S << N).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a3ea712cad49735ca672e1d2da1e68516)
func ImageFilterShiftLeftUint(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeftUint(_src1, _dest, _len, _n) == 0
}

// ImageFilterShiftLeft filters ShiftLeft: D = saturation255(S << N).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a29891456dee25b30c8da8f767d7545c5)
func ImageFilterShiftLeft(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeft(_src1, _dest, _len, _n) == 0
}

// ImageFilterBinarizeUsingThreshold filters using BinarizeUsingThreshold: D = (S >= T) ? 255:0.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a02d89f9fa47f1f5c2d969a9d86acb041)
func ImageFilterBinarizeUsingThreshold(src1, dest []byte, t byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_t := C.uchar(t)
	return C.SDL_imageFilterBinarizeUsingThreshold(_src1, _dest, _len, _t) == 0
}

// ImageFilterClipToRange filters using ClipToRange: D = (S >= Tmin) & (S <= Tmax) S:Tmin | Tmax.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#a46a5728f8857b0a06694828375527451)
func ImageFilterClipToRange(src1, dest []byte, tmin, tmax byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_tmin := C.uchar(tmin)
	_tmax := C.uchar(tmax)
	return C.SDL_imageFilterClipToRange(_src1, _dest, _len, _tmin, _tmax) == 0
}

// ImageFilterNormalizeLinear filters using NormalizeLinear: D = saturation255((Nmax - Nmin)/(Cmax - Cmin)*(S - Cmin) + Nmin).
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__image_filter_8h.html#ade0729be518dec0b26ec164ff4e63476)
func ImageFilterNormalizeLinear(src1, dest []byte, cmin, cmax, nmin, nmax int) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_cmin := C.int(cmin)
	_cmax := C.int(cmax)
	_nmin := C.int(nmin)
	_nmax := C.int(nmax)
	return C.SDL_imageFilterNormalizeLinear(_src1, _dest, _len, _cmin, _cmax, _nmin, _nmax) == 0
}

// RotoZoomSurface rotates and zooms a surface and optional anti-aliasing. Rotates and zoomes a 32bit or 8bit 'src' surface to newly created 'dst' surface. 'angle' is the rotation in degrees and 'zoom' a scaling factor. If 'smooth' is set then the destination 32bit surface is anti-aliased. If the surface is not 8bit or 32bit RGBA/ABGR it will be converted into a 32bit RGBA format on the fly.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__rotozoom_8h.html#a6f5f31a362f63370dc60049df14d6856)
func RotoZoomSurface(src *sdl.Surface, angle, zoom float64, smooth int) *sdl.Surface {
	_angle := C.double(angle)
	_zoom := C.double(zoom)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.rotozoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoom, _smooth)))
}

// RotoZoomSurfaceXY rotates and zooms a surface with different horizontal and vertival scaling factors and optional anti-aliasing. Rotates and zooms a 32bit or 8bit 'src' surface to newly created 'dst' surface. 'angle' is the rotation in degrees, 'zoomx and 'zoomy' scaling factors. If 'smooth' is set then the destination 32bit surface is anti-aliased. If the surface is not 8bit or 32bit RGBA/ABGR it will be converted into a 32bit RGBA format on the fly.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__rotozoom_8h.html#a592d84489ce544c050a9f3fe0e04f3f6)
func RotoZoomSurfaceXY(src *sdl.Surface, angle, zoomx, zoomy float64, smooth int) *sdl.Surface {
	_angle := C.double(angle)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.rotozoomSurfaceXY((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoomx, _zoomy, _smooth)))
}

// ZoomSurface zoom a surface by independent horizontal and vertical factors with optional smoothing. Zooms a 32bit or 8bit 'src' surface to newly created 'dst' surface. 'zoomx' and 'zoomy' are scaling factors for width and height. If 'smooth' is on then the destination 32bit surface is anti-aliased. If the surface is not 8bit or 32bit RGBA/ABGR it will be converted into a 32bit RGBA format on the fly. If zoom factors are negative, the image is flipped on the axes.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__rotozoom_8h.html#a0867857132421429994198cabacb0528)
func ZoomSurface(src *sdl.Surface, zoomx, zoomy float64, smooth int) *sdl.Surface {
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.zoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _zoomx, _zoomy, _smooth)))
}

// ZoomSurfaceSize calculates the size of the target surface for a gfx.ZoomSurface() call. The minimum size of the target surface is 1. The input factors can be positive or negative.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__rotozoom_8h.html#a87a121da75a099fd980295c759a7005d)
func ZoomSurfaceSize(width, height int32, zoomx, zoomy float64) (dstwidth, dstheight int) {
	_width := C.int(width)
	_height := C.int(height)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_dstwidth := (*C.int)(unsafe.Pointer(&dstwidth))
	_dstheight := (*C.int)(unsafe.Pointer(&dstheight))
	C.zoomSurfaceSize(_width, _height, _zoomx, _zoomy, _dstwidth, _dstheight)
	return dstwidth, dstheight
}

// ShrinkSurface shrink a surface by an integer ratio using averaging. Shrinks a 32bit or 8bit 'src' surface to a newly created 'dst' surface. 'factorx' and 'factory' are the shrinking ratios (i.e. 2=1/2 the size, 3=1/3 the size, etc.) The destination surface is antialiased by averaging the source box RGBA or Y information. If the surface is not 8bit or 32bit RGBA/ABGR it will be converted into a 32bit RGBA format on the fly. The input surface is not modified. The output surface is newly allocated.
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__rotozoom_8h.html#a9adfe732cbca348e3287096e7c67e72d)
func ShrinkSurface(src *sdl.Surface, factorx, factory int) *sdl.Surface {
	_factorx := C.int(factorx)
	_factory := C.int(factory)
	return (*sdl.Surface)(unsafe.Pointer(C.shrinkSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _factorx, _factory)))
}

// RotateSurface90Degrees rotates a 32 bit surface in increments of 90 degrees. Specialized 90 degree rotator which rotates a 'src' surface in 90 degree increments clockwise returning a new surface. Faster than rotozoomer since not scanning or interpolation takes place. Input surface must be 32 bit. (code contributed by J. Schiller, improved by C. Allport and A. Schiffler)
// (https://www.ferzkopp.net/Software/SDL2_gfx/Docs/html/_s_d_l2__rotozoom_8h.html#ac2858dec47549c8f82360568b5a29363)
func RotateSurface90Degrees(src *sdl.Surface, numClockwiseTurns int) *sdl.Surface {
	_numClockwiseTurns := C.int(numClockwiseTurns)
	return (*sdl.Surface)(unsafe.Pointer(C.rotateSurface90Degrees((*C.SDL_Surface)(unsafe.Pointer(src)), _numClockwiseTurns)))
}

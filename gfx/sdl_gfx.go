// Package gfx is an add-on for SDL2 that provides drawing of graphics primitives, rotozoomer, MMX image filters and framerate control.
package gfx

//#include <stdlib.h>
//#include "sdl_gfx_wrapper.h"
import "C"
import "github.com/veandco/go-sdl2/sdl"
import "unsafe"

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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__framerate_8c.html#a444ebaaaa6b1ceeafa921562bdab1a44)
func InitFramerate(manager *FPSmanager) {
	C.SDL_initFramerate(manager.cptr())
}

// SetFramerate sets the framerate in Hz.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__framerate_8c.html#afad4b503cf9719daced45fa4d9653d72)
func SetFramerate(manager *FPSmanager, rate uint32) bool {
	_rate := C.Uint32(rate)
	return C.SDL_setFramerate(manager.cptr(), _rate) == 0
}

// GetFramerate returns the current target framerate in Hz.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__framerate_8c.html#a575bb511d6f817ad846a788cbd08ae91)
func GetFramerate(manager *FPSmanager) (int, bool) {
	fps := int(C.SDL_getFramerate(manager.cptr()))
	return fps, fps >= 0
}

// GetFramecount returns the current framecount.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__framerate_8c.html#a96b13e26f8436222e866904a592a6eec)
func GetFramecount(manager *FPSmanager) (int, bool) {
	count := int(C.SDL_getFramecount(manager.cptr()))
	return count, count >= 0
}

// FramerateDelay delays execution to maintain a constant framerate and calculate fps.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__framerate_8c.html#afce13fa3dd37130deb4975d8b230c3ba)
func FramerateDelay(manager *FPSmanager) uint32 {
	return uint32(C.SDL_framerateDelay(manager.cptr()))
}

// PixelColor pixel draws with blending enabled if a<255.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#ae6f8690e5c5a85d3263c8e16727b34ef)
func PixelColor(renderer *sdl.Renderer, x, y int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.pixelColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _color) == 0
}

// PixelRGBA pixel draws with blending enabled if a<255.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a7b6f83bdef72f6b356664a93841381c0)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#ac211a904dce45093315e15b10c80d8ac)
func HlineColor(renderer *sdl.Renderer, x1, x2, y int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.hlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _color) == 0
}

// HlineRGBA draws horizontal line with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a6608a0d1d4c7e16fa1afcbd3eb5c3850)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a9b45060155a19fee24f998d7790f1d67)
func VlineColor(renderer *sdl.Renderer, x, y1, y2 int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.vlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _color) == 0
}

// VlineRGBA draws vertical line with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a8b79ac1e779755aee92b04f3a6cfc5d7)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a6ab25c393f6e5f8d68ea3365f6ea98d2)
func RectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.rectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// RectangleRGBA draws rectangle with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a40991c6eeb936d35d0a8e8aa95268f72)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a830dd9dcfa39f4718aa2c269060326d0)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a300272b3b799f09ca6cd5c541b19f07a)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a6bb30dfc32d0aee20271a0356a2e2fd0)
func BoxColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.boxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// BoxRGBA draws box (filled rectangle) with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a1864b3062793a7f7dd81aaf8c8abd6b0)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a718c4f31d1e145106959c2a77d5fee9d)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#aad706348fec18631d7bc48a2d91f5b4d)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#ad44c550fab3cb736eb049713ede94052)
func LineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.lineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// LineRGBA draws line with alpha blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a760139e11a9ae5defeb755ca0c794f5f)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a25c56f2def855db01dcf7ff7f7356182)
func AALineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int32, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.aalineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

// AALineRGBA draws anti-aliased line with alpha blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a25c56f2def855db01dcf7ff7f7356182)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a1494109358b4e4b7ec300d83e3f90300)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a8b24d64b51e23592c93abc2aa50c818e)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#aa99bd361cc947b448142720f2ca3320e)
func CircleColor(renderer *sdl.Renderer, x, y, rad int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.circleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

// CircleRGBA draws circle with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a7fe51d4c9426c8795e58c7ddd313b0a4)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a461b8ac31e00306aee5f8a4c242671d2)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a2aff993d0d8d64564e16145f401d3cf1)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#aad64361b01181e6aff940add96d23c61)
func AACircleColor(renderer *sdl.Renderer, x, y, rad int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.aacircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

// AACircleRGBA draws anti-aliased circle with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a332780885aa2cfdc2de34dcff8d67e8b)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a39147d1282ec814a1b9e31243aad0359)
func FilledCircleColor(renderer *sdl.Renderer, x, y, rad int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.filledCircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

// FilledCircleRGBA draws filled circle with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a562ba6b18fb70547cd50cb3bb0f70272)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a476cff7702f4be9090871e35859782f0)
func EllipseColor(renderer *sdl.Renderer, x, y, rx, ry int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.ellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

// EllipseRGBA draws ellipse with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a18c8a26c9009482eec40f9f4a6945fd1)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a1c7d20dcba8e0d7ce483f4c854c438be)
func AAEllipseColor(renderer *sdl.Renderer, x, y, rx, ry int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.aaellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

// AAEllipseRGBA draws anti-aliased ellipse with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#ab9f0f00d7fb2f04aa9ba1630e31a27bf)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a8fed50800f2f1bdfaa048698f5052f25)
func FilledEllipseColor(renderer *sdl.Renderer, x, y, rx, ry int32, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.filledEllipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

// FilledEllipseRGBA draws filled ellipse with blending.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a33595ad996dd0dcccde3abbcef540eec)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a3c2bc64deabda74933f31daba6bed7be)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a8442f2c2bedbe27c96d8d44319981992)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a2c30ee985b2513dc58d9b19d4e71562b)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a4ffdfd2834f3ef0fd0ee622b5f1d16b8)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a7465d08ef930ebb5442c7dd246fed4b5)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a45d6a7edcd8b25e1a60e39b7f60bda3f)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a78d4ed2372527f3b78f5893928b0f519)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a8f318d776ff1e3c6790405e0e59e5356)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a2d692dc25f3b579b386dff8dcd9cbc00)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#ae55541ec58990420dc6dc6b9d61f33d6)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a09950a50e8806e88bb20c543c58cc6a8)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a7d08522e52d8290c5c498ce435fa51f0)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#af22692175cb73329410cbcc7d7491c4d)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a40ef0b898905c190c193f0f55deb5a6c)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a65137af308ea878f28abc95419e8aef5)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#adfe8f9c42d29a090aae15eeb19b80d51)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a4b7fbf6cc366abdf345a25308d53e125)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#afacd57651ec0e0ccab60753636862cd0)
func SetFont(fontdata []byte, cw, ch uint32) {
	_fontdata := unsafe.Pointer(nil)
	if fontdata != nil {
		_fontdata = unsafe.Pointer(&fontdata[0])
	}
	_cw := C.Uint32(cw)
	_ch := C.Uint32(ch)
	C.gfxPrimitivesSetFont(_fontdata, _cw, _ch)
}

// SetFontRotation sets current global font character rotation steps. Default is 0 (no rotation). 1 = 90deg clockwise. 2 = 180deg clockwise. 3 = 270deg clockwise. Changing the rotation, will reset the character cache.(http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#aef6796a883f07d31bbf7c7df6d1153d2)
func SetFontRotation(rotation uint32) {
	_rotation := C.Uint32(rotation)
	C.gfxPrimitivesSetFontRotation(_rotation)
}

// CharacterColor draws a character of the currently set font. On first call for a particular character and color combination, the function needs to generate the character surface (slower). Subsequent calls blit a cached surface (fast). Uses alpha blending if A<255 in color.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#aef5fdeb16c4578d8cd50e106299e993e)
func CharacterColor(renderer *sdl.Renderer, x, y int32, c byte, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_color := C.Uint32(gfxColor(color))
	return C.characterColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _color) == 0
}

// CharacterRGBA draws a character of the currently set font.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a96379d2ce808aa642afb57bced0c670e)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a62d2ba55abc7673f2dfa29e6bbffefdf)
func StringColor(renderer *sdl.Renderer, x, y int32, s string, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_color := C.Uint32(gfxColor(color))
	return C.stringColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _color) == 0
}

// StringRGBA draws a string in the currently set font.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__gfx_primitives_8c.html#a6ca71826e311bdd9acf13b009256aa1c)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a798ce71024ee1a1d1b174fd60fe79917)
func ImageFilterMMXdetect() bool {
	return C.SDL_imageFilterMMXdetect() > 0
}

// ImageFilterMMXoff disables MMX check for filter functions and force to use non-MMX C based code.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a5dff661660755161bb4aaf6199cd1384)
func ImageFilterMMXoff() {
	C.SDL_imageFilterMMXoff()
}

// ImageFilterMMXon enables MMX check for filter functions and use MMX code if available.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a353ee234c3b51b33c4c5c4b30db5832d)
func ImageFilterMMXon() {
	C.SDL_imageFilterMMXon()
}

// ImageFilterAdd filters using Add: D = saturation255(S1 + S2).
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a9f06507eb0b63198dbd67495d61c9b20)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#ace072118fef77973210eb04fb4bfc779)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a3c01cf8576ea7a0dfc09dbaa953c9287)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a472909f904274255cd6793c520172e48)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#af4633031d40a9ea0956a2f3c6c87a384)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a5f3c9fd40426bb46eba5ac167505dcc5)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a80737f6427c7bdb30d39a92f6524fc14)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a30e685653eb1050c7d48feaeb8f801a1)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a85837ce1b5de1f907b6b9053922b5cbc)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a5cf1c477f4e32d02f74ee95d9f7b0021)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a0ea22f01c6a4724bac307da3e5355f58)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#ac3abfaa8ec2e88c3c4893588c5555856)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a812cb307cb60ef31f1ffe81a9eee6bb1)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a660543426c47dfec39a349eb3b8f905b)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a387fb6f0d48cc5d08f37f7f9b92d14b2)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#abb343ef95e22945e1d4d648b2e176e64)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a68851aed2dcc5dfd2f3b258236f3b88c)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a540d4625d76bcd03318c2a59ce650fdb)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a06f7a19d6e2fc89d7b48cc45d715806d)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a0713d6c267fba9756d6beae81e89f9e4)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a4561a73b249a26babc4c469ffbdae604)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a250e796fb2db470da0a78b74b78114e8)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a98372fea76310903abef7808db10d226)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#a951a062e15df290a137428e1e0f4d5ce)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#ab7224abc4ecc1b8a6f4441ef8379515f)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__image_filter_8c.html#ab018ace4db884cac953b06b09c00828b)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__rotozoom_8c.html#a5f64ed53eeee5f2667971c857698d1e5)
func RotoZoomSurface(src *sdl.Surface, angle, zoom float64, smooth int) *sdl.Surface {
	_angle := C.double(angle)
	_zoom := C.double(zoom)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.rotozoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoom, _smooth)))
}

// RotoZoomSurfaceXY rotates and zooms a surface with different horizontal and vertival scaling factors and optional anti-aliasing. Rotates and zooms a 32bit or 8bit 'src' surface to newly created 'dst' surface. 'angle' is the rotation in degrees, 'zoomx and 'zoomy' scaling factors. If 'smooth' is set then the destination 32bit surface is anti-aliased. If the surface is not 8bit or 32bit RGBA/ABGR it will be converted into a 32bit RGBA format on the fly.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__rotozoom_8c.html#aab98b5b0da4ea468bacf47f7b85f0ee2)
func RotoZoomSurfaceXY(src *sdl.Surface, angle, zoomx, zoomy float64, smooth int) *sdl.Surface {
	_angle := C.double(angle)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.rotozoomSurfaceXY((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoomx, _zoomy, _smooth)))
}

// ZoomSurface zoom a surface by independent horizontal and vertical factors with optional smoothing. Zooms a 32bit or 8bit 'src' surface to newly created 'dst' surface. 'zoomx' and 'zoomy' are scaling factors for width and height. If 'smooth' is on then the destination 32bit surface is anti-aliased. If the surface is not 8bit or 32bit RGBA/ABGR it will be converted into a 32bit RGBA format on the fly. If zoom factors are negative, the image is flipped on the axes.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__rotozoom_8c.html#abdd772b2f6b1f26134e4e90cda657a21)
func ZoomSurface(src *sdl.Surface, zoomx, zoomy float64, smooth int) *sdl.Surface {
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.zoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _zoomx, _zoomy, _smooth)))
}

// ZoomSurfaceSize calculates the size of the target surface for a gfx.ZoomSurface() call. The minimum size of the target surface is 1. The input factors can be positive or negative.
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__rotozoom_8c.html#a8ba40859c1a977dae87488dd8be1bf9a)
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
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__rotozoom_8c.html#aad3bf0cd89cc39ff874ffa778fa1495d)
func ShrinkSurface(src *sdl.Surface, factorx, factory int) *sdl.Surface {
	_factorx := C.int(factorx)
	_factory := C.int(factory)
	return (*sdl.Surface)(unsafe.Pointer(C.shrinkSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _factorx, _factory)))
}

// RotateSurface90Degrees rotates a 32 bit surface in increments of 90 degrees. Specialized 90 degree rotator which rotates a 'src' surface in 90 degree increments clockwise returning a new surface. Faster than rotozoomer since not scanning or interpolation takes place. Input surface must be 32 bit. (code contributed by J. Schiller, improved by C. Allport and A. Schiffler)
// (http://www.ferzkopp.net/Software/SDL_gfx-2.0/Docs/html/_s_d_l__rotozoom_8c.html#a77563d68634cb2624d4f2f0bcdc19e73)
func RotateSurface90Degrees(src *sdl.Surface, numClockwiseTurns int) *sdl.Surface {
	_numClockwiseTurns := C.int(numClockwiseTurns)
	return (*sdl.Surface)(unsafe.Pointer(C.rotateSurface90Degrees((*C.SDL_Surface)(unsafe.Pointer(src)), _numClockwiseTurns)))
}

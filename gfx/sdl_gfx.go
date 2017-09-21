package gfx

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_gfx
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gfx
//#include <stdlib.h>
//#include "sdl_gfx_wrapper.h"
import "C"
import "github.com/veandco/go-sdl2/sdl"
import "unsafe"

// FPS
const (
	FPS_UPPER_LIMIT = 200
	FPS_LOWER_LIMIT = 1
	FPS_DEFAULT     = 30
)

type FPSmanager struct {
	FrameCount uint32
	RateTicks  float32
	BaseTicks  uint32
	LastTicks  uint32
	Rate       uint32
}

const (
	SMOOTHING_OFF = 0
	SMOOTHING_ON  = 1
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

func InitFramerate(manager *FPSmanager) {
	C.SDL_initFramerate(manager.cptr())
}

func SetFramerate(manager *FPSmanager, rate uint32) bool {
	_rate := C.Uint32(rate)
	return C.SDL_setFramerate(manager.cptr(), _rate) == 0
}

func GetFramerate(manager *FPSmanager) (int, bool) {
	fps := int(C.SDL_getFramerate(manager.cptr()))
	return fps, fps >= 0
}

func GetFramecount(manager *FPSmanager) (int, bool) {
	count := int(C.SDL_getFramecount(manager.cptr()))
	return count, count >= 0
}

func FramerateDelay(manager *FPSmanager) uint32 {
	return uint32(C.SDL_framerateDelay(manager.cptr()))
}

func PixelColor(renderer *sdl.Renderer, x, y int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.pixelColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _color) == 0
}

func PixelRGBA(renderer *sdl.Renderer, x, y int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.pixelRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _r, _g, _b, _a) == 0
}

func HlineColor(renderer *sdl.Renderer, x1, x2, y int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.hlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _color) == 0
}

func HlineRGBA(renderer *sdl.Renderer, x1, x2, y int, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.hlineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _r, _g, _b, _a) == 0
}

func VlineColor(renderer *sdl.Renderer, x, y1, y2 int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.vlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _color) == 0
}

func VlineRGBA(renderer *sdl.Renderer, x, y1, y2 int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.vlineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _r, _g, _b, _a) == 0
}

func RectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.rectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func RectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) bool {
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

func RoundedRectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.roundedRectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _color) == 0
}

func RoundedRectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int, r, g, b, a uint8) bool {
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

func BoxColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.boxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func BoxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) bool {
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

func RoundedBoxColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.roundedBoxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _color) == 0
}

func RoundedBoxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, rad int, r, g, b, a uint8) bool {
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

func LineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.lineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func LineRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) bool {
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

func AALineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.aalineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func AALineRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) bool {
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

func ThickLineColor(renderer *sdl.Renderer, x1, y1, x2, y2, width int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_width := C.Uint8(width)
	_color := C.Uint32(gfxColor(color))
	return C.thickLineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _width, _color) == 0
}

func ThickLineRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, width int, r, g, b, a uint8) bool {
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

func CircleColor(renderer *sdl.Renderer, x, y, rad int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.circleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

func CircleRGBA(renderer *sdl.Renderer, x, y, rad int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.circleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

func ArcColor(renderer *sdl.Renderer, x, y, rad, start, end int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.arcColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

func ArcRGBA(renderer *sdl.Renderer, x, y, rad, start, end int, r, g, b, a uint8) bool {
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

func AACircleColor(renderer *sdl.Renderer, x, y, rad int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.aacircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

func AACircleRGBA(renderer *sdl.Renderer, x, y, rad int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aacircleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

func FilledCircleColor(renderer *sdl.Renderer, x, y, rad int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.filledCircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

func FilledCircleRGBA(renderer *sdl.Renderer, x, y, rad int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledCircleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

func EllipseColor(renderer *sdl.Renderer, x, y, rx, ry int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.ellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

func EllipseRGBA(renderer *sdl.Renderer, x, y, rx, ry int, r, g, b, a uint8) bool {
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

func AAEllipseColor(renderer *sdl.Renderer, x, y, rx, ry int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.aaellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

func AAEllipseRGBA(renderer *sdl.Renderer, x, y, rx, ry int, r, g, b, a uint8) bool {
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

func FilledEllipseColor(renderer *sdl.Renderer, x, y, rx, ry int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.filledEllipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

func FilledEllipseRGBA(renderer *sdl.Renderer, x, y, rx, ry int, r, g, b, a uint8) bool {
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

func PieColor(renderer *sdl.Renderer, x, y, rad, start, end int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.pieColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

func PieRGBA(renderer *sdl.Renderer, x, y, rad, start, end int, r, g, b, a uint8) bool {
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

func FilledPieColor(renderer *sdl.Renderer, x, y, rad, start, end int, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.filledPieColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

func FilledPieRGBA(renderer *sdl.Renderer, x, y, rad, start, end int, r, g, b, a uint8) bool {
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

func TrigonColor(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(gfxColor(color))
	return C.trigonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
}

func TrigonRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int, r, g, b, a uint8) bool {
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

func FilledTrigonColor(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int, color sdl.Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(gfxColor(color))
	return C.filledTrigonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
}

func FilledTrigonRGBA(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int, r, g, b, a uint8) bool {
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

func PolygonColor(renderer *sdl.Renderer, vx, vy []int16, color sdl.Color) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_color := C.Uint32(gfxColor(color))
	return C.polygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

func PolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.polygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func AAPolygonColor(renderer *sdl.Renderer, vx, vy []int16, color sdl.Color) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_color := C.Uint32(gfxColor(color))
	return C.aapolygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

func AAPolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aapolygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func FilledPolygonColor(renderer *sdl.Renderer, vx, vy []int16, color sdl.Color) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_color := C.Uint32(gfxColor(color))
	return C.filledPolygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

func FilledPolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledPolygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func TexturedPolygon(renderer *sdl.Renderer, vx, vy []int16, surface *sdl.Surface, textureDX, textureDY int) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_len := C.int(min(len(vx), len(vy)))
	_textureDX := C.int(textureDX)
	_textureDY := C.int(textureDY)
	return C.texturedPolygon((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _surface, _textureDX, _textureDY) == 0
}

func BezierColor(renderer *sdl.Renderer, vx, vy []int16, s int, color sdl.Color) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_s := C.int(s)
	_color := C.Uint32(gfxColor(color))
	return C.bezierColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _s, _color) == 0
}

func BezierRGBA(renderer *sdl.Renderer, vx, vy []int16, s int, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(min(len(vx), len(vy)))
	_s := C.int(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.bezierRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _s, _r, _g, _b, _a) == 0
}

func SetFont(fontdata []byte, cw, ch uint32) {
	_fontdata := unsafe.Pointer(nil)
	if fontdata != nil {
		_fontdata = unsafe.Pointer(&fontdata[0])
	}
	_cw := C.Uint32(cw)
	_ch := C.Uint32(ch)
	C.gfxPrimitivesSetFont(_fontdata, _cw, _ch)
}

func SetFontRotation(rotation uint32) {
	_rotation := C.Uint32(rotation)
	C.gfxPrimitivesSetFontRotation(_rotation)
}

func CharacterColor(renderer *sdl.Renderer, x, y int, c byte, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_color := C.Uint32(gfxColor(color))
	return C.characterColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _color) == 0
}

func CharacterRGBA(renderer *sdl.Renderer, x, y int, c, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.characterRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _r, _g, _b, _a) == 0
}

func StringColor(renderer *sdl.Renderer, x, y int, s string, color sdl.Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_color := C.Uint32(gfxColor(color))
	return C.stringColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _color) == 0
}

func StringRGBA(renderer *sdl.Renderer, x, y int, s string, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.stringRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _r, _g, _b, _a) == 0
}

func ImageFilterMMXdetect() bool {
	return C.SDL_imageFilterMMXdetect() > 0
}

func ImageFilterMMXoff() {
	C.SDL_imageFilterMMXoff()
}

func ImageFilterMMXon() {
	C.SDL_imageFilterMMXon()
}

func ImageFilterAdd(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterAdd(_src1, _src2, _dest, _len) == 0
}

func ImageFilterMean(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterMean(_src1, _src2, _dest, _len) == 0
}

func ImageFilterSub(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterSub(_src1, _src2, _dest, _len) == 0
}

func ImageFilterAbsDiff(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterAbsDiff(_src1, _src2, _dest, _len) == 0
}

func ImageFilterMult(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterMult(_src1, _src2, _dest, _len) == 0
}

func ImageFilterMultNor(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterMultNor(_src1, _src2, _dest, _len) == 0
}

func ImageFilterMultDivby2(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterMultDivby2(_src1, _src2, _dest, _len) == 0
}

func ImageFilterMultDivby4(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterMultDivby4(_src1, _src2, _dest, _len) == 0
}

func ImageFilterBitAnd(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterBitAnd(_src1, _src2, _dest, _len) == 0
}

func ImageFilterBitOr(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterBitOr(_src1, _src2, _dest, _len) == 0
}

func ImageFilterDiv(src1, src2, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	return C.SDL_imageFilterDiv(_src1, _src2, _dest, _len) == 0
}

func ImageFilterBitNegation(src1, dest []byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	return C.SDL_imageFilterBitNegation(_src1, _dest, _len) == 0
}

func ImageFilterAddByte(src1, dest []byte, c byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_c := C.uchar(c)
	return C.SDL_imageFilterAddByte(_src1, _dest, _len, _c) == 0
}

func ImageFilterAddUint(src1, dest []byte, c uint) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_c := C.uint(c)
	return C.SDL_imageFilterAddUint(_src1, _dest, _len, _c) == 0
}

func ImageFilterSubByte(src1, dest []byte, c byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_c := C.uchar(c)
	return C.SDL_imageFilterSubByte(_src1, _dest, _len, _c) == 0
}

func ImageFilterSubUint(src1, dest []byte, c uint) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_c := C.uint(c)
	return C.SDL_imageFilterSubUint(_src1, _dest, _len, _c) == 0
}

func ImageFilterShiftRight(src1, dest []byte, n byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftRight(_src1, _dest, _len, _n) == 0
}

func ImageFilterShiftRightUint(src1, dest []byte, n byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftRightUint(_src1, _dest, _len, _n) == 0
}

func ImageFilterMultByByte(src1, dest []byte, c byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_c := C.uchar(c)
	return C.SDL_imageFilterMultByByte(_src1, _dest, _len, _c) == 0
}

func ImageFilterShiftRightAndMultByByte(src1, dest []byte, n, c byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_n := C.uchar(n)
	_c := C.uchar(c)
	return C.SDL_imageFilterShiftRightAndMultByByte(_src1, _dest, _len, _n, _c) == 0
}

func ImageFilterShiftLeftByte(src1, dest []byte, n byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeftByte(_src1, _dest, _len, _n) == 0
}

func ImageFilterShiftLeftUint(src1, dest []byte, n byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeftUint(_src1, _dest, _len, _n) == 0
}

func ImageFilterShiftLeft(src1, dest []byte, n byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeft(_src1, _dest, _len, _n) == 0
}

func ImageFilterBinarizeUsingThreshold(src1, dest []byte, t byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_t := C.uchar(t)
	return C.SDL_imageFilterBinarizeUsingThreshold(_src1, _dest, _len, _t) == 0
}

func ImageFilterClipToRange(src1, dest []byte, tmin, tmax byte) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_tmin := C.uchar(tmin)
	_tmax := C.uchar(tmax)
	return C.SDL_imageFilterClipToRange(_src1, _dest, _len, _tmin, _tmax) == 0
}

func ImageFilterNormalizeLinear(src1, dest []byte, cmin, cmax, nmin, nmax int) bool {
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_len := C.uint(min(len(src1), len(dest)))
	_cmin := C.int(cmin)
	_cmax := C.int(cmax)
	_nmin := C.int(nmin)
	_nmax := C.int(nmax)
	return C.SDL_imageFilterNormalizeLinear(_src1, _dest, _len, _cmin, _cmax, _nmin, _nmax) == 0
}

func RotoZoomSurface(src *sdl.Surface, angle, zoom float64, smooth int) *sdl.Surface {
	_angle := C.double(angle)
	_zoom := C.double(zoom)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.rotozoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoom, _smooth)))
}

func RotoZoomSurfaceXY(src *sdl.Surface, angle, zoomx, zoomy float64, smooth int) *sdl.Surface {
	_angle := C.double(angle)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.rotozoomSurfaceXY((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoomx, _zoomy, _smooth)))
}

func ZoomSurface(src *sdl.Surface, zoomx, zoomy float64, smooth int) *sdl.Surface {
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*sdl.Surface)(unsafe.Pointer(C.zoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _zoomx, _zoomy, _smooth)))
}

func ZoomSurfaceSize(width, height int, zoomx, zoomy float64) (dstwidth int, dstheight int) {
	_width := C.int(width)
	_height := C.int(height)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_dstwidth := (*C.int)(unsafe.Pointer(&dstwidth))
	_dstheight := (*C.int)(unsafe.Pointer(&dstheight))
	C.zoomSurfaceSize(_width, _height, _zoomx, _zoomy, _dstwidth, _dstheight)
	return dstwidth, dstheight
}

func ShrinkSurface(src *sdl.Surface, factorx, factory int) *sdl.Surface {
	_factorx := C.int(factorx)
	_factory := C.int(factory)
	return (*sdl.Surface)(unsafe.Pointer(C.shrinkSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _factorx, _factory)))
}

func RotateSurface90Degrees(src *sdl.Surface, numClockwiseTurns int) *sdl.Surface {
	_numClockwiseTurns := C.int(numClockwiseTurns)
	return (*sdl.Surface)(unsafe.Pointer(C.rotateSurface90Degrees((*C.SDL_Surface)(unsafe.Pointer(src)), _numClockwiseTurns)))
}

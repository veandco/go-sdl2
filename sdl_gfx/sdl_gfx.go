package gfx

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_gfx
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gfx
//#include <stdlib.h>
//#include "sdl_gfx_wrapper.h"
import "C"
import "github.com/veandco/go-sdl2/sdl"
import "unsafe"
//import "errors"

// FPS
const (
	FPS_UPPER_LIMIT = 200
	FPS_LOWER_LIMIT = 200
	FPS_DEFAULT     = 30
)

type FPSManager C.FPSmanager

func PixelColor(renderer *sdl.Renderer, x, y int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_color := C.Uint32(color)
	return C.pixelColor(renderer, _x, _y, _color) == 0
}

func PixelRGBA(renderer *sdl.Renderer, x, y int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.pixelRGBA(renderer, _x, _y, _r, _g, _b, _a) == 0
}

func HlineColor(renderer *sdl.Renderer, x1, x2, y int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_color := C.Uint32(color)
	return C.hlineColor(renderer, _x1, _x2, _y, _color) == 0
}

func HlineRGBA(renderer *sdl.Renderer, x1, x2, y int, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.hlineRGBA(renderer, _x1, _x2, _y, _r, _g, _b, _a) == 0
}

func VlineColor(renderer *sdl.Renderer, x, y1, y2 int, color uint32) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(color)
	return C.vlineColor(renderer, _x, _y1, _y2, _color) == 0
}

func VlineRGBA(renderer *sdl.Renderer, x, y1, y2 int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.vlineRGBA(renderer, _x, _y1, _y2, _r, _g, _b, _a) == 0
}

func RectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(color)
	return C.rectangleColor(renderer, _x1, _y1, _x2, _y2, _color) == 0
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
	return C.rectangleRGBA(renderer, _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func RoundedRectangleColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(color)
	return C.roundedRectangleColor(renderer, _x1, _y1, _x2, _y2, _rad, _color) == 0
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
	return C.roundedRectangleRGBA(renderer, _x1, _y1, _x2, _y2, _rad, _r, _g, _b, _a) == 0
}

func BoxColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(color)
	return C.boxColor(renderer, _x1, _y1, _x2, _y2, _color) == 0
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
	return C.boxRGBA(renderer, _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func RoundedBoxColor(renderer *sdl.Renderer, x1, y1, x2, y2, rad int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(color)
	return C.roundedBoxColor(renderer, _x1, _y1, _x2, _y2, _rad, _color) == 0
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
	return C.roundedBoxRGBA(renderer, _x1, _y1, _x2, _y2, _rad, _r, _g, _b, _a) == 0
}

func LineColor(renderer *sdl.Renderer, x1, y1, x2, y2 int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(color)
	return C.lineColor(renderer, _x1, _y1, _x2, _y2, _color) == 0
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
	return C.lineRGBA(renderer, _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func AALineColor(renderer *sdl.Renderer, x1, y1, x2, y2, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(color)
	return C.aalineColor(renderer, _x1, _y1, _x2, _y2, _color) == 0
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
	return C.aalineRGBA(renderer, _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func ThickLineColor(renderer *sdl.Renderer, x1, y1, x2, y2, width int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_width := C.Uint8(width)
	_color := C.Uint32(color)
	return C.thickLineColor(renderer, _x1, _y1, _x2, _y2, _width, _color) == 0
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
	return C.thickLineRGBA(renderer, _x1, _y1, _x2, _y2, _width, _r, _g, _b, _a) == 0
}

func CircleColor(renderer *sdl.Renderer, x, y, rad int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(color)
	return C.circleColor(renderer, _x, _y, _rad, _color) == 0
}

func CircleRGBA(renderer *sdl.Renderer, x, y, rad int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.circleRGBA(renderer, _x, _y, _rad, _r, _g, _b, _a) == 0
}

func ArcColor(renderer *sdl.Renderer, x, y, rad, start, end int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(color)
	return C.arcColor(renderer, _x, _y, _rad, _start, _end, _color) == 0
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
	return C.arcRGBA(renderer, _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

func AACircleColor(renderer *sdl.Renderer, x, y, rad int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(color)
	return C.aacircleColor(renderer, _x, _y, _rad, _color) == 0
}

func AACircleRGBA(renderer *sdl.Renderer, x, y, rad int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aacircleRGBA(renderer, _x, _y, _rad, _r, _g, _b, _a) == 0
}

func FilledCircleColor(renderer *sdl.Renderer, x, y, rad int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(color)
	return C.filledCircleColor(renderer, _x, _y, _rad, _color) == 0
}

func FilledCircleRGBA(renderer *sdl.Renderer, x, y, rad int, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledCircleRGBA(renderer, _x, _y, _rad, _r, _g, _b, _a) == 0
}

func EllipseColor(renderer *sdl.Renderer, x, y, rx, ry int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(color)
	return C.ellipseColor(renderer, _x, _y, _rx, _ry, _color) == 0
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
	return C.ellipseRGBA(renderer, _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

func AAEllipseColor(renderer *sdl.Renderer, x, y, rx, ry, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(color)
	return C.aaellipseColor(renderer, _x, _y, _rx, _ry, _color) == 0
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
	return C.aaellipseRGBA(renderer, _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

func FilledEllipseColor(renderer *sdl.Renderer, x, y, rx, ry, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(color)
	return C.filledEllipseColor(renderer, _x, _y, _rx, _ry, _color) == 0
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
	return C.filledEllipseRGBA(renderer, _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

func PieColor(renderer *sdl.Renderer, x, y, rad, start, end int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(color)
	return C.pieColor(renderer, _x, _y, _rad, _start, _end, _color) == 0
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
	return C.pieRGBA(renderer, _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

func FilledPieColor(renderer *sdl.Renderer, x, y, rad, start, end int, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(color)
	return C.filledPieColor(renderer, _x, _y, _rad, _start, _end, _color) == 0
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
	return C.filledPieRGBA(renderer, _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

func TrigonColor(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(color)
	return C.trigonColor(renderer, _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
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
	return C.trigonRGBA(renderer, _x1, _y1, _x2, _y2, _x3, _y3, _r, _g, _b, _a) == 0
}

func FilledTrigonColor(renderer *sdl.Renderer, x1, y1, x2, y2, x3, y3 int, color uint32) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(color)
	return C.filledTrigonColor(renderer, _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
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
	return C.filledTrigonRGBA(renderer, _x1, _y1, _x2, _y2, _x3, _y3, _r, _g, _b, _a) == 0
}

func PolygonColor(renderer *sdl.Renderer, vx, vy []int16, color uint32) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_color := C.Uint32(color)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.polygonColor(renderer, _vx, _vy, _len, _color) == 0
}

func PolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.polygonRGBA(renderer, _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func AAPolygonColor(renderer *sdl.Renderer, vx, vy []int16, color uint32) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_color := C.Uint32(color)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.aapolygonColor(renderer, _vx, _vy, _len, _color) == 0
}

func AAPolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.aapolygonRGBA(renderer, _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func FilledPolygonColor(renderer *sdl.Renderer, vx, vy []int16, color uint32) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_color := C.Uint32(color)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.filledPolygonColor(renderer, _vx, _vy, _len, _color) == 0
}

func FilledPolygonRGBA(renderer *sdl.Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.filledPolygonRGBA(renderer, _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func TexturedPolygon(renderer *sdl.Renderer, vx, vy []int16, surface *sdl.Surface, textureDX, textureDY int) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_len := C.int(0)
	_textureDX := C.int(textureDX)
	_textureDY := C.int(textureDY)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.texturedPolygon(renderer, _vx, _vy, _len, _surface, _textureDX, _textureDY) == 0
}

func BezierColor(renderer *sdl.Renderer, vx, vy []int16, s int, color uint32) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_s := C.int(s)
	_color := C.Uint32(color)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.bezierColor(renderer, _vx, _vy, _len, _s, _color) == 0
}

func BezierRGBA(renderer *sdl.Renderer, vx, vy []int16, s int, r, g, b, a uint8) bool {
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_len := C.int(0)
	_s := C.int(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)

	if len(vx) < len(vy) {
		_len = C.int(len(vx))
	} else {
		_len = C.int(len(vy))
	}

	return C.bezierRGBA(renderer, _vx, _vy, _len, _s, _r, _g, _b, _a) == 0
}

func GFXPrimitiveSetFont(fontdata []byte, cw, ch uint32) {
	_fontdata := unsafe.Pointer(&fontdata[0])
	_cw := C.Uint32(cw)
	_ch := C.Uint32(ch)
	C.gfxPrimitivesSetFont(_fontdata, _cw, _ch)
}

func GFXPrimitivesSetFontRotation(rotation uint32) {
	_rotation := C.Uint32(rotation)
	C.gfxPrimitivesSetFontRotation(_rotation)
}

func CharacterColor(renderer *sdl.Renderer, x, y, c byte, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_color := C.Uint32(color)
	return C.characterColor(renderer, _x, _y, _c, _color) == 0
}

func CharacterRGBA(renderer *sdl.Renderer, x, y, c, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.characterRGBA(renderer, _x, _y, _c, _r, _g, _b, _a) == 0
}

func StringColor(renderer *sdl.Renderer, x, y int, s string, color uint32) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_color := C.Uint32(color)
	return C.stringColor(renderer, _x, _y, _s, _color) == 0
}

func StringRGBA(renderer *sdl.Renderer, x, y int, s string, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.stringRGBA(renderer, _x, _y, _s, _r, _g, _b, _a) == 0
}

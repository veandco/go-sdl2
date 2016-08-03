package gfx

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_gfx
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gfx
//#include <stdlib.h>
//#include "sdl_gfx_wrapper.h"
import "C"
import "github.com/veandco/go-sdl2/sdl"
//import "unsafe"
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
